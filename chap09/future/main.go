package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	SuccessFunc       func(string)
	FailFunc          func(error)
	ExecuteStringFunc func() (string, error)

	MaybeString struct {
		successFunc SuccessFunc
		failFunc    FailFunc
	}
)

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) *MaybeString {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
	return s
}

func main() {
	future := &MaybeString{}
	var wg sync.WaitGroup
	wg.Add(1)

	future.Execute(func() (string, error) {
		fmt.Println("This is message in Execute function.")
		time.Sleep(5 * time.Second)
		return "dsfgdhjkhgf", nil
	}).Success(func(str string) {
		fmt.Println("This is success message in Success function ", str)
		wg.Done()
	}).Fail(func(err error) {
		fmt.Println("This is fail message in Failure function.")
		wg.Done()
	})

	fmt.Println("This is waiting guy....")
	wg.Wait()
}
