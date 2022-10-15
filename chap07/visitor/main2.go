package main

import (
	"fmt"
	"io"
	"os"
)

type (
	Visitor1 interface {
		VisitA(*MessageA)
		VisitB(*MessageB)
	}

	Visitable1 interface {
		Accept(Visitor1)
	}
)

type (
	MessageA struct {
		Msg    string
		Output io.Writer
	}

	MessageB struct {
		Msg    string
		Output io.Writer
	}

	TestHelper struct {
		Received string
	}

	MsgFieldVisitorPrinter struct{}
	MessageVisitor         struct{}
)

func (m *MessageA) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "A: %s", m.Msg)
}

func (m *MessageA) Accept(v Visitor1) {
	v.VisitA(m)
}

func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}

	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}

func (m *MessageB) Accept(v Visitor1) {
	v.VisitB(m)
}

func (mf *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited A)")
}

func (mf *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s %s", m.Msg, "(Visited B)")
}

func (mf *MsgFieldVisitorPrinter) VisitA(m *MessageA) {
	fmt.Printf(m.Msg)
}

func (mf *MsgFieldVisitorPrinter) VisitB(m *MessageB) {
	fmt.Printf(m.Msg)
}

func (t *TestHelper) Write(p []byte) (int, error) {
	t.Received = string(p)
	return len(p), nil
}
