package entity

import (
	"net/http"
)

//Base HTTP errors
var (
	ErrBadRequest   = ErrorEntity{http.StatusBadRequest, 10001, "Bad Request", "Received unprocessable request or entity."}
	ErrInternal     = ErrorEntity{http.StatusInternalServerError, 10002, "Internal server error", "Something goes wrong. Try again later."}
	ErrNotFound     = ErrorEntity{http.StatusNotFound, 10003, "Not found", "Required entity or route was not found."}
	ErrUnauthorized = ErrorEntity{http.StatusUnauthorized, 10004, "User unauthorized", "Current user is unauthorized."}

	//Special debug errors
	ErrNotImplemented = ErrorEntity{http.StatusBadRequest, 20002, "Not implemented yet", "Sorry, requested functionality is not implemented yet"}

	//Auth errors
	AuthErrWrongType    = ErrorEntity{http.StatusBadRequest, 30001, "Unknown auth method called", "Requested auth method does not supported"}
	AuthErrNoCodeParam  = ErrorEntity{http.StatusBadRequest, 30002, "Missing \"code\" parameter", "Missing \"code\" parameter in request body"}
	AuthErrVkUserInfo   = ErrorEntity{http.StatusInternalServerError, 30003, "Can not get VK user info", "An error occurred while trying to get VK user info"}
	AuthErrSetToken     = ErrorEntity{http.StatusInternalServerError, 30004, "Error while setting the token", "An error occurred while trying to get generate auth token"}
	AuthErrNewUser      = ErrorEntity{http.StatusInternalServerError, 30005, "Error while creating new user", "An error occurred while trying to add new user to database"}
	AuthErrNewGamer     = ErrorEntity{http.StatusInternalServerError, 30005, "Error while creating new gamer", "An error occurred while trying to add new gamer to database"}
	AuthErrUserNotFound = ErrorEntity{http.StatusForbidden, 30006, "User not found", "User / password combination not found"}

	//Data errors
	DataErrBadUserInfo    = ErrorEntityEx{http.StatusBadRequest, 40001, "Bad user info", "Incorrect user information was received", nil}
	DataErrBadCountryCode = ErrorEntityEx{http.StatusBadRequest, 40002, "Bad country code", "Incorrect country code was received", nil}
	DataErrUserExists     = ErrorEntity{http.StatusBadRequest, 40003, "User already exists", "User with same email already registered"}
	DataErrFriendship     = ErrorEntity{http.StatusBadRequest, 40004, "Friendship is unavailable", "There is no possibility to create this friendship"}
	DataErrGetUserFriends = ErrorEntity{http.StatusInternalServerError, 40005, "Can not get user friends", "An error occurred while trying to get user friends"}
	DataErrQueryUsers     = ErrorEntity{http.StatusInternalServerError, 40006, "Can not query users", "An error occurred while trying to query users by filter"}

	//Database errors
	RegisterUserErr = ErrorEntity{http.StatusInternalServerError, 50001, "User register error", "User with this email already exists"}
	UpdateUserErr   = ErrorEntity{http.StatusInternalServerError, 50002, "DB error occurred", "An error occurred while trying to update user info"}
	GetCountryErr   = ErrorEntity{http.StatusInternalServerError, 50003, "DB error occurred", "An error occurred while trying to read country list"}
)

type ErrorEntity struct {
	HttpCode    int    `json:"-"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

type ErrorEntityEx struct {
	HttpCode    int         `json:"-"`
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Errors      interface{} `json:"errors"`
}
