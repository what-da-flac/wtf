package rabbit

type namer struct {
	name string
}

func newNamer(name string) *namer { return &namer{name: name} }

func (x *namer) exchangeName() string { return x.name + "_exchange" }
func (x *namer) keyName() string      { return x.name + "_key" }
func (x *namer) queueName() string    { return x.name + "_queue" }
