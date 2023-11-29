package orm

import (
	"github.com/google/uuid"
	"github.com/saichler/my.simple/go/common"
	"github.com/saichler/my.simple/go/orm/relational"
	"github.com/saichler/my.simple/go/utils/strng"
)

type ORM struct {
	plugin     common.IOrmPlugin
	introspect common.IIntrospect
}

func NewOrm(plugin common.IOrmPlugin, introspect common.IIntrospect) common.IORM {
	return &ORM{plugin: plugin, introspect: introspect}
}

func (o *ORM) Introspect() common.IIntrospect {
	return o.introspect
}

func (o *ORM) Persist(any interface{}) error {
	_, err := o.introspect.Inspect(any)
	if err != nil {
		return err
	}
	if o.plugin.RelationalData() {
		relationalData := relational.NewRelationalData(newOrmTransactionID())
		err = relationalData.AddInstances(any, o.introspect)
		if err != nil {
			return err
		}
		return o.plugin.Persist(relationalData)
	}
	return o.plugin.Persist(any)
}

func (o *ORM) Fetch(fetch common.IFetch) (interface{}, error) {
	data, err := o.plugin.Fetch(fetch)
	if o.plugin.RelationalData() {
		rdata := data.(*relational.RelationalData)
		return rdata.ToIstances(o.introspect)
	}
	return data, err
}

func newOrmTransactionID() string {
	return strng.New("orm-", uuid.New().String()).String()
}
