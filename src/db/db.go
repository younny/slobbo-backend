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
	GetPosts(pageID int) *types.PostList
	GetPostByID(id uint) *types.Post
	CreatePost(post *types.Post) error
	UpdatePost(post *types.Post) error
	DeletePost(id uint) error
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
	c.Client.AutoMigrate(&types.Post{})

	//todo more types..
}

func (c *Client) GetPosts(pageID int) *types.PostList {
	posts := &types.PostList{}
	c.Client.Where("id >= ?", pageID).Order("id").Limit(pageSize + 1).Find(&posts.Items)
	if len(posts.Items) == pageSize+1 {
		posts.NextPageID = posts.Items[len(posts.Items)-1].ID
		posts.Items = posts.Items[:pageSize]
	}
	return posts
}

func (c *Client) GetPostByID(id uint) *types.Post {
	post := &types.Post{}

	if err := c.Client.Where("id = ?", id).First(&post).Scan(post).Error; err != nil {
		return nil
	}

	return post
}

func (c *Client) CreatePost(post *types.Post) error {

	return c.Client.Create(&post).Error
}

func (c *Client) UpdatePost(post *types.Post) error {
	return c.Client.Save(&post).Error
}

func (c *Client) DeletePost(id uint) error {
	post := &types.Post{}
	if err := c.Client.Where("id = ?", id).First(&post).Error; err != nil {
		return err
	}

	return c.Client.Delete(&post).Error
}
