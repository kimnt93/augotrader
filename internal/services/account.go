package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"encoding/json"
	"fmt"
)

type LoginInfo struct {
	AccountId      string `json:"account_id"`
	ConsumerId     string `json:"consumer_id"`
	ConsumerSecret string `json:"consumer_secret"`
	PrivateKey     string `json:"private_key"`
	AuthToken      string `json:"auth_token"`
	IsPaperTrading bool   `json:"is_paper_trading"`
	IsDisabled     bool   `json:"is_disabled"`
}

func GetAllAccountIds() ([]string, error) {
	allAccountIds := []string{}

	loginInfo, err := GetAllLoginInfo()
	if err != nil {
		return allAccountIds, err
	}

	for _, info := range loginInfo {
		// Append account id
		allAccountIds = append(allAccountIds, info.AccountId)
	}
	return allAccountIds, nil
}

func SetLoginInfoByAccountId(accountId string, consumerId string, consumerSecret string, privateKey string, authToken string, isPaperTrading bool, is_disabled bool) (bool, error) {
	loginInfo := LoginInfo{
		AccountId:      accountId,
		ConsumerId:     consumerId,
		ConsumerSecret: consumerSecret,
		PrivateKey:     privateKey,
		AuthToken:      authToken,
		IsPaperTrading: isPaperTrading,
		IsDisabled:     is_disabled,
	}

	key := fmt.Sprintf("%s.%s", static.CFG_ACCOUNT_LOGIN_INFO, accountId)
	value, err := json.Marshal(loginInfo)
	if err != nil {
		return false, err
	}

	_, err = cache.SetKeyStr(key, value)
	if err != nil {
		return false, err
	}

	return true, nil
}

func UpdateLoginInfoByAccountId(accountId string, consumerId string, consumerSecret string, privateKey string, authToken string, isPaperTrading bool, is_disabled bool) (bool, error) {
	return SetLoginInfoByAccountId(accountId, consumerId, consumerSecret, privateKey, authToken, isPaperTrading, is_disabled)
}

func DeleteLoginInfoByAccountId(accountId string) (LoginInfo, error) {
	loginInfo, err := GetLoginInfoByAccountId(accountId)
	if err != nil {
		return loginInfo, err
	}

	key := fmt.Sprintf("%s.%s", static.CFG_ACCOUNT_LOGIN_INFO, accountId)
	_, err = cache.DeleteKey(key)

	if err != nil {
		return loginInfo, err
	}
	return loginInfo, nil
}

func GetAllLoginInfo() ([]LoginInfo, error) {
	var loginInfos []LoginInfo

	keys, err := cache.RedisClient.Keys(cache.Ctx, fmt.Sprintf("%s.*", static.CFG_ACCOUNT_LOGIN_INFO)).Result()
	if err != nil {
		return loginInfos, err
	}

	for _, key := range keys {
		var loginInfo LoginInfo
		jsonStr, err := cache.GetKeyStr(key)
		if err != nil {
			return loginInfos, err
		}

		err = json.Unmarshal([]byte(jsonStr), &loginInfo)
		if err != nil {
			return loginInfos, err
		}
		loginInfos = append(loginInfos, loginInfo)
	}

	return loginInfos, nil
}

func GetLoginInfoByAccountId(accountId string) (LoginInfo, error) {
	var loginInfo LoginInfo

	key := fmt.Sprintf("%s.%s", static.CFG_ACCOUNT_LOGIN_INFO, accountId)
	jsonStr, err := cache.GetKeyStr(key)
	if err != nil {
		return loginInfo, err
	}

	err = json.Unmarshal([]byte(jsonStr), &loginInfo)
	if err != nil {
		return loginInfo, err
	}

	return loginInfo, nil
}
