package Composite

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	c.SetName(name)
	return c
}


const (
	LeafNode = iota
	CompositeNode
)

type component struct {
	parent Component
	name string
}


func(c*component) Parent() Component {
	return c.parent
}
func(c*component)SetParent(parent Component) {
	c.parent = parent

}
func(c*component)Name() string {
	return c.name
}
func(c*component)SetName(name string) {
	c.name = name
}
func(c*component)AddChild(Component) {

}
func(c*component)Print(string) {

}

