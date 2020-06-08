package helpers

import (
	"text/template/parse"
)

func WalkParseTree(node parse.Node, cur []string) []string {
	if node.Type() == parse.NodeAction {
		cur = append(cur, node.String())
	}

	if ln, ok := node.(*parse.ListNode); ok {
		for _, n := range ln.Nodes {
			cur = WalkParseTree(n, cur)
		}
	}

	return cur
}
