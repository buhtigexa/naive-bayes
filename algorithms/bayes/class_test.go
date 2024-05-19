package bayes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClass_Add(t *testing.T) {
	class := newClass("id")
	var doc Document = Document{[]string{"a", "b", "c", "d", "e", "f"}, "test"}
	class.Add(doc)
	for _, k := range doc.terms {
		if _, ok := class.terms[k]; !ok {
			assert.Fail(t, k)
		}
	}
	assert.Equal(t, class.totalDocs, int32(1))
	assert.Equal(t, class.totalWords, int32(len(doc.terms)))

}
