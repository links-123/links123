package link

import (
	"fmt"
	"github.com/ic2hrmk/links123/app/gateways/link/errors"
	"testing"

	"github.com/go-resty/resty"
	"github.com/ic2hrmk/links123/shared/env"
	"github.com/ic2hrmk/links123/shared/gateway/representation"
)

//
// Looks how service reacts on case, when file is damaged
//
func TestImlink_Inconsistent(t *testing.T) {
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
		SetFile("file", testFilePath).
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

	if errResponse.Message != errors.ErrInconsistentJson {
		t.Fatalf("wrong error message [%s], expected [%s]", errResponse.Message, errors.ErrInconsistentJson)
	}
}
