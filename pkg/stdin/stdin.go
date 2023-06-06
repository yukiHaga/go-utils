package stdin

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func StrStdin() string {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return scanner.Text()
}

// Atoiは２値を返すので、そこのハンドリングを使う側でする
func IntStdin() (int, error) {
    stringInput := StrStdin()
    return strconv.Atoi(strings.TrimSpace(stringInput))
}

