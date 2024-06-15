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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CloudflaredTunnelSpec defines the desired state of CloudflaredTunnel
type CloudflaredTunnelSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	// Service to target with the Cloudflared Tunnel
	TargetService string `json:"targetService"`

	// +kubebuilder:validation:Required
	// Service Port to target with the Cloudflared Tunnel
	TargetPort int `json:"targetPort"`
}

// CloudflaredTunnelStatus defines the observed state of CloudflaredTunnel
type CloudflaredTunnelStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	TunnelID string `json:"tunnelId"`
	Message  string `json:"message"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CloudflaredTunnel is the Schema for the cloudflaredtunnels API
type CloudflaredTunnel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudflaredTunnelSpec   `json:"spec,omitempty"`
	Status CloudflaredTunnelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudflaredTunnelList contains a list of CloudflaredTunnel
type CloudflaredTunnelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudflaredTunnel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudflaredTunnel{}, &CloudflaredTunnelList{})
}
