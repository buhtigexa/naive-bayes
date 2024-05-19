package bayes

type Document struct {
	terms []string
	class string
}

type metric struct {
	frequency   int
	probability float64
}

type class struct {
	totalWords int32
	terms      map[string]*metric
	id         string
	totalDocs  int
	priorProb  float64
}

func newClass(id string) *class {
	return &class{
		id:    id,
		terms: make(map[string]*metric),
	}
}

func (c *class) Add(doc Document) {
	if len(doc.terms) == 0 {
		return
	}
	for _, w := range doc.terms {
		if _, ok := c.terms[w]; !ok {
			c.terms[w] = &metric{}
		}
		c.terms[w].frequency++
	}
	c.totalDocs++
	c.totalWords += int32(len(doc.terms))
}

func (c *class) probs() {
	for w, m := range c.terms {
		m.probability = float64(c.terms[w].frequency) / float64(c.totalWords)
	}
}

func (c *class) addWord(w string) {
	if _, ok := c.terms[w]; !ok {
		c.terms[w] = &metric{}
	}
	for _, m := range c.terms {
		m.frequency++
	}
	c.totalWords += int32(len(c.terms))
}
func (c *class) hasWord(w string) bool {
	if _, ok := c.terms[w]; !ok {
		return false
	}
	return true
}
