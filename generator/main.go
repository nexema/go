package generator

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	content, err := reader.ReadString('\n')
	if err != nil {
		println(err.Error())
		return
	}

	content = strings.TrimSpace(content)
	writer.WriteString("ok\n")
	writer.Flush()
}
