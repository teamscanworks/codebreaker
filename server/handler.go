package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/teamscanworks/codebreaker/types"
)

// Handler is the core object in the server package. It keeps an in-memory state
// of the chain-registry which can be updated using `Pull`. It handles requests
// for this data through the router.
type Handler struct {
	registryUrl string
	//rpcEndpoints string // todo: initialize wasm module rpc endpoints for each supported chain
	lastUpdated time.Time
	chains      []string
	chainList   map[string]types.ContractsRegistry
	log         *log.Logger
}

func NewHandler(registryUrl string, log *log.Logger) *Handler {
	return &Handler{
		registryUrl: registryUrl,
		lastUpdated: time.Unix(0, 0),
		chains:      make([]string, 0),
		chainList:   make(map[string]types.ContractsRegistry),
		log:         log,
	}
}

func (h Handler) Chains(res http.ResponseWriter, req *http.Request) {
	respondWithJSON(res, http.StatusOK, h.chains)
}

func (h Handler) ContractByCodeId(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	chainName, ok := vars["chain"]
	if !ok {
		respondWithError(res, http.StatusBadRequest, "Invalid request payload. No valid chain provided.")
		return
	}
	codeId, ok := vars["codeId"]
	if !ok {
		respondWithError(res, http.StatusBadRequest, "Invalid request payload. No valid code id provided.")
		return
	}

	contractRegistry, ok := h.chainList[chainName]
	if !ok {
		respondWithError(res, http.StatusNotFound, fmt.Sprintf("Chain information for %s could not be found.", chainName))
		return
	}

	contractInfo, exists := contractRegistry.Contracts[codeId]

	if !exists {
		respondWithError(res, http.StatusNotFound, fmt.Sprintf("Contract information for code-id %s on %s could not be found.", codeId, chainName))
		return
	}

	respondWithJSON(res, http.StatusOK, contractInfo)
}

/*
func (h Handler) ContractByAddress(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	chainName, ok := vars["chain"]
	if !ok {
		badRequest(res)
		return
	}
	address, ok := vars["address"]
	if !ok {
		badRequest(res)
		return
	}
	chain, ok := h.chainList[chainName]
	if !ok {
		resourceNotFound(res)
		return
	}

	// use rpc endpoint to get codeId from address
	codeId, err := getCodeIdfromAddress(address)

	if err != nil {
		// todo: error handling, rpc unavailable
	}

	contractInfo, exists := chain.Codes[codeId]

	if !exists {
		resourceNotFound(res) // todo: error handling, explicit "code does not exist"...
		return
	}

	respondWithJSON(res, contractInfo)
}
*/

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func resourceNotFound(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.WriteHeader(http.StatusNotFound)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func badRequest(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.WriteHeader(http.StatusBadRequest)
}
