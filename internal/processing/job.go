package processing

type Job struct {
	Action string
}

type JobPayload struct {
	Action string `json:"action"`
}

func NewJob(payload *JobPayload) Job {
	return Job{
		Action: payload.Action,
	}
}
