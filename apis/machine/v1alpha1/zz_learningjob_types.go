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

type LearningJobObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LearningJobParameters struct {

	// The id of the datasource to query.
	// +kubebuilder:validation:Required
	DatasourceID *float64 `json:"datasourceId" tf:"datasource_id,omitempty"`

	// The type of datasource being queried. Currently allowed values are prometheus, graphite, loki, postgres, and datadog.
	// +kubebuilder:validation:Required
	DatasourceType *string `json:"datasourceType" tf:"datasource_type,omitempty"`

	// A description of the job.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The hyperparameters used to fine tune the algorithm. See https://grafana.com/docs/grafana-cloud/machine-learning/models/ for the full list of available hyperparameters. Defaults to `map[]`.
	// +kubebuilder:validation:Optional
	HyperParams map[string]*string `json:"hyperParams,omitempty" tf:"hyper_params,omitempty"`

	// The data interval in seconds to train the data on. Defaults to `300`.
	// +kubebuilder:validation:Optional
	Interval *float64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// The metric used to query the job results.
	// +kubebuilder:validation:Required
	Metric *string `json:"metric" tf:"metric,omitempty"`

	// The name of the job.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// An object representing the query params to query Grafana with.
	// +kubebuilder:validation:Required
	QueryParams map[string]*string `json:"queryParams" tf:"query_params,omitempty"`

	// The data interval in seconds to train the data on. Defaults to `7776000`.
	// +kubebuilder:validation:Optional
	TrainingWindow *float64 `json:"trainingWindow,omitempty" tf:"training_window,omitempty"`
}

// LearningJobSpec defines the desired state of LearningJob
type LearningJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LearningJobParameters `json:"forProvider"`
}

// LearningJobStatus defines the observed state of LearningJob.
type LearningJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LearningJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LearningJob is the Schema for the LearningJobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafanajet}
type LearningJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LearningJobSpec   `json:"spec"`
	Status            LearningJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LearningJobList contains a list of LearningJobs
type LearningJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LearningJob `json:"items"`
}

// Repository type metadata.
var (
	LearningJob_Kind             = "LearningJob"
	LearningJob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LearningJob_Kind}.String()
	LearningJob_KindAPIVersion   = LearningJob_Kind + "." + CRDGroupVersion.String()
	LearningJob_GroupVersionKind = CRDGroupVersion.WithKind(LearningJob_Kind)
)

func init() {
	SchemeBuilder.Register(&LearningJob{}, &LearningJobList{})
}
