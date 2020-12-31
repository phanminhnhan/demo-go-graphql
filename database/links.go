package database

import (
	"github.com/go-pg/pg/v9"
	"graphql_postrgres/graph/model"
)

type LinkRepo struct{
	Db *pg.DB
}

func (l * LinkRepo) GetAllLinks()([]*model.Link, error){
	var links  []*model.Link
	err := l.Db.Model(&links).Select()
	if err != nil {
		return nil, err
	}
	return links, nil
}

func(l *LinkRepo) CreateLink(link *model.Link)(*model.Link, error){
	_, err := l.Db.Model(link).Returning("*").Insert()
	if err != nil{
		return nil, err
	}
	return link, nil
}

func(l *LinkRepo) GetbyId(id string)( *model.Link, error){
	var link model.Link
	err:= l.Db.Model(&link).Where("id = ?", id).First()
	if err != nil{
		return nil, err
	}
	return &link, nil
}

func(l *LinkRepo)Updatelink(link *model.Link)(*model.Link, error){
	_, err := l.Db.Model(link).Where("id = ?", link.ID).Update()
	return link , err
}

func  (l *LinkRepo)DeleteLink(link *model.Link) error{
	_, err := l.Db.Model(link).Where("id= ?", link.ID).Delete()
	return err
}