package stack

import "testing"

func TestBrowserHistory_To(t *testing.T) {
	history := NewBrowserHistory()
	history.To("11111111111")
	history.To("22222222222")
	history.To("33333333333")
	history.Print()

	history.Back()
	history.Print()

	history.Goahead()
	history.Goahead()
	history.Print()

	history.To("44444444444")
	history.Print()
}
