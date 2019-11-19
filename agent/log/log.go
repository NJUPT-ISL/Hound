package log

import (
	"fmt"
	"time"
)

func ErrPrint(err error) {
	fmt.Printf("[Hound Agent] - [%s] \"Error: %s\"\n", time.Now(), err)
}

func Print(log string) {
	fmt.Printf("[Hound Agent] - [%s] \" %s\"\n", time.Now(), log)
}
