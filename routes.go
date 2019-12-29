package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (a *app) getBlockChain(w http.ResponseWriter, r *http.Request) {
	res, err := json.MarshalIndent(a.BC.blocks, "", " ")
	if err != nil {
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(res)
}

func (a *app) addBlock(w http.ResponseWriter, r *http.Request) {
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not write Block: %v", err)
		w.Write([]byte("could not write block"))
		return
	}

	a.BC.AddBlock(data)
	resp, err := json.MarshalIndent(a.BC.blocks, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("could not marshal payload: %v", err)
		w.Write([]byte("could not write block"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
