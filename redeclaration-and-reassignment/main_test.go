package redeclarationandreassignment

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func intError() (int, error) {
	return 1, errors.New("error1")
}

func stringError() (string, error) {
	return "1", errors.New("error2")
}

func Test_ErrorPointerAddress(t *testing.T) {
	i, err := intError()
	fmt.Printf("%d: %+v => %p\n", i, err, err)

	s, err := stringError()
	fmt.Printf("%s: %+v => %p\n", s, err, err)

}

func Test_Openfile(t *testing.T) {
	open("/tmp/123123.txt")
}

func open(name string) error {
	f, err := os.Open(name)
	if err != nil {
		fmt.Printf("open failed: pointer address => %p\n", err)
		// return err
	}
	d, err := f.Stat()
	if err != nil {
		fmt.Printf("sate failed: pointer address => %p\n", err)
		f.Close()
		return err
	}

	fmt.Println(d.IsDir())

	return nil
}
