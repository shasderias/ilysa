package errors

import "fmt"

func Errorf(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}
