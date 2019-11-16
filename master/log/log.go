package log

import (
	"fmt"
	"time"
)

func ErrPrint(err error) string {
	return fmt.Sprintf("[Hound Master] - [%s] \"Error: %s\"\n", time.Now(), err)
}

func Print(log string) string {
	return fmt.Sprintf("[Hound Master] - [%s] \": %s\"\n", time.Now(), log)
}
