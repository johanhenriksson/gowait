package gowait

type Task struct {
	*Taskdef
	Cluster Cluster
	Node    Node
}

type Node struct {
	Parent *Client
}

func (t *Task) Spawn(image string) {

}
