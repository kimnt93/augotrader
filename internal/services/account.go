package services

import (
	"augotrader/internal/cache"
	"augotrader/internal/static"
	"augotrader/internal/types"
	"encoding/json"
	"fmt"
)

func GetAllAccounts() ([]types.AccountLoginSummary, error) {
	allAccounts := []types.AccountLoginSummary{}

	loginInfo, err := GetAllLoginInfo()
	if err != nil {
		return allAccounts, err
	}

	for _, info := range loginInfo {
		accountSummary := types.AccountLoginSummary{
			AccountId:  info.AccountId,
			IsDisabled: info.IsDisabled,
		}

		// Append account id
		allAccounts = append(allAccounts, accountSummary)
	}
	return allAccounts, nil
}

func SetLoginInfoByAccountId(accountId string, consumerId string, consumerSecret string, privateKey string, authToken string, isPaperTrading bool, is_disabled bool) (types.LoginInfo, error) {
	loginInfo := types.LoginInfo{
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
		return types.LoginInfo{}, err
	}

	_, err = cache.SetKeyStr(key, value)
	if err != nil {
		return types.LoginInfo{}, err
	}

	return loginInfo, nil
}

func UpdateLoginInfoByAccountId(accountId string, consumerId string, consumerSecret string, privateKey string, authToken string, isPaperTrading bool, is_disabled bool) (types.LoginInfo, error) {
	return SetLoginInfoByAccountId(accountId, consumerId, consumerSecret, privateKey, authToken, isPaperTrading, is_disabled)
}

func DeleteLoginInfoByAccountId(accountId string) (types.LoginInfo, error) {
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

func GetAllLoginInfo() ([]types.LoginInfo, error) {
	var loginInfos []types.LoginInfo

	keys, err := cache.RedisClient.Keys(cache.Ctx, fmt.Sprintf("%s.*", static.CFG_ACCOUNT_LOGIN_INFO)).Result()
	if err != nil {
		return loginInfos, err
	}

	for _, key := range keys {
		var loginInfo types.LoginInfo
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

func GetLoginInfoByAccountId(accountId string) (types.LoginInfo, error) {
	var loginInfo types.LoginInfo

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
