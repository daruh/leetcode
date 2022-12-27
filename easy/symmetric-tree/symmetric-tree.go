package same_tree

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [1,2,2,3,4,4,3,5,6,7,8,]
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	isSymmetricFlag := true
	bstList := []*TreeNode{root}
	for len(bstList) > 0 && isSymmetricFlag {
		var bstBranch []*TreeNode
		allNils := true
		for i := 0; i < len(bstList); i++ {
			node := bstList[i]
			if node == nil {
				continue
			}
			if node.Left != nil || node.Right != nil {
				allNils = false
			}
			bstBranch = append(bstBranch, node.Left, node.Right)
		}

		if len(bstBranch) == 0 || allNils {
			break
		}
		i := 0
		j := len(bstBranch) - 1
		for i < j {
			if bstBranch[i] == nil && bstBranch[j] == nil {
				i++
				j--
				continue
			}
			if bstBranch[i] != nil && bstBranch[j] != nil {
				if bstBranch[i].Val != bstBranch[j].Val {
					isSymmetricFlag = false
					break
				}
			}
			if (bstBranch[i] == nil && bstBranch[j] != nil) || (bstBranch[i] != nil && bstBranch[j] == nil) {
				isSymmetricFlag = false
				break
			}
			i++
			j--
		}

		bstList = bstBranch
	}
	return isSymmetricFlag

}
