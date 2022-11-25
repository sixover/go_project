package prefix_tree

type node struct {
	pass int
	end  int
	next map[rune]*node
}

func initPrefixTree() *node {
	return &node{
		pass: 0,
		end:  0,
		next: make(map[rune]*node),
	}
}

func (t *node) insert(keyWorld string) {
	if keyWorld == "" {
		return
	}
	convertList := []rune(keyWorld)
	tempRoot := t
	for _, resA := range convertList {
		if _, ok := tempRoot.next[resA]; !ok {
			tempRoot.next[resA] = &node{
				pass: 0,
				end:  0,
				next: make(map[rune]*node),
			}
		}
		tempRoot.pass++
		tempRoot, _ = tempRoot.next[resA]
	}
	tempRoot.pass++
	tempRoot.end++
}

func (t *node) search(keyWorld string) int {
	if keyWorld == "" {
		return 0
	}
	convertList := []rune(keyWorld)
	tempNode := t
	for _, res := range convertList {
		if _, ok := tempNode.next[res]; !ok {
			return 0
		}
		tempNode, _ = tempNode.next[res]
	}
	return tempNode.end
}

func (t *node) prefixSearch(keyWorld string) int {
	if keyWorld == "" {
		return 0
	}
	convertList := []rune(keyWorld)
	tempNode := t
	for _, res := range convertList {
		if _, ok := tempNode.next[res]; !ok {
			return 0
		}
		tempNode, _ = tempNode.next[res]
	}
	return tempNode.pass
}

func (t *node) delete(keyWorld string) {
	if keyWorld == "" {
		return
	}
	if t.search(keyWorld) == 0 {
		return
	}
	convertList := []rune(keyWorld)
	tempNode := t
	tempNode.pass--
	for _, res := range convertList {
		tempNextNode, _ := tempNode.next[res]
		tempNextNode.pass--
		if tempNextNode.pass == 0 {
			delete(tempNode.next, res)
			return
		}
		tempNode = tempNextNode
	}
	tempNode.end--
}
