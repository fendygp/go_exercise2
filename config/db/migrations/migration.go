package migrations

import "bitbucket.org/bridce/go_exercise2/config/db"

func init() {
	/* init automigrate DB
	Automigrate( your model )
	*/

}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
