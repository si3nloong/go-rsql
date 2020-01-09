package rsql

import (
	"log"
	"reflect"
	"strings"
)

// StructField :
type StructField struct {
	Name string
	Tag  *StructTag
	Type reflect.Type
}

// StructTag :
type StructTag struct {
	name   string
	values map[string]string
}

// Struct :
type Struct struct {
	Fields []*StructField
	Names  map[string]*StructField
}

func NewTag(name string, tag reflect.StructTag) *StructTag {
	paths := strings.Split(tag.Get(name), ",")
	t := new(StructTag)
	t.name = paths[0]
	t.values = make(map[string]string)
	for _, v := range paths[1:] {
		p := strings.SplitN(v, "=", 2)
		t.values[p[0]] = ""
		if len(p) > 1 {
			t.values[p[0]] = p[1]
		}
	}
	return t
}

func (t StructTag) Lookup(key string) (value string, ok bool) {
	value, ok = t.values[key]
	return
}

func getCodec(t reflect.Type) *Struct {
	fields := make([]*StructField, 0)
	codec := new(Struct)
	for i := 0; i < t.NumField(); i++ {
		fv := t.Field(i)
		log.Println(fv)

		tag := NewTag("rsql", fv.Tag)
		f := new(StructField)
		f.Name = fv.Name
		f.Tag = tag
		if tag.name != "" {
			f.Name = tag.name
		}
		f.Type = fv.Type

		fields = append(fields, f)
	}

	codec.Fields = fields
	codec.Names = make(map[string]*StructField)
	for _, f := range fields {
		codec.Names[f.Name] = f
	}
	// log.Println(fields)
	return codec
}
