package service

import (
	"github.com/jjeanjacques10/crud-golang/src/configuration/rest_err"
	"github.com/jjeanjacques10/crud-golang/src/model"
)

func (*userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
