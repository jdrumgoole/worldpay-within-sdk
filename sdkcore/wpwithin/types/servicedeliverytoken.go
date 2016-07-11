package types
import "time"

type ServiceDeliveryToken struct {

	Key string `json:"key"`
	Issued time.Time `json:"issued"`
	Expiry time.Time `json:"expiry"`
	RefundOnExpiry bool `json:"refundOnExpiry"`
	Signature []byte `json:"signature"`
}