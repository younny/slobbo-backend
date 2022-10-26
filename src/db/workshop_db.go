package db

import (
	"github.com/younny/slobbo-backend/src/types"
)

func (c *Client) GetWorkshops(pageID int) *types.WorkshopList {
	workshops := &types.WorkshopList{}

	c.Client.Where("id >= ?", pageID).Order("id").Limit(pageSize + 1).Find(&workshops.Items)
	if len(workshops.Items) == pageSize+1 {
		workshops.NextPageID = workshops.Items[len(workshops.Items)-1].ID
		workshops.Items = workshops.Items[:pageSize]
	}
	return workshops
}

func (c *Client) GetWorkshopByID(id uint) *types.Workshop {
	workshop := &types.Workshop{}

	if err := c.Client.Where("id = ?", id).First(&workshop).Scan(&workshop).Error; err != nil {
		return nil
	}

	return workshop
}

func (c *Client) CreateWorkshop(workshop *types.Workshop) error {
	return c.Client.Create(&workshop).Error
}

func (c *Client) UpdateWorkshop(workshop *types.Workshop) error {
	return c.Client.Save(&workshop).Error
}

func (c *Client) DeleteWorkshop(id uint) error {
	workshop := &types.Workshop{}
	if err := c.Client.Where("id == ?", id).First(&workshop).Error; err != nil {
		return err
	}

	return c.Client.Delete(&workshop).Error
}
