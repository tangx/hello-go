package main

import (
	"net/url"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	// password = `P_-ssw(a)rd23`
	// password = `!()_-~.`
	password = `_-~.`
)

func TestMain(t *testing.T) {

	s := url.PathEscape(password)
	t.Run("url.PathEscape", func(t *testing.T) {
		NewWithT(t).Expect(s).To(Equal(password))
	})

	s = url.QueryEscape(password)
	t.Run("url.QueryEscape", func(t *testing.T) {
		NewWithT(t).Expect(s).To(Equal(password), "wooa, it's same")
	})

}
