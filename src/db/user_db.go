package db

import (
	"github.com/younny/slobbo-backend/src/types"
)

func (c *Client) GetUsers() *types.UserList {
	users := &types.UserList{}

	if err := c.Client.Find(&users.Items).Error; err != nil {
		return nil
	}

	return users
}

func (c *Client) GetUserByID(id uint) *types.User {
	user := &types.User{}

	if err := c.Client.Where("id = ?", id).Take(&user).Error; err != nil {
		return nil
	}

	return user
}

func (c *Client) GetUserByEmail(email string) *types.User {
	user := &types.User{}

	if err := c.Client.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil
	}

	return user
}

func (c *Client) CreateUser(user *types.User) error {
	return c.Client.Create(&user).Error
}

func (c *Client) UpdateUser(user *types.User) error {
	return c.Client.Save(&user).Error
}

func (c *Client) DeleteUser(id uint) error {
	user := &types.User{}

	if err := c.Client.Where("id = ?", id).Take(&user).Error; err != nil {
		return nil
	}

	return c.Client.Delete(&user).Error
}
