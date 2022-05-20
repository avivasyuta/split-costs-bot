package domain

type Trip struct {
	ID   string
	Name string
}

type TripInput struct {
	Name *string
}
