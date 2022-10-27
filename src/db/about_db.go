package db

import (
	"github.com/younny/slobbo-backend/src/types"
)

func (c *Client) createAbout(about *types.About) error {
	if err := c.Client.Create(&about).Error; err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAbout() *types.About {
	about := &types.About{}
	if err := c.Client.First(&about, 1).Error; err != nil {
		c.createAbout(about)
	}
	return about
}

func (c *Client) UpdateAbout(about *types.About) error {
	return c.Client.Save(&about).Error
}
