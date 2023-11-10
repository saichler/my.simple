package idql_parser

import (
	"errors"
	"github.com/saichler/my.simple/go/utils/strng"
	"strings"
)

type Fetch struct {
	request     string
	elementType string
	only        []string
	criteria    *Criteria

	sort          string
	ascending     bool
	limit         int
	page          int
	caseSensitive bool

	shouldRegister   bool
	shouldUnregister bool
}

func NewFetch(request string) (*Fetch, error) {
	fetch := &Fetch{}
	fetch.request = request
	parsed := ParseFetch(request)
	_fetch := parsed[FETCH]
	_fetch = strings.TrimSpace(_fetch)
	if _fetch == "" {
		return nil, errors.New("no fetch word found in: " + request)
	}
	fetch.elementType = _fetch

	_only := parsed[ONLY]
	_only = strings.TrimSpace(_only)
	if _only != "" {
		sp := strings.Split(_only, ",")
		fetch.only = make([]string, len(sp))
		for i, v := range sp {
			fetch.only[i] = strings.TrimSpace(v)
		}
	}

	_criteria := parsed[CRITERIA]
	_criteria = strings.TrimSpace(_criteria)
	criteria, err := parseCriteria(_criteria)
	if err != nil {
		return nil, err
	}
	fetch.criteria = criteria

	return fetch, nil
}

func (f *Fetch) Criteria() *Criteria {
	return f.criteria
}

func (f *Fetch) ElementType() string {
	return f.elementType
}

func (f *Fetch) Only() []string {
	return f.only
}

func (f *Fetch) Sort() string {
	return f.sort
}

func (f *Fetch) String() string {
	s := &strng.String{}
	s.Add("fetch ")
	s.Add(f.elementType)
	if f.only != nil && len(f.only) > 0 {
		s.Add(" only ")
		first := true
		for _, fld := range f.only {
			if !first {
				s.Add(",")
			}
			first = false
			s.Add(fld)
		}
	}
	if f.criteria != nil {
		s.Add(" criteria ")
		s.Add(f.criteria.String())
	}
	return s.String()
}
