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
	totalWords int32
	terms      map[string]*Metric
	id         string
	totalDocs  int
	priorProb  float64
}

func newClass(id string) *Class {
	return &Class{
		id:    id,
		terms: make(map[string]*Metric),
	}
}

func (c *Class) Add(doc Document) {
	if len(doc.Terms) == 0 {
		return
	}
	for _, w := range doc.Terms {
		if _, ok := c.terms[w]; !ok {
			c.terms[w] = &Metric{}
		}
		c.terms[w].Frequency++
	}
	c.totalDocs++
	c.totalWords += int32(len(doc.Terms))
}

func (c *Class) probs() {
	for w, m := range c.terms {
		m.Probability = float64(c.terms[w].Frequency) / float64(c.totalWords)
	}
}

func (c *Class) getProb(w string) float64 {
	if metric, ok := c.terms[w]; ok {
		return metric.Probability
	}
	return 0.0001

}

func (c *Class) addWord(w string) {
	if _, ok := c.terms[w]; !ok {
		c.terms[w] = &Metric{}
	}
	for _, m := range c.terms {
		m.Frequency++
	}
	c.totalWords += int32(len(c.terms))
}
func (c *Class) hasWord(w string) bool {
	if _, ok := c.terms[w]; !ok {
		return false
	}
	return true
}
