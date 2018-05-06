// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// OutputClaim is an object representing the database table.
type OutputClaim struct {
	OutputID string      `boil:"output_id" json:"output_id" toml:"output_id" yaml:"output_id"`
	VoutID   uint        `boil:"vout_id" json:"vout_id" toml:"vout_id" yaml:"vout_id"`
	ClaimID  string      `boil:"claim_id" json:"claim_id" toml:"claim_id" yaml:"claim_id"`
	Type     null.String `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`

	R *outputClaimR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L outputClaimL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OutputClaimColumns = struct {
	OutputID string
	VoutID   string
	ClaimID  string
	Type     string
}{
	OutputID: "output_id",
	VoutID:   "vout_id",
	ClaimID:  "claim_id",
	Type:     "type",
}

// outputClaimR is where relationships are stored.
type outputClaimR struct {
	Output *Output
	Vout   *Output
	Claim  *Claim
}

// outputClaimL is where Load methods for each relationship are stored.
type outputClaimL struct{}

var (
	outputClaimColumns               = []string{"output_id", "vout_id", "claim_id", "type"}
	outputClaimColumnsWithoutDefault = []string{"output_id", "vout_id", "claim_id"}
	outputClaimColumnsWithDefault    = []string{"type"}
	outputClaimPrimaryKeyColumns     = []string{"output_id", "vout_id", "claim_id"}
)

type (
	// OutputClaimSlice is an alias for a slice of pointers to OutputClaim.
	// This should generally be used opposed to []OutputClaim.
	OutputClaimSlice []*OutputClaim

	outputClaimQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	outputClaimType                 = reflect.TypeOf(&OutputClaim{})
	outputClaimMapping              = queries.MakeStructMapping(outputClaimType)
	outputClaimPrimaryKeyMapping, _ = queries.BindMapping(outputClaimType, outputClaimMapping, outputClaimPrimaryKeyColumns)
	outputClaimInsertCacheMut       sync.RWMutex
	outputClaimInsertCache          = make(map[string]insertCache)
	outputClaimUpdateCacheMut       sync.RWMutex
	outputClaimUpdateCache          = make(map[string]updateCache)
	outputClaimUpsertCacheMut       sync.RWMutex
	outputClaimUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)

// OneP returns a single outputClaim record from the query, and panics on error.
func (q outputClaimQuery) OneP() *OutputClaim {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single outputClaim record from the query.
func (q outputClaimQuery) One() (*OutputClaim, error) {
	o := &OutputClaim{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for output_claim")
	}

	return o, nil
}

// AllP returns all OutputClaim records from the query, and panics on error.
func (q outputClaimQuery) AllP() OutputClaimSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all OutputClaim records from the query.
func (q outputClaimQuery) All() (OutputClaimSlice, error) {
	var o []*OutputClaim

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to OutputClaim slice")
	}

	return o, nil
}

// CountP returns the count of all OutputClaim records in the query, and panics on error.
func (q outputClaimQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all OutputClaim records in the query.
func (q outputClaimQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count output_claim rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q outputClaimQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q outputClaimQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if output_claim exists")
	}

	return count > 0, nil
}

// OutputG pointed to by the foreign key.
func (o *OutputClaim) OutputG(mods ...qm.QueryMod) outputQuery {
	return o.Output(boil.GetDB(), mods...)
}

// Output pointed to by the foreign key.
func (o *OutputClaim) Output(exec boil.Executor, mods ...qm.QueryMod) outputQuery {
	queryMods := []qm.QueryMod{
		qm.Where("transaction_hash=?", o.OutputID),
	}

	queryMods = append(queryMods, mods...)

	query := Outputs(exec, queryMods...)
	queries.SetFrom(query.Query, "`output`")

	return query
}

// VoutG pointed to by the foreign key.
func (o *OutputClaim) VoutG(mods ...qm.QueryMod) outputQuery {
	return o.Vout(boil.GetDB(), mods...)
}

// Vout pointed to by the foreign key.
func (o *OutputClaim) Vout(exec boil.Executor, mods ...qm.QueryMod) outputQuery {
	queryMods := []qm.QueryMod{
		qm.Where("vout=?", o.VoutID),
	}

	queryMods = append(queryMods, mods...)

	query := Outputs(exec, queryMods...)
	queries.SetFrom(query.Query, "`output`")

	return query
}

// ClaimG pointed to by the foreign key.
func (o *OutputClaim) ClaimG(mods ...qm.QueryMod) claimQuery {
	return o.Claim(boil.GetDB(), mods...)
}

// Claim pointed to by the foreign key.
func (o *OutputClaim) Claim(exec boil.Executor, mods ...qm.QueryMod) claimQuery {
	queryMods := []qm.QueryMod{
		qm.Where("claim_id=?", o.ClaimID),
	}

	queryMods = append(queryMods, mods...)

	query := Claims(exec, queryMods...)
	queries.SetFrom(query.Query, "`claim`")

	return query
} // LoadOutput allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (outputClaimL) LoadOutput(e boil.Executor, singular bool, maybeOutputClaim interface{}) error {
	var slice []*OutputClaim
	var object *OutputClaim

	count := 1
	if singular {
		object = maybeOutputClaim.(*OutputClaim)
	} else {
		slice = *maybeOutputClaim.(*[]*OutputClaim)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &outputClaimR{}
		}
		args[0] = object.OutputID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &outputClaimR{}
			}
			args[i] = obj.OutputID
		}
	}

	query := fmt.Sprintf(
		"select * from `output` where `transaction_hash` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Output")
	}
	defer results.Close()

	var resultSlice []*Output
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Output")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Output = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OutputID == foreign.TransactionHash {
				local.R.Output = foreign
				break
			}
		}
	}

	return nil
}

// LoadVout allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (outputClaimL) LoadVout(e boil.Executor, singular bool, maybeOutputClaim interface{}) error {
	var slice []*OutputClaim
	var object *OutputClaim

	count := 1
	if singular {
		object = maybeOutputClaim.(*OutputClaim)
	} else {
		slice = *maybeOutputClaim.(*[]*OutputClaim)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &outputClaimR{}
		}
		args[0] = object.VoutID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &outputClaimR{}
			}
			args[i] = obj.VoutID
		}
	}

	query := fmt.Sprintf(
		"select * from `output` where `vout` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Output")
	}
	defer results.Close()

	var resultSlice []*Output
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Output")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Vout = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.VoutID == foreign.Vout {
				local.R.Vout = foreign
				break
			}
		}
	}

	return nil
}

// LoadClaim allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (outputClaimL) LoadClaim(e boil.Executor, singular bool, maybeOutputClaim interface{}) error {
	var slice []*OutputClaim
	var object *OutputClaim

	count := 1
	if singular {
		object = maybeOutputClaim.(*OutputClaim)
	} else {
		slice = *maybeOutputClaim.(*[]*OutputClaim)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &outputClaimR{}
		}
		args[0] = object.ClaimID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &outputClaimR{}
			}
			args[i] = obj.ClaimID
		}
	}

	query := fmt.Sprintf(
		"select * from `claim` where `claim_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Claim")
	}
	defer results.Close()

	var resultSlice []*Claim
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Claim")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		object.R.Claim = resultSlice[0]
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ClaimID == foreign.ClaimID {
				local.R.Claim = foreign
				break
			}
		}
	}

	return nil
}

// SetOutputG of the output_claim to the related item.
// Sets o.R.Output to related.
// Adds o to related.R.OutputClaims.
// Uses the global database handle.
func (o *OutputClaim) SetOutputG(insert bool, related *Output) error {
	return o.SetOutput(boil.GetDB(), insert, related)
}

// SetOutputP of the output_claim to the related item.
// Sets o.R.Output to related.
// Adds o to related.R.OutputClaims.
// Panics on error.
func (o *OutputClaim) SetOutputP(exec boil.Executor, insert bool, related *Output) {
	if err := o.SetOutput(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetOutputGP of the output_claim to the related item.
// Sets o.R.Output to related.
// Adds o to related.R.OutputClaims.
// Uses the global database handle and panics on error.
func (o *OutputClaim) SetOutputGP(insert bool, related *Output) {
	if err := o.SetOutput(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetOutput of the output_claim to the related item.
// Sets o.R.Output to related.
// Adds o to related.R.OutputClaims.
func (o *OutputClaim) SetOutput(exec boil.Executor, insert bool, related *Output) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `output_claim` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"output_id"}),
		strmangle.WhereClause("`", "`", 0, outputClaimPrimaryKeyColumns),
	)
	values := []interface{}{related.TransactionHash, o.OutputID, o.VoutID, o.ClaimID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OutputID = related.TransactionHash

	if o.R == nil {
		o.R = &outputClaimR{
			Output: related,
		}
	} else {
		o.R.Output = related
	}

	if related.R == nil {
		related.R = &outputR{
			OutputClaims: OutputClaimSlice{o},
		}
	} else {
		related.R.OutputClaims = append(related.R.OutputClaims, o)
	}

	return nil
}

// SetVoutG of the output_claim to the related item.
// Sets o.R.Vout to related.
// Adds o to related.R.VoutOutputClaims.
// Uses the global database handle.
func (o *OutputClaim) SetVoutG(insert bool, related *Output) error {
	return o.SetVout(boil.GetDB(), insert, related)
}

// SetVoutP of the output_claim to the related item.
// Sets o.R.Vout to related.
// Adds o to related.R.VoutOutputClaims.
// Panics on error.
func (o *OutputClaim) SetVoutP(exec boil.Executor, insert bool, related *Output) {
	if err := o.SetVout(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetVoutGP of the output_claim to the related item.
// Sets o.R.Vout to related.
// Adds o to related.R.VoutOutputClaims.
// Uses the global database handle and panics on error.
func (o *OutputClaim) SetVoutGP(insert bool, related *Output) {
	if err := o.SetVout(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetVout of the output_claim to the related item.
// Sets o.R.Vout to related.
// Adds o to related.R.VoutOutputClaims.
func (o *OutputClaim) SetVout(exec boil.Executor, insert bool, related *Output) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `output_claim` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"vout_id"}),
		strmangle.WhereClause("`", "`", 0, outputClaimPrimaryKeyColumns),
	)
	values := []interface{}{related.Vout, o.OutputID, o.VoutID, o.ClaimID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.VoutID = related.Vout

	if o.R == nil {
		o.R = &outputClaimR{
			Vout: related,
		}
	} else {
		o.R.Vout = related
	}

	if related.R == nil {
		related.R = &outputR{
			VoutOutputClaims: OutputClaimSlice{o},
		}
	} else {
		related.R.VoutOutputClaims = append(related.R.VoutOutputClaims, o)
	}

	return nil
}

// SetClaimG of the output_claim to the related item.
// Sets o.R.Claim to related.
// Adds o to related.R.OutputClaims.
// Uses the global database handle.
func (o *OutputClaim) SetClaimG(insert bool, related *Claim) error {
	return o.SetClaim(boil.GetDB(), insert, related)
}

// SetClaimP of the output_claim to the related item.
// Sets o.R.Claim to related.
// Adds o to related.R.OutputClaims.
// Panics on error.
func (o *OutputClaim) SetClaimP(exec boil.Executor, insert bool, related *Claim) {
	if err := o.SetClaim(exec, insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetClaimGP of the output_claim to the related item.
// Sets o.R.Claim to related.
// Adds o to related.R.OutputClaims.
// Uses the global database handle and panics on error.
func (o *OutputClaim) SetClaimGP(insert bool, related *Claim) {
	if err := o.SetClaim(boil.GetDB(), insert, related); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetClaim of the output_claim to the related item.
// Sets o.R.Claim to related.
// Adds o to related.R.OutputClaims.
func (o *OutputClaim) SetClaim(exec boil.Executor, insert bool, related *Claim) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `output_claim` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"claim_id"}),
		strmangle.WhereClause("`", "`", 0, outputClaimPrimaryKeyColumns),
	)
	values := []interface{}{related.ClaimID, o.OutputID, o.VoutID, o.ClaimID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ClaimID = related.ClaimID

	if o.R == nil {
		o.R = &outputClaimR{
			Claim: related,
		}
	} else {
		o.R.Claim = related
	}

	if related.R == nil {
		related.R = &claimR{
			OutputClaims: OutputClaimSlice{o},
		}
	} else {
		related.R.OutputClaims = append(related.R.OutputClaims, o)
	}

	return nil
}

// OutputClaimsG retrieves all records.
func OutputClaimsG(mods ...qm.QueryMod) outputClaimQuery {
	return OutputClaims(boil.GetDB(), mods...)
}

// OutputClaims retrieves all the records using an executor.
func OutputClaims(exec boil.Executor, mods ...qm.QueryMod) outputClaimQuery {
	mods = append(mods, qm.From("`output_claim`"))
	return outputClaimQuery{NewQuery(exec, mods...)}
}

// FindOutputClaimG retrieves a single record by ID.
func FindOutputClaimG(outputID string, voutID uint, claimID string, selectCols ...string) (*OutputClaim, error) {
	return FindOutputClaim(boil.GetDB(), outputID, voutID, claimID, selectCols...)
}

// FindOutputClaimGP retrieves a single record by ID, and panics on error.
func FindOutputClaimGP(outputID string, voutID uint, claimID string, selectCols ...string) *OutputClaim {
	retobj, err := FindOutputClaim(boil.GetDB(), outputID, voutID, claimID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOutputClaim retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOutputClaim(exec boil.Executor, outputID string, voutID uint, claimID string, selectCols ...string) (*OutputClaim, error) {
	outputClaimObj := &OutputClaim{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `output_claim` where `output_id`=? AND `vout_id`=? AND `claim_id`=?", sel,
	)

	q := queries.Raw(exec, query, outputID, voutID, claimID)

	err := q.Bind(outputClaimObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from output_claim")
	}

	return outputClaimObj, nil
}

// FindOutputClaimP retrieves a single record by ID with an executor, and panics on error.
func FindOutputClaimP(exec boil.Executor, outputID string, voutID uint, claimID string, selectCols ...string) *OutputClaim {
	retobj, err := FindOutputClaim(exec, outputID, voutID, claimID, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OutputClaim) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *OutputClaim) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *OutputClaim) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *OutputClaim) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("model: no output_claim provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(outputClaimColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	outputClaimInsertCacheMut.RLock()
	cache, cached := outputClaimInsertCache[key]
	outputClaimInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			outputClaimColumns,
			outputClaimColumnsWithDefault,
			outputClaimColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(outputClaimType, outputClaimMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(outputClaimType, outputClaimMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `output_claim` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `output_claim` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `output_claim` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, outputClaimPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.Exec(cache.query, vals...)
	if err != nil {
		return errors.Wrap(err, "model: unable to insert into output_claim")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.OutputID,
		o.VoutID,
		o.ClaimID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for output_claim")
	}

CacheNoHooks:
	if !cached {
		outputClaimInsertCacheMut.Lock()
		outputClaimInsertCache[key] = cache
		outputClaimInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single OutputClaim record. See Update for
// whitelist behavior description.
func (o *OutputClaim) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single OutputClaim record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *OutputClaim) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the OutputClaim, and panics on error.
// See Update for whitelist behavior description.
func (o *OutputClaim) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the OutputClaim.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *OutputClaim) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	key := makeCacheKey(whitelist, nil)
	outputClaimUpdateCacheMut.RLock()
	cache, cached := outputClaimUpdateCache[key]
	outputClaimUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			outputClaimColumns,
			outputClaimPrimaryKeyColumns,
			whitelist,
		)

		if len(wl) == 0 {
			return errors.New("model: unable to update output_claim, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `output_claim` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, outputClaimPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(outputClaimType, outputClaimMapping, append(wl, outputClaimPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update output_claim row")
	}

	if !cached {
		outputClaimUpdateCacheMut.Lock()
		outputClaimUpdateCache[key] = cache
		outputClaimUpdateCacheMut.Unlock()
	}

	return nil
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q outputClaimQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q outputClaimQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "model: unable to update all for output_claim")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OutputClaimSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OutputClaimSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OutputClaimSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OutputClaimSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outputClaimPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `output_claim` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, outputClaimPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to update all in outputClaim slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OutputClaim) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *OutputClaim) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *OutputClaim) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *OutputClaim) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("model: no output_claim provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(outputClaimColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	outputClaimUpsertCacheMut.RLock()
	cache, cached := outputClaimUpsertCache[key]
	outputClaimUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			outputClaimColumns,
			outputClaimColumnsWithDefault,
			outputClaimColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			outputClaimColumns,
			outputClaimPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("model: unable to upsert output_claim, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "output_claim", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `output_claim` WHERE `output_id`=? AND `vout_id`=? AND `claim_id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(outputClaimType, outputClaimMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(outputClaimType, outputClaimMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.Exec(cache.query, vals...)
	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for output_claim")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.OutputID,
		o.VoutID,
		o.ClaimID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for output_claim")
	}

CacheNoHooks:
	if !cached {
		outputClaimUpsertCacheMut.Lock()
		outputClaimUpsertCache[key] = cache
		outputClaimUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteP deletes a single OutputClaim record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OutputClaim) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single OutputClaim record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OutputClaim) DeleteG() error {
	if o == nil {
		return errors.New("model: no OutputClaim provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single OutputClaim record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OutputClaim) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single OutputClaim record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OutputClaim) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no OutputClaim provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), outputClaimPrimaryKeyMapping)
	sql := "DELETE FROM `output_claim` WHERE `output_id`=? AND `vout_id`=? AND `claim_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete from output_claim")
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q outputClaimQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q outputClaimQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("model: no outputClaimQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from output_claim")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OutputClaimSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OutputClaimSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("model: no OutputClaim slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OutputClaimSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OutputClaimSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("model: no OutputClaim slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outputClaimPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `output_claim` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, outputClaimPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "model: unable to delete all from outputClaim slice")
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *OutputClaim) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *OutputClaim) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OutputClaim) ReloadG() error {
	if o == nil {
		return errors.New("model: no OutputClaim provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OutputClaim) Reload(exec boil.Executor) error {
	ret, err := FindOutputClaim(exec, o.OutputID, o.VoutID, o.ClaimID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OutputClaimSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OutputClaimSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OutputClaimSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("model: empty OutputClaimSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OutputClaimSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	outputClaims := OutputClaimSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), outputClaimPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `output_claim`.* FROM `output_claim` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, outputClaimPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&outputClaims)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in OutputClaimSlice")
	}

	*o = outputClaims

	return nil
}

// OutputClaimExists checks if the OutputClaim row exists.
func OutputClaimExists(exec boil.Executor, outputID string, voutID uint, claimID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `output_claim` where `output_id`=? AND `vout_id`=? AND `claim_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, outputID, voutID, claimID)
	}

	row := exec.QueryRow(sql, outputID, voutID, claimID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if output_claim exists")
	}

	return exists, nil
}

// OutputClaimExistsG checks if the OutputClaim row exists.
func OutputClaimExistsG(outputID string, voutID uint, claimID string) (bool, error) {
	return OutputClaimExists(boil.GetDB(), outputID, voutID, claimID)
}

// OutputClaimExistsGP checks if the OutputClaim row exists. Panics on error.
func OutputClaimExistsGP(outputID string, voutID uint, claimID string) bool {
	e, err := OutputClaimExists(boil.GetDB(), outputID, voutID, claimID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OutputClaimExistsP checks if the OutputClaim row exists. Panics on error.
func OutputClaimExistsP(exec boil.Executor, outputID string, voutID uint, claimID string) bool {
	e, err := OutputClaimExists(exec, outputID, voutID, claimID)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
