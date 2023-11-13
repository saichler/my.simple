package mdql_query

import (
	"errors"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/mdql/mdql_parser"
	"reflect"
	"strings"
)

type Fetch struct {
	pFetch     *mdql_parser.Fetch
	criteria   *Criteria
	only       []string
	introspect common.IIntrospect
}

func NewFetch(request string, introspect common.IIntrospect) (*Fetch, error) {
	pFetch, err := mdql_parser.NewFetch(request)
	if err != nil {
		return nil, err
	}
	fetch := &Fetch{}
	fetch.introspect = introspect
	fetch.pFetch = pFetch
	err = fetch.validateElement()
	if err != nil {
		return nil, err
	}

	err = fetch.validateOnly()
	if err != nil {
		return nil, err
	}

	criteria, err := newCriteria(pFetch.Criteria(), fetch.pFetch.ElementType(), introspect)
	if err != nil {
		return nil, err
	}
	fetch.criteria = criteria
	return fetch, nil
}

func (fetch *Fetch) validateElement() error {
	_, ok := fetch.introspect.Node(fetch.pFetch.ElementType())
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
			_, ok := fetch.introspect.Node(att)
			if !ok {
				return errors.New("Unknown Only request for " + att)
			}
		} else {
			_, ok := fetch.introspect.Node(fetch.pFetch.ElementType() + "." + att)
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
