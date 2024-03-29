package user_input_add_v1_request

type UserAddRequest struct {
	UserId    string `json:"userId" validate:"required"`
	Rut       string `json:"rut" validate:"required"`
	UserName  string `json:"userName" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Valid     bool   `json:"valid" validate:"required"`
	Profile   Profile
}

type Profile struct {
	ProfileId       int    `json:"profileId" validate:"required"`
	ProfileStatus   bool   `json:"profileStatus" validate:"required"`
	ProfileDateInit string `json:"profileDateInit" validate:"required"`
	ProfileDateEnd  string `json:"profileDateEnd" validate:"required"`
	ProfileAllTime  bool   `json:"profileAllTime" validate:"required"`
}
