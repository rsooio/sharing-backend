package auth

func IsRoot(uid uint) bool {
	return uid == 1
}
