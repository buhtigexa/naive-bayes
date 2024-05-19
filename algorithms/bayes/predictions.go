package bayes

type prediction struct {
	class string
	prob  float64
}

type Predictions []prediction

func (p Predictions) Len() int {
	return len(p)
}
func (p Predictions) Less(i, j int) bool {
	return p[i].prob < p[j].prob
}

func (p Predictions) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
