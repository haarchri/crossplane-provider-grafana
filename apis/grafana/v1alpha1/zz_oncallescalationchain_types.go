/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OncallEscalationChainObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OncallEscalationChainParameters struct {

	// The name of the escalation chain.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`
}

// OncallEscalationChainSpec defines the desired state of OncallEscalationChain
type OncallEscalationChainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OncallEscalationChainParameters `json:"forProvider"`
}

// OncallEscalationChainStatus defines the observed state of OncallEscalationChain.
type OncallEscalationChainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OncallEscalationChainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OncallEscalationChain is the Schema for the OncallEscalationChains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type OncallEscalationChain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OncallEscalationChainSpec   `json:"spec"`
	Status            OncallEscalationChainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OncallEscalationChainList contains a list of OncallEscalationChains
type OncallEscalationChainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OncallEscalationChain `json:"items"`
}

// Repository type metadata.
var (
	OncallEscalationChain_Kind             = "OncallEscalationChain"
	OncallEscalationChain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OncallEscalationChain_Kind}.String()
	OncallEscalationChain_KindAPIVersion   = OncallEscalationChain_Kind + "." + CRDGroupVersion.String()
	OncallEscalationChain_GroupVersionKind = CRDGroupVersion.WithKind(OncallEscalationChain_Kind)
)

func init() {
	SchemeBuilder.Register(&OncallEscalationChain{}, &OncallEscalationChainList{})
}
