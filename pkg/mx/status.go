package mx

import "strings"

type Status int

const (
	OK Status = iota
	IMPAIRED
	FAILED
	UNKNOWN
)

var statusText = map[Status]string{
	OK:       "ok",
	IMPAIRED: "impaired",
	FAILED:   "failed",
	UNKNOWN:  "unknown",
}

func (s Status) String() string {
	return strings.ToUpper(statusText[s])
}

func AsStatus(s string) Status {
	s = strings.ToLower(s)
	for k, v := range statusText {
		if v == s {
			return k
		}
	}
	return UNKNOWN
}
