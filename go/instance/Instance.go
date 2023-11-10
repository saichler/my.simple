package instance

import (
	"errors"
	"github.com/saichler/my.simple/go/introspect"
	"github.com/saichler/my.simple/go/introspect/model"
	"github.com/saichler/my.simple/go/utils/strng"
	"strings"
)

type Instance struct {
	parent *Instance
	node   *model.Node
	key    interface{}
	value  interface{}
	id     string
}

func NewInstance(node *model.Node, parent *Instance, key interface{}, value interface{}) *Instance {
	i := &Instance{}
	i.parent = parent
	i.node = node
	i.key = key
	i.value = value
	return i
}

func InstanceOf(instanceId string, i *introspect.Introspect) (*Instance, error) {
	instanceKey := NodeKey(instanceId)
	node, ok := i.Node(instanceKey)
	if !ok {
		return nil, errors.New("Unknown attribute " + instanceKey)
	}
	return newInstance(node, instanceId)
}

func (inst *Instance) Parent() *Instance {
	return inst.parent
}

func (inst *Instance) Node() *model.Node {
	return inst.node
}

func (inst *Instance) Key() interface{} {
	return inst.key
}

func (inst *Instance) Value() interface{} {
	return inst.value
}

func (inst *Instance) setKeyValue(instanceId string) (string, error) {
	id := instanceId
	dIndex := strings.LastIndex(instanceId, ".")
	if dIndex == -1 {
		return "", nil
	}
	beIndex := strings.LastIndex(instanceId, ">")
	if beIndex == -1 {
		return "", nil
	}
	for dIndex < beIndex {
		id = id[0:beIndex]
		dIndex = strings.LastIndex(id, ".")
		beIndex = strings.LastIndex(id, ">")
	}
	prefix := instanceId[0:dIndex]
	suffix := instanceId[dIndex+1:]
	bbIndex := strings.LastIndex(suffix, "<")
	if bbIndex == -1 {
		return prefix, nil
	}

	v := suffix[bbIndex+1 : len(suffix)-1]
	value, err := strng.FromString(v)
	if err != nil {
		return "", err
	}
	inst.key = value.Interface()
	return prefix, nil
}

func (inst *Instance) InstanceId() (string, error) {
	if inst.id != "" {
		return inst.id, nil
	}
	buff := &strng.String{}
	if inst.parent == nil {
		buff.Add(strings.ToLower(inst.node.TypeName))
		buff.Add(inst.node.CachedKey)
	} else {
		pi, err := inst.parent.InstanceId()
		if err != nil {
			return "", err
		}
		buff.Add(pi)
		buff.Add(".")
		buff.Add(strings.ToLower(inst.node.FieldName))
	}

	if inst.key != nil {
		keyStr := strng.String{TypesPrefix: true}
		keyStrVal, err := keyStr.StringOf(inst.key)
		if err != nil {
			return "", err
		}
		buff.Add("<")
		buff.Add(keyStrVal)
		buff.Add(">")
	}
	inst.id = buff.String()
	return inst.id, nil
}

func NodeKey(instanceId string) string {
	buff := &strng.String{}
	open := false
	for _, c := range instanceId {
		if c == '<' {
			open = true
		} else if c == '>' {
			open = false
		} else if !open {
			buff.Add(string(c))
		}
	}
	return buff.String()
}

func newInstance(node *model.Node, instancePath string) (*Instance, error) {
	inst := &Instance{}
	inst.node = node
	if node.Parent != nil {
		prefix, err := inst.setKeyValue(instancePath)
		if err != nil {
			return nil, err
		}
		pi, err := newInstance(node.Parent, prefix)
		if err != nil {
			return nil, err
		}
		inst.parent = pi
	} else {
		index1 := strings.Index(instancePath, "<")
		index2 := strings.Index(instancePath, ">")
		if index1 != -1 && index2 != -1 && index2 > index1 {
			keyVal, err := strng.FromString(instancePath[index1+1 : index2])
			if err != nil {
				return nil, err
			}
			inst.key = keyVal.Interface()
		}
	}
	return inst, nil
}
