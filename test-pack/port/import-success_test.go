package link

import (
	"fmt"
	"github.com/ic2hrmk/ship_links/shared/gateway/representation"
	"testing"

	"github.com/go-resty/resty"
	"github.com/ic2hrmk/ship_links/shared/env"
)

//
// Looks how service reacts on case, when file is OK
//
func TestImlink_Success(t *testing.T) {
	//
	// Get file fixture path
	//
	testFilePath, err := env.GetStringVar(correctFilePath)
	if err != nil {
		t.Fatalf("failed to get fixture file path: %s", err)
	}

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
	errResponse := &representation.ErrorResponse{}

	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFile("file", testFilePath).
		SetError(&errResponse).
		Post(fmt.Sprintf("%s/api/links/import", linkGatewayURL))

	if err != nil {
		t.Fatalf("failed to send upload request, %s", err)
	}

	if !resp.IsSuccess() {
		t.Fatalf("wrong server response - errors are detected, %v", resp)
	}

	if errResponse != nil && errResponse.Message != "" {
		t.Fatalf("error response is not nil, %v", *errResponse)
	}
}
