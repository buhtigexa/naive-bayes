package bayes

import (
	"fmt"
	"sort"
)

type NaiveBayes struct {
	classes   map[string]*class
	words     map[string]bool
	totalDocs int32
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

func (nb *NaiveBayes) add(doc Document) {
	if _, ok := nb.classes[doc.class]; !ok {
		nb.classes[doc.class] = newClass(doc.class)
	}
	nb.classes[doc.class].Add(doc)
	for w, _ := range nb.classes[doc.class].terms {
		nb.words[w] = true
	}
	nb.totalDocs++

}

func (nb *NaiveBayes) Train(documents []Document) *TrainResult {
	var trainResult = &TrainResult{}

	for _, doc := range documents {
		nb.add(doc)
	}

	nb.balance()

	trainResult.classes = make(map[string]class, len(nb.classes))

	for w, c := range nb.classes {
		c.probs()
		trainResult.Docs += c.totalDocs
		c.priorProb = float64(c.totalDocs) / float64(nb.totalDocs)
		trainResult.classes[w] = class{
			terms:     c.terms,
			id:        c.id,
			totalDocs: c.totalDocs,
			priorProb: c.priorProb,
		}
	}
	return trainResult
}

func (nb *NaiveBayes) Predict(doc Document) Predictions {
	var predictions Predictions
	for _, class := range nb.classes {
		probs := 1.0
		for _, word := range doc.terms {
			probs *= class.getProb(word)
		}
		predictions = append(predictions, prediction{class.id, class.priorProb * probs})
	}
	sort.Sort(sort.Reverse(predictions))
	return predictions
}

func (nb *NaiveBayes) balance() {
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
