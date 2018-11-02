package analyzer

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

const iosMessage = "[2017-11-21, 8:41:01 PM] Some Really Long Name: A really cool and long iosMessage to test for!"
const androidMessage = "11/21/17, 8:41 PM - Some Really Long Name: A really cool and long androidMessage to test for!"

var ios, android = iosParser{}, androidParser{}

func TestIosParser_Date(t *testing.T) {
	d := string(ios.Date(iosMessage))
	if d != "Nov 2017" {
		t.Fatalf("Date is %v, should be %v", d, "Nov 2017")
	}
}

func TestIosParser_Message(t *testing.T) {
	m := string(ios.Message(iosMessage))
	if m != "A really cool and long iosMessage to test for!" {
		t.Fatalf("Message is \"%v\", should be \"%v\"", m, "A really cool and long iosMessage to test for!")
	}
}

func TestIosParser_Hour(t *testing.T) {
	d := ios.Hour(iosMessage)
	if string(d) != "8:00 PM" {
		t.Fatalf("Hour is %v, should be %v", d, 20)
	}
}

func TestIosParser_Sender(t *testing.T) {
	s := string(ios.Sender(iosMessage))
	if s != "Some Really Long Name" {
		t.Fatalf("Name is %v, should be %v", s, "some Really Long Name")
	}
}

func TestAndroidParser_Date(t *testing.T) {
	d := string(android.Date(androidMessage))
	if d != "Nov 2017" {
		t.Fatalf("Date is %v, should be %v", d, "Nov 2017")
	}
}

func TestAndroidParser_Message(t *testing.T) {
	m := string(android.Message(androidMessage))
	if m != "A really cool and long androidMessage to test for!" {
		t.Fatalf("Message is \"%v\", should be \"%v\"", m, "A really cool and long androidMessage to test for!")
	}
}

func TestAndroidParser_Hour(t *testing.T) {
	d := android.Hour(androidMessage)
	if string(d) != "8:00 PM" {
		t.Fatalf("Hour is %v, should be %v", d, 20)
	}
}

func TestAndroidParser_Sender(t *testing.T) {
	s := string(android.Sender(androidMessage))
	if s != "Some Really Long Name" {
		t.Fatalf("Name is %v, should be %v", s, "some Really Long Name")
	}
}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
