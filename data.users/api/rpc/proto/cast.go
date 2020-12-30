package proto

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/timoth-y/scrapnote-api/data.users/core/model"
)

func (m *User) ToNative() *model.User {
	return &model.User{
		UniqueID:     m.UniqueID,
		Username:     m.Username,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		PhoneNumber:  m.PhoneNumber,
		Avatar:       m.Avatar,
		Confirmed:    m.Confirmed,
		RegisterDate: m.RegisterDate.AsTime(),
	}
}

func (m *User) FromNative(n *model.User) *User {
	m.UniqueID = n.UniqueID
	m.Username = n.Username
	m.Email = n.Email
	m.PasswordHash = n.PasswordHash
	m.FirstName = n.FirstName
	m.LastName = n.LastName
	m.PhoneNumber = n.PhoneNumber
	m.Avatar = n.Avatar
	m.Confirmed = n.Confirmed
	m.RegisterDate = timestamppb.New(n.RegisterDate)
	return m
}

func NativeToUsers(native []*model.User) []*User {
	users := make([]*User, 0)
	for _, user := range native {
		if user == nil {
			continue
		}
		users = append(users, (&User{}).FromNative(user))
	}
	return users
}

func UsersToNative(in []*User) []*model.User {
	users := make([]*model.User, 0)
	for _, user := range in {
		users = append(users, user.ToNative())
	}
	return users
}