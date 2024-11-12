package api

import (
	"github.com/elifgider/iyzipay-go/model"
	"github.com/elifgider/iyzipay-go/utils"
	"github.com/sirupsen/logrus"
)

func RetrievePayment(req utils.RequestConvertible, opt *model.Options) (*retrievePayment, error) {
	logrus.Error("-*-RetrievePayment")
	resp := &retrievePayment{}
	if err := request("POST", "/payment/detail", req, opt, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func CreatePayment(req utils.RequestConvertible, opt *model.Options) (*createPayment, error) {
	logrus.Error("-*-CreatePayment")
	resp := &createPayment{}
	if err := request("POST", "/payment/auth", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func InitializeThreedsPayment(req utils.RequestConvertible, opt *model.Options) (*initializeThreedsPayment, error) {
	logrus.Error("-*-InitializeThreedsPayment")
	resp := &initializeThreedsPayment{}
	if err := request("POST", "/payment/3dsecure/initialize", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateThreedsPayment(req utils.RequestConvertible, opt *model.Options) (*createPayment, error) {
	logrus.Error("-*-CreateThreedsPayment")
	resp := &createPayment{}
	if err := request("POST", "/payment/3dsecure/auth", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func CancelPayment(req utils.RequestConvertible, opt *model.Options) (*cancelPayment, error) {
	logrus.Error("-*-CancelPayment")
	resp := &cancelPayment{}
	if err := request("POST", "/payment/cancel", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func RefundPayment(req utils.RequestConvertible, opt *model.Options) (*refundPayment, error) {
	logrus.Error("-*-RefundPayment")
	resp := &refundPayment{}
	if err := request("POST", "/payment/refund", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
