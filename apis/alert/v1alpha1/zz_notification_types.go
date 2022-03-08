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

type NotificationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NotificationParameters struct {

	// Whether to disable sending resolve messages. Defaults to `false`.
	// +kubebuilder:validation:Optional
	DisableResolveMessage *bool `json:"disableResolveMessage,omitempty" tf:"disable_resolve_message,omitempty"`

	// Frequency of alert reminders. Frequency must be set if reminders are enabled. Defaults to ``.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Is this the default channel for all your alerts. Defaults to `false`.
	// +kubebuilder:validation:Optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default,omitempty"`

	// The name of the alert notification channel.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Additional secure settings, for full reference lookup [Grafana Supported Settings documentation](https://grafana.com/docs/grafana/latest/administration/provisioning/#supported-settings).
	// +kubebuilder:validation:Optional
	SecureSettingsSecretRef map[string]v1.SecretKeySelector `json:"secureSettingsSecretRef" tf:"-"`

	// Whether to send reminders for triggered alerts. Defaults to `false`.
	// +kubebuilder:validation:Optional
	SendReminder *bool `json:"sendReminder,omitempty" tf:"send_reminder,omitempty"`

	// Additional settings, for full reference see [Grafana HTTP API documentation](https://grafana.com/docs/grafana/latest/http_api/alerting_notification_channels/).
	// +kubebuilder:validation:Optional
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`

	// The type of the alert notification channel.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Unique identifier. If unset, this will be automatically generated.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

// NotificationSpec defines the desired state of Notification
type NotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationParameters `json:"forProvider"`
}

// NotificationStatus defines the observed state of Notification.
type NotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Notification is the Schema for the Notifications API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type Notification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationSpec   `json:"spec"`
	Status            NotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationList contains a list of Notifications
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notification `json:"items"`
}

// Repository type metadata.
var (
	Notification_Kind             = "Notification"
	Notification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Notification_Kind}.String()
	Notification_KindAPIVersion   = Notification_Kind + "." + CRDGroupVersion.String()
	Notification_GroupVersionKind = CRDGroupVersion.WithKind(Notification_Kind)
)

func init() {
	SchemeBuilder.Register(&Notification{}, &NotificationList{})
}
