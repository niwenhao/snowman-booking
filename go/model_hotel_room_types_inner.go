/*
 * スキー予約サービス
 *
 * このサービスはスキー場一覧、スキー場のホテルの一覧、スキー場のホテルに予約することできます。
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type HotelRoomTypesInner struct {

	// 部屋種類名称
	Name string `json:"name,omitempty"`

	// 部屋タイプの詳細
	Detail string `json:"detail,omitempty"`
}

// AssertHotelRoomTypesInnerRequired checks if the required fields are not zero-ed
func AssertHotelRoomTypesInnerRequired(obj HotelRoomTypesInner) error {
	return nil
}

// AssertHotelRoomTypesInnerConstraints checks if the values respects the defined constraints
func AssertHotelRoomTypesInnerConstraints(obj HotelRoomTypesInner) error {
	return nil
}
