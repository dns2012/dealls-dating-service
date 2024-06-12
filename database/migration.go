package main

import (
	"fmt"
	"github.com/dns2012/dealls-dating-service/app/domain/constant"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/manager"
	"github.com/dns2012/dealls-dating-service/config"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	passwordManager = manager.NewPasswordManager(bcrypt.DefaultCost)
)

func main() {
	db := config.GormDB()

	entities := []interface{}{
		&entity.User{},
		&entity.Profile{},
		&entity.Preference{},
		&entity.Package{},
		&entity.UserPackage{},
	}

	db.Migrator().DropTable(entities...)

	err := db.AutoMigrate(entities...)

	if err != nil {
		log.Fatalf("Failed to gorm migrate: ", err)
	}

	// NOTE: Setup Seeder For Base Development Data Example
	for i := 0; i < 20; i++ {
		userSeeder(db, i)
		if i > 1 && i <= 5 {
			preferenceSeeder(db, i)
		}
		if i >= 1 && i <= 3 {
			packageSeeder(db, i)
		}
	}
}

func userSeeder(db *gorm.DB, i int) {
	fullName := fmt.Sprintf("User%d Random", i+1)
	name := strings.Fields(fullName)
	password, _ := passwordManager.CreateHashPassword(fmt.Sprintf("password%d", i+1))
	parsedTime, _ := time.Parse("2006-01-02", "1997-06-30")

	db.Omit("premium_at").Create(&entity.User{
		Nickname: name[0],
		Email:    fmt.Sprintf("user%d@example.com", i+1),
		Password: password,
		Profile: entity.Profile{
			FullName: fullName,
			ImageUrl: constant.DUMMY_IMAGE_URL,
			BirthAt:  parsedTime,
			Gender: uint(rand.Intn(constant.GENDER_FEMALE)) + 1,
			Company:  "Dealls",
			JobTitle: "Backend Engineer",
		},
	})
}

func preferenceSeeder(db *gorm.DB, i int) {
	db.Create(&entity.Preference{
		UserID:           1,
		PreferenceUserID: uint(i),
		PreferenceType:   uint(rand.Intn(constant.PREFERENCE_TYPE_LIKE)) + 1,
	})
}

func packageSeeder(db *gorm.DB, i int) {
	unlimitedSwap := false
	totalSwap := i * 50
	if i == 3 {
		unlimitedSwap = true
		totalSwap = 0
	}
	db.Create(&entity.Package{
		Name:            fmt.Sprintf("Package %d", i),
		Description:     fmt.Sprintf("Package %d For Your Freedom Swap", i),
		Price: uint(i * 50000),
		UnlimitedSwap:   unlimitedSwap,
		TotalSwapPerDay: uint(totalSwap),
	})
}