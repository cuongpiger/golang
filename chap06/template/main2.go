package main

import "strings"

type (
	MessageRetriever interface {
		Message() string
	}

	Templater interface {
		first() string
		third() string
		ExecuteAlgorithm(MessageRetriever) string
	}
)

type (
	Template          struct{}
	AnonymousTemplate struct{}
	adapter           struct {
		myFunc func() string
	}
	TestStruct struct {
		Template
	}
)

func (t *Template) first() string {
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return t.first() + " " + m.Message() + " " + t.third()
}

func (a *AnonymousTemplate) first() string {
	return "hello"
}

func (a *AnonymousTemplate) third() string {
	return "template"
}

func (t *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{t.first(), f(), t.third()}, " ")
}

func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}

	return ""
}

func (m *TestStruct) Message() string {
	return "world"
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}
