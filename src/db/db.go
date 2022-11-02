package db

import (
	"github.com/younny/slobbo-backend/src/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	pageSize = 10
)

type ClientInterface interface {
	Ping() error
	Connect(connectionString string) error
	DropTable(arg0 ...interface{}) error
	AutoMigrate(arg0 ...interface{}) error

	GetUsers() *types.UserList
	GetUserByID(id uint) *types.User
	GetUserByEmail(email string) *types.User
	CreateUser(user *types.User) error
	UpdateUser(user *types.User) error
	DeleteUser(id uint) error

	GetAbout() *types.About
	CreateAbout(about *types.About) error
	UpdateAbout(about *types.About) error

	GetPosts(pageID int) *types.PostList
	GetPostByID(id uint) *types.Post
	CreatePost(post *types.Post) error
	UpdatePost(post *types.Post) error
	DeletePost(id uint) error

	GetWorkshops(pageID int) *types.WorkshopList
	GetWorkshopByID(id uint) *types.Workshop
	CreateWorkshop(workshop *types.Workshop) error
	UpdateWorkshop(workshop *types.Workshop) error
	DeleteWorkshop(id uint) error
}

type Client struct {
	Client *gorm.DB
}

func (c *Client) Ping() error {
	return c.Client.DB().Ping()
}

func (c *Client) Connect(connectionString string) error {
	var err error

	c.Client, err = gorm.Open("postgres", connectionString)

	if err != nil {
		return err
	}
	c.Client.LogMode(false)
	c.AutoMigrate(&types.User{}, &types.About{}, &types.Post{}, &types.Workshop{})
	return nil
}

func (c *Client) DropTable(arg0 ...interface{}) error {
	return c.Client.DropTableIfExists(arg0...).Error
}

func (c *Client) AutoMigrate(arg0 ...interface{}) error {
	return c.Client.AutoMigrate(arg0...).Error
}
