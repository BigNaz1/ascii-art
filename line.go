package ascii

import (
	"fmt"
	"bufio"
)


func Line(num int) string {
	str := ""
	f, e := os.Open("standard.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()

	f.Seek(0, 0)
	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		str, _ = content.ReadString('\n')
	}
	str = strings.TrimSuffix(str, "\n")
	return str
}