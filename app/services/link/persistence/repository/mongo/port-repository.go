package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/ic2hrmk/ship_links/app/services/link/persistence/model"
	"github.com/ic2hrmk/ship_links/app/services/link/persistence/repository"
)

const linksCollection = "links"

type LinkRepository struct {
	db *mgo.Database
}

func NewLinkRepository(db *mgo.Database) repository.LinkRepository {
	return &LinkRepository{db: db}
}

func (r *LinkRepository) collection() *mgo.Collection {
	return r.db.C(linksCollection)
}

func (r *LinkRepository) Save(record *model.Link) (*model.Link, error) {
	_, err := r.collection().Upsert(bson.M{"_id": record.LinkID}, record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *LinkRepository) SaveBulk(records []*model.Link) error {
	bulkOperator := r.collection().Bulk()

	for i := range records {
		bulkOperator.Upsert(bson.M{"_id": records[i].LinkID}, records[i])
	}

	if _, err := bulkOperator.Run(); err != nil {
		return err
	}

	return nil
}

func (r *LinkRepository) FindAll(limit, offset uint64) ([]*model.Link, error) {
	return r.prepareResultList(
		r.collection().Find(bson.M{}).Limit(int(limit)).Skip(int(offset)))
}

//nolint:megacheck
func (r *LinkRepository) prepareOneResult(query *mgo.Query) (*model.Link, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	record := &model.Link{}

	if count == 0 {
		return nil, repository.ErrLinkNotFound
	}

	err = query.One(&record)
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *LinkRepository) prepareResultList(query *mgo.Query) ([]*model.Link, error) {
	count, err := query.Count()
	if err != nil {
		return nil, err
	}

	var records []*model.Link

	if count == 0 {
		return records, nil
	}

	err = query.All(&records)
	if err != nil {
		return nil, err
	}

	return records, nil
}
