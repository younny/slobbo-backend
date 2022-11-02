package db

import (
	"github.com/younny/slobbo-backend/src/types"
)

func (c *Client) CreateAbout(about *types.About) error {
	if err := c.Client.Create(&about).Error; err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAbout() *types.About {
	about := &types.About{}
	if err := c.Client.First(&about).Error; err != nil {
		return nil
	}
	return about
}

func (c *Client) UpdateAbout(about *types.About) error {
	return c.Client.Where("id = ?", about.ID).Take(&types.About{}).UpdateColumns(&about).Error
}
