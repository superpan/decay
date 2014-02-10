package decay

import (
	"math"
	"time"
)

// Reddit's popular sort
func Reddit(decay float64) func(ups float64, downs float64, date time.Time) float64 {
	return func(ups float64, downs float64, date time.Time) float64 {
		s := ups - downs
		order := math.Log10(math.Max(math.Abs(s), 1))
		age := time.Now().Sub(date).Seconds()

		return order - age/decay
	}
}

// hacker news popular sort
func HackerNews(gravity float64) func(votes int, date time.Time) float64 {
	return func(votes int, date time.Time) float64 {
		var age, base float64
		age = time.Now().Sub(date).Hours()
		base = age + 2
		return float64(votes-1) / math.Pow(base, gravity)
	}
}

// wilson score popular sort
func WilsonScore(z float64) func(ups float64, downs float64) float64 {
	return func(ups, downs float64) float64 {
		var n, p, zzfn float64
		n = ups + downs
		if n == 0 {
			return 0
		}

		p = ups / n
		zzfn = z * z / (4 * n)
		return (p + 2*zzfn - z*math.Sqrt((zzfn/n+p*(1-p))/n)) / (1 + 4*zzfn)
	}
}
