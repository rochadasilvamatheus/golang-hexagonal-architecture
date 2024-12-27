package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type ProductInterface interface {
	isValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID string `valid:"uuidv4"`
	Name string `valid:"required"`
	Price float64 `valid:"float,optional"`
	Status string `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("Status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("Price must be greater or equal to 0")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("Price must be greater than zero to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
		return errors.New("Price must be zero to enable in order to have the product disabled")

}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}