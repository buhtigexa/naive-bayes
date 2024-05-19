package bayes

import "fmt"

type NaiveBayes struct {
	classes map[string]*class
	words   map[string]bool
}

func NewNaiveBayes() *NaiveBayes {
	return &NaiveBayes{
		classes: make(map[string]*class),
		words:   make(map[string]bool),
	}
}

type TrainResult struct {
	Docs    int
	classes map[string]class
}

func (nb *NaiveBayes) Train(documents []Document) *TrainResult {
	var trainResult = &TrainResult{}
	var totalDocs int32 = 0

	for _, doc := range documents {
		if _, ok := nb.classes[doc.class]; !ok {
			nb.classes[doc.class] = newClass(doc.class)
		}
		nb.classes[doc.class].Add(doc)
		for w, _ := range nb.classes[doc.class].terms {
			nb.words[w] = true
		}
		totalDocs++
	}

	nb.balanceModel()

	trainResult.classes = make(map[string]class, len(nb.classes))

	for w, c := range nb.classes {
		c.probs()
		trainResult.Docs += c.totalDocs
		c.priorProb = float64(c.totalDocs) / float64(totalDocs)
		trainResult.classes[w] = class{
			terms:     c.terms,
			id:        c.id,
			totalDocs: c.totalDocs,
			priorProb: c.priorProb,
		}
	}
	return trainResult
}

func (nb *NaiveBayes) balanceModel() {
	for w, _ := range nb.words {
		for _, class := range nb.classes {
			if !class.hasWord(w) {
				fmt.Printf("Class %s no tiene %s \n", class.id, w)
				for _, uclass := range nb.classes {
					uclass.addWord(w)
					fmt.Printf(" UPDATEANDO (%s,%s)\n ", uclass.id, w)
				}
			}
		}
	}
}
