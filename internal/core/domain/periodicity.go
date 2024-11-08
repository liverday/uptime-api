package domain

type Periodicity string

const (
	OneMinute      Periodicity = "1m"
	FiveMinutes    Periodicity = "5m"
	TenMinutes     Periodicity = "10m"
	FifteenMinutes Periodicity = "15m"
	ThirtyMinutes  Periodicity = "30m"
)

func (p Periodicity) String() string {
	return string(p)
}

var Periods = []Periodicity{OneMinute, FiveMinutes, TenMinutes, FifteenMinutes, ThirtyMinutes}
