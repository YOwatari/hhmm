package params

import "errors"

const OneDayMinutes int64 = 1 * 24 * 60

type DevideNumber struct {
	Value int64
}

func (d *DevideNumber) Set(n int64) error {
	if validate(n) {
		d.Value = n
	} else {
		return errors.New("not accepted")
	}

	return nil
}

func validate(n int64) bool {
	divisor := divisor(OneDayMinutes)
	return isDivisor(n, divisor)
}

func isDivisor(n int64, ns []int64) bool {
	for _, x := range ns {
		if x == n {
			return true
		}
	}
	return false
}

func divisor(n int64) []int64 {
	return []int64{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24, 30, 32, 36, 40, 45, 48, 60, 72, 80, 90, 96, 120, 144, 160, 180, 240, 288, 360, 480, 720, 1440}
}
