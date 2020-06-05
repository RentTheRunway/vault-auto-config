package state

// Represents a tree-like structure of the current configuration of vault
type ConfigState struct {
	Children map[string]*Node `yaml:"children,omitempty"`
}

type Nodes = map[string]*Node

type Node struct {
	Children map[string]*Node `yaml:"children,omitempty"`
	Config interface{} `yaml:"config,omitempty"`
	Parent *Node `yaml:"-"`
}

func NewConfigState() *ConfigState {
	return &ConfigState{
		Children: make(map[string]*Node),
	}
}

func NewNode() *Node {
	return &Node{
		Children: make(map[string]*Node),
	}
}

func (n *Node) AddNode(name string) *Node {
	node := NewNode()
	node.Parent = n
	n.Children[name] = node
	return node
}
