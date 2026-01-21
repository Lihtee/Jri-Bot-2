package frequency

type Frequency int

const (
	Year  Frequency = 1
	Month Frequency = 12
	Week  Frequency = Month * 4
	Day   Frequency = 365 // Use number of day in a year as basis for everything.
)
