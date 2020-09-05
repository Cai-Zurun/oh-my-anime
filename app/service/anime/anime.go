package anime

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"io/ioutil"
	"oh-my-anime_gf/app/model/anime"
	"github.com/gogf/gf/net/ghttp"
	"os"
)

type AddAnimeInput struct {
	Name string	`v:"required#名字不能为空"`
	Link string	`v:"required#链接不能为空"`
	Type string	`v:"required#类型不能为空"`
	Img  *ghttp.UploadFile
}

const ImgPath = "/Users/zrun/Img/animeImg/"		//部署到服务器时自定义修改

func AddAnime(data *AddAnimeInput) error{
	// 检查添加的类型是否已经存在
	if CheckAnime(data.Name, data.Type) {
		return errors.New(fmt.Sprintf("在%s类型下的 %s已经存在", data.Type,data.Name))
	}
	//添加的anime的type不存在，则添加该type
	if !CheckType(data.Type) {
		AddType(data.Type)
	}
	var entity *anime.Entity
	if err := gconv.Struct(data, &entity); err != nil {
		return err
	}
	entity.CreateTime = gtime.Now()
	if err := SaveFile(data.Img, entity); err !=nil {
		return err
	}
	if _, err := anime.Save(entity); err != nil {
		return err
	}
	return nil
}

//Anime存在返回true
func CheckAnime(Name string, Type string) bool {
	cnt, err := anime.FindCount("Name=? and Type=?", Name, Type)
	if err != nil {
		return false
	} else {
		return cnt != 0
	}
}

func SaveFile(File *ghttp.UploadFile, entity *anime.Entity)  error{
	name, err := File.Save(ImgPath)
	if err != nil {
		return err
	}
	entity.ImgPath = ImgPath + name
	return nil
}

//Where不知道该咋用
//func CheckAnime(data *Info) bool {
//	res := anime.Model.Where("Name", data.Name).And("Type", data.Type)
//	fmt.Println(res)
//	if res != nil {
//		return false
//	}
//	return true
//}

type GetAnimeInput struct {
	Type string `v:"required#类型不能为空"`
}

type AnimeOutput struct {
	id int
	Name string
	Link string
	Type string
	CreateTime *gtime.Time
	Img string
}

func GetAnime(Type string)  ([]AnimeOutput, error){
	if !CheckType(Type) {
		return nil, errors.New(fmt.Sprintf("%s类型不存在", Type))
	}
	AnimeArr, err := anime.FindAll("Type", Type)
	Animes := make([]AnimeOutput, len(AnimeArr))
	for i := 0; i <len(AnimeArr) ; i++ {
		if AnimeArr[i].ImgPath != "" {
			fmt.Println(AnimeArr[i].ImgPath)
			file, _ := os.Open(AnimeArr[i].ImgPath)
			buff, _ := ioutil.ReadAll(file)
			imgEnc := base64.StdEncoding.EncodeToString(buff)
			Animes[i].Img = imgEnc
			//TODO::搞懂原理
		}
		Animes[i].id = AnimeArr[i].Id
		Animes[i].Name = AnimeArr[i].Name
		Animes[i].Link = AnimeArr[i].Link
		Animes[i].Type = AnimeArr[i].Type
		Animes[i].CreateTime = AnimeArr[i].CreateTime
	}
	return Animes, err
}

type DeleteAnimeInput struct {
	Type string `v:"required#类型不能为空"`
	Name string `v:"required#名字不能为空"`
}

func DeleteAnime(data *DeleteAnimeInput) error {
	if !CheckType(data.Type) {
		return errors.New(fmt.Sprintf("%s 类型不存在", data.Type))
	}
	if !CheckAnime(data.Name, data.Type) {
		return errors.New(fmt.Sprintf("%s 动漫不存在", data.Name))
	}
	_, err := anime.Delete("Type=? and Name=?", data.Type, data.Name)
	if err != nil {
		return err
	}
	return nil
}

type UpdateAnimeInput struct {
	Type string `v:"required#类型不能为空"`
	Name string `v:"required#名字不能为空"`
	NewName string
	NewLink string
}

func UpdateAnime(data *UpdateAnimeInput)  error{
	if !CheckType(data.Type) {
		return errors.New(fmt.Sprintf("%s 类型不存在", data.Type))
	}
	if !CheckAnime(data.Name, data.Type) {
		return errors.New(fmt.Sprintf("%s 动漫不存在", data.Name))
	}
	if data.NewName != "" && data.NewLink != "" {
		if _, err := anime.Model.Data("Name =? , Link=?", data.NewName, data.NewLink).Where("Type", data.Type).And("Name", data.Name).Update(); err != nil {
			return err
		}
	}else if data.NewName != "" {
		if _, err := anime.Model.Data("Name", data.NewName).Where("Type", data.Type).And("Name", data.Name).Update(); err != nil {
			return err
		}
	}else if data.NewLink != "" {
		if _, err := anime.Model.Data("Link", data.NewLink).Where("Type", data.Type).And("Name", data.Name).Update(); err != nil {
			return err
		}
	}
	return nil
}