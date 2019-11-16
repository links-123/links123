package link

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty"
	"github.com/ic2hrmk/ship_links/app/gateways/link/representation"
	"github.com/ic2hrmk/ship_links/shared/env"
)

//
// Looks how service reacts on case, when file is damaged
//
func TestFindAll_Success(t *testing.T) {
	//
	// Get target URL
	//
	linkGatewayURL, err := env.GetStringVar(linkGatewayURL)
	if err != nil {
		t.Fatalf("failed to get target URL: %s", err)
	}

	//
	// Assemble request
	//
	findResponse := &representation.LinkListResponse{}
	errResponse := &representation.ErrorResponse{}

	resp, err := resty.R().
		SetResult(&findResponse).
		SetError(&errResponse).
		Get(fmt.Sprintf("%s/api/links", linkGatewayURL))

	if err != nil {
		t.Fatalf("failed to send upload request, %s", err)
	}

	if !resp.IsSuccess() {
		t.Fatalf("wrong server response - errors are detected, %v", resp)
	}

	if errResponse != nil && errResponse.Message != "" {
		t.Fatalf("error response is not nil, %v", errResponse)
	}

	if findResponse == nil {
		t.Fatalf("find response is nil")
	}
}
