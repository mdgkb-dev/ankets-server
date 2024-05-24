package humans

import (
	"context"
	"mdgkb/ankets-server/models"
)

func (s *Service) Create(c context.Context, item *models.Human) error {
	err := R.Create(c, item)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll(c context.Context) (models.HumansWithCount, error) {
	return R.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Human, error) {
	return R.Get(c, id)
}

func (s *Service) GetByHumanAccountID(c context.Context, id string) (*models.Human, error) {
	return R.GetByHumanAccountID(c, id)
}

func (s *Service) Update(c context.Context, item *models.Human) error {
	err := R.Update(c, item)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return R.Delete(c, id)
}
