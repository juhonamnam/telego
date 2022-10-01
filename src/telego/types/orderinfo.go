package types

type OrderInfo struct {
	Name            *string         `json:"name"`
	PhoneNumber     *string         `json:"phone_number"`
	Email           *string         `json:"email"`
	ShippingAddress *ShippingAddres `json:"shipping_address"`
}
