package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineSpec   `json:"spec"`
	Status PipelineStatus `json:"status"`
}

type PipelineSpec struct {
	Steps         []*PipelineStep
	RunningObject *Run
	TimeTrigger   *TimeTrigger
}

type PipelineStep struct {
	Command []string
	Args    []string
}

type TimeTrigger struct {
	CustomExpression string
}

type PipelineStatus struct {
}
