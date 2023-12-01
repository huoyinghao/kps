package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Run struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunSpec   `json:"spec"`
	Status RunStatus `json:"status"`
}

type RunSpec struct {
	Steps []*PipelineStep `json:"steps"`
}

type RunStatus struct {
	Status Status
}
