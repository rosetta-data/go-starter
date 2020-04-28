// Code generated by SQLBoiler 4.0.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testApplicationStates(t *testing.T) {
	t.Parallel()

	query := ApplicationStates()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testApplicationStatesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
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

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApplicationStatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ApplicationStates().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApplicationStatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ApplicationStateSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testApplicationStatesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ApplicationStateExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ApplicationState exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ApplicationStateExists to return true, but got false.")
	}
}

func testApplicationStatesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	applicationStateFound, err := FindApplicationState(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if applicationStateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testApplicationStatesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ApplicationStates().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testApplicationStatesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ApplicationStates().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testApplicationStatesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	applicationStateOne := &ApplicationState{}
	applicationStateTwo := &ApplicationState{}
	if err = randomize.Struct(seed, applicationStateOne, applicationStateDBTypes, false, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}
	if err = randomize.Struct(seed, applicationStateTwo, applicationStateDBTypes, false, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = applicationStateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = applicationStateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ApplicationStates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testApplicationStatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	applicationStateOne := &ApplicationState{}
	applicationStateTwo := &ApplicationState{}
	if err = randomize.Struct(seed, applicationStateOne, applicationStateDBTypes, false, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}
	if err = randomize.Struct(seed, applicationStateTwo, applicationStateDBTypes, false, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = applicationStateOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = applicationStateTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testApplicationStatesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testApplicationStatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(applicationStateColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testApplicationStateToManyApplicants(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationState
	var b, c Applicant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, applicantDBTypes, false, applicantColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, applicantDBTypes, false, applicantColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ApplicationStateID = a.ID
	c.ApplicationStateID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Applicants().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ApplicationStateID == b.ApplicationStateID {
			bFound = true
		}
		if v.ApplicationStateID == c.ApplicationStateID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ApplicationStateSlice{&a}
	if err = a.L.LoadApplicants(ctx, tx, false, (*[]*ApplicationState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Applicants); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Applicants = nil
	if err = a.L.LoadApplicants(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Applicants); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testApplicationStateToManyFromApplicationStateApplicationStateTransitions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationState
	var b, c ApplicationStateTransition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.FromApplicationStateID = a.ID
	c.FromApplicationStateID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.FromApplicationStateApplicationStateTransitions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.FromApplicationStateID == b.FromApplicationStateID {
			bFound = true
		}
		if v.FromApplicationStateID == c.FromApplicationStateID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ApplicationStateSlice{&a}
	if err = a.L.LoadFromApplicationStateApplicationStateTransitions(ctx, tx, false, (*[]*ApplicationState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FromApplicationStateApplicationStateTransitions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FromApplicationStateApplicationStateTransitions = nil
	if err = a.L.LoadFromApplicationStateApplicationStateTransitions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FromApplicationStateApplicationStateTransitions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testApplicationStateToManyToApplicationStateApplicationStateTransitions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationState
	var b, c ApplicationStateTransition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, applicationStateTransitionDBTypes, false, applicationStateTransitionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ToApplicationStateID = a.ID
	c.ToApplicationStateID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ToApplicationStateApplicationStateTransitions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ToApplicationStateID == b.ToApplicationStateID {
			bFound = true
		}
		if v.ToApplicationStateID == c.ToApplicationStateID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ApplicationStateSlice{&a}
	if err = a.L.LoadToApplicationStateApplicationStateTransitions(ctx, tx, false, (*[]*ApplicationState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ToApplicationStateApplicationStateTransitions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ToApplicationStateApplicationStateTransitions = nil
	if err = a.L.LoadToApplicationStateApplicationStateTransitions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ToApplicationStateApplicationStateTransitions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testApplicationStateToManyRoleApplicationStates(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationState
	var b, c RoleApplicationState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, roleApplicationStateDBTypes, false, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roleApplicationStateDBTypes, false, roleApplicationStateColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ApplicationStateID = a.ID
	c.ApplicationStateID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RoleApplicationStates().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ApplicationStateID == b.ApplicationStateID {
			bFound = true
		}
		if v.ApplicationStateID == c.ApplicationStateID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ApplicationStateSlice{&a}
	if err = a.L.LoadRoleApplicationStates(ctx, tx, false, (*[]*ApplicationState)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RoleApplicationStates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RoleApplicationStates = nil
	if err = a.L.LoadRoleApplicationStates(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RoleApplicationStates); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testApplicationStateToManyAddOpApplicants(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationState
	var b, c, d, e Applicant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Applicant{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, applicantDBTypes, false, strmangle.SetComplement(applicantPrimaryKeyColumns, applicantColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Applicant{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddApplicants(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ApplicationStateID {
			t.Error("foreign key was wrong value", a.ID, first.ApplicationStateID)
		}
		if a.ID != second.ApplicationStateID {
			t.Error("foreign key was wrong value", a.ID, second.ApplicationStateID)
		}

		if first.R.ApplicationState != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ApplicationState != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Applicants[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Applicants[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Applicants().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testApplicationStateToManyAddOpFromApplicationStateApplicationStateTransitions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationState
	var b, c, d, e ApplicationStateTransition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ApplicationStateTransition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, applicationStateTransitionDBTypes, false, strmangle.SetComplement(applicationStateTransitionPrimaryKeyColumns, applicationStateTransitionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ApplicationStateTransition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFromApplicationStateApplicationStateTransitions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.FromApplicationStateID {
			t.Error("foreign key was wrong value", a.ID, first.FromApplicationStateID)
		}
		if a.ID != second.FromApplicationStateID {
			t.Error("foreign key was wrong value", a.ID, second.FromApplicationStateID)
		}

		if first.R.FromApplicationState != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.FromApplicationState != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FromApplicationStateApplicationStateTransitions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FromApplicationStateApplicationStateTransitions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FromApplicationStateApplicationStateTransitions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testApplicationStateToManyAddOpToApplicationStateApplicationStateTransitions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationState
	var b, c, d, e ApplicationStateTransition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ApplicationStateTransition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, applicationStateTransitionDBTypes, false, strmangle.SetComplement(applicationStateTransitionPrimaryKeyColumns, applicationStateTransitionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ApplicationStateTransition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddToApplicationStateApplicationStateTransitions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ToApplicationStateID {
			t.Error("foreign key was wrong value", a.ID, first.ToApplicationStateID)
		}
		if a.ID != second.ToApplicationStateID {
			t.Error("foreign key was wrong value", a.ID, second.ToApplicationStateID)
		}

		if first.R.ToApplicationState != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ToApplicationState != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ToApplicationStateApplicationStateTransitions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ToApplicationStateApplicationStateTransitions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ToApplicationStateApplicationStateTransitions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testApplicationStateToManyAddOpRoleApplicationStates(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ApplicationState
	var b, c, d, e RoleApplicationState

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, applicationStateDBTypes, false, strmangle.SetComplement(applicationStatePrimaryKeyColumns, applicationStateColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RoleApplicationState{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, roleApplicationStateDBTypes, false, strmangle.SetComplement(roleApplicationStatePrimaryKeyColumns, roleApplicationStateColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*RoleApplicationState{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRoleApplicationStates(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ApplicationStateID {
			t.Error("foreign key was wrong value", a.ID, first.ApplicationStateID)
		}
		if a.ID != second.ApplicationStateID {
			t.Error("foreign key was wrong value", a.ID, second.ApplicationStateID)
		}

		if first.R.ApplicationState != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ApplicationState != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RoleApplicationStates[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RoleApplicationStates[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RoleApplicationStates().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testApplicationStatesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
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

func testApplicationStatesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ApplicationStateSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testApplicationStatesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ApplicationStates().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	applicationStateDBTypes = map[string]string{`ID`: `uuid`, `State`: `text`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                       = bytes.MinRead
)

func testApplicationStatesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(applicationStatePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(applicationStateAllColumns) == len(applicationStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testApplicationStatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(applicationStateAllColumns) == len(applicationStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ApplicationState{}
	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, applicationStateDBTypes, true, applicationStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(applicationStateAllColumns, applicationStatePrimaryKeyColumns) {
		fields = applicationStateAllColumns
	} else {
		fields = strmangle.SetComplement(
			applicationStateAllColumns,
			applicationStatePrimaryKeyColumns,
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

	slice := ApplicationStateSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testApplicationStatesUpsert(t *testing.T) {
	t.Parallel()

	if len(applicationStateAllColumns) == len(applicationStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ApplicationState{}
	if err = randomize.Struct(seed, &o, applicationStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ApplicationState: %s", err)
	}

	count, err := ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, applicationStateDBTypes, false, applicationStatePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ApplicationState struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ApplicationState: %s", err)
	}

	count, err = ApplicationStates().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
