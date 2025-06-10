package common

import "errors"

type CommonData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Implements Entity interface
func (c *CommonData) GetID() string       { return c.ID }
func (c *CommonData) SetID(id string)     { c.ID = id }
func (c *CommonData) GetName() string     { return c.Name }
func (c *CommonData) SetName(name string) { c.Name = name }

// Implements Validator interface
func (c *CommonData) Validate() error {
	if c.ID == "" {
		return errors.New("id is required")
	}
	if c.Name == "" {
		return errors.New("name is required")
	}
	return nil
}
