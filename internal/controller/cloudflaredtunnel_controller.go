/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	jensrotnecomv1alpha1 "github.com/jensrotne/cloudflared-operator/api/v1alpha1"
	"github.com/jensrotne/cloudflared-operator/internal/cloudflare"
)

// CloudflaredTunnelReconciler reconciles a CloudflaredTunnel object
type CloudflaredTunnelReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=jensrotne.com,resources=cloudflaredtunnels,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=jensrotne.com,resources=cloudflaredtunnels/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=jensrotne.com,resources=cloudflaredtunnels/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CloudflaredTunnel object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.2/pkg/reconcile
func (r *CloudflaredTunnelReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var tunnel jensrotnecomv1alpha1.CloudflaredTunnel

	if err := r.Get(ctx, req.NamespacedName, &tunnel); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, client.IgnoreNotFound(err)
		}

		log.Log.Error(err, "unable to fetch CloudflaredTunnel")
		return ctrl.Result{}, err
	}

	if err := r.handleFinalizer(ctx, tunnel); err != nil {
		log.Log.Error(err, "unable to handle finalizer")
		return ctrl.Result{}, err
	}

	if !tunnel.ObjectMeta.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, nil
	}

	cloudflareTunnel, err := GetOrCreateTunnel(&tunnel)

	if err != nil {
		log.Log.Error(err, "unable to get or create tunnel")
		return ctrl.Result{}, err
	}

	err = UpsertTunnelConfig(&tunnel, *cloudflareTunnel)

	if err != nil {
		log.Log.Error(err, "unable to upsert tunnel config")
		return ctrl.Result{}, err
	}

	err = UpsertTunnelDNSRecord(&tunnel, *cloudflareTunnel)

	if err != nil {
		log.Log.Error(err, "unable to upsert tunnel DNS record")
		return ctrl.Result{}, err
	}

	secret, err := GetOrCreateTunnelTokenSecret(ctx, r, &tunnel, *cloudflareTunnel)

	if err != nil {
		log.Log.Error(err, "unable to get or create tunnel token secret")
		return ctrl.Result{}, err
	}

	_, err = GetOrCreateDeployment(ctx, r, &tunnel, *cloudflareTunnel, secret)

	if err != nil {
		log.Log.Error(err, "unable to get or create deployment")
		return ctrl.Result{}, err
	}

	err = SetStatus(ctx, r, &tunnel, *cloudflareTunnel)

	if err != nil {
		log.Log.Error(err, "unable to set status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CloudflaredTunnelReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(controller.Options{MaxConcurrentReconciles: 10}).
		For(&jensrotnecomv1alpha1.CloudflaredTunnel{}).
		Complete(r)
}

func GetOrCreateTunnel(t *jensrotnecomv1alpha1.CloudflaredTunnel) (*cloudflare.CloudflareTunnel, error) {
	// Check if CRD status has tunnel ID
	if t.Status.TunnelID != "" {
		// Get tunnel from Cloudflare
		getTunnelResponse, err := cloudflare.GetTunnel(t.Status.TunnelID)

		if err != nil {
			return nil, err
		}

		if getTunnelResponse.Success {
			return &getTunnelResponse.Result, nil
		}
	}

	// List tunnels
	listOptions := map[string]string{
		"name":       t.Name,
		"is_deleted": "false",
	}

	listTunnelsResponse, err := cloudflare.ListTunnels(listOptions)

	if err != nil {
		return nil, err
	}

	if len(listTunnelsResponse.Result) == 0 {
		// Create tunnel
		createTunnelResponse, err := cloudflare.CreateTunnel(t.Name, "cloudflare", nil)

		if err != nil {
			return nil, err
		}

		if createTunnelResponse.Success {
			return &createTunnelResponse.Result, nil
		}
	}

	return &listTunnelsResponse.Result[0], nil
}

func UpsertTunnelConfig(t *jensrotnecomv1alpha1.CloudflaredTunnel, tunnel cloudflare.CloudflareTunnel) error {
	var config cloudflare.TunnelConfig
	changed := false

	// Get tunnel config
	getTunnelConfigResponse, err := tunnel.GetTunnelConfig()

	if err != nil {
		return err
	}

	if !getTunnelConfigResponse.Success {
		return fmt.Errorf("unable to get tunnel config")
	}

	if getTunnelConfigResponse.Result.Config == nil {
		config = cloudflare.TunnelConfig{}
	} else {
		config = *getTunnelConfigResponse.Result.Config
	}

	hostName := fmt.Sprintf("%s.%s", tunnel.ID, t.Spec.HostName)
	service := fmt.Sprintf("http://%s:%d", t.Spec.TargetService, t.Spec.TargetPort)

	if config.Ingress == nil {
		config.Ingress = []cloudflare.TunnelConfigIngress{
			{
				Hostname: &hostName,
				Service:  service,
			},
			{
				Service: "http_status:404",
			},
		}

		changed = true
	} else {
		if len(config.Ingress) < 2 && config.Ingress[0].Service == "http_status:404" {

			// Prepend ingress
			config.Ingress = append([]cloudflare.TunnelConfigIngress{
				{
					Hostname: &hostName,
					Service:  service,
				},
			}, config.Ingress...)

			changed = true
		} else if len(config.Ingress) < 2 && config.Ingress[0].Service != "http_status:404" {
			// Check if ingress is correct

			if *config.Ingress[0].Hostname != hostName {
				config.Ingress[0].Hostname = &hostName
			}

			if config.Ingress[0].Service != hostName {
				config.Ingress[0].Service = service
			}

			config.Ingress = append(config.Ingress, cloudflare.TunnelConfigIngress{
				Service: "http_status:404",
			})

			changed = true
		} else {
			// Check if ingress is correct
			if *config.Ingress[0].Hostname != hostName {
				config.Ingress[0].Hostname = &hostName
				changed = true
			}

			if config.Ingress[0].Service != service {
				config.Ingress[0].Service = service
				changed = true
			}
		}
	}

	if changed {
		// Put tunnel config
		putTunnelConfigResponse, err := tunnel.PutTunnelConfig(config)

		if err != nil {
			return err
		}

		if !putTunnelConfigResponse.Success {
			return fmt.Errorf("unable to put tunnel config")
		}
	}

	return nil
}

func UpsertTunnelDNSRecord(t *jensrotnecomv1alpha1.CloudflaredTunnel, tunnel cloudflare.CloudflareTunnel) error {
	// Check if DNS record exists
	record, err := cloudflare.GetDNSRecordIfExists(fmt.Sprintf("%s.%s", tunnel.ID, t.Spec.HostName))

	if err != nil {
		return err
	}

	if record == nil {
		tunnelDNS := fmt.Sprintf("%s.cfargotunnel.com", tunnel.ID)

		// Create DNS record
		_, err := cloudflare.CreateDNSCNAMERecord(tunnel.ID, tunnelDNS)

		if err != nil {
			return err
		}
	}

	return nil
}

func GetOrCreateTunnelTokenSecret(ctx context.Context, r *CloudflaredTunnelReconciler, t *jensrotnecomv1alpha1.CloudflaredTunnel, tunnel cloudflare.CloudflareTunnel) (*core.Secret, error) {
	// Get tunnel token
	tunnelTokenResponse, err := tunnel.GetTunnelToken()

	if err != nil {
		return nil, err
	}

	if !tunnelTokenResponse.Success {
		return nil, fmt.Errorf("unable to get tunnel token")
	}

	var secret core.Secret

	secretName := fmt.Sprintf("%s-tunnel-secret", tunnel.Name)

	err = r.Get(ctx, client.ObjectKey{Namespace: t.Namespace, Name: secretName}, &secret)

	if apierrors.IsNotFound(err) {
		secret = core.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: t.Namespace,
				Annotations: map[string]string{
					"tunnel.jensrotne.com/owner": t.Name,
				},
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: t.APIVersion,
						Kind:       t.Kind,
						Name:       t.Name,
						UID:        t.UID,
					},
				},
			},
			StringData: map[string]string{
				"secret": tunnelTokenResponse.Result,
			},
		}

		if err := r.Create(ctx, &secret); err != nil {
			return nil, err
		}
	} else {
		dataString := string(secret.Data["secret"])
		if dataString != tunnelTokenResponse.Result {
			secret.Data = map[string][]byte{
				"secret": []byte(tunnelTokenResponse.Result),
			}

			if err := r.Update(ctx, &secret); err != nil {
				return nil, err
			}
		}
	}

	return &secret, nil
}

func GetOrCreateDeployment(ctx context.Context, r *CloudflaredTunnelReconciler, t *jensrotnecomv1alpha1.CloudflaredTunnel, tunnel cloudflare.CloudflareTunnel, secret *core.Secret) (*appsv1.Deployment, error) {
	var deployment appsv1.Deployment

	err := r.Get(ctx, client.ObjectKey{Namespace: t.Namespace, Name: t.Name}, &deployment)

	if apierrors.IsNotFound(err) {
		deployment = appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      t.Name,
				Namespace: t.Namespace,
				Annotations: map[string]string{
					"tunnel.jensrotne.com/tunnel-id": tunnel.ID,
					"tunnel.jensrotne.com/owner":     t.Name,
				},
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: t.APIVersion,
						Kind:       t.Kind,
						Name:       t.Name,
						UID:        t.UID,
					},
				},
			},
			Spec: appsv1.DeploymentSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": t.Name,
					},
				},
				Template: core.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": t.Name,
						},
					},
					Spec: core.PodSpec{
						Containers: []core.Container{
							{
								Name:  "cloudflared",
								Image: "cloudflare/cloudflared:latest",
								Command: []string{
									"cloudflared",
								},
								Args: []string{
									"tunnel",
									"--metrics",
									"0.0.0.0:2000",
									"run",
									tunnel.ID,
								},
								Env: []core.EnvVar{
									{
										Name: "TUNNEL_TOKEN",
										ValueFrom: &core.EnvVarSource{
											SecretKeyRef: &core.SecretKeySelector{
												LocalObjectReference: core.LocalObjectReference{
													Name: secret.Name,
												},
												Key: "secret",
											},
										},
									},
								},
								LivenessProbe: &core.Probe{
									ProbeHandler: core.ProbeHandler{
										HTTPGet: &core.HTTPGetAction{
											Path: "/ready",
											Port: intstr.FromInt(2000),
										},
									},
									FailureThreshold:    1,
									PeriodSeconds:       10,
									InitialDelaySeconds: 10,
								},
							},
						},
					},
				},
			},
		}

		if err := r.Create(ctx, &deployment); err != nil {
			return nil, err
		}
	}

	return &deployment, nil
}

func SetStatus(ctx context.Context, r *CloudflaredTunnelReconciler, t *jensrotnecomv1alpha1.CloudflaredTunnel, tunnel cloudflare.CloudflareTunnel) error {
	// Update status
	t.Status.TunnelID = tunnel.ID

	if err := r.Status().Update(ctx, t); err != nil {
		return err
	}

	return nil
}

func CleanUpOwnedResources(ctx context.Context, r *CloudflaredTunnelReconciler, t *jensrotnecomv1alpha1.CloudflaredTunnel, tunnel cloudflare.CloudflareTunnel) error {
	// Delete deployment if it exists

	var deployment appsv1.Deployment

	err := r.Get(ctx, client.ObjectKey{Namespace: t.Namespace, Name: t.Name}, &deployment)

	if err == nil {
		if err := r.Delete(ctx, &deployment); err != nil {
			return err
		}
	}

	// Delete secret if it exists

	var secret core.Secret

	secretName := fmt.Sprintf("%s-tunnel-secret", tunnel.Name)

	err = r.Get(ctx, client.ObjectKey{Namespace: t.Namespace, Name: secretName}, &secret)

	if err == nil {
		if err := r.Delete(ctx, &secret); err != nil {
			return err
		}
	}

	// Delete DNS record if it exists

	record, err := cloudflare.GetDNSRecordIfExists(fmt.Sprintf("%s.%s", tunnel.ID, t.Spec.HostName))

	if err != nil {
		return err
	}

	if record != nil {
		if err := cloudflare.DeleteDNSRecord(record.ID); err != nil {
			return err
		}
	}

	// Delete tunnel

	_, err = cloudflare.DeleteTunnel(tunnel.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *CloudflaredTunnelReconciler) handleFinalizer(ctx context.Context, tunnel jensrotnecomv1alpha1.CloudflaredTunnel) error {
	finalizerName := "tunnel.jensrotne.com/finalizer"

	if tunnel.ObjectMeta.DeletionTimestamp.IsZero() {
		if !controllerutil.ContainsFinalizer(&tunnel, finalizerName) {
			controllerutil.AddFinalizer(&tunnel, finalizerName)

			return r.Update(ctx, &tunnel)
		}
	} else {
		if controllerutil.ContainsFinalizer(&tunnel, finalizerName) {
			if err := CleanUpOwnedResources(ctx, r, &tunnel, cloudflare.CloudflareTunnel{ID: tunnel.Status.TunnelID}); err != nil {
				return err
			}

			controllerutil.RemoveFinalizer(&tunnel, finalizerName)

			return r.Update(ctx, &tunnel)
		}
	}

	return nil
}
