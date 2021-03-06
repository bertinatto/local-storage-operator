package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeAPIServer provides information to configure an operator to manage kube-apiserver.
type KubeAPIServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	// +kubebuilder:validation:Required
	// +required
	Spec KubeAPIServerSpec `json:"spec"`
	// +optional
	Status KubeAPIServerStatus `json:"status"`
}

type KubeAPIServerSpec struct {
	StaticPodOperatorSpec `json:",inline"`
}

type KubeAPIServerStatus struct {
	StaticPodOperatorStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeAPIServerList is a collection of items
type KubeAPIServerList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ListMeta `json:"metadata"`
	// Items contains the items
	Items []KubeAPIServer `json:"items"`
}
