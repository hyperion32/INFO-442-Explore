package groupssrc

//GroupsStore represents all the database methods, each model's methods are grouped. UserId will be part of all structs in some way, so it is not needed as an input paramter unless there is no other input struct or input struct id.
type GroupsStore interface {
	//Category model methods
	GetCategories(query string) ([]*Category, error)
	SaveCategory(catid int, userid int) error
	UnsaveCategory(catid int, userid int) error
	GetSavedCategories(userid int) ([]*Category, error)

	//Group model methods
	CreateGroup(gp *Group) (*Group, error)
	SearchGroups(query string, userid int, page int, catid int) ([]*Group, error)
	GetGroup(gpid int, userid int) (*Group, error)
	DeleteGroup(gpid int) error
	SaveGroup(gpid int, userid int) error
	UnsaveGroup(gp int, userid int) error
	GetSavedGroups(userid int) ([]*Group, error)
	GetJoinedGroups(userid int) ([]*Group, error)
	GetAdminGroups(userid int) ([]*Group, error)
	GetGroupMembershipCount(gid int) (int, error)

	//GroupComment model methods
	CreateGroupComment(gc *GroupComment) (*GroupComment, error)
	GetGroupComment(gcid int) (*GroupComment, error)
	DeleteGroupComment(gcid int) error
	GetGroupCommentsByGroup(gid int, page int) ([]*GroupComment, error)
	GetChildrenComments(gc *GroupComment) (*GroupComment, error)

	//BlogComment model methods
	CreateBlogComment(bc *BlogComment) (*BlogComment, error)
	GetBlogComment(bcid int) (*BlogComment, error)
	DeleteBlogComment(bcid int) error
	GetBlogCommentsByBlog(bid int, page int) ([]*BlogComment, error)
	GetBlogChildrenComments(bc *BlogComment) (*BlogComment, error)

	//BlogPost model methods
	CreateBlogPost(bp *BlogPost) (*BlogPost, error)
	GetBlogPost(bpid int) (*BlogPost, error)
	GetBlogPostsByGroup(gid int, page int) ([]*BlogPost, error)
	DeleteBlogPost(bpid int) error

	//MembershipRequest model methods
	CreateMembershipRequest(gpid int, userid int) (*MembershipRequest, error)
	GetMembershipRequest(mrid int) (*MembershipRequest, error)
	GetMembershipRequests(gpid int) ([]*MembershipRequest, error)
	AcceptMembershipRequest(mrid int) error
	DeclineMembershipRequest(mrid int) error
	LeaveGroup(gid int, userid int) error
}
