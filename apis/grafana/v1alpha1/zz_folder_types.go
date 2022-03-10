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

type FolderObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FolderParameters struct {

	// The title of the folder.
	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`

	// Unique identifier.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

// FolderSpec defines the desired state of Folder
type FolderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderParameters `json:"forProvider"`
}

// FolderStatus defines the observed state of Folder.
type FolderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Folder is the Schema for the Folders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type Folder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderSpec   `json:"spec"`
	Status            FolderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderList contains a list of Folders
type FolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Folder `json:"items"`
}

// Repository type metadata.
var (
	Folder_Kind             = "Folder"
	Folder_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Folder_Kind}.String()
	Folder_KindAPIVersion   = Folder_Kind + "." + CRDGroupVersion.String()
	Folder_GroupVersionKind = CRDGroupVersion.WithKind(Folder_Kind)
)

func init() {
	SchemeBuilder.Register(&Folder{}, &FolderList{})
}
