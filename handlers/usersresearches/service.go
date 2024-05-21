package usersresearches

import (
	"context"
	"mdgkb/ankets-server/models"
)

func (s *Service) Create(c context.Context, item *models.UserResearch) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	// item.SetIDForChildren()
	// err = UserResearchgroup.CreateService(s.helper).UpsertMany(item.UserResearchGroups)
	if err != nil {
		return err
	}
	//err = UserResearchdiagnosis.CreateService(s.helper).CreateMany(item.UserResearchDiagnosis)
	//if err != nil {
	//	return err
	//}
	return err
}

func (s *Service) GetAll(c context.Context) (models.UsersResearchesWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.UserResearch, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.UserResearch) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	// item.SetIDForChildren()

	//UserResearchGroupService := UserResearchgroup.CreateService(s.helper)
	//err = UserResearchGroupService.UpsertMany(item.UserResearchGroups)
	//if err != nil {
	//	return err
	//}
	//err = UserResearchGroupService.DeleteMany(item.UsersResearchesForDelete)
	//if err != nil {
	//	return err
	//}

	//UserResearchDiagnosisService := UserResearchdiagnosis.CreateService(s.helper)
	//err = UserResearchDiagnosisService.UpsertMany(item.UserResearchDiagnosis)
	//if err != nil {
	//	return err
	//}
	//err = UserResearchDiagnosisService.DeleteMany(item.UserResearchDiagnosisForDelete)
	//if err != nil {
	//	return err
	//}
	return err
}

func (s *Service) Delete(c context.Context, id *string) error {
	return R.Delete(c, id)
}

// func (s *Service) GetUserResearchAndPatient(c context.Context, researchID string, patientID string) (*models.UserResearch, *models.Patient, error) {
// 	research, err := R.Get(c, researchID)
// 	if err != nil {
// 		return nil, nil, err
// 	}
//
// 	// patient, err := patients.S.Get(c, patientID)
// 	// if err != nil {
// 	// 	return nil, nil, err
// 	// }
// 	return research, nil, nil
// }
