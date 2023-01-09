package util

import (
	"fmt"
	"strings"
)

func GetBaseDomain(host string) string {
	addr, err := SplitHostPortDefault(host, "", "")
	if err != nil {
		return "byte.builders"
	}
	parts := strings.Split(addr.Host, ".")
	base := strings.Join(parts[1:], ".")
	return base
}

func GetBaseUrl(host string, production bool) string {
	var baseUrl string
	baseDomain := GetBaseDomain(host)
	if production {
		baseUrl = fmt.Sprintf("https://%v", baseDomain)
	} else {
		baseUrl = fmt.Sprintf("http://%v:5998", baseDomain)
	}
	return baseUrl
}

func GetApiUrl(host string, production bool) string {
	var apiUrl string
	baseDomain := GetBaseDomain(host)
	if production {
		apiUrl = fmt.Sprintf("https://api.%v", baseDomain)
	} else {
		apiUrl = fmt.Sprintf("http://api.%v:3003", baseDomain)
	}
	return apiUrl
}

func GetAccountsUrl(host string, production bool) string {
	var accountsUrl string
	baseDomain := GetBaseDomain(host)
	if production {
		accountsUrl = fmt.Sprintf("https://accounts.%v", baseDomain)
	} else {
		accountsUrl = fmt.Sprintf("http://accounts.%v:3000", baseDomain)
	}
	return accountsUrl
}
