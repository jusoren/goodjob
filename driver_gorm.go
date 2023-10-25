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

func (d *DriverGorm) CreateExecution(execution *Execution) error {
	return d.DB.Create(execution).Error
}

func (d *DriverGorm) UpdateExecution(execution *Execution) error {
	return d.DB.Save(execution).Error
}

func (d *DriverGorm) FindJobs(options FindJobsOptions) ([]Job, error) {
	var query = d.DB
	if options.Name != "" {
		query = query.Where("name = ?", options.Name)
	}
	if options.Status != "" {
		query = query.Where("status = ?", options.Status)
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

func (d *DriverGorm) GetPendingTimeoutAndErrorJobs() ([]Job, error) {
	var jobs []Job
	err := d.DB.Where("status = ? OR (status = ? AND is_timeout = ?)", "pending", "running", true).Find(&jobs).Error
	return jobs, err
}
