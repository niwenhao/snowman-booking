/*
 * スキー予約サービス
 *
 * このサービスはスキー場一覧、スキー場のホテルの一覧、スキー場のホテルに予約することできます。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type Booking struct {

	// お客様問い合わせ用注文番号
	OrderId string `json:"orderId"`

	// 予約するホテルの内部ID
	HotelId int32 `json:"hotelId,omitempty"`

	BookingGuest Guest `json:"bookingGuest"`

	// 同行者情報
	AdditionalGuests []BookingAdditionalGuestsInner `json:"additionalGuests,omitempty"`
}

// AssertBookingRequired checks if the required fields are not zero-ed
func AssertBookingRequired(obj Booking) error {
	elements := map[string]interface{}{
		"orderId": obj.OrderId,
		"bookingGuest": obj.BookingGuest,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGuestRequired(obj.BookingGuest); err != nil {
		return err
	}
	for _, el := range obj.AdditionalGuests {
		if err := AssertBookingAdditionalGuestsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertBookingConstraints checks if the values respects the defined constraints
func AssertBookingConstraints(obj Booking) error {
	return nil
}
