package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"442GroupsProject/team-e/backend/gateway/gatewaysrc"

	"github.com/gorilla/mux"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath, keyok := os.LookupEnv("TLSKEY")
	tlsCertPath, certok := os.LookupEnv("TLSCERT")

	if keyok == false || certok == false {
		os.Stdout.WriteString("key/certificate not set")
		os.Exit(1)
	}

	groupsaddr, ok := os.LookupEnv("GROUPSADDR")
	if ok == false {
		os.Stdout.WriteString("groupsaddr not set")
		os.Exit(1)
	}

	dsn, dsnok := os.LookupEnv("DSN")
	if dsnok == false {
		os.Stdout.WriteString("dsn not set")
		os.Exit(1)
	}

	if dsnok == false {
		os.Stdout.WriteString("dsn not set")
		os.Exit(1)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		os.Stdout.WriteString("error opening database")
		os.Exit(1)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Printf("error pinging database: %v\n", err)
	} else {
		fmt.Printf("successfully connected!\n")
	}

	ctx := &gatewaysrc.GatewayContext{}
	stdb := &gatewaysrc.SQLStore{}
	stdb.DB = db
	ctx.GStore = stdb

	groupsDirector := func(r *http.Request) {
		serverName := groupsaddr

		r.Header.Del("X-User")
		guser, rcode := GetGoogleUserInfo(r)
		if rcode == 200 {
			user, errGetUserInfo := ctx.GStore.GetUserInfo(guser)
			fmt.Println(user.UserID)
			if errGetUserInfo == nil && user != nil {
				encoded, _ := json.Marshal(user)
				r.Header.Set("X-User", string(encoded))
			} else {
				fmt.Println(errGetUserInfo.Error())
			}
		}

		r.Host = serverName
		r.URL.Host = serverName
		r.URL.Scheme = "http"
	}

	mux := mux.NewRouter()

	groupsProxy := &httputil.ReverseProxy{Director: groupsDirector}
	mux.Handle("/v1/categories", groupsProxy)
	mux.Handle("/v1/categories/{categoryID}", groupsProxy)
	mux.Handle("/v1/groups", groupsProxy)
	mux.Handle("/v1/groups/{groupID}", groupsProxy)
	mux.Handle("/v1/groups/{groupID}/save", groupsProxy)
	mux.Handle("/v1/search", groupsProxy)
	mux.Handle("/v1/groups/{groupID}/comments", groupsProxy)
	mux.Handle("/v1/groups/{groupID}/comments/{commentID}", groupsProxy)
	mux.Handle("/v1/groups/{groupID}/blog", groupsProxy)
	mux.Handle("/v1/groups/{groupID}/blog/{blogID}", groupsProxy)

	mux.Handle("/v1/groups/{groupID}/requests", groupsProxy)
	mux.Handle("/v1/groups/{groupID}/requests/{requestID}", groupsProxy)
	mux.Handle("/v1/groups/{groupsID}/blog/{blogID}/comments", groupsProxy)
	mux.Handle("/v1/groups/{groupsID}/blog/{blogID}/comments/{commentID}", groupsProxy)
	mux.Handle("/v1/admin", groupsProxy)

	mux.HandleFunc("/", gatewaysrc.HandleHome)
	mux.HandleFunc("/login", gatewaysrc.HandleLogin)
	mux.HandleFunc("/callback", gatewaysrc.HandleCallback)

	wrappedMux := gatewaysrc.AddFiveResponseHeaders(mux, "Access-Control-Allow-Origin", "*", "Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE", "Access-Control-Allow-Headers", "Content-Type, Authorization", "Access-Control-Expose-Headers", "Authorization", "Access-Control-Max-Age", "600")

	log.Printf("server is lsitening at %s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))
}

//GetGoogleUserInfo returns the google info for the auth token
func GetGoogleUserInfo(r *http.Request) (*gatewaysrc.GoogleUser, int) {
	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		query, ok := r.URL.Query()["auth"]
		if !ok || len(query[0]) < 1 {
			return nil, 666
		}
		accessToken = query[0]
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, resp.StatusCode
	}

	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		return nil, resp.StatusCode
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 777
	}

	defer resp.Body.Close()

	guser := &gatewaysrc.GoogleUser{}
	json.Unmarshal(responseData, &guser)

	fmt.Println(guser.GoogleID)

	return guser, 200
}
