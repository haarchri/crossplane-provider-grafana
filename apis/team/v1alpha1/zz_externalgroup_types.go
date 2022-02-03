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

type ExternalGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ExternalGroupParameters struct {

	// The team external groups list
	// +kubebuilder:validation:Required
	Groups []*string `json:"groups" tf:"groups,omitempty"`

	// The Team ID
	// +kubebuilder:validation:Required
	TeamID *int64 `json:"teamId" tf:"team_id,omitempty"`
}

// ExternalGroupSpec defines the desired state of ExternalGroup
type ExternalGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExternalGroupParameters `json:"forProvider"`
}

// ExternalGroupStatus defines the observed state of ExternalGroup.
type ExternalGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExternalGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalGroup is the Schema for the ExternalGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type ExternalGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExternalGroupSpec   `json:"spec"`
	Status            ExternalGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalGroupList contains a list of ExternalGroups
type ExternalGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalGroup `json:"items"`
}

// Repository type metadata.
var (
	ExternalGroup_Kind             = "ExternalGroup"
	ExternalGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExternalGroup_Kind}.String()
	ExternalGroup_KindAPIVersion   = ExternalGroup_Kind + "." + CRDGroupVersion.String()
	ExternalGroup_GroupVersionKind = CRDGroupVersion.WithKind(ExternalGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ExternalGroup{}, &ExternalGroupList{})
}
