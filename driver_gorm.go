package goodjob

import "gorm.io/gorm"

type DriverGorm struct {
	DB *gorm.DB
	Driver
}

func NewDriverGorm(db *gorm.DB) *DriverGorm {
	return &DriverGorm{DB: db}
}

func (d *DriverGorm) CreateJob(job *Job) error {
	return d.DB.Create(job).Error
}

func (d *DriverGorm) FindJobs(options FindJobsOptions) error {
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
	return query.Find(&[]Job{}).Error
}
