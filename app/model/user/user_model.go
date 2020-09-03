// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package user

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"
)

// arModel is a active record design model for table user operations.
type arModel struct {
	gmvc.M
}

var (
	// Table is the table name of user.
	Table = "user"
	// Model is the model object of user.
	Model = &arModel{g.DB("default").Table(Table).Safe()}
	// Columns defines and stores column names for table user.
	Columns = struct {
		Id          string // 用户ID             
        Passport    string // 账号               
        Password    string // 密码               
        Nickname    string // 昵称               
        CreateTime  string // 创建时间/注册时间
	}{
		Id:         "id",           
        Passport:   "passport",     
        Password:   "password",     
        Nickname:   "nickname",     
        CreateTime: "create_time",
	}
)

// FindOne is a convenience method for Model.FindOne.
// See Model.FindOne.
func FindOne(where ...interface{}) (*Entity, error) {
	return Model.FindOne(where...)
}

// FindAll is a convenience method for Model.FindAll.
// See Model.FindAll.
func FindAll(where ...interface{}) ([]*Entity, error) {
	return Model.FindAll(where...)
}

// FindValue is a convenience method for Model.FindValue.
// See Model.FindValue.
func FindValue(fieldsAndWhere ...interface{}) (gdb.Value, error) {
	return Model.FindValue(fieldsAndWhere...)
}

// FindArray is a convenience method for Model.FindArray.
// See Model.FindArray.
func FindArray(fieldsAndWhere ...interface{}) ([]gdb.Value, error) {
	return Model.FindArray(fieldsAndWhere...)
}

// FindCount is a convenience method for Model.FindCount.
// See Model.FindCount.
func FindCount(where ...interface{}) (int, error) {
	return Model.FindCount(where...)
}

// Insert is a convenience method for Model.Insert.
func Insert(data ...interface{}) (result sql.Result, err error) {
	return Model.Insert(data...)
}

// InsertIgnore is a convenience method for Model.InsertIgnore.
func InsertIgnore(data ...interface{}) (result sql.Result, err error) {
	return Model.InsertIgnore(data...)
}

// Replace is a convenience method for Model.Replace.
func Replace(data ...interface{}) (result sql.Result, err error) {
	return Model.Replace(data...)
}

// Save is a convenience method for Model.Save.
func Save(data ...interface{}) (result sql.Result, err error) {
	return Model.Save(data...)
}

// Update is a convenience method for Model.Update.
func Update(dataAndWhere ...interface{}) (result sql.Result, err error) {
	return Model.Update(dataAndWhere...)
}

// Delete is a convenience method for Model.Delete.
func Delete(where ...interface{}) (result sql.Result, err error) {
	return Model.Delete(where...)
}

// As sets an alias name for current table.
func (m *arModel) As(as string) *arModel {
	return &arModel{m.M.As(as)}
}

// TX sets the transaction for current operation.
func (m *arModel) TX(tx *gdb.TX) *arModel {
	return &arModel{m.M.TX(tx)}
}

// Master marks the following operation on master node.
func (m *arModel) Master() *arModel {
	return &arModel{m.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (m *arModel) Slave() *arModel {
	return &arModel{m.M.Slave()}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (m *arModel) LeftJoin(table ...string) *arModel {
	return &arModel{m.M.LeftJoin(table ...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (m *arModel) RightJoin(table ...string) *arModel {
	return &arModel{m.M.RightJoin(table ...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (m *arModel) InnerJoin(table ...string) *arModel {
	return &arModel{m.M.InnerJoin(table ...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
func (m *arModel) Fields(fields string) *arModel {
	return &arModel{m.M.Fields(fields)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
func (m *arModel) FieldsEx(fields string) *arModel {
	return &arModel{m.M.FieldsEx(fields)}
}

// Option sets the extra operation option for the model.
func (m *arModel) Option(option int) *arModel {
	return &arModel{m.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (m *arModel) OmitEmpty() *arModel {
	return &arModel{m.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (m *arModel) Filter() *arModel {
	return &arModel{m.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (m *arModel) Where(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.Where(where, args...)}
}

// And adds "AND" condition to the where statement.
func (m *arModel) And(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (m *arModel) Or(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (m *arModel) Group(groupBy string) *arModel {
	return &arModel{m.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (m *arModel) Order(orderBy ...string) *arModel {
	return &arModel{m.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (m *arModel) Limit(limit ...int) *arModel {
	return &arModel{m.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (m *arModel) Offset(offset int) *arModel {
	return &arModel{m.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (m *arModel) Page(page, limit int) *arModel {
	return &arModel{m.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (m *arModel) Batch(batch int) *arModel {
	return &arModel{m.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (m *arModel) Cache(duration time.Duration, name ...string) *arModel {
	return &arModel{m.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (m *arModel) Data(data ...interface{}) *arModel {
	return &arModel{m.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*Entity.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
func (m *arModel) All(where ...interface{}) ([]*Entity, error) {
	all, err := m.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*Entity
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *Entity.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
func (m *arModel) One(where ...interface{}) (*Entity, error) {
	one, err := m.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *Entity
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by Model.WherePri and Model.One.
// Also see Model.WherePri and Model.One.
func (m *arModel) FindOne(where ...interface{}) (*Entity, error) {
	one, err := m.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *Entity
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by Model.WherePri and Model.All.
// Also see Model.WherePri and Model.All.
func (m *arModel) FindAll(where ...interface{}) ([]*Entity, error) {
	all, err := m.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*Entity
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Chunk iterates the table with given size and callback function.
func (m *arModel) Chunk(limit int, callback func(entities []*Entity, err error) bool) {
	m.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*Entity
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (m *arModel) LockUpdate() *arModel {
	return &arModel{m.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (m *arModel) LockShared() *arModel {
	return &arModel{m.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (m *arModel) Unscoped() *arModel {
	return &arModel{m.M.Unscoped()}
}