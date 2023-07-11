/*
Copyright 2023 rakshitgondwal.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CivoClusterSpec defines the desired state of CivoCluster
type CivoClusterSpec struct {
	ClusterName string `json:"cluster_name"`
	APIKey string `json:"api_key"`
	Region string `json:"region"`
	NetworkID string `json:"networkid"`
}

// CivoClusterStatus defines the observed state of CivoCluster
type CivoClusterStatus struct {
	RunningStatus string `json:"running_status"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CivoCluster is the Schema for the civoclusters API
type CivoCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CivoClusterSpec   `json:"spec,omitempty"`
	Status CivoClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CivoClusterList contains a list of CivoCluster
type CivoClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CivoCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CivoCluster{}, &CivoClusterList{})
}
