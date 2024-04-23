package questions

import (
	"context"
	"mdgkb/ankets-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(c context.Context, item *models.Question) error {
	return R.Create(c, item)
}

func (s *Service) GetAll(c context.Context) (models.QuestionsWithCount, error) {
	items, err := R.GetAll(c)
	if err != nil {
		return items, err
	}
	return items, nil
}

func (s *Service) Get(c context.Context, id string) (*models.Question, error) {
	item, err := R.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.Question) error {
	return R.Update(c, item)
}

func (s *Service) GetAnthropometryQuestions(c context.Context) (models.Questions, error) {
	return R.GetAnthropometryQuestions(c)
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}

func (s *Service) UpsertMany(c context.Context, items models.Questions) error {
	if len(items) == 0 {
		return nil
	}

	err := R.upsertMany(c, items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	// registerPropertyRadioService := answervariants.CreateService(s.helper)
	// err = registerPropertyRadioService.UpsertMany(items.GetRegisterPropertyRadios())
	// if err != nil {
	// 	return err
	// }
	// err = registerPropertyRadioService.DeleteMany(items.GetRegisterPropertyRadioForDelete())
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s *Service) DeleteMany(c context.Context, idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return R.deleteMany(c, idPool)
}
