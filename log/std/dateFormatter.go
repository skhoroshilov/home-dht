package std

import (
	"fmt"
	"time"
)

// dateFormatter implements io.Writer interface writing dates in ISO-like format
// (YYYY-MMM-DD hh:mm:ss.zzz). It is used for formatting dates in golang log package.
type dateFormatter struct {
}

func (w dateFormatter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02 15:04:05.999") + " " + string(bytes))
}
