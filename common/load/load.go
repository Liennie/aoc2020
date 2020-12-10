package load

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Res struct {
	Line string
	Err  error
}

func File(filename string) <-chan Res {
	ch := make(chan Res)

	go func() {
		defer close(ch)

		file, err := os.Open(filename)
		if err != nil {
			ch <- Res{"", err}
			return
		}
		defer file.Close()

		r := bufio.NewReader(file)
		for {
			l, err := r.ReadString('\n')
			if len(l) > 0 {
				ch <- Res{strings.TrimSuffix(l, "\n"), nil}
			}
			if err != nil {
				if err != io.EOF {
					ch <- Res{"", fmt.Errorf("ReadString: %w", err)}
				}
				return
			}
		}
	}()

	return ch
}
