package models

import (
	utils "staff-service/Utils"
)

type BaseResponse struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type JWT struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (token *JWT) Create(username string) (err error) {
	token.Token, err = utils.GenerateJWT(username, utils.TOKEN)

	if err != nil {
		return err
	}

	token.RefreshToken, err = utils.GenerateJWT(username, utils.REFRESH_TOKEN)

	if err != nil {
		return err
	}

	return nil
}
