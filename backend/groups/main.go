package main

import (
	"442GroupsProject/team-e/backend/groups/groupssrc"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	dsn, dsnok := os.LookupEnv("DSN")
	if dsnok == false {
		os.Stdout.WriteString("dsn not set")
		os.Exit(1)
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		os.Stdout.WriteString(err.Error())
		os.Stdout.WriteString("error opening database")
		os.Exit(1)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Printf("error pinging database: %v\n", err)
	} else {
		fmt.Printf("successfully connected!\n")
	}

	ctx := &groupssrc.GroupContext{}
	stdb := &groupssrc.SQLStore{}
	stdb.DB = db
	ctx.GStore = stdb

	mux := mux.NewRouter()

	mux.HandleFunc("/v1/categories", ctx.CategoriesHandler)
	mux.HandleFunc("/v1/categories/{categoryID}", ctx.SavedCategoriesHandler)

	mux.HandleFunc("/v1/groups", ctx.CreateGroupHandler)
	mux.HandleFunc("/v1/groups/{groupID}", ctx.GroupHandler)
	mux.HandleFunc("/v1/groups/{groupID}/save", ctx.SavedGroupHandler)
	mux.HandleFunc("/v1/groups/{groupID}/requests", ctx.GenericGroupMembershipHandler)
	mux.HandleFunc("/v1/groups/{groupID}/requests/{requestID}", ctx.GroupMembershipHandler)
	mux.HandleFunc("/v1/groups/{groupID}/comments", ctx.GenericCommentHandler)
	mux.HandleFunc("/v1/groups/{groupID}/comments/{commentID}", ctx.CommentHandler)
	mux.HandleFunc("/v1/groups/{groupID}/blog", ctx.GenericBlogPostHandler)
	mux.HandleFunc("/v1/groups/{groupID}/blog/{blogID}", ctx.BlogPostHandler)
	mux.HandleFunc("/v1/groups/{groupsID}/blog/{blogID}/comments", ctx.GenericBlogCommentHandler)
	mux.HandleFunc("/v1/groups/{groupsID}/blog/{blogID}/comments/{commentID}", ctx.BlogCommentHandler)

	mux.HandleFunc("/v1/search", ctx.GroupSearchHandler)
	mux.HandleFunc("/v1/admin", ctx.AdminHandler)

	log.Printf("server is lsitening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
