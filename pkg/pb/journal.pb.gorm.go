// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/kodesmil/ks-model/journal.proto

// Generated with protoc-gen-gorm version: v0.19.0
// Anticipating compatibility with atlas-app-toolkit version: v0.19.2

package pb

import context "context"
import fmt "fmt"
import time "time"

import auth1 "github.com/kodesmil/atlas-app-toolkit/auth"
import errors1 "github.com/infobloxopen/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"
import ptypes1 "github.com/golang/protobuf/ptypes"
import query1 "github.com/infobloxopen/atlas-app-toolkit/query"
import resource1 "github.com/infobloxopen/atlas-app-toolkit/gorm/resource"

import math "math"
import _ "google.golang.org/genproto/protobuf/field_mask"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"
import _ "github.com/infobloxopen/atlas-app-toolkit/query"
import _ "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type JournalSubjectORM struct {
	Id   int64 `gorm:"type:serial;primary_key"`
	Key  string
	Name string
	Type int32
}

// TableName overrides the default tablename generated by GORM
func (JournalSubjectORM) TableName() string {
	return "journal_subjects"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *JournalSubject) ToORM(ctx context.Context) (JournalSubjectORM, error) {
	to := JournalSubjectORM{}
	var err error
	if prehook, ok := interface{}(m).(JournalSubjectWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.DecodeInt64(&JournalSubject{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Key = m.Key
	to.Name = m.Name
	to.Type = int32(m.Type)
	if posthook, ok := interface{}(m).(JournalSubjectWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *JournalSubjectORM) ToPB(ctx context.Context) (JournalSubject, error) {
	to := JournalSubject{}
	var err error
	if prehook, ok := interface{}(m).(JournalSubjectWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.Encode(&JournalSubject{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Key = m.Key
	to.Name = m.Name
	to.Type = JournalSubject_Type(m.Type)
	if posthook, ok := interface{}(m).(JournalSubjectWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type JournalSubject the arg will be the target, the caller the one being converted from

// JournalSubjectBeforeToORM called before default ToORM code
type JournalSubjectWithBeforeToORM interface {
	BeforeToORM(context.Context, *JournalSubjectORM) error
}

// JournalSubjectAfterToORM called after default ToORM code
type JournalSubjectWithAfterToORM interface {
	AfterToORM(context.Context, *JournalSubjectORM) error
}

// JournalSubjectBeforeToPB called before default ToPB code
type JournalSubjectWithBeforeToPB interface {
	BeforeToPB(context.Context, *JournalSubject) error
}

// JournalSubjectAfterToPB called after default ToPB code
type JournalSubjectWithAfterToPB interface {
	AfterToPB(context.Context, *JournalSubject) error
}

type JournalEntryORM struct {
	AccountID        string
	CreatedAt        *time.Time
	Id               int64 `gorm:"type:serial;primary_key"`
	JournalSubjectId int64
	Note             string
	Severity         int32
}

// TableName overrides the default tablename generated by GORM
func (JournalEntryORM) TableName() string {
	return "journal_entries"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *JournalEntry) ToORM(ctx context.Context) (JournalEntryORM, error) {
	to := JournalEntryORM{}
	var err error
	if prehook, ok := interface{}(m).(JournalEntryWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.DecodeInt64(&JournalEntry{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Severity = int32(m.Severity)
	to.Note = m.Note
	if m.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes1.Timestamp(m.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	to.JournalSubjectId = m.JournalSubjectId
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return to, err
	}
	to.AccountID = accountID
	if posthook, ok := interface{}(m).(JournalEntryWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *JournalEntryORM) ToPB(ctx context.Context) (JournalEntry, error) {
	to := JournalEntry{}
	var err error
	if prehook, ok := interface{}(m).(JournalEntryWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	if v, err := resource1.Encode(&JournalEntry{}, m.Id); err != nil {
		return to, err
	} else {
		to.Id = v
	}
	to.Severity = JournalEntry_Severity(m.Severity)
	to.Note = m.Note
	if m.CreatedAt != nil {
		if to.CreatedAt, err = ptypes1.TimestampProto(*m.CreatedAt); err != nil {
			return to, err
		}
	}
	to.JournalSubjectId = m.JournalSubjectId
	if posthook, ok := interface{}(m).(JournalEntryWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type JournalEntry the arg will be the target, the caller the one being converted from

// JournalEntryBeforeToORM called before default ToORM code
type JournalEntryWithBeforeToORM interface {
	BeforeToORM(context.Context, *JournalEntryORM) error
}

// JournalEntryAfterToORM called after default ToORM code
type JournalEntryWithAfterToORM interface {
	AfterToORM(context.Context, *JournalEntryORM) error
}

// JournalEntryBeforeToPB called before default ToPB code
type JournalEntryWithBeforeToPB interface {
	BeforeToPB(context.Context, *JournalEntry) error
}

// JournalEntryAfterToPB called after default ToPB code
type JournalEntryWithAfterToPB interface {
	AfterToPB(context.Context, *JournalEntry) error
}

// DefaultCreateJournalSubject executes a basic gorm create call
func DefaultCreateJournalSubject(ctx context.Context, in *JournalSubject, db *gorm1.DB) (*JournalSubject, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type JournalSubjectORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadJournalSubject executes a basic gorm read call
func DefaultReadJournalSubject(ctx context.Context, in *JournalSubject, db *gorm1.DB) (*JournalSubject, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &JournalSubjectORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := JournalSubjectORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(JournalSubjectORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type JournalSubjectORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteJournalSubject(ctx context.Context, in *JournalSubject, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&JournalSubjectORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type JournalSubjectORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteJournalSubjectSet(ctx context.Context, in []*JournalSubject, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []int64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&JournalSubjectORM{})).(JournalSubjectORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&JournalSubjectORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&JournalSubjectORM{})).(JournalSubjectORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type JournalSubjectORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*JournalSubject, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*JournalSubject, *gorm1.DB) error
}

// DefaultStrictUpdateJournalSubject clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateJournalSubject(ctx context.Context, in *JournalSubject, db *gorm1.DB) (*JournalSubject, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateJournalSubject")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &JournalSubjectORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type JournalSubjectORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchJournalSubject executes a basic gorm update call with patch behavior
func DefaultPatchJournalSubject(ctx context.Context, in *JournalSubject, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*JournalSubject, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj JournalSubject
	var err error
	if hook, ok := interface{}(&pbObj).(JournalSubjectWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadJournalSubject(ctx, &JournalSubject{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(JournalSubjectWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskJournalSubject(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(JournalSubjectWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateJournalSubject(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(JournalSubjectWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type JournalSubjectWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *JournalSubject, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *JournalSubject, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *JournalSubject, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type JournalSubjectWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *JournalSubject, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetJournalSubject executes a bulk gorm update call with patch behavior
func DefaultPatchSetJournalSubject(ctx context.Context, objects []*JournalSubject, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*JournalSubject, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*JournalSubject, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchJournalSubject(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskJournalSubject patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskJournalSubject(ctx context.Context, patchee *JournalSubject, patcher *JournalSubject, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*JournalSubject, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Key" {
			patchee.Key = patcher.Key
			continue
		}
		if f == prefix+"Name" {
			patchee.Name = patcher.Name
			continue
		}
		if f == prefix+"Type" {
			patchee.Type = patcher.Type
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListJournalSubject executes a gorm list call
func DefaultListJournalSubject(ctx context.Context, db *gorm1.DB, f *query1.Filtering, s *query1.Sorting, p *query1.Pagination, fs *query1.FieldSelection) ([]*JournalSubject, error) {
	in := JournalSubject{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &JournalSubjectORM{}, &JournalSubject{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []JournalSubjectORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalSubjectORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	pbResponse := []*JournalSubject{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type JournalSubjectORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type JournalSubjectORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type JournalSubjectORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]JournalSubjectORM, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) error
}

// DefaultCreateJournalEntry executes a basic gorm create call
func DefaultCreateJournalEntry(ctx context.Context, in *JournalEntry, db *gorm1.DB) (*JournalEntry, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type JournalEntryORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadJournalEntry executes a basic gorm read call
func DefaultReadJournalEntry(ctx context.Context, in *JournalEntry, db *gorm1.DB) (*JournalEntry, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &JournalEntryORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := JournalEntryORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(JournalEntryORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type JournalEntryORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteJournalEntry(ctx context.Context, in *JournalEntry, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&JournalEntryORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type JournalEntryORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteJournalEntrySet(ctx context.Context, in []*JournalEntry, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []int64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&JournalEntryORM{})).(JournalEntryORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	acctId, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return err
	}
	err = db.Where("account_id = ? AND id in (?)", acctId, keys).Delete(&JournalEntryORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&JournalEntryORM{})).(JournalEntryORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type JournalEntryORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*JournalEntry, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*JournalEntry, *gorm1.DB) error
}

// DefaultStrictUpdateJournalEntry clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateJournalEntry(ctx context.Context, in *JournalEntry, db *gorm1.DB) (*JournalEntry, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateJournalEntry")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	accountID, err := auth1.GetAccountID(ctx, nil)
	if err != nil {
		return nil, err
	}
	db = db.Where(map[string]interface{}{"account_id": accountID})
	lockedRow := &JournalEntryORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type JournalEntryORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchJournalEntry executes a basic gorm update call with patch behavior
func DefaultPatchJournalEntry(ctx context.Context, in *JournalEntry, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*JournalEntry, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj JournalEntry
	var err error
	if hook, ok := interface{}(&pbObj).(JournalEntryWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadJournalEntry(ctx, &JournalEntry{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(JournalEntryWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskJournalEntry(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(JournalEntryWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateJournalEntry(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(JournalEntryWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type JournalEntryWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *JournalEntry, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *JournalEntry, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *JournalEntry, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type JournalEntryWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *JournalEntry, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetJournalEntry executes a bulk gorm update call with patch behavior
func DefaultPatchSetJournalEntry(ctx context.Context, objects []*JournalEntry, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*JournalEntry, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*JournalEntry, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchJournalEntry(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskJournalEntry patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskJournalEntry(ctx context.Context, patchee *JournalEntry, patcher *JournalEntry, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*JournalEntry, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Severity" {
			patchee.Severity = patcher.Severity
			continue
		}
		if f == prefix+"Note" {
			patchee.Note = patcher.Note
			continue
		}
		if f == prefix+"CreatedAt" {
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if f == prefix+"JournalSubjectId" {
			patchee.JournalSubjectId = patcher.JournalSubjectId
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListJournalEntry executes a gorm list call
func DefaultListJournalEntry(ctx context.Context, db *gorm1.DB, f *query1.Filtering, s *query1.Sorting, p *query1.Pagination, fs *query1.FieldSelection) ([]*JournalEntry, error) {
	in := JournalEntry{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &JournalEntryORM{}, &JournalEntry{}, f, s, p, fs)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []JournalEntryORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(JournalEntryORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse, f, s, p, fs); err != nil {
			return nil, err
		}
	}
	pbResponse := []*JournalEntry{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type JournalEntryORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type JournalEntryORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) (*gorm1.DB, error)
}
type JournalEntryORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]JournalEntryORM, *query1.Filtering, *query1.Sorting, *query1.Pagination, *query1.FieldSelection) error
}
type JournalEntriesDefaultServer struct {
	DB *gorm1.DB
}

// Create ...
func (m *JournalEntriesDefaultServer) Create(ctx context.Context, in *CreateJournalEntryRequest) (*CreateJournalEntryResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithBeforeCreate); ok {
		var err error
		if db, err = custom.BeforeCreate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateJournalEntry(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateJournalEntryResponse{Result: res}
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithAfterCreate); ok {
		var err error
		if err = custom.AfterCreate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// JournalEntriesJournalEntryWithBeforeCreate called before DefaultCreateJournalEntry in the default Create handler
type JournalEntriesJournalEntryWithBeforeCreate interface {
	BeforeCreate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// JournalEntriesJournalEntryWithAfterCreate called before DefaultCreateJournalEntry in the default Create handler
type JournalEntriesJournalEntryWithAfterCreate interface {
	AfterCreate(context.Context, *CreateJournalEntryResponse, *gorm1.DB) error
}

// Read ...
func (m *JournalEntriesDefaultServer) Read(ctx context.Context, in *ReadJournalEntryRequest) (*ReadJournalEntryResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithBeforeRead); ok {
		var err error
		if db, err = custom.BeforeRead(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadJournalEntry(ctx, &JournalEntry{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &ReadJournalEntryResponse{Result: res}
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithAfterRead); ok {
		var err error
		if err = custom.AfterRead(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// JournalEntriesJournalEntryWithBeforeRead called before DefaultReadJournalEntry in the default Read handler
type JournalEntriesJournalEntryWithBeforeRead interface {
	BeforeRead(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// JournalEntriesJournalEntryWithAfterRead called before DefaultReadJournalEntry in the default Read handler
type JournalEntriesJournalEntryWithAfterRead interface {
	AfterRead(context.Context, *ReadJournalEntryResponse, *gorm1.DB) error
}

// Update ...
func (m *JournalEntriesDefaultServer) Update(ctx context.Context, in *UpdateJournalEntryRequest) (*UpdateJournalEntryResponse, error) {
	var err error
	var res *JournalEntry
	db := m.DB
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithBeforeUpdate); ok {
		var err error
		if db, err = custom.BeforeUpdate(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err = DefaultStrictUpdateJournalEntry(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &UpdateJournalEntryResponse{Result: res}
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithAfterUpdate); ok {
		var err error
		if err = custom.AfterUpdate(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// JournalEntriesJournalEntryWithBeforeUpdate called before DefaultUpdateJournalEntry in the default Update handler
type JournalEntriesJournalEntryWithBeforeUpdate interface {
	BeforeUpdate(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// JournalEntriesJournalEntryWithAfterUpdate called before DefaultUpdateJournalEntry in the default Update handler
type JournalEntriesJournalEntryWithAfterUpdate interface {
	AfterUpdate(context.Context, *UpdateJournalEntryResponse, *gorm1.DB) error
}

// Delete ...
func (m *JournalEntriesDefaultServer) Delete(ctx context.Context, in *DeleteJournalEntryRequest) (*DeleteJournalEntryResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithBeforeDelete); ok {
		var err error
		if db, err = custom.BeforeDelete(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteJournalEntry(ctx, &JournalEntry{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteJournalEntryResponse{}
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithAfterDelete); ok {
		var err error
		if err = custom.AfterDelete(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// JournalEntriesJournalEntryWithBeforeDelete called before DefaultDeleteJournalEntry in the default Delete handler
type JournalEntriesJournalEntryWithBeforeDelete interface {
	BeforeDelete(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// JournalEntriesJournalEntryWithAfterDelete called before DefaultDeleteJournalEntry in the default Delete handler
type JournalEntriesJournalEntryWithAfterDelete interface {
	AfterDelete(context.Context, *DeleteJournalEntryResponse, *gorm1.DB) error
}

// List ...
func (m *JournalEntriesDefaultServer) List(ctx context.Context, in *ListJournalEntryRequest) (*ListJournalEntryResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithBeforeList); ok {
		var err error
		if db, err = custom.BeforeList(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListJournalEntry(ctx, db, in.Filter, in.OrderBy, in.Paging, in.Fields)
	if err != nil {
		return nil, err
	}
	out := &ListJournalEntryResponse{Results: res}
	if custom, ok := interface{}(in).(JournalEntriesJournalEntryWithAfterList); ok {
		var err error
		if err = custom.AfterList(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// JournalEntriesJournalEntryWithBeforeList called before DefaultListJournalEntry in the default List handler
type JournalEntriesJournalEntryWithBeforeList interface {
	BeforeList(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// JournalEntriesJournalEntryWithAfterList called before DefaultListJournalEntry in the default List handler
type JournalEntriesJournalEntryWithAfterList interface {
	AfterList(context.Context, *ListJournalEntryResponse, *gorm1.DB) error
}
type JournalSubjectsDefaultServer struct {
	DB *gorm1.DB
}

// List ...
func (m *JournalSubjectsDefaultServer) List(ctx context.Context, in *ListJournalSubjectRequest) (*ListJournalSubjectResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(JournalSubjectsJournalSubjectWithBeforeList); ok {
		var err error
		if db, err = custom.BeforeList(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListJournalSubject(ctx, db, in.Filter, in.OrderBy, in.Paging, in.Fields)
	if err != nil {
		return nil, err
	}
	out := &ListJournalSubjectResponse{Results: res}
	if custom, ok := interface{}(in).(JournalSubjectsJournalSubjectWithAfterList); ok {
		var err error
		if err = custom.AfterList(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// JournalSubjectsJournalSubjectWithBeforeList called before DefaultListJournalSubject in the default List handler
type JournalSubjectsJournalSubjectWithBeforeList interface {
	BeforeList(context.Context, *gorm1.DB) (*gorm1.DB, error)
}

// JournalSubjectsJournalSubjectWithAfterList called before DefaultListJournalSubject in the default List handler
type JournalSubjectsJournalSubjectWithAfterList interface {
	AfterList(context.Context, *ListJournalSubjectResponse, *gorm1.DB) error
}
