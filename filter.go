package filter

import (
	"fmt"
)

// filter data by fields like "attr,path1/path2,objects(attr,path1/path2,objects(attr,path1/path2))"
func Filter(data *interface{}, fileds string) error {
	tree, err := Compile(fileds)
	if err != nil {
		return fmt.Errorf("fileds: %v", err)
	}
	*data = recursiveFilter(tree, *data)
	return nil
}

func recursiveFilter(t *Tree, data interface{}) interface{} {
	if t.Children == nil {
		return data
	} else if d, ok := data.(map[string]interface{}); ok {
		ret := map[string]interface{}{}
		for tk, tv := range t.Children {
			if dv, ok := d[tk]; ok {
				ret[tk] = recursiveFilter(tv, dv)
			}
		}
		return ret
	} else if d, ok := data.([]interface{}); ok {
		ret := make([]interface{}, len(d))
		for dk, dv := range d {
			if r := recursiveFilter(t, dv); r != nil {
				ret[dk] = r
			}
		}
		return ret
	} else {
		return nil
	}
}
