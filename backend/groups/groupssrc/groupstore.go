package groupssrc

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" //needed
)

//SQLStore holds db info
type SQLStore struct {
	DB *sql.DB
}

//CATEGORY DB METHODS

//GetCategories returns all categories
func (sqls *SQLStore) GetCategories(query string) ([]*Category, error) {
	cts := make([]*Category, 0)
	wcstring := string('%') + query + string('%')
	insq := "select * from category where category_name LIKE ? order by category_name"

	res, errQuery := sqls.DB.Query(insq, wcstring)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {

		ct := &Category{}
		errScan := res.Scan(&ct.CategoryID, &ct.CategoryName)
		if errScan != nil {
			return nil, errScan
		}
		cts = append(cts, ct)
	}

	return cts, nil
}

//SaveCategory allows for the ability to save category. Takes an array so multiple can be saved in the same request.
func (sqls *SQLStore) SaveCategory(catid int, userid int) error {
	insq := "insert into saved_category(category_id, user_id) values(?,?)"

	_, errExec := sqls.DB.Exec(insq, catid, userid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//UnsaveCategory unsaves a previously saved category
func (sqls *SQLStore) UnsaveCategory(catid int, userid int) error {
	insq := "delete from saved_category where category_id = ? and user_id = ?"

	_, errExec := sqls.DB.Exec(insq, catid, userid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//GetSavedCategories gets the saved categories for a given user
func (sqls *SQLStore) GetSavedCategories(userid int) ([]*Category, error) {
	cts := make([]*Category, 0)
	insq := "select sc.category_id, c.category_name from saved_category sc join category c on sc.category_id = c.category_id"

	res, errQuery := sqls.DB.Query(insq)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {
		ct := &Category{}
		errScan := res.Scan(&ct.CategoryID, &ct.CategoryName)
		if errScan != nil {
			return nil, errScan
		}
		cts = append(cts, ct)
	}

	return cts, nil
}

//GROUP DB METHODS

//CreateGroup creates a group
func (sqls *SQLStore) CreateGroup(gp *Group) (*Group, error) {
	insq := "insert into `group`(user_id, category_id, group_name, group_description, created_at) values(?,?,?,?,?)"

	res, errExec := sqls.DB.Exec(insq, gp.User.UserID, gp.Category.CategoryID, gp.GroupName, gp.GroupDescription, time.Now())
	if errExec != nil {
		return nil, errExec
	}

	gpid, errID := res.LastInsertId()
	if errID != nil {
		return nil, errID
	}
	gp.GroupID = int(gpid)

	return gp, nil
}

//SearchGroups searches groups with the given query term, and returns groups with similar group names to the query.
func (sqls *SQLStore) SearchGroups(query string, userid int, page int, catid int) ([]*Group, error) {
	gps := make([]*Group, 0)
	wcstring := string('%') + query + string('%')
	offset := 9 * (page - 1)

	var res *sql.Rows
	var errQuery error
	var insq string
	if catid != 0 {
		insq = "select g.group_id, g.user_id, g.category_id, g.group_name, g.group_description, g.created_at, c.category_name, u.first_name, u.last_name, u.photo_url from `group` g join category c on g.category_id = c.category_id join user u on g.user_id = u.user_id where g.group_name LIKE ? and g.category_id = ? order by g.group_name limit 9 offset ?"

		res, errQuery = sqls.DB.Query(insq, wcstring, catid, offset)
		if errQuery != nil {
			return nil, errQuery
		}
	} else {
		insq = "select g.group_id, g.user_id, g.category_id, g.group_name, g.group_description, g.created_at, c.category_name, u.first_name, u.last_name, u.photo_url from `group` g join category c on g.category_id = c.category_id join user u on g.user_id = u.user_id where g.group_name LIKE ? order by g.group_name limit 9 offset ?"

		res, errQuery = sqls.DB.Query(insq, wcstring, offset)
		if errQuery != nil {
			return nil, errQuery
		}
	}

	defer res.Close()

	for res.Next() {
		gp := &Group{}
		user := &User{}
		c := &Category{}
		errScan := res.Scan(&gp.GroupID, &user.UserID, &c.CategoryID, &gp.GroupName, &gp.GroupDescription, &gp.CreatedAt, &c.CategoryName, &user.FirstName, &user.LastName, &user.PhotoURL)
		if errScan != nil {
			return nil, errScan
		}
		gp.User = user
		gp.Category = c

		sgid := 0
		insq = "select sg_id from saved_group where group_id = ? and user_id = ?"
		errQuery = sqls.DB.QueryRow(insq, gp.GroupID, userid).Scan(&sgid)
		if errQuery != nil {
			if errQuery == sql.ErrNoRows {
				gp.IsSaved = false
			} else {
				return nil, errQuery
			}
		} else {
			gp.IsSaved = true
		}

		var state bool
		insq = "select state from membership where group_id = ? and user_id = ?"
		errQuery = sqls.DB.QueryRow(insq, gp.GroupID, userid).Scan(&state)
		if errQuery != nil {
			if errQuery != sql.ErrNoRows {
				return nil, errQuery
			}
			gp.JoinedState = "NA"
		} else {
			if state {
				gp.JoinedState = "Joined"
			} else {
				gp.JoinedState = "Pending"
			}
		}

		gps = append(gps, gp)
	}

	return gps, nil
}

//GetGroup gets a group by groupid and returns it
func (sqls *SQLStore) GetGroup(gpid int, userid int) (*Group, error) {
	gp := &Group{}
	user := &User{}
	c := &Category{}

	insq := "select g.group_id, g.user_id, g.category_id, g.group_name, g.group_description, g.created_at, c.category_name, u.first_name, u.last_name, u.photo_url from `group` g join category c on g.category_id = c.category_id join user u on g.user_id = u.user_id where group_id = ?"

	errQuery := sqls.DB.QueryRow(insq, gpid).Scan(&gp.GroupID, &user.UserID, &c.CategoryID, &gp.GroupName, &gp.GroupDescription, &gp.CreatedAt, &c.CategoryName, &user.FirstName, &user.LastName, &user.PhotoURL)
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errQuery
	}
	gp.User = user
	gp.Category = c

	if userid != 0 {
		insq = "select group_id from saved_group where group_id = ? and user_id = ?"
		errQuery = sqls.DB.QueryRow(insq, gpid, userid).Scan(&gp.GroupID)
		if errQuery != nil {
			if errQuery == sql.ErrNoRows {
				gp.IsSaved = false
			} else {
				return nil, errQuery
			}
		} else {
			gp.IsSaved = true
		}

		var state bool
		insq = "select state from membership where group_id = ? and user_id = ?"
		errQuery = sqls.DB.QueryRow(insq, gpid, userid).Scan(&state)
		if errQuery != nil {
			if errQuery != sql.ErrNoRows {
				return nil, errQuery
			}
			gp.JoinedState = "NA"
		} else {
			if state {
				gp.JoinedState = "Joined"
			} else {
				gp.JoinedState = "Pending"
			}
		}

	}

	return gp, nil
}

//DeleteGroup deletes a group by groupid
func (sqls *SQLStore) DeleteGroup(gpid int) error {
	//hopefully this is everything we gotta delete.

	insq := "delete from blog_post where group_id = ?"

	_, errExec := sqls.DB.Exec(insq, gpid)
	if errExec != nil {
		return errExec
	}

	insq = "delete from group_comment where group_id = ?"

	_, errExec = sqls.DB.Exec(insq, gpid)
	if errExec != nil {
		return errExec
	}

	insq = "delete from membership where group_id = ?"

	_, errExec = sqls.DB.Exec(insq, gpid)
	if errExec != nil {
		return errExec
	}

	insq = "delete from saved_group where group_id = ?"

	_, errExec = sqls.DB.Exec(insq, gpid)
	if errExec != nil {
		return errExec
	}

	insq = "delete from `group` where group_id = ?"

	_, errExec = sqls.DB.Exec(insq, gpid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//SaveGroup saves a specific group to a specific user's saved groups
func (sqls *SQLStore) SaveGroup(gpid int, userid int) error {
	insq := "insert into saved_group(group_id, user_id) values(?,?)"

	_, errExec := sqls.DB.Exec(insq, gpid, userid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//UnsaveGroup unsaves a specific group to a specific user's saved groups
func (sqls *SQLStore) UnsaveGroup(gpid int, userid int) error {
	insq := "delete from saved_group where group_id = ? and user_id = ?"

	_, errExec := sqls.DB.Exec(insq, gpid, userid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//GetSavedGroups retuns all groups a given user has saved
func (sqls *SQLStore) GetSavedGroups(userid int) ([]*Group, error) {
	gps := make([]*Group, 0)

	insq := "select sg.group_id, u.user_id, u.first_name, u.last_name, u.photo_url, c.category_id, c.category_name, g.group_name, g.group_description, g.created_at from saved_group sg join `group` g on sg.group_id = g.group_id join category c on g.category_id = c.category_id join user u on g.user_id = u.user_id where sg.user_id = ?"

	res, errQuery := sqls.DB.Query(insq, userid)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {
		gp := &Group{}
		user := &User{}
		c := &Category{}
		errScan := res.Scan(&gp.GroupID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &c.CategoryID, &c.CategoryName, &gp.GroupName, &gp.GroupDescription, &gp.CreatedAt)
		if errScan != nil {
			return nil, errScan
		}
		gp.User = user
		gps = append(gps, gp)
	}

	return gps, nil
}

//GetJoinedGroups gets all the groups you have joined
func (sqls *SQLStore) GetJoinedGroups(userid int) ([]*Group, error) {
	gps := make([]*Group, 0)

	insq := "select m.group_id, u.user_id, u.first_name, u.last_name, u.photo_url, c.category_id, c.category_name, g.group_name, g.group_description, g.created_at from membership m join `group` g on m.group_id = g.group_id join category c on g.category_id = c.category_id join user u on g.user_id = u.user_id where m.user_id = ? and state = true"

	res, errQuery := sqls.DB.Query(insq, userid)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {
		gp := &Group{}
		user := &User{}
		c := &Category{}
		errScan := res.Scan(&gp.GroupID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &c.CategoryID, &c.CategoryName, &gp.GroupName, &gp.GroupDescription, &gp.CreatedAt)
		if errScan != nil {
			return nil, errScan
		}
		gp.User = user
		gps = append(gps, gp)
	}

	return gps, nil
}

//GetAdminGroups gets all groups you created
func (sqls *SQLStore) GetAdminGroups(userid int) ([]*Group, error) {
	gps := make([]*Group, 0)

	insq := "select g.group_id, g.user_id, u.first_name, u.last_name, u.photo_url, g.category_id, c.category_name, g.group_name, g.group_description, g.created_at from `group` g join category c on g.category_id = c.category_id join user u on g.user_id = u.user_id where g.user_id = ?"

	res, errQuery := sqls.DB.Query(insq, userid)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {
		gp := &Group{}
		user := &User{}
		c := &Category{}
		errScan := res.Scan(&gp.GroupID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &c.CategoryID, &c.CategoryName, &gp.GroupName, &gp.GroupDescription, &gp.CreatedAt)
		if errScan != nil {
			return nil, errScan
		}
		gp.User = user
		gps = append(gps, gp)
	}

	return gps, nil
}

func (sqls *SQLStore) GetGroupMembershipCount(gid int) (int, error) {
	insq := "select * from membership where group_id = ? and state = true"

	res, errQuery := sqls.DB.Query(insq, gid)
	if errQuery != nil {
		return 0, errQuery
	}
	defer res.Close()

	count := 0
	for res.Next() {
		count += 1
	}

	return count, nil
}

//GROUPCOMMENT DB METHODS

//CreateGroupComment creates a group comment
func (sqls *SQLStore) CreateGroupComment(gc *GroupComment) (*GroupComment, error) {
	if gc.ReplyID.Int64 == 0 {
		insq := "insert into group_comment(user_id, group_id, comment_content, created_at, deleted) values(?,?,?,?,?)"

		res, errExec := sqls.DB.Exec(insq, gc.User.UserID, gc.GroupID, gc.CommentContent, time.Now(), false)
		if errExec != nil {
			return nil, errExec
		}

		gcid, errID := res.LastInsertId()
		if errID != nil {
			return nil, errID
		}
		gc.GroupCommentID = int(gcid)

		return gc, nil
	}

	insq := "insert into group_comment(user_id, group_id, reply_id, comment_content, created_at, deleted) values(?,?,?,?,?,?)"
	res, errExec := sqls.DB.Exec(insq, gc.User.UserID, gc.GroupID, gc.ReplyID, gc.CommentContent, time.Now(), false)
	if errExec != nil {
		return nil, errExec
	}

	gcid, errID := res.LastInsertId()
	if errID != nil {
		return nil, errID
	}
	gc.GroupCommentID = int(gcid)

	return gc, nil
}

//GetGroupComment gets a group comment by groupcomment id
func (sqls *SQLStore) GetGroupComment(gcid int) (*GroupComment, error) {
	gc := &GroupComment{}
	user := &User{}

	insq := "select gc.gc_id, gc.user_id, u.first_name, u.last_name, u.photo_url, gc.group_id, gc.reply_id, gc.comment_content, gc.created_at, gc.deleted from group_comment gc join user u on gc.user_id = u.user_id where gc.gc_id = ?"

	errQuery := sqls.DB.QueryRow(insq, gcid).Scan(&gc.GroupCommentID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &gc.GroupID, &gc.ReplyID, &gc.CommentContent, &gc.CreatedAt, &gc.Deleted)
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errQuery
	}
	gc.User = user

	return gc, nil
}

//GetGroupCommentsByGroup gets comments for a specified group
func (sqls *SQLStore) GetGroupCommentsByGroup(gid int, page int) ([]*GroupComment, error) {
	gcs := make([]*GroupComment, 0)
	offset := 3 * (page - 1)

	insq := "select gc.gc_id, gc.user_id, u.first_name, u.last_name, u.photo_url, gc.group_id, gc.reply_id, gc.comment_content, gc.created_at, gc.deleted from group_comment gc join user u on gc.user_id = u.user_id where gc.group_id = ? and gc.reply_id IS NULL order by gc.gc_id DESC limit 3 offset ?"

	res, errQuery := sqls.DB.Query(insq, gid, offset)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {
		gc := &GroupComment{}
		user := &User{}
		errScan := res.Scan(&gc.GroupCommentID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &gc.GroupID, &gc.ReplyID, &gc.CommentContent, &gc.CreatedAt, &gc.Deleted)
		if errScan != nil {
			return nil, errScan
		}
		gc.User = user

		gc, errRec := sqls.GetChildrenComments(gc)
		if errRec != nil {
			return nil, errRec
		}

		gcs = append(gcs, gc)
	}

	return gcs, nil
}

//GetChildrenComments is a recursive function that gets all child comments for a parent.
func (sqls *SQLStore) GetChildrenComments(gc *GroupComment) (*GroupComment, error) {
	children := make([]*GroupComment, 0)

	insq := "select gc.gc_id, gc.user_id, u.first_name, u.last_name, u.photo_url, gc.group_id, gc.reply_id, gc.comment_content, gc.created_at, gc.deleted from group_comment gc join user u on gc.user_id = u.user_id where gc.reply_id = ? order by gc.gc_id DESC"

	res, errQuery := sqls.DB.Query(insq, gc.GroupCommentID)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	count := 0
	for res.Next() {
		count++
		child := &GroupComment{}
		user := &User{}
		errScan := res.Scan(&child.GroupCommentID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &child.GroupID, &child.ReplyID, &child.CommentContent, &child.CreatedAt, &child.Deleted)
		if errScan != nil {
			return nil, errScan
		}
		child.User = user

		gc, errRec := sqls.GetChildrenComments(child)
		if errRec != nil {
			return nil, errRec
		}

		children = append(children, gc)
	}

	if count != 0 {
		gc.Children = children
	}

	return gc, nil
}

//DeleteGroupComment deletes a group comment
func (sqls *SQLStore) DeleteGroupComment(gcid int) error {
	insq := "update group_comment set deleted = true where gc_id = ?"

	_, errExec := sqls.DB.Exec(insq, gcid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//BLOGCOMMENT DB METHODS

//CreateBlogComment creates a blog comment
func (sqls *SQLStore) CreateBlogComment(bc *BlogComment) (*BlogComment, error) {
	if bc.ReplyID.Int64 == 0 {
		insq := "insert into blog_comment(user_id, bp_id, comment_content, created_at, deleted) values(?,?,?,?,?)"

		res, errExec := sqls.DB.Exec(insq, bc.User.UserID, bc.BlogPostID, bc.CommentContent, time.Now(), false)
		if errExec != nil {
			return nil, errExec
		}

		bcid, errID := res.LastInsertId()
		if errID != nil {
			return nil, errID
		}
		bc.BlogCommentID = int(bcid)

		return bc, nil
	}

	insq := "insert into blog_comment(user_id, bp_id, reply_id, comment_content, created_at, deleted) values(?,?,?,?,?,?)"

	res, errExec := sqls.DB.Exec(insq, bc.User.UserID, bc.BlogPostID, bc.ReplyID, bc.CommentContent, time.Now(), false)
	if errExec != nil {
		return nil, errExec
	}

	bcid, errID := res.LastInsertId()
	if errID != nil {
		return nil, errID
	}
	bc.BlogCommentID = int(bcid)

	return bc, nil
}

//GetBlogComment gets a blog comment
func (sqls *SQLStore) GetBlogComment(bcid int) (*BlogComment, error) {
	bc := &BlogComment{}
	user := &User{}

	insq := "select bc.bc_id, bc.user_id, u.first_name, u.last_name, u.photo_url, bc.bp_id, bc.reply_id, bc.comment_content, bc.created_at, bc.deleted from blog_comment bc join user u on bc.user_id = u.user_id where bc.bc_id = ?"

	errQuery := sqls.DB.QueryRow(insq, bcid).Scan(&bc.BlogCommentID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &bc.BlogPostID, &bc.ReplyID, &bc.CommentContent, &bc.CreatedAt, &bc.Deleted)
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errQuery
	}
	bc.User = user

	return bc, nil
}

//DeleteBlogComment deletes a specified blog comment
func (sqls *SQLStore) DeleteBlogComment(bcid int) error {
	insq := "update blog_comment set deleted = true where bc_id = ?"

	_, errExec := sqls.DB.Exec(insq, bcid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//GetBlogCommentsByBlog returns comments for a specific blog post
func (sqls *SQLStore) GetBlogCommentsByBlog(bid int, page int) ([]*BlogComment, error) {
	bcs := make([]*BlogComment, 0)
	offset := 3 * (page - 1)

	insq := "select bc.bc_id, bc.user_id, u.first_name, u.last_name, u.photo_url, bc.bp_id, bc.reply_id, bc.comment_content, bc.created_at, bc.deleted from blog_comment bc join user u on bc.user_id = u.user_id where bc.bp_id = ? and bc.reply_id IS NULL order by bc.bc_id DESC limit 3 offset ?"

	res, errQuery := sqls.DB.Query(insq, bid, offset)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {
		bc := &BlogComment{}
		user := &User{}
		errScan := res.Scan(&bc.BlogCommentID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &bc.BlogPostID, &bc.ReplyID, &bc.CommentContent, &bc.CreatedAt, &bc.Deleted)
		if errScan != nil {
			return nil, errScan
		}
		bc.User = user

		bc, errRec := sqls.GetBlogChildrenComments(bc)
		if errRec != nil {
			return nil, errRec
		}

		bcs = append(bcs, bc)
	}

	return bcs, nil
}

//GetBlogChildrenComments is a recurive methods that gets the children blog comments
func (sqls *SQLStore) GetBlogChildrenComments(bc *BlogComment) (*BlogComment, error) {
	children := make([]*BlogComment, 0)

	insq := "select bc.bc_id, bc.user_id, u.first_name, u.last_name, u.photo_url, bc.bp_id, bc.reply_id, bc.comment_content, bc.created_at, bc.deleted from blog_comment bc join user u on bc.user_id = u.user_id where bc.reply_id = ? order by bc.bc_id DESC"

	res, errQuery := sqls.DB.Query(insq, bc.BlogCommentID)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	count := 0
	for res.Next() {
		count++
		child := &BlogComment{}
		user := &User{}
		errScan := res.Scan(&child.BlogCommentID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &child.BlogPostID, &child.ReplyID, &child.CommentContent, &child.CreatedAt, &child.Deleted)
		if errScan != nil {
			return nil, errScan
		}
		child.User = user

		bc, errRec := sqls.GetBlogChildrenComments(child)
		if errRec != nil {
			return nil, errRec
		}

		children = append(children, bc)
	}

	if count != 0 {
		bc.Children = children
	}

	return bc, nil
}

//BLOGPOST DB METHODS

//CreateBlogPost creates a blog post
func (sqls *SQLStore) CreateBlogPost(bp *BlogPost) (*BlogPost, error) {
	insq := "insert into blog_post(user_id, group_id, post_title, post_content, created_at) values(?,?,?,?,?)"

	res, errExec := sqls.DB.Exec(insq, bp.User.UserID, bp.GroupID, bp.PostTitle, bp.PostContent, time.Now())
	if errExec != nil {
		return nil, errExec
	}

	bpid, errID := res.LastInsertId()
	if errID != nil {
		return nil, errID
	}
	bp.BlogPostID = int(bpid)

	return bp, nil
}

//GetBlogPost gets a blog post
func (sqls *SQLStore) GetBlogPost(bpid int) (*BlogPost, error) {
	bp := &BlogPost{}
	user := &User{}

	insq := "select bp.bp_id, bp.user_id, u.first_name, u.last_name, u.photo_url, bp.group_id, bp.post_title, bp.post_content, bp.created_at from blog_post bp join user u on bp.user_id = u.user_id where bp.bp_id = ?"

	errQuery := sqls.DB.QueryRow(insq, bpid).Scan(&bp.BlogPostID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &bp.GroupID, &bp.PostTitle, &bp.PostContent, &bp.CreatedAt)
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errQuery
	}
	bp.User = user

	return bp, nil
}

//GetBlogPostsByGroup gets blog posts for a group
func (sqls *SQLStore) GetBlogPostsByGroup(gid int, page int) ([]*BlogPost, error) {
	bps := make([]*BlogPost, 0)
	offset := 3 * (page - 1)

	insq := "select bp.bp_id, bp.user_id, u.first_name, u.last_name, u.photo_url, bp.group_id, bp.post_title, bp.post_content, bp.created_at from blog_post bp join user u on bp.user_id = u.user_id where bp.group_id = ? order by bp.bp_id DESC limit 3 offset ?"

	res, errQuery := sqls.DB.Query(insq, gid, offset)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {
		bp := &BlogPost{}
		user := &User{}
		errScan := res.Scan(&bp.BlogPostID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &bp.GroupID, &bp.PostTitle, &bp.PostContent, &bp.CreatedAt)
		if errScan != nil {
			return nil, errScan
		}
		bp.User = user
		//GetComments method when done

		bps = append(bps, bp)
	}

	return bps, nil
}

//DeleteBlogPost deletes a specified blog post
func (sqls *SQLStore) DeleteBlogPost(bpid int) error {

	insq := "delete from blog_comment where bp_id = ?"

	_, errExec := sqls.DB.Exec(insq, bpid)
	if errExec != nil {
		return errExec
	}

	insq = "delete from blog_post where bp_id = ?"

	_, errExec = sqls.DB.Exec(insq, bpid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//MEMEBERSHIPREQUEST DB METHODS

//CreateMembershipRequest creates a new membership request
func (sqls *SQLStore) CreateMembershipRequest(gpid int, userid int) (*MembershipRequest, error) {
	mr := &MembershipRequest{}
	insq := "insert into membership(user_id, group_id, updated_at, state) values(?,?,?,?)"

	res, errExec := sqls.DB.Exec(insq, userid, gpid, time.Now(), false)
	if errExec != nil {
		return nil, errExec
	}

	mid, errID := res.LastInsertId()
	if errID != nil {
		return nil, errID
	}
	mr.MembershipID = int(mid)
	user := &User{}
	user.UserID = userid
	mr.User = user
	mr.GroupID = gpid

	return mr, nil
}

//GetMembershipRequest gest a membership request for a given request
func (sqls *SQLStore) GetMembershipRequest(mrid int) (*MembershipRequest, error) {
	mr := &MembershipRequest{}
	user := &User{}

	insq := "select m.membership_id, m.user_id, u.first_name, u.last_name, u.photo_url, m.group_id, m.updated_at, m.state from membership m join user u on m.user_id = u.user_id where m.membership_id = ?"

	errQuery := sqls.DB.QueryRow(insq, mrid).Scan(&mr.MembershipID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &mr.GroupID, &mr.UpdatedAt, &mr.State)
	if errQuery != nil {
		if errQuery == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errQuery
	}
	mr.User = user

	return mr, nil
}

//GetMembershipRequests gets all membership requests for the given group
func (sqls *SQLStore) GetMembershipRequests(gpid int) ([]*MembershipRequest, error) {
	mrs := make([]*MembershipRequest, 0)
	insq := "select m.membership_id, m.user_id, u.first_name, u.last_name, u.photo_url, m.group_id, m.updated_at, m.state from membership m join user u on m.user_id = u.user_id where m.group_id = ? and state = false"

	res, errQuery := sqls.DB.Query(insq, gpid)
	if errQuery != nil {
		return nil, errQuery
	}
	defer res.Close()

	for res.Next() {
		mr := &MembershipRequest{}
		user := &User{}
		errScan := res.Scan(&mr.MembershipID, &user.UserID, &user.FirstName, &user.LastName, &user.PhotoURL, &mr.GroupID, &mr.UpdatedAt, &mr.State)
		if errScan != nil {
			return nil, errScan
		}
		mr.User = user
		mrs = append(mrs, mr)
	}

	return mrs, nil
}

//AcceptMembershipRequest accepts a request for a user
func (sqls *SQLStore) AcceptMembershipRequest(mrid int) error {
	insq := "update membership set state = true, updated_at = ? where membership_id = ?"

	_, errExec := sqls.DB.Exec(insq, time.Now(), mrid)
	if errExec != nil {
		return errExec
	}

	return nil
}

//DeclineMembershipRequest declines a request for a user
func (sqls *SQLStore) DeclineMembershipRequest(mrid int) error {

	insq := "delete from membership where membership_id = ?"

	_, errExec := sqls.DB.Exec(insq, mrid)
	if errExec != nil {
		return errExec
	}
	return nil
}

//LeaveGroup allows a user to leave a group
func (sqls *SQLStore) LeaveGroup(gid int, userid int) error {

	insq := "delete from membership where group_id = ? and user_id = ?"

	_, errExec := sqls.DB.Exec(insq, gid, userid)
	if errExec != nil {
		return errExec
	}
	return nil
}
