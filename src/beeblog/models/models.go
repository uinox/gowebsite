package models

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Tm struct {
	Name string
	Op   string
	Id   string
}

type TopicView struct {
	Topic   *Topic
	Replies []*Comment
}

type HomeData struct {
	Topics     []*Topic
	Categories []*Category
}

const (
	_DB_NAME   = "mysql"
	_DB_DRIVER = "root:Zhang1991#@tcp(127.0.0.1:3306)/bee?charset=utf8&parseTime=true&loc=Local"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"auto_now_add;type(datetime)"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"auto_now_add;type(datetime)"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id             int64
	Uid            int64
	Title          string
	Category       string
	Labels         string
	Content        string `orm:"size(5000)"`
	Attachment     string
	Created        time.Time `orm:"auto_now_add;type(datetime)"`
	Updated        time.Time `orm:"auto_now_add;type(datetime)"`
	Views          int64     `orm:"index"`
	Author         string
	ReplyTime      time.Time `orm:"auto_now_add;type(datetime)"`
	ReplyCount     int64
	ReplyLastUserd int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

type User struct {
	Id   int64
	Name string `orm:"size(100)"`
	Password string
	Email    string
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

type Student struct {
	Id   int64
	Name string
	Age  int64
	Addr string
}

func RegisterDB() {
	// orm.DefaultTimeLoc = time.UTC
	orm.RegisterDataBase("default", _DB_NAME, _DB_DRIVER, 30)

	orm.RegisterModel(new(User), new(Student), new(Category), new(Topic), new(Comment))

	orm.RunSyncdb("default", false, true)

}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("category")
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

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err

}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func AddTopic(title, category, label, content, attachment string) error {

	//处理标签
	label = "$" + strings.Join(strings.Split(label, " "), "$#") + "#"

	o := orm.NewOrm()

	topic := &Topic{
		Title:      title,
		Category:   category,
		Labels:     label,
		Content:    content,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	//!+category number

	cate := new(Category)

	topics, err := GetAllTopics(category, "", true)
	if err != nil {
		return err
	} else {
		cate.TopicCount = int64(len(topics))
		_, err = o.Update(cate)
	}

	// cate := new(Category)

	/* qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	} */

	//!-category number end

	return err
}

func GetAllTopics(category, label string, isDesc bool) (topics []*Topic, err error) {
	o := orm.NewOrm()

	topics = make([]*Topic, 0)

	qs := o.QueryTable("topic")

	if isDesc {
		if len(category) > 0 {
			qs = qs.Filter("category", category)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#")
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

	topic.Labels = strings.Replace(strings.Replace(topic.Labels, "#", " ", -1), "$", " ", -1)
	return topic, nil
}

func ModifyTopic(tid, title, category, label, content, attachment string) error {

	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	//处理标签
	label = "$" + strings.Join(strings.Split(label, " "), "$#") + "#"

	var oldCate, oldAttach string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldAttach = topic.Attachment
		topic.Title = title
		topic.Category = category
		topic.Labels = label
		topic.Content = content
		topic.Attachment = attachment
		topic.Updated = time.Now()
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}

	//update cate
	if len(oldCate) > 0 {
		cate := new(Category)
		/* qs := o.QueryTable("category")
		err := qs.Filter("title", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		} */
		topics, err := GetAllTopics(oldCate, label, true)
		if err != nil {
			return err
		} else {
			cate.TopicCount = int64(len(topics))
			_, err = o.Update(cate)
		}
	}

	//删除旧的附件
	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}

	cate := new(Category)
	/* qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	} */
	topics, err := GetAllTopics(category, "", true)
	if err != nil {
		return err
	} else {
		cate.TopicCount = int64(len(topics))
		_, err = o.Update(cate)
	}

	return nil
}

func DeleteTopic(tid, category string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	var oldCate string
	o := orm.NewOrm()

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}

	if len(oldCate) > 0 {
		cate := new(Category)
		// qs := o.QueryTable("category")
		// err = qs.Filter("title", oldCate).One(cate)
		// if err == nil {
		// 	cate.TopicCount--
		// 	_, err = o.Update(cate)

		// }

		topics, err := GetAllTopics(oldCate, "", true)
		if err != nil {
			return err
		} else {
			cate.TopicCount = int64(len(topics))
			_, err = o.Update(cate)
		}
	}

	//!+category number

	/* categoryT := new(Category)

	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(categoryT)
	if err != nil {
		return err
	}

	categoryT.TopicCount++
	_, err = o.Update(categoryT) */

	//!-category number end

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
	if err != nil {
		return err
	}

	//update replycount

	/* topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return err
	}
	topic.ReplyTime = time.Now()
	topic.ReplyCount++
	_, err = o.Update(topic) */

	//the way 2 for update replycount
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyCount++
		_, err = o.Update(topic)
	}

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

	var tidNum int64
	reply := &Comment{Id: ridNum}
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}

	}

	replies := make([]*Comment, 0)

	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {

		//!+ 自己处理的，但有问题，不知道具体错误情况
		topic := &Topic{Id: tidNum}
		if o.Read(topic) == nil {
			topic.ReplyTime = time.Now()
			topic.ReplyCount = 0
			_, err = o.Update(topic)
		}
		//!- end
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = replies[0].Created
		topic.ReplyCount = int64(len(replies))
		_, err = o.Update(topic)
	}

	return err
}

//自己写的
/* func DeleteReply(tid, rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	reply := &Comment{Id: ridNum}
	_, err = o.Delete(reply)

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return err
	}

	topic.ReplyTime = time.Now()
	topic.ReplyCount--
	_, err = o.Update(topic)

	return err
} */

func AddUser(name, password, email string) int {
	o := orm.NewOrm()

	user := &User{
		Name:     name,
		Password: password,
		Email:    email,
		Created:  time.Now(),
	}
	qs := o.QueryTable("user")
	err := qs.Filter("name", user.Name).One(user)
	fmt.Println(err)
	if err != nil {
		_, err = o.Insert(user)
		if err != nil {
			return 304
		}
		return 200
	}

	return 301
}

func GetUser(name string) (*User, error) {
	o := orm.NewOrm()

	user := new(User)

	qs := o.QueryTable("user")
	err := qs.Filter("name", name).One(user)
	if err != nil {
		fmt.Println("aaa")
		return nil, err
	}

	return user, err
}
