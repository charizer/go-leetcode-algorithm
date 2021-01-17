package lowest_common_ancestor_of_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// 根比两个节点都要大，则最近公共祖先在左子树中
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}else if root.Val < p.Val && root.Val < q.Val {
		// 根比两个节点都要小，则最近公共祖先在右子树中
		return lowestCommonAncestor(root.Right, p, q)
	}else{
		// 否则p,q分散在root的左、右子树中，则root即为最近公共祖先
		return root
	}
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		}else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		}else{
			return root
		}
	}
}
