package user_entity_interface

type UserEntityInterface interface {
	GetRut() string
	GetUserName() string
	GetPassword() string
	GetEmail() string
	GetFirstName() string
	GetLastName() string
	GetValid() bool
	GetProfile() ProfileEntityInterface

	SetRut(rut string) error
	SetUserName(name string) error
	SetPassword(password string) error
	SetEmail(email string) error
	SetFirstName(firstName string) error
	SetLastName(lastName string) error
	SetValid(valid bool) error
	SetProfile(profile ProfileEntityInterface) error
}

type ProfileEntityInterface interface {
	GetProfileId() int
	GetProfileStatus() bool
	GetProfileDateInit() string
	GetProfileDateEnd() string
	GetProfileAllTime() bool

	SetProfileId(profileId int) error
	SetProfileStatus(profileStatus bool) error
	SetProfileDateInit(profileDateInit string) error
	SetProfileDateEnd(profileDateEnd string) error
	SetProfileAllTime(profileAllTime bool) error
}
