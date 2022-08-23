package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(restoreIpAddresses("101023"))
}

func restoreIpAddresses(s string) (res []string) { //nolint:revive,stylecheck
	if len(s) < 4 || len(s) > 12 {
		return
	}

	dfs(s, 0, "", &res)

	return
}

const maxBlockNum = 4

func dfs(s string, block int, ipTmp string, result *[]string) {
	if block > maxBlockNum {
		return
	}

	if block == maxBlockNum {
		if len(s) == 0 {
			*result = append(*result, ipTmp)
		}

		return
	}

	if block != 0 {
		ipTmp += "."
	}

	for i := 1; i < 4 && i <= len(s); i++ {
		chunk := s[:i]

		if chunk == "0" || strToInt(chunk) < 256 {
			dfs(s[i:], block+1, ipTmp+chunk, result)
		}

		if chunk == "0" {
			return
		}
	}
}

const ASCIINumOffset = 48

// as we do know that input is digits-only were able to freak out a bit.
func strToInt(str string) (res int) {
	for i := 0; i < len(str); i++ {
		res += int(str[i]-ASCIINumOffset) * int(math.Pow10(len(str)-i-1))
	}

	return
}
