package days

type Day struct {
	value string
}

func (d Day) String() string { return d.value }

var (
	Monday  = Day{"Monday"}
	Tuesday = Day{"Tuesday"}
	Days    = []Day{Monday, Tuesday}
)
