package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

type IdentityResponse struct {
	Identity *sts.GetCallerIdentityOutput `json:"identity"`
}

func IdentityHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := session.NewSession()
	if err != nil {
		log.Println("failed to create session,", err)
		return
	}

	svc := sts.New(sess)

	var params *sts.GetCallerIdentityInput
	resp, err := svc.GetCallerIdentity(params)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{err.Error()})
		return
	}

	response := IdentityResponse{
		Identity: resp,
	}
	json.NewEncoder(w).Encode(response)
	return
}
