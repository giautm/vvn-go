// Code generated by ogen, DO NOT EDIT.

package api

import (
	ht "github.com/ogen-go/ogen/http"
)

type APIKey struct {
	APIKey string
}

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
// Ref: #/components/schemas/DocumentEnum
type DocumentEnum string

const (
	DocumentEnumCCCD             DocumentEnum = "CCCD"
	DocumentEnumNEWID            DocumentEnum = "NEW ID"
	DocumentEnumOLDID            DocumentEnum = "OLD ID"
	DocumentEnumPASSPORT         DocumentEnum = "PASSPORT"
	DocumentEnumDRIVERLICENSEOLD DocumentEnum = "DRIVER LICENSE OLD"
	DocumentEnumDRIVERLICENSEPET DocumentEnum = "DRIVER LICENSE PET"
	DocumentEnumCHIPID           DocumentEnum = "CHIP ID"
	DocumentEnumPOLICEID         DocumentEnum = "POLICE ID"
	DocumentEnumARMYID           DocumentEnum = "ARMY ID"
)

// Thông tin đánh giá giả mạo khuôn mặt. Có thể kiểm tra thêm 3 random pose hoặc 3
// ảnh thẳng mặt liên tiếp. ``` {
// "fake_code": "REAL",
// "fake_score": 0.0,
// "status": "REAL",
// "liveness_compare_scores": [float]
// } ```.
// Ref: #/components/schemas/FaceAntiSpoofStatus
type FaceAntiSpoofStatus struct {
	FakeCode OptSpoofStatusEnum "json:\"fake_code\""
	// Mức độ giả mạo ảnh chụp chân dung, khoảng giá trị [0-1.0]:
	// - 1.0: ảnh giả
	// - 0.0: ảnh thật
	// NOTE: By our experiments, threshold to judge an image is fake is fake_score > 0.63.
	FakeScore OptFloat64      "json:\"fake_score\""
	FakeType  OptFakeTypeEnum "json:\"fake_type\""
	// Trả về khi check_3_random_pose==1. Score matching giữa ảnh live với 3 ảnh đưa vào
	// kiểm tra.
	LivenessCompareScores OptString          "json:\"liveness_compare_scores\""
	Status                OptSpoofStatusEnum "json:\"status\""
}

type FaceFeature []float64

// Ref: #/components/schemas/FaceIDRecognitionInput
type FaceIDRecognitionInput struct {
	Image     string "json:\"image\""
	RequestID string "json:\"request_id\""
}

func (*FaceIDRecognitionInput) faceRecognitionReq() {}

// Ref: #/components/schemas/FaceIDRecognitionInput
type FaceIDRecognitionInputForm struct {
	Image     ht.MultipartFile "json:\"image\""
	RequestID string           "json:\"request_id\""
}

func (*FaceIDRecognitionInputForm) faceRecognitionReq() {}

// Ref: #/components/schemas/FaceIDRecognitionResult
type FaceIDRecognitionResult struct {
	Message Message "json:\"message\""
	// List of photos with faces similar to the one imported.
	RecognitionResult []FaceRecognitionResult "json:\"recognition_result\""
	// Time handle request.
	RecognitionTime OptFloat64 "json:\"recognition_time\""
}

func (*FaceIDRecognitionResult) faceRecognitionRes() {}

// Merged schema.
// Ref: #/components/schemas/FaceIDRegisterInput
type FaceIDRegisterInput struct {
	// Customer is unique name. Used in case many customers have the same real name.
	UniqueName string "json:\"unique_name\""
	// Optional field,request system overrides customer registration. Used in case this photo or similar
	// photo of the customer has been previously registered.
	// * `0` - Default registration is not required
	// * `1` - Registration required.
	Force OptInt "json:\"force\""
	// Image of consumer.
	Image string "json:\"image\""
	// Name of consumer.
	PersonName OptString "json:\"person_name\""
}

func (*FaceIDRegisterInput) faceRegisterReq() {}

// Merged schema.
// Ref: #/components/schemas/FaceIDRegisterInput
type FaceIDRegisterInputForm struct {
	// Customer is unique name. Used in case many customers have the same real name.
	UniqueName string "json:\"unique_name\""
	// Optional field,request system overrides customer registration. Used in case this photo or similar
	// photo of the customer has been previously registered.
	// * `0` - Default registration is not required
	// * `1` - Registration required.
	Force OptInt "json:\"force\""
	// Image of consumer.
	Image ht.MultipartFile "json:\"image\""
	// Name of consumer.
	PersonName OptString "json:\"person_name\""
}

func (*FaceIDRegisterInputForm) faceRegisterReq() {}

// Ref: #/components/schemas/FaceIDRegisterResult
type FaceIDRegisterResult struct {
	// Góc nghiêng của khuôn mặt ở ảnh.
	FaceCardAngle OptInt "json:\"face_card_angle\""
	// ID of face register in database.
	FaceID OptInt "json:\"face_id\""
	// Position of faces in photos. Format: [left, top, right, bottom].
	FaceLoc       []int      "json:\"face_loc\""
	MatchedScore  OptFloat64 "json:\"matched_score\""
	Message       Message    "json:\"message\""
	SamePersonThr OptFloat64 "json:\"same_person_thr\""
	UniqueName    OptString  "json:\"unique_name\""
}

func (*FaceIDRegisterResult) faceRegisterRes() {}

// Merged schema.
// Ref: #/components/schemas/FaceIDResult
type FaceIDResult struct {
	// Customer is unique name. Used in case many customers have the same real name.
	UniqueName   string           "json:\"unique_name\""
	CompareScore float64          "json:\"compare_score\""
	Info         FaceIDResultInfo "json:\"info\""
	// Name of consumer.
	PersonName OptString "json:\"person_name\""
}

type FaceIDResultInfo map[string]string

func (s *FaceIDResultInfo) init() FaceIDResultInfo {
	m := *s
	if m == nil {
		m = map[string]string{}
		*s = m
	}
	return m
}

// Ref: #/components/schemas/FaceIDVerificationInput
type FaceIDVerificationInput struct {
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

func (*FaceIDVerificationInput) faceVerificationReq() {}

// Ref: #/components/schemas/FaceIDVerificationInput
type FaceIDVerificationInputForm struct {
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

func (*FaceIDVerificationInputForm) faceVerificationReq() {}

// Ref: #/components/schemas/FaceIDVerificationResult
type FaceIDVerificationResult struct {
	FaceAntiSpoofStatus OptFaceAntiSpoofStatus "json:\"face_anti_spoof_status\""
	// Góc nghiêng của khuôn mặt ở ảnh card id.
	FaceCardAngle OptInt "json:\"face_card_angle\""
	// Góc nghiêng của khuôn mặt ở ảnh chụp chân dung.
	FaceLiveAngle OptInt "json:\"face_live_angle\""
	// Vị trí của khuôn mặt detect được trong ảnh card id.
	FaceLocCard *FaceLocation "json:\"face_loc_card\""
	// Vị trí của khuôn mặt detect được trong ảnh chụp chân dung.
	FaceLocLive *FaceLocation "json:\"face_loc_live\""
	// Feature vector của khuôn mặt ở ảnh thẻ.
	FeatureVectorFaceCard *FaceFeature "json:\"feature_vector_face_card\""
	// Feature vector của khuôn mặt ở ảnh chụp chân dung.
	FeatureVectorFaceLive *FaceFeature "json:\"feature_vector_face_live\""
	Message               Message      "json:\"message\""
	// Request ID.
	RequestID string "json:\"request_id\""
	// Độ tương đồng của ảnh khuôn mặt từ thẻ và ảnh chân dung [0-1].
	Sim OptFloat64 "json:\"sim\""
	// Thời gian thực hiện việc xác thực ở phía server (đơn vị mili giây - ms).
	VerificationTime OptFloat64              "json:\"verification_time\""
	VerifyResult     OptFaceVerifyResultEnum "json:\"verify_result\""
	// Kết quả xác thực dạng text.
	VerifyResultText OptString         "json:\"verify_result_text\""
	WearingMask      OptMaskResultEnum "json:\"wearing_mask\""
	// Điểm score khi nhận dạng khẩu trang, điểm càng cao thì độ tin tưởng đeo
	// khẩu trang càng lớn score trong khoảng [0-1].
	WearingMaskScore OptFloat64 "json:\"wearing_mask_score\""
}

func (*FaceIDVerificationResult) faceVerificationRes() {}

type FaceLocation []int

// Ref: #/components/schemas/FaceRecognitionResult
type FaceRecognitionResult struct {
	FaceLoc FaceLocation "json:\"face_loc\""
	// Size of face in pixels.
	FaceSize int "json:\"face_size\""
	// List of top-k faces with highest confidence.
	TopK []FaceIDResult "json:\"top_k\""
}

type FaceUnregisterOK struct {
	Message Message "json:\"message\""
}

func (*FaceUnregisterOK) faceUnregisterRes() {}

// Merged schema.
type FaceUnregisterReq struct {
	// Customer is unique name. Used in case many customers have the same real name.
	UniqueName string "json:\"unique_name\""
	// Customer is Face ID, returned in the result when registering a Customer is photo of the face in
	// the system. If face_id is not provided, all customer photos with unique_name will be removed from
	// the system.
	FaceID OptString "json:\"face_id\""
}

// Kết quả xác thực dạng số:
// * `0` - not same
// * `1` - may be same
// * `2` - same person.
// Ref: #/components/schemas/FaceVerifyResultEnum
type FaceVerifyResultEnum int

const (
	FaceVerifyResultEnum0 FaceVerifyResultEnum = 0
	FaceVerifyResultEnum1 FaceVerifyResultEnum = 1
	FaceVerifyResultEnum2 FaceVerifyResultEnum = 2
)

// Trả về loại lỗi giả mạo
// * `N/A` - ảnh selfie là ảnh thật
// * `SCREEN` - ảnh chụp lại từ màn hình
// * `RANDOM_POSE` - ảnh giả mạo do kiểm tra với 3 ảnh chụp quay các hướng
// * `STRAIGHT_POSE` - ảnh giả mạo do kiểm tra 3 ảnh chụp thẳng liên tiếp.
// Ref: #/components/schemas/FakeTypeEnum
type FakeTypeEnum string

const (
	FakeTypeEnumNSlashA      FakeTypeEnum = "N/A"
	FakeTypeEnumSCREEN       FakeTypeEnum = "SCREEN"
	FakeTypeEnumRANDOMPOSE   FakeTypeEnum = "RANDOM_POSE"
	FakeTypeEnumSTRAIGHTPOSE FakeTypeEnum = "STRAIGHT_POSE"
)

// Thông báo lỗi trả về từ Gateway.
// Ref: #/components/schemas/GatewayError
type GatewayError struct {
	// Thông điệp lỗi.
	Message string "json:\"message\""
}

func (*GatewayError) faceRecognitionRes()  {}
func (*GatewayError) faceRegisterRes()     {}
func (*GatewayError) faceUnregisterRes()   {}
func (*GatewayError) faceVerificationRes() {}
func (*GatewayError) oCRecognitionRes()    {}

// Kết quả kiểm tra văn bản:
// * `FAKE` - giả mạo
// * `CONER`
// * `REAL` - giấy tờ thực
// * `PUNCH` - Đục lỗ
// * `BW` - photocopy đen trắng
// * `REAL` - giấy tờ OK.
// Ref: #/components/schemas/IDCheckEnum
type IDCheckEnum string

const (
	IDCheckEnumBW    IDCheckEnum = "BW"
	IDCheckEnumCONER IDCheckEnum = "CONER"
	IDCheckEnumFAKE  IDCheckEnum = "FAKE"
	IDCheckEnumPUNCH IDCheckEnum = "PUNCH"
	IDCheckEnumREAL  IDCheckEnum = "REAL"
)

// Phân biệt mặt trước/mặt sau:
// * `0` - mặt trước
// * `1` - mặt sau.
// Ref: #/components/schemas/IDTypeEnum
type IDTypeEnum string

const (
	IDTypeEnum0 IDTypeEnum = "0"
	IDTypeEnum1 IDTypeEnum = "1"
)

// Ref: #/components/schemas/MaskResultEnum
type MaskResultEnum string

const (
	MaskResultEnumYES MaskResultEnum = "YES"
	MaskResultEnumNO  MaskResultEnum = "NO"
)

// Thông tin thêm về kết quả trả về. Ví dụ:
// ```{"error_message": "000", "error_code": "OK"}```.
// Ref: #/components/schemas/Message
type Message struct {
	// API version.
	APIVersion OptString "json:\"api_version\""
	// Copyright.
	CopyRight OptString "json:\"copy_right\""
	// Mã lỗi.
	ErrorCode string "json:\"error_code\""
	// Mô tả lỗi.
	ErrorMessage string "json:\"error_message\""
	// Information.
	Info OptString "json:\"info\""
}

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
	// Đặt điểm.
	Characteristics OptString "json:\"characteristics\""
	// Độ tin tưởng của trường đặt điểm (là chuỗi string của 1 mảng các giá
	// trị float từ 0-1).
	CharacteristicsConf OptString "json:\"characteristics_conf\""
	// Loại BLX.
	Class     OptString "json:\"class\""
	Copyright OptString "json:\"copyright\""
	// Quốc gia.
	Country OptString "json:\"country\""
	// Quận/Huyện.
	District OptString       "json:\"district\""
	Document OptDocumentEnum "json:\"document\""
	// Dân tộc.
	Ethnicity OptString "json:\"ethnicity\""
	// Độ tin tưởng của trường dân tộc (là chuỗi string của 1 mảng các giá trị
	// float từ 0-1).
	Ethnicityconf OptString "json:\"ethnicityconf\""
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
	Hometownconf OptString      "json:\"hometownconf\""
	ID           OptString      "json:\"id\""
	IDCheck      OptIDCheckEnum "json:\"id_check\""
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
	IDLogicMessage OptString     "json:\"id_logic_message\""
	IDType         OptIDTypeEnum "json:\"id_type\""
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
	OptinalData  OptString "json:\"optinal_data\""
	OptionalData OptString "json:\"optional_data\""
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
	// Đường phố.
	Street OptString "json:\"street\""
	// Tên đường.
	StreetName OptString "json:\"street_name\""
}

func (*OCRResult) oCRecognitionRes() {}

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

// NewOptDocumentEnum returns new OptDocumentEnum with value set to v.
func NewOptDocumentEnum(v DocumentEnum) OptDocumentEnum {
	return OptDocumentEnum{
		Value: v,
		Set:   true,
	}
}

// OptDocumentEnum is optional DocumentEnum.
type OptDocumentEnum struct {
	Value DocumentEnum
	Set   bool
}

// IsSet returns true if OptDocumentEnum was set.
func (o OptDocumentEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDocumentEnum) Reset() {
	var v DocumentEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDocumentEnum) SetTo(v DocumentEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDocumentEnum) Get() (v DocumentEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDocumentEnum) Or(d DocumentEnum) DocumentEnum {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

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

// NewOptFaceVerifyResultEnum returns new OptFaceVerifyResultEnum with value set to v.
func NewOptFaceVerifyResultEnum(v FaceVerifyResultEnum) OptFaceVerifyResultEnum {
	return OptFaceVerifyResultEnum{
		Value: v,
		Set:   true,
	}
}

// OptFaceVerifyResultEnum is optional FaceVerifyResultEnum.
type OptFaceVerifyResultEnum struct {
	Value FaceVerifyResultEnum
	Set   bool
}

// IsSet returns true if OptFaceVerifyResultEnum was set.
func (o OptFaceVerifyResultEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFaceVerifyResultEnum) Reset() {
	var v FaceVerifyResultEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFaceVerifyResultEnum) SetTo(v FaceVerifyResultEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFaceVerifyResultEnum) Get() (v FaceVerifyResultEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFaceVerifyResultEnum) Or(d FaceVerifyResultEnum) FaceVerifyResultEnum {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFakeTypeEnum returns new OptFakeTypeEnum with value set to v.
func NewOptFakeTypeEnum(v FakeTypeEnum) OptFakeTypeEnum {
	return OptFakeTypeEnum{
		Value: v,
		Set:   true,
	}
}

// OptFakeTypeEnum is optional FakeTypeEnum.
type OptFakeTypeEnum struct {
	Value FakeTypeEnum
	Set   bool
}

// IsSet returns true if OptFakeTypeEnum was set.
func (o OptFakeTypeEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFakeTypeEnum) Reset() {
	var v FakeTypeEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFakeTypeEnum) SetTo(v FakeTypeEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFakeTypeEnum) Get() (v FakeTypeEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFakeTypeEnum) Or(d FakeTypeEnum) FakeTypeEnum {
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

// NewOptIDCheckEnum returns new OptIDCheckEnum with value set to v.
func NewOptIDCheckEnum(v IDCheckEnum) OptIDCheckEnum {
	return OptIDCheckEnum{
		Value: v,
		Set:   true,
	}
}

// OptIDCheckEnum is optional IDCheckEnum.
type OptIDCheckEnum struct {
	Value IDCheckEnum
	Set   bool
}

// IsSet returns true if OptIDCheckEnum was set.
func (o OptIDCheckEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptIDCheckEnum) Reset() {
	var v IDCheckEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptIDCheckEnum) SetTo(v IDCheckEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptIDCheckEnum) Get() (v IDCheckEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptIDCheckEnum) Or(d IDCheckEnum) IDCheckEnum {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptIDTypeEnum returns new OptIDTypeEnum with value set to v.
func NewOptIDTypeEnum(v IDTypeEnum) OptIDTypeEnum {
	return OptIDTypeEnum{
		Value: v,
		Set:   true,
	}
}

// OptIDTypeEnum is optional IDTypeEnum.
type OptIDTypeEnum struct {
	Value IDTypeEnum
	Set   bool
}

// IsSet returns true if OptIDTypeEnum was set.
func (o OptIDTypeEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptIDTypeEnum) Reset() {
	var v IDTypeEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptIDTypeEnum) SetTo(v IDTypeEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptIDTypeEnum) Get() (v IDTypeEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptIDTypeEnum) Or(d IDTypeEnum) IDTypeEnum {
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

// NewOptMaskResultEnum returns new OptMaskResultEnum with value set to v.
func NewOptMaskResultEnum(v MaskResultEnum) OptMaskResultEnum {
	return OptMaskResultEnum{
		Value: v,
		Set:   true,
	}
}

// OptMaskResultEnum is optional MaskResultEnum.
type OptMaskResultEnum struct {
	Value MaskResultEnum
	Set   bool
}

// IsSet returns true if OptMaskResultEnum was set.
func (o OptMaskResultEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMaskResultEnum) Reset() {
	var v MaskResultEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMaskResultEnum) SetTo(v MaskResultEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMaskResultEnum) Get() (v MaskResultEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMaskResultEnum) Or(d MaskResultEnum) MaskResultEnum {
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

// NewOptSpoofStatusEnum returns new OptSpoofStatusEnum with value set to v.
func NewOptSpoofStatusEnum(v SpoofStatusEnum) OptSpoofStatusEnum {
	return OptSpoofStatusEnum{
		Value: v,
		Set:   true,
	}
}

// OptSpoofStatusEnum is optional SpoofStatusEnum.
type OptSpoofStatusEnum struct {
	Value SpoofStatusEnum
	Set   bool
}

// IsSet returns true if OptSpoofStatusEnum was set.
func (o OptSpoofStatusEnum) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptSpoofStatusEnum) Reset() {
	var v SpoofStatusEnum
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptSpoofStatusEnum) SetTo(v SpoofStatusEnum) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptSpoofStatusEnum) Get() (v SpoofStatusEnum, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptSpoofStatusEnum) Or(d SpoofStatusEnum) SpoofStatusEnum {
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

// Ref: #/components/schemas/SpoofStatusEnum
type SpoofStatusEnum string

const (
	SpoofStatusEnumFAKE SpoofStatusEnum = "FAKE"
	SpoofStatusEnumREAL SpoofStatusEnum = "REAL"
)
