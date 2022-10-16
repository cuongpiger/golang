package main

import (
	"testing"
)

func TestSubject(t *testing.T) {
	testObserver1 := &TestObserver{1, "default"}
	testObserver2 := &TestObserver{2, "default"}
	testObserver3 := &TestObserver{3, "default"}

	publisher := Publisher{}

	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(testObserver1)
		publisher.AddObserver(testObserver2)
		publisher.AddObserver(testObserver3)

		if len(publisher.ObserverList) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(testObserver2)

		if len(publisher.ObserverList) != 2 {
			t.Errorf("The size of the observer list is not the "+
				"expected. 3 != %d\n", len(publisher.ObserverList))
		}

		for _, observer := range publisher.ObserverList {
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
			}

			if testObserver.ID == 2 {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {

		if len(publisher.ObserverList) == 0 {
			t.Errorf("The list is empty. Nothing to test\n")
		}

		for _, observer := range publisher.ObserverList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if printObserver.Message != "default" {
				t.Errorf("The observer's Message field weren't"+
					" empty: %s\n", printObserver.Message)
			}
		}

		message := "Hello World!"
		publisher.NotifyObservers(message)

		for _, observer := range publisher.ObserverList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if printObserver.Message != message {
				t.Errorf("Expected message on observer %d was "+
					"not expected: '%s' != '%s'\n", printObserver.ID,
					printObserver.Message, message)
			}
		}
	})
}
