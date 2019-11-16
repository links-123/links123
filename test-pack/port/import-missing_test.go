package link

import (
	"fmt"
	"github.com/ic2hrmk/ship_links/app/gateways/link/errors"
	"testing"

	"github.com/go-resty/resty"
	"github.com/ic2hrmk/ship_links/shared/env"
	"github.com/ic2hrmk/ship_links/shared/gateway/representation"
)

//
// Looks how service reacts on case, when file isn't attached
//
func TestImlink_Missing(t *testing.T) {
	//
	// Get file fixture path
	//
	testFilePath, err := env.GetStringVar(inconsistentFilePath)
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
		SetFile("file2", testFilePath). // Add multipart file but with wrong param
		SetError(&errResponse).
		Post(fmt.Sprintf("%s/api/links/import", linkGatewayURL))

	if err != nil {
		t.Fatalf("failed to send upload request, %s", err)
	}

	if resp.IsSuccess() {
		t.Fatalf("wrong server response - errors are not detected")
	}

	if errResponse == nil {
		t.Fatalf("error response is nil")
	}

	if errResponse.Message != errors.ErrFileNotAttached {
		t.Fatalf("wrong error message [%s], expected [%s]", errResponse.Message, errors.ErrFileNotAttached)
	}
}
