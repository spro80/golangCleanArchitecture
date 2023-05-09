package user_entity

import (
	user_entities_interface "github.com/spro80/golangCleanArchitecture/app/domain/entity/user_entity/interfaces"
)

type userClient struct {
	user *User
}

type profileClient struct {
	profile *Profile
}

func NewUserEntity() user_entities_interface.UserEntityInterface {
	userClient := &userClient{
		user: &User{},
	}
	return userClient
}

func NewProfileEntity() user_entities_interface.ProfileEntityInterface {
	profileClient := &profileClient{
		profile: &Profile{},
	}
	return profileClient
}

func (u *userClient) GetUserId() string {
	return u.user.UserId
}

func (u *userClient) GetRut() string {
	return u.user.Rut
}

func (u *userClient) GetUserName() string {
	return u.user.UserName
}

func (u *userClient) GetPassword() string {
	return u.user.Password
}

func (u *userClient) GetEmail() string {
	return u.user.Email
}

func (u *userClient) GetFirstName() string {
	return u.user.FirstName
}

func (u *userClient) GetLastName() string {
	return u.user.LastName
}

func (u *userClient) GetValid() bool {
	return u.user.Valid
}

func (u *userClient) GetProfile() user_entities_interface.ProfileEntityInterface {
	return u.user.Profile
}

func (p *profileClient) GetProfileId() int {
	return p.profile.ProfileId
}

func (p *profileClient) GetProfileStatus() bool {
	return p.profile.ProfileStatus
}

func (p *profileClient) GetProfileDateInit() string {
	return p.profile.ProfileDateInit
}

func (p *profileClient) GetProfileDateEnd() string {
	return p.profile.ProfileDateEnd
}

func (p *profileClient) GetProfileAllTime() bool {
	return p.profile.ProfileAllTime
}

func (u *userClient) SetUserId(userId string) error {
	u.user.UserId = userId
	return nil
}

func (u *userClient) SetRut(rut string) error {
	u.user.Rut = rut
	return nil
}

func (u *userClient) SetUserName(name string) error {
	u.user.UserName = name
	return nil
}

func (u *userClient) SetPassword(password string) error {
	u.user.Password = password
	return nil
}

func (u *userClient) SetEmail(email string) error {
	u.user.Email = email
	return nil
}

func (u *userClient) SetFirstName(firstName string) error {
	u.user.FirstName = firstName
	return nil
}

func (u *userClient) SetLastName(lastName string) error {
	u.user.LastName = lastName
	return nil
}

func (u *userClient) SetValid(valid bool) error {
	u.user.Valid = valid
	return nil
}

func (u *userClient) SetProfile(profile user_entities_interface.ProfileEntityInterface) error {
	u.user.Profile = profile
	return nil
}

func (p *profileClient) SetProfileId(profileId int) error {
	p.profile.ProfileId = profileId
	return nil
}

func (p *profileClient) SetProfileStatus(profileStatus bool) error {
	p.profile.ProfileStatus = profileStatus
	return nil
}

func (p *profileClient) SetProfileDateInit(profileDateInit string) error {
	p.profile.ProfileDateInit = profileDateInit
	return nil
}

func (p *profileClient) SetProfileDateEnd(profileDateEnd string) error {
	p.profile.ProfileDateEnd = profileDateEnd
	return nil
}

func (p *profileClient) SetProfileAllTime(profileAllTime bool) error {
	p.profile.ProfileAllTime = profileAllTime
	return nil
}
