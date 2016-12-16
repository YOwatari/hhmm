package params

import "time"

const gap int64 = 0
const oneHourMinutes int64 = 60

type HHMM struct {
	All    []int64
	Before int64
	After  int64
}

func (h *HHMM) Set(n int64) error {
	h.All = make([]int64, n)

	var min int64 = OneDayMinutes / n
	for i, _ := range h.All {
		h.All[i] = carry(gap + int64(i)*min)
	}

	t := time.Now()
	now := int64(t.Hour()*100 + t.Minute())

	for i, v := range h.All {
		if v >= now {
			h.Before = h.All[i-1]
			h.After = v
			break
		}
	}

	return nil
}

func carry(n int64) int64 {
	d := n / oneHourMinutes
	m := n % oneHourMinutes

	return d*100 + m
}
