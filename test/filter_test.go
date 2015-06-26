package goosmtools

import (
	"github.com/gaffo/goosm"
	"testing"
)

func fail(tb testing.TB, reason string) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%d: %s\n\n", filepath.Base(file), line, reason)
	tb.FailNow()
}

func contains(tb testing.TB, s []string, e string) {
	for _, a := range s {
		if a == e {
			return
		}
	}
	fmt.Printf("Expected to contain %s\n", e)
	tb.FailNow()
}

// assert fails the test if the condition is false.
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: "+msg+"\n\n", append([]interface{}{filepath.Base(file), line}, v...)...)
		tb.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}