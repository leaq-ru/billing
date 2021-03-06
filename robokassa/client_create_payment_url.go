package robokassa

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

var ErrOnlyIntegerRub = errors.New("only integer rub allowed")

func (c Client) CreatePaymentURL(rkInvoiceID uint64, amount uint32) (paymentURL string, err error) {
	if !strings.HasSuffix(strconv.Itoa(int(amount)), "00") {
		err = ErrOnlyIntegerRub
		return
	}

	outSum := strconv.Itoa(int(amount / 100))

	strRKInvID := strconv.Itoa(int(rkInvoiceID))

	sha := sha512.New()
	_, err = sha.Write([]byte(strings.Join([]string{
		c.merchantLogin,
		outSum,
		strRKInvID,
		c.passwordOne,
	}, ":")))
	if err != nil {
		return
	}

	hash := strings.ToUpper(hex.EncodeToString(sha.Sum(nil)))

	u, err := url.Parse("https://auth.robokassa.ru/Merchant/Index.aspx")
	if err != nil {
		return
	}

	q := u.Query()
	q.Set("MerchantLogin", c.merchantLogin)
	q.Set("OutSum", outSum)
	q.Set("InvId", strRKInvID)
	q.Set("Description", "Пополнение баланса LEAQ")
	q.Set("SignatureValue", hash)
	q.Set("Culture", "ru")
	if isTest(c.isTest) {
		q.Set("IsTest", c.isTest)
	}
	u.RawQuery = q.Encode()

	paymentURL = u.String()
	return
}
