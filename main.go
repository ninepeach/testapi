package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func main() {
	// 创建一个长度为100的数组
	data := make([]string, 100)
	for i := 0; i < 100; i++ {
		data[i] = "Item " + strconv.Itoa(i)
	}

	// 初始化 Fiber 应用
	app := fiber.New()

	// 创建一个GET接口，获取所有数据
	app.Get("/items", func(c *fiber.Ctx) error {
		return c.JSON(data)
	})

	// 创建一个GET接口，根据ID查找数据
	app.Get("/items/:id", func(c *fiber.Ctx) error {
		// 获取ID参数并尝试转换为整数
		idStr := c.Params("id")

		// 校验参数是否为数字
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID: must be a number")
		}

		// 校验ID是否在数组范围内
		if id < 0 || id >= len(data) {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid ID: out of range")
		}

		// 返回找到的Item
		return c.JSON(data[id])
	})

	// 创建一个GET接口，返回JSON数组
	app.Get("/getdata", func(c *fiber.Ctx) error {
		// 模拟一个JSON数组数据
		jsonArray := []fiber.Map{
			{"id": 1, "name": "Alice", "age": 25},
			{"id": 2, "name": "Bob", "age": 30},
			{"id": 3, "name": "Charlie", "age": 22},
			{"id": 4, "name": "Diana", "age": 28},
		}

		// 返回JSON数组
		return c.JSON(jsonArray)
	})

	// 启动服务
	app.Listen(":3000")
}
