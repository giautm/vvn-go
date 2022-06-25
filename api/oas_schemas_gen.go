// Code generated by ogen, DO NOT EDIT.

package api

import (
	ht "github.com/ogen-go/ogen/http"
)

type APIKey struct {
	APIKey string
}

// Thông tin đánh giá giả mạo khuôn mặt. Có thể kiểm tra thêm 3 random pose hoặc 3
// ảnh thẳng mặt liên tiếp. ``` {
// "fake_code": "REAL",
// "fake_score": 0.0,
// "status": "REAL",
// "liveness_compare_scores": [float]
// } ```.
// Ref: #/components/schemas/FaceAntiSpoofStatus
type FaceAntiSpoofStatus struct {
	FakeCode OptFaceAntiSpoofStatusFakeCode "json:\"fake_code\""
	// Mức độ giả mạo ảnh chụp chân dung, khoảng giá trị [0-1.0]:
	// - 1.0: ảnh giả
	// - 0.0: ảnh thật
	// NOTE: By our experiments, threshold to judge an image is fake is fake_score > 0.63.
	FakeScore OptFloat64 "json:\"fake_score\""
	// Trả về loại lỗi giả mạo
	// * `N/A` - ảnh selfie là ảnh thật
	// * `SCREEN` - ảnh chụp lại từ màn hình
	// * `RANDOM_POSE` - ảnh giả mạo do kiểm tra với 3 ảnh chụp quay các hướng
	// * `STRAIGHT_POSE` - ảnh giả mạo do kiểm tra 3 ảnh chụp thẳng liên tiếp.
	FakeType OptFaceAntiSpoofStatusFakeType "json:\"fake_type\""
	// Trả về khi check_3_random_pose==1. Score matching giữa ảnh live với 3 ảnh đưa vào
	// kiểm tra.
	LivenessCompareScores OptString                    "json:\"liveness_compare_scores\""
	Status                OptFaceAntiSpoofStatusStatus "json:\"status\""
}

type FaceAntiSpoofStatusFakeCode string

const (
	FaceAntiSpoofStatusFakeCodeFAKE FaceAntiSpoofStatusFakeCode = "FAKE"
	FaceAntiSpoofStatusFakeCodeREAL FaceAntiSpoofStatusFakeCode = "REAL"
)

// Trả về loại lỗi giả mạo
// * `N/A` - ảnh selfie là ảnh thật
// * `SCREEN` - ảnh chụp lại từ màn hình
// * `RANDOM_POSE` - ảnh giả mạo do kiểm tra với 3 ảnh chụp quay các hướng
// * `STRAIGHT_POSE` - ảnh giả mạo do kiểm tra 3 ảnh chụp thẳng liên tiếp.
type FaceAntiSpoofStatusFakeType string

const (
	FaceAntiSpoofStatusFakeTypeNSlashA      FaceAntiSpoofStatusFakeType = "N/A"
	FaceAntiSpoofStatusFakeTypeSCREEN       FaceAntiSpoofStatusFakeType = "SCREEN"
	FaceAntiSpoofStatusFakeTypeRANDOMPOSE   FaceAntiSpoofStatusFakeType = "RANDOM_POSE"
	FaceAntiSpoofStatusFakeTypeSTRAIGHTPOSE FaceAntiSpoofStatusFakeType = "STRAIGHT_POSE"
)

type FaceAntiSpoofStatusStatus string

const (
	FaceAntiSpoofStatusStatusFAKE FaceAntiSpoofStatusStatus = "FAKE"
	FaceAntiSpoofStatusStatusREAL FaceAntiSpoofStatusStatus = "REAL"
)

// Thông báo lỗi trả về từ Gateway.
// Ref: #/components/schemas/GatewayError
type GatewayError struct {
	// Thông điệp lỗi.
	Message OptString "json:\"message\""
}

func (*GatewayError) newFaceIDVerificationRes() {}
func (*GatewayError) newOCRRecognitionRes()     {}

// Ref: #/components/schemas/OCRInput
type OCRInput struct {
	// Giá trị ngưỡng để kiểm tra id_full về mặt đầy đủ thông tin công dân, CMND..
	// ) yêu cầu request_id của mặt trước và hay không.
	IDFullThr OptFloat32 "json:\"id_full_thr\""
	Image     string     "json:\"image\""
	RequestID string     "json:\"request_id\""
}

// Ref: #/components/schemas/OCRInput
type OCRInputForm struct {
	// Giá trị ngưỡng để kiểm tra id_full về mặt đầy đủ thông tin công dân, CMND..
	// ) yêu cầu request_id của mặt trước và hay không.
	IDFullThr OptFloat32       "json:\"id_full_thr\""
	Image     ht.MultipartFile "json:\"image\""
	RequestID string           "json:\"request_id\""
}

// Ref: #/components/schemas/OCRResult
type OCRResult struct {
	// Địa chỉ thường trú.
	Address OptString "json:\"address\""
	// Độ tin tưởng trường địa chỉ thường trú (là chuỗi string của 1 mảng các
	// giá trị float từ 0-1).
	Addressconf OptString "json:\"addressconf\""
	// Ngày sinh.
	Birthday OptString "json:\"birthday\""
	// Độ tin tưởng trường ngày tháng năm sinh (là chuỗi string của 1 mảng các giá
	// trị float từ 0-1).
	Birthdayconf OptString "json:\"birthdayconf\""
	// Loại BLX.
	Class OptString "json:\"class\""
	// Quốc gia.
	Country OptString "json:\"country\""
	// Quận/Huyện.
	District OptString "json:\"district\""
	// Thông tin mô tả các loại giấy tờ:
	// * `CCCD` - (Căn cước công dân)
	// * `NEW ID` - (CMT 12 số)
	// * `OLD ID` - (CMT 9 số)
	// * `PASSPORT`
	// * `DRIVER LICENSE OLD`
	// * `DRIVER LICENSE PET`
	// * `CHIP ID`
	// * `POLICE ID`
	// * `ARMY ID`
	// *Chú ý*: Không phân biệt được mặt sau của CMT 12 số và CCCD -> mặt sau của 2
	// loại giấy tờ này hoàn toàn giống nhau về mặt hình ảnh và ý nghĩa.
	Document OptOCRResultDocument "json:\"document\""
	// Dân tộc.
	Ethnicity OptString "json:\"ethnicity\""
	// Thời hạn giấy tờ (Với BLX có trường hợp không thời hạn). Chú ý: trường
	// ngày hết hạn có dạng dd-mm-yyyy hoặc Không thời hạn --> nếu không phù hợp
	// với các dạng này(bị che, tẩy xóa) --> giá trị nhận dạng là N/A.
	Expiry OptString "json:\"expiry\""
	// Độ tin tưởng của trường hết hạn (là chuỗi string của 1 mảng các giá trị
	// float từ 0-1).
	Expiryconf OptString "json:\"expiryconf\""
	// Nguyên quán.
	Hometown OptString "json:\"hometown\""
	// Độ tin tưởng trường nguyên quán (là chuỗi string của 1 mảng các giá trị
	// float từ 0-1).
	Hometownconf OptString "json:\"hometownconf\""
	ID           OptString "json:\"id\""
	// Kết quả kiểm tra văn bản:
	// * `FAKE` - giả mạo
	// * `CONER`
	// * `REAL` - giấy tờ thực
	// * `PUNCH` - Đục lỗ
	// * `BW` - photocopy đen trắng
	// * `REAL` - giấy tờ OK.
	IDCheck OptOCRResultIDCheck "json:\"id_check\""
	// Check xem giấy tờ có full về mặt THÔNG TIN hay không
	// * `1` - Full
	// * `0` - Không full; bị chụp thiếu
	// Note: trường hợp số ID bị che (dẫn tới score < id_full_thr (mặc định 0.8) -->
	// not full) Giấy tờ có thể không đầy đủ về mặt hình ảnh(ví dụ không có
	// hình ảnh khuôn mặt), nhưng nếu vẫn đầy đủ thông tin -- > id_full =1.
	IDFull OptInt "json:\"id_full\""
	// ID logic:
	// * `0` - check not OK
	// * `1` - check OK.
	IDLogic OptString "json:\"id_logic\""
	// Nội dung kiểm tra logic:
	// * `OK` - logic trên giấy tờ đúng theo quy định
	// * `ID is expired` - giấy tờ đã hết hạn
	// * `Not match province code` - sai mã tỉnh
	// * `Not match sex code` - Giới tính trên giấy tờ và trong số ID không trùng khớp
	// * `Not match year code` - Năm sinh trên giấy tờ và trong số ID không trùng khớp
	// * `Expiry subtract birthday not good` - Ngày tháng năm sinh và ngày tháng hết hạn không
	// hợp lệ (đối với CCCD; ngày hết hạn là khi công dân đủ 25, 40, 60 tuổi hoặc
	// không có thời hạn)
	// * `ID can be fake` - trường hợp có thể đang bị giả mạo về chữ "CĂN CƯỚC
	// CÔNG DÂN/CHỨNG MINH NHÂN DÂN" đối với giấy tờ thẻ cứng.
	IDLogicMessage OptString "json:\"id_logic_message\""
	// Phân biệt mặt trước/mặt sau:
	// * `0` - mặt trước
	// * `1` - mặt sau.
	IDType OptString "json:\"id_type\""
	// Độ tin tưởng của trường ID (là chuỗi string của 1 mảng các giá trị float
	// từ 0-1)
	// *Chú ý*: độ tin tưởng của ký tự bất kì < 0.8 --> cảnh báo đầu vào kém.
	Idconf OptString "json:\"idconf\""
	// Nơi cấp.
	IssueBy OptString "json:\"issue_by\""
	// Độ tin tưởng của trường nơi cấp (là chuỗi string của 1 mảng các giá trị
	// float từ 0-1).
	IssueByConf OptString "json:\"issue_by_conf\""
	// Ngày cấp *Chú ý*: trường ngày cấp có dạng dd-mm-yyyy --> nếu không phù hợp
	// với các dạng này(bị che, tẩy xóa) --> giá trị nhận dạng là N/A.
	IssueDate OptString "json:\"issue_date\""
	// Độ tin tưởng của trường ngày cấp (là chuỗi string của 1 mảng các giá trị
	// float từ 0-1).
	IssueDateConf OptString "json:\"issue_date_conf\""
	// Tên khách hàng trên giấy tờ.
	Name OptString "json:\"name\""
	// Độ tin tưởng trường họ tên (là chuỗi string của 1 mảng các giá trị float
	// từ 0-1).
	Nameconf OptString "json:\"nameconf\""
	// Quốc gia.
	National OptString "json:\"national\""
	// Chỉ có ý nghĩa với passport. Số CMT của Vietnam.
	OptinalData OptString "json:\"optinal_data\""
	// Loại Passport (công vụ/thường - PK/P).
	PassportType OptString "json:\"passport_type\""
	Precinct     OptString "json:\"precinct\""
	// Tỉnh/TP.
	Province OptString "json:\"province\""
	// Tôn giáo.
	Religion OptString "json:\"religion\""
	// Độ tin tưởng của trường tôn giáo (là chuỗi string của 1 mảng các giá trị
	// float từ 0-1).
	Religionconf OptString "json:\"religionconf\""
	// Mã kết quả.
	ResultCode OptOCRResultResultCode "json:\"result_code\""
	// Container id.
	ServerName OptString "json:\"server_name\""
	// Server version.
	ServerVer OptString "json:\"server_ver\""
	// Giới tính.
	Sex OptString "json:\"sex\""
	// Độ tin tưởng trường giới tính (là chuỗi string của 1 mảng các giá trị float
	// từ 0-1).
	Sexconf OptString "json:\"sexconf\""
}

func (*OCRResult) newOCRRecognitionRes() {}

// Thông tin mô tả các loại giấy tờ:
// * `CCCD` - (Căn cước công dân)
// * `NEW ID` - (CMT 12 số)
// * `OLD ID` - (CMT 9 số)
// * `PASSPORT`
// * `DRIVER LICENSE OLD`
// * `DRIVER LICENSE PET`
// * `CHIP ID`
// * `POLICE ID`
// * `ARMY ID`
// *Chú ý*: Không phân biệt được mặt sau của CMT 12 số và CCCD -> mặt sau của 2
// loại giấy tờ này hoàn toàn giống nhau về mặt hình ảnh và ý nghĩa.
type OCRResultDocument string

const (
	OCRResultDocumentCCCD             OCRResultDocument = "CCCD"
	OCRResultDocumentNEWID            OCRResultDocument = "NEW ID"
	OCRResultDocumentOLDID            OCRResultDocument = "OLD ID"
	OCRResultDocumentPASSPORT         OCRResultDocument = "PASSPORT"
	OCRResultDocumentDRIVERLICENSEOLD OCRResultDocument = "DRIVER LICENSE OLD"
	OCRResultDocumentDRIVERLICENSEPET OCRResultDocument = "DRIVER LICENSE PET"
	OCRResultDocumentCHIPID           OCRResultDocument = "CHIP ID"
	OCRResultDocumentPOLICEID         OCRResultDocument = "POLICE ID"
	OCRResultDocumentARMYID           OCRResultDocument = "ARMY ID"
)

// Kết quả kiểm tra văn bản:
// * `FAKE` - giả mạo
// * `CONER`
// * `REAL` - giấy tờ thực
// * `PUNCH` - Đục lỗ
// * `BW` - photocopy đen trắng
// * `REAL` - giấy tờ OK.
type OCRResultIDCheck string

const (
	OCRResultIDCheckBW    OCRResultIDCheck = "BW"
	OCRResultIDCheckCONER OCRResultIDCheck = "CONER"
	OCRResultIDCheckFAKE  OCRResultIDCheck = "FAKE"
	OCRResultIDCheckPUNCH OCRResultIDCheck = "PUNCH"
	OCRResultIDCheckREAL  OCRResultIDCheck = "REAL"
)

// Mã kết quả.
type OCRResultResultCode int

const (
	OCRResultResultCode200 OCRResultResultCode = 200
	OCRResultResultCode500 OCRResultResultCode = 500
	OCRResultResultCode501 OCRResultResultCode = 501
	OCRResultResultCode401 OCRResultResultCode = 401
	OCRResultResultCode402 OCRResultResultCode = 402
	OCRResultResultCode201 OCRResultResultCode = 201
)

// NewOptFaceAntiSpoofStatus returns new OptFaceAntiSpoofStatus with value set to v.
func NewOptFaceAntiSpoofStatus(v FaceAntiSpoofStatus) OptFaceAntiSpoofStatus {
	return OptFaceAntiSpoofStatus{
		Value: v,
		Set:   true,
	}
}

// OptFaceAntiSpoofStatus is optional FaceAntiSpoofStatus.
type OptFaceAntiSpoofStatus struct {
	Value FaceAntiSpoofStatus
	Set   bool
}

// IsSet returns true if OptFaceAntiSpoofStatus was set.
func (o OptFaceAntiSpoofStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFaceAntiSpoofStatus) Reset() {
	var v FaceAntiSpoofStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFaceAntiSpoofStatus) SetTo(v FaceAntiSpoofStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFaceAntiSpoofStatus) Get() (v FaceAntiSpoofStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFaceAntiSpoofStatus) Or(d FaceAntiSpoofStatus) FaceAntiSpoofStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFaceAntiSpoofStatusFakeCode returns new OptFaceAntiSpoofStatusFakeCode with value set to v.
func NewOptFaceAntiSpoofStatusFakeCode(v FaceAntiSpoofStatusFakeCode) OptFaceAntiSpoofStatusFakeCode {
	return OptFaceAntiSpoofStatusFakeCode{
		Value: v,
		Set:   true,
	}
}

// OptFaceAntiSpoofStatusFakeCode is optional FaceAntiSpoofStatusFakeCode.
type OptFaceAntiSpoofStatusFakeCode struct {
	Value FaceAntiSpoofStatusFakeCode
	Set   bool
}

// IsSet returns true if OptFaceAntiSpoofStatusFakeCode was set.
func (o OptFaceAntiSpoofStatusFakeCode) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFaceAntiSpoofStatusFakeCode) Reset() {
	var v FaceAntiSpoofStatusFakeCode
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFaceAntiSpoofStatusFakeCode) SetTo(v FaceAntiSpoofStatusFakeCode) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFaceAntiSpoofStatusFakeCode) Get() (v FaceAntiSpoofStatusFakeCode, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFaceAntiSpoofStatusFakeCode) Or(d FaceAntiSpoofStatusFakeCode) FaceAntiSpoofStatusFakeCode {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFaceAntiSpoofStatusFakeType returns new OptFaceAntiSpoofStatusFakeType with value set to v.
func NewOptFaceAntiSpoofStatusFakeType(v FaceAntiSpoofStatusFakeType) OptFaceAntiSpoofStatusFakeType {
	return OptFaceAntiSpoofStatusFakeType{
		Value: v,
		Set:   true,
	}
}

// OptFaceAntiSpoofStatusFakeType is optional FaceAntiSpoofStatusFakeType.
type OptFaceAntiSpoofStatusFakeType struct {
	Value FaceAntiSpoofStatusFakeType
	Set   bool
}

// IsSet returns true if OptFaceAntiSpoofStatusFakeType was set.
func (o OptFaceAntiSpoofStatusFakeType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFaceAntiSpoofStatusFakeType) Reset() {
	var v FaceAntiSpoofStatusFakeType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFaceAntiSpoofStatusFakeType) SetTo(v FaceAntiSpoofStatusFakeType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFaceAntiSpoofStatusFakeType) Get() (v FaceAntiSpoofStatusFakeType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFaceAntiSpoofStatusFakeType) Or(d FaceAntiSpoofStatusFakeType) FaceAntiSpoofStatusFakeType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFaceAntiSpoofStatusStatus returns new OptFaceAntiSpoofStatusStatus with value set to v.
func NewOptFaceAntiSpoofStatusStatus(v FaceAntiSpoofStatusStatus) OptFaceAntiSpoofStatusStatus {
	return OptFaceAntiSpoofStatusStatus{
		Value: v,
		Set:   true,
	}
}

// OptFaceAntiSpoofStatusStatus is optional FaceAntiSpoofStatusStatus.
type OptFaceAntiSpoofStatusStatus struct {
	Value FaceAntiSpoofStatusStatus
	Set   bool
}

// IsSet returns true if OptFaceAntiSpoofStatusStatus was set.
func (o OptFaceAntiSpoofStatusStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFaceAntiSpoofStatusStatus) Reset() {
	var v FaceAntiSpoofStatusStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFaceAntiSpoofStatusStatus) SetTo(v FaceAntiSpoofStatusStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFaceAntiSpoofStatusStatus) Get() (v FaceAntiSpoofStatusStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFaceAntiSpoofStatusStatus) Or(d FaceAntiSpoofStatusStatus) FaceAntiSpoofStatusStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat32 returns new OptFloat32 with value set to v.
func NewOptFloat32(v float32) OptFloat32 {
	return OptFloat32{
		Value: v,
		Set:   true,
	}
}

// OptFloat32 is optional float32.
type OptFloat32 struct {
	Value float32
	Set   bool
}

// IsSet returns true if OptFloat32 was set.
func (o OptFloat32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat32) Reset() {
	var v float32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat32) SetTo(v float32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat32) Get() (v float32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat32) Or(d float32) float32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMultipartFile returns new OptMultipartFile with value set to v.
func NewOptMultipartFile(v ht.MultipartFile) OptMultipartFile {
	return OptMultipartFile{
		Value: v,
		Set:   true,
	}
}

// OptMultipartFile is optional ht.MultipartFile.
type OptMultipartFile struct {
	Value ht.MultipartFile
	Set   bool
}

// IsSet returns true if OptMultipartFile was set.
func (o OptMultipartFile) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMultipartFile) Reset() {
	var v ht.MultipartFile
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMultipartFile) SetTo(v ht.MultipartFile) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMultipartFile) Get() (v ht.MultipartFile, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMultipartFile) Or(d ht.MultipartFile) ht.MultipartFile {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptOCRResultDocument returns new OptOCRResultDocument with value set to v.
func NewOptOCRResultDocument(v OCRResultDocument) OptOCRResultDocument {
	return OptOCRResultDocument{
		Value: v,
		Set:   true,
	}
}

// OptOCRResultDocument is optional OCRResultDocument.
type OptOCRResultDocument struct {
	Value OCRResultDocument
	Set   bool
}

// IsSet returns true if OptOCRResultDocument was set.
func (o OptOCRResultDocument) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOCRResultDocument) Reset() {
	var v OCRResultDocument
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOCRResultDocument) SetTo(v OCRResultDocument) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOCRResultDocument) Get() (v OCRResultDocument, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOCRResultDocument) Or(d OCRResultDocument) OCRResultDocument {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptOCRResultIDCheck returns new OptOCRResultIDCheck with value set to v.
func NewOptOCRResultIDCheck(v OCRResultIDCheck) OptOCRResultIDCheck {
	return OptOCRResultIDCheck{
		Value: v,
		Set:   true,
	}
}

// OptOCRResultIDCheck is optional OCRResultIDCheck.
type OptOCRResultIDCheck struct {
	Value OCRResultIDCheck
	Set   bool
}

// IsSet returns true if OptOCRResultIDCheck was set.
func (o OptOCRResultIDCheck) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOCRResultIDCheck) Reset() {
	var v OCRResultIDCheck
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOCRResultIDCheck) SetTo(v OCRResultIDCheck) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOCRResultIDCheck) Get() (v OCRResultIDCheck, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOCRResultIDCheck) Or(d OCRResultIDCheck) OCRResultIDCheck {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptOCRResultResultCode returns new OptOCRResultResultCode with value set to v.
func NewOptOCRResultResultCode(v OCRResultResultCode) OptOCRResultResultCode {
	return OptOCRResultResultCode{
		Value: v,
		Set:   true,
	}
}

// OptOCRResultResultCode is optional OCRResultResultCode.
type OptOCRResultResultCode struct {
	Value OCRResultResultCode
	Set   bool
}

// IsSet returns true if OptOCRResultResultCode was set.
func (o OptOCRResultResultCode) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptOCRResultResultCode) Reset() {
	var v OCRResultResultCode
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptOCRResultResultCode) SetTo(v OCRResultResultCode) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptOCRResultResultCode) Get() (v OCRResultResultCode, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptOCRResultResultCode) Or(d OCRResultResultCode) OCRResultResultCode {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptVerificationMessage returns new OptVerificationMessage with value set to v.
func NewOptVerificationMessage(v VerificationMessage) OptVerificationMessage {
	return OptVerificationMessage{
		Value: v,
		Set:   true,
	}
}

// OptVerificationMessage is optional VerificationMessage.
type OptVerificationMessage struct {
	Value VerificationMessage
	Set   bool
}

// IsSet returns true if OptVerificationMessage was set.
func (o OptVerificationMessage) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptVerificationMessage) Reset() {
	var v VerificationMessage
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptVerificationMessage) SetTo(v VerificationMessage) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptVerificationMessage) Get() (v VerificationMessage, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptVerificationMessage) Or(d VerificationMessage) VerificationMessage {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptVerificationResultVerifyResult returns new OptVerificationResultVerifyResult with value set to v.
func NewOptVerificationResultVerifyResult(v VerificationResultVerifyResult) OptVerificationResultVerifyResult {
	return OptVerificationResultVerifyResult{
		Value: v,
		Set:   true,
	}
}

// OptVerificationResultVerifyResult is optional VerificationResultVerifyResult.
type OptVerificationResultVerifyResult struct {
	Value VerificationResultVerifyResult
	Set   bool
}

// IsSet returns true if OptVerificationResultVerifyResult was set.
func (o OptVerificationResultVerifyResult) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptVerificationResultVerifyResult) Reset() {
	var v VerificationResultVerifyResult
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptVerificationResultVerifyResult) SetTo(v VerificationResultVerifyResult) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptVerificationResultVerifyResult) Get() (v VerificationResultVerifyResult, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptVerificationResultVerifyResult) Or(d VerificationResultVerifyResult) VerificationResultVerifyResult {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptVerificationResultWearingMask returns new OptVerificationResultWearingMask with value set to v.
func NewOptVerificationResultWearingMask(v VerificationResultWearingMask) OptVerificationResultWearingMask {
	return OptVerificationResultWearingMask{
		Value: v,
		Set:   true,
	}
}

// OptVerificationResultWearingMask is optional VerificationResultWearingMask.
type OptVerificationResultWearingMask struct {
	Value VerificationResultWearingMask
	Set   bool
}

// IsSet returns true if OptVerificationResultWearingMask was set.
func (o OptVerificationResultWearingMask) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptVerificationResultWearingMask) Reset() {
	var v VerificationResultWearingMask
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptVerificationResultWearingMask) SetTo(v VerificationResultWearingMask) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptVerificationResultWearingMask) Get() (v VerificationResultWearingMask, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptVerificationResultWearingMask) Or(d VerificationResultWearingMask) VerificationResultWearingMask {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/VerificationInput
type VerificationInput struct {
	// Kiểm tra các ảnh live image có phải cùng một khuôn mặt hay không.
	Check3RandomPose OptInt "json:\"check_3_random_pose\""
	// Kiểm tra các ảnh khuôn mặt có fake hay không.
	Check3StraightPose OptInt "json:\"check_3_straight_pose\""
	// Giá trị dùng để check xem khuôn mặt live có phải fake.
	FakeThreshold OptFloat64 "json:\"fake_threshold\""
	// Ảnh card id (chứng minh thư, căn cước công dân) của người dùng.
	ImageCard string "json:\"image_card\""
	// Ảnh chụp khuôn mặt.
	ImageLive string "json:\"image_live\""
	// Ảnh chụp khuôn mặt ở các góc độ khác nhau để tăng độ chính xác khi xác
	// thực khuôn mặt.
	ImageLive1 OptString "json:\"image_live1\""
	// Ảnh chụp khuôn mặt ở các góc độ khác nhau để tăng độ chính xác khi xác
	// thực khuôn mặt.
	ImageLive2 OptString "json:\"image_live2\""
	// Ảnh chụp khuôn mặt ở các góc độ khác nhau để tăng độ chính xác khi xác
	// thực khuôn mặt.
	ImageLive3 OptString "json:\"image_live3\""
	// Threshold để kiểm tra xem có đeo khẩu trang hay không.
	MaskThreshold OptFloat64 "json:\"mask_threshold\""
	// Request ID.
	RequestID string "json:\"request_id\""
	// Trả về face feature vector hay không.
	ReturnFeature OptInt "json:\"return_feature\""
	// Độ tương đồng <= sim_threshold_level1 => không cùng 1 người sim_threshold_level1 <
	// độ tương đồng < sim_threshold_level2 ==> có thể cùng 1 người độ tương đồng
	// >= sim_threshold_level2 ==> cùng 1 người.
	SimThresholdLevel1 OptFloat64 "json:\"sim_threshold_level1\""
	// Độ tương đồng <= sim_threshold_level1 => không cùng 1 người sim_threshold_level1 <
	// độ tương đồng < sim_threshold_level2 ==> có thể cùng 1 người độ tương đồng
	// >= sim_threshold_level2 ==> cùng 1 người.
	SimThresholdLevel2 OptFloat64 "json:\"sim_threshold_level2\""
}

func (*VerificationInput) newFaceIDVerificationReq() {}

// Ref: #/components/schemas/VerificationInput
type VerificationInputForm struct {
	// Kiểm tra các ảnh live image có phải cùng một khuôn mặt hay không.
	Check3RandomPose OptInt "json:\"check_3_random_pose\""
	// Kiểm tra các ảnh khuôn mặt có fake hay không.
	Check3StraightPose OptInt "json:\"check_3_straight_pose\""
	// Giá trị dùng để check xem khuôn mặt live có phải fake.
	FakeThreshold OptFloat64 "json:\"fake_threshold\""
	// Ảnh card id (chứng minh thư, căn cước công dân) của người dùng.
	ImageCard ht.MultipartFile "json:\"image_card\""
	// Ảnh chụp khuôn mặt.
	ImageLive ht.MultipartFile "json:\"image_live\""
	// Ảnh chụp khuôn mặt ở các góc độ khác nhau để tăng độ chính xác khi xác
	// thực khuôn mặt.
	ImageLive1 OptMultipartFile "json:\"image_live1\""
	// Ảnh chụp khuôn mặt ở các góc độ khác nhau để tăng độ chính xác khi xác
	// thực khuôn mặt.
	ImageLive2 OptMultipartFile "json:\"image_live2\""
	// Ảnh chụp khuôn mặt ở các góc độ khác nhau để tăng độ chính xác khi xác
	// thực khuôn mặt.
	ImageLive3 OptMultipartFile "json:\"image_live3\""
	// Threshold để kiểm tra xem có đeo khẩu trang hay không.
	MaskThreshold OptFloat64 "json:\"mask_threshold\""
	// Request ID.
	RequestID string "json:\"request_id\""
	// Trả về face feature vector hay không.
	ReturnFeature OptInt "json:\"return_feature\""
	// Độ tương đồng <= sim_threshold_level1 => không cùng 1 người sim_threshold_level1 <
	// độ tương đồng < sim_threshold_level2 ==> có thể cùng 1 người độ tương đồng
	// >= sim_threshold_level2 ==> cùng 1 người.
	SimThresholdLevel1 OptFloat64 "json:\"sim_threshold_level1\""
	// Độ tương đồng <= sim_threshold_level1 => không cùng 1 người sim_threshold_level1 <
	// độ tương đồng < sim_threshold_level2 ==> có thể cùng 1 người độ tương đồng
	// >= sim_threshold_level2 ==> cùng 1 người.
	SimThresholdLevel2 OptFloat64 "json:\"sim_threshold_level2\""
}

func (*VerificationInputForm) newFaceIDVerificationReq() {}

// Thông tin thêm về kết quả trả về. Ví dụ: ```{"api_version": 1.0.4, "error_message":
// "000", "error_code": "OK"}```.
// Ref: #/components/schemas/VerificationMessage
type VerificationMessage struct {
	// API version.
	APIVersion OptString "json:\"api_version\""
	// Mã lỗi.
	ErrorCode OptString "json:\"error_code\""
	// Mô tả lỗi.
	ErrorMessage OptString "json:\"error_message\""
}

// Ref: #/components/schemas/VerificationResult
type VerificationResult struct {
	FaceAntiSpoofStatus OptFaceAntiSpoofStatus "json:\"face_anti_spoof_status\""
	// Feature vector của khuôn mặt ở ảnh thẻ.
	FeatureVectorFaceCard []float64 "json:\"feature_vector_face_card\""
	// Feature vector của khuôn mặt ở ảnh chụp chân dung.
	FeatureVectorFaceLive []float64              "json:\"feature_vector_face_live\""
	Message               OptVerificationMessage "json:\"message\""
	// Độ tương đồng của ảnh khuôn mặt từ thẻ và ảnh chân dung [0-1].
	Sim OptFloat64 "json:\"sim\""
	// Thời gian thực hiện việc xác thực ở phía server (đơn vị mili giây - ms).
	VerificationTime OptInt "json:\"verification_time\""
	// Kết quả xác thực dạng số:
	// * `0` - not same
	// * `1` - may be same
	// * `2` - same person.
	VerifyResult OptVerificationResultVerifyResult "json:\"verify_result\""
	// Kết quả xác thực dạng text.
	VerifyResultText OptString                        "json:\"verify_result_text\""
	WearingMask      OptVerificationResultWearingMask "json:\"wearing_mask\""
	// Điểm score khi nhận dạng khẩu trang, điểm càng cao thì độ tin tưởng đeo
	// khẩu trang càng lớn score trong khoảng [0-1].
	WearingMaskScore OptFloat64 "json:\"wearing_mask_score\""
}

func (*VerificationResult) newFaceIDVerificationRes() {}

// Kết quả xác thực dạng số:
// * `0` - not same
// * `1` - may be same
// * `2` - same person.
type VerificationResultVerifyResult int

const (
	VerificationResultVerifyResult0 VerificationResultVerifyResult = 0
	VerificationResultVerifyResult1 VerificationResultVerifyResult = 1
	VerificationResultVerifyResult2 VerificationResultVerifyResult = 2
)

type VerificationResultWearingMask string

const (
	VerificationResultWearingMaskYES VerificationResultWearingMask = "YES"
	VerificationResultWearingMaskNO  VerificationResultWearingMask = "NO"
)
