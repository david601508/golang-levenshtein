package diff

type EditOp struct {
	Action byte
	Src    string
	Dst    string
	SrcPos int
	DstPos int
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func LevenshteinDistance(str1, str2 string) int {
	len1 := len(str1)
	len2 := len(str2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[len1][len2]
}

func LevenshteinEditOps(str1, str2 string) []EditOp {
	len1 := len(str1)
	len2 := len(str2)

	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	ops := make([]EditOp, 0)
	i, j := len1, len2
	for i > 0 || j > 0 {
		if i > 0 && dp[i][j] == dp[i-1][j]+1 {
			ops = append(ops, EditOp{Action: 'D',
				Src:    string(str1[i-1]),
				Dst:    "",
				SrcPos: i - 1,
				DstPos: j})
			i--
		} else if j > 0 && dp[i][j] == dp[i][j-1]+1 {
			ops = append(ops, EditOp{Action: 'I',
				Src:    "",
				Dst:    string(str2[j-1]),
				SrcPos: i,
				DstPos: j - 1})
			j--
		} else if i > 0 && j > 0 {
			if dp[i][j] == dp[i-1][j-1]+1 {
				ops = append(ops, EditOp{Action: 'R',
					Src:    string(str1[i-1]),
					Dst:    string(str2[j-1]),
					SrcPos: i - 1,
					DstPos: j - 1})
				i--
				j--
			} else {
				i--
				j--
			}
		}
	}

	reverseOps := make([]EditOp, len(ops))
	for i, j := 0, len(ops)-1; j >= 0; i, j = i+1, j-1 {
		reverseOps[i] = ops[j]
	}

	return reverseOps
}
