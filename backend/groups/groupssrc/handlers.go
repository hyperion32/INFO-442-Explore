package groupssrc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
)

//All handlers take in a header that has the user making the request in addition to other inputs. This allows the back-end to know who is making the request; needed for creating new groups/comments/posts / knowing if the user is in a group/pending request in a group/is group admin etc.
//All inputs will be either part of the request url or in the request body. Specifying how exactly feels too implementation-y to be explained here.

//CategoriesHandler is the CategoryHandler, handles requests about categories.
func (ctx *GroupContext) CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	//GET: load categories
	//input: none
	//output: category struct(s), 200 status code
	if r.Method == http.MethodGet {
		rCat := &ReturnCategories{}

		userHeader := r.Header.Get("X-User")
		if len(userHeader) != 0 {
			user := &User{}
			errDecode := json.Unmarshal([]byte(userHeader), &user)
			if errDecode != nil {
				http.Error(w, "Error getting user", http.StatusInternalServerError)
				return
			}

			scts, errGetCts := ctx.GStore.GetSavedCategories(user.UserID)
			if errGetCts != nil {
				http.Error(w, errGetCts.Error(), http.StatusInternalServerError)
				return
			}

			rCat.SavedCategories = scts
		} else {
			scts := make([]*Category, 0)
			rCat.SavedCategories = scts
		}

		query, _ := r.URL.Query()["query"]
		if len(query) < 1 {
			query = append(query, "")
		}

		cts, errDB := ctx.GStore.GetCategories(query[0])
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}
		rCat.Categories = cts

		encoded, errEncode := json.Marshal(rCat)
		if errEncode != nil {
			http.Error(w, "Error encoding user to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//SavedCategoriesHandler handles requests to save and unsave categories
func (ctx *GroupContext) SavedCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	userHeader := r.Header.Get("X-User")
	if len(userHeader) == 0 {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user := &User{}
	errDecode := json.Unmarshal([]byte(userHeader), &user)
	if errDecode != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	pathid := path.Base(r.URL.Path)
	split := strings.Split(pathid, "&")
	pathid = split[0]
	pid, errConv := strconv.Atoi(pathid)
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		errSave := ctx.GStore.SaveCategory(pid, user.UserID)
		if errSave != nil {
			http.Error(w, "Error saving category", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("category saved"))
	} else if r.Method == http.MethodDelete {
		errDelete := ctx.GStore.UnsaveCategory(pid, user.UserID)
		if errDelete != nil {
			http.Error(w, "Error unsaving category", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("category unsaved"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//GroupSearchHandler loads groups based on a search query.
func (ctx *GroupContext) GroupSearchHandler(w http.ResponseWriter, r *http.Request) {
	//GET: Load search results
	//inputs: search query
	//outputs: group struct(s), 200 status code
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	user := &User{}

	userHeader := r.Header.Get("X-User")
	if len(userHeader) > 0 {
		errDecode := json.Unmarshal([]byte(userHeader), &user)
		if errDecode != nil {
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}
	} else {
		user.UserID = 0
	}

	query, _ := r.URL.Query()["query"]
	if len(query) < 1 {
		query = append(query, "")
	}

	category, _ := r.URL.Query()["category"]
	if len(category) < 1 {
		category = append(category, "0")
	}
	catnum, errConv := strconv.Atoi(category[0])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	page, _ := r.URL.Query()["page"]
	if len(page) < 1 {
		page = append(page, "1")
	}
	pagenum, errConv := strconv.Atoi(page[0])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	groups, errDB := ctx.GStore.SearchGroups(query[0], user.UserID, pagenum, catnum)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}

	encoded, errEncode := json.Marshal(groups)
	if errEncode != nil {
		http.Error(w, "Error encoding user to JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}

//CreateGroupHandler is the create group controller, handles requests for creating a new group.
func (ctx *GroupContext) CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	//POST: Create new group
	//inputs: group struct
	//output: group struct, 201 status code
	userHeader := r.Header.Get("X-User")
	if len(userHeader) == 0 {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user := &User{}
	errDecode := json.Unmarshal([]byte(userHeader), &user)
	if errDecode != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	group := &Group{}
	errDecode = json.NewDecoder(r.Body).Decode(&group)
	if errDecode != nil {
		http.Error(w, "Bad input", http.StatusBadRequest)
		return
	}

	group.User = user
	group, errDB := ctx.GStore.CreateGroup(group)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}

	encoded, errEncode := json.Marshal(group)
	if errEncode != nil {
		http.Error(w, "Error encoding user to JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(encoded)
}

//GroupHandler is the group controller, handles requests for existing groups.
func (ctx *GroupContext) GroupHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}

	userHeader := r.Header.Get("X-User")
	if len(userHeader) > 0 {
		errDecode := json.Unmarshal([]byte(userHeader), &user)
		if errDecode != nil {
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}
	} else {
		user.UserID = 0
	}

	pathid := path.Base(r.URL.Path)
	split := strings.Split(pathid, "&")
	pathid = split[0]
	gid, errConv := strconv.Atoi(pathid)
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	group, errDB := ctx.GStore.GetGroup(gid, user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}

	//GET: Load group information
	//inputs: group id
	//outputs: group struct, 200 status code
	if r.Method == http.MethodGet {
		comments, errDB := ctx.GStore.GetGroupCommentsByGroup(gid, 1)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}
		group.Comments = comments

		if group.User.UserID == user.UserID {
			group.IsAdmin = true
		} else {
			group.IsAdmin = false
		}

		count, errDB := ctx.GStore.GetGroupMembershipCount(gid)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}
		group.NumMembers = count

		encoded, errEncode := json.Marshal(group)
		if errEncode != nil {
			http.Error(w, "Error encoding user to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
		//DELETE: Delete existing group
		//inputs: group id
		//output: 200 status code
	} else if r.Method == http.MethodDelete {
		if group.User.UserID != user.UserID {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}

		errDelete := ctx.GStore.DeleteGroup(gid)
		if errDelete != nil {
			http.Error(w, errDelete.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Group deleted successfully"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//SavedGroupHandler is the SavedGroupController, handles request for a user's existing saved groups and to save a new group
func (ctx *GroupContext) SavedGroupHandler(w http.ResponseWriter, r *http.Request) {
	userHeader := r.Header.Get("X-User")
	if len(userHeader) == 0 {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user := &User{}
	errDecode := json.Unmarshal([]byte(userHeader), &user)
	if errDecode != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		errSave := ctx.GStore.SaveGroup(gid, user.UserID)
		if errSave != nil {
			http.Error(w, "Error saving group", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("group saved"))
		//DELETE: Remove a group from saved groups
		//inputs: group id
		//outputs: 200 status code
	} else if r.Method == http.MethodDelete {
		errDelete := ctx.GStore.UnsaveGroup(gid, user.UserID)
		if errDelete != nil {
			http.Error(w, "Error unsaving group", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("group unsaved"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//GenericCommentHandler handles requests to create a new group comment or gets comments from that group
func (ctx *GroupContext) GenericCommentHandler(w http.ResponseWriter, r *http.Request) {
	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	group, errDB := ctx.GStore.GetGroup(gid, 0)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}

	//Gets group comments
	if r.Method == http.MethodGet {
		page, _ := r.URL.Query()["page"]
		if len(page) < 1 {
			page = append(page, "1")
		}
		pagenum, errConv := strconv.Atoi(page[0])
		if errConv != nil {
			http.Error(w, "Not an integer", http.StatusBadRequest)
			return
		}

		comments, errDB := ctx.GStore.GetGroupCommentsByGroup(gid, pagenum)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		encoded, errEncode := json.Marshal(comments)
		if errEncode != nil {
			http.Error(w, "Error encoding user to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
	} else if r.Method == http.MethodPost {
		//POST: Create a new comment
		//inputs: comment struct
		//output: comment struct, 201 status code
		userHeader := r.Header.Get("X-User")
		if len(userHeader) == 0 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		user := &User{}
		errDecode := json.Unmarshal([]byte(userHeader), &user)
		if errDecode != nil {
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}

		gc := &GroupComment{}
		errDecode = json.NewDecoder(r.Body).Decode(&gc)
		if errDecode != nil {
			http.Error(w, "Bad input", http.StatusBadRequest)
			return
		}
		gc.GroupID = gid
		gc.User = user

		gc, errDB := ctx.GStore.CreateGroupComment(gc)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		encoded, errEncode := json.Marshal(gc)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(encoded)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//CommentHandler is the comment controller, handles requests for an existing comment
func (ctx *GroupContext) CommentHandler(w http.ResponseWriter, r *http.Request) {
	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	pathid = path.Base(r.URL.Path)
	split = strings.Split(pathid, "&")
	pathid = split[0]
	gcid, errConv := strconv.Atoi(pathid)
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	gc, errDB := ctx.GStore.GetGroupComment(gcid)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if gc.GroupID != gid {
		http.Error(w, "Comment is not for the specified group", http.StatusBadRequest)
		return
	}
	if gc == nil {
		http.Error(w, "Comment does not exist", http.StatusBadRequest)
		return
	}

	//GET: Load comment information
	//inputs: comment id
	//outputs: comment struct, 200 status code
	if r.Method == http.MethodGet {
		encoded, errEncode := json.Marshal(gc)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
		//DELETE: Delete existing comment
		//inputs: comment id
		//output: 200 status code
	} else if r.Method == http.MethodDelete {
		userHeader := r.Header.Get("X-User")
		if len(userHeader) == 0 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		user := &User{}
		errDecode := json.Unmarshal([]byte(userHeader), &user)
		if errDecode != nil {
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}

		if gc.User.UserID != user.UserID {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}

		errDelete := ctx.GStore.DeleteGroupComment(gcid)
		if errDelete != nil {
			http.Error(w, errDelete.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Comment deleted successfully"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//GenericBlogPostHandler handles requests to create new blog posts
func (ctx *GroupContext) GenericBlogPostHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	userHeader := r.Header.Get("X-User")
	if len(userHeader) > 0 {
		errDecode := json.Unmarshal([]byte(userHeader), &user)
		if errDecode != nil {
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}
	} else {
		user.UserID = 0
	}

	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	group, errDB := ctx.GStore.GetGroup(gid, user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {
		page, _ := r.URL.Query()["page"]
		if len(page) < 1 {
			page = append(page, "1")
		}
		pagenum, errConv := strconv.Atoi(page[0])
		if errConv != nil {
			http.Error(w, "Not an integer", http.StatusBadRequest)
			return
		}

		blogs, errDB := ctx.GStore.GetBlogPostsByGroup(gid, pagenum)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		encoded, errEncode := json.Marshal(blogs)
		if errEncode != nil {
			http.Error(w, "Error encoding user to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
		//POST: Create a new blog post
		//inputs: blog post struct
		//output: blog post struct, 201 status code
	} else if r.Method == http.MethodPost {
		fmt.Println(group.User.UserID)
		fmt.Println(user.UserID)
		if group.JoinedState != "Joined" && group.User.UserID != user.UserID {
			http.Error(w, "Access forbidden, not in group", http.StatusForbidden)
			return
		}

		bp := &BlogPost{}
		errDecode := json.NewDecoder(r.Body).Decode(&bp)
		if errDecode != nil {
			http.Error(w, "Bad input", http.StatusBadRequest)
			return
		}
		bp.GroupID = gid
		bp.User = user

		bp, errDB = ctx.GStore.CreateBlogPost(bp)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		encoded, errEncode := json.Marshal(bp)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(encoded)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

//BlogPostHandler is the blog controller, handles requests for an existing blog post
func (ctx *GroupContext) BlogPostHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	userHeader := r.Header.Get("X-User")
	fmt.Println(len(userHeader))
	if len(userHeader) > 0 {
		errDecode := json.Unmarshal([]byte(userHeader), &user)
		if errDecode != nil {
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}
	} else {
		user.UserID = 0
	}

	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	pathid = path.Base(r.URL.Path)
	split = strings.Split(pathid, "&")
	pathid = split[0]
	bpid, errConv := strconv.Atoi(pathid)
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	group, errDB := ctx.GStore.GetGroup(gid, user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}

	bp, errDB := ctx.GStore.GetBlogPost(bpid)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if bp == nil {
		http.Error(w, "Blog post does not exist", http.StatusBadRequest)
		return
	}
	if bp.GroupID != gid {
		http.Error(w, "Blog post is not for the specified group", http.StatusBadRequest)
		return
	}

	//GET: Load blog post information
	//inputs: blog post id
	//outputs: blog struct, 200 status code
	if r.Method == http.MethodGet {
		encoded, errEncode := json.Marshal(bp)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
		//DELETE: Deletes an existing blog post
		//inputs: blog post id
		//output: 200 status code
	} else if r.Method == http.MethodDelete {
		fmt.Println(group.User.UserID)
		fmt.Println(user.UserID)
		if group.JoinedState != "Joined" && user.UserID != group.User.UserID { //2nd is if user is group admin
			http.Error(w, "Access forbidden, not in group", http.StatusForbidden)
			return
		}

		errDelete := ctx.GStore.DeleteBlogPost(bpid)
		if errDelete != nil {
			http.Error(w, errDelete.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Comment deleted successfully"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//CreateBlogCommentHandler handles requests to create new blog comments
func (ctx *GroupContext) CreateBlogCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	userHeader := r.Header.Get("X-User")
	if len(userHeader) == 0 {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user := &User{}
	errDecode := json.Unmarshal([]byte(userHeader), &user)
	if errDecode != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	pathid := path.Base(r.URL.Path)
	split := strings.Split(pathid, "&")
	pathid = split[0]
	/*pid, errConv := strconv.Atoi(pathid)
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}
	*/
}

//GenericBlogCommentHandler handles requests the specificed blog comment
func (ctx *GroupContext) GenericBlogCommentHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	userHeader := r.Header.Get("X-User")
	if len(userHeader) > 0 {
		errDecode := json.Unmarshal([]byte(userHeader), &user)
		if errDecode != nil {
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}
	} else {
		user.UserID = 0
	}

	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}
	bpid, errConv := strconv.Atoi(split[5])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	group, errDB := ctx.GStore.GetGroup(gid, user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}

	bp, errDB := ctx.GStore.GetBlogPost(bpid)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if bp == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}
	if bp.GroupID != group.GroupID {
		http.Error(w, "Blog post is not in the specified group", http.StatusBadRequest)
		return
	}

	//Gets group comments
	if r.Method == http.MethodGet {
		page, _ := r.URL.Query()["page"]
		if len(page) < 1 {
			page = append(page, "1")
		}
		pagenum, errConv := strconv.Atoi(page[0])
		if errConv != nil {
			http.Error(w, "Not an integer", http.StatusBadRequest)
			return
		}

		comments, errDB := ctx.GStore.GetBlogCommentsByBlog(bpid, pagenum)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		encoded, errEncode := json.Marshal(comments)
		if errEncode != nil {
			http.Error(w, "Error encoding user to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
	} else if r.Method == http.MethodPost {
		//POST: Create a new comment
		//inputs: comment struct
		//output: comment struct, 201 status code
		if user.UserID < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		if group.JoinedState != "Joined" && group.User.UserID != user.UserID {
			http.Error(w, "Cannot comment unless in group", http.StatusUnauthorized)
			return
		}

		bc := &BlogComment{}
		errDecode := json.NewDecoder(r.Body).Decode(&bc)
		if errDecode != nil {
			http.Error(w, "Bad input", http.StatusBadRequest)
			return
		}
		bc.BlogPostID = bpid
		bc.User = user

		bc, errDB := ctx.GStore.CreateBlogComment(bc)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		encoded, errEncode := json.Marshal(bc)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(encoded)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//BlogCommentHandler handles requests the specificed blog comment
func (ctx *GroupContext) BlogCommentHandler(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	userHeader := r.Header.Get("X-User")
	if len(userHeader) > 0 {
		errDecode := json.Unmarshal([]byte(userHeader), &user)
		if errDecode != nil {
			http.Error(w, "Error getting user", http.StatusInternalServerError)
			return
		}
	} else {
		user.UserID = 0
	}

	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}
	bpid, errConv := strconv.Atoi(split[5])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}
	pathid = path.Base(r.URL.Path)
	split = strings.Split(pathid, "&")
	pathid = split[0]
	bcid, errConv := strconv.Atoi(pathid)
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	group, errDB := ctx.GStore.GetGroup(gid, user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}

	bp, errDB := ctx.GStore.GetBlogPost(bpid)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if bp == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}
	if bp.GroupID != group.GroupID {
		http.Error(w, "Blog post is not in the specified group", http.StatusBadRequest)
		return
	}

	bc, errDB := ctx.GStore.GetBlogComment(bcid)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if bc.BlogPostID != bpid {
		http.Error(w, "Blog comment is not for the specified blog", http.StatusBadRequest)
		return
	}
	if bc == nil {
		http.Error(w, "Comment does not exist", http.StatusBadRequest)
		return
	}

	//GET: Load comment information
	//inputs: comment id
	//outputs: comment struct, 200 status code
	if r.Method == http.MethodGet {
		encoded, errEncode := json.Marshal(bc)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
		//DELETE: Delete existing comment
		//inputs: comment id
		//output: 200 status code
	} else if r.Method == http.MethodDelete {
		//you created comment / you are group admin
		if bc.User.UserID != user.UserID && bc.User.UserID != group.User.UserID {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}

		errDelete := ctx.GStore.DeleteBlogComment(bcid)
		if errDelete != nil {
			http.Error(w, errDelete.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Comment deleted successfully"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//GenericGroupMembershipHandler handles requests to specific membership requests
func (ctx *GroupContext) GenericGroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
	userHeader := r.Header.Get("X-User")
	if len(userHeader) == 0 {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user := &User{}
	errDecode := json.Unmarshal([]byte(userHeader), &user)
	if errDecode != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	//GET: Get all current membership requests
	//inputs: group id
	//outputs: membership request struct(s), 200 status code
	if r.Method == http.MethodGet {
		group, errDB := ctx.GStore.GetGroup(gid, user.UserID)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}
		if group == nil {
			http.Error(w, "Group does not exist", http.StatusBadRequest)
			return
		}
		if group.User.UserID != user.UserID {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		mrs, errDB := ctx.GStore.GetMembershipRequests(gid)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		encoded, errEncode := json.Marshal(mrs)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
		//POST: Create a new request for a group
	} else if r.Method == http.MethodPost {
		mr, errDB := ctx.GStore.CreateMembershipRequest(gid, user.UserID)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		encoded, errEncode := json.Marshal(mr)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(encoded)
	} else if r.Method == http.MethodDelete {
		errDB := ctx.GStore.LeaveGroup(gid, user.UserID)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("request recinded"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//GroupMembershipHandler is the membership controller, handles requests to join a group and requests to confirm or reject someone from the group
func (ctx *GroupContext) GroupMembershipHandler(w http.ResponseWriter, r *http.Request) {
	userHeader := r.Header.Get("X-User")
	if len(userHeader) == 0 {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user := &User{}
	errDecode := json.Unmarshal([]byte(userHeader), &user)
	if errDecode != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	pathid := r.URL.Path
	split := strings.Split(pathid, "/")
	gid, errConv := strconv.Atoi(split[3])
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	group, errDB := ctx.GStore.GetGroup(gid, user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if group == nil {
		http.Error(w, "Group does not exist", http.StatusBadRequest)
		return
	}

	pathid = path.Base(r.URL.Path)
	split = strings.Split(pathid, "&")
	pathid = split[0]
	mrid, errConv := strconv.Atoi(pathid)
	if errConv != nil {
		http.Error(w, "Not an integer", http.StatusBadRequest)
		return
	}

	mr, errDB := ctx.GStore.GetMembershipRequest(mrid)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	if mr == nil {
		http.Error(w, "Request does not exist", http.StatusBadRequest)
		return
	}
	if mr.GroupID != group.GroupID {
		http.Error(w, "request is not in group", http.StatusBadRequest)
		return
	}

	//GET: Get a membership request by id
	//inputs: group id
	//outputs: membership request struct(s), 200 status code
	if r.Method == http.MethodGet {
		encoded, errEncode := json.Marshal(mr)
		if errEncode != nil {
			http.Error(w, "Error encoding to JSON", http.StatusBadRequest)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(encoded)
		//PATCH: Update membership request from pending to confirmed or
		//inputs: group id, membership request id
		//outputs: 200 status code
	} else if r.Method == http.MethodPatch {
		if group.User.UserID != user.UserID {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		errDB := ctx.GStore.AcceptMembershipRequest(mrid)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("request accepted"))
	} else if r.Method == http.MethodDelete {
		if group.User.UserID != user.UserID {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		errDB := ctx.GStore.DeclineMembershipRequest(mrid)
		if errDB != nil {
			http.Error(w, errDB.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("request declined"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

//AdminHandler handles requests to get admin groups info
func (ctx *GroupContext) AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	userHeader := r.Header.Get("X-User")
	if len(userHeader) == 0 {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	user := &User{}
	errDecode := json.Unmarshal([]byte(userHeader), &user)
	if errDecode != nil {
		http.Error(w, "Error getting user", http.StatusInternalServerError)
		return
	}

	ra := &ReturnAdmin{}

	ag, errDB := ctx.GStore.GetAdminGroups(user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	ra.AdminGroups = ag

	jg, errDB := ctx.GStore.GetJoinedGroups(user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	ra.JoinedGroups = jg

	sg, errDB := ctx.GStore.GetSavedGroups(user.UserID)
	if errDB != nil {
		http.Error(w, errDB.Error(), http.StatusInternalServerError)
		return
	}
	ra.SavedGroups = sg

	encoded, errEncode := json.Marshal(ra)
	if errEncode != nil {
		http.Error(w, "Error encoding user to JSON", http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encoded)
}
