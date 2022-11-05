package main

import "fmt"

type (
	// Department is the handler interface
	Department interface {
		execute(*Patient)
		setNext(Department)
	}

	// Reception is concrete handler
	Reception struct {
		next Department
	}

	Doctor struct {
		next Department
	}

	Medical struct {
		next Department
	}

	Cashier struct {
		next Department
	}

	Patient struct {
		name              string
		registrationDone  bool
		doctorCheckUpDone bool
		medicineDone      bool
		paymentDone       bool
	}
)

// Reception's collection of methods
func (s *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration is already done")
		s.next.execute(p)
		return
	}

	fmt.Println("Reception registering patient")
	p.registrationDone = true
	s.next.execute(p)
}

func (s *Reception) setNext(next Department) {
	s.next = next
}

// Doctor's collection of methods
func (s *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup is already done")
		s.next.execute(p)
		return
	}

	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	s.next.execute(p)
}

func (s *Doctor) setNext(next Department) {
	s.next = next
}

// Medical's collection of methods
func (s *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine is already given")
		s.next.execute(p)
		return
	}

	fmt.Println("Medical giving medicine")
	p.medicineDone = true
	s.next.execute(p)
}

func (s *Medical) setNext(next Department) {
	s.next = next
}

// Cashier's collection of methods
func (s *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment is already done")
		return
	}

	fmt.Println("Cashier collecting payment")
	p.paymentDone = true
}

func (s *Cashier) setNext(next Department) {
	s.next = next
}

// main function
func main() {
	cashier := new(Cashier)

	// set next for medical department
	medical := new(Medical)
	medical.setNext(cashier)

	// set next for doctor department
	doctor := new(Doctor)
	doctor.setNext(medical)

	// set next for reception department
	reception := new(Reception)
	reception.setNext(doctor)

	patient := &Patient{name: "John"}
	reception.execute(patient)
}
