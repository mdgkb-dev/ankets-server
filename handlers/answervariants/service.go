package answervariants

import (
	"mdgkb/ankets-server/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.AnswerVariant) error {
	return R.create(item)
}

func (s *Service) GetAll() ([]*models.AnswerVariant, error) {
	items, err := R.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.AnswerVariant, error) {
	item, err := R.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.AnswerVariant) error {
	return R.update(item)
}

func (s *Service) Delete(id *string) error {
	return R.delete(id)
}

func (s *Service) UpsertMany(items models.AnswerVariants) error {
	if len(items) == 0 {
		return nil
	}
	err := R.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIDForChildren()
	//err = registerPropertyOthersService.UpsertMany(items.GetRegisterPropertyOthers())
	//if err != nil {
	//	return err
	//}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return R.deleteMany(idPool)
}
