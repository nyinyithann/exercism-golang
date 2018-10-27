package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Parent   int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if records[0].ID != records[0].Parent {
		return nil, fmt.Errorf("Root node should not have parent")
	}

	prev := records[0].ID
	for i := 1; i < len(records); i++ {
		if records[i].ID <= records[i].Parent {
			return nil, fmt.Errorf("ID should be greater than Parent ID")
		}
		if records[i].ID-prev != 1 {
			return nil, fmt.Errorf("Non-Continuous")
		}
		prev++
	}

	dupe := make(map[int]*Record)
	for _, r := range records {
		temp := dupe[r.ID]
		if temp == nil {
			dupe[r.ID] = &r
			continue
		}
		if temp.Parent == r.Parent {
			return nil, fmt.Errorf("Node is a duplicate - {ID: %d, Parent: %d}", r.ID, r.Parent)
		}
	}

	root := &Node{ID: records[0].ID}

	nodes := make([]*Node, 0, len(records)-1)
	for _, n := range records[1:] {
		nodes = append(nodes, &Node{ID: n.ID, Parent: n.Parent})
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Parent >= nodes[j].Parent
	})

	for _, n := range nodes {

		idx := sort.Search(len(nodes), func(i int) bool {
			return nodes[i].ID <= n.Parent
		})
		if idx != len(nodes) {
			n.Parent = 0
			nodes[idx].Children = append(nodes[idx].Children, n)
			sort.Slice(nodes[idx].Children, func(i, j int) bool {
				return nodes[idx].Children[i].ID <= nodes[idx].Children[j].ID
			})
		} else if n.Parent == root.ID {
			n.Parent = 0
			root.Children = append(root.Children, n)
		} else {
			return nil, fmt.Errorf("Orphan node found")
		}
	}
	sort.Slice(root.Children, func(i, j int) bool {
		return root.Children[i].ID <= root.Children[j].ID
	})
	return root, nil
}
