package Objects

type Status string

const (
	Done    Status = "Done"
	Active         = "Active"
	Unknown        = "Unknown"
)

func (s Status) String() string {
	switch s {
	case Done:
		return "Done"
	case Active:
		return "Active"
	}
	return ""
}

func CreateStatus(sString string) Status {
	switch sString {
	case "Active":
		return Active
	case "Done":
		return Done
		//case "":
		//	return active
	}
	return Unknown
}
