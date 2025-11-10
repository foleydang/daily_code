# daily_code

## 项目目的

**daily_code** 是一个专注于日常算法练习和编程技能提升的项目，旨在通过持续的代码实践来强化算法思维、数据结构理解和编程能力。该项目主要包含：

- **算法实现**：常见算法（如排序、搜索、动态规划等）的多种实现方式
- **LeetCode 题目解答**：涵盖 LeetCode Hot 100 等经典题目的详细解答和优化思路
- **多语言支持**：主要使用 Go 语言开发，同时包含 Java 相关实现
- **每日练习记录**：按日期组织的代码练习，方便追踪学习进度和复习

## 项目结构

```
├── .gitignore        # Git 忽略文件配置
├── README.md         # 项目说明文档
├── _2025/           # 2025年的代码练习
│   └── test.go      # 测试代码
├── go.mod           # Go 模块定义文件
├── java/            # Java 相关代码
│   └── src/         # Java 源代码目录
│       └── year2025/ # 2025年的Java练习
└── main.go          # 项目主入口文件
```

## 技术栈

- **主要语言**：Go 1.24.7
- **辅助语言**：Java
- **开发工具**：支持 Go 和 Java 的 IDE（如 IntelliJ IDEA、VS Code 等）

## 使用方法

### 运行 Go 代码

1. 确保已安装 Go 1.24.7 或更高版本
2. 克隆项目到本地：
   ```bash
   git clone https://github.com/foleydang/daily_code.git
   cd daily_code
   ```
3. 运行主程序：
   ```bash
   go run main.go
   ```
4. 运行特定日期的练习代码：
   ```bash
   go run _2025/1013.go
   ```

### 运行测试

项目包含简单的测试功能，可以验证算法实现的正确性：

```bash
# 运行主程序中的测试
go run main.go
```

## 项目内容

### 已实现的算法

- **排序算法**：快速排序（quickSort）
- **动态规划**：斐波那契数列（fibonacci）
- **数组操作**：两数之和（twoSum）

### LeetCode Hot 100 题目列表

#### 数组与字符串
1. [两数之和](https://leetcode-cn.com/problems/two-sum/) - 简单
2. [两数相加](https://leetcode-cn.com/problems/add-two-numbers/) - 中等
3. [无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/) - 中等
4. [寻找两个正序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/) - 困难
5. [最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/) - 中等
6. [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/) - 中等
7. [三数之和](https://leetcode-cn.com/problems/3sum/) - 中等
8. [删除有序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/) - 简单
9. [移除元素](https://leetcode-cn.com/problems/remove-element/) - 简单
10. [实现 strStr()](https://leetcode-cn.com/problems/implement-strstr/) - 简单
11. [最大子数组和](https://leetcode-cn.com/problems/maximum-subarray/) - 简单
12. [买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/) - 简单
13. [买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/) - 简单
14. [跳跃游戏](https://leetcode-cn.com/problems/jump-game/) - 中等
15. [合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/) - 简单
16. [旋转数组](https://leetcode-cn.com/problems/rotate-array/) - 中等
17. [买卖股票的最佳时机 III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/) - 困难
18. [买卖股票的最佳时机 IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/) - 困难
19. [接雨水](https://leetcode-cn.com/problems/trapping-rain-water/) - 困难
20. [删除排序数组中的重复项 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/) - 中等

#### 链表
21. [反转链表](https://leetcode-cn.com/problems/reverse-linked-list/) - 简单
22. [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/) - 简单
23. [环形链表](https://leetcode-cn.com/problems/linked-list-cycle/) - 简单
24. [环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/) - 中等
25. [删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/) - 中等
26. [两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/) - 中等
27. [K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/) - 困难
28. [随机链表的复制](https://leetcode-cn.com/problems/copy-list-with-random-pointer/) - 中等
29. [相交链表](https://leetcode-cn.com/problems/intersection-of-two-linked-lists/) - 简单
30. [回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/) - 简单
31. [排序链表](https://leetcode-cn.com/problems/sort-list/) - 中等
32. [LRU 缓存](https://leetcode-cn.com/problems/lru-cache/) - 中等

### 树
33. [二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/) - 简单
34. [二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/) - 简单
35. [二叉树的后序遍历](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/) - 简单
36. [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/) - 中等
37. [对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/) - 简单
38. [二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/) - 简单
39. [二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/) - 简单
40. [翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/) - 简单
41. [路径总和](https://leetcode-cn.com/problems/path-sum/) - 简单
42. [路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/) - 中等
43. [从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) - 中等
44. [从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) - 中等
45. [验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/) - 中等
46. [二叉搜索树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) - 简单
47. [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/) - 中等
48. [二叉搜索树中的插入操作](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/) - 中等
49. [删除二叉搜索树中的节点](https://leetcode-cn.com/problems/delete-node-in-a-bst/) - 中等
50. [平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/) - 简单
51. [二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/) - 中等
52. [将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/) - 简单
53. [不同的二叉搜索树](https://leetcode-cn.com/problems/unique-binary-search-trees/) - 中等
54. [不同的二叉搜索树 II](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/) - 中等
55. [填充每个节点的下一个右侧节点指针](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/) - 中等
56. [填充每个节点的下一个右侧节点指针 II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/) - 中等
57. [最大二叉树](https://leetcode-cn.com/problems/maximum-binary-tree/) - 中等
58. [二叉搜索树的最小绝对差](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/) - 简单

### 动态规划
59. [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/) - 简单
60. [打家劫舍](https://leetcode-cn.com/problems/house-robber/) - 中等
61. [打家劫舍 II](https://leetcode-cn.com/problems/house-robber-ii/) - 中等
62. [零钱兑换](https://leetcode-cn.com/problems/coin-change/) - 中等
63. [最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/) - 中等
64. [最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/) - 中等
65. [编辑距离](https://leetcode-cn.com/problems/edit-distance/) - 困难
66. [买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) - 中等
67. [单词拆分](https://leetcode-cn.com/problems/word-break/) - 中等
68. [乘积最大子数组](https://leetcode-cn.com/problems/maximum-product-subarray/) - 中等
69. [跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/) - 中等
70. [不同路径](https://leetcode-cn.com/problems/unique-paths/) - 中等
71. [不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/) - 中等
72. [最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/) - 中等
73. [三角形最小路径和](https://leetcode-cn.com/problems/triangle/) - 中等
74. [最大矩形](https://leetcode-cn.com/problems/maximal-rectangle/) - 困难
75. [最大正方形](https://leetcode-cn.com/problems/maximal-square/) - 中等
76. [分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/) - 中等
77. [目标和](https://leetcode-cn.com/problems/target-sum/) - 中等
78. [一和零](https://leetcode-cn.com/problems/ones-and-zeroes/) - 中等

### 回溯算法
79. [全排列](https://leetcode-cn.com/problems/permutations/) - 中等
80. [全排列 II](https://leetcode-cn.com/problems/permutations-ii/) - 中等
81. [子集](https://leetcode-cn.com/problems/subsets/) - 中等
82. [子集 II](https://leetcode-cn.com/problems/subsets-ii/) - 中等
83. [组合总和](https://leetcode-cn.com/problems/combination-sum/) - 中等
84. [组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/) - 中等
85. [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/) - 中等
86. [括号生成](https://leetcode-cn.com/problems/generate-parentheses/) - 中等
87. [解数独](https://leetcode-cn.com/problems/sudoku-solver/) - 困难
88. [N 皇后](https://leetcode-cn.com/problems/n-queens/) - 困难
89. [N 皇后 II](https://leetcode-cn.com/problems/n-queens-ii/) - 困难

### 排序与查找
90. [二分查找](https://leetcode-cn.com/problems/binary-search/) - 简单
91. [在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/) - 中等
92. [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/) - 中等
93. [搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/) - 中等
94. [寻找峰值](https://leetcode-cn.com/problems/find-peak-element/) - 中等
95. [合并区间](https://leetcode-cn.com/problems/merge-intervals/) - 中等
96. [前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/) - 中等
97. [数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/) - 中等

### 其他
98. [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/) - 中等
99. [跳跃游戏](https://leetcode-cn.com/problems/jump-game/) - 中等
100. [最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/) - 困难

## 学习目标

通过参与本项目，您将能够：

- 掌握常见算法和数据结构的实现原理
- 提高解决复杂编程问题的能力
- 熟悉 LeetCode 等平台的经典题目
- 培养良好的代码风格和编程习惯
- 提升算法思维和逻辑推理能力

## 贡献指南

欢迎对本项目进行贡献！如果您有兴趣参与：

1. Fork 本项目
2. 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 联系方式

如有问题或建议，欢迎通过以下方式联系：

- GitHub Issues: [提交 Issue](https://github.com/yourusername/daily_code/issues)

---

**持续学习，每日进步！** 🚀
