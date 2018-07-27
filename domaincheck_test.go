package domaincheck

import (
	"testing"
)

var (
	validNames = []string{
		"example.com",
		"good.example.com",
		"*.good.example.com",
		"*.good-example.com",
		"www.example.com",
		"*.xn--rger-koa.example.com",
	}

	invalidNames = []string{
		"",
		".",
		"*.",
		"@",
		"*",
		"example.123~.com",
		"example.().com",
		"example.<>.com",
		"example.!@#$%^&.com",
		"a@",
		"@a",
		"a.*",
		"*.a",
		"*..",
		"*.com",
		"*.*.example.com",
		"bad.*.example.com",
		"*.-bad-example.com",
		"*.bad-example-.too",
		"bad.   .example.com",
		"bad. www  .example.com",
		"bad.www\t.example.com",
		"bad.with spaces in middle.com",
		"..invalid",
		"invalid-too",
		"com",
		"*.invalid",
		"also-invalid..",
		"192.168.0.1",
		" 192.168.0.1 ",
		" 2001:db8::2:1 ",
	}

	validWildcards = []string{
		"*.good-wildcard.com",
		"*.сильныйцветок.рф",
	}

	invalidWildcards = []string{
		"example.com",
		"good.example.com",
		"*.",
		"*..",
		"*.com",
		"*.*.example.com",
		"good.*.example.com",
		"*.-bad-example.com",
		"*.bad-example-.too",
		"..invalid",
		"invalid-too",
		"*.invalid",
		"also-invalid..",
	}
)

func TestDomainCheck(t *testing.T) {
	for _, domain := range validNames {
		if Valid(domain) != true {
			t.Fatal("fail a good case", domain)
		}
	}

	for _, domain := range invalidNames {
		if Valid(domain) != false {
			t.Fatal("missed a bad case", domain)
		}
	}
}

func TestWildcard(t *testing.T) {
	for _, domain := range validWildcards {
		if ValidWildcard(domain) != true {
			t.Fatal("fail a good case", domain)
		}
	}

	for _, domain := range invalidWildcards {
		if ValidWildcard(domain) != false {
			t.Fatal("missed a bad case", domain)
		}
	}
}

func TestStemDomain(t *testing.T) {
	for _, domain := range validNames {
		if ValidWildcard(domain) {
			if StemDomain(domain) != domain[2:] {
				t.Fatal("fail a good wildcard case", domain)
			}
		} else if StemDomain(domain) != domain {
			t.Fatal("fail a good non-wildcard case", domain)
		}
	}

	for _, domain := range invalidWildcards {
		if Valid(domain) {
			if StemDomain(domain) != domain {
				t.Fatal("fail a good non-wildcard case", domain)
			}
		} else if StemDomain(domain) != "" {
			t.Fatal("missed a bad case", domain)
		}
	}

	for _, domain := range invalidNames {
		if StemDomain(domain) != "" {
			t.Fatal("missed a bad case", domain)
		}
	}
}

func TestPunycodeName(t *testing.T) {
	for _, domain := range invalidNames {
		if PunycodeName(domain) != "" {
			t.Fatal("bad puny code conversion, didn't catch bad name", domain)
		}
	}

	p1 := PunycodeName("*.сильныйцветок.рф")

	if p1 != "*.xn--b1aghehcoluq5b8cl.xn--p1ai" {
		t.Fatal("bad puny code conversion", p1)
	}

	p2 := PunycodeName("随便.com")
	if p2 != "xn--wtq493o.com" {
		t.Fatal("bad puny code conversion")
	}

	p3 := PunycodeName("example.com")
	if p3 != "example.com" {
		t.Fatal("bad puny code conversion")
	}
}
