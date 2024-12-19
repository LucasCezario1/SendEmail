package endpoints

import (
	"SendEmail/internal/contract"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

)

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaignRequest{
		Name:    "Teste",
		Content: "HI LUCAS",
		Emails:  []string{"test@test.com"},
	}
	serviceMock := new(internalMock.)
	serviceMock.On("Create", mock.MatchedBy(func(req contract.NewCampaignRequest) bool {
		if req.Name == body.Name &&
			req.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	handler := Handler{CampaignService: serviceMock}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)

	rr := httptest.NewRecorder()
	_, status, err := handler.CampaignsPost(rr, req)

	assert.Equal(status, http.StatusCreated)
	assert.Nil(err)
}

func Test_CampaignsPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaignRequest{
		Name:    "Teste",
		Content: "HI LUCAS",
		Emails:  []string{"test@test.com"},
	}
	serviceMock := new(serviceMock)
	serviceMock.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: serviceMock}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)

	rr := httptest.NewRecorder()
	_, _, err := handler.CampaignsPost(rr, req)

	assert.NotNil(err)

}
