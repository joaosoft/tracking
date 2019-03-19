package tracking

import (
	"github.com/joaosoft/dbr"
)

type StoragePostgres struct {
	config *TrackingConfig
	db     *dbr.Dbr
}

func NewStoragePostgres(config *TrackingConfig) (*StoragePostgres, error) {
	dbr, err := dbr.New(dbr.WithConfiguration(config.Dbr))
	if err != nil {
		return nil, err
	}

	return &StoragePostgres{
		config: config,
		db:     dbr,
	}, nil
}

func (storage *StoragePostgres) AddEvent(event *Event) error {

	// category
	exists, category, err := storage.GetCategoryByName(event.Category)
	if err != nil {
		return err
	}

	if !exists {
		category = &Category{
			IdCategory: genUI(),
			Name:       event.Category,
		}

		if err = storage.AddCategory(category); err != nil {
			return err
		}
	}

	// action
	exists, action, err := storage.GetActionByName(event.Action)
	if err != nil {
		return err
	}

	if !exists {
		action = &Action{
			IdAction: genUI(),
			Name:     event.Action,
		}

		if err = storage.AddAction(action); err != nil {
			return err
		}
	}

	// association
	if err = storage.UpsertCategoryActionAssociation(category, action); err != nil {
		return err
	}

	// event
	event.IdEvent = genUI()
	event.FkCategory = category.IdCategory
	event.FkAction = action.IdAction

	_, err = storage.db.
		Insert().
		Into("event").
		Record(event).
		Exec()

	if err != nil {
		return err
	}

	return nil
}

func (storage *StoragePostgres) GetCategoryByName(name string) (exists bool, category *Category, err error) {
	category = &Category{}

	count, err := storage.db.
		Select("*").
		From("category").
		Load(category)

	if err != nil {
		return false, nil, err
	}

	if count == 0 {
		return false, nil, nil
	}

	return true, category, nil
}

func (storage *StoragePostgres) AddCategory(category *Category) error {
	err := storage.db.
		Insert().
		Into("category").
		Record(category).
		Return("id_category").
		Load(&category.IdCategory)

	if err != nil {
		return err
	}

	return nil
}

func (storage *StoragePostgres) GetActionByName(name string) (exists bool, action *Action, err error) {
	action = &Action{}

	count, err := storage.db.
		Select("*").
		From("action").
		Load(action)

	if err != nil {
		return false, nil, err
	}

	if count == 0 {
		return false, nil, nil
	}

	return true, action, nil
}

func (storage *StoragePostgres) AddAction(action *Action) error {
	err := storage.db.
		Insert().
		Into("action").
		Record(action).
		Return("id_action").
		Load(&action.IdAction)

	if err != nil {
		return err
	}

	return nil
}

func (storage *StoragePostgres) UpsertCategoryActionAssociation(category *Category, action *Action) error {
	_, err := storage.db.
		Insert().
		Into("category_action").
		Columns([]interface{}{"fk_category", "fk_action"}...).
		Values(category.IdCategory, action.IdAction).
		OnConflict([]interface{}{"fk_category", "fk_action"}...).
		DoNothing().
		Exec()

	if err != nil {
		return err
	}

	return nil
}
