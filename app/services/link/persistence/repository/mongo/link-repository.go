package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pkg/errors"

	"github.com/links-123/links123/app/services/link/persistence/model"
	"github.com/links-123/links123/app/services/link/persistence/repository"
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
		return nil, errors.WithMessage(repository.ErrStorageFailure, err.Error())
	}

	return record, nil
}

func (r *LinkRepository) Delete(linkID, userID string) error {
	err := r.collection().Remove(
		bson.M{
			"_id":    linkID,
			"userID": userID,
		})

	if err != nil && err != mgo.ErrNotFound {
		return errors.WithMessage(repository.ErrStorageFailure, err.Error())
	}

	return nil
}

func (r *LinkRepository) FindByUserID(userID string, limit, offset uint64) ([]*model.Link, uint64, error) {
	userLinksQuery := r.collection().
		Find(bson.M{
			"userID": userID,
		})

	totalLinks, err := r.countResults(userLinksQuery)
	if err != nil {
		return nil, 0, err
	}

	links, err := r.prepareResultList(
		userLinksQuery.
			Limit(int(limit)).
			Skip(int(offset)),
	)
	if err != nil {
		return nil, 0, err
	}

	return links, totalLinks, nil
}

//nolint:megacheck
func (r *LinkRepository) prepareOneResult(query *mgo.Query) (*model.Link, error) {
	count, err := query.Count()
	if err != nil {
		return nil, errors.WithMessage(repository.ErrStorageFailure, err.Error())
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
	var records []*model.Link

	if err := query.All(&records); err != nil {
		return nil, errors.WithMessage(repository.ErrStorageFailure, err.Error())
	}

	return records, nil
}

func (r *LinkRepository) countResults(query *mgo.Query) (uint64, error) {
	results, err := query.Count()
	if err != nil {
		return 0, errors.WithMessage(repository.ErrStorageFailure, err.Error())
	}

	return uint64(results), nil
}
