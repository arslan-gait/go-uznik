package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	var m, n int
	scanf("%d %d\n", &m, &n)

	// creating dp array
	dp := make([][]uint8, n)
	for i := range dp {
		dp[i] = make([]uint8, m)
	}

	dp[0][0] = 1

	// creating storage array
	arr := make([][]rune, n)
	for i := range arr {
		arr[i] = make([]rune, m)
	}

	var cell rune

	for i := m - 1; i >= 0; i-- {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				scanf("%c", &cell)
				arr[j][i] = cell
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				scanf("%c", &cell)
				arr[j][i] = cell
			}
		}
		scanf("%с", &cell)
	}

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < m; j++ {
	// 		fmt.Print(string(arr[i][j]), " ")
	// 	}
	// 	fmt.Println(" ")
	// }

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] != '0' {
				// if not the first row
				if i != 0 {
					// then plus one path from the left
					dp[i][j] += dp[i-1][j]
				}
				// if not the first column
				if j != 0 {
					// then plus one path from the up
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}

	res := dp[n-1][m-1]
	if res == 0 {
		fmt.Println("impossible")
	} else {
		fmt.Println(res)
	}
}
