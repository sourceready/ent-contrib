package entgql

import "context"

const (
	// CascadeQueryOneKey is the context key to enable cascading query for one operation.
	CascadeQueryOneKey = "ent_cascade_query_one"
	// CascadeQueryAllKey is the context key to enable cascading query for all operations.
	CascadeQueryAllKey = "ent_cascade_query_all"
)

func IsCascadeQueryOne(ctx context.Context) bool {
	return ctx.Value(CascadeQueryOneKey) == "true"
}

func IsCascadeQueryAll(ctx context.Context) bool {
	return ctx.Value(CascadeQueryAllKey) == "true"
}
