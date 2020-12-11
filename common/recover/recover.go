package recover

import (
	"fmt"
)

func Err(f func(err error)) {
	if e := recover(); e != nil {
		switch e.(type) {
		case error:
			f(e.(error))
		default:
			f(fmt.Errorf("%v", e))
		}
	}
}
