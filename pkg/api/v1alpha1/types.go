package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// Clusters is configuration of remote clusters
// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Clusters struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:",inline"`
	Clusters          []Cluster `json:"clusters,omitempty"`
}

// Cluster represents remote cluster api access
type Cluster struct {
	Name string `json:"name"`
}

func init() {
	SchemeBuilder.Register(&Clusters{})
}
