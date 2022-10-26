package db

import "github.com/younny/slobbo-backend/src/types"

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
