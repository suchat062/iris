package usecase

import (
	"iris/src/modules/profile/model"
	"iris/src/modules/profile/repository"
)

type profileUsecaseImpl struct {
	profileRepository repository.ProfileRepository
}

func NewProfileUsecase(profileRepository repository.ProfileRepository) *profileUsecaseImpl {
	return &profileUsecaseImpl{profileRepository:profileRepository}
}

func (profileUsecase *profileUsecaseImpl) SaveProfile(profile *model.Profile) (*model.Profile, error) {
	err := profileUsecase.profileRepository.Save(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (profileUsecase *profileUsecaseImpl) UpdateProfile(id string, profile *model.Profile) (*model.Profile, error) {
	err := profileUsecase.profileRepository.Update(id, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (profileUsecase *profileUsecaseImpl) GetByID(id string) (*model.Profile, error) {
	profile ,err := profileUsecase.profileRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (profileUsecase *profileUsecaseImpl) GetByEmail(email string) (*model.Profile, error) {
	profile, err := profileUsecase.profileRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return profile, nil
}