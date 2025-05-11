package sqs

type Model struct {
	DeliveryDelaySeconds     int64   `yaml:"DeliveryDelaySeconds"`
	Name                     string  `yaml:"Name"`
	RemoveOnDestroy          bool    `yaml:"RemoveOnDestroy"`
	RetentionPeriodSeconds   int64   `yaml:"RetentionPeriodSeconds"`
	VisibilityTimeoutSeconds float64 `yaml:"VisibilityTimeoutSeconds"`
}
