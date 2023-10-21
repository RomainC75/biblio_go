package seeder

// import (
// "fmt"
// "log"

// articleModel "gitub.com/RomainC75/biblio/internal/modules/article/models"
// userModel "gitub.com/RomainC75/biblio/internal/modules/user/models"
// "gitub.com/RomainC75/biblio/pkg/database"

// "golang.org/x/crypto/bcrypt"
// )

func Seed() {
	// // Seeder logic
	// db := database.Connection()

	// pass := "secret"
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	// if err != nil {
	// 	log.Fatal("hass pass error ")
	// 	return
	// }
	// user := userModel.User{Name: "Random Name", Email: "random@email.com", Password: string(hashedPassword)}
	// db.Create(&user)

	// log.Printf("user created successfully with email adress: %s \n", user.Email)

	// for i := 1; i <= 10; i++ {
	// 	article := articleModel.Article{Title: fmt.Sprintf("Random Tile %d", i), Content: fmt.Sprintf("Content %d", i), UserId: 1}
	// 	db.Create(&article)

	// 	log.Printf("Article created successfully with title : %s \n", article.Title)
	// }
	// log.Println("Seeder done")
}
