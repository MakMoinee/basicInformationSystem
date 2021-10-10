package response

import (
	"encoding/json"
	"net/http"

	"github.com/MakMoinee/basicInformationSystem/internal/common"
)

// Response Data Model
type Response struct {
	Data interface{} `json:"data"`
}

// returns success response
func Success(data interface{}, w http.ResponseWriter) {
	payload, err := json.Marshal(data)
	if err != nil {
		w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set(common.ContentTypeKey, common.ContentTypeValue)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
