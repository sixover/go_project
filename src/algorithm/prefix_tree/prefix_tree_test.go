package prefix_tree

import "testing"

func TestPrefixTree(t *testing.T) {
	root := initPrefixTree()
	strList := []string{"abc", "abd", "bce", "abcd", "bcf"}
	for _, str := range strList {
		root.insert(str)
	}
	t.Log("ab prefix: ", root.prefixSearch("ab"))
	t.Log("abc search: ", root.search("abc"))
	root.delete("abd")
	t.Log("ab prefix: ", root.prefixSearch("ab"))
}
