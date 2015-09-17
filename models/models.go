// Package models provides ...
package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/blog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"imdex"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Category        string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name,
		Created:   time.Now(),
		TopicTime: time.Now()}
	qs := o.QueryTable("Category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func AddTopic(title, category, content string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:     title,
		Category:  category,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(),
	}

	_, err := o.Insert(topic)
	IncreaseCategory(topic.Id)
	return err
}

func GetAllTopics(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")

	var err error

	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("Category", cate)
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)

	}
	return topics, err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(tid, title, category, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	DecreaseCategory(tidNum)
	o := orm.NewOrm()
	Topic := &Topic{Id: tidNum}
	if o.Read(Topic) == nil {
		Topic.Title = title
		Topic.Category = category
		Topic.Content = content
		Topic.Updated = time.Now()
		o.Update(Topic)
	}
	IncreaseCategory(tidNum)
	return nil
}

func DeleteTopic(tid string) error {
	cid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	DecreaseCategory(cid)
	o := orm.NewOrm()
	topic := &Topic{Id: cid}
	_, err = o.Delete(topic)
	return err
}

func AddReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	reply := &Comment{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}

	o := orm.NewOrm()
	_, err = o.Insert(reply)
	return err
}

func GetAllReplies(tid string) (replies []*Comment, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	replies = make([]*Comment, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}

func DeleteReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	reply := &Comment{Id: ridNum}

	_, err = o.Delete(reply)
	return err
}

func DecreaseCategory(tid int64) error {
	topic := &Topic{Id: tid}
	o := orm.NewOrm()
	err := o.Read(topic)
	if err != nil {
		return err
	}

	cate := &Category{Title: topic.Category}

	if o.Read(cate, "Title") != nil {
		return err
	}

	if cate.TopicCount == 1 {
		_, err = o.Delete(cate)
		if err != nil {
			return err
		}
	} else {
		cate.TopicCount--
	}
	_, err = o.Update(cate)
	if err != nil {
		return err
	}
	return nil
}

func IncreaseCategory(tid int64) error {

	o := orm.NewOrm()

	topic := &Topic{Id: tid}
	err := o.Read(topic)
	if err != nil {
		return err
	}

	cate := &Category{Title: topic.Category}

	err = o.Read(cate, "Title")
	switch err {
	case orm.ErrNoRows:
		cate.Title = topic.Category
		cate.Created = time.Now()
		cate.TopicCount = 1
		cate.TopicTime = time.Now()
		_, err = o.Insert(cate)
		if err != nil {
			return err
		}
	case nil:
		cate.TopicCount++
		_, err = o.Update(cate)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteTopics(cid string) error {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cidNum}

	err = o.Read(cate)
	if err != nil {
		return err
	}

	topics := make([]Topic, 0)
	qs := o.QueryTable("topic")
	_, err = qs.Filter("Category", cate.Title).All(&topics)
	if err != nil {
		return err
	}
	for _, topic := range topics {
		_, err = o.Delete(&topic)
		if err != nil {
			return err
		}
	}
	return nil
}
