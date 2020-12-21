package groupssrc

import (
	"database/sql"
	"time"
)

//User is the UserModel, stores information about a user
//No user model methods in groups, however there is one in gateway.
type User struct {
	UserID    int    `json:"id"`
	GoogleID  string `json:"googleId,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhotoURL  string `json:"photoUrl"`
}

//Category is the CategoryModel, stores information about a category
type Category struct {
	CategoryID   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
}

//SavedCategory is a saved category, saved either from the inital survey or in-site in another fasion
type SavedCategory struct {
	SavedCategoryID int       `json:"savedCategoryId"`
	Category        *Category `json:"category"`
	User            *User     `json:"user"`
}

//ReturnCategories is the return struct type for the getCategories handler
type ReturnCategories struct {
	SavedCategories []*Category `json:"savedCategories"` //normal category because we don't need to send user info to front-end.
	Categories      []*Category `json:"categories"`
}

//Group is the GroupModel, stores information about a group
type Group struct {
	GroupID          int             `json:"groupId"`
	User             *User           `json:"user"`
	Category         *Category       `json:"category"`
	GroupName        string          `json:"groupName"`
	GroupDescription string          `json:"groupDescription"`
	CreatedAt        time.Time       `json:"createdAt"`
	IsSaved          bool            `json:"isSaved,omitempty"`
	JoinedState      string          `json:"joinedStatus,omitempty"`
	IsAdmin          bool            `json:"isAdmin,omitempty"`
	NumMembers       int             `json:"numMembers"`
	Comments         []*GroupComment `json:"comments,omitempty"`
}

//GroupComment is the CommentModel, stores information about a comment. Represents comments on group pages
type GroupComment struct {
	GroupCommentID int             `json:"groupCommentId"`
	User           *User           `json:"user"`
	GroupID        int             `json:"groupId"`
	ReplyID        sql.NullInt64   `json:"replyId"` //commentID of comment being replied to, if any
	CommentContent string          `json:"commentContent"`
	CreatedAt      time.Time       `json:"createdAt"`
	Deleted        bool            `json:"deleted"`
	Children       []*GroupComment `json:"children,omitempty"`
}

//BlogPost is the BlogModel, stores information about a blog post
type BlogPost struct {
	BlogPostID  int       `json:"blogPostId"`
	User        *User     `json:"user"`
	GroupID     int       `json:"groupId"`
	PostTitle   string    `json:"postTitle"`
	PostContent string    `json:"postContent"`
	CreatedAt   time.Time `json:"createdAt"`
}

//BlogComment is the CommentModel, stores information about a comment. Represents comments on blog posts
type BlogComment struct {
	BlogCommentID  int            `json:"blogCommentId"`
	User           *User          `json:"user"`
	BlogPostID     int            `json:"blogPostId"`
	ReplyID        sql.NullInt64  `json:"replyId"`
	CommentContent string         `json:"commentContent"`
	CreatedAt      time.Time      `json:"createdAt"`
	Deleted        bool           `json:"deleted"`
	Children       []*BlogComment `json:"children,omitempty"`
}

//MembershipRequest is the MembershipRequestModel, stores information about the status of a membership requets, as well as saved groups
type MembershipRequest struct {
	MembershipID int       `json:"membershipID"`
	User         *User     `json:"user"`
	GroupID      int       `json:"groupId"`
	UpdatedAt    time.Time `json:"updatedAt"`
	State        bool      `json:"state"` //false = pending, true = accepted, missing = declined/no request
}

//ReturnAdmin contains admin groups info (admin, joined, saved) to be returned
type ReturnAdmin struct {
	AdminGroups  []*Group `json:"adminGroups"`
	JoinedGroups []*Group `json:"joinedGroups"`
	SavedGroups  []*Group `json:"savedGroups"`
}

//GroupContext contains the db context
type GroupContext struct {
	GStore GroupsStore
}
