package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	vault "github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	// Read a secret from the default mount path for KV v2
	secretPath := "kv/data/my-secret"
	secret, err := client.Logical().Read(secretPath)
	if err != nil {
		log.Fatalf("Failed to read database credentials from Vault: %v", err)
	}
	if secret == nil {
		log.Fatalf("No credentials found at path: %s", secretPath)
	}

	// get data from secret
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalf("Failed to parse secret data")
	}

	// save postgres_username and postgres_password to the variable
	postgresUsername, usernameOk := data["postgres_username"].(string)
	postgresPassword, passwordOk := data["postgres_password"].(string)

	if !usernameOk || !passwordOk {
		log.Fatalf("Failed to retrieve username or password from secret data")
	}

	fmt.Printf("Username: %s\n", postgresUsername)
	fmt.Printf("Password: %s\n", postgresPassword)

	// port := 5432

	// // Connect to PostgreSQL
	// dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=vault port=%d sslmode=disable TimeZone=Asia/Shanghai", postgresUsername, postgresPassword, port)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("failed to connect database: %v", err)
	// }

	// connection, err := db.DB()
	// if err != nil {
	// 	log.Fatalf("failed to get database connection: %v", err)
	// }
	// defer connection.Close()

	// // Migrate the schema (auto-create table if not exists)
	// err = db.AutoMigrate(&User{})
	// if err != nil {
	// 	log.Fatalf("failed to migrate schema: %v", err)
	// }

	// // Initialize Fiber
	// app := fiber.New()
	// app.Use(cors.New())

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	// // Endpoint get all users
	// app.Get("/users", func(c *fiber.Ctx) error {
	// 	var users []User
	// 	result := db.Find(&users) // Menyimpan hasil pencarian ke dalam variabel result
	// 	if result.Error != nil {  // Menangani error jika ada
	// 		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	// 	}
	// 	return c.JSON(users)
	// })

	// //endpoint create user
	// app.Post("/users", func(c *fiber.Ctx) error {
	// 	var user User
	// 	if err := c.BodyParser(&user); err != nil {
	// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	// 	}
	// 	db.Create(&user)
	// 	return c.JSON(user)
	// })

	// // Jalankan aplikasi Fiber
	// err = app.Listen(":3000")
	// if err != nil {
	// 	log.Fatalf("failed to start server: %v", err)
	// }
}
