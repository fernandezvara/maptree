package maptree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {

	assert := assert.New(t)

	assert.Equal("a", "a", "Must be equal")

	aTree := New()

	aTree.Set("a", "metaa", "dataa")
	aTree.Set("a/a", "metaaa", "dataaa")

	assert.Len(aTree.Tree().Leafs, 1, "Must have 1 leaf")
	assert.True(aTree.Set("this/is/a/long/path", "metadata", "data"), "This must be a new record.")
	assert.False(aTree.Set("this/is/a/long/path", "metadata'", "data'"), "This must be an update.")
	assert.Len(aTree.Tree().Leafs, 2, "Must have 2 sub leaves")

	assert.False(aTree.Delete("non-existent-entry"), "This must be false, entry does not exists")
	assert.False(aTree.Delete("this/is/a/non-existent/path"), "This must be false, entry does not exists")
	assert.True(aTree.Delete("this/is/a/long/path"), "This must be true, entry does exists")

	aTree.Separator(".")
	assert.Equal(aTree.separator, ".")
	assert.True(aTree.Delete("this.is"), "Must be true")
	assert.Len(aTree.Tree().Leafs, 2, "Must have 2 sub leaves")
	assert.True(aTree.Delete("this"), "Must be true")
	assert.Len(aTree.Tree().Leafs, 1, "Must have 1 leaf")
	assert.True(aTree.Delete("a"), "Must be true, it deletes the rest of the tree")
	assert.Len(aTree.Tree().Leafs, 0, "Must be empty")

}
