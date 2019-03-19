/**
 * @author terry.gmx
 * This code need another class in the same package: TreeNode.java
 */
public class MaxDepth {
	int depth = 0;

	public int maxDepth(TreeNode root) {
		if (root == null) {
			return 0;
		}
		getDepth(root, 0);
		return depth;
	}

	public void getDepth(TreeNode node, int d) {
		d++;
		if (node.left == null && node.right == null) {
			if (depth < d) {
				depth = d;
			}
			return;
		} else {
			if (node.left != null) {
				getDepth(node.left, d);
			}
			if (node.right != null) {
				getDepth(node.right, d);
			}
		}
	}
}
