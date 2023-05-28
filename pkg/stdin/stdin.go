package stdin

import (
	"bufio"
	"os"
)

func StrStdin() string {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return scanner.Text()
}