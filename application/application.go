package application

var (
	fmt_greet    = "Hi, I'm %s!"
	SRC_SYOBOCAL = "SYOBOCAL"
	SRC_LOCAL    = "LOCAL"
)

type Application interface {
	Greet() string
}

func Get(src string) Application {
	if src == SRC_SYOBOCAL {
		return Syobocal{"Syobocal"}
	}
	return Local{"Local"}
}
