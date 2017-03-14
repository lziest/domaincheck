package domaincheck

import (
	"strings"
)

func Valid(domain string) bool {
	// strip "*." prefix if exists
	if domain[:2] == "*." {
		domain = domain[2:]
	}

	// should not see any '*' any moreA
	if strings.ContainsAny(domain, "*") {
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

func ValidWildcard(domain string) bool {
	if Valid(domain) && domain[:2] == "*." {
		return true
	}
	return false
}

func StemDomain(domain string) string {
	if ValidWildcard(domain) {
		return domain[2:]
	}

	if Valid(domain) {
		return domain
	}

	return ""
}
