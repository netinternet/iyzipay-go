package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"

	model "github.com/netinternet/iyzipay-go/model"
	utils "github.com/netinternet/iyzipay-go/utils"
)

func request(method, url string, request utils.RequestConvertible, option *model.Options, respStruct interface{}) error {
	requestBody, err := json.Marshal(request.GetJsonObject())
	if err != nil {
		return err
	}

	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, option.GetBaseUrl()+url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	headers := getHttpHeaders(request, option, url, string(requestBody))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &respStruct); err != nil {
		return err
	}

	return nil
}

func getHttpHeaders(_ utils.RequestConvertible, option *model.Options, uriPath, requestBody string) map[string]string {
	header := make(map[string]string)
	header["Accept"] = "application/json"
	header["Content-type"] = "application/json"

	rnd := utils.RandString(8)
	encryptedData := generateEncryptedData(rnd, uriPath, requestBody, option.GetApiSecret())
	base64EncodedAuthorization := prepareAuthStringV2(option, rnd, encryptedData)

	header["Authorization"] = "IYZWSv2" + " " + base64EncodedAuthorization
	header["x-iyzi-rnd"] = rnd
	header["x-iyzi-client-version"] = "iyzipay-php-2.0.51"
	return header
}

func generateEncryptedData(randomKey, uriPath, requestBody, secretKey string) string {
	// HMACSHA256 algoritmasını kullanarak şifreleme
	data := randomKey + uriPath + requestBody
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
func prepareAuthStringV2(options *model.Options, rnd, encryptedData string) string {
	authString := "apiKey:" + options.GetApiKey() + "&randomKey:" + rnd + "&signature:" + encryptedData
	return base64.URLEncoding.EncodeToString([]byte(authString))
}
