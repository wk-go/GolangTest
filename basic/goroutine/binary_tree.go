package main
import "fmt"

type Stack struct {
	queue map[int]*Tree
}

func (self *Stack) Push(tree *Tree) {
	self.queue[len(self.queue)] = tree
}

func (self *Stack) Pop() *Tree {
	tree, ok := self.queue[len(self.queue) - 1]
	if ok {
		delete(self.queue, len(self.queue) - 1)
		return tree
	}
	return nil
}
func (self *Stack) Empty() bool {
	return len(self.queue) == 0
}


type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}
// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。迭代遍历二叉树
func Walk(t *Tree, ch chan int) {
	stack := &Stack{make(map[int]*Tree)}
	var t_curr *Tree
	for {
		//总是把左节点压入堆栈
		for t != nil {
			stack.Push(t)
			t = t.Left
		}
		for !stack.Empty() {
			t_curr = stack.Pop()
			ch <- t_curr.Value
			if t_curr.Right != nil {
				t = t_curr.Right
				break
			}
		}
		if stack.Empty() && t == nil {
			break
		}
	}
	close(ch)
}

//递归遍历二叉树
func Walk_recursion(t *Tree) {
	if t.Left != nil {
		Walk_recursion(t.Left)
	}
	fmt.Println(t.Value)
	if t.Right != nil {
		Walk_recursion(t.Right)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	ch1,ch2 := make(chan int),make(chan int)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2{
			return false
		}
		if !ok1 || !ok2 {
			break
		}
		if v1 != v2{
			return false
		}
	}
	return true
}

func main() {
	var tree1, tree2 Tree
	ch := make(chan int)

	fmt.Println("tree1:")
	tree1 = Tree{Value:3}
	tree1.Left = &Tree{
		&Tree{Value:1},
		1,
		&Tree{Value:2},
	}
	tree1.Right = &Tree{
		&Tree{Value:5},
		8,
		&Tree{Value:13},
	}
	Walk_recursion(&tree1)

	/*go Walk(&tree1, ch)
	fmt.Println()
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(v)
	}*/

	fmt.Println("\ntree2:")
	tree2 = Tree{
		Left:&Tree{
			Left:&Tree{
				Left:&Tree{
					Value:1,
				},
				Value:1,
				Right:&Tree{
					Value:2,
				},
			},
			Value:3,
			Right:&Tree{Value:5},
		},
		Value:8,
		Right:&Tree{Value:13},
	}
	/*Walk_recursion(&tree2)
	fmt.Println()*/

	ch = make(chan int)
	go Walk(&tree2, ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println(v)
	}

	fmt.Println("tree1 == tree2 is ",Same(&tree1,&tree2))

}