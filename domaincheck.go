package domaincheck

import (
	"net"
	"strings"
	"unicode"

	"golang.org/x/net/idna"
)

// Valid return true if domain is valid
func Valid(domain string) bool {
	trimmed := strings.TrimSpace(domain)
	if len(trimmed) < len(domain) {
		return false // reject domain name with leading or trailing spaces
	}
	// reject IP address
	if net.ParseIP(domain) != nil {
		return false
	}

	// strip "*." prefix if exists
	if len(domain) >= 2 && domain[:2] == "*." {
		domain = domain[2:]
	}

	// should not see any special characters any more
	if strings.ContainsAny(domain, "_~`!@#$%^&*()=+{}[]|\\;:'\",<>/?") {
		return false
	}

	// should not contain spaces anywhere
	if strings.Contains(domain, " ") {
		return false
	}

	tokens := strings.Split(domain, ".")
	// should have at least two tokens
	if len(tokens) < 2 {
		return false
	}

	for _, token := range tokens {
		trimmed := strings.TrimSpace(token) // trim space-filled token
		if len(trimmed) < len(token) {
			return false // reject any token containing space
		}

		if len(token) == 0 {
			return false // consecutive '.' is forbidden
		}

		// token begins or ends with '-' is bad
		if token[0] == '-' || token[len(token)-1] == '-' {
			return false
		}
	}
	for _, char := range domain {
		if !unicode.IsPrint(char) {
			return false
		}
	}

	return true
}

// ValidWildcard returns true if domain is a valid wildcard one
func ValidWildcard(domain string) bool {
	// valid domain should be at least 3 characters long.
	if Valid(domain) && domain[:2] == "*." {
		return true
	}
	return false
}

// StemDomain returns the stem domain with wildcard prefix stripped (if any).
func StemDomain(domain string) string {
	if ValidWildcard(domain) {
		return domain[2:]
	}

	if Valid(domain) {
		return domain
	}

	return ""
}

// PunycodeName returns a punycoded domain name, with wildcard properly prefixed.
func PunycodeName(domain string) string {
	// only convert stem domain
	stem := StemDomain(domain)
	puny, err := idna.ToASCII(stem) // get Punycoded domain name
	if err != nil {
		return ""
	}

	// treat wildcard
	if puny != "" && ValidWildcard(domain) {
		puny = "*." + puny
	}

	return puny
}
