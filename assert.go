package assert

import (
	"fmt"
	"os"
	"reflect"
	"runtime/debug"
	"syscall"
)

var DefaultMode = Panic

const (
	Panic PanicMode = iota
	Exit
	SIGTERM
	SKIP
)

type AssertionPanic string
type PanicMode int

func log(err error, arg ...any) {

	s := prettyStack{}

	stackTrace := debug.Stack()
	out, err := s.parse(stackTrace, err, arg...)

	if err == nil {
		os.Stderr.Write(out)
	} else {
		// print stdlib output as a fallback
		os.Stderr.Write(stackTrace)
	}
}

func handlePanic(mode PanicMode) {
	switch mode {
	case Panic:
		panic(AssertionPanic("ASSERTION FAILED"))
	case Exit:
		os.Exit(1)
	case SIGTERM:
		// err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
		err := sigterm()
		if err != nil {
			fmt.Fprintf(os.Stderr, "PANIC: %s\n", err)
			os.Exit(1)
		}

		// still kill the current go routine
		panic(AssertionPanic("SYSCALL SIGTERM PANIC"))
	}
}

func sigterm() error {

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		return err
	}

	return p.Signal(syscall.SIGTERM)
}

func NoErr(err error, mode PanicMode, arg ...any) {
	if err != nil {

		log(err, arg...)
		handlePanic(mode)
	}
}

func AssertBool(expected, got bool, mode PanicMode, arg ...any) {
	if expected != got {
		err := fmt.Errorf("assertion err: expected %t got %t", expected, got)
		log(err, arg...)
		handlePanic(mode)
	}
}

func NotNil(intf any, mode PanicMode, arg ...any) {

	if IsNill(intf) {
		log(nil, arg...)
		handlePanic(mode)
	}
}

// IsNill returns true intf is nil
func IsNill(intf any) bool {
	if intf == nil {
		return true
	}

	if canIsNil(intf) {
		return reflect.ValueOf(intf).IsNil()
	}

	return false
}

func canIsNil(intf any) bool {

	v := reflect.ValueOf(intf)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
		fallthrough
	case reflect.Interface, reflect.Slice:
		return true
	default:
		return false
	}

}

func StrNotEmpty(str string, mode PanicMode, arg ...any) {
	if str == "" {
		log(nil, arg...)
		handlePanic(mode)
	}
}

// Must assert err is not nil with
// [DefaultMode]
func Must[T any](v T, err error) T {
	NoErr(err, DefaultMode)
	return v
}

// Must assert err is not nil with
// [DefaultMode]
func MustOk[T any](v T, ok bool) T {
	if !ok {
		log(nil)
		handlePanic(DefaultMode)
	}
	return v
}
