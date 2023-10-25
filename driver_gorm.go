package goodjob

import (
	"errors"

	"gorm.io/gorm"
)

type DriverGorm struct {
	DB *gorm.DB
}

func NewDriverGorm(db *gorm.DB) *DriverGorm {
	return &DriverGorm{DB: db}
}

func (d *DriverGorm) CreateJob(job *Job) error {
	if job.Name == "" {
		return errors.New("name is required")
	}

	return d.DB.Create(job).Error
}

func (d *DriverGorm) FindJobs(options FindJobsOptions) ([]Job, error) {
	var query = d.DB
	if options.Name != "" {
		query = query.Where("name = ?", options.Name)
	}
	if options.Status != "" {
		query = query.Where("status = ?", options.Status)
	}
	if options.IsTimeout != nil {
		query = query.Where("is_timeout = ?", options.IsTimeout)
	}

	result := []Job{}

	if err := query.Find(&result).Error; err != nil {
		return []Job{}, err
	}

	return result, nil
}

func (d *DriverGorm) FindOneJob(options FindOneJobOptions) (Job, error) {
	var job Job
	var query = d.DB
	if options.ID != "" {
		query = query.Where("id = ?", options.ID)
	}
	if options.Name != "" {
		query = query.Where("name = ?", options.Name)
	}
	err := query.First(&job).Error
	return job, err
}
