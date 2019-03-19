import java.util.ArrayList;
import java.util.List;

/**
 * @author terry.guo
 *
 */
public class PreorderTraversalSolution {

	/**
	 * @param args
	 */
	public List<Integer> preorderTraversal(TreeNode root) {
		
		List<Integer> l = new ArrayList<Integer>();
		if(root == null) {
			return l;
		}
		l.add(root.val);
		if(root.left != null) {
			l.addAll(preorderTraversal(root.left));
		}
		if(root.right!= null) {
			l.addAll(preorderTraversal(root.right));
		}
		
		return l;

	}
}
