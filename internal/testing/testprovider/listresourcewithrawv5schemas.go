// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package testprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/list"
)

var _ list.ListResource = &ListResourceWithRawV5Schemas{}
var _ list.ListResourceWithRawV5Schemas = &ListResourceWithRawV5Schemas{}

// Declarative list.ListResourceWithRawV5Schemas for unit testing.
type ListResourceWithRawV5Schemas struct {
	*ListResource

	// ListResourceWithRawV5Schemas interface methods
	RawV5SchemasMethod func(context.Context, list.RawV5SchemaRequest, *list.RawV5SchemaResponse)
}

// RawV5Schemas satisfies the list.ListResourceWithRawV5Schemas interface.
func (r *ListResourceWithRawV5Schemas) RawV5Schemas(ctx context.Context, req list.RawV5SchemaRequest, resp *list.RawV5SchemaResponse) {
	if r.RawV5SchemasMethod == nil {
		return
	}

	r.RawV5SchemasMethod(ctx, req, resp)
}
