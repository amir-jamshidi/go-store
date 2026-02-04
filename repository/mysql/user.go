package mysql

import "room-reserve/entity"

func (db DB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	return false, nil
}
func (db DB) RegisterUser(u entity.User) (entity.User, error) {
	return entity.User{}, nil
}
