package v1alpha1

type Status string

const (
	Complete Status = "complete"
	Running  Status = "running"
	Failed   Status = "failed"
	Waiting  Status = "waiting"
	Unknown  Status = "unknown"
)
