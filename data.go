package common

import "errors"

type CommonData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	TenantID string `json:"tenantId"`
}

// Implements Entity interface
func (c *CommonData) GetID() string         { return c.ID }
func (c *CommonData) SetID(id string)       { c.ID = id }
func (c *CommonData) GetName() string       { return c.Name }
func (c *CommonData) SetName(name string)   { c.Name = name }
func (c *CommonData) GetTenantID() string   { return c.TenantID }
func (c *CommonData) SetTenantID(id string) { c.TenantID = id }

// Implements Validator interface
func (c *CommonData) Validate() error {
	if c.ID == "" {
		return errors.New("id is required")
	}
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.TenantID == "" {
		return errors.New("tenantId is required")
	}
	return nil
}
