// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent/category"
	"entgo.io/contrib/entgql/internal/todogotype/ent/friendship"
	"entgo.io/contrib/entgql/internal/todogotype/ent/group"
	"entgo.io/contrib/entgql/internal/todogotype/ent/pet"
	"entgo.io/contrib/entgql/internal/todogotype/ent/schema/bigintgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent/schema/uintgql"
	"entgo.io/contrib/entgql/internal/todogotype/ent/todo"
	"entgo.io/contrib/entgql/internal/todogotype/ent/user"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     string   `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string   `json:"type,omitempty"` // edge type.
	Name string   `json:"name,omitempty"` // edge name.
	IDs  []string `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (c Category) marshalID() string {
	var buf bytes.Buffer
	c.ID.MarshalGQL(&buf)
	return buf.String()
}

func (c *Category) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     c.marshalID(),
		Type:   "Category",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(c.Text); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "text",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Status); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "category.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Config); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "*schematype.CategoryConfig",
		Name:  "config",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Duration); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "time.Duration",
		Name:  "duration",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Count); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "uint64",
		Name:  "count",
		Value: string(buf),
	}
	if buf, err = json.Marshal(c.Strings); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "[]string",
		Name:  "strings",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Todo",
		Name: "todos",
	}
	err = c.QueryTodos().
		Select(todo.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (f *Friendship) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     f.ID,
		Type:   "Friendship",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(f.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.UserID); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "user_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(f.FriendID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "friend_id",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user",
	}
	err = f.QueryUser().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "friend",
	}
	err = f.QueryFriend().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (gr *Group) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gr.ID,
		Type:   "Group",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(gr.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "users",
	}
	err = gr.QueryUsers().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (pe Pet) marshalID() string {
	var buf bytes.Buffer
	pe.ID.MarshalGQL(&buf)
	return buf.String()
}

func (pe *Pet) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pe.marshalID(),
		Type:   "Pet",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(pe.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	return node, nil
}

func (t *Todo) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     t.ID,
		Type:   "Todo",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(t.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Status); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "todo.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Priority); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "priority",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.Text); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "text",
		Value: string(buf),
	}
	if buf, err = json.Marshal(t.CategoryID); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "bigintgql.BigInt",
		Name:  "category_id",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Todo",
		Name: "parent",
	}
	err = t.QueryParent().
		Select(todo.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Todo",
		Name: "children",
	}
	err = t.QueryChildren().
		Select(todo.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Category",
		Name: "category",
	}
	err = t.QueryCategory().
		Select(category.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(u.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Group",
		Name: "groups",
	}
	err = u.QueryGroups().
		Select(group.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "friends",
	}
	err = u.QueryFriends().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "Friendship",
		Name: "friendships",
	}
	err = u.QueryFriendships().
		Select(friendship.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id string) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, string) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, string) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, string) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id string) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id string, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id string) (Noder, error) {
	switch table {
	case category.Table:
		var uid bigintgql.BigInt
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Category.Query().
			Where(category.ID(uid))
		query, err := query.CollectFields(ctx, "Category")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case friendship.Table:
		query := c.Friendship.Query().
			Where(friendship.ID(id))
		query, err := query.CollectFields(ctx, "Friendship")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case group.Table:
		query := c.Group.Query().
			Where(group.ID(id))
		query, err := query.CollectFields(ctx, "Group")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case pet.Table:
		var uid uintgql.Uint64
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Pet.Query().
			Where(pet.ID(uid))
		query, err := query.CollectFields(ctx, "Pet")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case todo.Table:
		query := c.Todo.Query().
			Where(todo.ID(id))
		query, err := query.CollectFields(ctx, "Todo")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []string, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]string)
	id2idx := make(map[string][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []string) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[string][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case category.Table:
		uids := make([]bigintgql.BigInt, len(ids))
		for i, id := range ids {
			if err := uids[i].UnmarshalGQL(id); err != nil {
				return nil, err
			}
		}
		query := c.Category.Query().
			Where(category.IDIn(uids...))
		query, err := query.CollectFields(ctx, "Category")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.marshalID()] {
				*noder = node
			}
		}
	case friendship.Table:
		query := c.Friendship.Query().
			Where(friendship.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Friendship")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case group.Table:
		query := c.Group.Query().
			Where(group.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Group")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case pet.Table:
		uids := make([]uintgql.Uint64, len(ids))
		for i, id := range ids {
			if err := uids[i].UnmarshalGQL(id); err != nil {
				return nil, err
			}
		}
		query := c.Pet.Query().
			Where(pet.IDIn(uids...))
		query, err := query.CollectFields(ctx, "Pet")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.marshalID()] {
				*noder = node
			}
		}
	case todo.Table:
		query := c.Todo.Query().
			Where(todo.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Todo")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}