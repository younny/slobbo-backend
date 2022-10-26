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

	GetAbout() *types.About
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
	c.autoMigrate()
	return nil
}

func (c *Client) autoMigrate() {
	c.Client.AutoMigrate(&types.About{})
	c.Client.AutoMigrate(&types.Post{})
	c.Client.AutoMigrate(&types.Workshop{})
}
