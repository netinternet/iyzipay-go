package api

import (
	"github.com/elifgider/iyzipay-go/model"
	"github.com/elifgider/iyzipay-go/utils"
	"github.com/sirupsen/logrus"
)

func InstallmentInfo(req utils.RequestConvertible, opt *model.Options) (*installmentInfo, error) {
	logrus.Error("-*InstallmentInfo")
	resp := &installmentInfo{}
	if err := request("POST", "/payment/iyzipos/installment", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func BinNumber(req utils.RequestConvertible, opt *model.Options) (*binNumber, error) {
	logrus.Error("-*BinNumber")
	resp := &binNumber{}
	if err := request("POST", "/payment/bin/check", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
