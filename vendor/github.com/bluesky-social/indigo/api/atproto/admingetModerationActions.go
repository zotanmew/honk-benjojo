// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.getModerationActions

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// AdminGetModerationActions_Output is the output of a com.atproto.admin.getModerationActions call.
type AdminGetModerationActions_Output struct {
	Actions []*AdminDefs_ActionView `json:"actions" cborgen:"actions"`
	Cursor  *string                 `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
}

// AdminGetModerationActions calls the XRPC method "com.atproto.admin.getModerationActions".
func AdminGetModerationActions(ctx context.Context, c *xrpc.Client, cursor string, limit int64, subject string) (*AdminGetModerationActions_Output, error) {
	var out AdminGetModerationActions_Output

	params := map[string]interface{}{
		"cursor":  cursor,
		"limit":   limit,
		"subject": subject,
	}
	if err := c.Do(ctx, xrpc.Query, "", "com.atproto.admin.getModerationActions", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}