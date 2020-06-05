package state

// Represents a tree-like structure of the current configuration of vault
type ConfigState struct {
	Children map[string]*Node `yaml:"children,omitempty"`
}

// Represents a collection of child nodes in the config state
type Nodes = map[string]*Node

// Represents a node in the config state
type Node struct {
	Children map[string]*Node `yaml:"children,omitempty"`
	Config   interface{}      `yaml:"config,omitempty"`
	Parent   *Node            `yaml:"-"`
}

// Creates a new instance of ConfigState
func NewConfigState() *ConfigState {
	return &ConfigState{
		Children: make(map[string]*Node),
	}
}

// Creates a new node for ConfigState
func NewNode() *Node {
	return &Node{
		Children: make(map[string]*Node),
	}
}

// Adds a node to an existing node as a child with the given name
func (n *Node) AddNode(name string) *Node {
	node := NewNode()
	node.Parent = n
	n.Children[name] = node
	return node
}
