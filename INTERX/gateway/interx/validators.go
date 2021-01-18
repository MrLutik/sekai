package interx

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/KiraCore/sekai/INTERX/common"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterValidatorsQueryRoutes registers validators query routers.
func RegisterValidatorsQueryRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(common.QueryValidators, QueryValidators(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", common.QueryValidators, "This is an API to query validators.", true)
}

func queryValidatorsHandle(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	queries := r.URL.Query()
	key := queries["key"]
	offset := queries["offset"]
	limit := queries["limit"]
	countTotal := queries["count_total"]

	var events = make([]string, 0, 4)
	if len(key) == 1 {
		events = append(events, fmt.Sprintf("pagination.key=%s", key[0]))
	}
	if len(offset) == 1 {
		events = append(events, fmt.Sprintf("pagination.offset=%s", offset[0]))
	}
	if len(limit) == 1 {
		events = append(events, fmt.Sprintf("pagination.limit=%s", limit[0]))
	}
	if len(countTotal) == 1 {
		events = append(events, fmt.Sprintf("pagination.count_total=%s", countTotal[0]))
	}

	r.URL.RawQuery = strings.Join(events, "&")
	fmt.Println(r.URL.RawQuery)

	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryValidators is a function to list validators.
func QueryValidators(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		response.Response, response.Error, statusCode = queryValidatorsHandle(r, gwCosmosmux)

		common.WrapResponse(w, request, *response, statusCode, false)
	}
}