package test

import (
	"encoding/json"
	"fmt"
	"github.com/fire1220/goutils/marshaljson"
	"testing"
	"time"
)

type GoodInfo struct {
	Title    string    `json:"title" default:"ABC"`
	Like     string    `json:"like"`
	PlayTime time.Time `json:"play_time" datetime:"2006-01-02 15:04:05"`
}

func (t GoodInfo) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(t)
}

type Good struct {
	ID      int32       `json:"id" default:"456"`
	Float1  float64     `json:"float1" default:"11.1"`
	Float2  float64     `json:"float2" default:"-11.1"`
	Float3  float64     `json:"float3" defaultString:"hello"`
	Uint    uint        `json:"uint" default:"-111"`
	Bool1   bool        `json:"bool1" default:"false"`
	Bool2   bool        `json:"bool2" default:"true"`
	Bool3   bool        `json:"bool3" defaultString:"hello"`
	Slice1  []bool      `json:"slice1" default:"[]"`
	Slice2  []bool      `json:"slice2" default:"{}"`
	Slice3  []bool      `json:"slice3" defaultString:"hello"`
	Map1    map[int]int `json:"map1" default:"[]"`
	Map2    map[int]int `json:"map2" default:"hello"`
	Struct1 GoodInfo    `json:"struct1"`
	Struct2 GoodInfo    `json:"struct2" default:"{}"`
	Struct3 GoodInfo    `json:"struct3" default:"hello"`
	Ptr1    *GoodInfo   `json:"ptr1" default:"{}"`
	Ptr2    *GoodInfo   `json:"ptr2" default:"[]"`
	Ptr3    *GoodInfo   `json:"ptr3" default:"hello"`
	String1 string      `json:"string1" defaultString:"hello"`
	String2 string      `json:"string2" default:"hello"`
	Name    string      `json:"name" default:"123"`
	Time1   time.Time   `json:"time1" datetime:"2006-01-02 15:04:05"`
	Time2   time.Time   `json:"time2" datetime:"2006-01-02"`
	Time3   time.Time   `json:"time3" datetime:"15:04:05" default:"-"`
	Time4   time.Time   `json:"time4" datetime:"2006-01-02 15:04:05" default:"0000-00-00"`
	Time5   time.Time   `json:"time5" default:""`
	Time6   time.Time   `json:"time6" default:"-"`
}

func (t Good) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(t)
}

func TestMarshal(t *testing.T) {
	good := Good{
		ID:    0,
		Name:  "",
		Time1: time.Now(),
		Time2: time.Now(),
		Time3: time.Now(),
		Time4: time.Now(),
	}
	bytes, err := json.Marshal(good)
	fmt.Printf("%s\n", bytes)
	fmt.Println(err)
	goodList := make([]Good, 0)
	goodList = append(goodList, good)
	bytes, err = json.Marshal(goodList)
	fmt.Printf("%s\n", bytes)
	fmt.Println(err)
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	TeacherId int       `json:"teacher_id"`
	CreatedAt time.Time `json:"created_at" datetime:"2006-01-02 15:04:05" default:"0000-00-00"`
}

// 如果有结构体嵌套的情况，需要把每个结构图都实现 MarshalJSON 方法，
// 否则会把子集的结构体的 MarshalJSON继承到父级里，导致结构图替换时候缺少父级字段
// 或者不用嵌套的方式，用 User User `json:"user"` 的方式
type UserWithTeacher struct {
	// User        User      `json:"user"`
	User
	TeacherName string    `json:"teacher_name"`
	PlayTime    time.Time `json:"play_time" datetime:"2006-01-02 15:04:05"`
}

func (t User) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(t)
}

func (t UserWithTeacher) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(t)
}

func TestMarshal2(t *testing.T) {
	user := User{
		ID:        1,
		Name:      "李四",
		TeacherId: 1,
		CreatedAt: time.Now(),
	}
	teacher := UserWithTeacher{
		User:        user,
		TeacherName: "张三",
		PlayTime:    time.Now(),
	}
	bytes, err := json.Marshal(teacher)
	fmt.Printf("%s\n", bytes)
	fmt.Println(err)
}
