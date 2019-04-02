package repository

import "iris/src/modules/profile/model"

type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindByEmail(string) (*model.Profile, error)
	findAll() (model.Profiles, error)
}
