package utils

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Error is a struct of error obj
// code is for error code
// status code is the first 3 digits of the error code
type Error struct {
	Code             int                    `json:"-"`
	CodeStr          string                 `json:"code"`
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

func NewLegacyErrorMessage(code int, message string, data map[string]interface{}, localizedMessage Message) *Error {
	return &Error{Code: code, Message: message, Data: data, LocalizedMessage: localizedMessage}
}

func FromLegacyError(err Error) error {
	statusCode := HttpToGRPCStatus[err.Code/100]

	if statusCode == 0 {
		statusCode = HttpToGRPCStatus[http.StatusBadRequest]
	}
	return status.Error(statusCode, fmt.Sprintf("%d|%s|%s|%s", err.Code, err.Message, err.LocalizedMessage.English, err.LocalizedMessage.Indonesia))
}

func NewError(code int, message, en, id string) error {
	statusCode := HttpToGRPCStatus[code/100]

	if statusCode == 0 {
		statusCode = HttpToGRPCStatus[http.StatusBadRequest]
	}
	return status.Error(statusCode, fmt.Sprintf("%d|%s|%s|%s", code, message, en, id))
}

func NewBadRequestError(err error) error {
	return NewError(40000, err.Error(), "", "")
}

var (
	ErrPromotionSingleUseNotExists               = NewError(40000, "Sorry! promotion single use not exists!", "Sorry! promotion single use not exists!", "Maaf! Promosi sekali pakai tidak ada!")
	ErrInternal                                  = NewError(50000, "Sorry. Unable to complete your request", "Sorry. Unable to complete your request", "Maaf. Tidak dapat menyelesaikan permintaan Anda")
	ErrAuth                                      = NewError(40006, "Failed to authenticate to Promotion Portal", "Failed to authenticate to Promotion Portal", "Gagal melakukan autentikasi ke Portal Promosi")
	ErrTokenExpired                              = NewError(40101, "Sorry, token is expired", "Sorry, token is expired", "Maaf, token telah kedaluwarsa")
	ErrTokenInvalid                              = NewError(40102, "Soory, token is invalid", "Soory, token is invalid", "Maaf, token tidak valid")
	ErrPromoCodeIsUsed                           = NewError(40000, "promo code already to used!", "promo code already to used!", "promo code telah digunakan!")
	ErrBadRequest                                = NewError(40000, "Oops! Something went wrong!", "Oops! Something went wrong!", "Ups! Ada yang salah!")
	ErrInvalidInput                              = NewError(40001, "Oops! Something went wrong!", "Oops! Something went wrong!", "Ups! Ada yang salah!")
	ErrPromotionNotExists                        = NewError(40002, "Sorry! promotion not exists!", "Sorry! promotion not exists!", "Maaf! Promosi tidak ada!")
	ErrPromotionAlreadyUsed                      = NewError(40003, "Promotion is Already to used!", "Promotion is Already to used!", "Promosi sudah digunakan!")
	ErrPromotionVersion                          = NewError(40004, "Sorry, you can't create promotion version", "Sorry, you can't create promotion version", "Maaf, Anda tidak dapat membuat versi promosi")
	ErrStatusPromotionEdit                       = NewError(40005, "This promo code is not allowed to update because the status isn't `Terjadwal`", "This promo code is not allowed to update because the status isn't `Terjadwal`", "Kode promo ini tidak diizinkan untuk diperbarui karena statusnya bukan `Terjadwal`")
	ErrGetPromotionTimeRangeByPromoID            = NewError(40006, "Failed to get promotion time ranges BY Promo ID", "Failed to get promotion time ranges BY Promo ID", "Gagal mendapatkan rentang waktu promosi BERDASARKAN ID Promo")
	ErrAccountInvalid                            = NewError(40007, "Your account name is invalid", "Your account name is invalid", "Nama akun Anda tidak valid")
	ErrCodeRequired                              = NewError(40008, "code is required", "code is required", "kode dibutuhkan")
	ErrIdentifierRequired                        = NewError(40009, "identifier is required", "identifier is required", "identifier dibutuhkan")
	ErrServiceTypeRequired                       = NewError(40010, "service_type is required", "service_type is required", "service_type dibutuhkan")
	ErrPaymentTypeRequired                       = NewError(40011, "payment_type is required", "payment_type is required", "payment_type dibutuhkan")
	ErrPromotionIDInvalid                        = NewError(40012, "promotion id not valid, using other id", "promotion id not valid, using other id", "ID promosi tidak valid, gunakan ID lain")
	ErrReactivatedPromoInvalidBecauseOverPeriode = NewError(40013, "failed to reactivated promotion code, because promo code is expired", "failed to reactivated promotion code, because promo code is expired", "gagal mengaktifkan kembali kode promo, karena kode promo telah kedaluwarsa")
	ErrReactivatedPromoInvalidBecauseLimitBudget = NewError(40014, "failed to reactivated promotion code, because promo code has used all the budget", "failed to reactivated promotion code, because promo code has used all the budget", "gagal mengaktifkan kembali kode promo, karena kode promo telah menggunakan semua anggaran")
	ErrPromoLimitBudget                          = NewError(40015, "the budget for this promo has reached the limit", "the budget for this promo has reached the limit", "anggaran untuk promo ini telah mencapai batasnya")
	ErrGetPromotionRules                         = NewError(40016, "Failed to get promotion rules", "Failed to get promotion rules", "Gagal mendapatkan aturan promosi")
	ErrInvalidStatus                             = NewError(40017, "invalid promotion status", "invalid promotion status", "status promosi tidak valid")
	ErrInvalidPromotionID                        = NewError(40018, "invalid promotion id or promotion not exists", "invalid promotion id or promotion not exists", "ID promosi tidak valid atau promosi tidak ada")
	ErrIdentifierMissing                         = NewError(40019, "identifier is required", "identifier is required", "identifier dibutuhkan")
	ErrMaxReferralCodeUsage                      = NewError(40020, "You have reached max usage of referal code", "You have reached max usage of referal code", "Anda telah mencapai batas penggunaan kode referal")
	ErrMaxPromotionCodeUsage                     = NewError(40021, "Your promotion max usage has been reached", "Your promotion max usage has been reached", "Penggunaan maksimal promosi Anda telah tercapai")
	ErrPromotionExpired                          = NewError(40022, "Your promotion has expired", "Your promotion has expired", "Promosi Anda telah kedaluwarsa")
	ErrPromotionServiceTypeInvalid               = NewError(40023, "Invalid service type", "Invalid service type", "Jenis layanan tidak valid untuk promosi ini")
	ErrPromotionPaymentTypeInvalid               = NewError(40024, "Invalid payment type", "Invalid payment type", "Jenis pembayaran tidak valid untuk promosi ini")
	ErrPromotionDeactivated                      = NewError(40025, "Your promotion code is deactivated", "Your promotion code is deactivated", "Kode promo Anda dinonaktifkan")
	ErrPromotionInvalid                          = NewError(40025, "Your promotion code is invalid", "Your promotion code is invalid", "Kode promo Anda tidak valid")
	ErrPromotionInvalidPickupTime                = NewError(40026, "Your promotion code is invalid for current pickup time", "Your promotion code is invalid for current pickup time", "Kode promo Anda tidak valid untuk waktu pengambilan saat ini")
	ErrPromotionInvalidPickupLoc                 = NewError(40027, "Your pickup location is invalid for this promotion", "Your pickup location is invalid for this promotion", "Lokasi pengambilan Anda tidak valid untuk promosi ini")
	ErrPromotionInvalidDropoffLoc                = NewError(40028, "Your dropoff location is invalid for this promotion", "Your dropoff location is invalid for this promotion", "Lokasi penurunan Anda tidak valid untuk promosi ini")
	ErrMaxReferralUsage                          = NewError(40029, "You have reached max usage of referral code", "You have reached max usage of referral code", "Anda telah mencapai batas penggunaan kode referral")
	ErrPromoCodeInvalid                          = NewError(40030, "Your promotion code is invalid", "Your promotion code is invalid", "Kode promo Anda tidak valid")
	ErrPromoOverlaps                             = NewError(40031, "Promotion is overlapping", "Promotion is overlapping", "Promosi tumpang tindih")
	ErrDiscountValueInvalid                      = NewError(40032, "Invalid discount value", "Invalid discount value", "Nilai diskon tidak valid")
	ErrDiscountTypeInvalid                       = NewError(40033, "Invalid discount type", "Invalid discount type", "Jenis diskon tidak valid")
	ErrEndDateInvalid                            = NewError(40034, "End Date could not be before Start Date", "End Date could not be before Start Date", "Tanggal Akhir tidak boleh sebelum Tanggal Mulai")
	ErrServiceTypeInvalid                        = NewError(40035, "Invalid service type", "Invalid service type", "Jenis layanan tidak valid")
	ErrPaymentTypeInvalid                        = NewError(40036, "Invalid payment type", "Invalid payment type", "Jenis pembayaran tidak valid")
	ErrPromoCodeMissing                          = NewError(40037, "promotion_code is required", "promotion_code is required", "kode promo dibutuhkan")
	ErrIdentifierInvalid                         = NewError(40038, "Kamu tidak bisa menggunakan kode promo ini", "You can't use this promo code", "Kamu tidak bisa menggunakan kode promo ini")
	ErrServiceUnavailable                        = NewError(40039, "Service Unavailable", "Service Unavailable", "Layanan tidak tersedia saat ini, silakan coba lagi nanti")
	ErrKeyRequired                               = NewError(40040, "key is required", "key is required", "kunci dibutuhkan")
	ErrInvalidReferalLimitUsageReached           = NewError(40041, "You have reached the maximum usage of the referal code", "You have reached the maximum usage of the referal code", "Anda telah mencapai batas penggunaan maksimal kode referal")
	ErrInvalidLimitUsageReached                  = NewError(40042, "You have reached usage limit for this promo", "You have reached usage limit for this promo", "Anda telah mencapai batas penggunaan untuk promo ini")
	ErrInvalidCodeExpired                        = NewError(40043, "This promo code has expired", "This promo code has expired", "Kode promo ini telah kedaluwarsa")
	ErrInvalidPromo                              = NewError(40044, "This promo code is invalid", "This promo code is invalid", "Kode promo ini tidak valid")
	ErrInvalidTime                               = NewError(40045, "The promo time period is invalid", "The promo time period is invalid", "Periode waktu promo tidak valid")
	ErrInvalidLocation                           = NewError(40046, "Invalid destination/pickup for this promo", "Invalid destination/pickup for this promo", "Tujuan/pengambilan tidak valid untuk promo ini")
	ErrInvalidPayment                            = NewError(40047, "Invalid payment method for this promo", "Invalid payment method for this promo", "Metode pembayaran tidak valid untuk promo ini")
	ErrInvalidBudget                             = NewError(40048, "Promo has reached the usage limit", "Promo has reached the usage limit", "Promo telah mencapai batas penggunaan")
	ErrInvalidCalculateBudget                    = NewError(40049, "Promo has reached the usage limit", "Promo has reached the usage limit", "Promo telah mencapai batas penggunaan")
	ErrInvalidQuotaRedeem                        = NewError(40050, "Promo has reached the usage limit", "Promo has reached the usage limit", "Promo telah mencapai batas penggunaan")
	ErrInvalidMaxRedemmPeriode                   = NewError(40051, "Promo has reached the usage limit", "Promo has reached the usage limit", "Promo telah mencapai batas penggunaan")
	ErrInvalidServiceType                        = NewError(40052, "You cannot use this promo for type service", "You cannot use this promo for type service", "Anda tidak dapat menggunakan promo ini untuk jenis layanan")
	ErrInvalidNotActive                          = NewError(40053, "This promo code is no longer active", "This promo code is no longer active", "Kode promo ini sudah tidak aktif")
	ErrInvalidPromoDeactivate                    = NewError(40054, "This promo code has been deactivated", "This promo code has been deactivated", "Kode promo ini telah dinonaktifkan")
	ErrInvalidRuleType                           = NewError(40055, "The configuration doesn't match with promo code TnC", "The configuration doesn't match with promo code TnC", "Konfigurasi tidak sesuai dengan ketentuan kode promo")
	ErrInvalidQuotaRedeemPeriod                  = NewError(40056, "Promo has reached the usage limit", "Promo has reached the usage limit this period", "Promo telah mencapai batas penggunaan periode ini")
	ErrInvalidBookingType                        = NewError(40057, "You cannot use this promo for the selected booking type", "You cannot use this promo for the selected booking type", "Anda tidak dapat menggunakan promo ini untuk tipe pemesanan ini")
	ErrInvalidMinimumTrx                         = NewError(40058, "The minimum transaction for this promo is not reached", "The minimum transaction for this promo is not reached", "Minimum pembayaran tidak tercapai untuk menggunakan promo ini")
	ErrInvalidNewUser                            = NewError(40059, "This promo code only available for new users only", "This promo code only available for new users only", "Kode promo ini hanya berlaku untuk user baru saja.")
	ErrGeneralBadRequest                         = NewError(40060, "You can't use this promo code", "You can't use this promo code", "Kamu tidak dapat menggunakan kode promo ini")
	ErrInvalidPromotionCategory                  = NewError(40061, "Invalid promotion category", "Invalid promotion category", "Kategori promosi tidak sesuai")
	ErrSegmentation                              = NewError(40062, "You are not eligible to claim this promo!", "You are not eligible to claim this promo!", "Kamu tidak dapat menggunakan promo ini!")
)

var HttpToGRPCStatus = map[int]codes.Code{
	http.StatusOK:                            codes.OK,
	http.StatusCreated:                       codes.OK,
	http.StatusAccepted:                      codes.OK,
	http.StatusNonAuthoritativeInfo:          codes.OK, // No exact match, using OK
	http.StatusNoContent:                     codes.OK, // No exact match, using OK
	http.StatusResetContent:                  codes.OK, // No exact match, using OK
	http.StatusPartialContent:                codes.OK, // No exact match, using OK
	http.StatusMultiStatus:                   codes.OK, // No exact match, using OK
	http.StatusAlreadyReported:               codes.OK, // No exact match, using OK
	http.StatusIMUsed:                        codes.OK, // No exact match, using OK
	http.StatusBadRequest:                    codes.InvalidArgument,
	http.StatusUnauthorized:                  codes.Unauthenticated,
	http.StatusPaymentRequired:               codes.Unimplemented, // Not an official HTTP status
	http.StatusForbidden:                     codes.PermissionDenied,
	http.StatusNotFound:                      codes.NotFound,
	http.StatusMethodNotAllowed:              codes.PermissionDenied, // No exact match, using PermissionDenied
	http.StatusNotAcceptable:                 codes.InvalidArgument,  // No exact match, using InvalidArgument
	http.StatusProxyAuthRequired:             codes.PermissionDenied, // No exact match, using PermissionDenied
	http.StatusRequestTimeout:                codes.DeadlineExceeded,
	http.StatusConflict:                      codes.Aborted,
	http.StatusGone:                          codes.FailedPrecondition,
	http.StatusLengthRequired:                codes.InvalidArgument, // No exact match, using InvalidArgument
	http.StatusPreconditionFailed:            codes.FailedPrecondition,
	http.StatusRequestEntityTooLarge:         codes.InvalidArgument,    // No exact match, using InvalidArgument
	http.StatusRequestURITooLong:             codes.InvalidArgument,    // No exact match, using InvalidArgument
	http.StatusUnsupportedMediaType:          codes.InvalidArgument,    // No exact match, using InvalidArgument
	http.StatusRequestedRangeNotSatisfiable:  codes.InvalidArgument,    // No exact match, using InvalidArgument
	http.StatusExpectationFailed:             codes.InvalidArgument,    // No exact match, using InvalidArgument
	http.StatusTeapot:                        codes.FailedPrecondition, // No exact match, using FailedPrecondition
	http.StatusMisdirectedRequest:            codes.Unknown,            // No exact match, using Unknown
	http.StatusUnprocessableEntity:           codes.InvalidArgument,    // No exact match, using InvalidArgument
	http.StatusLocked:                        codes.FailedPrecondition, // No exact match, using FailedPrecondition
	http.StatusFailedDependency:              codes.FailedPrecondition,
	http.StatusTooEarly:                      codes.Unknown,            // No exact match, using Unknown
	http.StatusUpgradeRequired:               codes.FailedPrecondition, // No exact match, using FailedPrecondition
	http.StatusPreconditionRequired:          codes.FailedPrecondition,
	http.StatusTooManyRequests:               codes.ResourceExhausted,
	http.StatusRequestHeaderFieldsTooLarge:   codes.InvalidArgument, // No exact match, using InvalidArgument
	http.StatusUnavailableForLegalReasons:    codes.PermissionDenied,
	http.StatusInternalServerError:           codes.Internal,
	http.StatusNotImplemented:                codes.Unimplemented,
	http.StatusBadGateway:                    codes.Internal, // No exact match, using Internal
	http.StatusServiceUnavailable:            codes.Unavailable,
	http.StatusGatewayTimeout:                codes.DeadlineExceeded,  // No exact match, using DeadlineExceeded
	http.StatusHTTPVersionNotSupported:       codes.Unimplemented,     // No exact match, using Unimplemented
	http.StatusVariantAlsoNegotiates:         codes.Unknown,           // No exact match, using Unknown
	http.StatusInsufficientStorage:           codes.ResourceExhausted, // No exact match, using ResourceExhausted
	http.StatusLoopDetected:                  codes.Aborted,           // No exact match, using Aborted
	http.StatusNotExtended:                   codes.Unknown,           // No exact match, using Unknown
	http.StatusNetworkAuthenticationRequired: codes.PermissionDenied,
}
