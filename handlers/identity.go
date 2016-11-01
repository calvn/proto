package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func IdentityHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := session.NewSession()
	if err != nil {
		log.Println("failed to create session,", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	svc := sts.New(sess)

	resp, err := svc.GetCallerIdentity(nil)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&ErrorResponse{err.Error()})
		return
	}

	json.NewEncoder(w).Encode(resp)
	return
}
