package anime

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"oh-my-anime_gf/app/model/anime"
	"oh-my-anime_gf/app/model/anime_type"
)

type AddTypeInput struct {
	Type string `v:"required#类型不能为空"`
}

func AddType(Type string) error{
	// 检查添加的类型是否已经存在
	if CheckType(Type) {
		return errors.New(fmt.Sprintf("类型 %s 已经存在", Type))
	}
	//使用g.Map是为了下面转换成struct
	data := g.Map{
		"Type" : Type,
	}
	var entity *anime_type.Entity
	if err := gconv.Struct(data, &entity); err != nil {
		return err
	}
	if _, err := anime_type.Save(entity); err != nil {
		return err
	}
	return nil
}

//Type存在则返回true
func CheckType(Type string) bool {
	cnt, err := anime_type.FindCount("Type", Type)
	if err != nil {
		return false
	} else {
		return cnt != 0
	}
}

func GetType()  []gdb.Value{
	TypeArr, _ := anime_type.FindArray("type")
	return 	TypeArr
}

type DeleteTypeInput struct {
	Type string `v:"required#类型不能为空"`
}

func DeleteType(data *DeleteTypeInput)  error{
	if !CheckType(data.Type) {
		return errors.New(fmt.Sprintf("%s 类型不存在", data.Type))
	}
	if _, err := anime_type.Delete("Type", data.Type); err !=nil {
		return err
	}
	if _, err := anime.Delete("Type", data.Type); err !=nil {
		return err
	}
	return nil
}

type UpdateTypeInput struct {
	Type string `v:"required#类型不能为空"`
	NewType string `v:"required#新类型不能为空"`
}

func UpdateType(data *UpdateTypeInput)  error{
	if !CheckType(data.Type) {
		return errors.New(fmt.Sprintf("%s 类型不存在", data.Type))
	}
	if _, err := anime_type.Model.Data("Type", data.NewType).Where("Type", data.Type).Update(); err != nil {
		return err
	}
	if _, err := anime.Model.Data("Type", data.NewType).Where("Type", data.Type).Update(); err != nil {
		return err
	}
	return nil
}

type UpdateSeqInput struct {
	Order string `v:"required#指令不能为空"`
	Type string `v:"required#类型不能为空"`
}

func UpdateSeq(data *UpdateSeqInput) error {
	if data.Order != "up" && data.Order != "down" {
		return errors.New("order必须为up或down")
	}else if data.Order == "up" {

	}else {

	}
}
