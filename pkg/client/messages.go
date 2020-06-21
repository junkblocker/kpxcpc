package client

const ChangePublicKeysAction = "change-public-keys"
const AssociateAction = "associate"
const GetLoginsAction = "get-logins"
const TestAssociateAction = "test-associate"

type Request struct {
	Action   string      `json:"action"`
	ClientID Base64Bytes `json:"clientID"`
	Nonce    Base64Bytes `json:"nonce"`
	Message  Base64Bytes `json:"message,omitempty"`
}

type Response struct {
	Message Base64Bytes `json:"message"`
	Nonce   Base64Bytes `json:"nonce"`
	Error   *string     `json:"error"`
	Code    *int        `json:"errorCode,string"`
	Success *bool       `json:"success,string"`
	Version *string     `json:"version"`
	Hash    *string     `json:"hash"`
	ID      string      `json:"id"`
}

type AssociateResponseMessage struct {
	Response
}

type TestAssociateResponseMessage struct {
	Response
}

type ChangePublicKeysRequest struct {
	Request
	PulicKey Base64Bytes `json:"publicKey"`
}

type ChangePublicKeysResponse struct {
	Response
	PulicKey Base64Bytes `json:"publicKey"`
}

type AssociateMessage struct {
	Action string      `json:"action"`
	Key    Base64Bytes `json:"key"`   // client pubkey
	IDKey  Base64Bytes `json:"idKey"` // new id pubkey
}

type TestAssociateMessage struct {
	Action string `json:"action"`
	DBKey
}

type GetLoginsMessage struct {
	Action    string  `json:"action"`
	URL       string  `json:"url"`
	SubmitURL string  `json:"submitUrl,omitempty"`
	HTTPAuth  string  `json:"httpAuth,omitempty"`
	Keys      []DBKey `json:"keys,omitempty"`
}

type DBKey struct {
	ID  string      `json:"id"`  // saved DB identifier
	Key Base64Bytes `json:"key"` // saved pubkey
}

type GetLoginsResponseMessage struct {
	Response
	Count   int          `json:"count"`
	Entries []LoginEntry `json:"entries"`
}

type LoginEntry struct {
	Login        string              `json:"login"`
	Name         string              `json:"name"`
	Password     string              `json:"password"`
	UUID         string              `json:"uuid"`
	StringFields []map[string]string `json:"stringFields"`
}