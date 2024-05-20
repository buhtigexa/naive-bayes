package bayes

type Document struct {
	Terms []string
	Class string
}

func NewDocument(terms []string, class string) Document {
	return Document{terms, class}
}

type Metric struct {
	Frequency   int
	Probability float64
}

type Class struct {
	TotalWords int32
	Terms      map[string]*Metric
	Id         string
	TotalDocs  int
	PriorProb  float64
}

func newClass(id string) *Class {
	return &Class{
		Id:    id,
		Terms: make(map[string]*Metric),
	}
}

func (c *Class) Add(doc Document) {
	if len(doc.Terms) == 0 {
		return
	}
	for _, w := range doc.Terms {
		if _, ok := c.Terms[w]; !ok {
			c.Terms[w] = &Metric{}
		}
		c.Terms[w].Frequency++
	}
	c.TotalDocs++
	c.TotalWords += int32(len(doc.Terms))
}

func (c *Class) probs() {
	for w, m := range c.Terms {
		m.Probability = float64(c.Terms[w].Frequency) / float64(c.TotalWords)
	}
}

func (c *Class) getProb(w string) float64 {
	if metric, ok := c.Terms[w]; ok {
		return metric.Probability
	}
	return 0.0001

}

func (c *Class) addWord(w string) {
	if _, ok := c.Terms[w]; !ok {
		c.Terms[w] = &Metric{}
	}
	for _, m := range c.Terms {
		m.Frequency++
	}
	c.TotalWords += int32(len(c.Terms))
}
func (c *Class) hasWord(w string) bool {
	if _, ok := c.Terms[w]; !ok {
		return false
	}
	return true
}
