package res

type ResourceFlag uint64

const (
	// basic resources
	DetailChange ResourceFlag = 1 << iota
	CommentPost
	CommentRevoke
	ArticlePost
	ArticleRevise
	ArticleRevoke

	basic_end
)

const (
	// user manager resources
	UserCreate = basic_end << iota
	UserDelete
	UserUpdate
	UserRetrive

	user_end
)

const (
	// role manager resources
	RoleCreate = user_end << iota
	RoleDelete
	RoleUpdate
	RoleRetrive

	role_end
)

const (
	// comment manager resources
	CommentCreate = role_end << iota
	CommentDelete
	CommentUpdate
	CommentRetrive

	comment_end
)

const (
	// article manager resources
	ArticleCreate = comment_end << iota
	ArticleDelete
	ArticleUpdate
	ArticleRetrive

	article_end
)

const (
	BasicResources = basic_end - 1

	UserResources    = user_end - basic_end
	RoleResources    = role_end - user_end
	CommentResources = comment_end - role_end
	ArticleResources = article_end - comment_end
)
