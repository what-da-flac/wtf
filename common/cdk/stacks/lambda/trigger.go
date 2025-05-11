package lambda

type triggerType string

const (
	TriggerTypeSQS triggerType = "SQS"
)

type Trigger struct {
	Type triggerType `yaml:"Type"`
}

func (x *Trigger) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}
	if err := unmarshal(&raw); err != nil {
		return err
	}
	val, ok := raw["Type"]
	if !ok {
		return nil
	}
	if v, ok := val.(string); ok {
		//nolint:gocritic
		switch v {
		case string(TriggerTypeSQS):
			x.Type = TriggerTypeSQS
		}
	}
	return nil
}
