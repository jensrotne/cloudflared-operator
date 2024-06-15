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
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/log"

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

	// Check if Tunnel exists in Cloudflare
	listOptions := map[string]string{
		"name":       tunnel.Name,
		"is_deleted": "false",
	}

	cloudflareTunnels := cloudflare.ListTunnels(listOptions)

	var cloudflareTunnel cloudflare.CloudflareTunnel

	if len(cloudflareTunnels.Result) == 0 {
		// Create Tunnel

		// Generate random base64 encoded secret
		secret := make([]byte, 32)
		_, err := rand.Read(secret)
		if err != nil {
			log.Log.Error(err, "unable to generate random secret")
			return ctrl.Result{}, err
		}

		tunnelSecret := base64.StdEncoding.EncodeToString(secret)

		res := cloudflare.CreateTunnel(tunnel.Name, "cloudflare", tunnelSecret)

		if res.Success {
			cloudflareTunnel = res.Result
		} else {
			log.Log.Error(fmt.Errorf("unable to create tunnel"), "unable to create tunnel")
			return ctrl.Result{}, fmt.Errorf("unable to create tunnel")
		}

	} else {
		cloudflareTunnel = cloudflareTunnels.Result[0]
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
