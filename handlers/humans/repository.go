package humans

import (
	"context"
	"mdgkb/ankets-server/models"
)

func (r *Repository) GetAll(c context.Context) (item models.HumansWithCount, err error) {
	item.Humans = make(models.Humans, 0)
	q := r.helper.DB.IDB(c).NewSelect().
		Model(&item.Humans).
		Relation("Human").
		Relation("HumanAccount")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	item.Count, err = q.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Human, error) {
	item := models.Human{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Human").
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) GetByHumanAccountID(c context.Context, id string) (*models.Human, error) {
	item := models.Human{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.user_account_id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Human{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Human) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) Create(c context.Context, item *models.Human) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}
