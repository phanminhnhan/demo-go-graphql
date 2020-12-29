package database


import (
	"context"
	"graphql_postrgres/graph/model"
)

type LinkRepo struct{
	sql *Sql
}

func (l * LinkRepo) GetAllLinks(ctx context.Context)([]*model.Link, error){
	var links []*model.Link
	err := l.sql.Db.SelectContext(ctx, &links, "select *from links")
	if err != nil{
		return links, nil
	}
	return links, nil
}