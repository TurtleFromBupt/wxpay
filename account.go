package wxpay

import (
	"io/ioutil"
)

type Account struct {
	appID     string
	mchID     string
	apiKey    string
	certData  []byte
	isSandbox bool
}

// 创建微信支付账号
func NewAccount(appID string, mchID string, apiKey string, isSanbox bool) *Account {
	return &Account{
		appID:     appID,
		mchID:     mchID,
		apiKey:    apiKey,
		isSandbox: isSanbox,
	}
}

// set cert file
func (a *Account) SetCertFile(certPath string) error {
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		return err
	}
	a.certData = certData
	return nil
}

// set cert data
func (a *Account) SetCertData(certData []byte) {
	a.certData = certData
}
