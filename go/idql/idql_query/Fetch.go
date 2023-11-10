package idql_query

import (
	"errors"
	"github.com/saichler/my.simple/go/idql/idql_parser"
	"github.com/saichler/my.simple/go/introspect"
	"reflect"
	"strings"
)

type Fetch struct {
	pFetch       *idql_parser.Fetch
	criteria     *Criteria
	only         []string
	introspector *introspect.Introspect
}

func NewFetch(request string, introspector *introspect.Introspect) (*Fetch, error) {
	pFetch, err := idql_parser.NewFetch(request)
	if err != nil {
		return nil, err
	}
	fetch := &Fetch{}
	fetch.introspector = introspector
	fetch.pFetch = pFetch
	err = fetch.validateElement()
	if err != nil {
		return nil, err
	}

	err = fetch.validateOnly()
	if err != nil {
		return nil, err
	}

	criteria, err := newCriteria(pFetch.Criteria(), fetch.pFetch.ElementType())
	if err != nil {
		return nil, err
	}
	fetch.criteria = criteria
	return fetch, nil
}

func (fetch *Fetch) validateElement() error {
	_, ok := fetch.introspector.Node(fetch.pFetch.ElementType())
	if ok {
		return nil
	}
	return errors.New("Element Type " + fetch.pFetch.ElementType() + " does not exist")
}

func (fetch *Fetch) Criteria() *Criteria {
	return fetch.criteria
}

func (fetch *Fetch) validateOnly() error {
	only := fetch.pFetch.Only()
	for _, att := range only {
		if strings.HasPrefix(att, fetch.pFetch.ElementType()) {
			_, ok := fetch.introspector.Node(att)
			if !ok {
				return errors.New("Unknown Only request for " + att)
			}
		} else {
			_, ok := fetch.introspector.Node(fetch.pFetch.ElementType() + "." + att)
			if !ok {
				return errors.New("Unknown Only request for " + att)
			}
		}
	}
	return nil
}

func (fetch *Fetch) Match(any interface{}) (bool, error) {
	val := reflect.ValueOf(any)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.IsValid() && strings.ToLower(val.Type().Name()) != fetch.pFetch.ElementType() {
		return false, errors.New("Element Type does not match")
	}
	return fetch.criteria.Match(reflect.ValueOf(any))
}

func (fetch *Fetch) String() string {
	return fetch.pFetch.String()
}
