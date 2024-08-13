package utils

// Error is a struct of error obj
// code is for error code
// status code is the first 3 digits of the error code
type Error struct {
	Code             int                    `json:"code"`
	Message          string                 `json:"message"`
	LocalizedMessage Message                `json:"localized_message,omitempty"`
	Data             map[string]interface{} `json:"data,omitempty"`
}

type Message struct {
	English   string `json:"en"`
	Indonesia string `json:"id"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) CodeErr() int {
	return e.Code
}

func (e *Error) DataErr() map[string]interface{} {
	return e.Data
}

func (e *Error) WithData(data map[string]interface{}) *Error {
	return &Error{
		Code:             e.Code,
		Message:          e.Message,
		LocalizedMessage: e.LocalizedMessage,
		Data:             data,
	}
}

func NewLegacyErrorMessage(code int, message string, id, en string) *Error {
	return &Error{Code: code, Message: message, LocalizedMessage: Message{
		English:   en,
		Indonesia: id,
	}}
}

var (
	ErrPromotionSingleUseNotExists               = NewLegacyErrorMessage(40000, "Sorry! promotion single use not exists!", "Sorry! promotion single use not exists!", "Maaf! Promosi sekali pakai tidak ada!")
	ErrInternal                                  = NewLegacyErrorMessage(50000, "Sorry. Unable to complete your request", "Sorry. Unable to complete your request", "Maaf. Tidak dapat menyelesaikan permintaan Anda")
	ErrAuth                                      = NewLegacyErrorMessage(40006, "Failed to authenticate to Promotion Portal", "Failed to authenticate to Promotion Portal", "Gagal melakukan autentikasi ke Portal Promosi")
	ErrTokenExpired                              = NewLegacyErrorMessage(40101, "Sorry, token is expired", "Sorry, token is expired", "Maaf, token telah kedaluwarsa")
	ErrTokenInvalid                              = NewLegacyErrorMessage(40102, "Soory, token is invalid", "Soory, token is invalid", "Maaf, token tidak valid")
	ErrPromoCodeIsUsed                           = NewLegacyErrorMessage(40000, "promo code already to used!", "promo code already to used!", "promo code telah digunakan!")
	ErrBadRequest                                = NewLegacyErrorMessage(40000, "Oops! Something went wrong!", "Oops! Something went wrong!", "Ups! Ada yang salah!")
	ErrInvalidInput                              = NewLegacyErrorMessage(40001, "Oops! Something went wrong!", "Oops! Something went wrong!", "Ups! Ada yang salah!")
	ErrPromotionNotExists                        = NewLegacyErrorMessage(40002, "Sorry! promotion not exists!", "Sorry! promotion not exists!", "Maaf! Promosi tidak ada!")
	ErrPromotionAlreadyUsed                      = NewLegacyErrorMessage(40003, "Promotion is Already to used!", "Promotion is Already to used!", "Promosi sudah digunakan!")
	ErrPromotionVersion                          = NewLegacyErrorMessage(40004, "Sorry, you can't create promotion version", "Sorry, you can't create promotion version", "Maaf, Anda tidak dapat membuat versi promosi")
	ErrStatusPromotionEdit                       = NewLegacyErrorMessage(40005, "This promo code is not allowed to update because the status isn't `Terjadwal`", "This promo code is not allowed to update because the status isn't `Terjadwal`", "Kode promo ini tidak diizinkan untuk diperbarui karena statusnya bukan `Terjadwal`")
	ErrGetPromotionTimeRangeByPromoID            = NewLegacyErrorMessage(40006, "Failed to get promotion time ranges BY Promo ID", "Failed to get promotion time ranges BY Promo ID", "Gagal mendapatkan rentang waktu promosi BERDASARKAN ID Promo")
	ErrAccountInvalid                            = NewLegacyErrorMessage(40007, "Your account name is invalid", "Your account name is invalid", "Nama akun Anda tidak valid")
	ErrCodeRequired                              = NewLegacyErrorMessage(40008, "code is required", "code is required", "kode dibutuhkan")
	ErrIdentifierRequired                        = NewLegacyErrorMessage(40009, "identifier is required", "identifier is required", "identifier dibutuhkan")
	ErrServiceTypeRequired                       = NewLegacyErrorMessage(40010, "service_type is required", "service_type is required", "service_type dibutuhkan")
	ErrPaymentTypeRequired                       = NewLegacyErrorMessage(40011, "payment_type is required", "payment_type is required", "payment_type dibutuhkan")
	ErrPromotionIDInvalid                        = NewLegacyErrorMessage(40012, "promotion id not valid, using other id", "promotion id not valid, using other id", "ID promosi tidak valid, gunakan ID lain")
	ErrReactivatedPromoInvalidBecauseOverPeriode = NewLegacyErrorMessage(40013, "failed to reactivated promotion code, because promo code is expired", "failed to reactivated promotion code, because promo code is expired", "gagal mengaktifkan kembali kode promo, karena kode promo telah kedaluwarsa")
	ErrReactivatedPromoInvalidBecauseLimitBudget = NewLegacyErrorMessage(40014, "failed to reactivated promotion code, because promo code has used all the budget", "failed to reactivated promotion code, because promo code has used all the budget", "gagal mengaktifkan kembali kode promo, karena kode promo telah menggunakan semua anggaran")
	ErrPromoLimitBudget                          = NewLegacyErrorMessage(40015, "the budget for this promo has reached the limit", "the budget for this promo has reached the limit", "anggaran untuk promo ini telah mencapai batasnya")
	ErrGetPromotionRules                         = NewLegacyErrorMessage(40016, "Failed to get promotion rules", "Failed to get promotion rules", "Gagal mendapatkan aturan promosi")
	ErrInvalidStatus                             = NewLegacyErrorMessage(40017, "invalid promotion status", "invalid promotion status", "status promosi tidak valid")
	ErrInvalidPromotionID                        = NewLegacyErrorMessage(40018, "invalid promotion id or promotion not exists", "invalid promotion id or promotion not exists", "ID promosi tidak valid atau promosi tidak ada")
	ErrIdentifierMissing                         = NewLegacyErrorMessage(40019, "identifier is required", "identifier is required", "identifier dibutuhkan")
	ErrMaxReferralCodeUsage                      = NewLegacyErrorMessage(40020, "You have reached max usage of referal code", "You have reached max usage of referal code", "Anda telah mencapai batas penggunaan kode referal")
	ErrMaxPromotionCodeUsage                     = NewLegacyErrorMessage(40021, "Your promotion max usage has been reached", "Your promotion max usage has been reached", "Penggunaan maksimal promosi Anda telah tercapai")
	ErrPromotionExpired                          = NewLegacyErrorMessage(40022, "Your promotion has expired", "Your promotion has expired", "Promosi Anda telah kedaluwarsa")
	ErrPromotionServiceTypeInvalid               = NewLegacyErrorMessage(40023, "Invalid service type", "Invalid service type", "Jenis layanan tidak valid untuk promosi ini")
	ErrPromotionPaymentTypeInvalid               = NewLegacyErrorMessage(40024, "Invalid payment type", "Invalid payment type", "Jenis pembayaran tidak valid untuk promosi ini")
	ErrPromotionDeactivated                      = NewLegacyErrorMessage(40025, "Your promotion code is deactivated", "Your promotion code is deactivated", "Kode promo Anda dinonaktifkan")
	ErrPromotionInvalid                          = NewLegacyErrorMessage(40025, "Your promotion code is invalid", "Your promotion code is invalid", "Kode promo Anda tidak valid")
	ErrPromotionInvalidPickupTime                = NewLegacyErrorMessage(40026, "Your promotion code is invalid for current pickup time", "Your promotion code is invalid for current pickup time", "Kode promo Anda tidak valid untuk waktu pengambilan saat ini")
	ErrPromotionInvalidPickupLoc                 = NewLegacyErrorMessage(40027, "Your pickup location is invalid for this promotion", "Your pickup location is invalid for this promotion", "Lokasi pengambilan Anda tidak valid untuk promosi ini")
	ErrPromotionInvalidDropoffLoc                = NewLegacyErrorMessage(40028, "Your dropoff location is invalid for this promotion", "Your dropoff location is invalid for this promotion", "Lokasi penurunan Anda tidak valid untuk promosi ini")
	ErrMaxReferralUsage                          = NewLegacyErrorMessage(40029, "You have reached max usage of referral code", "You have reached max usage of referral code", "Anda telah mencapai batas penggunaan kode referral")
	ErrPromoCodeInvalid                          = NewLegacyErrorMessage(40030, "Your promotion code is invalid", "Your promotion code is invalid", "Kode promo Anda tidak valid")
	ErrPromoOverlaps                             = NewLegacyErrorMessage(40031, "Promotion is overlapping", "Promotion is overlapping", "Promosi tumpang tindih")
	ErrDiscountValueInvalid                      = NewLegacyErrorMessage(40032, "Invalid discount value", "Invalid discount value", "Nilai diskon tidak valid")
	ErrDiscountTypeInvalid                       = NewLegacyErrorMessage(40033, "Invalid discount type", "Invalid discount type", "Jenis diskon tidak valid")
	ErrEndDateInvalid                            = NewLegacyErrorMessage(40034, "End Date could not be before Start Date", "End Date could not be before Start Date", "Tanggal Akhir tidak boleh sebelum Tanggal Mulai")
	ErrServiceTypeInvalid                        = NewLegacyErrorMessage(40035, "Invalid service type", "Invalid service type", "Jenis layanan tidak valid")
	ErrPaymentTypeInvalid                        = NewLegacyErrorMessage(40036, "Invalid payment type", "Invalid payment type", "Jenis pembayaran tidak valid")
	ErrPromoCodeMissing                          = NewLegacyErrorMessage(40037, "promotion_code is required", "promotion_code is required", "kode promo dibutuhkan")
	ErrIdentifierInvalid                         = NewLegacyErrorMessage(40038, "Kamu tidak bisa menggunakan kode promo ini", "You can't use this promo code", "Kamu tidak bisa menggunakan kode promo ini")
	ErrServiceUnavailable                        = NewLegacyErrorMessage(40039, "Service Unavailable", "Service Unavailable", "Layanan tidak tersedia saat ini, silakan coba lagi nanti")
	ErrKeyRequired                               = NewLegacyErrorMessage(40040, "key is required", "key is required", "kunci dibutuhkan")
	ErrInvalidReferalLimitUsageReached           = NewLegacyErrorMessage(40041, "You have reached the maximum usage of the referal code", "You have reached the maximum usage of the referal code", "Anda telah mencapai batas penggunaan maksimal kode referal")
	ErrInvalidLimitUsageReached                  = NewLegacyErrorMessage(40042, "You have reached usage limit for this promo", "You have reached usage limit for this promo", "Anda telah mencapai batas penggunaan untuk promo ini")
	ErrInvalidCodeExpired                        = NewLegacyErrorMessage(40043, "This promo code has expired", "This promo code has expired", "Kode promo ini telah kedaluwarsa")
	ErrInvalidPromo                              = NewLegacyErrorMessage(40044, "This promo code is invalid", "This promo code is invalid", "Kode promo ini tidak valid")
	ErrInvalidTime                               = NewLegacyErrorMessage(40045, "The promo time period is invalid", "The promo time period is invalid", "Periode waktu promo tidak valid")
	ErrInvalidLocation                           = NewLegacyErrorMessage(40046, "Invalid destination/pickup for this promo", "Invalid destination/pickup for this promo", "Tujuan/pengambilan tidak valid untuk promo ini")
	ErrInvalidPayment                            = NewLegacyErrorMessage(40047, "Invalid payment method for this promo", "Invalid payment method for this promo", "Metode pembayaran tidak valid untuk promo ini")
	ErrInvalidBudget                             = NewLegacyErrorMessage(40048, "Promo has reached the usage limit", "Promo has reached the usage limit", "Promo telah mencapai batas penggunaan")
	ErrInvalidCalculateBudget                    = NewLegacyErrorMessage(40049, "Promo has reached the usage limit", "Promo has reached the usage limit", "Promo telah mencapai batas penggunaan")
	ErrInvalidQuotaRedeem                        = NewLegacyErrorMessage(40050, "Promo has reached the usage limit", "Promo has reached the usage limit", "Promo telah mencapai batas penggunaan")
	ErrInvalidMaxRedemmPeriode                   = NewLegacyErrorMessage(40051, "Promo has reached the usage limit", "Promo has reached the usage limit", "Promo telah mencapai batas penggunaan")
	ErrInvalidServiceType                        = NewLegacyErrorMessage(40052, "You cannot use this promo for type service", "You cannot use this promo for type service", "Anda tidak dapat menggunakan promo ini untuk jenis layanan")
	ErrInvalidNotActive                          = NewLegacyErrorMessage(40053, "This promo code is no longer active", "This promo code is no longer active", "Kode promo ini sudah tidak aktif")
	ErrInvalidPromoDeactivate                    = NewLegacyErrorMessage(40054, "This promo code has been deactivated", "This promo code has been deactivated", "Kode promo ini telah dinonaktifkan")
	ErrInvalidRuleType                           = NewLegacyErrorMessage(40055, "The configuration doesn't match with promo code TnC", "The configuration doesn't match with promo code TnC", "Konfigurasi tidak sesuai dengan ketentuan kode promo")
	ErrInvalidQuotaRedeemPeriod                  = NewLegacyErrorMessage(40056, "Promo has reached the usage limit", "Promo has reached the usage limit this period", "Promo telah mencapai batas penggunaan periode ini")
	ErrInvalidBookingType                        = NewLegacyErrorMessage(40057, "You cannot use this promo for the selected booking type", "You cannot use this promo for the selected booking type", "Anda tidak dapat menggunakan promo ini untuk tipe pemesanan ini")
	ErrInvalidMinimumTrx                         = NewLegacyErrorMessage(40058, "The minimum transaction for this promo is not reached", "The minimum transaction for this promo is not reached", "Minimum pembayaran tidak tercapai untuk menggunakan promo ini")
	ErrInvalidNewUser                            = NewLegacyErrorMessage(40059, "This promo code only available for new users only", "This promo code only available for new users only", "Kode promo ini hanya berlaku untuk user baru saja.")
	ErrGeneralBadRequest                         = NewLegacyErrorMessage(40060, "You can't use this promo code", "You can't use this promo code", "Kamu tidak dapat menggunakan kode promo ini")
	ErrInvalidPromotionCategory                  = NewLegacyErrorMessage(40061, "Invalid promotion category", "Invalid promotion category", "Kategori promosi tidak sesuai")
	ErrSegmentation                              = NewLegacyErrorMessage(40062, "You are not eligible to claim this promo!", "You are not eligible to claim this promo!", "Kamu tidak dapat menggunakan promo ini!")
)
