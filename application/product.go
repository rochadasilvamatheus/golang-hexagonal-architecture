package application

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
	ID string
	Name string
	Price float64
	Status string
}

func (p *Product) isValid() (bool, error) {
	if p.Price <= 0 {
		return false, errors.New("Price must be greater than 0")
	}

	if p.Name == "" {
		return false, errors.New("Name is required")
	}

	return true, nil
}

func (p *Product) Enable() error {
	p.Status = ENABLED
	return nil
}

func (p *Product) Disable() error {
	p.Status = DISABLED
	return nil
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