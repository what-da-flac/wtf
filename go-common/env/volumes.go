package env

type Volumes struct {
	Media Names
}

func newVolumes() Volumes {
	return Volumes{
		Media: "/media",
	}
}
