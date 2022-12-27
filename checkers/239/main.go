package main

import (
	"fmt"
	"strconv"
	"strings"
	"testlib"
	"testlib/core"
)

var nk int = 0
var ch = make(map[int]int)
var temp_ch = make(map[int]int)
var inputStream = testlib.GetReader(testlib.TaskInput)
var outputStream = testlib.GetReader(testlib.TaskOutput)
var inputFromUserStream = testlib.GetReader(testlib.InputFromUser)

//var outputToUserStream = testlib.GetWriter(testlib.OutputToUser)

type Pair struct {
	first  int
	second int
}

type Node struct {
	key    int
	height byte
	left   *Node
	right  *Node
}

func height(p *Node) byte {
	if p == nil {
		return 0
	} else {
		return p.height
	}
}

func bfactor(p *Node) int {
	return int(height(p.right)) - int(height(p.left))
}

func fixheight(p *Node) {
	var hl byte = height(p.left)
	var hr byte = height(p.right)
	if hl > hr {
		p.height = hl + 1
	} else {
		p.height = hr + 1
	}
}

func rotateright(p *Node) *Node {
	q := p.left
	p.left = q.right
	q.right = p
	fixheight(p)
	fixheight(q)
	return q
}

func rotateleft(q *Node) *Node {
	p := q.right
	q.right = p.left
	p.left = q
	fixheight(q)
	fixheight(p)
	return p
}

func balance(p *Node) *Node {
	fixheight(p)
	if bfactor(p) == 2 {
		if bfactor(p.right) < 0 {
			p.right = rotateright(p.right)
		}
		return rotateleft(p)
	}
	if bfactor(p) == -2 {
		if bfactor(p) == -2 {
			if bfactor(p.left) > 0 {
				p.left = rotateleft(p.left)
			}
			return rotateright(p)
		}
	}
	return p
}

func insert(p *Node, k int) *Node {
	if p == nil {
		temp := Node{
			key:    k,
			height: 1,
			left:   nil,
			right:  nil,
		}
		return &temp
	}
	if k < p.key {
		p.left = insert(p.left, k)
	} else {
		p.right = insert(p.right, k)
	}
	return balance(p)
}

func findmin(p *Node) *Node {
	if p.left == nil {
		return p
	} else {
		return findmin(p.left)
	}
}

func removemin(p *Node) *Node {
	if p.left == nil {
		return p.right
	}
	p.left = removemin(p.left)
	return balance(p)
}

func remove(p *Node, k int) *Node {
	if p == nil {
		return nil
	}
	if k < p.key {
		p.left = remove(p.left, k)
	} else if k > p.key {
		p.right = remove(p.right, k)
	} else { // k == p->key
		q := p.left
		r := p.right
		p = nil // deleting
		if r == nil {
			return q
		}
		min := findmin(r)
		min.right = removemin(r)
		min.left = q
		return balance(min)
	}
	return balance(p)
}

func input(k *Node, before *Node, isleft bool) {
	x := core.ReadInt(inputFromUserStream)
	if x == -1 {
		return
	}
	nk++
	if k == nil {
		k := Node{
			key:    x,
			height: 1,
			left:   nil,
			right:  nil,
		}
		if isleft {
			before.left = &k
		} else {
			before.right = &k
		}
	} else {
		k.key = x
	}
	input(k.left, k, true)
	input(k.right, k, false)
}

func check(k *Node) {
	if k == nil {
		return
	}
	if bfactor(k) > 1 {
		testlib.WriteAnswer(fmt.Sprintf("Not balanced"), testlib.WA)
		return
	}
	if k.left != nil {
		if k.key < k.left.key {
			testlib.WriteAnswer(fmt.Sprintf("Tree not AVL"), testlib.WA)
			return
		} else {
			check(k.left)
		}
	}
	if k.right != nil {
		if k.key < k.right.key {
			testlib.WriteAnswer(fmt.Sprintf("Tree not AVL"), testlib.WA)
			return
		} else {
			check(k.right)
		}
	}
	return
}

func i2b(i_0 int) bool {
	if i_0 == 0 {
		return false
	} else {
		return true
	}
}

func createMatrix0(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func checkValueOutOfSegmen(val, a, b int) bool {
	if val < a || val > b {
		return true
	} else {
		return false
	}
}

func main() {
	defer func() {
		_ = testlib.GetStream(testlib.TaskOutput).Close()
		_ = testlib.GetStream(testlib.InputFromUser).Close()
	}()
	//var cnt float64 = 0.0
	n := core.ReadInt(inputStream)
	var x int
	var q int
	var m int
	var firstRow string = core.ReadLine(inputStream, true)
	firstRowSplited := strings.Split(firstRow, " ")
	for i := 0; i < len(firstRowSplited); i++ {
		x, _ = strconv.Atoi(firstRowSplited[i])
		if x != -1 {
			ch[x]++
		}
	}
	m = core.ReadInt(inputStream)
	for i := 0; i < m; i++ {
		nk = 0
		q = core.ReadInt(inputStream)
		x = core.ReadInt(inputStream)
		if q == 1 {
			ch[x]++
			n++
		} else {
			ch[x]--
			n--
		}
		temp_ch = ch
		k := Node{
			key:    0,
			height: 1,
			left:   nil,
			right:  nil,
		}
		input(&k, nil, false)
		if nk != n {
			testlib.WriteAnswer(fmt.Sprintf("Less vertices %d %d %d", x, nk, n), testlib.WA)
			return
		}
		check(&k)
		ch = temp_ch
	}
	testlib.WriteAnswer(fmt.Sprintf("OK"), testlib.OK)
}
