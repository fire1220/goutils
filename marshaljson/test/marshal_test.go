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
	ID          int32       `json:"id" default:"456"`
	ValFloat    float64     `json:"val_float" default:"111"`
	Val         uint        `json:"val" default:"-111"`
	ValBool     bool        `json:"val_bool" default:"true"`
	ValSlice    []bool      `json:"val_slice" default:"[]"`
	ValSlice2   []bool      `json:"val_slice2" default:"{}"`
	ValMap      map[int]int `json:"val_map" default:"[]"`
	ValStruct   GoodInfo    `json:"val_struct"`
	ValStruct2  GoodInfo    `json:"val_struct2" default:"{}"`
	ValStruct3  GoodInfo    `json:"val_struct3" default:"hello"`
	ValPtr      *GoodInfo   `json:"val_ptr" default:"{}"`
	ValPtr2     *GoodInfo   `json:"val_ptr2" default:"[]"`
	ValPtr3     *GoodInfo   `json:"val_ptr3" default:"hello"`
	Name        string      `json:"name" default:"123"`
	PlayTime    time.Time   `json:"play_time" datetime:"2006-01-02 15:04:05"`
	ExecuteTime time.Time   `json:"execute_time" datetime:"2006-01-02" default:"-"`
	CreatedAt   time.Time   `json:"created_at" datetime:"2006-01-02 15:04:05" default:"0000-00-00"`
	UpdatedAt   time.Time   `json:"updated_at" default:""`
}

func (t Good) MarshalJSON() ([]byte, error) {
	return marshaljson.MarshalFormat(t)
}

func TestMarshal(t *testing.T) {
	good := Good{ID: 0, Name: "", PlayTime: time.Now(), ExecuteTime: time.Now()}
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
