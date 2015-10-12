package filter
import (
	"encoding/json"
	"testing"
	"strings"
)

func TestFieldsFilter(t *testing.T)  {
	var filterString = "id,test/a,object/attr/name"
	var _ = ",url/a,id,object(site,attachments/url,attr/name,content/name),test"
	var jsonString = `{
	"id": true,
	"test": {"a":false,"b":123444444444444444444444444444},
 "url": "https://",
 "object": {
  "site": "leenanxi.com",
  "content":{
  	"name": "leenanxi",
  	"key": "youku"
  },
  "attr":[{
  	"name": "leenanxi",
  	"key": "youku"
  }],
  "attachments": [
   "ttt",{
    "url": "https",
    "name": "110908"
   },
   {
    "url": "http",
    "name": "0000"
   }
  ]
 }
}`
	dec := json.NewDecoder(strings.NewReader(jsonString))
	dec.UseNumber()
	var data interface{}
	if err := dec.Decode(&data); err != nil {
		panic(err)
	}
	e := Filter(&data, filterString)
	if e != nil {
		t.Log(e)
		t.Fail()
	}
	bytes, _ := json.MarshalIndent(data, "", "\t")
	t.Log("The Filter Result：", string(bytes))
}