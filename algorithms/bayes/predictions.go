package bayes

type Prediction struct {
	Class string
	Prob  float64
}

type Predictions []Prediction

func (p Predictions) Len() int {
	return len(p)
}
func (p Predictions) Less(i, j int) bool {
	return p[i].Prob < p[j].Prob
}

func (p Predictions) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
