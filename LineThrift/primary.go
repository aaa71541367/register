package LineThrift

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"

	"github.com/shillbie/register/thrift"
)

type AuthErrorCode int64

const (
	AuthErrorCode_INTERNAL_ERROR              AuthErrorCode = 0
	AuthErrorCode_ILLEGAL_ARGUMENT            AuthErrorCode = 1
	AuthErrorCode_VERIFICATION_FAILED         AuthErrorCode = 2
	AuthErrorCode_NOT_FOUND                   AuthErrorCode = 3
	AuthErrorCode_RETRY_LATER                 AuthErrorCode = 4
	AuthErrorCode_HUMAN_VERIFICATION_REQUIRED AuthErrorCode = 5
	AuthErrorCode_INVALID_CONTEXT             AuthErrorCode = 100
	AuthErrorCode_APP_UPGRADE_REQUIRED        AuthErrorCode = 101
)

func (p AuthErrorCode) String() string {
	switch p {
	case AuthErrorCode_INTERNAL_ERROR:
		return "INTERNAL_ERROR"
	case AuthErrorCode_ILLEGAL_ARGUMENT:
		return "ILLEGAL_ARGUMENT"
	case AuthErrorCode_VERIFICATION_FAILED:
		return "VERIFICATION_FAILED"
	case AuthErrorCode_NOT_FOUND:
		return "NOT_FOUND"
	case AuthErrorCode_RETRY_LATER:
		return "RETRY_LATER"
	case AuthErrorCode_HUMAN_VERIFICATION_REQUIRED:
		return "HUMAN_VERIFICATION_REQUIRED"
	case AuthErrorCode_INVALID_CONTEXT:
		return "INVALID_CONTEXT"
	case AuthErrorCode_APP_UPGRADE_REQUIRED:
		return "APP_UPGRADE_REQUIRED"
	}
	return "<UNSET>"
}

func AuthErrorCodeFromString(s string) (AuthErrorCode, error) {
	switch s {
	case "INTERNAL_ERROR":
		return AuthErrorCode_INTERNAL_ERROR, nil
	case "ILLEGAL_ARGUMENT":
		return AuthErrorCode_ILLEGAL_ARGUMENT, nil
	case "VERIFICATION_FAILED":
		return AuthErrorCode_VERIFICATION_FAILED, nil
	case "NOT_FOUND":
		return AuthErrorCode_NOT_FOUND, nil
	case "RETRY_LATER":
		return AuthErrorCode_RETRY_LATER, nil
	case "HUMAN_VERIFICATION_REQUIRED":
		return AuthErrorCode_HUMAN_VERIFICATION_REQUIRED, nil
	case "INVALID_CONTEXT":
		return AuthErrorCode_INVALID_CONTEXT, nil
	case "APP_UPGRADE_REQUIRED":
		return AuthErrorCode_APP_UPGRADE_REQUIRED, nil
	}
	return AuthErrorCode(0), fmt.Errorf("not a valid AuthErrorCode string")
}

func AuthErrorCodePtr(v AuthErrorCode) *AuthErrorCode { return &v }

func (p AuthErrorCode) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *AuthErrorCode) UnmarshalText(text []byte) error {
	q, err := AuthErrorCodeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *AuthErrorCode) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = AuthErrorCode(v)
	return nil
}

func (p *AuthErrorCode) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type AvailableMethods int64

const (
	AvailableMethods_UNKNOWN AvailableMethods = 0
	AvailableMethods_SMS     AvailableMethods = 1
	AvailableMethods_IVR     AvailableMethods = 2
	AvailableMethods_SMSPULL AvailableMethods = 3
)

func (p AvailableMethods) String() string {
	switch p {
	case AvailableMethods_UNKNOWN:
		return "UNKNOWN"
	case AvailableMethods_SMS:
		return "SMS"
	case AvailableMethods_IVR:
		return "IVR"
	case AvailableMethods_SMSPULL:
		return "SMSPULL"
	}
	return "<UNSET>"
}

func AvailableMethodsFromString(s string) (AvailableMethods, error) {
	switch s {
	case "UNKNOWN":
		return AvailableMethods_UNKNOWN, nil
	case "SMS":
		return AvailableMethods_SMS, nil
	case "IVR":
		return AvailableMethods_IVR, nil
	case "SMSPULL":
		return AvailableMethods_SMSPULL, nil
	}
	return AvailableMethods(0), fmt.Errorf("not a valid AvailableMethods string")
}

func AvailableMethodsPtr(v AvailableMethods) *AvailableMethods { return &v }

func (p AvailableMethods) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *AvailableMethods) UnmarshalText(text []byte) error {
	q, err := AvailableMethodsFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *AvailableMethods) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = AvailableMethods(v)
	return nil
}

func (p *AvailableMethods) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type SecondAuthMethod int64

const (
	SecondAuthMethod_UNKNOWN   SecondAuthMethod = 0
	SecondAuthMethod_SKIP      SecondAuthMethod = 1
	SecondAuthMethod_WEB_BASED SecondAuthMethod = 2
)

func (p SecondAuthMethod) String() string {
	switch p {
	case SecondAuthMethod_UNKNOWN:
		return "UNKNOWN"
	case SecondAuthMethod_SKIP:
		return "SKIP"
	case SecondAuthMethod_WEB_BASED:
		return "WEB_BASED"
	}
	return "<UNSET>"
}

func SecondAuthMethodFromString(s string) (SecondAuthMethod, error) {
	switch s {
	case "UNKNOWN":
		return SecondAuthMethod_UNKNOWN, nil
	case "SKIP":
		return SecondAuthMethod_SKIP, nil
	case "WEB_BASED":
		return SecondAuthMethod_WEB_BASED, nil
	}
	return SecondAuthMethod(0), fmt.Errorf("not a valid SecondAuthMethod string")
}

func SecondAuthMethodPtr(v SecondAuthMethod) *SecondAuthMethod { return &v }

func (p SecondAuthMethod) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *SecondAuthMethod) UnmarshalText(text []byte) error {
	q, err := SecondAuthMethodFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *SecondAuthMethod) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = SecondAuthMethod(v)
	return nil
}

func (p *SecondAuthMethod) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type AccountIdentifierType int64

const (
	AccountIdentifierType_UNKNOWN      AccountIdentifierType = 0
	AccountIdentifierType_PHONE_NUMBER AccountIdentifierType = 1
	AccountIdentifierType_EMAIL        AccountIdentifierType = 2
)

func (p AccountIdentifierType) String() string {
	switch p {
	case AccountIdentifierType_UNKNOWN:
		return "UNKNOWN"
	case AccountIdentifierType_PHONE_NUMBER:
		return "PHONE_NUMBER"
	case AccountIdentifierType_EMAIL:
		return "EMAIL"
	}
	return "<UNSET>"
}

func AccountIdentifierTypeFromString(s string) (AccountIdentifierType, error) {
	switch s {
	case "UNKNOWN":
		return AccountIdentifierType_UNKNOWN, nil
	case "PHONE_NUMBER":
		return AccountIdentifierType_PHONE_NUMBER, nil
	case "EMAIL":
		return AccountIdentifierType_EMAIL, nil
	}
	return AccountIdentifierType(0), fmt.Errorf("not a valid AccountIdentifierType string")
}

func AccountIdentifierTypePtr(v AccountIdentifierType) *AccountIdentifierType { return &v }

func (p AccountIdentifierType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *AccountIdentifierType) UnmarshalText(text []byte) error {
	q, err := AccountIdentifierTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *AccountIdentifierType) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = AccountIdentifierType(v)
	return nil
}

func (p *AccountIdentifierType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type EncryptionKeyVersion int64

const (
	EncryptionKeyVersion_UNKNOWN EncryptionKeyVersion = 0
	EncryptionKeyVersion_V1      EncryptionKeyVersion = 1
)

func (p EncryptionKeyVersion) String() string {
	switch p {
	case EncryptionKeyVersion_UNKNOWN:
		return "UNKNOWN"
	case EncryptionKeyVersion_V1:
		return "V1"
	}
	return "<UNSET>"
}

func EncryptionKeyVersionFromString(s string) (EncryptionKeyVersion, error) {
	switch s {
	case "UNKNOWN":
		return EncryptionKeyVersion_UNKNOWN, nil
	case "V1":
		return EncryptionKeyVersion_V1, nil
	}
	return EncryptionKeyVersion(0), fmt.Errorf("not a valid EncryptionKeyVersion string")
}

func EncryptionKeyVersionPtr(v EncryptionKeyVersion) *EncryptionKeyVersion { return &v }

func (p EncryptionKeyVersion) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *EncryptionKeyVersion) UnmarshalText(text []byte) error {
	q, err := EncryptionKeyVersionFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *EncryptionKeyVersion) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = EncryptionKeyVersion(v)
	return nil
}

func (p *EncryptionKeyVersion) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type SocialLoginType int64

const (
	SocialLoginType_UNKNOWN  SocialLoginType = 0
	SocialLoginType_FACEBOOK SocialLoginType = 1
	SocialLoginType_APPLE    SocialLoginType = 2
)

func (p SocialLoginType) String() string {
	switch p {
	case SocialLoginType_UNKNOWN:
		return "UNKNOWN"
	case SocialLoginType_FACEBOOK:
		return "FACEBOOK"
	case SocialLoginType_APPLE:
		return "APPLE"
	}
	return "<UNSET>"
}

func SocialLoginTypeFromString(s string) (SocialLoginType, error) {
	switch s {
	case "UNKNOWN":
		return SocialLoginType_UNKNOWN, nil
	case "FACEBOOK":
		return SocialLoginType_FACEBOOK, nil
	case "APPLE":
		return SocialLoginType_APPLE, nil
	}
	return SocialLoginType(0), fmt.Errorf("not a valid SocialLoginType string")
}

func SocialLoginTypePtr(v SocialLoginType) *SocialLoginType { return &v }

func (p SocialLoginType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *SocialLoginType) UnmarshalText(text []byte) error {
	q, err := SocialLoginTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *SocialLoginType) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = SocialLoginType(v)
	return nil
}

func (p *SocialLoginType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type VerifMethod int64

const (
	VerifMethod_UNKNOWN     VerifMethod = 0
	VerifMethod_UNAVAILABLE VerifMethod = 1
	VerifMethod_SMS         VerifMethod = 2
	VerifMethod_TTS         VerifMethod = 3
)

func (p VerifMethod) String() string {
	switch p {
	case VerifMethod_UNKNOWN:
		return "UNKNOWN"
	case VerifMethod_UNAVAILABLE:
		return "UNAVAILABLE"
	case VerifMethod_SMS:
		return "SMS"
	case VerifMethod_TTS:
		return "TTS"
	}
	return "<UNSET>"
}

func VerifMethodFromString(s string) (VerifMethod, error) {
	switch s {
	case "UNKNOWN":
		return VerifMethod_UNKNOWN, nil
	case "UNAVAILABLE":
		return VerifMethod_UNAVAILABLE, nil
	case "SMS":
		return VerifMethod_SMS, nil
	case "TTS":
		return VerifMethod_TTS, nil
	}
	return VerifMethod(0), fmt.Errorf("not a valid VerifMethod string")
}

func VerifMethodPtr(v VerifMethod) *VerifMethod { return &v }

func (p VerifMethod) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *VerifMethod) UnmarshalText(text []byte) error {
	q, err := VerifMethodFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *VerifMethod) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = VerifMethod(v)
	return nil
}

func (p *VerifMethod) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type AvailableMethod int64

const (
	AvailableMethod_UNKNOWN   AvailableMethod = 0
	AvailableMethod_SKIP      AvailableMethod = 1
	AvailableMethod_PASSWORD  AvailableMethod = 2
	AvailableMethod_WEB_BASED AvailableMethod = 3
)

func (p AvailableMethod) String() string {
	switch p {
	case AvailableMethod_UNKNOWN:
		return "UNKNOWN"
	case AvailableMethod_SKIP:
		return "SKIP"
	case AvailableMethod_PASSWORD:
		return "PASSWORD"
	case AvailableMethod_WEB_BASED:
		return "WEB_BASED"
	}
	return "<UNSET>"
}

func AvailableMethodFromString(s string) (AvailableMethod, error) {
	switch s {
	case "UNKNOWN":
		return AvailableMethod_UNKNOWN, nil
	case "SKIP":
		return AvailableMethod_SKIP, nil
	case "PASSWORD":
		return AvailableMethod_PASSWORD, nil
	case "WEB_BASED":
		return AvailableMethod_WEB_BASED, nil
	}
	return AvailableMethod(0), fmt.Errorf("not a valid AvailableMethod string")
}

func AvailableMethodPtr(v AvailableMethod) *AvailableMethod { return &v }

func (p AvailableMethod) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *AvailableMethod) UnmarshalText(text []byte) error {
	q, err := AvailableMethodFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *AvailableMethod) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = AvailableMethod(v)
	return nil
}

func (p *AvailableMethod) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// Attributes:
//  - Code
//  - AlertMessage
//  - WebAuthDetails
type AuthException struct {
	Code         AuthErrorCode `thrift:"code,1" db:"code" json:"code"`
	AlertMessage string        `thrift:"alertMessage,2" db:"alertMessage" json:"alertMessage"`
	// unused fields # 3 to 10
	WebAuthDetails *WebAuthDetails `thrift:"webAuthDetails,11" db:"webAuthDetails" json:"webAuthDetails"`
}

func NewAuthException() *AuthException {
	return &AuthException{}
}

func (p *AuthException) GetCode() AuthErrorCode {
	return p.Code
}

func (p *AuthException) GetAlertMessage() string {
	return p.AlertMessage
}

var AuthException_WebAuthDetails_DEFAULT *WebAuthDetails

func (p *AuthException) GetWebAuthDetails() *WebAuthDetails {
	if !p.IsSetWebAuthDetails() {
		return AuthException_WebAuthDetails_DEFAULT
	}
	return p.WebAuthDetails
}
func (p *AuthException) IsSetWebAuthDetails() bool {
	return p.WebAuthDetails != nil
}

func (p *AuthException) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 11:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField11(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AuthException) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := AuthErrorCode(v)
		p.Code = temp
	}
	return nil
}

func (p *AuthException) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AlertMessage = v
	}
	return nil
}

func (p *AuthException) ReadField11(ctx context.Context, iprot thrift.TProtocol) error {
	p.WebAuthDetails = &WebAuthDetails{}
	if err := p.WebAuthDetails.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.WebAuthDetails), err)
	}
	return nil
}

func (p *AuthException) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "AuthException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField11(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AuthException) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "code", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Code)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err)
	}
	return err
}

func (p *AuthException) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "alertMessage", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:alertMessage: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AlertMessage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.alertMessage (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:alertMessage: ", p), err)
	}
	return err
}

func (p *AuthException) writeField11(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "webAuthDetails", thrift.STRUCT, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:webAuthDetails: ", p), err)
	}
	if err := p.WebAuthDetails.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.WebAuthDetails), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:webAuthDetails: ", p), err)
	}
	return err
}

func (p *AuthException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuthException(%+v)", *p)
}

func (p *AuthException) Error() string {
	return p.String()
}

// Attributes:
//  - AuthToken
type RegisterPrimaryResponse struct {
	AuthToken string `thrift:"authToken,1" db:"authToken" json:"authToken"`
}

func NewRegisterPrimaryResponse() *RegisterPrimaryResponse {
	return &RegisterPrimaryResponse{}
}

func (p *RegisterPrimaryResponse) GetAuthToken() string {
	return p.AuthToken
}
func (p *RegisterPrimaryResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RegisterPrimaryResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthToken = v
	}
	return nil
}

func (p *RegisterPrimaryResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "RegisterPrimaryResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RegisterPrimaryResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authToken", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authToken (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authToken: ", p), err)
	}
	return err
}

func (p *RegisterPrimaryResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RegisterPrimaryResponse(%+v)", *p)
}

// Attributes:
//  - AuthToken
type MigratePrimaryResponse struct {
	AuthToken string `thrift:"authToken,1" db:"authToken" json:"authToken"`
}

func NewMigratePrimaryResponse() *MigratePrimaryResponse {
	return &MigratePrimaryResponse{}
}

func (p *MigratePrimaryResponse) GetAuthToken() string {
	return p.AuthToken
}
func (p *MigratePrimaryResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MigratePrimaryResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthToken = v
	}
	return nil
}

func (p *MigratePrimaryResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "MigratePrimaryResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MigratePrimaryResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authToken", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authToken (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authToken: ", p), err)
	}
	return err
}

func (p *MigratePrimaryResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MigratePrimaryResponse(%+v)", *p)
}

// Attributes:
//  - SecondAuthMethod
type GetSecondAuthMethodResponse struct {
	SecondAuthMethod SecondAuthMethod `thrift:"secondAuthMethod,1" db:"secondAuthMethod" json:"secondAuthMethod"`
}

func NewGetSecondAuthMethodResponse() *GetSecondAuthMethodResponse {
	return &GetSecondAuthMethodResponse{}
}

func (p *GetSecondAuthMethodResponse) GetSecondAuthMethod() SecondAuthMethod {
	return p.SecondAuthMethod
}
func (p *GetSecondAuthMethodResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetSecondAuthMethodResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := SecondAuthMethod(v)
		p.SecondAuthMethod = temp
	}
	return nil
}

func (p *GetSecondAuthMethodResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "GetSecondAuthMethodResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetSecondAuthMethodResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "secondAuthMethod", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:secondAuthMethod: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.SecondAuthMethod)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.secondAuthMethod (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:secondAuthMethod: ", p), err)
	}
	return err
}

func (p *GetSecondAuthMethodResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetSecondAuthMethodResponse(%+v)", *p)
}

// Attributes:
//  - BaseUrl
//  - Token
type WebAuthDetails struct {
	BaseUrl string `thrift:"baseUrl,1" db:"baseUrl" json:"baseUrl"`
	Token   string `thrift:"token,2" db:"token" json:"token"`
}

func NewWebAuthDetails() *WebAuthDetails {
	return &WebAuthDetails{}
}

func (p *WebAuthDetails) GetBaseUrl() string {
	return p.BaseUrl
}

func (p *WebAuthDetails) GetToken() string {
	return p.Token
}
func (p *WebAuthDetails) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *WebAuthDetails) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.BaseUrl = v
	}
	return nil
}

func (p *WebAuthDetails) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Token = v
	}
	return nil
}

func (p *WebAuthDetails) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "WebAuthDetails"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *WebAuthDetails) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "baseUrl", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:baseUrl: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.BaseUrl)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.baseUrl (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:baseUrl: ", p), err)
	}
	return err
}

func (p *WebAuthDetails) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "token", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:token: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Token)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.token (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:token: ", p), err)
	}
	return err
}

func (p *WebAuthDetails) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("WebAuthDetails(%+v)", *p)
}

// Attributes:
//  - WebAuthDetails
type IssueWebAuthDetailsForSecondAuthResponse struct {
	WebAuthDetails *WebAuthDetails `thrift:"webAuthDetails,1" db:"webAuthDetails" json:"webAuthDetails"`
}

func NewIssueWebAuthDetailsForSecondAuthResponse() *IssueWebAuthDetailsForSecondAuthResponse {
	return &IssueWebAuthDetailsForSecondAuthResponse{}
}

var IssueWebAuthDetailsForSecondAuthResponse_WebAuthDetails_DEFAULT *WebAuthDetails

func (p *IssueWebAuthDetailsForSecondAuthResponse) GetWebAuthDetails() *WebAuthDetails {
	if !p.IsSetWebAuthDetails() {
		return IssueWebAuthDetailsForSecondAuthResponse_WebAuthDetails_DEFAULT
	}
	return p.WebAuthDetails
}
func (p *IssueWebAuthDetailsForSecondAuthResponse) IsSetWebAuthDetails() bool {
	return p.WebAuthDetails != nil
}

func (p *IssueWebAuthDetailsForSecondAuthResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IssueWebAuthDetailsForSecondAuthResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.WebAuthDetails = &WebAuthDetails{}
	if err := p.WebAuthDetails.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.WebAuthDetails), err)
	}
	return nil
}

func (p *IssueWebAuthDetailsForSecondAuthResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "IssueWebAuthDetailsForSecondAuthResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IssueWebAuthDetailsForSecondAuthResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "webAuthDetails", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:webAuthDetails: ", p), err)
	}
	if err := p.WebAuthDetails.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.WebAuthDetails), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:webAuthDetails: ", p), err)
	}
	return err
}

func (p *IssueWebAuthDetailsForSecondAuthResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IssueWebAuthDetailsForSecondAuthResponse(%+v)", *p)
}

// Attributes:
//  - WebAuthDetails
type IssueWebAuthDetailsForAcctVerifResponse struct {
	WebAuthDetails *WebAuthDetails `thrift:"webAuthDetails,1" db:"webAuthDetails" json:"webAuthDetails"`
}

func NewIssueWebAuthDetailsForAcctVerifResponse() *IssueWebAuthDetailsForAcctVerifResponse {
	return &IssueWebAuthDetailsForAcctVerifResponse{}
}

var IssueWebAuthDetailsForAcctVerifResponse_WebAuthDetails_DEFAULT *WebAuthDetails

func (p *IssueWebAuthDetailsForAcctVerifResponse) GetWebAuthDetails() *WebAuthDetails {
	if !p.IsSetWebAuthDetails() {
		return IssueWebAuthDetailsForAcctVerifResponse_WebAuthDetails_DEFAULT
	}
	return p.WebAuthDetails
}
func (p *IssueWebAuthDetailsForAcctVerifResponse) IsSetWebAuthDetails() bool {
	return p.WebAuthDetails != nil
}

func (p *IssueWebAuthDetailsForAcctVerifResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IssueWebAuthDetailsForAcctVerifResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.WebAuthDetails = &WebAuthDetails{}
	if err := p.WebAuthDetails.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.WebAuthDetails), err)
	}
	return nil
}

func (p *IssueWebAuthDetailsForAcctVerifResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "IssueWebAuthDetailsForAcctVerifResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IssueWebAuthDetailsForAcctVerifResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "webAuthDetails", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:webAuthDetails: ", p), err)
	}
	if err := p.WebAuthDetails.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.WebAuthDetails), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:webAuthDetails: ", p), err)
	}
	return err
}

func (p *IssueWebAuthDetailsForAcctVerifResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IssueWebAuthDetailsForAcctVerifResponse(%+v)", *p)
}

// Attributes:
//  - Type
//  - Identifier
//  - CountryCode
type AccountIdentifier struct {
	Type       AccountIdentifierType `thrift:"type,1" db:"type" json:"type"`
	Identifier string                `thrift:"identifier,2" db:"identifier" json:"identifier"`
	// unused fields # 3 to 10
	CountryCode string `thrift:"countryCode,11" db:"countryCode" json:"countryCode"`
}

func NewAccountIdentifier() *AccountIdentifier {
	return &AccountIdentifier{}
}

func (p *AccountIdentifier) GetType() AccountIdentifierType {
	return p.Type
}

func (p *AccountIdentifier) GetIdentifier() string {
	return p.Identifier
}

func (p *AccountIdentifier) GetCountryCode() string {
	return p.CountryCode
}
func (p *AccountIdentifier) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 11:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField11(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountIdentifier) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := AccountIdentifierType(v)
		p.Type = temp
	}
	return nil
}

func (p *AccountIdentifier) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Identifier = v
	}
	return nil
}

func (p *AccountIdentifier) ReadField11(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 11: ", err)
	} else {
		p.CountryCode = v
	}
	return nil
}

func (p *AccountIdentifier) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "AccountIdentifier"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField11(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountIdentifier) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "type", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:type: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Type)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.type (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:type: ", p), err)
	}
	return err
}

func (p *AccountIdentifier) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "identifier", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:identifier: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Identifier)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.identifier (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:identifier: ", p), err)
	}
	return err
}

func (p *AccountIdentifier) writeField11(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "countryCode", thrift.STRING, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:countryCode: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.CountryCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.countryCode (11) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:countryCode: ", p), err)
	}
	return err
}

func (p *AccountIdentifier) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountIdentifier(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type OpenSessionResponse struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewOpenSessionResponse() *OpenSessionResponse {
	return &OpenSessionResponse{}
}

func (p *OpenSessionResponse) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *OpenSessionResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *OpenSessionResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *OpenSessionResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "OpenSessionResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OpenSessionResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *OpenSessionResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OpenSessionResponse(%+v)", *p)
}

// Attributes:
//  - MetaData
type OpenSessionRequest struct {
	MetaData map[string]string `thrift:"metaData,1" db:"metaData" json:"metaData"`
}

func NewOpenSessionRequest() *OpenSessionRequest {
	return &OpenSessionRequest{}
}

func (p *OpenSessionRequest) GetMetaData() map[string]string {
	return p.MetaData
}
func (p *OpenSessionRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.MAP {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *OpenSessionRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]string, size)
	p.MetaData = tMap
	for i := 0; i < size; i++ {
		var _key0 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key0 = v
		}
		var _val1 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_val1 = v
		}
		p.MetaData[_key0] = _val1
	}
	if err := iprot.ReadMapEnd(ctx); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *OpenSessionRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "OpenSessionRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OpenSessionRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "metaData", thrift.MAP, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:metaData: ", p), err)
	}
	if err := oprot.WriteMapBegin(ctx, thrift.STRING, thrift.STRING, len(p.MetaData)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.MetaData {
		if err := oprot.WriteString(ctx, string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteMapEnd(ctx); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:metaData: ", p), err)
	}
	return err
}

func (p *OpenSessionRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OpenSessionRequest(%+v)", *p)
}

// Attributes:
//  - UserProfile
type VerifyAccountUsingPwdResponse struct {
	// unused field # 1
	UserProfile *UserProfile `thrift:"userProfile,2" db:"userProfile" json:"userProfile"`
}

func NewVerifyAccountUsingPwdResponse() *VerifyAccountUsingPwdResponse {
	return &VerifyAccountUsingPwdResponse{}
}

var VerifyAccountUsingPwdResponse_UserProfile_DEFAULT *UserProfile

func (p *VerifyAccountUsingPwdResponse) GetUserProfile() *UserProfile {
	if !p.IsSetUserProfile() {
		return VerifyAccountUsingPwdResponse_UserProfile_DEFAULT
	}
	return p.UserProfile
}
func (p *VerifyAccountUsingPwdResponse) IsSetUserProfile() bool {
	return p.UserProfile != nil
}

func (p *VerifyAccountUsingPwdResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *VerifyAccountUsingPwdResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserProfile = &UserProfile{}
	if err := p.UserProfile.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserProfile), err)
	}
	return nil
}

func (p *VerifyAccountUsingPwdResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "VerifyAccountUsingPwdResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifyAccountUsingPwdResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userProfile", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:userProfile: ", p), err)
	}
	if err := p.UserProfile.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserProfile), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:userProfile: ", p), err)
	}
	return err
}

func (p *VerifyAccountUsingPwdResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifyAccountUsingPwdResponse(%+v)", *p)
}

// Attributes:
//  - EncryptionKeyVersion
//  - CipherText
type EncryptedPassword struct {
	EncryptionKeyVersion EncryptionKeyVersion `thrift:"encryptionKeyVersion,1" db:"encryptionKeyVersion" json:"encryptionKeyVersion"`
	CipherText           string               `thrift:"cipherText,2" db:"cipherText" json:"cipherText"`
}

func NewEncryptedPassword() *EncryptedPassword {
	return &EncryptedPassword{}
}

func (p *EncryptedPassword) GetEncryptionKeyVersion() EncryptionKeyVersion {
	return p.EncryptionKeyVersion
}

func (p *EncryptedPassword) GetCipherText() string {
	return p.CipherText
}
func (p *EncryptedPassword) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EncryptedPassword) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := EncryptionKeyVersion(v)
		p.EncryptionKeyVersion = temp
	}
	return nil
}

func (p *EncryptedPassword) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.CipherText = v
	}
	return nil
}

func (p *EncryptedPassword) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "EncryptedPassword"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EncryptedPassword) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "encryptionKeyVersion", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:encryptionKeyVersion: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.EncryptionKeyVersion)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.encryptionKeyVersion (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:encryptionKeyVersion: ", p), err)
	}
	return err
}

func (p *EncryptedPassword) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "cipherText", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:cipherText: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.CipherText)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.cipherText (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:cipherText: ", p), err)
	}
	return err
}

func (p *EncryptedPassword) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EncryptedPassword(%+v)", *p)
}

// Attributes:
//  - AccountExist
//  - UserProfile
//  - SameUdidFromAccount
type VerifySocialLoginResponse struct {
	// unused field # 1
	AccountExist bool `thrift:"accountExist,2" db:"accountExist" json:"accountExist"`
	// unused fields # 3 to 10
	UserProfile         *UserProfile `thrift:"userProfile,11" db:"userProfile" json:"userProfile"`
	SameUdidFromAccount bool         `thrift:"sameUdidFromAccount,12" db:"sameUdidFromAccount" json:"sameUdidFromAccount"`
}

func NewVerifySocialLoginResponse() *VerifySocialLoginResponse {
	return &VerifySocialLoginResponse{}
}

func (p *VerifySocialLoginResponse) GetAccountExist() bool {
	return p.AccountExist
}

var VerifySocialLoginResponse_UserProfile_DEFAULT *UserProfile

func (p *VerifySocialLoginResponse) GetUserProfile() *UserProfile {
	if !p.IsSetUserProfile() {
		return VerifySocialLoginResponse_UserProfile_DEFAULT
	}
	return p.UserProfile
}

func (p *VerifySocialLoginResponse) GetSameUdidFromAccount() bool {
	return p.SameUdidFromAccount
}
func (p *VerifySocialLoginResponse) IsSetUserProfile() bool {
	return p.UserProfile != nil
}

func (p *VerifySocialLoginResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 11:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField11(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 12:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField12(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *VerifySocialLoginResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AccountExist = v
	}
	return nil
}

func (p *VerifySocialLoginResponse) ReadField11(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserProfile = &UserProfile{}
	if err := p.UserProfile.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserProfile), err)
	}
	return nil
}

func (p *VerifySocialLoginResponse) ReadField12(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 12: ", err)
	} else {
		p.SameUdidFromAccount = v
	}
	return nil
}

func (p *VerifySocialLoginResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "VerifySocialLoginResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField11(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField12(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifySocialLoginResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountExist", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:accountExist: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.AccountExist)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accountExist (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:accountExist: ", p), err)
	}
	return err
}

func (p *VerifySocialLoginResponse) writeField11(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userProfile", thrift.STRUCT, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:userProfile: ", p), err)
	}
	if err := p.UserProfile.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserProfile), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:userProfile: ", p), err)
	}
	return err
}

func (p *VerifySocialLoginResponse) writeField12(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sameUdidFromAccount", thrift.BOOL, 12); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 12:sameUdidFromAccount: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.SameUdidFromAccount)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sameUdidFromAccount (12) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 12:sameUdidFromAccount: ", p), err)
	}
	return err
}

func (p *VerifySocialLoginResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifySocialLoginResponse(%+v)", *p)
}

// Attributes:
//  - Udid
//  - DeviceModel
type Device struct {
	Udid        string `thrift:"udid,1" db:"udid" json:"udid"`
	DeviceModel string `thrift:"deviceModel,2" db:"deviceModel" json:"deviceModel"`
}

func NewDevice() *Device {
	return &Device{}
}

func (p *Device) GetUdid() string {
	return p.Udid
}

func (p *Device) GetDeviceModel() string {
	return p.DeviceModel
}
func (p *Device) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Device) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Udid = v
	}
	return nil
}

func (p *Device) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.DeviceModel = v
	}
	return nil
}

func (p *Device) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Device"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Device) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "udid", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:udid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Udid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.udid (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:udid: ", p), err)
	}
	return err
}

func (p *Device) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "deviceModel", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:deviceModel: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.DeviceModel)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.deviceModel (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:deviceModel: ", p), err)
	}
	return err
}

func (p *Device) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Device(%+v)", *p)
}

// Attributes:
//  - DeviceModel
//  - Udid
type Device2 struct {
	DeviceModel string `thrift:"deviceModel,1" db:"deviceModel" json:"deviceModel"`
	Udid        string `thrift:"udid,2" db:"udid" json:"udid"`
}

func NewDevice2() *Device2 {
	return &Device2{}
}

func (p *Device2) GetDeviceModel() string {
	return p.DeviceModel
}

func (p *Device2) GetUdid() string {
	return p.Udid
}
func (p *Device2) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Device2) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.DeviceModel = v
	}
	return nil
}

func (p *Device2) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Udid = v
	}
	return nil
}

func (p *Device2) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Device"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Device2) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "deviceModel", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:deviceModel: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.DeviceModel)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.deviceModel (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:deviceModel: ", p), err)
	}
	return err
}

func (p *Device2) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "udid", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:udid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Udid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.udid (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:udid: ", p), err)
	}
	return err
}

func (p *Device2) Equals(other *Device2) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.DeviceModel != other.DeviceModel {
		return false
	}
	if p.Udid != other.Udid {
		return false
	}
	return true
}

func (p *Device2) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Device(%+v)", *p)
}

// Attributes:
//  - Type
//  - AccessToken
//  - CountryCode
type SocialLogin struct {
	Type        SocialLoginType `thrift:"type,1" db:"type" json:"type"`
	AccessToken string          `thrift:"accessToken,2" db:"accessToken" json:"accessToken"`
	CountryCode string          `thrift:"countryCode,3" db:"countryCode" json:"countryCode"`
}

func NewSocialLogin() *SocialLogin {
	return &SocialLogin{}
}

func (p *SocialLogin) GetType() SocialLoginType {
	return p.Type
}

func (p *SocialLogin) GetAccessToken() string {
	return p.AccessToken
}

func (p *SocialLogin) GetCountryCode() string {
	return p.CountryCode
}
func (p *SocialLogin) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SocialLogin) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := SocialLoginType(v)
		p.Type = temp
	}
	return nil
}

func (p *SocialLogin) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AccessToken = v
	}
	return nil
}

func (p *SocialLogin) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.CountryCode = v
	}
	return nil
}

func (p *SocialLogin) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "SocialLogin"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SocialLogin) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "type", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:type: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Type)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.type (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:type: ", p), err)
	}
	return err
}

func (p *SocialLogin) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accessToken", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:accessToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AccessToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accessToken (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:accessToken: ", p), err)
	}
	return err
}

func (p *SocialLogin) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "countryCode", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:countryCode: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.CountryCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.countryCode (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:countryCode: ", p), err)
	}
	return err
}

func (p *SocialLogin) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SocialLogin(%+v)", *p)
}

// Attributes:
//  - AvailableMethods
//  - PrettifiedPhoneNumber
type GetPhoneVerifMethodResponse struct {
	AvailableMethods []VerifMethod `thrift:"availableMethods,1" db:"availableMethods" json:"availableMethods"`
	// unused field # 2
	PrettifiedPhoneNumber string `thrift:"prettifiedPhoneNumber,3" db:"prettifiedPhoneNumber" json:"prettifiedPhoneNumber"`
}

func NewGetPhoneVerifMethodResponse() *GetPhoneVerifMethodResponse {
	return &GetPhoneVerifMethodResponse{}
}

func (p *GetPhoneVerifMethodResponse) GetAvailableMethods() []VerifMethod {
	return p.AvailableMethods
}

func (p *GetPhoneVerifMethodResponse) GetPrettifiedPhoneNumber() string {
	return p.PrettifiedPhoneNumber
}
func (p *GetPhoneVerifMethodResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetPhoneVerifMethodResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]VerifMethod, 0, size)
	p.AvailableMethods = tSlice
	for i := 0; i < size; i++ {
		var _elem2 VerifMethod
		if v, err := iprot.ReadI32(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			temp := VerifMethod(v)
			_elem2 = temp
		}
		p.AvailableMethods = append(p.AvailableMethods, _elem2)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *GetPhoneVerifMethodResponse) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.PrettifiedPhoneNumber = v
	}
	return nil
}

func (p *GetPhoneVerifMethodResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "GetPhoneVerifMethodResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetPhoneVerifMethodResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "availableMethods", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:availableMethods: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.I32, len(p.AvailableMethods)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.AvailableMethods {
		if err := oprot.WriteI32(ctx, int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:availableMethods: ", p), err)
	}
	return err
}

func (p *GetPhoneVerifMethodResponse) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "prettifiedPhoneNumber", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:prettifiedPhoneNumber: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PrettifiedPhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.prettifiedPhoneNumber (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:prettifiedPhoneNumber: ", p), err)
	}
	return err
}

func (p *GetPhoneVerifMethodResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetPhoneVerifMethodResponse(%+v)", *p)
}

// Attributes:
//  - PhoneNumber
//  - CountryCode
type UserPhoneNumber struct {
	PhoneNumber string `thrift:"phoneNumber,1" db:"phoneNumber" json:"phoneNumber"`
	CountryCode string `thrift:"countryCode,2" db:"countryCode" json:"countryCode"`
}

func NewUserPhoneNumber() *UserPhoneNumber {
	return &UserPhoneNumber{}
}

func (p *UserPhoneNumber) GetPhoneNumber() string {
	return p.PhoneNumber
}

func (p *UserPhoneNumber) GetCountryCode() string {
	return p.CountryCode
}
func (p *UserPhoneNumber) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *UserPhoneNumber) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.PhoneNumber = v
	}
	return nil
}

func (p *UserPhoneNumber) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.CountryCode = v
	}
	return nil
}

func (p *UserPhoneNumber) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "UserPhoneNumber"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *UserPhoneNumber) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "phoneNumber", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:phoneNumber: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.phoneNumber (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:phoneNumber: ", p), err)
	}
	return err
}

func (p *UserPhoneNumber) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "countryCode", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:countryCode: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.CountryCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.countryCode (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:countryCode: ", p), err)
	}
	return err
}

func (p *UserPhoneNumber) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserPhoneNumber(%+v)", *p)
}

// Attributes:
//  - AvailableMethods
type SendPinCodeForPhoneResponse struct {
	AvailableMethods []VerifMethod `thrift:"availableMethods,1" db:"availableMethods" json:"availableMethods"`
}

func NewSendPinCodeForPhoneResponse() *SendPinCodeForPhoneResponse {
	return &SendPinCodeForPhoneResponse{}
}

func (p *SendPinCodeForPhoneResponse) GetAvailableMethods() []VerifMethod {
	return p.AvailableMethods
}
func (p *SendPinCodeForPhoneResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SendPinCodeForPhoneResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]VerifMethod, 0, size)
	p.AvailableMethods = tSlice
	for i := 0; i < size; i++ {
		var _elem3 VerifMethod
		if v, err := iprot.ReadI32(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			temp := VerifMethod(v)
			_elem3 = temp
		}
		p.AvailableMethods = append(p.AvailableMethods, _elem3)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *SendPinCodeForPhoneResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "SendPinCodeForPhoneResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SendPinCodeForPhoneResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "availableMethods", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:availableMethods: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.I32, len(p.AvailableMethods)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.AvailableMethods {
		if err := oprot.WriteI32(ctx, int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:availableMethods: ", p), err)
	}
	return err
}

func (p *SendPinCodeForPhoneResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SendPinCodeForPhoneResponse(%+v)", *p)
}

// Attributes:
//  - AccountExist
//  - SameUdidFromAccount
//  - UserProfile
type VerifyPhoneResponse struct {
	// unused field # 1
	AccountExist        bool `thrift:"accountExist,2" db:"accountExist" json:"accountExist"`
	SameUdidFromAccount bool `thrift:"sameUdidFromAccount,3" db:"sameUdidFromAccount" json:"sameUdidFromAccount"`
	// unused fields # 4 to 10
	UserProfile *UserProfile `thrift:"userProfile,11" db:"userProfile" json:"userProfile"`
}

func NewVerifyPhoneResponse() *VerifyPhoneResponse {
	return &VerifyPhoneResponse{}
}

func (p *VerifyPhoneResponse) GetAccountExist() bool {
	return p.AccountExist
}

func (p *VerifyPhoneResponse) GetSameUdidFromAccount() bool {
	return p.SameUdidFromAccount
}

var VerifyPhoneResponse_UserProfile_DEFAULT *UserProfile

func (p *VerifyPhoneResponse) GetUserProfile() *UserProfile {
	if !p.IsSetUserProfile() {
		return VerifyPhoneResponse_UserProfile_DEFAULT
	}
	return p.UserProfile
}
func (p *VerifyPhoneResponse) IsSetUserProfile() bool {
	return p.UserProfile != nil
}

func (p *VerifyPhoneResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 11:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField11(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *VerifyPhoneResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AccountExist = v
	}
	return nil
}

func (p *VerifyPhoneResponse) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.SameUdidFromAccount = v
	}
	return nil
}

func (p *VerifyPhoneResponse) ReadField11(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserProfile = &UserProfile{}
	if err := p.UserProfile.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserProfile), err)
	}
	return nil
}

func (p *VerifyPhoneResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "VerifyPhoneResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField11(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifyPhoneResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountExist", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:accountExist: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.AccountExist)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accountExist (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:accountExist: ", p), err)
	}
	return err
}

func (p *VerifyPhoneResponse) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sameUdidFromAccount", thrift.BOOL, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:sameUdidFromAccount: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.SameUdidFromAccount)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sameUdidFromAccount (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:sameUdidFromAccount: ", p), err)
	}
	return err
}

func (p *VerifyPhoneResponse) writeField11(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userProfile", thrift.STRUCT, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:userProfile: ", p), err)
	}
	if err := p.UserProfile.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserProfile), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:userProfile: ", p), err)
	}
	return err
}

func (p *VerifyPhoneResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifyPhoneResponse(%+v)", *p)
}

// Attributes:
//  - PublicKey
//  - Nonce
type ExchangeEncryptionKeyResponse struct {
	PublicKey string `thrift:"publicKey,1" db:"publicKey" json:"publicKey"`
	Nonce     string `thrift:"nonce,2" db:"nonce" json:"nonce"`
}

func NewExchangeEncryptionKeyResponse() *ExchangeEncryptionKeyResponse {
	return &ExchangeEncryptionKeyResponse{}
}

func (p *ExchangeEncryptionKeyResponse) GetPublicKey() string {
	return p.PublicKey
}

func (p *ExchangeEncryptionKeyResponse) GetNonce() string {
	return p.Nonce
}
func (p *ExchangeEncryptionKeyResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ExchangeEncryptionKeyResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.PublicKey = v
	}
	return nil
}

func (p *ExchangeEncryptionKeyResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Nonce = v
	}
	return nil
}

func (p *ExchangeEncryptionKeyResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ExchangeEncryptionKeyResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ExchangeEncryptionKeyResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "publicKey", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:publicKey: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PublicKey)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.publicKey (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:publicKey: ", p), err)
	}
	return err
}

func (p *ExchangeEncryptionKeyResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "nonce", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:nonce: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Nonce)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nonce (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:nonce: ", p), err)
	}
	return err
}

func (p *ExchangeEncryptionKeyResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ExchangeEncryptionKeyResponse(%+v)", *p)
}

// Attributes:
//  - AuthKeyVersion
//  - PublicKey
//  - Nonce
type ExchangeEncryptionKeyRequest struct {
	AuthKeyVersion EncryptionKeyVersion `thrift:"authKeyVersion,1" db:"authKeyVersion" json:"authKeyVersion"`
	PublicKey      string               `thrift:"publicKey,2" db:"publicKey" json:"publicKey"`
	Nonce          string               `thrift:"nonce,3" db:"nonce" json:"nonce"`
}

func NewExchangeEncryptionKeyRequest() *ExchangeEncryptionKeyRequest {
	return &ExchangeEncryptionKeyRequest{}
}

func (p *ExchangeEncryptionKeyRequest) GetAuthKeyVersion() EncryptionKeyVersion {
	return p.AuthKeyVersion
}

func (p *ExchangeEncryptionKeyRequest) GetPublicKey() string {
	return p.PublicKey
}

func (p *ExchangeEncryptionKeyRequest) GetNonce() string {
	return p.Nonce
}
func (p *ExchangeEncryptionKeyRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ExchangeEncryptionKeyRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := EncryptionKeyVersion(v)
		p.AuthKeyVersion = temp
	}
	return nil
}

func (p *ExchangeEncryptionKeyRequest) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.PublicKey = v
	}
	return nil
}

func (p *ExchangeEncryptionKeyRequest) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Nonce = v
	}
	return nil
}

func (p *ExchangeEncryptionKeyRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ExchangeEncryptionKeyRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ExchangeEncryptionKeyRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authKeyVersion", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authKeyVersion: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.AuthKeyVersion)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authKeyVersion (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authKeyVersion: ", p), err)
	}
	return err
}

func (p *ExchangeEncryptionKeyRequest) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "publicKey", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:publicKey: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PublicKey)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.publicKey (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:publicKey: ", p), err)
	}
	return err
}

func (p *ExchangeEncryptionKeyRequest) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "nonce", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:nonce: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Nonce)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nonce (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:nonce: ", p), err)
	}
	return err
}

func (p *ExchangeEncryptionKeyRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ExchangeEncryptionKeyRequest(%+v)", *p)
}

// Attributes:
//  - CountryCode
//  - CountryInEEA
//  - CountrySetOfEEA
type GetCountryInfoResponse struct {
	CountryCode     string   `thrift:"countryCode,1" db:"countryCode" json:"countryCode"`
	CountryInEEA    bool     `thrift:"countryInEEA,2" db:"countryInEEA" json:"countryInEEA"`
	CountrySetOfEEA []string `thrift:"countrySetOfEEA,3" db:"countrySetOfEEA" json:"countrySetOfEEA"`
}

func NewGetCountryInfoResponse() *GetCountryInfoResponse {
	return &GetCountryInfoResponse{}
}

func (p *GetCountryInfoResponse) GetCountryCode() string {
	return p.CountryCode
}

func (p *GetCountryInfoResponse) GetCountryInEEA() bool {
	return p.CountryInEEA
}

func (p *GetCountryInfoResponse) GetCountrySetOfEEA() []string {
	return p.CountrySetOfEEA
}
func (p *GetCountryInfoResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.SET {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetCountryInfoResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.CountryCode = v
	}
	return nil
}

func (p *GetCountryInfoResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.CountryInEEA = v
	}
	return nil
}

func (p *GetCountryInfoResponse) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadSetBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading set begin: ", err)
	}
	tSet := make([]string, 0, size)
	p.CountrySetOfEEA = tSet
	for i := 0; i < size; i++ {
		var _elem4 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem4 = v
		}
		p.CountrySetOfEEA = append(p.CountrySetOfEEA, _elem4)
	}
	if err := iprot.ReadSetEnd(ctx); err != nil {
		return thrift.PrependError("error reading set end: ", err)
	}
	return nil
}

func (p *GetCountryInfoResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "GetCountryInfoResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetCountryInfoResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "countryCode", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:countryCode: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.CountryCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.countryCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:countryCode: ", p), err)
	}
	return err
}

func (p *GetCountryInfoResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "countryInEEA", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:countryInEEA: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.CountryInEEA)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.countryInEEA (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:countryInEEA: ", p), err)
	}
	return err
}

func (p *GetCountryInfoResponse) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "countrySetOfEEA", thrift.SET, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:countrySetOfEEA: ", p), err)
	}
	if err := oprot.WriteSetBegin(ctx, thrift.STRING, len(p.CountrySetOfEEA)); err != nil {
		return thrift.PrependError("error writing set begin: ", err)
	}
	for i := 0; i < len(p.CountrySetOfEEA); i++ {
		for j := i + 1; j < len(p.CountrySetOfEEA); j++ {
			if reflect.DeepEqual(p.CountrySetOfEEA[i], p.CountrySetOfEEA[j]) {
				return thrift.PrependError("", fmt.Errorf("%T error writing set field: slice is not unique", p.CountrySetOfEEA[i]))
			}
		}
	}
	for _, v := range p.CountrySetOfEEA {
		if err := oprot.WriteString(ctx, string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteSetEnd(ctx); err != nil {
		return thrift.PrependError("error writing set end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:countrySetOfEEA: ", p), err)
	}
	return err
}

func (p *GetCountryInfoResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetCountryInfoResponse(%+v)", *p)
}

// Attributes:
//  - CountryCode
//  - Hni
//  - CarrierName
type SimCard struct {
	CountryCode string `thrift:"countryCode,1" db:"countryCode" json:"countryCode"`
	Hni         string `thrift:"hni,2" db:"hni" json:"hni"`
	CarrierName string `thrift:"carrierName,3" db:"carrierName" json:"carrierName"`
}

func NewSimCard() *SimCard {
	return &SimCard{}
}

func (p *SimCard) GetCountryCode() string {
	return p.CountryCode
}

func (p *SimCard) GetHni() string {
	return p.Hni
}

func (p *SimCard) GetCarrierName() string {
	return p.CarrierName
}
func (p *SimCard) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SimCard) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.CountryCode = v
	}
	return nil
}

func (p *SimCard) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Hni = v
	}
	return nil
}

func (p *SimCard) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.CarrierName = v
	}
	return nil
}

func (p *SimCard) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "SimCard"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SimCard) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "countryCode", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:countryCode: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.CountryCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.countryCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:countryCode: ", p), err)
	}
	return err
}

func (p *SimCard) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "hni", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:hni: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Hni)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.hni (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:hni: ", p), err)
	}
	return err
}

func (p *SimCard) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "carrierName", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:carrierName: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.CarrierName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.carrierName (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:carrierName: ", p), err)
	}
	return err
}

func (p *SimCard) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SimCard(%+v)", *p)
}

// Attributes:
//  - UserProfile
type GetUserProfileResponse struct {
	UserProfile *UserProfile `thrift:"userProfile,1" db:"userProfile" json:"userProfile"`
}

func NewGetUserProfileResponse() *GetUserProfileResponse {
	return &GetUserProfileResponse{}
}

var GetUserProfileResponse_UserProfile_DEFAULT *UserProfile

func (p *GetUserProfileResponse) GetUserProfile() *UserProfile {
	if !p.IsSetUserProfile() {
		return GetUserProfileResponse_UserProfile_DEFAULT
	}
	return p.UserProfile
}
func (p *GetUserProfileResponse) IsSetUserProfile() bool {
	return p.UserProfile != nil
}

func (p *GetUserProfileResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetUserProfileResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserProfile = &UserProfile{}
	if err := p.UserProfile.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserProfile), err)
	}
	return nil
}

func (p *GetUserProfileResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "GetUserProfileResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetUserProfileResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userProfile", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:userProfile: ", p), err)
	}
	if err := p.UserProfile.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserProfile), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:userProfile: ", p), err)
	}
	return err
}

func (p *GetUserProfileResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetUserProfileResponse(%+v)", *p)
}

// Attributes:
//  - AvailableMethod
//  - SameAccountFromAuthFactor
type GetAcctVerifMethodResponse struct {
	AvailableMethod           AvailableMethod `thrift:"availableMethod,1" db:"availableMethod" json:"availableMethod"`
	SameAccountFromAuthFactor bool            `thrift:"sameAccountFromAuthFactor,2" db:"sameAccountFromAuthFactor" json:"sameAccountFromAuthFactor"`
}

func NewGetAcctVerifMethodResponse() *GetAcctVerifMethodResponse {
	return &GetAcctVerifMethodResponse{}
}

func (p *GetAcctVerifMethodResponse) GetAvailableMethod() AvailableMethod {
	return p.AvailableMethod
}

func (p *GetAcctVerifMethodResponse) GetSameAccountFromAuthFactor() bool {
	return p.SameAccountFromAuthFactor
}
func (p *GetAcctVerifMethodResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetAcctVerifMethodResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := AvailableMethod(v)
		p.AvailableMethod = temp
	}
	return nil
}

func (p *GetAcctVerifMethodResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.SameAccountFromAuthFactor = v
	}
	return nil
}

func (p *GetAcctVerifMethodResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "GetAcctVerifMethodResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetAcctVerifMethodResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "availableMethod", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:availableMethod: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.AvailableMethod)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.availableMethod (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:availableMethod: ", p), err)
	}
	return err
}

func (p *GetAcctVerifMethodResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sameAccountFromAuthFactor", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:sameAccountFromAuthFactor: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.SameAccountFromAuthFactor)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sameAccountFromAuthFactor (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:sameAccountFromAuthFactor: ", p), err)
	}
	return err
}

func (p *GetAcctVerifMethodResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetAcctVerifMethodResponse(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - UserPhoneNumber
type FetchPhonePinCodeMsgRequest struct {
	AuthSessionId   string           `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	UserPhoneNumber *UserPhoneNumber `thrift:"userPhoneNumber,2" db:"userPhoneNumber" json:"userPhoneNumber"`
}

func NewFetchPhonePinCodeMsgRequest() *FetchPhonePinCodeMsgRequest {
	return &FetchPhonePinCodeMsgRequest{}
}

func (p *FetchPhonePinCodeMsgRequest) GetAuthSessionId() string {
	return p.AuthSessionId
}

var FetchPhonePinCodeMsgRequest_UserPhoneNumber_DEFAULT *UserPhoneNumber

func (p *FetchPhonePinCodeMsgRequest) GetUserPhoneNumber() *UserPhoneNumber {
	if !p.IsSetUserPhoneNumber() {
		return FetchPhonePinCodeMsgRequest_UserPhoneNumber_DEFAULT
	}
	return p.UserPhoneNumber
}
func (p *FetchPhonePinCodeMsgRequest) IsSetUserPhoneNumber() bool {
	return p.UserPhoneNumber != nil
}

func (p *FetchPhonePinCodeMsgRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FetchPhonePinCodeMsgRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *FetchPhonePinCodeMsgRequest) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserPhoneNumber = &UserPhoneNumber{}
	if err := p.UserPhoneNumber.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserPhoneNumber), err)
	}
	return nil
}

func (p *FetchPhonePinCodeMsgRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "FetchPhonePinCodeMsgRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FetchPhonePinCodeMsgRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *FetchPhonePinCodeMsgRequest) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userPhoneNumber", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:userPhoneNumber: ", p), err)
	}
	if err := p.UserPhoneNumber.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserPhoneNumber), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:userPhoneNumber: ", p), err)
	}
	return err
}

func (p *FetchPhonePinCodeMsgRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FetchPhonePinCodeMsgRequest(%+v)", *p)
}

// Attributes:
//  - PinCodeMessage
//  - DestinationPhoneNumber
type FetchPhonePinCodeMsgResponse struct {
	PinCodeMessage         string `thrift:"pinCodeMessage,1" db:"pinCodeMessage" json:"pinCodeMessage"`
	DestinationPhoneNumber string `thrift:"destinationPhoneNumber,2" db:"destinationPhoneNumber" json:"destinationPhoneNumber"`
}

func NewFetchPhonePinCodeMsgResponse() *FetchPhonePinCodeMsgResponse {
	return &FetchPhonePinCodeMsgResponse{}
}

func (p *FetchPhonePinCodeMsgResponse) GetPinCodeMessage() string {
	return p.PinCodeMessage
}

func (p *FetchPhonePinCodeMsgResponse) GetDestinationPhoneNumber() string {
	return p.DestinationPhoneNumber
}
func (p *FetchPhonePinCodeMsgResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FetchPhonePinCodeMsgResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.PinCodeMessage = v
	}
	return nil
}

func (p *FetchPhonePinCodeMsgResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.DestinationPhoneNumber = v
	}
	return nil
}

func (p *FetchPhonePinCodeMsgResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "FetchPhonePinCodeMsgResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FetchPhonePinCodeMsgResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "pinCodeMessage", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pinCodeMessage: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PinCodeMessage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.pinCodeMessage (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pinCodeMessage: ", p), err)
	}
	return err
}

func (p *FetchPhonePinCodeMsgResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "destinationPhoneNumber", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:destinationPhoneNumber: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.DestinationPhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.destinationPhoneNumber (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:destinationPhoneNumber: ", p), err)
	}
	return err
}

func (p *FetchPhonePinCodeMsgResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FetchPhonePinCodeMsgResponse(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - Device
//  - UserPhoneNumber
type GetPhoneVerifMethodV2Request struct {
	AuthSessionId   string           `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	Device          *Device          `thrift:"device,2" db:"device" json:"device"`
	UserPhoneNumber *UserPhoneNumber `thrift:"userPhoneNumber,3" db:"userPhoneNumber" json:"userPhoneNumber"`
}

func NewGetPhoneVerifMethodV2Request() *GetPhoneVerifMethodV2Request {
	return &GetPhoneVerifMethodV2Request{}
}

func (p *GetPhoneVerifMethodV2Request) GetAuthSessionId() string {
	return p.AuthSessionId
}

var GetPhoneVerifMethodV2Request_Device_DEFAULT *Device

func (p *GetPhoneVerifMethodV2Request) GetDevice() *Device {
	if !p.IsSetDevice() {
		return GetPhoneVerifMethodV2Request_Device_DEFAULT
	}
	return p.Device
}

var GetPhoneVerifMethodV2Request_UserPhoneNumber_DEFAULT *UserPhoneNumber

func (p *GetPhoneVerifMethodV2Request) GetUserPhoneNumber() *UserPhoneNumber {
	if !p.IsSetUserPhoneNumber() {
		return GetPhoneVerifMethodV2Request_UserPhoneNumber_DEFAULT
	}
	return p.UserPhoneNumber
}
func (p *GetPhoneVerifMethodV2Request) IsSetDevice() bool {
	return p.Device != nil
}

func (p *GetPhoneVerifMethodV2Request) IsSetUserPhoneNumber() bool {
	return p.UserPhoneNumber != nil
}

func (p *GetPhoneVerifMethodV2Request) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Request) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Request) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Device = &Device{}
	if err := p.Device.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Device), err)
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Request) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserPhoneNumber = &UserPhoneNumber{}
	if err := p.UserPhoneNumber.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserPhoneNumber), err)
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Request) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "GetPhoneVerifMethodV2Request"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Request) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *GetPhoneVerifMethodV2Request) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "device", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:device: ", p), err)
	}
	if err := p.Device.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Device), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:device: ", p), err)
	}
	return err
}

func (p *GetPhoneVerifMethodV2Request) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userPhoneNumber", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:userPhoneNumber: ", p), err)
	}
	if err := p.UserPhoneNumber.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserPhoneNumber), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:userPhoneNumber: ", p), err)
	}
	return err
}

func (p *GetPhoneVerifMethodV2Request) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetPhoneVerifMethodV2Request(%+v)", *p)
}

// Attributes:
//  - AvailableMethods
//  - PrettifiedPhoneNumber
type GetPhoneVerifMethodV2Response struct {
	AvailableMethods []AvailableMethods `thrift:"availableMethods,1" db:"availableMethods" json:"availableMethods"`
	// unused field # 2
	PrettifiedPhoneNumber string `thrift:"prettifiedPhoneNumber,3" db:"prettifiedPhoneNumber" json:"prettifiedPhoneNumber"`
}

func NewGetPhoneVerifMethodV2Response() *GetPhoneVerifMethodV2Response {
	return &GetPhoneVerifMethodV2Response{}
}

func (p *GetPhoneVerifMethodV2Response) GetAvailableMethods() []AvailableMethods {
	return p.AvailableMethods
}

func (p *GetPhoneVerifMethodV2Response) GetPrettifiedPhoneNumber() string {
	return p.PrettifiedPhoneNumber
}
func (p *GetPhoneVerifMethodV2Response) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Response) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]AvailableMethods, 0, size)
	p.AvailableMethods = tSlice
	for i := 0; i < size; i++ {
		var _elem1 AvailableMethods
		if v, err := iprot.ReadI32(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			temp := AvailableMethods(v)
			_elem1 = temp
		}
		p.AvailableMethods = append(p.AvailableMethods, _elem1)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Response) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.PrettifiedPhoneNumber = v
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Response) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "GetPhoneVerifMethodV2Response"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetPhoneVerifMethodV2Response) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "availableMethods", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:availableMethods: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.I32, len(p.AvailableMethods)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.AvailableMethods {
		if err := oprot.WriteI32(ctx, int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:availableMethods: ", p), err)
	}
	return err
}

func (p *GetPhoneVerifMethodV2Response) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "prettifiedPhoneNumber", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:prettifiedPhoneNumber: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PrettifiedPhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.prettifiedPhoneNumber (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:prettifiedPhoneNumber: ", p), err)
	}
	return err
}

func (p *GetPhoneVerifMethodV2Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetPhoneVerifMethodV2Response(%+v)", *p)
}

// Attributes:
//  - HasLastKey
//  - StartTimeMillis
//  - EndTimeMillis
type E2EEKeyBackupInfoForMigration struct {
	HasLastKey      bool  `thrift:"hasLastKey,1" db:"hasLastKey" json:"hasLastKey"`
	StartTimeMillis int64 `thrift:"startTimeMillis,2" db:"startTimeMillis" json:"startTimeMillis"`
	EndTimeMillis   int64 `thrift:"endTimeMillis,3" db:"endTimeMillis" json:"endTimeMillis"`
}

func NewE2EEKeyBackupInfoForMigration() *E2EEKeyBackupInfoForMigration {
	return &E2EEKeyBackupInfoForMigration{}
}

func (p *E2EEKeyBackupInfoForMigration) GetHasLastKey() bool {
	return p.HasLastKey
}

func (p *E2EEKeyBackupInfoForMigration) GetStartTimeMillis() int64 {
	return p.StartTimeMillis
}

func (p *E2EEKeyBackupInfoForMigration) GetEndTimeMillis() int64 {
	return p.EndTimeMillis
}
func (p *E2EEKeyBackupInfoForMigration) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *E2EEKeyBackupInfoForMigration) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.HasLastKey = v
	}
	return nil
}

func (p *E2EEKeyBackupInfoForMigration) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.StartTimeMillis = v
	}
	return nil
}

func (p *E2EEKeyBackupInfoForMigration) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.EndTimeMillis = v
	}
	return nil
}

func (p *E2EEKeyBackupInfoForMigration) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "E2EEKeyBackupInfoForMigration"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *E2EEKeyBackupInfoForMigration) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "hasLastKey", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:hasLastKey: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.HasLastKey)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.hasLastKey (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:hasLastKey: ", p), err)
	}
	return err
}

func (p *E2EEKeyBackupInfoForMigration) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "startTimeMillis", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:startTimeMillis: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.StartTimeMillis)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.startTimeMillis (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:startTimeMillis: ", p), err)
	}
	return err
}

func (p *E2EEKeyBackupInfoForMigration) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "endTimeMillis", thrift.I64, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:endTimeMillis: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.EndTimeMillis)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.endTimeMillis (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:endTimeMillis: ", p), err)
	}
	return err
}

func (p *E2EEKeyBackupInfoForMigration) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("E2EEKeyBackupInfoForMigration(%+v)", *p)
}

// Attributes:
//  - AppTypeDifferentFromPrevDevice
//  - E2eeKeyBackupServiceConfig
//  - E2eeKeyBackupInfo
type GetSessionContentBeforeMigCompletionResponse struct {
	AppTypeDifferentFromPrevDevice bool                           `thrift:"appTypeDifferentFromPrevDevice,1" db:"appTypeDifferentFromPrevDevice" json:"appTypeDifferentFromPrevDevice"`
	E2eeKeyBackupServiceConfig     bool                           `thrift:"e2eeKeyBackupServiceConfig,2" db:"e2eeKeyBackupServiceConfig" json:"e2eeKeyBackupServiceConfig"`
	E2eeKeyBackupInfo              *E2EEKeyBackupInfoForMigration `thrift:"e2eeKeyBackupInfo,3" db:"e2eeKeyBackupInfo" json:"e2eeKeyBackupInfo"`
}

func NewGetSessionContentBeforeMigCompletionResponse() *GetSessionContentBeforeMigCompletionResponse {
	return &GetSessionContentBeforeMigCompletionResponse{}
}

func (p *GetSessionContentBeforeMigCompletionResponse) GetAppTypeDifferentFromPrevDevice() bool {
	return p.AppTypeDifferentFromPrevDevice
}

func (p *GetSessionContentBeforeMigCompletionResponse) GetE2eeKeyBackupServiceConfig() bool {
	return p.E2eeKeyBackupServiceConfig
}

var GetSessionContentBeforeMigCompletionResponse_E2eeKeyBackupInfo_DEFAULT *E2EEKeyBackupInfoForMigration

func (p *GetSessionContentBeforeMigCompletionResponse) GetE2eeKeyBackupInfo() *E2EEKeyBackupInfoForMigration {
	if !p.IsSetE2eeKeyBackupInfo() {
		return GetSessionContentBeforeMigCompletionResponse_E2eeKeyBackupInfo_DEFAULT
	}
	return p.E2eeKeyBackupInfo
}
func (p *GetSessionContentBeforeMigCompletionResponse) IsSetE2eeKeyBackupInfo() bool {
	return p.E2eeKeyBackupInfo != nil
}

func (p *GetSessionContentBeforeMigCompletionResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *GetSessionContentBeforeMigCompletionResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AppTypeDifferentFromPrevDevice = v
	}
	return nil
}

func (p *GetSessionContentBeforeMigCompletionResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.E2eeKeyBackupServiceConfig = v
	}
	return nil
}

func (p *GetSessionContentBeforeMigCompletionResponse) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.E2eeKeyBackupInfo = &E2EEKeyBackupInfoForMigration{}
	if err := p.E2eeKeyBackupInfo.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E2eeKeyBackupInfo), err)
	}
	return nil
}

func (p *GetSessionContentBeforeMigCompletionResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "GetSessionContentBeforeMigCompletionResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *GetSessionContentBeforeMigCompletionResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "appTypeDifferentFromPrevDevice", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:appTypeDifferentFromPrevDevice: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.AppTypeDifferentFromPrevDevice)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.appTypeDifferentFromPrevDevice (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:appTypeDifferentFromPrevDevice: ", p), err)
	}
	return err
}

func (p *GetSessionContentBeforeMigCompletionResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "e2eeKeyBackupServiceConfig", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:e2eeKeyBackupServiceConfig: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.E2eeKeyBackupServiceConfig)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.e2eeKeyBackupServiceConfig (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:e2eeKeyBackupServiceConfig: ", p), err)
	}
	return err
}

func (p *GetSessionContentBeforeMigCompletionResponse) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "e2eeKeyBackupInfo", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:e2eeKeyBackupInfo: ", p), err)
	}
	if err := p.E2eeKeyBackupInfo.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E2eeKeyBackupInfo), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:e2eeKeyBackupInfo: ", p), err)
	}
	return err
}

func (p *GetSessionContentBeforeMigCompletionResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetSessionContentBeforeMigCompletionResponse(%+v)", *p)
}

// Attributes:
//  - InitialDelayInMillis
//  - MaxDelayInMillis
//  - Multiplier
//  - JitterRate
type RefreshApiRetryPolicy struct {
	InitialDelayInMillis int64   `thrift:"initialDelayInMillis,1" db:"initialDelayInMillis" json:"initialDelayInMillis"`
	MaxDelayInMillis     int64   `thrift:"maxDelayInMillis,2" db:"maxDelayInMillis" json:"maxDelayInMillis"`
	Multiplier           float64 `thrift:"multiplier,3" db:"multiplier" json:"multiplier"`
	JitterRate           float64 `thrift:"jitterRate,4" db:"jitterRate" json:"jitterRate"`
}

func NewRefreshApiRetryPolicy() *RefreshApiRetryPolicy {
	return &RefreshApiRetryPolicy{}
}

func (p *RefreshApiRetryPolicy) GetInitialDelayInMillis() int64 {
	return p.InitialDelayInMillis
}

func (p *RefreshApiRetryPolicy) GetMaxDelayInMillis() int64 {
	return p.MaxDelayInMillis
}

func (p *RefreshApiRetryPolicy) GetMultiplier() float64 {
	return p.Multiplier
}

func (p *RefreshApiRetryPolicy) GetJitterRate() float64 {
	return p.JitterRate
}
func (p *RefreshApiRetryPolicy) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.DOUBLE {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.DOUBLE {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RefreshApiRetryPolicy) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.InitialDelayInMillis = v
	}
	return nil
}

func (p *RefreshApiRetryPolicy) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.MaxDelayInMillis = v
	}
	return nil
}

func (p *RefreshApiRetryPolicy) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Multiplier = v
	}
	return nil
}

func (p *RefreshApiRetryPolicy) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadDouble(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.JitterRate = v
	}
	return nil
}

func (p *RefreshApiRetryPolicy) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "RefreshApiRetryPolicy"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RefreshApiRetryPolicy) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "initialDelayInMillis", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:initialDelayInMillis: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.InitialDelayInMillis)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.initialDelayInMillis (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:initialDelayInMillis: ", p), err)
	}
	return err
}

func (p *RefreshApiRetryPolicy) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "maxDelayInMillis", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:maxDelayInMillis: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.MaxDelayInMillis)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.maxDelayInMillis (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:maxDelayInMillis: ", p), err)
	}
	return err
}

func (p *RefreshApiRetryPolicy) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "multiplier", thrift.DOUBLE, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:multiplier: ", p), err)
	}
	if err := oprot.WriteDouble(ctx, float64(p.Multiplier)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.multiplier (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:multiplier: ", p), err)
	}
	return err
}

func (p *RefreshApiRetryPolicy) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "jitterRate", thrift.DOUBLE, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:jitterRate: ", p), err)
	}
	if err := oprot.WriteDouble(ctx, float64(p.JitterRate)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.jitterRate (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:jitterRate: ", p), err)
	}
	return err
}

func (p *RefreshApiRetryPolicy) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefreshApiRetryPolicy(%+v)", *p)
}

// Attributes:
//  - AccessToken
//  - RefreshToken
//  - DurationUntilRefreshInSec
//  - RefreshApiRetryPolicy
//  - LoginSessionId
//  - TokenIssueTimeEpochSec
type TokenV3IssueResult_ struct {
	AccessToken               string                 `thrift:"accessToken,1" db:"accessToken" json:"accessToken"`
	RefreshToken              string                 `thrift:"refreshToken,2" db:"refreshToken" json:"refreshToken"`
	DurationUntilRefreshInSec int64                  `thrift:"durationUntilRefreshInSec,3" db:"durationUntilRefreshInSec" json:"durationUntilRefreshInSec"`
	RefreshApiRetryPolicy     *RefreshApiRetryPolicy `thrift:"refreshApiRetryPolicy,4" db:"refreshApiRetryPolicy" json:"refreshApiRetryPolicy"`
	LoginSessionId            string                 `thrift:"loginSessionId,5" db:"loginSessionId" json:"loginSessionId"`
	TokenIssueTimeEpochSec    int64                  `thrift:"tokenIssueTimeEpochSec,6" db:"tokenIssueTimeEpochSec" json:"tokenIssueTimeEpochSec"`
}

func NewTokenV3IssueResult_() *TokenV3IssueResult_ {
	return &TokenV3IssueResult_{}
}

func (p *TokenV3IssueResult_) GetAccessToken() string {
	return p.AccessToken
}

func (p *TokenV3IssueResult_) GetRefreshToken() string {
	return p.RefreshToken
}

func (p *TokenV3IssueResult_) GetDurationUntilRefreshInSec() int64 {
	return p.DurationUntilRefreshInSec
}

var TokenV3IssueResult__RefreshApiRetryPolicy_DEFAULT *RefreshApiRetryPolicy

func (p *TokenV3IssueResult_) GetRefreshApiRetryPolicy() *RefreshApiRetryPolicy {
	if !p.IsSetRefreshApiRetryPolicy() {
		return TokenV3IssueResult__RefreshApiRetryPolicy_DEFAULT
	}
	return p.RefreshApiRetryPolicy
}

func (p *TokenV3IssueResult_) GetLoginSessionId() string {
	return p.LoginSessionId
}

func (p *TokenV3IssueResult_) GetTokenIssueTimeEpochSec() int64 {
	return p.TokenIssueTimeEpochSec
}
func (p *TokenV3IssueResult_) IsSetRefreshApiRetryPolicy() bool {
	return p.RefreshApiRetryPolicy != nil
}

func (p *TokenV3IssueResult_) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField5(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField6(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenV3IssueResult_) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AccessToken = v
	}
	return nil
}

func (p *TokenV3IssueResult_) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.RefreshToken = v
	}
	return nil
}

func (p *TokenV3IssueResult_) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.DurationUntilRefreshInSec = v
	}
	return nil
}

func (p *TokenV3IssueResult_) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	p.RefreshApiRetryPolicy = &RefreshApiRetryPolicy{}
	if err := p.RefreshApiRetryPolicy.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RefreshApiRetryPolicy), err)
	}
	return nil
}

func (p *TokenV3IssueResult_) ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.LoginSessionId = v
	}
	return nil
}

func (p *TokenV3IssueResult_) ReadField6(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.TokenIssueTimeEpochSec = v
	}
	return nil
}

func (p *TokenV3IssueResult_) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "TokenV3IssueResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField5(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField6(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenV3IssueResult_) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accessToken", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:accessToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AccessToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accessToken (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:accessToken: ", p), err)
	}
	return err
}

func (p *TokenV3IssueResult_) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "refreshToken", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:refreshToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.RefreshToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.refreshToken (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:refreshToken: ", p), err)
	}
	return err
}

func (p *TokenV3IssueResult_) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "durationUntilRefreshInSec", thrift.I64, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:durationUntilRefreshInSec: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.DurationUntilRefreshInSec)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.durationUntilRefreshInSec (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:durationUntilRefreshInSec: ", p), err)
	}
	return err
}

func (p *TokenV3IssueResult_) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "refreshApiRetryPolicy", thrift.STRUCT, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:refreshApiRetryPolicy: ", p), err)
	}
	if err := p.RefreshApiRetryPolicy.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RefreshApiRetryPolicy), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:refreshApiRetryPolicy: ", p), err)
	}
	return err
}

func (p *TokenV3IssueResult_) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "loginSessionId", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:loginSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.LoginSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.loginSessionId (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:loginSessionId: ", p), err)
	}
	return err
}

func (p *TokenV3IssueResult_) writeField6(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "tokenIssueTimeEpochSec", thrift.I64, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:tokenIssueTimeEpochSec: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.TokenIssueTimeEpochSec)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.tokenIssueTimeEpochSec (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:tokenIssueTimeEpochSec: ", p), err)
	}
	return err
}

func (p *TokenV3IssueResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenV3IssueResult_(%+v)", *p)
}

// Attributes:
//  - AuthToken
//  - TokenV3IssueResult_
//  - CountryCode
//  - PrettifiedFormatPhoneNumber
//  - LocalFormatPhoneNumber
//  - Mid
type MigratePrimaryWithTokenV3Response struct {
	AuthToken                   string               `thrift:"authToken,1" db:"authToken" json:"authToken"`
	TokenV3IssueResult_         *TokenV3IssueResult_ `thrift:"tokenV3IssueResult,2" db:"tokenV3IssueResult" json:"tokenV3IssueResult"`
	CountryCode                 string               `thrift:"countryCode,3" db:"countryCode" json:"countryCode"`
	PrettifiedFormatPhoneNumber string               `thrift:"prettifiedFormatPhoneNumber,4" db:"prettifiedFormatPhoneNumber" json:"prettifiedFormatPhoneNumber"`
	LocalFormatPhoneNumber      string               `thrift:"localFormatPhoneNumber,5" db:"localFormatPhoneNumber" json:"localFormatPhoneNumber"`
	Mid                         string               `thrift:"mid,6" db:"mid" json:"mid"`
}

func NewMigratePrimaryWithTokenV3Response() *MigratePrimaryWithTokenV3Response {
	return &MigratePrimaryWithTokenV3Response{}
}

func (p *MigratePrimaryWithTokenV3Response) GetAuthToken() string {
	return p.AuthToken
}

var MigratePrimaryWithTokenV3Response_TokenV3IssueResult__DEFAULT *TokenV3IssueResult_

func (p *MigratePrimaryWithTokenV3Response) GetTokenV3IssueResult_() *TokenV3IssueResult_ {
	if !p.IsSetTokenV3IssueResult_() {
		return MigratePrimaryWithTokenV3Response_TokenV3IssueResult__DEFAULT
	}
	return p.TokenV3IssueResult_
}

func (p *MigratePrimaryWithTokenV3Response) GetCountryCode() string {
	return p.CountryCode
}

func (p *MigratePrimaryWithTokenV3Response) GetPrettifiedFormatPhoneNumber() string {
	return p.PrettifiedFormatPhoneNumber
}

func (p *MigratePrimaryWithTokenV3Response) GetLocalFormatPhoneNumber() string {
	return p.LocalFormatPhoneNumber
}

func (p *MigratePrimaryWithTokenV3Response) GetMid() string {
	return p.Mid
}
func (p *MigratePrimaryWithTokenV3Response) IsSetTokenV3IssueResult_() bool {
	return p.TokenV3IssueResult_ != nil
}

func (p *MigratePrimaryWithTokenV3Response) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField5(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 6:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField6(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MigratePrimaryWithTokenV3Response) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthToken = v
	}
	return nil
}

func (p *MigratePrimaryWithTokenV3Response) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.TokenV3IssueResult_ = &TokenV3IssueResult_{}
	if err := p.TokenV3IssueResult_.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TokenV3IssueResult_), err)
	}
	return nil
}

func (p *MigratePrimaryWithTokenV3Response) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.CountryCode = v
	}
	return nil
}

func (p *MigratePrimaryWithTokenV3Response) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.PrettifiedFormatPhoneNumber = v
	}
	return nil
}

func (p *MigratePrimaryWithTokenV3Response) ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.LocalFormatPhoneNumber = v
	}
	return nil
}

func (p *MigratePrimaryWithTokenV3Response) ReadField6(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 6: ", err)
	} else {
		p.Mid = v
	}
	return nil
}

func (p *MigratePrimaryWithTokenV3Response) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "MigratePrimaryWithTokenV3Response"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField5(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField6(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MigratePrimaryWithTokenV3Response) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authToken", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authToken (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authToken: ", p), err)
	}
	return err
}

func (p *MigratePrimaryWithTokenV3Response) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "tokenV3IssueResult", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tokenV3IssueResult: ", p), err)
	}
	if err := p.TokenV3IssueResult_.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TokenV3IssueResult_), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tokenV3IssueResult: ", p), err)
	}
	return err
}

func (p *MigratePrimaryWithTokenV3Response) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "countryCode", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:countryCode: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.CountryCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.countryCode (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:countryCode: ", p), err)
	}
	return err
}

func (p *MigratePrimaryWithTokenV3Response) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "prettifiedFormatPhoneNumber", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:prettifiedFormatPhoneNumber: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PrettifiedFormatPhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.prettifiedFormatPhoneNumber (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:prettifiedFormatPhoneNumber: ", p), err)
	}
	return err
}

func (p *MigratePrimaryWithTokenV3Response) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "localFormatPhoneNumber", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:localFormatPhoneNumber: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.LocalFormatPhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.localFormatPhoneNumber (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:localFormatPhoneNumber: ", p), err)
	}
	return err
}

func (p *MigratePrimaryWithTokenV3Response) writeField6(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "mid", thrift.STRING, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:mid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Mid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.mid (6) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:mid: ", p), err)
	}
	return err
}

func (p *MigratePrimaryWithTokenV3Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MigratePrimaryWithTokenV3Response(%+v)", *p)
}

// Attributes:
//  - AuthToken
//  - TokenV3IssueResult_
//  - Mid
type RegisterPrimaryWithTokenV3Response struct {
	AuthToken           string               `thrift:"authToken,1" db:"authToken" json:"authToken"`
	TokenV3IssueResult_ *TokenV3IssueResult_ `thrift:"tokenV3IssueResult,2" db:"tokenV3IssueResult" json:"tokenV3IssueResult"`
	Mid                 string               `thrift:"mid,3" db:"mid" json:"mid"`
}

func NewRegisterPrimaryWithTokenV3Response() *RegisterPrimaryWithTokenV3Response {
	return &RegisterPrimaryWithTokenV3Response{}
}

func (p *RegisterPrimaryWithTokenV3Response) GetAuthToken() string {
	return p.AuthToken
}

var RegisterPrimaryWithTokenV3Response_TokenV3IssueResult__DEFAULT *TokenV3IssueResult_

func (p *RegisterPrimaryWithTokenV3Response) GetTokenV3IssueResult_() *TokenV3IssueResult_ {
	if !p.IsSetTokenV3IssueResult_() {
		return RegisterPrimaryWithTokenV3Response_TokenV3IssueResult__DEFAULT
	}
	return p.TokenV3IssueResult_
}

func (p *RegisterPrimaryWithTokenV3Response) GetMid() string {
	return p.Mid
}
func (p *RegisterPrimaryWithTokenV3Response) IsSetTokenV3IssueResult_() bool {
	return p.TokenV3IssueResult_ != nil
}

func (p *RegisterPrimaryWithTokenV3Response) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RegisterPrimaryWithTokenV3Response) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthToken = v
	}
	return nil
}

func (p *RegisterPrimaryWithTokenV3Response) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.TokenV3IssueResult_ = &TokenV3IssueResult_{}
	if err := p.TokenV3IssueResult_.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TokenV3IssueResult_), err)
	}
	return nil
}

func (p *RegisterPrimaryWithTokenV3Response) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Mid = v
	}
	return nil
}

func (p *RegisterPrimaryWithTokenV3Response) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "RegisterPrimaryWithTokenV3Response"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RegisterPrimaryWithTokenV3Response) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authToken", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authToken (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authToken: ", p), err)
	}
	return err
}

func (p *RegisterPrimaryWithTokenV3Response) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "tokenV3IssueResult", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tokenV3IssueResult: ", p), err)
	}
	if err := p.TokenV3IssueResult_.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TokenV3IssueResult_), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tokenV3IssueResult: ", p), err)
	}
	return err
}

func (p *RegisterPrimaryWithTokenV3Response) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "mid", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:mid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Mid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.mid (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:mid: ", p), err)
	}
	return err
}

func (p *RegisterPrimaryWithTokenV3Response) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RegisterPrimaryWithTokenV3Response(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - UserPhoneNumber
//  - VerifMethod
type ReqToSendPhonePinCodeRequest struct {
	AuthSessionId   string           `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	UserPhoneNumber *UserPhoneNumber `thrift:"userPhoneNumber,2" db:"userPhoneNumber" json:"userPhoneNumber"`
	VerifMethod     VerifMethod      `thrift:"verifMethod,3" db:"verifMethod" json:"verifMethod"`
}

func NewReqToSendPhonePinCodeRequest() *ReqToSendPhonePinCodeRequest {
	return &ReqToSendPhonePinCodeRequest{}
}

func (p *ReqToSendPhonePinCodeRequest) GetAuthSessionId() string {
	return p.AuthSessionId
}

var ReqToSendPhonePinCodeRequest_UserPhoneNumber_DEFAULT *UserPhoneNumber

func (p *ReqToSendPhonePinCodeRequest) GetUserPhoneNumber() *UserPhoneNumber {
	if !p.IsSetUserPhoneNumber() {
		return ReqToSendPhonePinCodeRequest_UserPhoneNumber_DEFAULT
	}
	return p.UserPhoneNumber
}

func (p *ReqToSendPhonePinCodeRequest) GetVerifMethod() VerifMethod {
	return p.VerifMethod
}
func (p *ReqToSendPhonePinCodeRequest) IsSetUserPhoneNumber() bool {
	return p.UserPhoneNumber != nil
}

func (p *ReqToSendPhonePinCodeRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ReqToSendPhonePinCodeRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *ReqToSendPhonePinCodeRequest) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserPhoneNumber = &UserPhoneNumber{}
	if err := p.UserPhoneNumber.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserPhoneNumber), err)
	}
	return nil
}

func (p *ReqToSendPhonePinCodeRequest) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		temp := VerifMethod(v)
		p.VerifMethod = temp
	}
	return nil
}

func (p *ReqToSendPhonePinCodeRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ReqToSendPhonePinCodeRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ReqToSendPhonePinCodeRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *ReqToSendPhonePinCodeRequest) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userPhoneNumber", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:userPhoneNumber: ", p), err)
	}
	if err := p.UserPhoneNumber.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserPhoneNumber), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:userPhoneNumber: ", p), err)
	}
	return err
}

func (p *ReqToSendPhonePinCodeRequest) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "verifMethod", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:verifMethod: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.VerifMethod)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.verifMethod (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:verifMethod: ", p), err)
	}
	return err
}

func (p *ReqToSendPhonePinCodeRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReqToSendPhonePinCodeRequest(%+v)", *p)
}

// Attributes:
//  - AvailableMethods
type ReqToSendPhonePinCodeResponse struct {
	AvailableMethods []AvailableMethods `thrift:"availableMethods,1" db:"availableMethods" json:"availableMethods"`
}

func NewReqToSendPhonePinCodeResponse() *ReqToSendPhonePinCodeResponse {
	return &ReqToSendPhonePinCodeResponse{}
}

func (p *ReqToSendPhonePinCodeResponse) GetAvailableMethods() []AvailableMethods {
	return p.AvailableMethods
}
func (p *ReqToSendPhonePinCodeResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ReqToSendPhonePinCodeResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]AvailableMethods, 0, size)
	p.AvailableMethods = tSlice
	for i := 0; i < size; i++ {
		var _elem4 AvailableMethods
		if v, err := iprot.ReadI32(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			temp := AvailableMethods(v)
			_elem4 = temp
		}
		p.AvailableMethods = append(p.AvailableMethods, _elem4)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *ReqToSendPhonePinCodeResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ReqToSendPhonePinCodeResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ReqToSendPhonePinCodeResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "availableMethods", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:availableMethods: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.I32, len(p.AvailableMethods)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.AvailableMethods {
		if err := oprot.WriteI32(ctx, int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:availableMethods: ", p), err)
	}
	return err
}

func (p *ReqToSendPhonePinCodeResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReqToSendPhonePinCodeResponse(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - UserPhoneNumber
//  - PinCode
type VerifyPhonePinCodeRequest struct {
	AuthSessionId   string           `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	UserPhoneNumber *UserPhoneNumber `thrift:"userPhoneNumber,2" db:"userPhoneNumber" json:"userPhoneNumber"`
	PinCode         string           `thrift:"pinCode,3" db:"pinCode" json:"pinCode"`
}

func NewVerifyPhonePinCodeRequest() *VerifyPhonePinCodeRequest {
	return &VerifyPhonePinCodeRequest{}
}

func (p *VerifyPhonePinCodeRequest) GetAuthSessionId() string {
	return p.AuthSessionId
}

var VerifyPhonePinCodeRequest_UserPhoneNumber_DEFAULT *UserPhoneNumber

func (p *VerifyPhonePinCodeRequest) GetUserPhoneNumber() *UserPhoneNumber {
	if !p.IsSetUserPhoneNumber() {
		return VerifyPhonePinCodeRequest_UserPhoneNumber_DEFAULT
	}
	return p.UserPhoneNumber
}

func (p *VerifyPhonePinCodeRequest) GetPinCode() string {
	return p.PinCode
}
func (p *VerifyPhonePinCodeRequest) IsSetUserPhoneNumber() bool {
	return p.UserPhoneNumber != nil
}

func (p *VerifyPhonePinCodeRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *VerifyPhonePinCodeRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *VerifyPhonePinCodeRequest) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserPhoneNumber = &UserPhoneNumber{}
	if err := p.UserPhoneNumber.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserPhoneNumber), err)
	}
	return nil
}

func (p *VerifyPhonePinCodeRequest) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.PinCode = v
	}
	return nil
}

func (p *VerifyPhonePinCodeRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "VerifyPhonePinCodeRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifyPhonePinCodeRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *VerifyPhonePinCodeRequest) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userPhoneNumber", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:userPhoneNumber: ", p), err)
	}
	if err := p.UserPhoneNumber.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserPhoneNumber), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:userPhoneNumber: ", p), err)
	}
	return err
}

func (p *VerifyPhonePinCodeRequest) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "pinCode", thrift.STRING, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:pinCode: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PinCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.pinCode (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:pinCode: ", p), err)
	}
	return err
}

func (p *VerifyPhonePinCodeRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifyPhonePinCodeRequest(%+v)", *p)
}

// Attributes:
//  - AccountExist
//  - SameUdidFromAccount
//  - UserProfile
type VerifyPhonePinCodeResponse struct {
	AccountExist        bool `thrift:"accountExist,1" db:"accountExist" json:"accountExist"`
	SameUdidFromAccount bool `thrift:"sameUdidFromAccount,2" db:"sameUdidFromAccount" json:"sameUdidFromAccount"`
	// unused fields # 3 to 10
	UserProfile *UserProfile `thrift:"userProfile,11" db:"userProfile" json:"userProfile"`
}

func NewVerifyPhonePinCodeResponse() *VerifyPhonePinCodeResponse {
	return &VerifyPhonePinCodeResponse{}
}

func (p *VerifyPhonePinCodeResponse) GetAccountExist() bool {
	return p.AccountExist
}

func (p *VerifyPhonePinCodeResponse) GetSameUdidFromAccount() bool {
	return p.SameUdidFromAccount
}

var VerifyPhonePinCodeResponse_UserProfile_DEFAULT *UserProfile

func (p *VerifyPhonePinCodeResponse) GetUserProfile() *UserProfile {
	if !p.IsSetUserProfile() {
		return VerifyPhonePinCodeResponse_UserProfile_DEFAULT
	}
	return p.UserProfile
}
func (p *VerifyPhonePinCodeResponse) IsSetUserProfile() bool {
	return p.UserProfile != nil
}

func (p *VerifyPhonePinCodeResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 11:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField11(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *VerifyPhonePinCodeResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AccountExist = v
	}
	return nil
}

func (p *VerifyPhonePinCodeResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.SameUdidFromAccount = v
	}
	return nil
}

func (p *VerifyPhonePinCodeResponse) ReadField11(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserProfile = &UserProfile{}
	if err := p.UserProfile.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserProfile), err)
	}
	return nil
}

func (p *VerifyPhonePinCodeResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "VerifyPhonePinCodeResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField11(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifyPhonePinCodeResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountExist", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:accountExist: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.AccountExist)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accountExist (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:accountExist: ", p), err)
	}
	return err
}

func (p *VerifyPhonePinCodeResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sameUdidFromAccount", thrift.BOOL, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:sameUdidFromAccount: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.SameUdidFromAccount)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sameUdidFromAccount (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:sameUdidFromAccount: ", p), err)
	}
	return err
}

func (p *VerifyPhonePinCodeResponse) writeField11(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userProfile", thrift.STRUCT, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:userProfile: ", p), err)
	}
	if err := p.UserProfile.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserProfile), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:userProfile: ", p), err)
	}
	return err
}

func (p *VerifyPhonePinCodeResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifyPhonePinCodeResponse(%+v)", *p)
}

type PrimaryAccountInitService interface {
	// Parameters:
	//  - AuthSessionId
	RegisterPrimaryUsingPhone(ctx context.Context, authSessionId string) (r *RegisterPrimaryResponse, err error)
	// Parameters:
	//  - AuthSessionId
	RegisterPrimaryUsingSocialLogin(ctx context.Context, authSessionId string) (r *RegisterPrimaryResponse, err error)
	// Parameters:
	//  - AuthSessionId
	MigratePrimaryUsingSocialLogin(ctx context.Context, authSessionId string) (r *MigratePrimaryResponse, err error)
	// Parameters:
	//  - AuthSessionId
	MigratePrimaryUsingPhone(ctx context.Context, authSessionId string) (r *MigratePrimaryResponse, err error)
	// Parameters:
	//  - AuthSessionId
	GetSecondAuthMethod(ctx context.Context, authSessionId string) (r *GetSecondAuthMethodResponse, err error)
	// Parameters:
	//  - AuthSessionId
	IssueWebAuthDetailsForSecondAuth(ctx context.Context, authSessionId string) (r *IssueWebAuthDetailsForSecondAuthResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - AccountIdentifier
	IssueWebAuthDetailsForAcctVerif(ctx context.Context, authSessionId string, accountIdentifier *AccountIdentifier) (r *IssueWebAuthDetailsForAcctVerifResponse, err error)
	// Parameters:
	//  - Request
	OpenSession(ctx context.Context, request *OpenSessionRequest) (r string, err error)
	// Parameters:
	//  - AuthSessionId
	//  - AccountIdentifier
	//  - EncryptedPassword
	VerifyAccountUsingPwd(ctx context.Context, authSessionId string, accountIdentifier *AccountIdentifier, encryptedPassword *EncryptedPassword) (r *VerifyAccountUsingPwdResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - Device
	//  - SocialLogin
	VerifySocialLogin(ctx context.Context, authSessionId string, device *Device, socialLogin *SocialLogin) (r *VerifySocialLoginResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - Device
	//  - UserPhoneNumber
	GetPhoneVerifMethod(ctx context.Context, authSessionId string, device *Device, userPhoneNumber *UserPhoneNumber) (r *GetPhoneVerifMethodResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - Device
	//  - UserPhoneNumber
	//  - VerifMethod
	SendPinCodeForPhone(ctx context.Context, authSessionId string, device *Device, userPhoneNumber *UserPhoneNumber, verifMethod VerifMethod) (r *SendPinCodeForPhoneResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - Device
	//  - UserPhoneNumber
	//  - PinCode
	VerifyPhone(ctx context.Context, authSessionId string, device *Device, userPhoneNumber *UserPhoneNumber, pinCode string) (r *VerifyPhoneResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - EncryptedPassword
	SetPassword(ctx context.Context, authSessionId string, encryptedPassword *EncryptedPassword) (err error)
	// Parameters:
	//  - AuthSessionId
	//  - Request
	ExchangeEncryptionKey(ctx context.Context, authSessionId string, request *ExchangeEncryptionKeyRequest) (r *ExchangeEncryptionKeyResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - SimCard
	GetCountryInfo(ctx context.Context, authSessionId string, simCard *SimCard) (r *GetCountryInfoResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - AccountIdentifier
	GetUserProfile(ctx context.Context, authSessionId string, accountIdentifier *AccountIdentifier) (r *GetUserProfileResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - AccountIdentifier
	GetAcctVerifMethod(ctx context.Context, authSessionId string, accountIdentifier *AccountIdentifier) (r *GetAcctVerifMethodResponse, err error)
	// Parameters:
	//  - AuthSessionId
	//  - DisplayName
	ValidateProfile(ctx context.Context, authSessionId string, displayName string) (err error)
	// Parameters:
	//  - Request
	FetchPhonePinCodeMsg(ctx context.Context, request *FetchPhonePinCodeMsgRequest) (r *FetchPhonePinCodeMsgResponse, err error)
	// Parameters:
	//  - Request
	GetPhoneVerifMethodV2(ctx context.Context, request *GetPhoneVerifMethodV2Request) (r *GetPhoneVerifMethodV2Response, err error)
	// Parameters:
	//  - AuthSessionId
	GetSessionContentBeforeMigCompletion(ctx context.Context, authSessionId string) (r *GetSessionContentBeforeMigCompletionResponse, err error)
	// Parameters:
	//  - AuthSessionId
	MigratePrimaryUsingEapAccountWithTokenV3(ctx context.Context, authSessionId string) (r *MigratePrimaryWithTokenV3Response, err error)
	// Parameters:
	//  - AuthSessionId
	MigratePrimaryUsingPhoneWithTokenV3(ctx context.Context, authSessionId string) (r *MigratePrimaryWithTokenV3Response, err error)
	// Parameters:
	//  - AuthSessionId
	RegisterPrimaryUsingPhoneWithTokenV3(ctx context.Context, authSessionId string) (r *RegisterPrimaryWithTokenV3Response, err error)
	// Parameters:
	//  - Request
	RequestToSendPhonePinCode(ctx context.Context, request *ReqToSendPhonePinCodeRequest) (r *ReqToSendPhonePinCodeResponse, err error)
	// Parameters:
	//  - Request
	VerifyPhonePinCode(ctx context.Context, request *VerifyPhonePinCodeRequest) (r *VerifyPhonePinCodeResponse, err error)
	// Parameters:
	//  - Request
	LookupAvailableEap(ctx context.Context, request *LookupAvailableEapRequest) (_r *LookupAvailableEapResponse, _err error)
}

type PrimaryAccountInitServiceClient struct {
	c thrift.TClient
}

func NewPrimaryAccountInitServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PrimaryAccountInitServiceClient {
	return &PrimaryAccountInitServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPrimaryAccountInitServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PrimaryAccountInitServiceClient {
	return &PrimaryAccountInitServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPrimaryAccountInitServiceClient(c thrift.TClient) *PrimaryAccountInitServiceClient {
	return &PrimaryAccountInitServiceClient{
		c: c,
	}
}

func (p *PrimaryAccountInitServiceClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - Request
func (p *PrimaryAccountInitServiceClient) FetchPhonePinCodeMsg(ctx context.Context, request *FetchPhonePinCodeMsgRequest) (r *FetchPhonePinCodeMsgResponse, err error) {
	var _args7 PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs
	_args7.Request = request
	var _result8 PrimaryAccountInitServiceFetchPhonePinCodeMsgResult
	if _, err = p.Client_().Call(ctx, "fetchPhonePinCodeMsg", &_args7, &_result8); err != nil {
		return
	}
	switch {
	case _result8.E != nil:
		return r, _result8.E
	}

	return _result8.GetSuccess(), nil
}

// Parameters:
//  - Request
func (p *PrimaryAccountInitServiceClient) GetPhoneVerifMethodV2(ctx context.Context, request *GetPhoneVerifMethodV2Request) (r *GetPhoneVerifMethodV2Response, err error) {
	var _args13 PrimaryAccountInitServiceGetPhoneVerifMethodV2Args
	_args13.Request = request
	var _result14 PrimaryAccountInitServiceGetPhoneVerifMethodV2Result
	if _, err = p.Client_().Call(ctx, "getPhoneVerifMethodV2", &_args13, &_result14); err != nil {
		return
	}
	switch {
	case _result14.E != nil:
		return r, _result14.E
	}

	return _result14.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) GetSessionContentBeforeMigCompletion(ctx context.Context, authSessionId string) (r *GetSessionContentBeforeMigCompletionResponse, err error) {
	var _args17 PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs
	_args17.AuthSessionId = authSessionId
	var _result18 PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult
	if _, err = p.Client_().Call(ctx, "getSessionContentBeforeMigCompletion", &_args17, &_result18); err != nil {
		return
	}
	switch {
	case _result18.E != nil:
		return r, _result18.E
	}

	return _result18.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) MigratePrimaryUsingEapAccountWithTokenV3(ctx context.Context, authSessionId string) (r *MigratePrimaryWithTokenV3Response, err error) {
	var _args25 PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args
	_args25.AuthSessionId = authSessionId
	var _result26 PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result
	if _, err = p.Client_().Call(ctx, "migratePrimaryUsingEapAccountWithTokenV3", &_args25, &_result26); err != nil {
		return
	}
	switch {
	case _result26.E != nil:
		return r, _result26.E
	}

	return _result26.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) MigratePrimaryUsingPhoneWithTokenV3(ctx context.Context, authSessionId string) (r *MigratePrimaryWithTokenV3Response, err error) {
	var _args27 PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args
	_args27.AuthSessionId = authSessionId
	var _result28 PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result
	if _, err = p.Client_().Call(ctx, "migratePrimaryUsingPhoneWithTokenV3", &_args27, &_result28); err != nil {
		return
	}
	switch {
	case _result28.E != nil:
		return r, _result28.E
	}

	return _result28.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) RegisterPrimaryUsingPhoneWithTokenV3(ctx context.Context, authSessionId string) (r *RegisterPrimaryWithTokenV3Response, err error) {
	var _args31 PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args
	_args31.AuthSessionId = authSessionId
	var _result32 PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result
	if _, err = p.Client_().Call(ctx, "registerPrimaryUsingPhoneWithTokenV3", &_args31, &_result32); err != nil {
		return
	}
	switch {
	case _result32.E != nil:
		return r, _result32.E
	}

	return _result32.GetSuccess(), nil
}

// Parameters:
//  - Request
func (p *PrimaryAccountInitServiceClient) RequestToSendPhonePinCode(ctx context.Context, request *ReqToSendPhonePinCodeRequest) (r *ReqToSendPhonePinCodeResponse, err error) {
	var _args35 PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs
	_args35.Request = request
	var _result36 PrimaryAccountInitServiceRequestToSendPhonePinCodeResult
	if _, err = p.Client_().Call(ctx, "requestToSendPhonePinCode", &_args35, &_result36); err != nil {
		return
	}
	switch {
	case _result36.E != nil:
		return r, _result36.E
	}

	return _result36.GetSuccess(), nil
}

// Parameters:
//  - Request
func (p *PrimaryAccountInitServiceClient) VerifyPhonePinCode(ctx context.Context, request *VerifyPhonePinCodeRequest) (r *VerifyPhonePinCodeResponse, err error) {
	var _args43 PrimaryAccountInitServiceVerifyPhonePinCodeArgs
	_args43.Request = request
	var _result44 PrimaryAccountInitServiceVerifyPhonePinCodeResult
	if _, err = p.Client_().Call(ctx, "verifyPhonePinCode", &_args43, &_result44); err != nil {
		return
	}
	switch {
	case _result44.E != nil:
		return r, _result44.E
	}

	return _result44.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) RegisterPrimaryUsingPhone(ctx context.Context, authSessionId string) (r *RegisterPrimaryResponse, err error) {
	var _args5 PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs
	_args5.AuthSessionId = authSessionId
	var _result6 PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult
	if _, err = p.Client_().Call(ctx, "registerPrimaryUsingPhone", &_args5, &_result6); err != nil {
		return
	}
	switch {
	case _result6.E != nil:
		return r, _result6.E
	}

	return _result6.GetSuccess(), nil
}

// Parameters:
//  - Request
func (p *PrimaryAccountInitServiceClient) LookupAvailableEap(ctx context.Context, request *LookupAvailableEapRequest) (_r *LookupAvailableEapResponse, _err error) {
	var _args88 PrimaryAccountInitServiceLookupAvailableEapArgs
	_args88.Request = request
	var _result90 PrimaryAccountInitServiceLookupAvailableEapResult
	_, _err = p.Client_().Call(ctx, "lookupAvailableEap", &_args88, &_result90)
	if _err != nil {
		return
	}
	switch {
	case _result90.E != nil:
		return _r, _result90.E
	}

	if _ret91 := _result90.GetSuccess(); _ret91 != nil {
		return _ret91, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "lookupAvailableEap failed: unknown result")
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) RegisterPrimaryUsingSocialLogin(ctx context.Context, authSessionId string) (r *RegisterPrimaryResponse, err error) {
	var _args7 PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs
	_args7.AuthSessionId = authSessionId
	var _result8 PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult
	if _, err = p.Client_().Call(ctx, "registerPrimaryUsingSocialLogin", &_args7, &_result8); err != nil {
		return
	}
	switch {
	case _result8.E != nil:
		return r, _result8.E
	}

	return _result8.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) MigratePrimaryUsingSocialLogin(ctx context.Context, authSessionId string) (r *MigratePrimaryResponse, err error) {
	var _args9 PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs
	_args9.AuthSessionId = authSessionId
	var _result10 PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult
	if _, err = p.Client_().Call(ctx, "migratePrimaryUsingSocialLogin", &_args9, &_result10); err != nil {
		return
	}
	switch {
	case _result10.E != nil:
		return r, _result10.E
	}

	return _result10.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) MigratePrimaryUsingPhone(ctx context.Context, authSessionId string) (r *MigratePrimaryResponse, err error) {
	var _args11 PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs
	_args11.AuthSessionId = authSessionId
	var _result12 PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult
	if _, err = p.Client_().Call(ctx, "migratePrimaryUsingPhone", &_args11, &_result12); err != nil {
		return
	}
	switch {
	case _result12.E != nil:
		return r, _result12.E
	}

	return _result12.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) GetSecondAuthMethod(ctx context.Context, authSessionId string) (r *GetSecondAuthMethodResponse, err error) {
	var _args13 PrimaryAccountInitServiceGetSecondAuthMethodArgs
	_args13.AuthSessionId = authSessionId
	var _result14 PrimaryAccountInitServiceGetSecondAuthMethodResult
	if _, err = p.Client_().Call(ctx, "getSecondAuthMethod", &_args13, &_result14); err != nil {
		return
	}
	switch {
	case _result14.E != nil:
		return r, _result14.E
	}

	return _result14.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
func (p *PrimaryAccountInitServiceClient) IssueWebAuthDetailsForSecondAuth(ctx context.Context, authSessionId string) (r *IssueWebAuthDetailsForSecondAuthResponse, err error) {
	var _args15 PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs
	_args15.AuthSessionId = authSessionId
	var _result16 PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult
	if _, err = p.Client_().Call(ctx, "issueWebAuthDetailsForSecondAuth", &_args15, &_result16); err != nil {
		return
	}
	switch {
	case _result16.E != nil:
		return r, _result16.E
	}

	return _result16.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - AccountIdentifier
func (p *PrimaryAccountInitServiceClient) IssueWebAuthDetailsForAcctVerif(ctx context.Context, authSessionId string, accountIdentifier *AccountIdentifier) (r *IssueWebAuthDetailsForAcctVerifResponse, err error) {
	var _args17 PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs
	_args17.AuthSessionId = authSessionId
	_args17.AccountIdentifier = accountIdentifier
	var _result18 PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult
	if _, err = p.Client_().Call(ctx, "issueWebAuthDetailsForAcctVerif", &_args17, &_result18); err != nil {
		return
	}
	switch {
	case _result18.E != nil:
		return r, _result18.E
	}

	return _result18.GetSuccess(), nil
}

// Parameters:
//  - Request
func (p *PrimaryAccountInitServiceClient) OpenSession(ctx context.Context, request *OpenSessionRequest) (r string, err error) {
	var _args19 PrimaryAccountInitServiceOpenSessionArgs
	_args19.Request = request
	var _result20 PrimaryAccountInitServiceOpenSessionResult
	if _, err = p.Client_().Call(ctx, "openSession", &_args19, &_result20); err != nil {
		return
	}
	switch {
	case _result20.E != nil:
		return r, _result20.E
	}

	return _result20.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - AccountIdentifier
//  - EncryptedPassword
func (p *PrimaryAccountInitServiceClient) VerifyAccountUsingPwd(ctx context.Context, authSessionId string, accountIdentifier *AccountIdentifier, encryptedPassword *EncryptedPassword) (r *VerifyAccountUsingPwdResponse, err error) {
	var _args21 PrimaryAccountInitServiceVerifyAccountUsingPwdArgs
	_args21.AuthSessionId = authSessionId
	_args21.AccountIdentifier = accountIdentifier
	_args21.EncryptedPassword = encryptedPassword
	var _result22 PrimaryAccountInitServiceVerifyAccountUsingPwdResult
	if _, err = p.Client_().Call(ctx, "verifyAccountUsingPwd", &_args21, &_result22); err != nil {
		return
	}
	switch {
	case _result22.E != nil:
		return r, _result22.E
	}

	return _result22.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - Device
//  - SocialLogin
func (p *PrimaryAccountInitServiceClient) VerifySocialLogin(ctx context.Context, authSessionId string, device *Device, socialLogin *SocialLogin) (r *VerifySocialLoginResponse, err error) {
	var _args23 PrimaryAccountInitServiceVerifySocialLoginArgs
	_args23.AuthSessionId = authSessionId
	_args23.Device = device
	_args23.SocialLogin = socialLogin
	var _result24 PrimaryAccountInitServiceVerifySocialLoginResult
	if _, err = p.Client_().Call(ctx, "verifySocialLogin", &_args23, &_result24); err != nil {
		return
	}
	switch {
	case _result24.E != nil:
		return r, _result24.E
	}

	return _result24.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - Device
//  - UserPhoneNumber
func (p *PrimaryAccountInitServiceClient) GetPhoneVerifMethod(ctx context.Context, authSessionId string, device *Device, userPhoneNumber *UserPhoneNumber) (r *GetPhoneVerifMethodResponse, err error) {
	var _args25 PrimaryAccountInitServiceGetPhoneVerifMethodArgs
	_args25.AuthSessionId = authSessionId
	_args25.Device = device
	_args25.UserPhoneNumber = userPhoneNumber
	var _result26 PrimaryAccountInitServiceGetPhoneVerifMethodResult
	if _, err = p.Client_().Call(ctx, "getPhoneVerifMethod", &_args25, &_result26); err != nil {
		return
	}
	switch {
	case _result26.E != nil:
		return r, _result26.E
	}

	return _result26.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - Device
//  - UserPhoneNumber
//  - VerifMethod
func (p *PrimaryAccountInitServiceClient) SendPinCodeForPhone(ctx context.Context, authSessionId string, device *Device, userPhoneNumber *UserPhoneNumber, verifMethod VerifMethod) (r *SendPinCodeForPhoneResponse, err error) {
	var _args27 PrimaryAccountInitServiceSendPinCodeForPhoneArgs
	_args27.AuthSessionId = authSessionId
	_args27.Device = device
	_args27.UserPhoneNumber = userPhoneNumber
	_args27.VerifMethod = verifMethod
	var _result28 PrimaryAccountInitServiceSendPinCodeForPhoneResult
	if _, err = p.Client_().Call(ctx, "sendPinCodeForPhone", &_args27, &_result28); err != nil {
		return
	}
	switch {
	case _result28.E != nil:
		return r, _result28.E
	}

	return _result28.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - Device
//  - UserPhoneNumber
//  - PinCode
func (p *PrimaryAccountInitServiceClient) VerifyPhone(ctx context.Context, authSessionId string, device *Device, userPhoneNumber *UserPhoneNumber, pinCode string) (r *VerifyPhoneResponse, err error) {
	var _args29 PrimaryAccountInitServiceVerifyPhoneArgs
	_args29.AuthSessionId = authSessionId
	_args29.Device = device
	_args29.UserPhoneNumber = userPhoneNumber
	_args29.PinCode = pinCode
	var _result30 PrimaryAccountInitServiceVerifyPhoneResult
	if _, err = p.Client_().Call(ctx, "verifyPhone", &_args29, &_result30); err != nil {
		return
	}
	switch {
	case _result30.E != nil:
		return r, _result30.E
	}

	return _result30.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - EncryptedPassword
func (p *PrimaryAccountInitServiceClient) SetPassword(ctx context.Context, authSessionId string, encryptedPassword *EncryptedPassword) (err error) {
	var _args31 PrimaryAccountInitServiceSetPasswordArgs
	_args31.AuthSessionId = authSessionId
	_args31.EncryptedPassword = encryptedPassword
	var _result32 PrimaryAccountInitServiceSetPasswordResult
	if _, err = p.Client_().Call(ctx, "setPassword", &_args31, &_result32); err != nil {
		return
	}
	switch {
	case _result32.E != nil:
		return _result32.E
	}

	return nil
}

// Parameters:
//  - AuthSessionId
//  - Request
func (p *PrimaryAccountInitServiceClient) ExchangeEncryptionKey(ctx context.Context, authSessionId string, request *ExchangeEncryptionKeyRequest) (r *ExchangeEncryptionKeyResponse, err error) {
	var _args33 PrimaryAccountInitServiceExchangeEncryptionKeyArgs
	_args33.AuthSessionId = authSessionId
	_args33.Request = request
	var _result34 PrimaryAccountInitServiceExchangeEncryptionKeyResult
	if _, err = p.Client_().Call(ctx, "exchangeEncryptionKey", &_args33, &_result34); err != nil {
		return
	}
	switch {
	case _result34.E != nil:
		return r, _result34.E
	}

	return _result34.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - SimCard
func (p *PrimaryAccountInitServiceClient) GetCountryInfo(ctx context.Context, authSessionId string, simCard *SimCard) (r *GetCountryInfoResponse, err error) {
	var _args35 PrimaryAccountInitServiceGetCountryInfoArgs
	_args35.AuthSessionId = authSessionId
	_args35.SimCard = simCard
	var _result36 PrimaryAccountInitServiceGetCountryInfoResult
	if _, err = p.Client_().Call(ctx, "getCountryInfo", &_args35, &_result36); err != nil {
		return
	}
	switch {
	case _result36.E != nil:
		return r, _result36.E
	}

	return _result36.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - AccountIdentifier
func (p *PrimaryAccountInitServiceClient) GetUserProfile(ctx context.Context, authSessionId string, accountIdentifier *AccountIdentifier) (r *GetUserProfileResponse, err error) {
	var _args37 PrimaryAccountInitServiceGetUserProfileArgs
	_args37.AuthSessionId = authSessionId
	_args37.AccountIdentifier = accountIdentifier
	var _result38 PrimaryAccountInitServiceGetUserProfileResult
	if _, err = p.Client_().Call(ctx, "getUserProfile", &_args37, &_result38); err != nil {
		return
	}
	switch {
	case _result38.E != nil:
		return r, _result38.E
	}

	return _result38.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - AccountIdentifier
func (p *PrimaryAccountInitServiceClient) GetAcctVerifMethod(ctx context.Context, authSessionId string, accountIdentifier *AccountIdentifier) (r *GetAcctVerifMethodResponse, err error) {
	var _args39 PrimaryAccountInitServiceGetAcctVerifMethodArgs
	_args39.AuthSessionId = authSessionId
	_args39.AccountIdentifier = accountIdentifier
	var _result40 PrimaryAccountInitServiceGetAcctVerifMethodResult
	if _, err = p.Client_().Call(ctx, "getAcctVerifMethod", &_args39, &_result40); err != nil {
		return
	}
	switch {
	case _result40.E != nil:
		return r, _result40.E
	}

	return _result40.GetSuccess(), nil
}

// Parameters:
//  - AuthSessionId
//  - DisplayName
func (p *PrimaryAccountInitServiceClient) ValidateProfile(ctx context.Context, authSessionId string, displayName string) (err error) {
	var _args41 PrimaryAccountInitServiceValidateProfileArgs
	_args41.AuthSessionId = authSessionId
	_args41.DisplayName = displayName
	var _result42 PrimaryAccountInitServiceValidateProfileResult
	if _, err = p.Client_().Call(ctx, "validateProfile", &_args41, &_result42); err != nil {
		return
	}
	switch {
	case _result42.E != nil:
		return _result42.E
	}

	return nil
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs struct {
	// unused field # 1
	AuthSessionId string `thrift:"authSessionId,2" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs() *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs {
	return &PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs{}
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "registerPrimaryUsingPhone_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceRegisterPrimaryUsingPhoneArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult struct {
	Success *RegisterPrimaryResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException           `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult() *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult {
	return &PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult{}
}

var PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult_Success_DEFAULT *RegisterPrimaryResponse

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) GetSuccess() *RegisterPrimaryResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &RegisterPrimaryResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "registerPrimaryUsingPhone_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceRegisterPrimaryUsingPhoneResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs struct {
	// unused field # 1
	AuthSessionId string `thrift:"authSessionId,2" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs() *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs {
	return &PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs{}
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "registerPrimaryUsingSocialLogin_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult struct {
	Success *RegisterPrimaryResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException           `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult() *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult {
	return &PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult{}
}

var PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult_Success_DEFAULT *RegisterPrimaryResponse

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) GetSuccess() *RegisterPrimaryResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &RegisterPrimaryResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "registerPrimaryUsingSocialLogin_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceRegisterPrimaryUsingSocialLoginResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs() *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs {
	return &PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs{}
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingSocialLogin_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult struct {
	Success *MigratePrimaryResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException          `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult() *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult {
	return &PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult{}
}

var PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult_Success_DEFAULT *MigratePrimaryResponse

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) GetSuccess() *MigratePrimaryResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &MigratePrimaryResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingSocialLogin_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceMigratePrimaryUsingSocialLoginResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs() *PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs {
	return &PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs{}
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingPhone_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceMigratePrimaryUsingPhoneArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult struct {
	Success *MigratePrimaryResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException          `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceMigratePrimaryUsingPhoneResult() *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult {
	return &PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult{}
}

var PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult_Success_DEFAULT *MigratePrimaryResponse

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) GetSuccess() *MigratePrimaryResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &MigratePrimaryResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingPhone_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceMigratePrimaryUsingPhoneResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceGetSecondAuthMethodArgs struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceGetSecondAuthMethodArgs() *PrimaryAccountInitServiceGetSecondAuthMethodArgs {
	return &PrimaryAccountInitServiceGetSecondAuthMethodArgs{}
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceGetSecondAuthMethodArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getSecondAuthMethod_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetSecondAuthMethodArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceGetSecondAuthMethodResult struct {
	Success *GetSecondAuthMethodResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException               `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceGetSecondAuthMethodResult() *PrimaryAccountInitServiceGetSecondAuthMethodResult {
	return &PrimaryAccountInitServiceGetSecondAuthMethodResult{}
}

var PrimaryAccountInitServiceGetSecondAuthMethodResult_Success_DEFAULT *GetSecondAuthMethodResponse

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) GetSuccess() *GetSecondAuthMethodResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceGetSecondAuthMethodResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceGetSecondAuthMethodResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceGetSecondAuthMethodResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &GetSecondAuthMethodResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getSecondAuthMethod_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetSecondAuthMethodResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetSecondAuthMethodResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs() *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs {
	return &PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs{}
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "issueWebAuthDetailsForSecondAuth_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult struct {
	Success *IssueWebAuthDetailsForSecondAuthResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                            `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult() *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult {
	return &PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult{}
}

var PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult_Success_DEFAULT *IssueWebAuthDetailsForSecondAuthResponse

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) GetSuccess() *IssueWebAuthDetailsForSecondAuthResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &IssueWebAuthDetailsForSecondAuthResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "issueWebAuthDetailsForSecondAuth_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceIssueWebAuthDetailsForSecondAuthResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - AccountIdentifier
type PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs struct {
	AuthSessionId     string             `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	AccountIdentifier *AccountIdentifier `thrift:"accountIdentifier,2" db:"accountIdentifier" json:"accountIdentifier"`
}

func NewPrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs() *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs {
	return &PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs{}
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs_AccountIdentifier_DEFAULT *AccountIdentifier

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) GetAccountIdentifier() *AccountIdentifier {
	if !p.IsSetAccountIdentifier() {
		return PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs_AccountIdentifier_DEFAULT
	}
	return p.AccountIdentifier
}
func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) IsSetAccountIdentifier() bool {
	return p.AccountIdentifier != nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.AccountIdentifier = &AccountIdentifier{}
	if err := p.AccountIdentifier.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.AccountIdentifier), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "issueWebAuthDetailsForAcctVerif_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountIdentifier", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:accountIdentifier: ", p), err)
	}
	if err := p.AccountIdentifier.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.AccountIdentifier), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:accountIdentifier: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult struct {
	Success *IssueWebAuthDetailsForAcctVerifResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                           `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult() *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult {
	return &PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult{}
}

var PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult_Success_DEFAULT *IssueWebAuthDetailsForAcctVerifResponse

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) GetSuccess() *IssueWebAuthDetailsForAcctVerifResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &IssueWebAuthDetailsForAcctVerifResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "issueWebAuthDetailsForAcctVerif_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceIssueWebAuthDetailsForAcctVerifResult(%+v)", *p)
}

// Attributes:
//  - Request
type PrimaryAccountInitServiceOpenSessionArgs struct {
	Request *OpenSessionRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryAccountInitServiceOpenSessionArgs() *PrimaryAccountInitServiceOpenSessionArgs {
	return &PrimaryAccountInitServiceOpenSessionArgs{}
}

var PrimaryAccountInitServiceOpenSessionArgs_Request_DEFAULT *OpenSessionRequest

func (p *PrimaryAccountInitServiceOpenSessionArgs) GetRequest() *OpenSessionRequest {
	if !p.IsSetRequest() {
		return PrimaryAccountInitServiceOpenSessionArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryAccountInitServiceOpenSessionArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryAccountInitServiceOpenSessionArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceOpenSessionArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &OpenSessionRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceOpenSessionArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "openSession_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceOpenSessionArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceOpenSessionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceOpenSessionArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceOpenSessionResult struct {
	Success *string        `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceOpenSessionResult() *PrimaryAccountInitServiceOpenSessionResult {
	return &PrimaryAccountInitServiceOpenSessionResult{}
}

var PrimaryAccountInitServiceOpenSessionResult_Success_DEFAULT string

func (p *PrimaryAccountInitServiceOpenSessionResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceOpenSessionResult_Success_DEFAULT
	}
	return *p.Success
}

var PrimaryAccountInitServiceOpenSessionResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceOpenSessionResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceOpenSessionResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceOpenSessionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceOpenSessionResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceOpenSessionResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceOpenSessionResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *PrimaryAccountInitServiceOpenSessionResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceOpenSessionResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "openSession_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceOpenSessionResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(ctx, string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceOpenSessionResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceOpenSessionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceOpenSessionResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - AccountIdentifier
//  - EncryptedPassword
type PrimaryAccountInitServiceVerifyAccountUsingPwdArgs struct {
	AuthSessionId     string             `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	AccountIdentifier *AccountIdentifier `thrift:"accountIdentifier,2" db:"accountIdentifier" json:"accountIdentifier"`
	EncryptedPassword *EncryptedPassword `thrift:"encryptedPassword,3" db:"encryptedPassword" json:"encryptedPassword"`
}

func NewPrimaryAccountInitServiceVerifyAccountUsingPwdArgs() *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs {
	return &PrimaryAccountInitServiceVerifyAccountUsingPwdArgs{}
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceVerifyAccountUsingPwdArgs_AccountIdentifier_DEFAULT *AccountIdentifier

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) GetAccountIdentifier() *AccountIdentifier {
	if !p.IsSetAccountIdentifier() {
		return PrimaryAccountInitServiceVerifyAccountUsingPwdArgs_AccountIdentifier_DEFAULT
	}
	return p.AccountIdentifier
}

var PrimaryAccountInitServiceVerifyAccountUsingPwdArgs_EncryptedPassword_DEFAULT *EncryptedPassword

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) GetEncryptedPassword() *EncryptedPassword {
	if !p.IsSetEncryptedPassword() {
		return PrimaryAccountInitServiceVerifyAccountUsingPwdArgs_EncryptedPassword_DEFAULT
	}
	return p.EncryptedPassword
}
func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) IsSetAccountIdentifier() bool {
	return p.AccountIdentifier != nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) IsSetEncryptedPassword() bool {
	return p.EncryptedPassword != nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.AccountIdentifier = &AccountIdentifier{}
	if err := p.AccountIdentifier.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.AccountIdentifier), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.EncryptedPassword = &EncryptedPassword{}
	if err := p.EncryptedPassword.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.EncryptedPassword), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifyAccountUsingPwd_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountIdentifier", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:accountIdentifier: ", p), err)
	}
	if err := p.AccountIdentifier.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.AccountIdentifier), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:accountIdentifier: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "encryptedPassword", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:encryptedPassword: ", p), err)
	}
	if err := p.EncryptedPassword.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.EncryptedPassword), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:encryptedPassword: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceVerifyAccountUsingPwdArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceVerifyAccountUsingPwdResult struct {
	Success *VerifyAccountUsingPwdResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                 `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceVerifyAccountUsingPwdResult() *PrimaryAccountInitServiceVerifyAccountUsingPwdResult {
	return &PrimaryAccountInitServiceVerifyAccountUsingPwdResult{}
}

var PrimaryAccountInitServiceVerifyAccountUsingPwdResult_Success_DEFAULT *VerifyAccountUsingPwdResponse

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) GetSuccess() *VerifyAccountUsingPwdResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceVerifyAccountUsingPwdResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceVerifyAccountUsingPwdResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceVerifyAccountUsingPwdResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &VerifyAccountUsingPwdResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifyAccountUsingPwd_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyAccountUsingPwdResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceVerifyAccountUsingPwdResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - Device
//  - SocialLogin
type PrimaryAccountInitServiceVerifySocialLoginArgs struct {
	AuthSessionId string       `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	Device        *Device      `thrift:"device,2" db:"device" json:"device"`
	SocialLogin   *SocialLogin `thrift:"socialLogin,3" db:"socialLogin" json:"socialLogin"`
}

func NewPrimaryAccountInitServiceVerifySocialLoginArgs() *PrimaryAccountInitServiceVerifySocialLoginArgs {
	return &PrimaryAccountInitServiceVerifySocialLoginArgs{}
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceVerifySocialLoginArgs_Device_DEFAULT *Device

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) GetDevice() *Device {
	if !p.IsSetDevice() {
		return PrimaryAccountInitServiceVerifySocialLoginArgs_Device_DEFAULT
	}
	return p.Device
}

var PrimaryAccountInitServiceVerifySocialLoginArgs_SocialLogin_DEFAULT *SocialLogin

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) GetSocialLogin() *SocialLogin {
	if !p.IsSetSocialLogin() {
		return PrimaryAccountInitServiceVerifySocialLoginArgs_SocialLogin_DEFAULT
	}
	return p.SocialLogin
}
func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) IsSetDevice() bool {
	return p.Device != nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) IsSetSocialLogin() bool {
	return p.SocialLogin != nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Device = &Device{}
	if err := p.Device.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Device), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.SocialLogin = &SocialLogin{}
	if err := p.SocialLogin.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.SocialLogin), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifySocialLogin_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "device", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:device: ", p), err)
	}
	if err := p.Device.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Device), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:device: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "socialLogin", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:socialLogin: ", p), err)
	}
	if err := p.SocialLogin.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.SocialLogin), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:socialLogin: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifySocialLoginArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceVerifySocialLoginArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceVerifySocialLoginResult struct {
	Success *VerifySocialLoginResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException             `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceVerifySocialLoginResult() *PrimaryAccountInitServiceVerifySocialLoginResult {
	return &PrimaryAccountInitServiceVerifySocialLoginResult{}
}

var PrimaryAccountInitServiceVerifySocialLoginResult_Success_DEFAULT *VerifySocialLoginResponse

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) GetSuccess() *VerifySocialLoginResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceVerifySocialLoginResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceVerifySocialLoginResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceVerifySocialLoginResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceVerifySocialLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &VerifySocialLoginResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifySocialLogin_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifySocialLoginResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceVerifySocialLoginResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - Device
//  - UserPhoneNumber
type PrimaryAccountInitServiceGetPhoneVerifMethodArgs struct {
	AuthSessionId   string           `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	Device          *Device          `thrift:"device,2" db:"device" json:"device"`
	UserPhoneNumber *UserPhoneNumber `thrift:"userPhoneNumber,3" db:"userPhoneNumber" json:"userPhoneNumber"`
}

func NewPrimaryAccountInitServiceGetPhoneVerifMethodArgs() *PrimaryAccountInitServiceGetPhoneVerifMethodArgs {
	return &PrimaryAccountInitServiceGetPhoneVerifMethodArgs{}
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceGetPhoneVerifMethodArgs_Device_DEFAULT *Device

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) GetDevice() *Device {
	if !p.IsSetDevice() {
		return PrimaryAccountInitServiceGetPhoneVerifMethodArgs_Device_DEFAULT
	}
	return p.Device
}

var PrimaryAccountInitServiceGetPhoneVerifMethodArgs_UserPhoneNumber_DEFAULT *UserPhoneNumber

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) GetUserPhoneNumber() *UserPhoneNumber {
	if !p.IsSetUserPhoneNumber() {
		return PrimaryAccountInitServiceGetPhoneVerifMethodArgs_UserPhoneNumber_DEFAULT
	}
	return p.UserPhoneNumber
}
func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) IsSetDevice() bool {
	return p.Device != nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) IsSetUserPhoneNumber() bool {
	return p.UserPhoneNumber != nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Device = &Device{}
	if err := p.Device.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Device), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserPhoneNumber = &UserPhoneNumber{}
	if err := p.UserPhoneNumber.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserPhoneNumber), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getPhoneVerifMethod_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "device", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:device: ", p), err)
	}
	if err := p.Device.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Device), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:device: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userPhoneNumber", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:userPhoneNumber: ", p), err)
	}
	if err := p.UserPhoneNumber.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserPhoneNumber), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:userPhoneNumber: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetPhoneVerifMethodArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceGetPhoneVerifMethodResult struct {
	Success *GetPhoneVerifMethodResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException               `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceGetPhoneVerifMethodResult() *PrimaryAccountInitServiceGetPhoneVerifMethodResult {
	return &PrimaryAccountInitServiceGetPhoneVerifMethodResult{}
}

var PrimaryAccountInitServiceGetPhoneVerifMethodResult_Success_DEFAULT *GetPhoneVerifMethodResponse

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) GetSuccess() *GetPhoneVerifMethodResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceGetPhoneVerifMethodResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceGetPhoneVerifMethodResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceGetPhoneVerifMethodResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &GetPhoneVerifMethodResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getPhoneVerifMethod_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetPhoneVerifMethodResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - Device
//  - UserPhoneNumber
//  - VerifMethod
type PrimaryAccountInitServiceSendPinCodeForPhoneArgs struct {
	AuthSessionId   string           `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	Device          *Device          `thrift:"device,2" db:"device" json:"device"`
	UserPhoneNumber *UserPhoneNumber `thrift:"userPhoneNumber,3" db:"userPhoneNumber" json:"userPhoneNumber"`
	VerifMethod     VerifMethod      `thrift:"verifMethod,4" db:"verifMethod" json:"verifMethod"`
}

func NewPrimaryAccountInitServiceSendPinCodeForPhoneArgs() *PrimaryAccountInitServiceSendPinCodeForPhoneArgs {
	return &PrimaryAccountInitServiceSendPinCodeForPhoneArgs{}
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceSendPinCodeForPhoneArgs_Device_DEFAULT *Device

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) GetDevice() *Device {
	if !p.IsSetDevice() {
		return PrimaryAccountInitServiceSendPinCodeForPhoneArgs_Device_DEFAULT
	}
	return p.Device
}

var PrimaryAccountInitServiceSendPinCodeForPhoneArgs_UserPhoneNumber_DEFAULT *UserPhoneNumber

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) GetUserPhoneNumber() *UserPhoneNumber {
	if !p.IsSetUserPhoneNumber() {
		return PrimaryAccountInitServiceSendPinCodeForPhoneArgs_UserPhoneNumber_DEFAULT
	}
	return p.UserPhoneNumber
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) GetVerifMethod() VerifMethod {
	return p.VerifMethod
}
func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) IsSetDevice() bool {
	return p.Device != nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) IsSetUserPhoneNumber() bool {
	return p.UserPhoneNumber != nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Device = &Device{}
	if err := p.Device.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Device), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserPhoneNumber = &UserPhoneNumber{}
	if err := p.UserPhoneNumber.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserPhoneNumber), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		temp := VerifMethod(v)
		p.VerifMethod = temp
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "sendPinCodeForPhone_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "device", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:device: ", p), err)
	}
	if err := p.Device.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Device), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:device: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userPhoneNumber", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:userPhoneNumber: ", p), err)
	}
	if err := p.UserPhoneNumber.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserPhoneNumber), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:userPhoneNumber: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "verifMethod", thrift.I32, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:verifMethod: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.VerifMethod)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.verifMethod (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:verifMethod: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceSendPinCodeForPhoneArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceSendPinCodeForPhoneResult struct {
	Success *SendPinCodeForPhoneResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException               `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceSendPinCodeForPhoneResult() *PrimaryAccountInitServiceSendPinCodeForPhoneResult {
	return &PrimaryAccountInitServiceSendPinCodeForPhoneResult{}
}

var PrimaryAccountInitServiceSendPinCodeForPhoneResult_Success_DEFAULT *SendPinCodeForPhoneResponse

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) GetSuccess() *SendPinCodeForPhoneResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceSendPinCodeForPhoneResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceSendPinCodeForPhoneResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceSendPinCodeForPhoneResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &SendPinCodeForPhoneResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "sendPinCodeForPhone_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceSendPinCodeForPhoneResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceSendPinCodeForPhoneResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - Device
//  - UserPhoneNumber
//  - PinCode
type PrimaryAccountInitServiceVerifyPhoneArgs struct {
	AuthSessionId   string           `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	Device          *Device          `thrift:"device,2" db:"device" json:"device"`
	UserPhoneNumber *UserPhoneNumber `thrift:"userPhoneNumber,3" db:"userPhoneNumber" json:"userPhoneNumber"`
	PinCode         string           `thrift:"pinCode,4" db:"pinCode" json:"pinCode"`
}

func NewPrimaryAccountInitServiceVerifyPhoneArgs() *PrimaryAccountInitServiceVerifyPhoneArgs {
	return &PrimaryAccountInitServiceVerifyPhoneArgs{}
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceVerifyPhoneArgs_Device_DEFAULT *Device

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) GetDevice() *Device {
	if !p.IsSetDevice() {
		return PrimaryAccountInitServiceVerifyPhoneArgs_Device_DEFAULT
	}
	return p.Device
}

var PrimaryAccountInitServiceVerifyPhoneArgs_UserPhoneNumber_DEFAULT *UserPhoneNumber

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) GetUserPhoneNumber() *UserPhoneNumber {
	if !p.IsSetUserPhoneNumber() {
		return PrimaryAccountInitServiceVerifyPhoneArgs_UserPhoneNumber_DEFAULT
	}
	return p.UserPhoneNumber
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) GetPinCode() string {
	return p.PinCode
}
func (p *PrimaryAccountInitServiceVerifyPhoneArgs) IsSetDevice() bool {
	return p.Device != nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) IsSetUserPhoneNumber() bool {
	return p.UserPhoneNumber != nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Device = &Device{}
	if err := p.Device.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Device), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserPhoneNumber = &UserPhoneNumber{}
	if err := p.UserPhoneNumber.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserPhoneNumber), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.PinCode = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifyPhone_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "device", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:device: ", p), err)
	}
	if err := p.Device.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Device), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:device: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userPhoneNumber", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:userPhoneNumber: ", p), err)
	}
	if err := p.UserPhoneNumber.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserPhoneNumber), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:userPhoneNumber: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "pinCode", thrift.STRING, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:pinCode: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PinCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.pinCode (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:pinCode: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhoneArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceVerifyPhoneArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceVerifyPhoneResult struct {
	Success *VerifyPhoneResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException       `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceVerifyPhoneResult() *PrimaryAccountInitServiceVerifyPhoneResult {
	return &PrimaryAccountInitServiceVerifyPhoneResult{}
}

var PrimaryAccountInitServiceVerifyPhoneResult_Success_DEFAULT *VerifyPhoneResponse

func (p *PrimaryAccountInitServiceVerifyPhoneResult) GetSuccess() *VerifyPhoneResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceVerifyPhoneResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceVerifyPhoneResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceVerifyPhoneResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceVerifyPhoneResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceVerifyPhoneResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &VerifyPhoneResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifyPhone_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhoneResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhoneResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhoneResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceVerifyPhoneResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - EncryptedPassword
type PrimaryAccountInitServiceSetPasswordArgs struct {
	AuthSessionId     string             `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	EncryptedPassword *EncryptedPassword `thrift:"encryptedPassword,2" db:"encryptedPassword" json:"encryptedPassword"`
}

func NewPrimaryAccountInitServiceSetPasswordArgs() *PrimaryAccountInitServiceSetPasswordArgs {
	return &PrimaryAccountInitServiceSetPasswordArgs{}
}

func (p *PrimaryAccountInitServiceSetPasswordArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceSetPasswordArgs_EncryptedPassword_DEFAULT *EncryptedPassword

func (p *PrimaryAccountInitServiceSetPasswordArgs) GetEncryptedPassword() *EncryptedPassword {
	if !p.IsSetEncryptedPassword() {
		return PrimaryAccountInitServiceSetPasswordArgs_EncryptedPassword_DEFAULT
	}
	return p.EncryptedPassword
}
func (p *PrimaryAccountInitServiceSetPasswordArgs) IsSetEncryptedPassword() bool {
	return p.EncryptedPassword != nil
}

func (p *PrimaryAccountInitServiceSetPasswordArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSetPasswordArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceSetPasswordArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.EncryptedPassword = &EncryptedPassword{}
	if err := p.EncryptedPassword.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.EncryptedPassword), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSetPasswordArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "setPassword_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSetPasswordArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceSetPasswordArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "encryptedPassword", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:encryptedPassword: ", p), err)
	}
	if err := p.EncryptedPassword.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.EncryptedPassword), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:encryptedPassword: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceSetPasswordArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceSetPasswordArgs(%+v)", *p)
}

// Attributes:
//  - E
type PrimaryAccountInitServiceSetPasswordResult struct {
	E *AuthException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceSetPasswordResult() *PrimaryAccountInitServiceSetPasswordResult {
	return &PrimaryAccountInitServiceSetPasswordResult{}
}

var PrimaryAccountInitServiceSetPasswordResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceSetPasswordResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceSetPasswordResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceSetPasswordResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceSetPasswordResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSetPasswordResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSetPasswordResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "setPassword_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceSetPasswordResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceSetPasswordResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceSetPasswordResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - Request
type PrimaryAccountInitServiceExchangeEncryptionKeyArgs struct {
	AuthSessionId string                        `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	Request       *ExchangeEncryptionKeyRequest `thrift:"request,2" db:"request" json:"request"`
}

func NewPrimaryAccountInitServiceExchangeEncryptionKeyArgs() *PrimaryAccountInitServiceExchangeEncryptionKeyArgs {
	return &PrimaryAccountInitServiceExchangeEncryptionKeyArgs{}
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceExchangeEncryptionKeyArgs_Request_DEFAULT *ExchangeEncryptionKeyRequest

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) GetRequest() *ExchangeEncryptionKeyRequest {
	if !p.IsSetRequest() {
		return PrimaryAccountInitServiceExchangeEncryptionKeyArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &ExchangeEncryptionKeyRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "exchangeEncryptionKey_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:request: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceExchangeEncryptionKeyArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceExchangeEncryptionKeyResult struct {
	Success *ExchangeEncryptionKeyResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                 `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceExchangeEncryptionKeyResult() *PrimaryAccountInitServiceExchangeEncryptionKeyResult {
	return &PrimaryAccountInitServiceExchangeEncryptionKeyResult{}
}

var PrimaryAccountInitServiceExchangeEncryptionKeyResult_Success_DEFAULT *ExchangeEncryptionKeyResponse

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) GetSuccess() *ExchangeEncryptionKeyResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceExchangeEncryptionKeyResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceExchangeEncryptionKeyResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceExchangeEncryptionKeyResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &ExchangeEncryptionKeyResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "exchangeEncryptionKey_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceExchangeEncryptionKeyResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceExchangeEncryptionKeyResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - SimCard
type PrimaryAccountInitServiceGetCountryInfoArgs struct {
	AuthSessionId string   `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	SimCard       *SimCard `thrift:"simCard,2" db:"simCard" json:"simCard"`
}

func NewPrimaryAccountInitServiceGetCountryInfoArgs() *PrimaryAccountInitServiceGetCountryInfoArgs {
	return &PrimaryAccountInitServiceGetCountryInfoArgs{}
}

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceGetCountryInfoArgs_SimCard_DEFAULT *SimCard

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) GetSimCard() *SimCard {
	if !p.IsSetSimCard() {
		return PrimaryAccountInitServiceGetCountryInfoArgs_SimCard_DEFAULT
	}
	return p.SimCard
}
func (p *PrimaryAccountInitServiceGetCountryInfoArgs) IsSetSimCard() bool {
	return p.SimCard != nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.SimCard = &SimCard{}
	if err := p.SimCard.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.SimCard), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getCountryInfo_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "simCard", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:simCard: ", p), err)
	}
	if err := p.SimCard.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.SimCard), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:simCard: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetCountryInfoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetCountryInfoArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceGetCountryInfoResult struct {
	Success *GetCountryInfoResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException          `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceGetCountryInfoResult() *PrimaryAccountInitServiceGetCountryInfoResult {
	return &PrimaryAccountInitServiceGetCountryInfoResult{}
}

var PrimaryAccountInitServiceGetCountryInfoResult_Success_DEFAULT *GetCountryInfoResponse

func (p *PrimaryAccountInitServiceGetCountryInfoResult) GetSuccess() *GetCountryInfoResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceGetCountryInfoResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceGetCountryInfoResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceGetCountryInfoResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceGetCountryInfoResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceGetCountryInfoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &GetCountryInfoResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getCountryInfo_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetCountryInfoResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetCountryInfoResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetCountryInfoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetCountryInfoResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - AccountIdentifier
type PrimaryAccountInitServiceGetUserProfileArgs struct {
	AuthSessionId     string             `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	AccountIdentifier *AccountIdentifier `thrift:"accountIdentifier,2" db:"accountIdentifier" json:"accountIdentifier"`
}

func NewPrimaryAccountInitServiceGetUserProfileArgs() *PrimaryAccountInitServiceGetUserProfileArgs {
	return &PrimaryAccountInitServiceGetUserProfileArgs{}
}

func (p *PrimaryAccountInitServiceGetUserProfileArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceGetUserProfileArgs_AccountIdentifier_DEFAULT *AccountIdentifier

func (p *PrimaryAccountInitServiceGetUserProfileArgs) GetAccountIdentifier() *AccountIdentifier {
	if !p.IsSetAccountIdentifier() {
		return PrimaryAccountInitServiceGetUserProfileArgs_AccountIdentifier_DEFAULT
	}
	return p.AccountIdentifier
}
func (p *PrimaryAccountInitServiceGetUserProfileArgs) IsSetAccountIdentifier() bool {
	return p.AccountIdentifier != nil
}

func (p *PrimaryAccountInitServiceGetUserProfileArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetUserProfileArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetUserProfileArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.AccountIdentifier = &AccountIdentifier{}
	if err := p.AccountIdentifier.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.AccountIdentifier), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetUserProfileArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getUserProfile_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetUserProfileArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetUserProfileArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountIdentifier", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:accountIdentifier: ", p), err)
	}
	if err := p.AccountIdentifier.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.AccountIdentifier), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:accountIdentifier: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetUserProfileArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetUserProfileArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceGetUserProfileResult struct {
	Success *GetUserProfileResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException          `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceGetUserProfileResult() *PrimaryAccountInitServiceGetUserProfileResult {
	return &PrimaryAccountInitServiceGetUserProfileResult{}
}

var PrimaryAccountInitServiceGetUserProfileResult_Success_DEFAULT *GetUserProfileResponse

func (p *PrimaryAccountInitServiceGetUserProfileResult) GetSuccess() *GetUserProfileResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceGetUserProfileResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceGetUserProfileResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceGetUserProfileResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceGetUserProfileResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceGetUserProfileResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceGetUserProfileResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceGetUserProfileResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetUserProfileResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &GetUserProfileResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetUserProfileResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetUserProfileResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getUserProfile_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetUserProfileResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetUserProfileResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetUserProfileResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetUserProfileResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - AccountIdentifier
type PrimaryAccountInitServiceGetAcctVerifMethodArgs struct {
	AuthSessionId     string             `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	AccountIdentifier *AccountIdentifier `thrift:"accountIdentifier,2" db:"accountIdentifier" json:"accountIdentifier"`
}

func NewPrimaryAccountInitServiceGetAcctVerifMethodArgs() *PrimaryAccountInitServiceGetAcctVerifMethodArgs {
	return &PrimaryAccountInitServiceGetAcctVerifMethodArgs{}
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

var PrimaryAccountInitServiceGetAcctVerifMethodArgs_AccountIdentifier_DEFAULT *AccountIdentifier

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) GetAccountIdentifier() *AccountIdentifier {
	if !p.IsSetAccountIdentifier() {
		return PrimaryAccountInitServiceGetAcctVerifMethodArgs_AccountIdentifier_DEFAULT
	}
	return p.AccountIdentifier
}
func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) IsSetAccountIdentifier() bool {
	return p.AccountIdentifier != nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.AccountIdentifier = &AccountIdentifier{}
	if err := p.AccountIdentifier.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.AccountIdentifier), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getAcctVerifMethod_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountIdentifier", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:accountIdentifier: ", p), err)
	}
	if err := p.AccountIdentifier.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.AccountIdentifier), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:accountIdentifier: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetAcctVerifMethodArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceGetAcctVerifMethodResult struct {
	Success *GetAcctVerifMethodResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException              `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceGetAcctVerifMethodResult() *PrimaryAccountInitServiceGetAcctVerifMethodResult {
	return &PrimaryAccountInitServiceGetAcctVerifMethodResult{}
}

var PrimaryAccountInitServiceGetAcctVerifMethodResult_Success_DEFAULT *GetAcctVerifMethodResponse

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) GetSuccess() *GetAcctVerifMethodResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceGetAcctVerifMethodResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceGetAcctVerifMethodResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceGetAcctVerifMethodResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &GetAcctVerifMethodResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getAcctVerifMethod_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetAcctVerifMethodResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetAcctVerifMethodResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - DisplayName
type PrimaryAccountInitServiceValidateProfileArgs struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	DisplayName   string `thrift:"displayName,2" db:"displayName" json:"displayName"`
}

func NewPrimaryAccountInitServiceValidateProfileArgs() *PrimaryAccountInitServiceValidateProfileArgs {
	return &PrimaryAccountInitServiceValidateProfileArgs{}
}

func (p *PrimaryAccountInitServiceValidateProfileArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}

func (p *PrimaryAccountInitServiceValidateProfileArgs) GetDisplayName() string {
	return p.DisplayName
}
func (p *PrimaryAccountInitServiceValidateProfileArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceValidateProfileArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceValidateProfileArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.DisplayName = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceValidateProfileArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "validateProfile_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceValidateProfileArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceValidateProfileArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "displayName", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:displayName: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.DisplayName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.displayName (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:displayName: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceValidateProfileArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceValidateProfileArgs(%+v)", *p)
}

// Attributes:
//  - E
type PrimaryAccountInitServiceValidateProfileResult struct {
	E *AuthException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceValidateProfileResult() *PrimaryAccountInitServiceValidateProfileResult {
	return &PrimaryAccountInitServiceValidateProfileResult{}
}

var PrimaryAccountInitServiceValidateProfileResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceValidateProfileResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceValidateProfileResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceValidateProfileResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceValidateProfileResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceValidateProfileResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceValidateProfileResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "validateProfile_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceValidateProfileResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceValidateProfileResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceValidateProfileResult(%+v)", *p)
}

// Attributes:
//  - Request
type PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs struct {
	Request *FetchPhonePinCodeMsgRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryAccountInitServiceFetchPhonePinCodeMsgArgs() *PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs {
	return &PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs{}
}

var PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs_Request_DEFAULT *FetchPhonePinCodeMsgRequest

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs) GetRequest() *FetchPhonePinCodeMsgRequest {
	if !p.IsSetRequest() {
		return PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &FetchPhonePinCodeMsgRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "fetchPhonePinCodeMsg_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceFetchPhonePinCodeMsgArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceFetchPhonePinCodeMsgResult struct {
	Success *FetchPhonePinCodeMsgResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceFetchPhonePinCodeMsgResult() *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult {
	return &PrimaryAccountInitServiceFetchPhonePinCodeMsgResult{}
}

var PrimaryAccountInitServiceFetchPhonePinCodeMsgResult_Success_DEFAULT *FetchPhonePinCodeMsgResponse

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) GetSuccess() *FetchPhonePinCodeMsgResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceFetchPhonePinCodeMsgResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceFetchPhonePinCodeMsgResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceFetchPhonePinCodeMsgResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &FetchPhonePinCodeMsgResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "fetchPhonePinCodeMsg_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceFetchPhonePinCodeMsgResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceFetchPhonePinCodeMsgResult(%+v)", *p)
}

// Attributes:
//  - Request
type PrimaryAccountInitServiceGetPhoneVerifMethodV2Args struct {
	Request *GetPhoneVerifMethodV2Request `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryAccountInitServiceGetPhoneVerifMethodV2Args() *PrimaryAccountInitServiceGetPhoneVerifMethodV2Args {
	return &PrimaryAccountInitServiceGetPhoneVerifMethodV2Args{}
}

var PrimaryAccountInitServiceGetPhoneVerifMethodV2Args_Request_DEFAULT *GetPhoneVerifMethodV2Request

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Args) GetRequest() *GetPhoneVerifMethodV2Request {
	if !p.IsSetRequest() {
		return PrimaryAccountInitServiceGetPhoneVerifMethodV2Args_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Args) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Args) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Args) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &GetPhoneVerifMethodV2Request{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Args) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getPhoneVerifMethodV2_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Args) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Args) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetPhoneVerifMethodV2Args(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceGetPhoneVerifMethodV2Result struct {
	Success *GetPhoneVerifMethodV2Response `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                 `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceGetPhoneVerifMethodV2Result() *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result {
	return &PrimaryAccountInitServiceGetPhoneVerifMethodV2Result{}
}

var PrimaryAccountInitServiceGetPhoneVerifMethodV2Result_Success_DEFAULT *GetPhoneVerifMethodV2Response

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) GetSuccess() *GetPhoneVerifMethodV2Response {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceGetPhoneVerifMethodV2Result_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceGetPhoneVerifMethodV2Result_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceGetPhoneVerifMethodV2Result_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &GetPhoneVerifMethodV2Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getPhoneVerifMethodV2_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetPhoneVerifMethodV2Result) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetPhoneVerifMethodV2Result(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs() *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs {
	return &PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs{}
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getSessionContentBeforeMigCompletion_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult struct {
	Success *GetSessionContentBeforeMigCompletionResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                                `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult() *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult {
	return &PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult{}
}

var PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult_Success_DEFAULT *GetSessionContentBeforeMigCompletionResponse

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) GetSuccess() *GetSessionContentBeforeMigCompletionResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &GetSessionContentBeforeMigCompletionResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "getSessionContentBeforeMigCompletion_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceGetSessionContentBeforeMigCompletionResult(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args() *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args {
	return &PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args{}
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingEapAccountWithTokenV3_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Args(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result struct {
	Success *MigratePrimaryWithTokenV3Response `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                     `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result() *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result {
	return &PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result{}
}

var PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result_Success_DEFAULT *MigratePrimaryWithTokenV3Response

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) GetSuccess() *MigratePrimaryWithTokenV3Response {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &MigratePrimaryWithTokenV3Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingEapAccountWithTokenV3_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceMigratePrimaryUsingEapAccountWithTokenV3Result(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args() *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args {
	return &PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args{}
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingPhoneWithTokenV3_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Args(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result struct {
	Success *MigratePrimaryWithTokenV3Response `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                     `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result() *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result {
	return &PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result{}
}

var PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result_Success_DEFAULT *MigratePrimaryWithTokenV3Response

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) GetSuccess() *MigratePrimaryWithTokenV3Response {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &MigratePrimaryWithTokenV3Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingPhoneWithTokenV3_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceMigratePrimaryUsingPhoneWithTokenV3Result(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args struct {
	// unused field # 1
	AuthSessionId string `thrift:"authSessionId,2" db:"authSessionId" json:"authSessionId"`
}

func NewPrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args() *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args {
	return &PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args{}
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "registerPrimaryUsingPhoneWithTokenV3_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:authSessionId: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Args(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result struct {
	Success *RegisterPrimaryWithTokenV3Response `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                      `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result() *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result {
	return &PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result{}
}

var PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result_Success_DEFAULT *RegisterPrimaryWithTokenV3Response

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) GetSuccess() *RegisterPrimaryWithTokenV3Response {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &RegisterPrimaryWithTokenV3Response{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "registerPrimaryUsingPhoneWithTokenV3_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceRegisterPrimaryUsingPhoneWithTokenV3Result(%+v)", *p)
}

// Attributes:
//  - Request
type PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs struct {
	Request *ReqToSendPhonePinCodeRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryAccountInitServiceRequestToSendPhonePinCodeArgs() *PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs {
	return &PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs{}
}

var PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs_Request_DEFAULT *ReqToSendPhonePinCodeRequest

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs) GetRequest() *ReqToSendPhonePinCodeRequest {
	if !p.IsSetRequest() {
		return PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &ReqToSendPhonePinCodeRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "requestToSendPhonePinCode_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceRequestToSendPhonePinCodeArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceRequestToSendPhonePinCodeResult struct {
	Success *ReqToSendPhonePinCodeResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException                 `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceRequestToSendPhonePinCodeResult() *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult {
	return &PrimaryAccountInitServiceRequestToSendPhonePinCodeResult{}
}

var PrimaryAccountInitServiceRequestToSendPhonePinCodeResult_Success_DEFAULT *ReqToSendPhonePinCodeResponse

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) GetSuccess() *ReqToSendPhonePinCodeResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceRequestToSendPhonePinCodeResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceRequestToSendPhonePinCodeResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceRequestToSendPhonePinCodeResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &ReqToSendPhonePinCodeResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "requestToSendPhonePinCode_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceRequestToSendPhonePinCodeResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceRequestToSendPhonePinCodeResult(%+v)", *p)
}

// Attributes:
//  - Request
type PrimaryAccountInitServiceVerifyPhonePinCodeArgs struct {
	Request *VerifyPhonePinCodeRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryAccountInitServiceVerifyPhonePinCodeArgs() *PrimaryAccountInitServiceVerifyPhonePinCodeArgs {
	return &PrimaryAccountInitServiceVerifyPhonePinCodeArgs{}
}

var PrimaryAccountInitServiceVerifyPhonePinCodeArgs_Request_DEFAULT *VerifyPhonePinCodeRequest

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeArgs) GetRequest() *VerifyPhonePinCodeRequest {
	if !p.IsSetRequest() {
		return PrimaryAccountInitServiceVerifyPhonePinCodeArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryAccountInitServiceVerifyPhonePinCodeArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &VerifyPhonePinCodeRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifyPhonePinCode_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceVerifyPhonePinCodeArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceVerifyPhonePinCodeResult struct {
	Success *VerifyPhonePinCodeResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException              `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceVerifyPhonePinCodeResult() *PrimaryAccountInitServiceVerifyPhonePinCodeResult {
	return &PrimaryAccountInitServiceVerifyPhonePinCodeResult{}
}

var PrimaryAccountInitServiceVerifyPhonePinCodeResult_Success_DEFAULT *VerifyPhonePinCodeResponse

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) GetSuccess() *VerifyPhonePinCodeResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceVerifyPhonePinCodeResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceVerifyPhonePinCodeResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceVerifyPhonePinCodeResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &VerifyPhonePinCodeResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifyPhonePinCode_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceVerifyPhonePinCodeResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceVerifyPhonePinCodeResult(%+v)", *p)
}

// Attributes:
//  - ErrorCode
type AccessTokenRefreshException struct {
	ErrorCode int64 `thrift:"errorCode,1" db:"errorCode" json:"errorCode"`
}

func NewAccessTokenRefreshException() *AccessTokenRefreshException {
	return &AccessTokenRefreshException{}
}

func (p *AccessTokenRefreshException) GetErrorCode() int64 {
	return p.ErrorCode
}
func (p *AccessTokenRefreshException) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccessTokenRefreshException) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ErrorCode = v
	}
	return nil
}

func (p *AccessTokenRefreshException) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "AccessTokenRefreshException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccessTokenRefreshException) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "errorCode", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:errorCode: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.ErrorCode)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.errorCode (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:errorCode: ", p), err)
	}
	return err
}

func (p *AccessTokenRefreshException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccessTokenRefreshException(%+v)", *p)
}

func (p *AccessTokenRefreshException) Error() string {
	return p.String()
}

// Attributes:
//  - RefreshToken
type RefreshAccessTokenRequest struct {
	RefreshToken string `thrift:"refreshToken,1" db:"refreshToken" json:"refreshToken"`
}

func NewRefreshAccessTokenRequest() *RefreshAccessTokenRequest {
	return &RefreshAccessTokenRequest{}
}

func (p *RefreshAccessTokenRequest) GetRefreshToken() string {
	return p.RefreshToken
}
func (p *RefreshAccessTokenRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RefreshAccessTokenRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.RefreshToken = v
	}
	return nil
}

func (p *RefreshAccessTokenRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "RefreshAccessTokenRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RefreshAccessTokenRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "refreshToken", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:refreshToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.RefreshToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.refreshToken (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:refreshToken: ", p), err)
	}
	return err
}

func (p *RefreshAccessTokenRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefreshAccessTokenRequest(%+v)", *p)
}

// Attributes:
//  - AccessToken
//  - DurationUntilRefreshInSec
//  - RetryPolicy
//  - TokenIssueTimeEpochSec
//  - RefreshToken
type RefreshAccessTokenResponse struct {
	AccessToken               string                 `thrift:"accessToken,1" db:"accessToken" json:"accessToken"`
	DurationUntilRefreshInSec int64                  `thrift:"durationUntilRefreshInSec,2" db:"durationUntilRefreshInSec" json:"durationUntilRefreshInSec"`
	RetryPolicy               *RefreshApiRetryPolicy `thrift:"retryPolicy,3" db:"retryPolicy" json:"retryPolicy"`
	TokenIssueTimeEpochSec    int64                  `thrift:"tokenIssueTimeEpochSec,4" db:"tokenIssueTimeEpochSec" json:"tokenIssueTimeEpochSec"`
	RefreshToken              string                 `thrift:"refreshToken,5" db:"refreshToken" json:"refreshToken"`
}

func NewRefreshAccessTokenResponse() *RefreshAccessTokenResponse {
	return &RefreshAccessTokenResponse{}
}

func (p *RefreshAccessTokenResponse) GetAccessToken() string {
	return p.AccessToken
}

func (p *RefreshAccessTokenResponse) GetDurationUntilRefreshInSec() int64 {
	return p.DurationUntilRefreshInSec
}

var RefreshAccessTokenResponse_RetryPolicy_DEFAULT *RefreshApiRetryPolicy

func (p *RefreshAccessTokenResponse) GetRetryPolicy() *RefreshApiRetryPolicy {
	if !p.IsSetRetryPolicy() {
		return RefreshAccessTokenResponse_RetryPolicy_DEFAULT
	}
	return p.RetryPolicy
}

func (p *RefreshAccessTokenResponse) GetTokenIssueTimeEpochSec() int64 {
	return p.TokenIssueTimeEpochSec
}

func (p *RefreshAccessTokenResponse) GetRefreshToken() string {
	return p.RefreshToken
}
func (p *RefreshAccessTokenResponse) IsSetRetryPolicy() bool {
	return p.RetryPolicy != nil
}

func (p *RefreshAccessTokenResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField5(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RefreshAccessTokenResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AccessToken = v
	}
	return nil
}

func (p *RefreshAccessTokenResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.DurationUntilRefreshInSec = v
	}
	return nil
}

func (p *RefreshAccessTokenResponse) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.RetryPolicy = &RefreshApiRetryPolicy{}
	if err := p.RetryPolicy.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.RetryPolicy), err)
	}
	return nil
}

func (p *RefreshAccessTokenResponse) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.TokenIssueTimeEpochSec = v
	}
	return nil
}

func (p *RefreshAccessTokenResponse) ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.RefreshToken = v
	}
	return nil
}

func (p *RefreshAccessTokenResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "RefreshAccessTokenResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField5(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RefreshAccessTokenResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accessToken", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:accessToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AccessToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accessToken (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:accessToken: ", p), err)
	}
	return err
}

func (p *RefreshAccessTokenResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "durationUntilRefreshInSec", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:durationUntilRefreshInSec: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.DurationUntilRefreshInSec)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.durationUntilRefreshInSec (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:durationUntilRefreshInSec: ", p), err)
	}
	return err
}

func (p *RefreshAccessTokenResponse) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "retryPolicy", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:retryPolicy: ", p), err)
	}
	if err := p.RetryPolicy.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.RetryPolicy), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:retryPolicy: ", p), err)
	}
	return err
}

func (p *RefreshAccessTokenResponse) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "tokenIssueTimeEpochSec", thrift.I64, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:tokenIssueTimeEpochSec: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.TokenIssueTimeEpochSec)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.tokenIssueTimeEpochSec (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:tokenIssueTimeEpochSec: ", p), err)
	}
	return err
}

func (p *RefreshAccessTokenResponse) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "refreshToken", thrift.STRING, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:refreshToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.RefreshToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.refreshToken (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:refreshToken: ", p), err)
	}
	return err
}

func (p *RefreshAccessTokenResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RefreshAccessTokenResponse(%+v)", *p)
}

// Attributes:
//  - AccessToken
type ReportRefreshedAccessTokenRequest struct {
	AccessToken string `thrift:"accessToken,1" db:"accessToken" json:"accessToken"`
}

func NewReportRefreshedAccessTokenRequest() *ReportRefreshedAccessTokenRequest {
	return &ReportRefreshedAccessTokenRequest{}
}

func (p *ReportRefreshedAccessTokenRequest) GetAccessToken() string {
	return p.AccessToken
}
func (p *ReportRefreshedAccessTokenRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ReportRefreshedAccessTokenRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AccessToken = v
	}
	return nil
}

func (p *ReportRefreshedAccessTokenRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ReportRefreshedAccessTokenRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ReportRefreshedAccessTokenRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accessToken", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:accessToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AccessToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accessToken (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:accessToken: ", p), err)
	}
	return err
}

func (p *ReportRefreshedAccessTokenRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReportRefreshedAccessTokenRequest(%+v)", *p)
}

type AccessTokenRefreshService interface {
	// Parameters:
	//  - Request
	Refresh(ctx context.Context, request *RefreshAccessTokenRequest) (r *RefreshAccessTokenResponse, err error)
	// Parameters:
	//  - Request
	ReportRefreshedAccessToken(ctx context.Context, request *ReportRefreshedAccessTokenRequest) (err error)
}

type AccessTokenRefreshServiceClient struct {
	c thrift.TClient
}

func NewAccessTokenRefreshServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AccessTokenRefreshServiceClient {
	return &AccessTokenRefreshServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewAccessTokenRefreshServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AccessTokenRefreshServiceClient {
	return &AccessTokenRefreshServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewAccessTokenRefreshServiceClient(c thrift.TClient) *AccessTokenRefreshServiceClient {
	return &AccessTokenRefreshServiceClient{
		c: c,
	}
}

func (p *AccessTokenRefreshServiceClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - Request
func (p *AccessTokenRefreshServiceClient) Refresh(ctx context.Context, request *RefreshAccessTokenRequest) (r *RefreshAccessTokenResponse, err error) {
	var _args156 AccessTokenRefreshServiceRefreshArgs
	_args156.Request = request
	var _result157 AccessTokenRefreshServiceRefreshResult
	if _, err = p.Client_().Call(ctx, "refresh", &_args156, &_result157); err != nil {
		return
	}
	switch {
	case _result157.E != nil:
		return r, _result157.E
	}

	return _result157.GetSuccess(), nil
}

// Parameters:
//  - Request
func (p *AccessTokenRefreshServiceClient) ReportRefreshedAccessToken(ctx context.Context, request *ReportRefreshedAccessTokenRequest) (err error) {
	var _args158 AccessTokenRefreshServiceReportRefreshedAccessTokenArgs
	_args158.Request = request
	var _result159 AccessTokenRefreshServiceReportRefreshedAccessTokenResult
	if _, err = p.Client_().Call(ctx, "reportRefreshedAccessToken", &_args158, &_result159); err != nil {
		return
	}
	switch {
	case _result159.E != nil:
		return _result159.E
	}

	return nil
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type AccessTokenRefreshServiceRefreshArgs struct {
	Request *RefreshAccessTokenRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewAccessTokenRefreshServiceRefreshArgs() *AccessTokenRefreshServiceRefreshArgs {
	return &AccessTokenRefreshServiceRefreshArgs{}
}

var AccessTokenRefreshServiceRefreshArgs_Request_DEFAULT *RefreshAccessTokenRequest

func (p *AccessTokenRefreshServiceRefreshArgs) GetRequest() *RefreshAccessTokenRequest {
	if !p.IsSetRequest() {
		return AccessTokenRefreshServiceRefreshArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *AccessTokenRefreshServiceRefreshArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *AccessTokenRefreshServiceRefreshArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceRefreshArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &RefreshAccessTokenRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceRefreshArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "refresh_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceRefreshArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *AccessTokenRefreshServiceRefreshArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccessTokenRefreshServiceRefreshArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type AccessTokenRefreshServiceRefreshResult struct {
	Success *RefreshAccessTokenResponse  `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AccessTokenRefreshException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewAccessTokenRefreshServiceRefreshResult() *AccessTokenRefreshServiceRefreshResult {
	return &AccessTokenRefreshServiceRefreshResult{}
}

var AccessTokenRefreshServiceRefreshResult_Success_DEFAULT *RefreshAccessTokenResponse

func (p *AccessTokenRefreshServiceRefreshResult) GetSuccess() *RefreshAccessTokenResponse {
	if !p.IsSetSuccess() {
		return AccessTokenRefreshServiceRefreshResult_Success_DEFAULT
	}
	return p.Success
}

var AccessTokenRefreshServiceRefreshResult_E_DEFAULT *AccessTokenRefreshException

func (p *AccessTokenRefreshServiceRefreshResult) GetE() *AccessTokenRefreshException {
	if !p.IsSetE() {
		return AccessTokenRefreshServiceRefreshResult_E_DEFAULT
	}
	return p.E
}
func (p *AccessTokenRefreshServiceRefreshResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccessTokenRefreshServiceRefreshResult) IsSetE() bool {
	return p.E != nil
}

func (p *AccessTokenRefreshServiceRefreshResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceRefreshResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &RefreshAccessTokenResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceRefreshResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AccessTokenRefreshException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceRefreshResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "refresh_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceRefreshResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AccessTokenRefreshServiceRefreshResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *AccessTokenRefreshServiceRefreshResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccessTokenRefreshServiceRefreshResult(%+v)", *p)
}

// Attributes:
//  - Request
type AccessTokenRefreshServiceReportRefreshedAccessTokenArgs struct {
	Request *ReportRefreshedAccessTokenRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewAccessTokenRefreshServiceReportRefreshedAccessTokenArgs() *AccessTokenRefreshServiceReportRefreshedAccessTokenArgs {
	return &AccessTokenRefreshServiceReportRefreshedAccessTokenArgs{}
}

var AccessTokenRefreshServiceReportRefreshedAccessTokenArgs_Request_DEFAULT *ReportRefreshedAccessTokenRequest

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenArgs) GetRequest() *ReportRefreshedAccessTokenRequest {
	if !p.IsSetRequest() {
		return AccessTokenRefreshServiceReportRefreshedAccessTokenArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &ReportRefreshedAccessTokenRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "reportRefreshedAccessToken_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccessTokenRefreshServiceReportRefreshedAccessTokenArgs(%+v)", *p)
}

// Attributes:
//  - E
type AccessTokenRefreshServiceReportRefreshedAccessTokenResult struct {
	E *AccessTokenRefreshException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewAccessTokenRefreshServiceReportRefreshedAccessTokenResult() *AccessTokenRefreshServiceReportRefreshedAccessTokenResult {
	return &AccessTokenRefreshServiceReportRefreshedAccessTokenResult{}
}

var AccessTokenRefreshServiceReportRefreshedAccessTokenResult_E_DEFAULT *AccessTokenRefreshException

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenResult) GetE() *AccessTokenRefreshException {
	if !p.IsSetE() {
		return AccessTokenRefreshServiceReportRefreshedAccessTokenResult_E_DEFAULT
	}
	return p.E
}
func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenResult) IsSetE() bool {
	return p.E != nil
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AccessTokenRefreshException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "reportRefreshedAccessToken_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *AccessTokenRefreshServiceReportRefreshedAccessTokenResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccessTokenRefreshServiceReportRefreshedAccessTokenResult(%+v)", *p)
}

type AccountEapConnectErrorCode int64

const (
	AccountEapConnectErrorCode_INTERNAL_ERROR              AccountEapConnectErrorCode = 0
	AccountEapConnectErrorCode_ILLEGAL_ARGUMENT            AccountEapConnectErrorCode = 1
	AccountEapConnectErrorCode_VERIFICATION_FAILED         AccountEapConnectErrorCode = 2
	AccountEapConnectErrorCode_RETRY_LATER                 AccountEapConnectErrorCode = 4
	AccountEapConnectErrorCode_HUMAN_VERIFICATION_REQUIRED AccountEapConnectErrorCode = 5
	AccountEapConnectErrorCode_APP_UPGRADE_REQUIRED        AccountEapConnectErrorCode = 101
)

func (p AccountEapConnectErrorCode) String() string {
	switch p {
	case AccountEapConnectErrorCode_INTERNAL_ERROR:
		return "INTERNAL_ERROR"
	case AccountEapConnectErrorCode_ILLEGAL_ARGUMENT:
		return "ILLEGAL_ARGUMENT"
	case AccountEapConnectErrorCode_VERIFICATION_FAILED:
		return "VERIFICATION_FAILED"
	case AccountEapConnectErrorCode_RETRY_LATER:
		return "RETRY_LATER"
	case AccountEapConnectErrorCode_HUMAN_VERIFICATION_REQUIRED:
		return "HUMAN_VERIFICATION_REQUIRED"
	case AccountEapConnectErrorCode_APP_UPGRADE_REQUIRED:
		return "APP_UPGRADE_REQUIRED"
	}
	return "<UNSET>"
}

func AccountEapConnectErrorCodeFromString(s string) (AccountEapConnectErrorCode, error) {
	switch s {
	case "INTERNAL_ERROR":
		return AccountEapConnectErrorCode_INTERNAL_ERROR, nil
	case "ILLEGAL_ARGUMENT":
		return AccountEapConnectErrorCode_ILLEGAL_ARGUMENT, nil
	case "VERIFICATION_FAILED":
		return AccountEapConnectErrorCode_VERIFICATION_FAILED, nil
	case "RETRY_LATER":
		return AccountEapConnectErrorCode_RETRY_LATER, nil
	case "HUMAN_VERIFICATION_REQUIRED":
		return AccountEapConnectErrorCode_HUMAN_VERIFICATION_REQUIRED, nil
	case "APP_UPGRADE_REQUIRED":
		return AccountEapConnectErrorCode_APP_UPGRADE_REQUIRED, nil
	}
	return AccountEapConnectErrorCode(0), fmt.Errorf("not a valid AccountEapConnectErrorCode string")
}

func AccountEapConnectErrorCodePtr(v AccountEapConnectErrorCode) *AccountEapConnectErrorCode {
	return &v
}

func (p AccountEapConnectErrorCode) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *AccountEapConnectErrorCode) UnmarshalText(text []byte) error {
	q, err := AccountEapConnectErrorCodeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *AccountEapConnectErrorCode) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = AccountEapConnectErrorCode(v)
	return nil
}

func (p *AccountEapConnectErrorCode) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type EapType int64

const (
	EapType_UNKNOWN  EapType = 0
	EapType_FACEBOOK EapType = 1
	EapType_APPLE    EapType = 2
)

func (p EapType) String() string {
	switch p {
	case EapType_UNKNOWN:
		return "UNKNOWN"
	case EapType_FACEBOOK:
		return "FACEBOOK"
	case EapType_APPLE:
		return "APPLE"
	}
	return "<UNSET>"
}

func EapTypeFromString(s string) (EapType, error) {
	switch s {
	case "UNKNOWN":
		return EapType_UNKNOWN, nil
	case "FACEBOOK":
		return EapType_FACEBOOK, nil
	case "APPLE":
		return EapType_APPLE, nil
	}
	return EapType(0), fmt.Errorf("not a valid EapType string")
}

func EapTypePtr(v EapType) *EapType { return &v }

func (p EapType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *EapType) UnmarshalText(text []byte) error {
	q, err := EapTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *EapType) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = EapType(v)
	return nil
}

func (p *EapType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// Attributes:
//  - Code
//  - AlertMessage
//  - WebAuthDetails
type AccountEapConnectException struct {
	Code         AccountEapConnectErrorCode `thrift:"code,1" db:"code" json:"code"`
	AlertMessage string                     `thrift:"alertMessage,2" db:"alertMessage" json:"alertMessage"`
	// unused fields # 3 to 10
	WebAuthDetails *WebAuthDetails `thrift:"webAuthDetails,11" db:"webAuthDetails" json:"webAuthDetails"`
}

func NewAccountEapConnectException() *AccountEapConnectException {
	return &AccountEapConnectException{}
}

func (p *AccountEapConnectException) GetCode() AccountEapConnectErrorCode {
	return p.Code
}

func (p *AccountEapConnectException) GetAlertMessage() string {
	return p.AlertMessage
}

var AccountEapConnectException_WebAuthDetails_DEFAULT *WebAuthDetails

func (p *AccountEapConnectException) GetWebAuthDetails() *WebAuthDetails {
	if !p.IsSetWebAuthDetails() {
		return AccountEapConnectException_WebAuthDetails_DEFAULT
	}
	return p.WebAuthDetails
}
func (p *AccountEapConnectException) IsSetWebAuthDetails() bool {
	return p.WebAuthDetails != nil
}

func (p *AccountEapConnectException) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 11:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField11(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountEapConnectException) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := AccountEapConnectErrorCode(v)
		p.Code = temp
	}
	return nil
}

func (p *AccountEapConnectException) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AlertMessage = v
	}
	return nil
}

func (p *AccountEapConnectException) ReadField11(ctx context.Context, iprot thrift.TProtocol) error {
	p.WebAuthDetails = &WebAuthDetails{}
	if err := p.WebAuthDetails.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.WebAuthDetails), err)
	}
	return nil
}

func (p *AccountEapConnectException) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "AccountEapConnectException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField11(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountEapConnectException) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "code", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Code)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err)
	}
	return err
}

func (p *AccountEapConnectException) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "alertMessage", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:alertMessage: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AlertMessage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.alertMessage (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:alertMessage: ", p), err)
	}
	return err
}

func (p *AccountEapConnectException) writeField11(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "webAuthDetails", thrift.STRUCT, 11); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:webAuthDetails: ", p), err)
	}
	if err := p.WebAuthDetails.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.WebAuthDetails), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 11:webAuthDetails: ", p), err)
	}
	return err
}

func (p *AccountEapConnectException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountEapConnectException(%+v)", *p)
}

func (p *AccountEapConnectException) Error() string {
	return p.String()
}

// Attributes:
//  - AuthSessionId
type ConnectEapAccountRequest struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewConnectEapAccountRequest() *ConnectEapAccountRequest {
	return &ConnectEapAccountRequest{}
}

func (p *ConnectEapAccountRequest) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *ConnectEapAccountRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ConnectEapAccountRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *ConnectEapAccountRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ConnectEapAccountRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ConnectEapAccountRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *ConnectEapAccountRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ConnectEapAccountRequest(%+v)", *p)
}

// Attributes:
//  - EapType
type DisconnectEapAccountRequest struct {
	EapType EapType `thrift:"eapType,1" db:"eapType" json:"eapType"`
}

func NewDisconnectEapAccountRequest() *DisconnectEapAccountRequest {
	return &DisconnectEapAccountRequest{}
}

func (p *DisconnectEapAccountRequest) GetEapType() EapType {
	return p.EapType
}
func (p *DisconnectEapAccountRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *DisconnectEapAccountRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := EapType(v)
		p.EapType = temp
	}
	return nil
}

func (p *DisconnectEapAccountRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "DisconnectEapAccountRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *DisconnectEapAccountRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "eapType", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:eapType: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.EapType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.eapType (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:eapType: ", p), err)
	}
	return err
}

func (p *DisconnectEapAccountRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DisconnectEapAccountRequest(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type OpenSessionEapResponse struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewOpenSessionEapResponse() *OpenSessionEapResponse {
	return &OpenSessionEapResponse{}
}

func (p *OpenSessionEapResponse) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *OpenSessionEapResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *OpenSessionEapResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *OpenSessionEapResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "OpenSessionEapResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OpenSessionEapResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *OpenSessionEapResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OpenSessionEapResponse(%+v)", *p)
}

// Attributes:
//  - Device
type OpenSessionEapRequest struct {
	Device *Device `thrift:"device,1" db:"device" json:"device"`
}

func NewOpenSessionEapRequest() *OpenSessionEapRequest {
	return &OpenSessionEapRequest{}
}

var OpenSessionEapRequest_Device_DEFAULT *Device

func (p *OpenSessionEapRequest) GetDevice() *Device {
	if !p.IsSetDevice() {
		return OpenSessionEapRequest_Device_DEFAULT
	}
	return p.Device
}
func (p *OpenSessionEapRequest) IsSetDevice() bool {
	return p.Device != nil
}

func (p *OpenSessionEapRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *OpenSessionEapRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Device = &Device{}
	if err := p.Device.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Device), err)
	}
	return nil
}

func (p *OpenSessionEapRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "OpenSessionEapRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *OpenSessionEapRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "device", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:device: ", p), err)
	}
	if err := p.Device.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Device), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:device: ", p), err)
	}
	return err
}

func (p *OpenSessionEapRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("OpenSessionEapRequest(%+v)", *p)
}

// Attributes:
//  - AccountExists
type VerifyEapLoginResponse struct {
	AccountExists bool `thrift:"accountExists,1" db:"accountExists" json:"accountExists"`
}

func NewVerifyEapLoginResponse() *VerifyEapLoginResponse {
	return &VerifyEapLoginResponse{}
}

func (p *VerifyEapLoginResponse) GetAccountExists() bool {
	return p.AccountExists
}
func (p *VerifyEapLoginResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *VerifyEapLoginResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AccountExists = v
	}
	return nil
}

func (p *VerifyEapLoginResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "VerifyEapLoginResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifyEapLoginResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountExists", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:accountExists: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.AccountExists)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accountExists (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:accountExists: ", p), err)
	}
	return err
}

func (p *VerifyEapLoginResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifyEapLoginResponse(%+v)", *p)
}

// Attributes:
//  - Type
//  - AccessToken
type EapLogin struct {
	Type        EapType `thrift:"type,1" db:"type" json:"type"`
	AccessToken string  `thrift:"accessToken,2" db:"accessToken" json:"accessToken"`
}

func NewEapLogin() *EapLogin {
	return &EapLogin{}
}

func (p *EapLogin) GetType() EapType {
	return p.Type
}

func (p *EapLogin) GetAccessToken() string {
	return p.AccessToken
}
func (p *EapLogin) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EapLogin) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := EapType(v)
		p.Type = temp
	}
	return nil
}

func (p *EapLogin) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AccessToken = v
	}
	return nil
}

func (p *EapLogin) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "EapLogin"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EapLogin) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "type", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:type: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Type)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.type (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:type: ", p), err)
	}
	return err
}

func (p *EapLogin) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accessToken", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:accessToken: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AccessToken)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.accessToken (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:accessToken: ", p), err)
	}
	return err
}

func (p *EapLogin) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EapLogin(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
//  - EapLogin
type VerifyEapLoginRequest struct {
	AuthSessionId string    `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
	EapLogin      *EapLogin `thrift:"eapLogin,2" db:"eapLogin" json:"eapLogin"`
}

func NewVerifyEapLoginRequest() *VerifyEapLoginRequest {
	return &VerifyEapLoginRequest{}
}

func (p *VerifyEapLoginRequest) GetAuthSessionId() string {
	return p.AuthSessionId
}

var VerifyEapLoginRequest_EapLogin_DEFAULT *EapLogin

func (p *VerifyEapLoginRequest) GetEapLogin() *EapLogin {
	if !p.IsSetEapLogin() {
		return VerifyEapLoginRequest_EapLogin_DEFAULT
	}
	return p.EapLogin
}
func (p *VerifyEapLoginRequest) IsSetEapLogin() bool {
	return p.EapLogin != nil
}

func (p *VerifyEapLoginRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *VerifyEapLoginRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *VerifyEapLoginRequest) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.EapLogin = &EapLogin{}
	if err := p.EapLogin.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.EapLogin), err)
	}
	return nil
}

func (p *VerifyEapLoginRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "VerifyEapLoginRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *VerifyEapLoginRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *VerifyEapLoginRequest) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "eapLogin", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:eapLogin: ", p), err)
	}
	if err := p.EapLogin.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.EapLogin), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:eapLogin: ", p), err)
	}
	return err
}

func (p *VerifyEapLoginRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("VerifyEapLoginRequest(%+v)", *p)
}

type AccountAuthFactorEapConnectService interface {
	// Parameters:
	//  - Request
	ConnectEapAccount(ctx context.Context, request *ConnectEapAccountRequest) (err error)
	// Parameters:
	//  - Request
	DisconnectEapAccount(ctx context.Context, request *DisconnectEapAccountRequest) (err error)
	// Parameters:
	//  - Request
	OpenSession(ctx context.Context, request *OpenSessionEapRequest) (r *OpenSessionEapResponse, err error)
	// Parameters:
	//  - Request
	VerifyEapLogin(ctx context.Context, request *VerifyEapLoginRequest) (r *VerifyEapLoginResponse, err error)
}

type AccountAuthFactorEapConnectServiceClient struct {
	c thrift.TClient
}

func NewAccountAuthFactorEapConnectServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AccountAuthFactorEapConnectServiceClient {
	return &AccountAuthFactorEapConnectServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewAccountAuthFactorEapConnectServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AccountAuthFactorEapConnectServiceClient {
	return &AccountAuthFactorEapConnectServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewAccountAuthFactorEapConnectServiceClient(c thrift.TClient) *AccountAuthFactorEapConnectServiceClient {
	return &AccountAuthFactorEapConnectServiceClient{
		c: c,
	}
}

func (p *AccountAuthFactorEapConnectServiceClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - Request
func (p *AccountAuthFactorEapConnectServiceClient) ConnectEapAccount(ctx context.Context, request *ConnectEapAccountRequest) (err error) {
	var _args0 AccountAuthFactorEapConnectServiceConnectEapAccountArgs
	_args0.Request = request
	var _result1 AccountAuthFactorEapConnectServiceConnectEapAccountResult
	if _, err = p.Client_().Call(ctx, "connectEapAccount", &_args0, &_result1); err != nil {
		return
	}
	switch {
	case _result1.E != nil:
		return _result1.E
	}

	return nil
}

// Parameters:
//  - Request
func (p *AccountAuthFactorEapConnectServiceClient) DisconnectEapAccount(ctx context.Context, request *DisconnectEapAccountRequest) (err error) {
	var _args2 AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs
	_args2.Request = request
	var _result3 AccountAuthFactorEapConnectServiceDisconnectEapAccountResult
	if _, err = p.Client_().Call(ctx, "disconnectEapAccount", &_args2, &_result3); err != nil {
		return
	}
	switch {
	case _result3.E != nil:
		return _result3.E
	}

	return nil
}

// Parameters:
//  - Request
func (p *AccountAuthFactorEapConnectServiceClient) OpenSession(ctx context.Context, request *OpenSessionEapRequest) (r *OpenSessionEapResponse, err error) {
	var _args4 AccountAuthFactorEapConnectServiceOpenSessionArgs
	_args4.Request = request
	var _result5 AccountAuthFactorEapConnectServiceOpenSessionResult
	if _, err = p.Client_().Call(ctx, "openSession", &_args4, &_result5); err != nil {
		return
	}
	switch {
	case _result5.E != nil:
		return r, _result5.E
	}

	return _result5.GetSuccess(), nil
}

// Parameters:
//  - Request
func (p *AccountAuthFactorEapConnectServiceClient) VerifyEapLogin(ctx context.Context, request *VerifyEapLoginRequest) (r *VerifyEapLoginResponse, err error) {
	var _args6 AccountAuthFactorEapConnectServiceVerifyEapLoginArgs
	_args6.Request = request
	var _result7 AccountAuthFactorEapConnectServiceVerifyEapLoginResult
	if _, err = p.Client_().Call(ctx, "verifyEapLogin", &_args6, &_result7); err != nil {
		return
	}
	switch {
	case _result7.E != nil:
		return r, _result7.E
	}

	return _result7.GetSuccess(), nil
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type AccountAuthFactorEapConnectServiceConnectEapAccountArgs struct {
	Request *ConnectEapAccountRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewAccountAuthFactorEapConnectServiceConnectEapAccountArgs() *AccountAuthFactorEapConnectServiceConnectEapAccountArgs {
	return &AccountAuthFactorEapConnectServiceConnectEapAccountArgs{}
}

var AccountAuthFactorEapConnectServiceConnectEapAccountArgs_Request_DEFAULT *ConnectEapAccountRequest

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountArgs) GetRequest() *ConnectEapAccountRequest {
	if !p.IsSetRequest() {
		return AccountAuthFactorEapConnectServiceConnectEapAccountArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *AccountAuthFactorEapConnectServiceConnectEapAccountArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &ConnectEapAccountRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "connectEapAccount_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountAuthFactorEapConnectServiceConnectEapAccountArgs(%+v)", *p)
}

// Attributes:
//  - E
type AccountAuthFactorEapConnectServiceConnectEapAccountResult struct {
	E *AccountEapConnectException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewAccountAuthFactorEapConnectServiceConnectEapAccountResult() *AccountAuthFactorEapConnectServiceConnectEapAccountResult {
	return &AccountAuthFactorEapConnectServiceConnectEapAccountResult{}
}

var AccountAuthFactorEapConnectServiceConnectEapAccountResult_E_DEFAULT *AccountEapConnectException

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountResult) GetE() *AccountEapConnectException {
	if !p.IsSetE() {
		return AccountAuthFactorEapConnectServiceConnectEapAccountResult_E_DEFAULT
	}
	return p.E
}
func (p *AccountAuthFactorEapConnectServiceConnectEapAccountResult) IsSetE() bool {
	return p.E != nil
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AccountEapConnectException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "connectEapAccount_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceConnectEapAccountResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountAuthFactorEapConnectServiceConnectEapAccountResult(%+v)", *p)
}

// Attributes:
//  - Request
type AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs struct {
	Request *DisconnectEapAccountRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewAccountAuthFactorEapConnectServiceDisconnectEapAccountArgs() *AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs {
	return &AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs{}
}

var AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs_Request_DEFAULT *DisconnectEapAccountRequest

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs) GetRequest() *DisconnectEapAccountRequest {
	if !p.IsSetRequest() {
		return AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &DisconnectEapAccountRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "disconnectEapAccount_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountAuthFactorEapConnectServiceDisconnectEapAccountArgs(%+v)", *p)
}

// Attributes:
//  - E
type AccountAuthFactorEapConnectServiceDisconnectEapAccountResult struct {
	E *AccountEapConnectException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewAccountAuthFactorEapConnectServiceDisconnectEapAccountResult() *AccountAuthFactorEapConnectServiceDisconnectEapAccountResult {
	return &AccountAuthFactorEapConnectServiceDisconnectEapAccountResult{}
}

var AccountAuthFactorEapConnectServiceDisconnectEapAccountResult_E_DEFAULT *AccountEapConnectException

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountResult) GetE() *AccountEapConnectException {
	if !p.IsSetE() {
		return AccountAuthFactorEapConnectServiceDisconnectEapAccountResult_E_DEFAULT
	}
	return p.E
}
func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountResult) IsSetE() bool {
	return p.E != nil
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AccountEapConnectException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "disconnectEapAccount_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceDisconnectEapAccountResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountAuthFactorEapConnectServiceDisconnectEapAccountResult(%+v)", *p)
}

// Attributes:
//  - Request
type AccountAuthFactorEapConnectServiceOpenSessionArgs struct {
	Request *OpenSessionEapRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewAccountAuthFactorEapConnectServiceOpenSessionArgs() *AccountAuthFactorEapConnectServiceOpenSessionArgs {
	return &AccountAuthFactorEapConnectServiceOpenSessionArgs{}
}

var AccountAuthFactorEapConnectServiceOpenSessionArgs_Request_DEFAULT *OpenSessionEapRequest

func (p *AccountAuthFactorEapConnectServiceOpenSessionArgs) GetRequest() *OpenSessionEapRequest {
	if !p.IsSetRequest() {
		return AccountAuthFactorEapConnectServiceOpenSessionArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *AccountAuthFactorEapConnectServiceOpenSessionArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &OpenSessionEapRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "openSession_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountAuthFactorEapConnectServiceOpenSessionArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type AccountAuthFactorEapConnectServiceOpenSessionResult struct {
	Success *OpenSessionEapResponse     `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AccountEapConnectException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewAccountAuthFactorEapConnectServiceOpenSessionResult() *AccountAuthFactorEapConnectServiceOpenSessionResult {
	return &AccountAuthFactorEapConnectServiceOpenSessionResult{}
}

var AccountAuthFactorEapConnectServiceOpenSessionResult_Success_DEFAULT *OpenSessionEapResponse

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) GetSuccess() *OpenSessionEapResponse {
	if !p.IsSetSuccess() {
		return AccountAuthFactorEapConnectServiceOpenSessionResult_Success_DEFAULT
	}
	return p.Success
}

var AccountAuthFactorEapConnectServiceOpenSessionResult_E_DEFAULT *AccountEapConnectException

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) GetE() *AccountEapConnectException {
	if !p.IsSetE() {
		return AccountAuthFactorEapConnectServiceOpenSessionResult_E_DEFAULT
	}
	return p.E
}
func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) IsSetE() bool {
	return p.E != nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &OpenSessionEapResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AccountEapConnectException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "openSession_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceOpenSessionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountAuthFactorEapConnectServiceOpenSessionResult(%+v)", *p)
}

// Attributes:
//  - Request
type AccountAuthFactorEapConnectServiceVerifyEapLoginArgs struct {
	Request *VerifyEapLoginRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewAccountAuthFactorEapConnectServiceVerifyEapLoginArgs() *AccountAuthFactorEapConnectServiceVerifyEapLoginArgs {
	return &AccountAuthFactorEapConnectServiceVerifyEapLoginArgs{}
}

var AccountAuthFactorEapConnectServiceVerifyEapLoginArgs_Request_DEFAULT *VerifyEapLoginRequest

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginArgs) GetRequest() *VerifyEapLoginRequest {
	if !p.IsSetRequest() {
		return AccountAuthFactorEapConnectServiceVerifyEapLoginArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &VerifyEapLoginRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifyEapLogin_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountAuthFactorEapConnectServiceVerifyEapLoginArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type AccountAuthFactorEapConnectServiceVerifyEapLoginResult struct {
	Success *VerifyEapLoginResponse     `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AccountEapConnectException `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewAccountAuthFactorEapConnectServiceVerifyEapLoginResult() *AccountAuthFactorEapConnectServiceVerifyEapLoginResult {
	return &AccountAuthFactorEapConnectServiceVerifyEapLoginResult{}
}

var AccountAuthFactorEapConnectServiceVerifyEapLoginResult_Success_DEFAULT *VerifyEapLoginResponse

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) GetSuccess() *VerifyEapLoginResponse {
	if !p.IsSetSuccess() {
		return AccountAuthFactorEapConnectServiceVerifyEapLoginResult_Success_DEFAULT
	}
	return p.Success
}

var AccountAuthFactorEapConnectServiceVerifyEapLoginResult_E_DEFAULT *AccountEapConnectException

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) GetE() *AccountEapConnectException {
	if !p.IsSetE() {
		return AccountAuthFactorEapConnectServiceVerifyEapLoginResult_E_DEFAULT
	}
	return p.E
}
func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) IsSetE() bool {
	return p.E != nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &VerifyEapLoginResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AccountEapConnectException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "verifyEapLogin_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *AccountAuthFactorEapConnectServiceVerifyEapLoginResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AccountAuthFactorEapConnectServiceVerifyEapLoginResult(%+v)", *p)
}

type PrimaryQrCodeMigrationExceptionCode int64

const (
	PrimaryQrCodeMigrationExceptionCode_INTERNAL_ERROR   PrimaryQrCodeMigrationExceptionCode = 0
	PrimaryQrCodeMigrationExceptionCode_ILLEGAL_ARGUMENT PrimaryQrCodeMigrationExceptionCode = 1
	PrimaryQrCodeMigrationExceptionCode_NOT_FOUND        PrimaryQrCodeMigrationExceptionCode = 2
	PrimaryQrCodeMigrationExceptionCode_RETRY_LATER      PrimaryQrCodeMigrationExceptionCode = 3
	PrimaryQrCodeMigrationExceptionCode_INVALID_CONTEXT  PrimaryQrCodeMigrationExceptionCode = 100
	PrimaryQrCodeMigrationExceptionCode_NOT_SUPPORTED    PrimaryQrCodeMigrationExceptionCode = 101
)

func (p PrimaryQrCodeMigrationExceptionCode) String() string {
	switch p {
	case PrimaryQrCodeMigrationExceptionCode_INTERNAL_ERROR:
		return "INTERNAL_ERROR"
	case PrimaryQrCodeMigrationExceptionCode_ILLEGAL_ARGUMENT:
		return "ILLEGAL_ARGUMENT"
	case PrimaryQrCodeMigrationExceptionCode_NOT_FOUND:
		return "NOT_FOUND"
	case PrimaryQrCodeMigrationExceptionCode_RETRY_LATER:
		return "RETRY_LATER"
	case PrimaryQrCodeMigrationExceptionCode_INVALID_CONTEXT:
		return "INVALID_CONTEXT"
	case PrimaryQrCodeMigrationExceptionCode_NOT_SUPPORTED:
		return "NOT_SUPPORTED"
	}
	return "<UNSET>"
}

func PrimaryQrCodeMigrationExceptionCodeFromString(s string) (PrimaryQrCodeMigrationExceptionCode, error) {
	switch s {
	case "INTERNAL_ERROR":
		return PrimaryQrCodeMigrationExceptionCode_INTERNAL_ERROR, nil
	case "ILLEGAL_ARGUMENT":
		return PrimaryQrCodeMigrationExceptionCode_ILLEGAL_ARGUMENT, nil
	case "NOT_FOUND":
		return PrimaryQrCodeMigrationExceptionCode_NOT_FOUND, nil
	case "RETRY_LATER":
		return PrimaryQrCodeMigrationExceptionCode_RETRY_LATER, nil
	case "INVALID_CONTEXT":
		return PrimaryQrCodeMigrationExceptionCode_INVALID_CONTEXT, nil
	case "NOT_SUPPORTED":
		return PrimaryQrCodeMigrationExceptionCode_NOT_SUPPORTED, nil
	}
	return PrimaryQrCodeMigrationExceptionCode(0), fmt.Errorf("not a valid PrimaryQrCodeMigrationExceptionCode string")
}

func PrimaryQrCodeMigrationExceptionCodePtr(v PrimaryQrCodeMigrationExceptionCode) *PrimaryQrCodeMigrationExceptionCode {
	return &v
}

func (p PrimaryQrCodeMigrationExceptionCode) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *PrimaryQrCodeMigrationExceptionCode) UnmarshalText(text []byte) error {
	q, err := PrimaryQrCodeMigrationExceptionCodeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *PrimaryQrCodeMigrationExceptionCode) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = PrimaryQrCodeMigrationExceptionCode(v)
	return nil
}

func (p *PrimaryQrCodeMigrationExceptionCode) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type TokenAuthExceptionCode int64

const (
	TokenAuthExceptionCode_AUTHENTICATION_FAILED TokenAuthExceptionCode = 1
	TokenAuthExceptionCode_INVALID_STATE         TokenAuthExceptionCode = 2
	TokenAuthExceptionCode_NOT_AUTHORIZED_DEVICE TokenAuthExceptionCode = 3
	TokenAuthExceptionCode_MUST_REFRESH_V3_TOKEN TokenAuthExceptionCode = 4
)

func (p TokenAuthExceptionCode) String() string {
	switch p {
	case TokenAuthExceptionCode_AUTHENTICATION_FAILED:
		return "AUTHENTICATION_FAILED"
	case TokenAuthExceptionCode_INVALID_STATE:
		return "INVALID_STATE"
	case TokenAuthExceptionCode_NOT_AUTHORIZED_DEVICE:
		return "NOT_AUTHORIZED_DEVICE"
	case TokenAuthExceptionCode_MUST_REFRESH_V3_TOKEN:
		return "MUST_REFRESH_V3_TOKEN"
	}
	return "<UNSET>"
}

func TokenAuthExceptionCodeFromString(s string) (TokenAuthExceptionCode, error) {
	switch s {
	case "AUTHENTICATION_FAILED":
		return TokenAuthExceptionCode_AUTHENTICATION_FAILED, nil
	case "INVALID_STATE":
		return TokenAuthExceptionCode_INVALID_STATE, nil
	case "NOT_AUTHORIZED_DEVICE":
		return TokenAuthExceptionCode_NOT_AUTHORIZED_DEVICE, nil
	case "MUST_REFRESH_V3_TOKEN":
		return TokenAuthExceptionCode_MUST_REFRESH_V3_TOKEN, nil
	}
	return TokenAuthExceptionCode(0), fmt.Errorf("not a valid TokenAuthExceptionCode string")
}

func TokenAuthExceptionCodePtr(v TokenAuthExceptionCode) *TokenAuthExceptionCode { return &v }

func (p TokenAuthExceptionCode) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *TokenAuthExceptionCode) UnmarshalText(text []byte) error {
	q, err := TokenAuthExceptionCodeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

func (p *TokenAuthExceptionCode) Scan(value interface{}) error {
	v, ok := value.(int64)
	if !ok {
		return errors.New("Scan value is not int64")
	}
	*p = TokenAuthExceptionCode(v)
	return nil
}

func (p *TokenAuthExceptionCode) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

// Attributes:
//  - Code
//  - Reason
type TokenAuthException struct {
	Code   TokenAuthExceptionCode `thrift:"code,1" db:"code" json:"code"`
	Reason string                 `thrift:"reason,2" db:"reason" json:"reason"`
}

func NewTokenAuthException() *TokenAuthException {
	return &TokenAuthException{}
}

func (p *TokenAuthException) GetCode() TokenAuthExceptionCode {
	return p.Code
}

func (p *TokenAuthException) GetReason() string {
	return p.Reason
}
func (p *TokenAuthException) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenAuthException) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := TokenAuthExceptionCode(v)
		p.Code = temp
	}
	return nil
}

func (p *TokenAuthException) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Reason = v
	}
	return nil
}

func (p *TokenAuthException) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "TokenAuthException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenAuthException) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "code", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Code)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err)
	}
	return err
}

func (p *TokenAuthException) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "reason", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:reason: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Reason)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.reason (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:reason: ", p), err)
	}
	return err
}

func (p *TokenAuthException) Equals(other *TokenAuthException) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Code != other.Code {
		return false
	}
	if p.Reason != other.Reason {
		return false
	}
	return true
}

func (p *TokenAuthException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenAuthException(%+v)", *p)
}

func (p *TokenAuthException) Error() string {
	return p.String()
}

func (TokenAuthException) TExceptionType() thrift.TExceptionType {
	return thrift.TExceptionTypeCompiled
}

var _ thrift.TException = (*TokenAuthException)(nil)

// Attributes:
//  - Code
//  - ErrorMessage
type PrimaryQrCodeMigrationException struct {
	Code         PrimaryQrCodeMigrationExceptionCode `thrift:"code,1" db:"code" json:"code"`
	ErrorMessage string                              `thrift:"errorMessage,2" db:"errorMessage" json:"errorMessage"`
}

func NewPrimaryQrCodeMigrationException() *PrimaryQrCodeMigrationException {
	return &PrimaryQrCodeMigrationException{}
}

func (p *PrimaryQrCodeMigrationException) GetCode() PrimaryQrCodeMigrationExceptionCode {
	return p.Code
}

func (p *PrimaryQrCodeMigrationException) GetErrorMessage() string {
	return p.ErrorMessage
}
func (p *PrimaryQrCodeMigrationException) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationException) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		temp := PrimaryQrCodeMigrationExceptionCode(v)
		p.Code = temp
	}
	return nil
}

func (p *PrimaryQrCodeMigrationException) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.ErrorMessage = v
	}
	return nil
}

func (p *PrimaryQrCodeMigrationException) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "PrimaryQrCodeMigrationException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationException) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "code", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.Code)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err)
	}
	return err
}

func (p *PrimaryQrCodeMigrationException) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "errorMessage", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:errorMessage: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.ErrorMessage)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.errorMessage (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:errorMessage: ", p), err)
	}
	return err
}

func (p *PrimaryQrCodeMigrationException) Equals(other *PrimaryQrCodeMigrationException) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Code != other.Code {
		return false
	}
	if p.ErrorMessage != other.ErrorMessage {
		return false
	}
	return true
}

func (p *PrimaryQrCodeMigrationException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationException(%+v)", *p)
}

func (p *PrimaryQrCodeMigrationException) Error() string {
	return p.String()
}

func (PrimaryQrCodeMigrationException) TExceptionType() thrift.TExceptionType {
	return thrift.TExceptionTypeCompiled
}

var _ thrift.TException = (*PrimaryQrCodeMigrationException)(nil)

// Attributes:
//  - RecoveryKey
//  - BackupBlobPayload
type EncryptedBinary struct {
	RecoveryKey       []byte `thrift:"recoveryKey,1" db:"recoveryKey" json:"recoveryKey"`
	BackupBlobPayload []byte `thrift:"backupBlobPayload,2" db:"backupBlobPayload" json:"backupBlobPayload"`
}

func NewEncryptedBinary() *EncryptedBinary {
	return &EncryptedBinary{}
}

func (p *EncryptedBinary) GetRecoveryKey() []byte {
	return p.RecoveryKey
}

func (p *EncryptedBinary) GetBackupBlobPayload() []byte {
	return p.BackupBlobPayload
}
func (p *EncryptedBinary) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EncryptedBinary) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.RecoveryKey = v
	}
	return nil
}

func (p *EncryptedBinary) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.BackupBlobPayload = v
	}
	return nil
}

func (p *EncryptedBinary) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "EncryptedBinary"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EncryptedBinary) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "recoveryKey", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:recoveryKey: ", p), err)
	}
	if err := oprot.WriteBinary(ctx, p.RecoveryKey); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.recoveryKey (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:recoveryKey: ", p), err)
	}
	return err
}

func (p *EncryptedBinary) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "backupBlobPayload", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:backupBlobPayload: ", p), err)
	}
	if err := oprot.WriteBinary(ctx, p.BackupBlobPayload); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.backupBlobPayload (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:backupBlobPayload: ", p), err)
	}
	return err
}

func (p *EncryptedBinary) Equals(other *EncryptedBinary) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	return true
}

func (p *EncryptedBinary) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EncryptedBinary(%+v)", *p)
}

// Attributes:
//  - SessionId
//  - EncryptedSecureChannelPayload
type SendEncryptedE2EEKeyRequest struct {
	SessionId                     string           `thrift:"sessionId,1" db:"sessionId" json:"sessionId"`
	EncryptedSecureChannelPayload *EncryptedBinary `thrift:"encryptedSecureChannelPayload,2" db:"encryptedSecureChannelPayload" json:"encryptedSecureChannelPayload"`
}

func NewSendEncryptedE2EEKeyRequest() *SendEncryptedE2EEKeyRequest {
	return &SendEncryptedE2EEKeyRequest{}
}

func (p *SendEncryptedE2EEKeyRequest) GetSessionId() string {
	return p.SessionId
}

var SendEncryptedE2EEKeyRequest_EncryptedSecureChannelPayload_DEFAULT *EncryptedBinary

func (p *SendEncryptedE2EEKeyRequest) GetEncryptedSecureChannelPayload() *EncryptedBinary {
	if !p.IsSetEncryptedSecureChannelPayload() {
		return SendEncryptedE2EEKeyRequest_EncryptedSecureChannelPayload_DEFAULT
	}
	return p.EncryptedSecureChannelPayload
}
func (p *SendEncryptedE2EEKeyRequest) IsSetEncryptedSecureChannelPayload() bool {
	return p.EncryptedSecureChannelPayload != nil
}

func (p *SendEncryptedE2EEKeyRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SendEncryptedE2EEKeyRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.SessionId = v
	}
	return nil
}

func (p *SendEncryptedE2EEKeyRequest) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.EncryptedSecureChannelPayload = &EncryptedBinary{}
	if err := p.EncryptedSecureChannelPayload.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.EncryptedSecureChannelPayload), err)
	}
	return nil
}

func (p *SendEncryptedE2EEKeyRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "SendEncryptedE2EEKeyRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SendEncryptedE2EEKeyRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:sessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.SessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:sessionId: ", p), err)
	}
	return err
}

func (p *SendEncryptedE2EEKeyRequest) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "encryptedSecureChannelPayload", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:encryptedSecureChannelPayload: ", p), err)
	}
	if err := p.EncryptedSecureChannelPayload.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.EncryptedSecureChannelPayload), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:encryptedSecureChannelPayload: ", p), err)
	}
	return err
}

func (p *SendEncryptedE2EEKeyRequest) Equals(other *SendEncryptedE2EEKeyRequest) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.SessionId != other.SessionId {
		return false
	}
	if !p.EncryptedSecureChannelPayload.Equals(other.EncryptedSecureChannelPayload) {
		return false
	}
	return true
}

func (p *SendEncryptedE2EEKeyRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SendEncryptedE2EEKeyRequest(%+v)", *p)
}

type CreateSessionRequest struct {
}

func NewCreateSessionRequest() *CreateSessionRequest {
	return &CreateSessionRequest{}
}

func (p *CreateSessionRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(ctx, fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CreateSessionRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "CreateSessionRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CreateSessionRequest) Equals(other *CreateSessionRequest) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	return true
}

func (p *CreateSessionRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateSessionRequest(%+v)", *p)
}

// Attributes:
//  - SessionId
type CreateQRMigrationSessionResponse struct {
	SessionId string `thrift:"sessionId,1" db:"sessionId" json:"sessionId"`
}

func NewCreateQRMigrationSessionResponse() *CreateQRMigrationSessionResponse {
	return &CreateQRMigrationSessionResponse{}
}

func (p *CreateQRMigrationSessionResponse) GetSessionId() string {
	return p.SessionId
}
func (p *CreateQRMigrationSessionResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CreateQRMigrationSessionResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.SessionId = v
	}
	return nil
}

func (p *CreateQRMigrationSessionResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "CreateQRMigrationSessionResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CreateQRMigrationSessionResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:sessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.SessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:sessionId: ", p), err)
	}
	return err
}

func (p *CreateQRMigrationSessionResponse) Equals(other *CreateQRMigrationSessionResponse) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.SessionId != other.SessionId {
		return false
	}
	return true
}

func (p *CreateQRMigrationSessionResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CreateQRMigrationSessionResponse(%+v)", *p)
}

// Attributes:
//  - NewDevicePublicKey_
//  - EncryptedQrIdentifier
type SecureChannelData struct {
	NewDevicePublicKey_   []byte `thrift:"newDevicePublicKey,1" db:"newDevicePublicKey" json:"newDevicePublicKey"`
	EncryptedQrIdentifier string `thrift:"encryptedQrIdentifier,2" db:"encryptedQrIdentifier" json:"encryptedQrIdentifier"`
}

func NewSecureChannelData() *SecureChannelData {
	return &SecureChannelData{}
}

func (p *SecureChannelData) GetNewDevicePublicKey_() []byte {
	return p.NewDevicePublicKey_
}

func (p *SecureChannelData) GetEncryptedQrIdentifier() string {
	return p.EncryptedQrIdentifier
}
func (p *SecureChannelData) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SecureChannelData) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.NewDevicePublicKey_ = v
	}
	return nil
}

func (p *SecureChannelData) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.EncryptedQrIdentifier = v
	}
	return nil
}

func (p *SecureChannelData) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "SecureChannelData"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SecureChannelData) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "newDevicePublicKey", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:newDevicePublicKey: ", p), err)
	}
	if err := oprot.WriteBinary(ctx, p.NewDevicePublicKey_); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.newDevicePublicKey (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:newDevicePublicKey: ", p), err)
	}
	return err
}

func (p *SecureChannelData) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "encryptedQrIdentifier", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:encryptedQrIdentifier: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.EncryptedQrIdentifier)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.encryptedQrIdentifier (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:encryptedQrIdentifier: ", p), err)
	}
	return err
}

func (p *SecureChannelData) Equals(other *SecureChannelData) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.EncryptedQrIdentifier != other.EncryptedQrIdentifier {
		return false
	}
	return true
}

func (p *SecureChannelData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SecureChannelData(%+v)", *p)
}

// Attributes:
//  - SessionId
//  - SecureChannelData
type CheckIfEncryptedE2EEKeyReceivedRequest struct {
	SessionId         string             `thrift:"sessionId,1" db:"sessionId" json:"sessionId"`
	SecureChannelData *SecureChannelData `thrift:"secureChannelData,2" db:"secureChannelData" json:"secureChannelData"`
}

func NewCheckIfEncryptedE2EEKeyReceivedRequest() *CheckIfEncryptedE2EEKeyReceivedRequest {
	return &CheckIfEncryptedE2EEKeyReceivedRequest{}
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) GetSessionId() string {
	return p.SessionId
}

var CheckIfEncryptedE2EEKeyReceivedRequest_SecureChannelData_DEFAULT *SecureChannelData

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) GetSecureChannelData() *SecureChannelData {
	if !p.IsSetSecureChannelData() {
		return CheckIfEncryptedE2EEKeyReceivedRequest_SecureChannelData_DEFAULT
	}
	return p.SecureChannelData
}
func (p *CheckIfEncryptedE2EEKeyReceivedRequest) IsSetSecureChannelData() bool {
	return p.SecureChannelData != nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.SessionId = v
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.SecureChannelData = &SecureChannelData{}
	if err := p.SecureChannelData.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.SecureChannelData), err)
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "CheckIfEncryptedE2EEKeyReceivedRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:sessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.SessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:sessionId: ", p), err)
	}
	return err
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "secureChannelData", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:secureChannelData: ", p), err)
	}
	if err := p.SecureChannelData.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.SecureChannelData), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:secureChannelData: ", p), err)
	}
	return err
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) Equals(other *CheckIfEncryptedE2EEKeyReceivedRequest) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.SessionId != other.SessionId {
		return false
	}
	if !p.SecureChannelData.Equals(other.SecureChannelData) {
		return false
	}
	return true
}

func (p *CheckIfEncryptedE2EEKeyReceivedRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CheckIfEncryptedE2EEKeyReceivedRequest(%+v)", *p)
}

// Attributes:
//  - DisplayName
//  - ProfileImageUrl
type UserProfile struct {
	DisplayName     string `thrift:"displayName,1" db:"displayName" json:"displayName"`
	ProfileImageUrl string `thrift:"profileImageUrl,2" db:"profileImageUrl" json:"profileImageUrl"`
}

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}

func (p *UserProfile) GetDisplayName() string {
	return p.DisplayName
}

func (p *UserProfile) GetProfileImageUrl() string {
	return p.ProfileImageUrl
}
func (p *UserProfile) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *UserProfile) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.DisplayName = v
	}
	return nil
}

func (p *UserProfile) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.ProfileImageUrl = v
	}
	return nil
}

func (p *UserProfile) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "UserProfile"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *UserProfile) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "displayName", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:displayName: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.DisplayName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.displayName (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:displayName: ", p), err)
	}
	return err
}

func (p *UserProfile) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "profileImageUrl", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:profileImageUrl: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.ProfileImageUrl)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.profileImageUrl (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:profileImageUrl: ", p), err)
	}
	return err
}

func (p *UserProfile) Equals(other *UserProfile) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.DisplayName != other.DisplayName {
		return false
	}
	if p.ProfileImageUrl != other.ProfileImageUrl {
		return false
	}
	return true
}

func (p *UserProfile) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UserProfile(%+v)", *p)
}

// Attributes:
//  - Nonce
//  - EncryptedSecureChannelPayload
//  - UserProfile
//  - AppTypeDifferentFromPrevDevice
//  - E2eeKeyBackupServiceConfig
type CheckIfEncryptedE2EEKeyReceivedResponse struct {
	Nonce                          string           `thrift:"nonce,1" db:"nonce" json:"nonce"`
	EncryptedSecureChannelPayload  *EncryptedBinary `thrift:"encryptedSecureChannelPayload,2" db:"encryptedSecureChannelPayload" json:"encryptedSecureChannelPayload"`
	UserProfile                    *UserProfile     `thrift:"userProfile,3" db:"userProfile" json:"userProfile"`
	AppTypeDifferentFromPrevDevice bool             `thrift:"appTypeDifferentFromPrevDevice,4" db:"appTypeDifferentFromPrevDevice" json:"appTypeDifferentFromPrevDevice"`
	E2eeKeyBackupServiceConfig     bool             `thrift:"e2eeKeyBackupServiceConfig,5" db:"e2eeKeyBackupServiceConfig" json:"e2eeKeyBackupServiceConfig"`
}

func NewCheckIfEncryptedE2EEKeyReceivedResponse() *CheckIfEncryptedE2EEKeyReceivedResponse {
	return &CheckIfEncryptedE2EEKeyReceivedResponse{}
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) GetNonce() string {
	return p.Nonce
}

var CheckIfEncryptedE2EEKeyReceivedResponse_EncryptedSecureChannelPayload_DEFAULT *EncryptedBinary

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) GetEncryptedSecureChannelPayload() *EncryptedBinary {
	if !p.IsSetEncryptedSecureChannelPayload() {
		return CheckIfEncryptedE2EEKeyReceivedResponse_EncryptedSecureChannelPayload_DEFAULT
	}
	return p.EncryptedSecureChannelPayload
}

var CheckIfEncryptedE2EEKeyReceivedResponse_UserProfile_DEFAULT *UserProfile

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) GetUserProfile() *UserProfile {
	if !p.IsSetUserProfile() {
		return CheckIfEncryptedE2EEKeyReceivedResponse_UserProfile_DEFAULT
	}
	return p.UserProfile
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) GetAppTypeDifferentFromPrevDevice() bool {
	return p.AppTypeDifferentFromPrevDevice
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) GetE2eeKeyBackupServiceConfig() bool {
	return p.E2eeKeyBackupServiceConfig
}
func (p *CheckIfEncryptedE2EEKeyReceivedResponse) IsSetEncryptedSecureChannelPayload() bool {
	return p.EncryptedSecureChannelPayload != nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) IsSetUserProfile() bool {
	return p.UserProfile != nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField5(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Nonce = v
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.EncryptedSecureChannelPayload = &EncryptedBinary{}
	if err := p.EncryptedSecureChannelPayload.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.EncryptedSecureChannelPayload), err)
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.UserProfile = &UserProfile{}
	if err := p.UserProfile.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.UserProfile), err)
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.AppTypeDifferentFromPrevDevice = v
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 5: ", err)
	} else {
		p.E2eeKeyBackupServiceConfig = v
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "CheckIfEncryptedE2EEKeyReceivedResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField5(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "nonce", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:nonce: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Nonce)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nonce (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:nonce: ", p), err)
	}
	return err
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "encryptedSecureChannelPayload", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:encryptedSecureChannelPayload: ", p), err)
	}
	if err := p.EncryptedSecureChannelPayload.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.EncryptedSecureChannelPayload), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:encryptedSecureChannelPayload: ", p), err)
	}
	return err
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "userProfile", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:userProfile: ", p), err)
	}
	if err := p.UserProfile.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.UserProfile), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:userProfile: ", p), err)
	}
	return err
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "appTypeDifferentFromPrevDevice", thrift.BOOL, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:appTypeDifferentFromPrevDevice: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.AppTypeDifferentFromPrevDevice)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.appTypeDifferentFromPrevDevice (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:appTypeDifferentFromPrevDevice: ", p), err)
	}
	return err
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "e2eeKeyBackupServiceConfig", thrift.BOOL, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:e2eeKeyBackupServiceConfig: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.E2eeKeyBackupServiceConfig)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.e2eeKeyBackupServiceConfig (5) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:e2eeKeyBackupServiceConfig: ", p), err)
	}
	return err
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) Equals(other *CheckIfEncryptedE2EEKeyReceivedResponse) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Nonce != other.Nonce {
		return false
	}
	if !p.EncryptedSecureChannelPayload.Equals(other.EncryptedSecureChannelPayload) {
		return false
	}
	if !p.UserProfile.Equals(other.UserProfile) {
		return false
	}
	if p.AppTypeDifferentFromPrevDevice != other.AppTypeDifferentFromPrevDevice {
		return false
	}
	if p.E2eeKeyBackupServiceConfig != other.E2eeKeyBackupServiceConfig {
		return false
	}
	return true
}

func (p *CheckIfEncryptedE2EEKeyReceivedResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CheckIfEncryptedE2EEKeyReceivedResponse(%+v)", *p)
}

type PrimaryQrCodeMigrationPreparationService interface {
	// Parameters:
	//  - Request
	CreateQRMigrationSession(ctx context.Context, request *CreateSessionRequest) (_r *CreateQRMigrationSessionResponse, _err error)
	// Parameters:
	//  - Request
	SendEncryptedE2EEKey(ctx context.Context, request *SendEncryptedE2EEKeyRequest) (_err error)
}

type PrimaryQrCodeMigrationPreparationServiceClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewPrimaryQrCodeMigrationPreparationServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PrimaryQrCodeMigrationPreparationServiceClient {
	return &PrimaryQrCodeMigrationPreparationServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPrimaryQrCodeMigrationPreparationServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PrimaryQrCodeMigrationPreparationServiceClient {
	return &PrimaryQrCodeMigrationPreparationServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPrimaryQrCodeMigrationPreparationServiceClient(c thrift.TClient) *PrimaryQrCodeMigrationPreparationServiceClient {
	return &PrimaryQrCodeMigrationPreparationServiceClient{
		c: c,
	}
}

func (p *PrimaryQrCodeMigrationPreparationServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *PrimaryQrCodeMigrationPreparationServiceClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *PrimaryQrCodeMigrationPreparationServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//  - Request
func (p *PrimaryQrCodeMigrationPreparationServiceClient) CreateQRMigrationSession(ctx context.Context, request *CreateSessionRequest) (_r *CreateQRMigrationSessionResponse, _err error) {
	var _args0 PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs
	_args0.Request = request
	var _result2 PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult
	var _meta1 thrift.ResponseMeta
	_meta1, _err = p.Client_().Call(ctx, "createSession", &_args0, &_result2)
	p.SetLastResponseMeta_(_meta1)
	if _err != nil {
		return
	}
	switch {
	case _result2.Pqme != nil:
		return _r, _result2.Pqme
	case _result2.Tae != nil:
		return _r, _result2.Tae
	}

	if _ret3 := _result2.GetSuccess(); _ret3 != nil {
		return _ret3, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "createQRMigrationSession failed: unknown result")
}

// Parameters:
//  - Request
func (p *PrimaryQrCodeMigrationPreparationServiceClient) SendEncryptedE2EEKey(ctx context.Context, request *SendEncryptedE2EEKeyRequest) (_err error) {
	var _args4 PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs
	_args4.Request = request
	var _result6 PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult
	var _meta5 thrift.ResponseMeta
	_meta5, _err = p.Client_().Call(ctx, "sendEncryptedE2EEKey", &_args4, &_result6)
	p.SetLastResponseMeta_(_meta5)
	if _err != nil {
		return
	}
	switch {
	case _result6.Pqme != nil:
		return _result6.Pqme
	case _result6.Tae != nil:
		return _result6.Tae
	}

	return nil
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs struct {
	Request *CreateSessionRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs() *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs {
	return &PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs{}
}

var PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs_Request_DEFAULT *CreateSessionRequest

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs) GetRequest() *CreateSessionRequest {
	if !p.IsSetRequest() {
		return PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &CreateSessionRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "createQRMigrationSession_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - Pqme
//  - Tae
type PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult struct {
	Success *CreateQRMigrationSessionResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	Pqme    *PrimaryQrCodeMigrationException  `thrift:"pqme,1" db:"pqme" json:"pqme,omitempty"`
	Tae     *TokenAuthException               `thrift:"tae,2" db:"tae" json:"tae,omitempty"`
}

func NewPrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult() *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult {
	return &PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult{}
}

var PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult_Success_DEFAULT *CreateQRMigrationSessionResponse

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) GetSuccess() *CreateQRMigrationSessionResponse {
	if !p.IsSetSuccess() {
		return PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult_Pqme_DEFAULT *PrimaryQrCodeMigrationException

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) GetPqme() *PrimaryQrCodeMigrationException {
	if !p.IsSetPqme() {
		return PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult_Pqme_DEFAULT
	}
	return p.Pqme
}

var PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult_Tae_DEFAULT *TokenAuthException

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) GetTae() *TokenAuthException {
	if !p.IsSetTae() {
		return PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult_Tae_DEFAULT
	}
	return p.Tae
}
func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) IsSetPqme() bool {
	return p.Pqme != nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) IsSetTae() bool {
	return p.Tae != nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &CreateQRMigrationSessionResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Pqme = &PrimaryQrCodeMigrationException{}
	if err := p.Pqme.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Pqme), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Tae = &TokenAuthException{}
	if err := p.Tae.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tae), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "createQRMigrationSession_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetPqme() {
		if err := oprot.WriteFieldBegin(ctx, "pqme", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pqme: ", p), err)
		}
		if err := p.Pqme.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Pqme), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pqme: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetTae() {
		if err := oprot.WriteFieldBegin(ctx, "tae", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tae: ", p), err)
		}
		if err := p.Tae.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tae), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tae: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationPreparationServiceCreateQRMigrationSessionResult(%+v)", *p)
}

// Attributes:
//  - Request
type PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs struct {
	Request *SendEncryptedE2EEKeyRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs() *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs {
	return &PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs{}
}

var PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs_Request_DEFAULT *SendEncryptedE2EEKeyRequest

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs) GetRequest() *SendEncryptedE2EEKeyRequest {
	if !p.IsSetRequest() {
		return PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &SendEncryptedE2EEKeyRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "sendEncryptedE2EEKey_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyArgs(%+v)", *p)
}

// Attributes:
//  - Pqme
//  - Tae
type PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult struct {
	Pqme *PrimaryQrCodeMigrationException `thrift:"pqme,1" db:"pqme" json:"pqme,omitempty"`
	Tae  *TokenAuthException              `thrift:"tae,2" db:"tae" json:"tae,omitempty"`
}

func NewPrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult() *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult {
	return &PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult{}
}

var PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult_Pqme_DEFAULT *PrimaryQrCodeMigrationException

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) GetPqme() *PrimaryQrCodeMigrationException {
	if !p.IsSetPqme() {
		return PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult_Pqme_DEFAULT
	}
	return p.Pqme
}

var PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult_Tae_DEFAULT *TokenAuthException

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) GetTae() *TokenAuthException {
	if !p.IsSetTae() {
		return PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult_Tae_DEFAULT
	}
	return p.Tae
}
func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) IsSetPqme() bool {
	return p.Pqme != nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) IsSetTae() bool {
	return p.Tae != nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Pqme = &PrimaryQrCodeMigrationException{}
	if err := p.Pqme.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Pqme), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Tae = &TokenAuthException{}
	if err := p.Tae.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Tae), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "sendEncryptedE2EEKey_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetPqme() {
		if err := oprot.WriteFieldBegin(ctx, "pqme", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pqme: ", p), err)
		}
		if err := p.Pqme.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Pqme), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pqme: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetTae() {
		if err := oprot.WriteFieldBegin(ctx, "tae", thrift.STRUCT, 2); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tae: ", p), err)
		}
		if err := p.Tae.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Tae), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tae: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationPreparationServiceSendEncryptedE2EEKeyResult(%+v)", *p)
}

type PrimaryQrCodeMigrationLongPollingService interface {
	// Parameters:
	//  - Request
	CheckIfEncryptedE2EEKeyReceived(ctx context.Context, request *CheckIfEncryptedE2EEKeyReceivedRequest) (_r *CheckIfEncryptedE2EEKeyReceivedResponse, _err error)
}

type PrimaryQrCodeMigrationLongPollingServiceClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewPrimaryQrCodeMigrationLongPollingServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PrimaryQrCodeMigrationLongPollingServiceClient {
	return &PrimaryQrCodeMigrationLongPollingServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPrimaryQrCodeMigrationLongPollingServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PrimaryQrCodeMigrationLongPollingServiceClient {
	return &PrimaryQrCodeMigrationLongPollingServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPrimaryQrCodeMigrationLongPollingServiceClient(c thrift.TClient) *PrimaryQrCodeMigrationLongPollingServiceClient {
	return &PrimaryQrCodeMigrationLongPollingServiceClient{
		c: c,
	}
}

func (p *PrimaryQrCodeMigrationLongPollingServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *PrimaryQrCodeMigrationLongPollingServiceClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *PrimaryQrCodeMigrationLongPollingServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//  - Request
func (p *PrimaryQrCodeMigrationLongPollingServiceClient) CheckIfEncryptedE2EEKeyReceived(ctx context.Context, request *CheckIfEncryptedE2EEKeyReceivedRequest) (_r *CheckIfEncryptedE2EEKeyReceivedResponse, _err error) {
	var _args21 PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs
	_args21.Request = request
	var _result23 PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult
	var _meta22 thrift.ResponseMeta
	_meta22, _err = p.Client_().Call(ctx, "checkIfEncryptedE2EEKeyReceived", &_args21, &_result23)
	p.SetLastResponseMeta_(_meta22)
	if _err != nil {
		return
	}
	switch {
	case _result23.E != nil:
		return _r, _result23.E
	}

	if _ret24 := _result23.GetSuccess(); _ret24 != nil {
		return _ret24, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "checkIfEncryptedE2EEKeyReceived failed: unknown result")
}

// Attributes:
//  - SessionId
//  - Nonce
//  - NewDevice_
type MigratePrimaryUsingQrCodeRequest struct {
	SessionId  string   `thrift:"sessionId,1" db:"sessionId" json:"sessionId"`
	Nonce      string   `thrift:"nonce,2" db:"nonce" json:"nonce"`
	NewDevice_ *Device2 `thrift:"newDevice,3" db:"newDevice" json:"newDevice"`
}

func NewMigratePrimaryUsingQrCodeRequest() *MigratePrimaryUsingQrCodeRequest {
	return &MigratePrimaryUsingQrCodeRequest{}
}

func (p *MigratePrimaryUsingQrCodeRequest) GetSessionId() string {
	return p.SessionId
}

func (p *MigratePrimaryUsingQrCodeRequest) GetNonce() string {
	return p.Nonce
}

var MigratePrimaryUsingQrCodeRequest_NewDevice__DEFAULT *Device2

func (p *MigratePrimaryUsingQrCodeRequest) GetNewDevice_() *Device2 {
	if !p.IsSetNewDevice_() {
		return MigratePrimaryUsingQrCodeRequest_NewDevice__DEFAULT
	}
	return p.NewDevice_
}
func (p *MigratePrimaryUsingQrCodeRequest) IsSetNewDevice_() bool {
	return p.NewDevice_ != nil
}

func (p *MigratePrimaryUsingQrCodeRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.SessionId = v
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeRequest) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Nonce = v
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeRequest) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.NewDevice_ = &Device2{}
	if err := p.NewDevice_.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.NewDevice_), err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "MigratePrimaryUsingQrCodeRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "sessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:sessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.SessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.sessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:sessionId: ", p), err)
	}
	return err
}

func (p *MigratePrimaryUsingQrCodeRequest) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "nonce", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:nonce: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Nonce)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.nonce (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:nonce: ", p), err)
	}
	return err
}

func (p *MigratePrimaryUsingQrCodeRequest) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "newDevice", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:newDevice: ", p), err)
	}
	if err := p.NewDevice_.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.NewDevice_), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:newDevice: ", p), err)
	}
	return err
}

func (p *MigratePrimaryUsingQrCodeRequest) Equals(other *MigratePrimaryUsingQrCodeRequest) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.SessionId != other.SessionId {
		return false
	}
	if p.Nonce != other.Nonce {
		return false
	}
	return true
}

func (p *MigratePrimaryUsingQrCodeRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MigratePrimaryUsingQrCodeRequest(%+v)", *p)
}

// Attributes:
//  - AuthSessionId
type LookupAvailableEapRequest struct {
	AuthSessionId string `thrift:"authSessionId,1" db:"authSessionId" json:"authSessionId"`
}

func NewLookupAvailableEapRequest() *LookupAvailableEapRequest {
	return &LookupAvailableEapRequest{}
}

func (p *LookupAvailableEapRequest) GetAuthSessionId() string {
	return p.AuthSessionId
}
func (p *LookupAvailableEapRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LookupAvailableEapRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AuthSessionId = v
	}
	return nil
}

func (p *LookupAvailableEapRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "LookupAvailableEapRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LookupAvailableEapRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "authSessionId", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:authSessionId: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.AuthSessionId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.authSessionId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:authSessionId: ", p), err)
	}
	return err
}

func (p *LookupAvailableEapRequest) Equals(other *LookupAvailableEapRequest) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.AuthSessionId != other.AuthSessionId {
		return false
	}
	return true
}

func (p *LookupAvailableEapRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LookupAvailableEapRequest(%+v)", *p)
}

// Attributes:
//  - AvailableEaps
type LookupAvailableEapResponse struct {
	AvailableEaps []SocialLoginType `thrift:"availableEaps,1" db:"availableEaps" json:"availableEaps"`
}

func NewLookupAvailableEapResponse() *LookupAvailableEapResponse {
	return &LookupAvailableEapResponse{}
}

func (p *LookupAvailableEapResponse) GetAvailableEaps() []SocialLoginType {
	return p.AvailableEaps
}
func (p *LookupAvailableEapResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LookupAvailableEapResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]SocialLoginType, 0, size)
	p.AvailableEaps = tSlice
	for i := 0; i < size; i++ {
		var _elem9 SocialLoginType
		if v, err := iprot.ReadI32(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			temp := SocialLoginType(v)
			_elem9 = temp
		}
		p.AvailableEaps = append(p.AvailableEaps, _elem9)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *LookupAvailableEapResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "LookupAvailableEapResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LookupAvailableEapResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "availableEaps", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:availableEaps: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.I32, len(p.AvailableEaps)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.AvailableEaps {
		if err := oprot.WriteI32(ctx, int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:availableEaps: ", p), err)
	}
	return err
}

func (p *LookupAvailableEapResponse) Equals(other *LookupAvailableEapResponse) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if len(p.AvailableEaps) != len(other.AvailableEaps) {
		return false
	}
	for i, _tgt := range p.AvailableEaps {
		_src10 := other.AvailableEaps[i]
		if _tgt != _src10 {
			return false
		}
	}
	return true
}

func (p *LookupAvailableEapResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LookupAvailableEapResponse(%+v)", *p)
}

// Attributes:
//  - Mid
//  - TokenV3IssueResult_
//  - TokenV1IssueResult_
//  - AccountCountryCode
//  - FormattedPhoneNumbers
type MigratePrimaryUsingQrCodeResponse struct {
	Mid                   string                 `thrift:"mid,1" db:"mid" json:"mid"`
	TokenV3IssueResult_   *TokenV3IssueResult_   `thrift:"tokenV3IssueResult,2" db:"tokenV3IssueResult" json:"tokenV3IssueResult"`
	TokenV1IssueResult_   *TokenV1IssueResult_   `thrift:"tokenV1IssueResult,3" db:"tokenV1IssueResult" json:"tokenV1IssueResult"`
	AccountCountryCode    *CountryCode           `thrift:"accountCountryCode,4" db:"accountCountryCode" json:"accountCountryCode"`
	FormattedPhoneNumbers *FormattedPhoneNumbers `thrift:"formattedPhoneNumbers,5" db:"formattedPhoneNumbers" json:"formattedPhoneNumbers"`
}

func NewMigratePrimaryUsingQrCodeResponse() *MigratePrimaryUsingQrCodeResponse {
	return &MigratePrimaryUsingQrCodeResponse{}
}

func (p *MigratePrimaryUsingQrCodeResponse) GetMid() string {
	return p.Mid
}

var MigratePrimaryUsingQrCodeResponse_TokenV3IssueResult__DEFAULT *TokenV3IssueResult_

func (p *MigratePrimaryUsingQrCodeResponse) GetTokenV3IssueResult_() *TokenV3IssueResult_ {
	if !p.IsSetTokenV3IssueResult_() {
		return MigratePrimaryUsingQrCodeResponse_TokenV3IssueResult__DEFAULT
	}
	return p.TokenV3IssueResult_
}

var MigratePrimaryUsingQrCodeResponse_TokenV1IssueResult__DEFAULT *TokenV1IssueResult_

func (p *MigratePrimaryUsingQrCodeResponse) GetTokenV1IssueResult_() *TokenV1IssueResult_ {
	if !p.IsSetTokenV1IssueResult_() {
		return MigratePrimaryUsingQrCodeResponse_TokenV1IssueResult__DEFAULT
	}
	return p.TokenV1IssueResult_
}

var MigratePrimaryUsingQrCodeResponse_AccountCountryCode_DEFAULT *CountryCode

func (p *MigratePrimaryUsingQrCodeResponse) GetAccountCountryCode() *CountryCode {
	if !p.IsSetAccountCountryCode() {
		return MigratePrimaryUsingQrCodeResponse_AccountCountryCode_DEFAULT
	}
	return p.AccountCountryCode
}

var MigratePrimaryUsingQrCodeResponse_FormattedPhoneNumbers_DEFAULT *FormattedPhoneNumbers

func (p *MigratePrimaryUsingQrCodeResponse) GetFormattedPhoneNumbers() *FormattedPhoneNumbers {
	if !p.IsSetFormattedPhoneNumbers() {
		return MigratePrimaryUsingQrCodeResponse_FormattedPhoneNumbers_DEFAULT
	}
	return p.FormattedPhoneNumbers
}
func (p *MigratePrimaryUsingQrCodeResponse) IsSetTokenV3IssueResult_() bool {
	return p.TokenV3IssueResult_ != nil
}

func (p *MigratePrimaryUsingQrCodeResponse) IsSetTokenV1IssueResult_() bool {
	return p.TokenV1IssueResult_ != nil
}

func (p *MigratePrimaryUsingQrCodeResponse) IsSetAccountCountryCode() bool {
	return p.AccountCountryCode != nil
}

func (p *MigratePrimaryUsingQrCodeResponse) IsSetFormattedPhoneNumbers() bool {
	return p.FormattedPhoneNumbers != nil
}

func (p *MigratePrimaryUsingQrCodeResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 4:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField4(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 5:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField5(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Mid = v
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.TokenV3IssueResult_ = &TokenV3IssueResult_{}
	if err := p.TokenV3IssueResult_.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TokenV3IssueResult_), err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeResponse) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	p.TokenV1IssueResult_ = &TokenV1IssueResult_{}
	if err := p.TokenV1IssueResult_.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.TokenV1IssueResult_), err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeResponse) ReadField4(ctx context.Context, iprot thrift.TProtocol) error {
	p.AccountCountryCode = &CountryCode{}
	if err := p.AccountCountryCode.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.AccountCountryCode), err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeResponse) ReadField5(ctx context.Context, iprot thrift.TProtocol) error {
	p.FormattedPhoneNumbers = &FormattedPhoneNumbers{}
	if err := p.FormattedPhoneNumbers.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.FormattedPhoneNumbers), err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "MigratePrimaryUsingQrCodeResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField4(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField5(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *MigratePrimaryUsingQrCodeResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "mid", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:mid: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Mid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.mid (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:mid: ", p), err)
	}
	return err
}

func (p *MigratePrimaryUsingQrCodeResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "tokenV3IssueResult", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:tokenV3IssueResult: ", p), err)
	}
	if err := p.TokenV3IssueResult_.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TokenV3IssueResult_), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:tokenV3IssueResult: ", p), err)
	}
	return err
}

func (p *MigratePrimaryUsingQrCodeResponse) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "tokenV1IssueResult", thrift.STRUCT, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:tokenV1IssueResult: ", p), err)
	}
	if err := p.TokenV1IssueResult_.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.TokenV1IssueResult_), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:tokenV1IssueResult: ", p), err)
	}
	return err
}

func (p *MigratePrimaryUsingQrCodeResponse) writeField4(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "accountCountryCode", thrift.STRUCT, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:accountCountryCode: ", p), err)
	}
	if err := p.AccountCountryCode.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.AccountCountryCode), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:accountCountryCode: ", p), err)
	}
	return err
}

func (p *MigratePrimaryUsingQrCodeResponse) writeField5(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "formattedPhoneNumbers", thrift.STRUCT, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:formattedPhoneNumbers: ", p), err)
	}
	if err := p.FormattedPhoneNumbers.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.FormattedPhoneNumbers), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:formattedPhoneNumbers: ", p), err)
	}
	return err
}

func (p *MigratePrimaryUsingQrCodeResponse) Equals(other *MigratePrimaryUsingQrCodeResponse) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Mid != other.Mid {
		return false
	}
	if !p.TokenV1IssueResult_.Equals(other.TokenV1IssueResult_) {
		return false
	}
	if !p.AccountCountryCode.Equals(other.AccountCountryCode) {
		return false
	}
	if !p.FormattedPhoneNumbers.Equals(other.FormattedPhoneNumbers) {
		return false
	}
	return true
}

func (p *MigratePrimaryUsingQrCodeResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MigratePrimaryUsingQrCodeResponse(%+v)", *p)
}

// Attributes:
//  - TokenSecret
type TokenV1IssueResult_ struct {
	TokenSecret string `thrift:"tokenSecret,1" db:"tokenSecret" json:"tokenSecret"`
}

func NewTokenV1IssueResult_() *TokenV1IssueResult_ {
	return &TokenV1IssueResult_{}
}

func (p *TokenV1IssueResult_) GetTokenSecret() string {
	return p.TokenSecret
}
func (p *TokenV1IssueResult_) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *TokenV1IssueResult_) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.TokenSecret = v
	}
	return nil
}

func (p *TokenV1IssueResult_) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "TokenV1IssueResult"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *TokenV1IssueResult_) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "tokenSecret", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:tokenSecret: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.TokenSecret)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.tokenSecret (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:tokenSecret: ", p), err)
	}
	return err
}

func (p *TokenV1IssueResult_) Equals(other *TokenV1IssueResult_) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.TokenSecret != other.TokenSecret {
		return false
	}
	return true
}

func (p *TokenV1IssueResult_) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("TokenV1IssueResult_(%+v)", *p)
}

// Attributes:
//  - Code
type CountryCode struct {
	Code string `thrift:"code,1" db:"code" json:"code"`
}

func NewCountryCode() *CountryCode {
	return &CountryCode{}
}

func (p *CountryCode) GetCode() string {
	return p.Code
}
func (p *CountryCode) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *CountryCode) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Code = v
	}
	return nil
}

func (p *CountryCode) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "CountryCode"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *CountryCode) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "code", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:code: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Code)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.code (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:code: ", p), err)
	}
	return err
}

func (p *CountryCode) Equals(other *CountryCode) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Code != other.Code {
		return false
	}
	return true
}

func (p *CountryCode) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CountryCode(%+v)", *p)
}

// Attributes:
//  - LocalFormatPhoneNumber
//  - PrettifiedFormatPhoneNumber
type FormattedPhoneNumbers struct {
	LocalFormatPhoneNumber      string `thrift:"localFormatPhoneNumber,1" db:"localFormatPhoneNumber" json:"localFormatPhoneNumber"`
	PrettifiedFormatPhoneNumber string `thrift:"prettifiedFormatPhoneNumber,2" db:"prettifiedFormatPhoneNumber" json:"prettifiedFormatPhoneNumber"`
}

func NewFormattedPhoneNumbers() *FormattedPhoneNumbers {
	return &FormattedPhoneNumbers{}
}

func (p *FormattedPhoneNumbers) GetLocalFormatPhoneNumber() string {
	return p.LocalFormatPhoneNumber
}

func (p *FormattedPhoneNumbers) GetPrettifiedFormatPhoneNumber() string {
	return p.PrettifiedFormatPhoneNumber
}
func (p *FormattedPhoneNumbers) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *FormattedPhoneNumbers) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.LocalFormatPhoneNumber = v
	}
	return nil
}

func (p *FormattedPhoneNumbers) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.PrettifiedFormatPhoneNumber = v
	}
	return nil
}

func (p *FormattedPhoneNumbers) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "FormattedPhoneNumbers"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *FormattedPhoneNumbers) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "localFormatPhoneNumber", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:localFormatPhoneNumber: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.LocalFormatPhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.localFormatPhoneNumber (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:localFormatPhoneNumber: ", p), err)
	}
	return err
}

func (p *FormattedPhoneNumbers) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "prettifiedFormatPhoneNumber", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:prettifiedFormatPhoneNumber: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.PrettifiedFormatPhoneNumber)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.prettifiedFormatPhoneNumber (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:prettifiedFormatPhoneNumber: ", p), err)
	}
	return err
}

func (p *FormattedPhoneNumbers) Equals(other *FormattedPhoneNumbers) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.LocalFormatPhoneNumber != other.LocalFormatPhoneNumber {
		return false
	}
	if p.PrettifiedFormatPhoneNumber != other.PrettifiedFormatPhoneNumber {
		return false
	}
	return true
}

func (p *FormattedPhoneNumbers) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FormattedPhoneNumbers(%+v)", *p)
}

type PrimaryQrCodeMigrationService interface {
	// Parameters:
	//  - Request
	MigratePrimaryUsingQrCode(ctx context.Context, request *MigratePrimaryUsingQrCodeRequest) (_r *MigratePrimaryUsingQrCodeResponse, _err error)
}

type PrimaryQrCodeMigrationServiceClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewPrimaryQrCodeMigrationServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *PrimaryQrCodeMigrationServiceClient {
	return &PrimaryQrCodeMigrationServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewPrimaryQrCodeMigrationServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *PrimaryQrCodeMigrationServiceClient {
	return &PrimaryQrCodeMigrationServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewPrimaryQrCodeMigrationServiceClient(c thrift.TClient) *PrimaryQrCodeMigrationServiceClient {
	return &PrimaryQrCodeMigrationServiceClient{
		c: c,
	}
}

func (p *PrimaryQrCodeMigrationServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *PrimaryQrCodeMigrationServiceClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *PrimaryQrCodeMigrationServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//  - Request
func (p *PrimaryQrCodeMigrationServiceClient) MigratePrimaryUsingQrCode(ctx context.Context, request *MigratePrimaryUsingQrCodeRequest) (_r *MigratePrimaryUsingQrCodeResponse, _err error) {
	var _args33 PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs
	_args33.Request = request
	var _result35 PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult
	var _meta34 thrift.ResponseMeta
	_meta34, _err = p.Client_().Call(ctx, "migratePrimaryUsingQrCode", &_args33, &_result35)
	p.SetLastResponseMeta_(_meta34)
	if _err != nil {
		return
	}
	switch {
	case _result35.E != nil:
		return _r, _result35.E
	}

	if _ret36 := _result35.GetSuccess(); _ret36 != nil {
		return _ret36, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "migratePrimaryUsingQrCode failed: unknown result")
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Request
type PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs struct {
	Request *MigratePrimaryUsingQrCodeRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs() *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs {
	return &PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs{}
}

var PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs_Request_DEFAULT *MigratePrimaryUsingQrCodeRequest

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs) GetRequest() *MigratePrimaryUsingQrCodeRequest {
	if !p.IsSetRequest() {
		return PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &MigratePrimaryUsingQrCodeRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingQrCode_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult struct {
	Success *MigratePrimaryUsingQrCodeResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *PrimaryQrCodeMigrationException   `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult() *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult {
	return &PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult{}
}

var PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult_Success_DEFAULT *MigratePrimaryUsingQrCodeResponse

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) GetSuccess() *MigratePrimaryUsingQrCodeResponse {
	if !p.IsSetSuccess() {
		return PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult_E_DEFAULT *PrimaryQrCodeMigrationException

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) GetE() *PrimaryQrCodeMigrationException {
	if !p.IsSetE() {
		return PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &MigratePrimaryUsingQrCodeResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &PrimaryQrCodeMigrationException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "migratePrimaryUsingQrCode_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationServiceMigratePrimaryUsingQrCodeResult(%+v)", *p)
}

// Attributes:
//  - Request
type PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs struct {
	Request *CheckIfEncryptedE2EEKeyReceivedRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs() *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs {
	return &PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs{}
}

var PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs_Request_DEFAULT *CheckIfEncryptedE2EEKeyReceivedRequest

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs) GetRequest() *CheckIfEncryptedE2EEKeyReceivedRequest {
	if !p.IsSetRequest() {
		return PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &CheckIfEncryptedE2EEKeyReceivedRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "checkIfEncryptedE2EEKeyReceived_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult struct {
	Success *CheckIfEncryptedE2EEKeyReceivedResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *PrimaryQrCodeMigrationException         `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult() *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult {
	return &PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult{}
}

var PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult_Success_DEFAULT *CheckIfEncryptedE2EEKeyReceivedResponse

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) GetSuccess() *CheckIfEncryptedE2EEKeyReceivedResponse {
	if !p.IsSetSuccess() {
		return PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult_E_DEFAULT *PrimaryQrCodeMigrationException

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) GetE() *PrimaryQrCodeMigrationException {
	if !p.IsSetE() {
		return PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &CheckIfEncryptedE2EEKeyReceivedResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &PrimaryQrCodeMigrationException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "checkIfEncryptedE2EEKeyReceived_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryQrCodeMigrationLongPollingServiceCheckIfEncryptedE2EEKeyReceivedResult(%+v)", *p)
}

// Attributes:
//  - Request
type PrimaryAccountInitServiceLookupAvailableEapArgs struct {
	Request *LookupAvailableEapRequest `thrift:"request,1" db:"request" json:"request"`
}

func NewPrimaryAccountInitServiceLookupAvailableEapArgs() *PrimaryAccountInitServiceLookupAvailableEapArgs {
	return &PrimaryAccountInitServiceLookupAvailableEapArgs{}
}

var PrimaryAccountInitServiceLookupAvailableEapArgs_Request_DEFAULT *LookupAvailableEapRequest

func (p *PrimaryAccountInitServiceLookupAvailableEapArgs) GetRequest() *LookupAvailableEapRequest {
	if !p.IsSetRequest() {
		return PrimaryAccountInitServiceLookupAvailableEapArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *PrimaryAccountInitServiceLookupAvailableEapArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Request = &LookupAvailableEapRequest{}
	if err := p.Request.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "lookupAvailableEap_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *PrimaryAccountInitServiceLookupAvailableEapArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceLookupAvailableEapArgs(%+v)", *p)
}

// Attributes:
//  - Success
//  - E
type PrimaryAccountInitServiceLookupAvailableEapResult struct {
	Success *LookupAvailableEapResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
	E       *AuthException              `thrift:"e,1" db:"e" json:"e,omitempty"`
}

func NewPrimaryAccountInitServiceLookupAvailableEapResult() *PrimaryAccountInitServiceLookupAvailableEapResult {
	return &PrimaryAccountInitServiceLookupAvailableEapResult{}
}

var PrimaryAccountInitServiceLookupAvailableEapResult_Success_DEFAULT *LookupAvailableEapResponse

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) GetSuccess() *LookupAvailableEapResponse {
	if !p.IsSetSuccess() {
		return PrimaryAccountInitServiceLookupAvailableEapResult_Success_DEFAULT
	}
	return p.Success
}

var PrimaryAccountInitServiceLookupAvailableEapResult_E_DEFAULT *AuthException

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) GetE() *AuthException {
	if !p.IsSetE() {
		return PrimaryAccountInitServiceLookupAvailableEapResult_E_DEFAULT
	}
	return p.E
}
func (p *PrimaryAccountInitServiceLookupAvailableEapResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) IsSetE() bool {
	return p.E != nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &LookupAvailableEapResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.E = &AuthException{}
	if err := p.E.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.E), err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "lookupAvailableEap_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin(ctx, "e", thrift.STRUCT, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:e: ", p), err)
		}
		if err := p.E.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.E), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:e: ", p), err)
		}
	}
	return err
}

func (p *PrimaryAccountInitServiceLookupAvailableEapResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PrimaryAccountInitServiceLookupAvailableEapResult(%+v)", *p)
}
