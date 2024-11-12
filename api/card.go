package api

import (
	"github.com/elifgider/iyzipay-go/model"
	"github.com/elifgider/iyzipay-go/utils"
	"github.com/sirupsen/logrus"
)

func RetriveCards(req utils.RequestConvertible, opt *model.Options) (*retriveCards, error) {
	logrus.Error("-*-RetriveCards")
	resp := &retriveCards{}
	if err := request("POST", "/cardstorage/cards", req, opt, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func CreateCard(req utils.RequestConvertible, opt *model.Options) (*createCard, error) {
	logrus.Error("-*-CreateCard")
	resp := &createCard{}
	if err := request("POST", "/cardstorage/card", req, opt, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func DeleteCard(req utils.RequestConvertible, opt *model.Options) (*baseResponse, error) {
	logrus.Error("-*-DeleteCard")
	resp := &baseResponse{}
	if err := request("DELETE", "/cardstorage/card", req, opt, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
