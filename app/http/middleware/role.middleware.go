package middleware

var Role roleMiddleware

type roleMiddleware struct{}

func checkRole(check string, roles []string) bool {
	for _, v := range roles {
		if v == check {
			return true
		}
	}
	return false
}

func (*roleMiddleware) IsAdmin(roles []string) bool {
	return checkRole("admin", roles)
}

func (*roleMiddleware) IsUser(roles []string) bool {
	return checkRole("user", roles)
}
