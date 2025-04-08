package domains

type FileStatus string

func (x FileStatus) String() string { return string(x) }

const (
	FileCreated FileStatus = "created"
)
