package codex

type ResCode int

const (
	CodeSuccess ResCode = 200

	CodeInvalidParams ResCode = 400

	CodeNeedLogin       ResCode = 401
	CodeCasbinError     ResCode = 401001
	CodeBusinessTypeErr ResCode = 401002

	CodeInvalidToken            ResCode = 403
	CodeWrongPassword           ResCode = 403001
	CodeWrongUserNameOrPassword ResCode = 403002
	CodeWrongSmsCode            ResCode = 403003
	CodeSendSmsCodeErr          ResCode = 403004
	CodeCacheSmsCodeErr         ResCode = 403005
	CodeDelCacheSmsCodeErr      ResCode = 403006
	CodePasswordNotEqual        ResCode = 403007
	CodeCacheEmailCodeErr       ResCode = 403008
	CodeSendEmailErr            ResCode = 403009
	CodeWrongEmailCode          ResCode = 403010
	CodeConfirmPasswordError    ResCode = 403011
	CodeOldPasswordError        ResCode = 403012
	CodeEmailCodeExpired        ResCode = 403013
	CodeInviteLinkExpired       ResCode = 403014
	CodeSMSCodeExpired          ResCode = 403015
	CodeRoleTypeErr             ResCode = 403016
	CodeParamsNilError          ResCode = 403017
	CodeDatabaseError           ResCode = 403018

	CodeUserNotExist ResCode = 404001
	CodeUserExist    ResCode = 409002
	CodeEmailExist   ResCode = 409003
	CodePhoneExist   ResCode = 409004
	CodeCompanyExist ResCode = 409005
	CodeQueryUserErr ResCode = 409006

	CodeInternalErr      ResCode = 500
	CodeGenTokenErr      ResCode = 500001
	CodeParseFormFileErr ResCode = 500002
	CodeFileFormErr      ResCode = 500003
	CodeCreateFileErr    ResCode = 500004
)
