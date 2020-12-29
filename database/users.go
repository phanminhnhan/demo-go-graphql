package database

import (
	"context"
	"graphql_postrgres/graph/model"
)





type UserRepo struct{
	sql *Sql
}

func (l * LinkRepo) GetUserById(ctx context.Context, id int)(*model.User, error){
	var user *model.User
	err := l.sql.Db.SelectContext(ctx, user, "select username from users where id = $1", id )
	if err != nil {
		return user, err
	}
	return user, nil
}