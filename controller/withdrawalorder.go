// controller/withdrawal.go

package controller

import (
	"net/http"
	"one-api/model" // 替换为您项目的正确导入路径
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetWithdrawalOrdersEndpoint(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// 如果无法将参数转换为整数，则尝试从上下文中获取整数类型的ID
		id = c.GetInt("id")
	}

	// 假设 model.GetUserById 函数会根据提供的ID获取用户信息
	user, err := model.GetUserById(id, true)
	if err != nil {
		// 如果发生错误，返回错误信息
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	orders, err := model.GetUserWithdrawalOrders(uint(user.Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "无法获取提现订单",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    orders,
	})
}

func GetAllWithdrawalOrdersEndpoint(c *gin.Context) {

	orders, err := model.GetAllWithdrawalOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "无法获取提现订单",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    orders,
	})
}

func UpdateWithdrawalOrderStatusEndpoint(c *gin.Context) {

	var request struct {
		OrderID     uint `json:"order_id" binding:"required"`
		Status      int  `json:"status" binding:"required"`
		ProcessorID uint `json:"processor_id"` // 处理该订单的管理员ID
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "请求参数错误: " + err.Error(),
		})
		return
	}

	err := model.UpdateWithdrawalOrderStatus(request.OrderID, request.Status, request.ProcessorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "更新提现订单状态失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "提现订单状态更新成功",
	})
}
