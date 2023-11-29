package sqlbase

import (
	"fmt"
	"github.com/saichler/my.simple/go/common"
)

func (plugin *OrmSqlBasePlugin) Fetch(fetch common.IFetch) (interface{}, error) {
	for _, tname := range plugin.cache.Tables() {
		fmt.Println(tname)
	}
	return nil, nil
}
