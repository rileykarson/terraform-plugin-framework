// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package testprovider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/list"
)

var _ list.ListResource = &ListResourceWithRawV6Schemas{}
var _ list.ListResourceWithRawV6Schemas = &ListResourceWithRawV6Schemas{}

// Declarative list.ListResourceWithRawV6Schemas for unit testing.
type ListResourceWithRawV6Schemas struct {
	*ListResource

	// ListResourceWithRawV6Schemas interface methods
	RawV6SchemasMethod func(context.Context, list.RawV6SchemaRequest, *list.RawV6SchemaResponse)
}

// RawV6Schemas satisfies the list.ListResourceWithRawV6Schemas interface.
func (r *ListResourceWithRawV6Schemas) RawV6Schemas(ctx context.Context, req list.RawV6SchemaRequest, resp *list.RawV6SchemaResponse) {
	if r.RawV6SchemasMethod == nil {
		return
	}

	r.RawV6SchemasMethod(ctx, req, resp)
}
