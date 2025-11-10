package year2025.day1021;

import java.util.*;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int val) { this.val = val; }
}

public class Solution {
    // 最长括号长度
     public static int longestValidParentheses(String s) {
        if (s == null || s.isEmpty()) {
            return 0;
        }
        int n = s.length();
        int[] dp = new int[n];
        int max = 0;
        for (int i = 1; i < n; i++) {
            if (s.charAt(i) == ')') {
                if (s.charAt(i - 1) == '(') {
                    // 匹配到了一个有效括号对
                    dp[i] = i >= 2 ? dp[i - 2] + 2 : 2;
                } else if (i - dp[i - 1] > 0 && s.charAt(i - dp[i - 1] - 1) == '(') {
                    // 匹配到了一个有效括号对，且之前还有有效括号对
                    dp[i] = dp[i - 1] + (i - dp[i - 1] >= 2 ? dp[i - dp[i - 1] - 2] : 0) + 2;
                }
                max = Math.max(max, dp[i]);
            }
        }
        return max;
    }

    // 二叉树层次遍历
    public static List<List<Integer>> levelOrder(TreeNode root) {
       List<List<Integer>> result = new ArrayList<>();
       if (root == null){
           return result;
       }
       Queue<TreeNode> queue = new PriorityQueue<>();
       queue.offer(root);
       while(!queue.isEmpty()){
           int size = queue.size();
           List<Integer> level = new ArrayList<>();
           for (int i = 0; i < size; i++){
               TreeNode cur = queue.poll();
               level.add(cur.val);
               if (cur.left != null){
                   queue.offer(cur.left);
               }
               if (cur.right != null){
                   queue.offer(cur.right);
               }
           }
           result.add(level);
       }
       return result;
    }

    // quick sort
     public static void quickSort(int[] nums, int left, int right) {
        if (left >= right) {
            return;
        }
        int pivot = nums[left + (right - left) / 2];
        int i = left, j = right;
        while (i <= j) {
            while (nums[i] < pivot) {
                i++;
            }
            while (nums[j] > pivot) {
                j--;
            }
            if (i <= j) {
                int temp = nums[i];
                nums[i] = nums[j];
                nums[j] = temp;
                i++;
                j--;
            }
        }
        quickSort(nums, left, j);
        quickSort(nums, i, right);
    }

    public static void main(String[] args) {
        System.out.println(longestValidParentheses("(()"));
    }
}
