package resources

import (
	"fmt"

	"code.nurture.farm/nurture.trade/razorpaygoclient/constants"
	"code.nurture.farm/nurture.trade/razorpaygoclient/requests"
)

//Addon ...
type Addon struct {
	Request *requests.Request
}

// Fetch fetches addon having the given addonID.
func (addon *Addon) Fetch(addonID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.ADDON_URL, addonID)
	return addon.Request.Get(url, queryParams, extraHeaders)
}

// Delete deletes the addon having the given addonID.
func (addon *Addon) Delete(addonID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.ADDON_URL, addonID)
	return addon.Request.Delete(url, queryParams, extraHeaders)
}
