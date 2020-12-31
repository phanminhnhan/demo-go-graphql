package database

import (
	"github.com/go-pg/pg/v9"
	"graphql_postrgres/graph/model"
)


type UserRepo struct{
	Db *pg.DB
}

func (u *UserRepo) GetUserById(id string ) (*model.User, error){
	var user model.User
	err := u.Db.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u * UserRepo) GetAllUsers()([]*model.User, error){
	var users  []*model.User
	_, err := u.Db.Query(&users,`select * from users inner join links on users.id = links.userid` )
	if err != nil {
		return nil, err
	}
	return users, nil
}


func(u *UserRepo) CreateUser(user *model.User)(*model.User, error){
	_, err := u.Db.Model(user).Returning("*").Insert()
	if err != nil{
		return nil, err
	}
	return user, nil
}