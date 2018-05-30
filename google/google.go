package google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/thommil/animals-go-common/model"
)

// Configuration definition for facebook providers
type Configuration struct {
	// URL used to check token (:idToken replaced)
	URL string
	// Issuer to check
	ISS string
	// Audience to check
	AUD string
}

type tokenInfo struct {
	AZP        string `json:"azp"`
	AUD        string `json:"aud"`
	SUB        string `json:"sub"`
	EXP        string `json:"exp"`
	ISS        string `json:"iss"`
	IAT        string `json:"iat"`
	Name       string `json:"name"`
	Picture    string `json:"picture"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Locale     string `json:"locale"`
	ALG        string `json:"alg"`
	KID        string `json:"kid"`
}

// Provider allows to check user entry against Google JWT token
type Provider struct {
	Database      *mgo.Database
	Configuration *Configuration
}

// Authenticate implementation of Provider API
func (provider Provider) Authenticate(credentials interface{}) (*model.User, error) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}

	response, err := httpClient.Get(strings.Replace(provider.Configuration.URL, ":idToken", credentials.(string), 1))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	token := &tokenInfo{}
	json.NewDecoder(response.Body).Decode(token)
	if token.ISS == provider.Configuration.ISS && token.AUD == provider.Configuration.AUD {
		query := model.FindAccount(provider.Database, &model.Account{ExternalID: token.SUB})
		count, err := query.Count()
		if err != nil {
			return nil, err
		}
		if count > 0 {
			//Found return user
			account := &model.Account{}
			if query.One(account) != nil {
				return nil, err
			}
			return model.FindUserByID(provider.Database, account.UserID)
		}

		//Not found, create account & user
		user, err := model.CreateOrUpdateUser(provider.Database, &model.User{Username: token.Name})
		if err != nil {
			return nil, err
		}
		model.CreateOrUpdateAccount(provider.Database, &model.Account{Provider: "google", ExternalID: token.SUB, UserID: user.ID.Hex()})
		return user, nil
	}
	return nil, fmt.Errorf("Bad 'iss' or 'aud' claim in Google token")
}
