package leveldbstore

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"

	"calendar/event"
)

type leveldbEventRepository struct {
	db *leveldb.DB
}

func NewLevelDBEventRepository(db *leveldb.DB) event.EventRepository {
	return &leveldbEventRepository{
		db: db,
	}
}

func NewLevelDB(path string) (*leveldb.DB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (l *leveldbEventRepository) Create(user_id uint64, e event.Event) (event.Event, error) {
	userPrefix := itob(user_id)
	eventID, err := l.db.Get(userPrefix, nil)
	if err != nil && err != leveldb.ErrNotFound {
		return event.Event{}, err
	}

	e.ID = binary.BigEndian.Uint64(eventID) + 1
	eventKey := append(userPrefix, itob(e.ID)...)
	eventValue, err := json.Marshal(e)
	if err != nil {
		return event.Event{}, err
	}

	err = l.db.Put(eventKey, eventValue, nil)
	if err != nil {
		return event.Event{}, err
	}

	return e, nil
}

func (l *leveldbEventRepository) Update(user_id uint64, e event.Event) error {
	eventKey := append(itob(user_id), itob(e.ID)...)
	eventValue, err := json.Marshal(e)
	if err != nil {
		return err
	}

	return l.db.Put(eventKey, eventValue, nil)
}

func (l *leveldbEventRepository) Delete(user_id uint64, event_id uint64) error {
	eventKey := append(itob(user_id), itob(event_id)...)
	return l.db.Delete(eventKey, nil)
}

func (l *leveldbEventRepository) GetForDay(user_id uint64, day time.Time) ([]event.Event, error) {
	var events []event.Event
	prefix := itob(user_id)
	iter := l.db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		var ev event.Event
		err := json.Unmarshal(iter.Value(), &ev)
		if err != nil {
			return nil, err
		}

		if ev.Date.Year() == day.Year() && ev.Date.Month() == day.Month() && ev.Date.Day() == day.Day() {
			events = append(events, ev)
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (l *leveldbEventRepository) GetForWeek(user_id uint64, week time.Time) ([]event.Event, error) {
	startOfWeek := startOfISOWeek(week)
	endOfWeek := startOfWeek.AddDate(0, 0, 7)
	var events []event.Event
	prefix := itob(user_id)
	iter := l.db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		var ev event.Event
		err := json.Unmarshal(iter.Value(), &ev)
		if err != nil {
			return nil, err
		}
		if ev.Date.After(startOfWeek) && ev.Date.Before(endOfWeek) {
			events = append(events, ev)
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (l *leveldbEventRepository) GetForMonth(user_id uint64, month time.Time) ([]event.Event, error) {
	var events []event.Event
	prefix := itob(user_id)
	iter := l.db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		var ev event.Event
		err := json.Unmarshal(iter.Value(), &ev)
		if err != nil {
			return nil, err
		}
		if ev.Date.Year() == month.Year() && ev.Date.Month() == month.Month() {
			events = append(events, ev)
		}
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, err
	}
	return events, nil
}

// Helper functions

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// startOfISOWeek returns the start of the week (Monday) for the given time.
func startOfISOWeek(t time.Time) time.Time {
	year, week := t.ISOWeek()
	// Get the first day of the year
	startOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	// Get the ISO Weekday of the first day of the year
	isoWeekStartDay := int(startOfYear.Weekday())
	if isoWeekStartDay == 0 {
		isoWeekStartDay = 7
	}
	// Calculate the offset to the start of the ISO Week
	offset := (7 - isoWeekStartDay) + (week-1)*7
	// Calculate the date of the Monday of the ISO Week
	return startOfYear.AddDate(0, 0, offset)
}
