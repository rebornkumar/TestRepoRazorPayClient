package resources_test

import (
	"encoding/json"
	"testing"

	"code.nurture.farm/nurture.trade/razorpaygoclient/constants"
	"code.nurture.farm/nurture.trade/razorpaygoclient/utils"
	"github.com/stretchr/testify/assert"
)

const TestCardID = "fake_card_id"

func TestCardFetch(t *testing.T) {
	url := constants.CARD_URL + "/" + TestCardID
	teardown, fixture := utils.StartMockServer(url, "fake_card")
	defer teardown()
	body, err := utils.Client.Card.Fetch(TestCardID, nil, nil)
	jsonByteArray, _ := json.Marshal(body)
	assert.Equal(t, err, nil)
	utils.TestResponse(jsonByteArray, []byte(fixture), t)
}
