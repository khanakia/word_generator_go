package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Word struct {
	ent.Schema
}

func (Word) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
	}
}

func (Word) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		// index.Fields("name"),
	}
}
