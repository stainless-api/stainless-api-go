// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stainlessv0

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/stainless-v0-go/internal/requestconfig"
	"github.com/stainless-sdks/stainless-v0-go/option"
	"github.com/stainless-sdks/stainless-v0-go/packages/param"
)

// WebhookPostmanService contains methods and other services that help with
// interacting with the stainless-v0 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookPostmanService] method instead.
type WebhookPostmanService struct {
	Options []option.RequestOption
}

// NewWebhookPostmanService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebhookPostmanService(opts ...option.RequestOption) (r WebhookPostmanService) {
	r = WebhookPostmanService{}
	r.Options = opts
	return
}

// TODO
func (r *WebhookPostmanService) NewNotification(ctx context.Context, body WebhookPostmanNewNotificationParams, opts ...option.RequestOption) (res *WebhookPostmanNewNotificationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v0/webhooks/postman/notifications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type WebhookPostmanNewNotificationResponse = any

type WebhookPostmanNewNotificationParams struct {
	CollectionID string `json:"collectionId,required"`
	paramObj
}

// IsPresent returns true if the field's value is not omitted and not the JSON
// "null". To check if this field is omitted, use [param.IsOmitted].
func (f WebhookPostmanNewNotificationParams) IsPresent() bool {
	return !param.IsOmitted(f) && !f.IsNull()
}

func (r WebhookPostmanNewNotificationParams) MarshalJSON() (data []byte, err error) {
	type shadow WebhookPostmanNewNotificationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
