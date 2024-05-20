package bayes

import (
	"fmt"
	"sort"
)

type NaiveBayes struct {
	classes   map[string]*Class
	words     map[string]bool
	totalDocs int32
}

func NewNaiveBayes() *NaiveBayes {
	return &NaiveBayes{
		classes: make(map[string]*Class),
		words:   make(map[string]bool),
	}
}

type TrainResult struct {
	Docs    int
	Classes map[string]Class
}

func (nb *NaiveBayes) add(doc Document) {
	if _, ok := nb.classes[doc.Class]; !ok {
		nb.classes[doc.Class] = newClass(doc.Class)
	}
	nb.classes[doc.Class].Add(doc)
	for w, _ := range nb.classes[doc.Class].Terms {
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

	trainResult.Classes = make(map[string]Class, len(nb.classes))

	for w, c := range nb.classes {
		c.probs()
		trainResult.Docs += c.TotalDocs
		c.PriorProb = float64(c.TotalDocs) / float64(nb.totalDocs)
		trainResult.Classes[w] = Class{
			Terms:     c.Terms,
			Id:        c.Id,
			TotalDocs: c.TotalDocs,
			PriorProb: c.PriorProb,
		}
	}
	return trainResult
}

func (nb *NaiveBayes) Predict(doc Document) Predictions {
	var predictions Predictions
	for _, class := range nb.classes {
		probs := 1.0
		for _, word := range doc.Terms {
			probs *= class.getProb(word)
		}
		predictions = append(predictions, Prediction{class.Id, class.PriorProb * probs})
	}
	sort.Sort(sort.Reverse(predictions))
	return predictions
}

func (nb *NaiveBayes) balance() {
	for w, _ := range nb.words {
		for _, class := range nb.classes {
			if !class.hasWord(w) {
				fmt.Printf("Class %s no tiene %s \n", class.Id, w)
				for _, uclass := range nb.classes {
					uclass.addWord(w)
					fmt.Printf(" UPDATEANDO (%s,%s)\n ", uclass.Id, w)
				}
			}
		}
	}
}
