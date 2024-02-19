package admin

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/models"
	"github.com/patih1/fwg17-go-backend/src/services"
)

func ListAllPromo(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit
	result, err := models.FindAllPromo(limit, offset)

	pageInfo := services.PageInfo{
		Page:      page,
		Limit:     limit,
		LastPage:  int(math.Ceil(float64(result.Count) / float64(limit))),
		TotalData: result.Count,
	}

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	if pageInfo.Page <= 0 || pageInfo.Page > pageInfo.LastPage {
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "not found",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseWPage{
		Success:  true,
		Message:  "list all products",
		Pageinfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailPromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.FindOnePromo(&id)

	if err != nil {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "detail user",
		Results: user,
	})
}

func CreatePromo(c *gin.Context) {
	data := services.Promo{}

	c.ShouldBind(&data)

	var col []string
	var values []string

	val := reflect.ValueOf(&data)
	types := val.Type()

	for i := 0; i < val.Elem().NumField(); i++ {
		x := fmt.Sprintf(`%v`, types.Elem().Field(i).Name)

		if x != "Id" && x != "CreatedAt" && x != "UpdatedAt" {

			if !val.Elem().Field(i).IsNil() {
				if fmt.Sprint(types.Elem().Field(i).Type) == "*int" || fmt.Sprint(types.Elem().Field(i).Type) == "*float64" {
					values = append(values, fmt.Sprint(val.Elem().FieldByName(x).Elem().Interface()))
				} else if fmt.Sprint(types.Elem().Field(i).Type) == "*bool" {
					values = append(values, "true")
				} else {
					values = append(values, fmt.Sprint(`'`, val.Elem().FieldByName(x).Elem().Interface(), `'`))
				}
				col = append(col, fmt.Sprint(`"`, strings.ToLower(x[:1])+x[1:], `"`))

			}
		}
	}
	user, err := models.CreatePromo(col, values)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "create user successfully",
		Results: user,
	})
}

func UpdatePromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data := services.Promo{}

	c.ShouldBind(&data)

	var values []string

	val := reflect.ValueOf(&data)
	types := val.Type()

	for i := 0; i < val.Elem().NumField(); i++ {
		x := fmt.Sprintf(`%v`, types.Elem().Field(i).Name)

		if x != "Id" && x != "CreatedAt" && x != "UpdatedAt" {

			if !val.Elem().Field(i).IsNil() {
				if fmt.Sprint(types.Elem().Field(i).Type) == "*int" || fmt.Sprint(types.Elem().Field(i).Type) == "*float64" {
					y := fmt.Sprint(`"`, strings.ToLower(x[:1])+x[1:], `"=`, val.Elem().FieldByName(x).Elem().Interface())
					values = append(values, y)
				} else if fmt.Sprint(types.Elem().Field(i).Type) == "*bool" {
					y := fmt.Sprint(`"`, strings.ToLower(x[:1])+x[1:], `"=`, "true")
					values = append(values, y)
				} else {
					y := fmt.Sprint(`"`, strings.ToLower(x[:1])+x[1:], `"='`, val.Elem().FieldByName(x).Elem().Interface(), `'`)
					values = append(values, y)
				}
			}
		}
	}
	user, err := models.UpdatePromo(values, id)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "update user successfully",
		Results: user,
	})
}

func DeletePromo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.DeletePromo(id)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "delete user successfully",
		Results: user,
	})
}
