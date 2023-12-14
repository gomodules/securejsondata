package util

import (
	"fmt"
	"strings"
)

const prodDomain = "appscode.com"

func GetBaseDomain(host string) string {
	addr, err := SplitHostPortDefault(host, "", "")
	if err != nil {
		return prodDomain
	}
	parts := strings.Split(addr.Host, ".")
	base := strings.Join(parts[1:], ".")
	return base
}

func platformDomain(domain string) string {
	if domain == prodDomain {
		return fmt.Sprintf("home.%s", domain)
	}
	return domain
}

func GetPlatformUrl(host string, production bool) string {
	var baseUrl string
	baseDomain := GetBaseDomain(host)
	if production {
		baseUrl = fmt.Sprintf("https://%v", platformDomain(baseDomain))
	} else {
		baseUrl = fmt.Sprintf("http://%v:8080", baseDomain)
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
