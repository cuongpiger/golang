package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCast(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}

	t.Logf("LOG: %s", msg)
}

func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("A payment method of type 'DebitCard' must exist")
	}

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}

	t.Logf("LOG: %s", msg)
}

func TestGetPaymentMethodNonExist(t *testing.T) {
	_, err := GetPaymentMethod(3)
	if err == nil {
		t.Fatal("A payment method with ID 3 must return an error")
	}

	t.Log("LOG:", err)
}
