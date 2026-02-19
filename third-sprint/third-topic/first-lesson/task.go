package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type SliceError []error

func (errors SliceError) Error() string {
	var out []string
	for _, err := range errors {
		out = append(out, err.Error())
	}

	return strings.Join(out, `;`)
}

func MyCheck(input string) error {
	var (
		err      SliceError
		spaces   int
		hasDigit bool
	)

	if len([]rune(input)) >= 20 {
		err = SliceError{errors.New(`Line is too long`)}
	}

	for _, ch := range input {
		if ch == ' ' {
			spaces++
		} else if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
	}

	if hasDigit {
		err = append(err, errors.New(`found numbers`))
	}

	if spaces < 2 {
		err = append(err, errors.New(`no 2 spaces`))
	}

	if len(err) == 0 {
		return nil
	}

	return err
}

func main() {
	for {
		fmt.Printf("Enter string (to quit press q): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`The string passed the check`)
		}
	}
}
