package bayes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClass_Add(t *testing.T) {
	class := newClass("id")
	var doc Document = Document{[]string{"a", "b", "c", "d", "e", "f"}, "test"}
	class.Add(doc)
	for _, k := range doc.Terms {
		if _, ok := class.terms[k]; !ok {
			assert.Fail(t, k)
		}
	}
	assert.Equal(t, 1, class.totalDocs)
	assert.Equal(t, int32(len(doc.Terms)), class.totalWords)
}
