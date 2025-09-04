package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"time"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("title"),
		field.String("content"),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{dialect.MySQL: "timestamp"}),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{dialect.MySQL: "timestamp"}),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return nil
}
