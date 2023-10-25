package goodjob

type Driver interface {
	CreateJob(*Job) error
	FindJobs(options FindJobsOptions) ([]Job, error)
	FindOneJob(options FindOneJobOptions) (Job, error)
}

type FindJobsOptions struct {
	Name      string
	Status    string
	IsTimeout *bool
}

type FindOneJobOptions struct {
	ID   string
	Name string
}
