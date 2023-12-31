// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Friendship is an object representing the database table.
type Friendship struct {
	FriendshipID int    `boil:"friendship_id" json:"friendship_id" toml:"friendship_id" yaml:"friendship_id"`
	UserEmail1   string `boil:"user_email_1" json:"user_email_1" toml:"user_email_1" yaml:"user_email_1"`
	UserEmail2   string `boil:"user_email_2" json:"user_email_2" toml:"user_email_2" yaml:"user_email_2"`

	R *friendshipR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L friendshipL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FriendshipColumns = struct {
	FriendshipID string
	UserEmail1   string
	UserEmail2   string
}{
	FriendshipID: "friendship_id",
	UserEmail1:   "user_email_1",
	UserEmail2:   "user_email_2",
}

var FriendshipTableColumns = struct {
	FriendshipID string
	UserEmail1   string
	UserEmail2   string
}{
	FriendshipID: "friendships.friendship_id",
	UserEmail1:   "friendships.user_email_1",
	UserEmail2:   "friendships.user_email_2",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var FriendshipWhere = struct {
	FriendshipID whereHelperint
	UserEmail1   whereHelperstring
	UserEmail2   whereHelperstring
}{
	FriendshipID: whereHelperint{field: "\"friendships\".\"friendship_id\""},
	UserEmail1:   whereHelperstring{field: "\"friendships\".\"user_email_1\""},
	UserEmail2:   whereHelperstring{field: "\"friendships\".\"user_email_2\""},
}

// FriendshipRels is where relationship names are stored.
var FriendshipRels = struct {
}{}

// friendshipR is where relationships are stored.
type friendshipR struct {
}

// NewStruct creates a new relationship struct
func (*friendshipR) NewStruct() *friendshipR {
	return &friendshipR{}
}

// friendshipL is where Load methods for each relationship are stored.
type friendshipL struct{}

var (
	friendshipAllColumns            = []string{"friendship_id", "user_email_1", "user_email_2"}
	friendshipColumnsWithoutDefault = []string{"user_email_1", "user_email_2"}
	friendshipColumnsWithDefault    = []string{"friendship_id"}
	friendshipPrimaryKeyColumns     = []string{"friendship_id"}
	friendshipGeneratedColumns      = []string{}
)

type (
	// FriendshipSlice is an alias for a slice of pointers to Friendship.
	// This should almost always be used instead of []Friendship.
	FriendshipSlice []*Friendship

	friendshipQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	friendshipType                 = reflect.TypeOf(&Friendship{})
	friendshipMapping              = queries.MakeStructMapping(friendshipType)
	friendshipPrimaryKeyMapping, _ = queries.BindMapping(friendshipType, friendshipMapping, friendshipPrimaryKeyColumns)
	friendshipInsertCacheMut       sync.RWMutex
	friendshipInsertCache          = make(map[string]insertCache)
	friendshipUpdateCacheMut       sync.RWMutex
	friendshipUpdateCache          = make(map[string]updateCache)
	friendshipUpsertCacheMut       sync.RWMutex
	friendshipUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single friendship record from the query.
func (q friendshipQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Friendship, error) {
	o := &Friendship{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: failed to execute a one query for friendships")
	}

	return o, nil
}

// All returns all Friendship records from the query.
func (q friendshipQuery) All(ctx context.Context, exec boil.ContextExecutor) (FriendshipSlice, error) {
	var o []*Friendship

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "orm: failed to assign all query results to Friendship slice")
	}

	return o, nil
}

// Count returns the count of all Friendship records in the query.
func (q friendshipQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to count friendships rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q friendshipQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "orm: failed to check if friendships exists")
	}

	return count > 0, nil
}

// Friendships retrieves all the records using an executor.
func Friendships(mods ...qm.QueryMod) friendshipQuery {
	mods = append(mods, qm.From("\"friendships\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"friendships\".*"})
	}

	return friendshipQuery{q}
}

// FindFriendship retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFriendship(ctx context.Context, exec boil.ContextExecutor, friendshipID int, selectCols ...string) (*Friendship, error) {
	friendshipObj := &Friendship{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"friendships\" where \"friendship_id\"=$1", sel,
	)

	q := queries.Raw(query, friendshipID)

	err := q.Bind(ctx, exec, friendshipObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "orm: unable to select from friendships")
	}

	return friendshipObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Friendship) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no friendships provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(friendshipColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	friendshipInsertCacheMut.RLock()
	cache, cached := friendshipInsertCache[key]
	friendshipInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			friendshipAllColumns,
			friendshipColumnsWithDefault,
			friendshipColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(friendshipType, friendshipMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"friendships\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"friendships\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "orm: unable to insert into friendships")
	}

	if !cached {
		friendshipInsertCacheMut.Lock()
		friendshipInsertCache[key] = cache
		friendshipInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Friendship.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Friendship) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	friendshipUpdateCacheMut.RLock()
	cache, cached := friendshipUpdateCache[key]
	friendshipUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			friendshipAllColumns,
			friendshipPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("orm: unable to update friendships, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"friendships\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, friendshipPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, append(wl, friendshipPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update friendships row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by update for friendships")
	}

	if !cached {
		friendshipUpdateCacheMut.Lock()
		friendshipUpdateCache[key] = cache
		friendshipUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q friendshipQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all for friendships")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected for friendships")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FriendshipSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("orm: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"friendships\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, friendshipPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to update all in friendship slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to retrieve rows affected all in update all friendship")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Friendship) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("orm: no friendships provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(friendshipColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	friendshipUpsertCacheMut.RLock()
	cache, cached := friendshipUpsertCache[key]
	friendshipUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			friendshipAllColumns,
			friendshipColumnsWithDefault,
			friendshipColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			friendshipAllColumns,
			friendshipPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("orm: unable to upsert friendships, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(friendshipPrimaryKeyColumns))
			copy(conflict, friendshipPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"friendships\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(friendshipType, friendshipMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(friendshipType, friendshipMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "orm: unable to upsert friendships")
	}

	if !cached {
		friendshipUpsertCacheMut.Lock()
		friendshipUpsertCache[key] = cache
		friendshipUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Friendship record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Friendship) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("orm: no Friendship provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), friendshipPrimaryKeyMapping)
	sql := "DELETE FROM \"friendships\" WHERE \"friendship_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete from friendships")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by delete for friendships")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q friendshipQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("orm: no friendshipQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from friendships")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for friendships")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FriendshipSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"friendships\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, friendshipPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "orm: unable to delete all from friendship slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "orm: failed to get rows affected by deleteall for friendships")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Friendship) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFriendship(ctx, exec, o.FriendshipID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FriendshipSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FriendshipSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), friendshipPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"friendships\".* FROM \"friendships\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, friendshipPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "orm: unable to reload all in FriendshipSlice")
	}

	*o = slice

	return nil
}

// FriendshipExists checks if the Friendship row exists.
func FriendshipExists(ctx context.Context, exec boil.ContextExecutor, friendshipID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"friendships\" where \"friendship_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, friendshipID)
	}
	row := exec.QueryRowContext(ctx, sql, friendshipID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "orm: unable to check if friendships exists")
	}

	return exists, nil
}

// Exists checks if the Friendship row exists.
func (o *Friendship) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FriendshipExists(ctx, exec, o.FriendshipID)
}
