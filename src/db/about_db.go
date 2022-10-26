package db

import "github.com/younny/slobbo-backend/src/types"

func (c *Client) GetAbout() *types.About {
	about := &types.About{}
	if err := c.Client.First(&about).Error; err != nil {
		return nil
	}

	return about
}

func (c *Client) UpdateAbout(about *types.About) error {
	if err := c.Client.Create(&about).Error; err != nil {
		return c.Client.Save(&about).Error
	}

	return c.Client.Save(&about).Error
}
