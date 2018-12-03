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
	d, err := ios.Date(iosMessage)
	if err != nil {
		t.Fatalf("Got error %s", err)
	}
	if string(d) != "Nov 2017" {
		t.Fatalf("Date is %v, should be %v", d, "Nov 2017")
	}
}

func TestIosParser_Message(t *testing.T) {
	m, err := ios.Message(iosMessage)
	if err != nil {
		t.Fatalf("Got error %s", err)
	}
	if string(m) != "A really cool and long iosMessage to test for!" {
		t.Fatalf("Message is \"%v\", should be \"%v\"", m, "A really cool and long iosMessage to test for!")
	}
}

func TestIosParser_Hour(t *testing.T) {
	d, err := ios.Hour(iosMessage)
	if err != nil {
		t.Fatalf("Got error %s", err)
	}
	if string(d) != "8:00 PM" {
		t.Fatalf("Hour is %v, should be %v", d, 20)
	}
}

func TestIosParser_Sender(t *testing.T) {
	s, err := ios.Sender(iosMessage)
	if err != nil {
		t.Fatalf("Got error %s", err)
	}
	if string(s) != "Some Really Long Name" {
		t.Fatalf("Name is %v, should be %v", s, "some Really Long Name")
	}
}

func TestIosParser_Valid(t *testing.T) {
	v := ios.Valid(iosMessage)

	if !v {
		t.Fatalf("Message is not valid, should be valid")
	}
}

func TestAndroidParser_Date(t *testing.T) {
	d, err := android.Date(androidMessage)
	if err != nil {
		t.Fatalf("Got error %s", err)
	}
	if string(d) != "Nov 2017" {
		t.Fatalf("Date is %v, should be %v", d, "Nov 2017")
	}
}

func TestAndroidParser_Message(t *testing.T) {
	m, err := android.Message(androidMessage)
	if err != nil {
		t.Fatalf("Got error %s", err)
	}
	if string(m) != "A really cool and long androidMessage to test for!" {
		t.Fatalf("Message is \"%v\", should be \"%v\"", m, "A really cool and long androidMessage to test for!")
	}
}

func TestAndroidParser_Hour(t *testing.T) {
	d, err := android.Hour(androidMessage)
	if err != nil {
		t.Fatalf("Got error %s", err)
	}
	if string(d) != "8:00 PM" {
		t.Fatalf("Hour is %v, should be %v", d, 20)
	}
}

func TestAndroidParser_Sender(t *testing.T) {
	s, err := android.Sender(androidMessage)
	if err != nil {
		t.Fatalf("Got error %s", err)
	}
	if string(s) != "Some Really Long Name" {
		t.Fatalf("Name is %v, should be %v", s, "some Really Long Name")
	}
}

func TestAndroidParser_Valid(t *testing.T) {
	v := android.Valid(androidMessage)
	if !v {
		t.Fatalf("Message is not valid, should be valid")
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
