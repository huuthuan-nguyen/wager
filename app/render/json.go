package render

import (
	"encoding/json"
	"github.com/huuthuan-nguyen/wager/app/transformer"
	"github.com/huuthuan-nguyen/wager/app/utils"
	"net/http"
)

// JSON /**
func JSON(w http.ResponseWriter, r *http.Request, payload interface{}) {
	transformerManager := transformer.Manager{
		Serializer: &transformer.JSONSerializer{},
	}

	payloadStruct := transformerManager.CreateData(payload)
	payloadResponse, err := json.Marshal(payloadStruct)
	if err != nil {
		utils.PanicInternalServerError()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payloadResponse)
}
