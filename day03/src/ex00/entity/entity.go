package entity

type Restaurant struct {
	ID        int
	Name      string
	Address   string
	Phone     string
	Longitude float64
	Latitude  float64
}

func New(id int, name, addr, phone string, lon, lat float64) Restaurant {
	return Restaurant{
		ID:        id,
		Name:      name,
		Address:   addr,
		Phone:     phone,
		Longitude: lon,
		Latitude:  lat,
	}
}
