
### 定义

中位数将一个集合分为长度相等的两个子集合，其中一个子集合的元素都大于另一个子集合。

- 条件1：`len(left_part)=len(right_part)`
- 条件2：`max(left_part) ≤ min(right_part)`

则可以得到中位数`Median = (max(left_part) + min(right_part)) / 2`


left_A = A[0]....A[i - 1] 长度为 i
right_A = A[i]....A[m - 1] 长度为 m - i
left_B = B[0]....B[j - 1] 长度为 j
right_B = B[j]....B[n - 1] 长度为 n - j

- 条件1：i + j == m + n - i - j
- 条件2：A[i - 1] <= B[j] && B[j - 1] <= A[i]
