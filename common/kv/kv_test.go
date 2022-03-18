package kv

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {

	str := ``
	obj := FromString(str)
	kvType := obj.GetType()
	if kvType != UndefinedType {
		t.Fatalf("将空字符串系列化成JSON，得到的应该是undefined，但测试得到的却是%s", obj.StringValue())
	}

	str = `hello world`
	obj = FromString(str)
	kvType = obj.GetType()
	if kvType != UndefinedType {
		t.Fatalf("将不带引号的字符串%s系列化成JSON，得到的应该是undefined，但测试得到的却是%s", str, obj.StringValue())
	}

	str = `hello world`
	obj = FromString(fmt.Sprintf("%q", str))
	if obj.GetString() != str {
		t.Fatalf("将带引号的字符串%q系列化成JSON，得到的应该是%s，但测试得到的却是%s", str, str, obj.StringValue())
	}

	str = `null`
	obj = FromString(str)
	if obj.GetNull() != nil {
		t.Fatalf("将%s系列化成JSON，得到的应该是null，但测试得到的却是%s", str, obj.StringValue())
	}

	str = `false`
	obj = FromString(str)
	if obj.GetBoolean() != false {
		t.Fatalf("将%s系列化成JSON，得到的应该是false，但测试得到的却是%s", str, obj.StringValue())
	}

	str = `true`
	obj = FromString(str)
	if obj.GetBoolean() != true {
		t.Fatalf("将%s系列化成JSON，得到的应该是true，但测试得到的却是%s", str, obj.StringValue())
	}

	num := 100.00011
	obj = FromString(fmt.Sprintf("%g", num))
	if obj.GetNumber() != num {
		t.Fatalf("将%g系列化成JSON，得到的应该是%g，但测试得到的却是%s", num, num, obj.StringValue())
	}

}

func TestObject(t *testing.T) {
	str := `
		{
			"name":"Jim Green",
			"age":34,
			"height":180,
			"others":{
				"skills":[
					{
						"golang":{"year":3,"lv":"heigh"}
					},
					"C#",
					"nodejs"
				]
			}
		}
	`

	obj := FromString(str)
	val := obj.GetProperty("others").GetProperty("skills").GetElement(1).GetString()
	if val != "C#" {
		t.Fatalf("/others/skills[1] should be C#,but get %s,obj.ToJson() = %q", val, obj.ToJson())
	}

	lvPath := "/others/skills[0]/golang/lv"
	val = obj.Select(lvPath).GetString()
	if val != "heigh" {
		t.Fatalf("%s should be heigh,but get %s,obj.ToJson()=%q", lvPath, val, obj.ToJson())
	}

	lvPath = "errorpath"
	val = obj.Select(lvPath).GetString()
	if val != "" {
		t.Fatalf("%s should be empty,but get %s,obj.ToJson()=%q", lvPath, val, obj.ToJson())
	}
}

func TestArray(t *testing.T) {
	str := `
	[
		{
			"name":"Jim Green",
			"age":34,
			"height":180,
			"others":{
				"skills":[
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					{
						"golang":{"year":3,"lv":"heigh"}
					},
					"C#",
					"nodejs"
				]
			}
		},
		"hello"
	]
`

	obj := FromString(str)
	val := obj.GetElement(0).GetProperty("others").GetProperty("skills").GetElement(14).GetString()
	if val != "C#" {
		t.Fatalf("[0]/others/skills[14] should be C#,but get %s,obj.ToJson() = %q", val, obj.ToJson())
	}

	lvPath := "[0]/others/skills[13]/golang/lv"
	val = obj.Select(lvPath).GetString()
	if val != "heigh" {
		t.Fatalf("%s should be heigh,but get %s,obj.ToJson()=%q", lvPath, val, obj.ToJson())
	}

	elevth := "[0]/others/skills[12]"
	val = obj.Select(elevth).GetString()
	if val != "hello" {
		t.Fatalf("%s should be hello,but get %s,obj.ToJson()=%q", elevth, val, obj.ToJson())
	}

	lvPath = "errorpath"
	val = obj.Select(lvPath).GetString()
	if val != "" {
		t.Fatalf("%s should be empty,but get %s,obj.ToJson()=%q", lvPath, val, obj.ToJson())
	}
}

func BenchmarkSelect(b *testing.B) {
	str := `
	[
		{
			"name":"Jim Green",
			"age":34,
			"height":180,
			"others":{
				"skills":[
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					"hello",
					{
						"golang":{"year":3,"lv":"heigh"}
					},
					"C#",
					"nodejs"
				]
			}
		},
		"hello"
	]
`
	for i := 0; i < b.N; i++ {

		obj := FromString(str)
		lvPath := "[0]/others/skills[13]/golang/lv"
		val := obj.Select(lvPath).GetString()
		if val != "heigh" {
			fmt.Printf("%s should be heigh,but get %s,obj.ToJson()=%q", lvPath, val, obj.ToJson())
		}
	}
}
