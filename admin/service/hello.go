package service

import "fmt"

func Hello(name string) string {
	return fmt.Sprint("hello, ", name)
}
