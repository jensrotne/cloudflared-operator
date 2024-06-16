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
		log.Log.Error(err, "unable to fetch CloudflaredTunnel")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Check if Tunnel exists in Cloudflare, if not create it
	listOptions := map[string]string{
		"name":       tunnel.Name,
		"is_deleted": "false",
	}

	cloudflareTunnels := cloudflare.ListTunnels(listOptions)

	var cloudflareTunnel cloudflare.CloudflareTunnel

	if len(cloudflareTunnels.Result) == 0 {
		log.Log.Info("Tunnel not found in Cloudflare. Creating...", "tunnel", tunnel.Name)
		// Create Tunnel

		res := cloudflare.CreateTunnel(tunnel.Name, "cloudflare", nil)

		if res.Success {
			cloudflareTunnel = res.Result

			log.Log.Info("Tunnel created", "tunnel", cloudflareTunnel.Name)
		} else {
			log.Log.Error(fmt.Errorf("unable to create tunnel"), "unable to create tunnel")
			return ctrl.Result{}, fmt.Errorf("unable to create tunnel")
		}

	} else {
		cloudflareTunnel = cloudflareTunnels.Result[0]
	}

	tunnelSecret := cloudflare.GetTunnelToken(cloudflareTunnel.ID)

	if !tunnelSecret.Success {
		log.Log.Error(fmt.Errorf("unable to get tunnel token"), "unable to get tunnel token")
		return ctrl.Result{}, fmt.Errorf("unable to get tunnel token")
	}

	// Create Tunnel secret if not exists
	var secret core.Secret

	secretName := fmt.Sprintf("%s-tunnel-secret", tunnel.Name)

	err := r.Get(ctx, client.ObjectKey{Namespace: tunnel.Namespace, Name: secretName}, &secret)

	if apierrors.IsNotFound(err) {
		log.Log.Info("Secret not found. Creating", "secret", secretName)

		secret = core.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: tunnel.Namespace,
				Annotations: map[string]string{
					"tunnel.jensrotne.com/owner": tunnel.Name,
				},
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: tunnel.APIVersion,
						Kind:       tunnel.Kind,
						Name:       tunnel.Name,
						UID:        tunnel.UID,
					},
				},
			},
			StringData: map[string]string{
				"secret": tunnelSecret.Result,
			},
		}

		if err := r.Create(ctx, &secret); err != nil {
			log.Log.Error(err, "unable to create secret")
			return ctrl.Result{}, err
		}
	} else {
		if secret.StringData["secret"] != tunnelSecret.Result {
			log.Log.Info("Secret found but secret does not match. Updating", "secret", secretName)

			secret.StringData["secret"] = tunnelSecret.Result

			if err := r.Update(ctx, &secret); err != nil {
				log.Log.Error(err, "unable to update secret")
				return ctrl.Result{}, err
			}
		}
	}

	// Check if Deployment exists, if not create it
	var deploy appsv1.Deployment

	err = r.Get(ctx, client.ObjectKey{Namespace: tunnel.Namespace, Name: tunnel.Name}, &deploy)

	if apierrors.IsNotFound(err) {
		log.Log.Info("Deployment not found. Creating...", "deployment", tunnel.Name)

		deploy = appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      tunnel.Name,
				Namespace: tunnel.Namespace,
				Annotations: map[string]string{
					"tunnel.jensrotne.com/tunnel-id": cloudflareTunnel.ID,
					"tunnel.jensrotne.com/owner":     tunnel.Name,
				},
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: tunnel.APIVersion,
						Kind:       tunnel.Kind,
						Name:       tunnel.Name,
						UID:        tunnel.UID,
					},
				},
			},
			Spec: appsv1.DeploymentSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						"app": tunnel.Name,
					},
				},
				Template: core.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: map[string]string{
							"app": tunnel.Name,
						},
					},
					Spec: core.PodSpec{
						Containers: []core.Container{
							{
								Name:  "cloudflared",
								Image: "cloudflare/cloudflared:latest",
								Command: []string{
									"cloudflared",
									"tunnel",
									"run",
								},
								Args: []string{
									"--token",
									tunnelSecret.Result,
								},
								LivenessProbe: &core.Probe{
									ProbeHandler: core.ProbeHandler{
										HTTPGet: &core.HTTPGetAction{
											Path: "/ready",
											Port: intstr.FromInt(8080),
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

		if err := r.Create(ctx, &deploy); err != nil {
			log.Log.Error(err, "unable to create Deployment")
			return ctrl.Result{}, err
		}
	}

	// Update status
	tunnel.Status.TunnelID = cloudflareTunnel.ID

	if err := r.Status().Update(ctx, &tunnel); err != nil {
		log.Log.Error(err, "unable to update CloudflaredTunnel status")
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
