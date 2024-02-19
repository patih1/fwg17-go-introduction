package customer

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patih1/fwg17-go-backend/src/controllers"
	"github.com/patih1/fwg17-go-backend/src/models"
	"github.com/patih1/fwg17-go-backend/src/services"
)

func CreateOrders(c *gin.Context) {
	order := services.Orders{}

	c.ShouldBind(&order)

	var col []string
	var values []string

	user := controllers.GetLogedId(c)
	currentDate := fmt.Sprintf("#%v", time.Now().Format("02-01-2006"))

	order.UserId = &user
	order.OrderNumber = &currentDate

	val := reflect.ValueOf(&order)
	types := val.Type()

	for i := 0; i < val.Elem().NumField(); i++ {
		x := fmt.Sprintf(`%v`, types.Elem().Field(i).Name)

		if x != "Id" && x != "CreatedAt" && x != "UpdatedAt" {

			if !val.Elem().Field(i).IsNil() {
				if fmt.Sprint(types.Elem().Field(i).Type) == "*int" {
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
	order, err := models.CreateOrders(col, values)

	col = nil
	values = nil

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, &services.ResponseOnly{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	orderDetail := services.OrderDetails{}
	c.ShouldBind(&orderDetail)

	val = reflect.ValueOf(&orderDetail)
	types = val.Type()

	price, _ := models.FindProductPrice(orderDetail.ProductId)
	variantPrice, _ := models.FindPVAdditionalPrice(orderDetail.ProductVariantId)
	sizePrice, _ := models.FindPSAdditionalPrice(orderDetail.ProductSizeId)
	quantity := *orderDetail.Quantity

	orderDetail.OrderId = &order.Id
	total := (*price.BasePrice + *variantPrice.AdditionalPrice + *sizePrice.AdditionalPrice) * quantity
	orderDetail.SubTotal = &total

	for i := 0; i < val.Elem().NumField(); i++ {
		x := fmt.Sprintf(`%v`, types.Elem().Field(i).Name)

		if x != "Id" && x != "CreatedAt" && x != "UpdatedAt" {

			if !val.Elem().Field(i).IsNil() {
				if fmt.Sprint(types.Elem().Field(i).Type) == "*int" {
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

	tax := total * 1 / 10

	models.CreateOrderDetails(col, values)
	result, _ := models.UpdateOrdersCS(order.Id, total, tax)

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "create order successfully",
		Results: result,
	})
}

func HistoryOrder(c *gin.Context) {
	id := controllers.GetLogedId(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset := (page - 1) * limit
	result, err := models.FindHistoryOrder(limit, offset, id)

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

	if pageInfo.TotalData == 0 {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "not data",
		})
		return
	}

	if pageInfo.Page <= 0 || pageInfo.Page > pageInfo.LastPage {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "not found",
		})
		return
	}

	c.JSON(http.StatusOK, &services.ResponseWPage{
		Success:  true,
		Message:  "list all order",
		Pageinfo: pageInfo,
		Results:  result.Data,
	})
}

func DetailHistory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := models.FindHistoryOD(&id)

	if err != nil {
		c.JSON(http.StatusNotFound, &services.ResponseOnly{
			Success: false,
			Message: "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, &services.Response{
		Success: true,
		Message: "order detail",
		Results: user,
	})
}
