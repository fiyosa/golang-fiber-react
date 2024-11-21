package resource_user

type Show struct {
	Id          string   `json:"id" example:""`
	Username    string   `json:"username" example:""`
	Name        string   `json:"name" example:""`
	Roles       []string `json:"roles" example:""`
	Permissions []string `json:"permissions" example:""`
	CreatedAt   string   `json:"created_at" example:""`
	UpdatedAt   string   `json:"updated_at" example:""`
}
