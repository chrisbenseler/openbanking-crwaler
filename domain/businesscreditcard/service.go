package businesscreditcard

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/subentities"
)

//Service electronicchannel service
type Service interface {
	DeleteAllFromInstitution(string) common.CustomError
	InsertMany([]Entity, string) common.CustomError
	FindByInstitution(string, int) ([]Entity, *subentities.Pagination, common.CustomError)
}

type service struct {
	repository Repository
}

//NewService create a new service for branches
func NewService(repository Repository) Service {

	return &service{
		repository: repository,
	}
}

//DeleteAllFromInstitution update branches from institution
func (s *service) DeleteAllFromInstitution(InstitutionID string) common.CustomError {

	deleteErr := s.repository.DeleteMany(InstitutionID)

	if deleteErr != nil {
		return common.NewInternalServerError("Could not delete businessCreditCards from institution", deleteErr)
	}

	return nil
}

//InsertMany insert many businessCreditCards in instition
func (s *service) InsertMany(businessCreditCards []Entity, institututionID string) common.CustomError {

	for _, businessCreditCard := range businessCreditCards {
		businessCreditCard.InstitutionID = institututionID
		s.repository.Save(businessCreditCard)
	}

	return nil
}

//FindByInstitution insert many branches in instition
func (s *service) FindByInstitution(institututionID string, page int) ([]Entity, *subentities.Pagination, common.CustomError) {
	return s.repository.FindByInstitution(institututionID, page)
}
