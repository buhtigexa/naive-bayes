package bayes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createCorpus() []Document {
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
	corpus := createCorpus()
	result := nb.Train(corpus)
	assert.NotNil(t, result)
	assert.Equal(t, result.Docs, len(corpus))
	assert.InEpsilon(t, result.Classes["normal"].terms["launch"].probability, 0.190, 0.01)
	assert.InEpsilon(t, result.Classes["spam"].terms["launch"].probability, 0.090, 0.09)
}

func TestNaiveBayes_Predict(t *testing.T) {
	test := Document{[]string{"launch", "money", "money", "money"}, ""}
	nb := NewNaiveBayes()
	corpus := createCorpus()
	_ = nb.Train(corpus)
	prediction := nb.Predict(test)
	assert.Equal(t, "spam", prediction[0].class)

	test = Document{[]string{"dear", "friend"}, ""}

	prediction = nb.Predict(test)
	assert.Equal(t, "normal", prediction[0].class)

}
