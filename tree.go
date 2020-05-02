package maptree

import (
	"strings"
)

const (
	// DefaultSeparator is the string that denotes each part of the path
	DefaultSeparator = "/"
)

// MapTree manages data on a map to help returning data hierarchy
type MapTree struct {
	root      *Leaf
	separator string
}

// Leaf is the unit that contains data and sub leaves
type Leaf struct {
	Leafs    map[string]Leaf `json:"items,omitempty"`
	Metadata interface{}     `json:"meta,omitempty"`
	Data     interface{}     `json:"data,omitempty"`
}

// New returns an empty MapTree
func New() *MapTree {
	return &MapTree{
		root: &Leaf{
			Leafs: make(map[string]Leaf),
		},
		separator: DefaultSeparator,
	}
}

// Separator allows to change the separator with a custom one
func (m *MapTree) Separator(separator string) {

	m.separator = separator

}

// Tree returns the root leaf that contains the tree
func (m *MapTree) Tree() *Leaf {

	return m.root
}

// Set adds a new leaf on the tree
func (m *MapTree) Set(key string, metadata, data interface{}) bool {

	return m.recursiveInsert(m.root, key, metadata, data, 0)

}

// Delete the data from a leaf on the tree
func (m *MapTree) Delete(key string) bool {

	return m.recursiveRemoval(m.root, key, 0)

}

func (m *MapTree) recursiveInsert(l *Leaf, key string, metadata, data interface{}, depth int) bool {

	var (
		parts    []string
		thisLeaf Leaf
		ok       bool
		isNew    bool
	)

	parts = strings.Split(key, m.separator)

	_, ok = l.Leafs[parts[depth]]
	if !ok {
		// create leaf
		l.Leafs[parts[depth]] = Leaf{
			Leafs: make(map[string]Leaf),
		}
		isNew = true
	}

	thisLeaf = l.Leafs[parts[depth]]

	// is the final leaf?
	if depth == len(parts)-1 {
		thisLeaf.Metadata = metadata
		thisLeaf.Data = data
		l.Leafs[parts[depth]] = thisLeaf
	} else {
		m.recursiveInsert(&thisLeaf, key, metadata, data, depth+1)
	}

	return isNew

}

func (m *MapTree) recursiveRemoval(l *Leaf, key string, depth int) bool {

	var (
		parts    []string
		thisLeaf Leaf
		ok       bool
	)

	parts = strings.Split(key, m.separator)

	// we try to remove now
	if depth == len(parts)-1 {
		// full subtree?
		_, ok = l.Leafs[parts[depth]]

		if ok {
			delete(l.Leafs, parts[depth])
		}

		return ok // return if deleted

	}

	// subkey exists? then walk inside
	thisLeaf, ok = l.Leafs[parts[depth]]
	if ok {
		return m.recursiveRemoval(&thisLeaf, key, depth+1)
	}

	return ok

}
