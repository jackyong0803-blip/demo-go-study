package admin

import (
	"demo-go-study/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {

	//查询数据库
	userList := []models.User{}

	// models.DB.Find(&userList)

	// c.JSON(200, gin.H{
	// 	"result": userList,
	// })

	//查询age大于20的用户
	//userList := []models.User{}
	//models.DB.Where("age<20").Find(&userList)
	// models.DB.Find(&userList)

	// 查询id大于3小于8的数据
	// var start = 3
	// var end = 8
	// models.DB.Where("id>? AND id<?", start, end).Find(&userList)

	//使用in查询  id在3,5,6中的数据
	//models.DB.Where("id in (?)", []int{3, 5, 6}).Find(&userList)

	//使用like 模糊查询
	//models.DB.Where("username like (?)", "%浩%").Find(&userList)

	// Or 查询
	//models.DB.Where("id=? OR id=?", 2, 3).Find(&userList)
	//models.DB.Where("id=?", 2).Or("id=?", 3).Or("id=?", 4).Find(&userList)

	//使用 select 返回指定字段
	//models.DB.Select("id,username,email").Find(&userList)
	//只返回指定的字段
	// type UserInfo struct {
	// 	Id       int    `json:"id"`
	// 	Username string `json:"username"`
	// }

	// var userList []UserInfo
	// models.DB.Model(&models.User{}).Select("id,username").Find(&userList)

	// c.JSON(200, gin.H{
	// 	"result": userList,
	// })

	//排序   order("id desc")
	models.DB.Order("age asc").Find(&userList)
	c.JSON(200, gin.H{
		"result": userList,
	})

}
func (con UserController) Add(c *gin.Context) {

	user := models.User{
		Username: "王五",
		Age:      32,
		Email:    "weang@qq.con",
		AddTime:  int(models.GetUnix()),
	}

	err := models.DB.Create(&user).Error
	if err != nil {
		fmt.Println("创建用户失败:", err)
		c.String(500, "增加数据失败: "+err.Error())
		return
	}
	fmt.Println("创建的用户:", user)
	c.String(200, "增加数据成功，用户ID: %d", user.Id)
}
func (con UserController) Edit(c *gin.Context) {
	//保存所有字段

	// // 查询id等于6的数据
	// user := models.User{Id: 6}
	// models.DB.Find(&user)
	// //更新数据
	// user.Username = "哈哈"
	// user.Email = "itying@qqq.com"
	// user.AddTime = int(models.GetUnix())
	// models.DB.Save(&user)

	//更新单个列
	// user := models.User{}
	// models.DB.Model(&user).Where("id = ?", 6).Update("username", "哈哈哈哈哈哈")

	user := models.User{}
	models.DB.Where("id = ?", 6).Find(&user)
	user.Username = "哈哈"
	user.Email = "aaa@qqq.com"
	user.AddTime = int(models.GetUnix())
	err := models.DB.Save(&user).Error
	if err != nil {
		fmt.Println("修改用户失败:", err)
		c.String(500, "修改用户失败: "+err.Error())
		return
	}
	fmt.Println("修改用户成功:", user)
	c.String(200, "修改用户成功")
}
func (con UserController) Delete(c *gin.Context) {

	// user := models.User{Id: 6}
	// models.DB.Delete(&user)

	//删除数据
	user := models.User{}
	models.DB.Where("username = ?", "gorm").Delete(&user)

	c.String(200, "删除用户")
}
