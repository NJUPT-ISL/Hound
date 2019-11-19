package log

import (
	"fmt"
	"time"
)

func ErrPrint(err error) {
	fmt.Printf("[Hound Master] - [%s] \"Error: %s\"\n", time.Now(), err)
}

func Print(log string) {
	fmt.Printf("[Hound Master] - [%s] \" %s\"\n", time.Now(), log)
}
