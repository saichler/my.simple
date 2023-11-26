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

func NewOrm(plugin common.IOrmPlugin, introspect common.IIntrospect) *ORM {
	return &ORM{plugin: plugin, introspect: introspect}
}

func (o *ORM) Introspect() common.IIntrospect {
	return o.introspect
}

func (o *ORM) Write(any interface{}) error {
	_, err := o.introspect.Inspect(any)
	if err != nil {
		return err
	}
	if o.plugin.SQL() {
		relationalData := relational.NewRelationalData(newOrmTransactionID())
		err = relationalData.AddInstances(any, o.introspect)
		if err != nil {
			return err
		}
		return o.plugin.Write(relationalData, o)
	}
	return o.plugin.Write(any, o)
}

func newOrmTransactionID() string {
	return strng.New("orm-", uuid.New().String()).String()
}
