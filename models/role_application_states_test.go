// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRoleApplicationStates(t *testing.T) {
	t.Parallel()

	query := RoleApplicationStates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRoleApplicationStatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleApplicationStatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RoleApplicationStates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleApplicationStatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleApplicationStateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoleApplicationStatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RoleApplicationStateExists(ctx, tx, o.RoleID, o.ApplicationStateID)
	if err != nil {
		t.Errorf("Unable to check if RoleApplicationState exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RoleApplicationStateExists to return true, but got false.")
	}
}

func testRoleApplicationStatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	roleApplicationStateFound, err := FindRoleApplicationState(ctx, tx, o.RoleID, o.ApplicationStateID)
	if err != nil {
		t.Error(err)
	}

	if roleApplicationStateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRoleApplicationStatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RoleApplicationStates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRoleApplicationStatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RoleApplicationStates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRoleApplicationStatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	roleApplicationStateOne := &RoleApplicationState{}
	roleApplicationStateTwo := &RoleApplicationState{}
	if err = randomize.Struct(seed, roleApplicationStateOne, roleApplicationStateDBTypes, false, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}
	if err = randomize.Struct(seed, roleApplicationStateTwo, roleApplicationStateDBTypes, false, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roleApplicationStateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleApplicationStateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoleApplicationStates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRoleApplicationStatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	roleApplicationStateOne := &RoleApplicationState{}
	roleApplicationStateTwo := &RoleApplicationState{}
	if err = randomize.Struct(seed, roleApplicationStateOne, roleApplicationStateDBTypes, false, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}
	if err = randomize.Struct(seed, roleApplicationStateTwo, roleApplicationStateDBTypes, false, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roleApplicationStateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roleApplicationStateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func roleApplicationStateBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func roleApplicationStateAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func roleApplicationStateAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func roleApplicationStateBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func roleApplicationStateAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func roleApplicationStateBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func roleApplicationStateAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func roleApplicationStateBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func roleApplicationStateAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RoleApplicationState) error {
	*o = RoleApplicationState{}
	return nil
}

func testRoleApplicationStatesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RoleApplicationState{}
	o := &RoleApplicationState{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState object: %s", err)
	}

	AddRoleApplicationStateHook(boil.BeforeInsertHook, roleApplicationStateBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateBeforeInsertHooks = []RoleApplicationStateHook{}

	AddRoleApplicationStateHook(boil.AfterInsertHook, roleApplicationStateAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateAfterInsertHooks = []RoleApplicationStateHook{}

	AddRoleApplicationStateHook(boil.AfterSelectHook, roleApplicationStateAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateAfterSelectHooks = []RoleApplicationStateHook{}

	AddRoleApplicationStateHook(boil.BeforeUpdateHook, roleApplicationStateBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateBeforeUpdateHooks = []RoleApplicationStateHook{}

	AddRoleApplicationStateHook(boil.AfterUpdateHook, roleApplicationStateAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateAfterUpdateHooks = []RoleApplicationStateHook{}

	AddRoleApplicationStateHook(boil.BeforeDeleteHook, roleApplicationStateBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateBeforeDeleteHooks = []RoleApplicationStateHook{}

	AddRoleApplicationStateHook(boil.AfterDeleteHook, roleApplicationStateAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateAfterDeleteHooks = []RoleApplicationStateHook{}

	AddRoleApplicationStateHook(boil.BeforeUpsertHook, roleApplicationStateBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateBeforeUpsertHooks = []RoleApplicationStateHook{}

	AddRoleApplicationStateHook(boil.AfterUpsertHook, roleApplicationStateAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	roleApplicationStateAfterUpsertHooks = []RoleApplicationStateHook{}
}

func testRoleApplicationStatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoleApplicationStatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(roleApplicationStateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoleApplicationStateToOneApplicationStateUsingApplicationState(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoleApplicationState
	var foreign ApplicationState

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, roleApplicationStateDBTypes, false, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, applicationStateDBTypes, false, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ApplicationStateID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ApplicationState().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RoleApplicationStateSlice{&local}
	if err = local.L.LoadApplicationState(ctx, tx, false, (*[]*RoleApplicationState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ApplicationState == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ApplicationState = nil
	if err = local.L.LoadApplicationState(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ApplicationState == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRoleApplicationStateToOneRoleUsingRole(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoleApplicationState
	var foreign Role

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, roleApplicationStateDBTypes, false, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, roleDBTypes, false, roleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Role struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RoleID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Role().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RoleApplicationStateSlice{&local}
	if err = local.L.LoadRole(ctx, tx, false, (*[]*RoleApplicationState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Role = nil
	if err = local.L.LoadRole(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Role == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRoleApplicationStateToOneSetOpApplicationStateUsingApplicationState(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoleApplicationState
	var b, c ApplicationState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleApplicationStateDBTypes, false, strmangle.SetComplement(roleApplicationStatePrimaryKeyColumns, roleApplicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ApplicationState{&b, &c} {
		err = a.SetApplicationState(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ApplicationState != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoleApplicationStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ApplicationStateID != x.ID {
			t.Error("foreign key was wrong value", a.ApplicationStateID)
		}

		if exists, err := RoleApplicationStateExists(ctx, tx, a.RoleID, a.ApplicationStateID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testRoleApplicationStateToOneSetOpRoleUsingRole(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoleApplicationState
	var b, c Role

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roleApplicationStateDBTypes, false, strmangle.SetComplement(roleApplicationStatePrimaryKeyColumns, roleApplicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleDBTypes, false, strmangle.SetComplement(rolePrimaryKeyColumns, roleColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Role{&b, &c} {
		err = a.SetRole(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Role != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoleApplicationStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RoleID != x.ID {
			t.Error("foreign key was wrong value", a.RoleID)
		}

		if exists, err := RoleApplicationStateExists(ctx, tx, a.RoleID, a.ApplicationStateID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testRoleApplicationStatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoleApplicationStatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoleApplicationStateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoleApplicationStatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoleApplicationStates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	roleApplicationStateDBTypes = map[string]string{`RoleID`: `uuid`, `ApplicationStateID`: `uuid`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                           = bytes.MinRead
)

func testRoleApplicationStatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(roleApplicationStatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(roleApplicationStateAllColumns) == len(roleApplicationStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRoleApplicationStatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(roleApplicationStateAllColumns) == len(roleApplicationStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoleApplicationState{}
	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roleApplicationStateDBTypes, true, roleApplicationStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(roleApplicationStateAllColumns, roleApplicationStatePrimaryKeyColumns) {
		fields = roleApplicationStateAllColumns
	} else {
		fields = strmangle.SetComplement(
			roleApplicationStateAllColumns,
			roleApplicationStatePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RoleApplicationStateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRoleApplicationStatesUpsert(t *testing.T) {
	t.Parallel()

	if len(roleApplicationStateAllColumns) == len(roleApplicationStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RoleApplicationState{}
	if err = randomize.Struct(seed, &o, roleApplicationStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoleApplicationState: %s", err)
	}

	count, err := RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, roleApplicationStateDBTypes, false, roleApplicationStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoleApplicationState struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoleApplicationState: %s", err)
	}

	count, err = RoleApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
