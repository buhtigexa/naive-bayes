package bayes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getCorpus() []Document {
	doc1 := Document{[]string{"dear", "friend", "launch", "money"}, "normal"}
	doc2 := Document{[]string{"dear", "friend", "launch"}, "normal"}
	doc3 := Document{[]string{"dear", "friend", "launch"}, "normal"}
	doc4 := Document{[]string{"dear", "friend"}, "normal"}
	doc5 := Document{[]string{"dear", "friend"}, "normal"}
	doc6 := Document{[]string{"dear"}, "normal"}
	doc7 := Document{[]string{"dear"}, "normal"}
	doc8 := Document{[]string{"dear"}, "normal"}

	doc9 := Document{[]string{"dear", "dear", "friend", "money"}, "spam"}
	doc10 := Document{[]string{"money"}, "spam"}
	doc11 := Document{[]string{"money"}, "spam"}
	doc12 := Document{[]string{"money"}, "spam"}

	corpus := []Document{doc1, doc2, doc3, doc4, doc5, doc6, doc7, doc8, doc9, doc10, doc11, doc12}

	return corpus
}

func TestNaiveBayes_Train(t *testing.T) {
	nb := NewNaiveBayes()
	docs := getCorpus()
	result := nb.Train(docs)
	assert.NotNil(t, result)
	assert.Equal(t, result.Docs, len(docs))
	assert.InEpsilon(t, result.classes["normal"].terms["launch"].probability, 0.190, 0.01)
	assert.InEpsilon(t, result.classes["spam"].terms["launch"].probability, 0.090, 0.09)
}
