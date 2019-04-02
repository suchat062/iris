package repository

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"iris/src/modules/profile/model"
	"time"
)

type profileRepositoryMongo struct {
	db *mgo.Database
	collection string
}

func NewProfileRepositoryMongo(db *mgo.Database, collection string) *profileRepositoryMongo {
	return &profileRepositoryMongo{
		db: db,
		collection: collection,
	}
}

func (r *profileRepositoryMongo) Save(profile *model.Profile) error {
	err := r.db.C(r.collection).Insert(profile)
	return err
}

func (r *profileRepositoryMongo) Update(id string, profile *model.Profile) error {
	profile.UpdatedAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"id" : id}, profile)
	return err
}

func (r *profileRepositoryMongo) Delete(id string) error {
	err := r.db.C(r.collection).Remove(bson.M{"id" : id})
	return err
}

func (r *profileRepositoryMongo) FindByID(id string) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.C(r.collection).FindId(bson.M{"id" : id}).One(&profile)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *profileRepositoryMongo) FindByEmail(email string) (*model.Profile, error) {
	var profile model.Profile
	err := r.db.C(r.collection).Find(bson.M{"email" : email}).One(&profile)

	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepositoryMongo) findAll() (*model.Profiles, error) {
	var profiles model.Profiles
	err := r.db.C(r.collection).Find(bson.M{}).All(&profiles)

	if err != nil {
		return nil, err
	}
	return &profiles, nil
}
