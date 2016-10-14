package db

import (
	"testing"
	"time"
	"os"
	"path"
	"io/ioutil"
	"strings"
	"os/exec"
	"path/filepath"
	"gopkg.in/baa.v1"
	"../entity"
)


//初始化配置文件
func initConfig() {
	currPath, _ := os.Getwd()
	file, _ := os.Open(path.Join(currPath, "../conf/app.ini"))
	defer file.Close()
	resu, _ := ioutil.ReadAll(file)
	strResu := string(resu);

	splits := strings.Split(strResu, "\n")

	ret := make(map[string]string)
	for _, v := range splits {

		if strings.Trim(v, " ") == "" || v[0] == '#' {
			continue
		}
		keyval := strings.Split(v, "=")

		ret[strings.Trim(keyval[0], " ")] = strings.Trim(strings.Trim(keyval[1], " "), "\"")
	}

	b := baa.Default()
	b.SetDI("config", ret)
}



//[UserInfoDAL.Login]
func Test_userLogin(t *testing.T) {
	user, err := UserInfoDAL.Login("hhx", "", "123")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("查询完成", user)
	}
}

//[ArticlesDAL.GetListByCategoryId]
func Test_GetArticleListByCid(t *testing.T) {
	ret, err := ArticlesDAL.GetListByCategoryId(0, 0, 2, 0)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(len(ret), ret)
	}
}


//[ArticlesDAL.GetDetail]
func Test_getdetail(t *testing.T) {
	initConfig()
	ret, err := ArticlesDAL.GetDetail(2)
	t.Log(err)
	t.Log(ret)
}


//[UserInfoDAL.GetById]测试获取用户信息
func Test_getuser(t *testing.T) {
	t.Log(UserInfoDAL.GetById(1))
	t.Log(UserInfoDAL.GetById(1111))
}

//[UserInfoDAL.regist]测试注册用户
func Test_regist(t *testing.T) {
	initConfig()
	user := &entity.UserInfo{}
	user.EmailAddr = "22222"
	user.LoginName = "11122"
	user.LoginPwd = "333"
	user.NickName = "hhasd"

	UserInfoDAL.Regist(user)
	t.Log(user)
}

//[UserInfoDAL.regist]测试注册用户
func Test_checkEmailAndLoginname(t *testing.T) {
	initConfig()
	user := entity.UserInfo{LoginName:"hhxa", EmailAddr:"222a"}
	ret := UserInfoDAL.EmailOrLoginNameExist(user)
	t.Log(ret)

}

//[CategoriesDAL.GetListName]测试获取用户信息
func Test_getcategoryListName(t *testing.T) {
	initConfig()
	t.Log(CategoriesDAL.GetListName())

	//Test_testpath(t)
}

func Test_testpath(t *testing.T) {
	execDirAbsPath, _ := os.Getwd()
	t.Log("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)

	execFileRelativePath, _ := exec.LookPath(os.Args[0])
	t.Log("执行程序与命令执行目录的相对路径　　　　:", execFileRelativePath)

	execDirRelativePath, _ := path.Split(execFileRelativePath)
	t.Log("执行程序所在目录与命令执行目录的相对路径:", execDirRelativePath)

	execFileAbsPath, _ := filepath.Abs(execFileRelativePath)
	t.Log("执行程序的绝对路径　　　　　　　　　　　:", execFileAbsPath)

	execDirAbsPath, _ = filepath.Abs(execDirRelativePath)
	t.Log("执行程序所在目录的绝对路径　　　　　　　:", execDirAbsPath)

	os.Chdir(execDirRelativePath) //进入目录
	enteredDirAbsPath, _ := os.Getwd()
	t.Log("所进入目录的绝对路径　　　　　　　　　　:", enteredDirAbsPath)
}

//时间戳测试
func Test_timetamp(t *testing.T) {
	now := time.Now()
	_, offset := now.Zone()
	timetamp := time.Now().Unix() - int64(offset)

	t.Log(timetamp)
}

//时间戳测试
func Test_getCname(t *testing.T) {
	cname := CategoriesDAL.GetNameById(1)

	t.Log(cname)
}

func Test_io(t *testing.T) {

	currPath, _ := os.Getwd()
	file, _ := os.Open(path.Join(currPath, "../conf/app.ini"))
	defer file.Close()
	resu, _ := ioutil.ReadAll(file)
	strResu := string(resu);

	splits := strings.Split(strResu, "\n")

	ret := make(map[string]string)
	for _, v := range splits {
		if strings.Trim(v, " ") == "" {
			continue
		}
		keyval := strings.Split(v, "=")

		ret[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], " ")
	}

	t.Log(ret["mysqldb"])
}

func Test_timeadd(t *testing.T) {
	currTime := time.Now()

	t.Log(currTime)
	t.Log(currTime.Add(1e9))
}