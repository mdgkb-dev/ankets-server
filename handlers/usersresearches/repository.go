package usersresearches

import (
	"context"

	"mdgkb/ankets-server/models"

	"github.com/pro-assistance/pro-assister/middleware"

	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.UserResearch) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.UsersResearchesWithCount, err error) {
	items.UsersResearches = make(models.UsersResearches, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items.UsersResearches).
		Relation("Research")
	// Relation("Formulas.FormulaResults")

	// query.Join("join researches_domains on researches_domains.research_id = researches.id and researches_domains.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))
	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)

	items.Count, err = query.ScanAndCount(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.UserResearch, error) {
	item := models.UserResearch{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("Research").
		Where("?TableAlias.id = ?", id).Scan(c)
	if err != nil {
		return nil, err
	}
	return &item, err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.UserResearch{}).Where("id = ?", *id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.UserResearch) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) GetForExport(c context.Context, idPool []string) (items models.UsersResearches, err error) {
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Relation("Questions", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("questions.item_order")
		}).
		Relation("Questions.AnswerVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("answer_variants.item_order")
		}).
		Relation("Questions.QuestionExamples").
		Relation("Questions.ValueType").
		Relation("Questions.QuestionVariants", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("question_variants.name")
		}).
		Relation("Questions.Children.ValueType").
		Relation("Questions.Children.AnswerVariants").
		Relation("Formulas.FormulaResults")

	if len(idPool) > 0 {
		query = query.Where("?TableAlias.id in (?)", bun.In(idPool))
	}

	query.Join("join researches_domains on researches_domains.research_id = researches.id and researches_domains.domain_id in (?)", bun.In(middleware.ClaimDomainIDS.FromContextSlice(c)))
	// r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	err = query.Scan(c)
	return items, err
}
