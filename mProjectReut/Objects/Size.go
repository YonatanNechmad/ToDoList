package Objects

type Size string

const (
	Small       Size = "Small"
	Medium           = "Medium"
	Large            = "Large"
	UnknownSize      = "UnknownSize"
)

func CreateSize(size string) Size {
	switch size {
	case "Small":
		return Small
	case "Medium":
		return Medium
	case "Large":
		return Large
	}
	return UnknownSize
}

func (s Size) String() string {
	switch s {
	case Small:
		return "Small"
	case Medium:
		return "Medium"
	case Large:
		return "Large"
	}
	return ""
}
