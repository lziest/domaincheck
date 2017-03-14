package domaincheck

import (
	"strings"
)

// Valid return true if domain is valid
func Valid(domain string) bool {
	// strip "*." prefix if exists
	if len(domain) >= 2 && domain[:2] == "*." {
		domain = domain[2:]
	}

	// should not see any '*' any more
	if strings.ContainsAny(domain, "*") {
		return false
	}

	// should not contain '@'
	if strings.ContainsAny(domain, "@") {
		return false
	}

	tokens := strings.Split(domain, ".")
	// should have at least two tokens
	if len(tokens) < 2 {
		return false
	}

	for _, token := range tokens {
		if len(token) == 0 {
			return false // consecutive '.' is forbidden
		}

		// token begins or ends with '-' is bad
		if token[0] == '-' || token[len(token)-1] == '-' {
			return false
		}
	}

	return true
}

// ValidWildcard returns true if domain is a valid wildcard one
func ValidWildcard(domain string) bool {
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
