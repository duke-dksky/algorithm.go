
### 规律

- 第k层有k个元素
- 每层第一个以及最后一个元素值为1
- 对于第k(k > 2)层第n(n > 1 && n < k)个元素A[k][n],A[k][n] = A[k-1][n-1] + A[k-1][n]