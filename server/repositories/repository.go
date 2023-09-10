package repositories

type repository struct {
	// db *gorm.DB
}

func MakeRepository() *repository {
	return &repository{
		// db: db,
	}
}
