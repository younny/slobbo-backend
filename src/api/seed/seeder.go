package seed

import (
	"github.com/younny/slobbo-backend/src/types"
	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
)

var posts = []types.Post{
	{
		Title:     "Finibus Bonorum et Malorum",
		SubTitle:  "China",
		Body:      "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?",
		Author:    "Slobbo",
		Category:  0,
		Thumbnail: "https://i.picsum.photos/id/290/457/300.jpg?hmac=vVVo8oVh4CrY7ddKOqBKNaTs1a9KJkHD2biSJKshv1g",
	},
	{
		Title:     "Sed ut perspiciatis",
		SubTitle:  "Korea",
		Body:      "Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?",
		Author:    "Slobbo",
		Category:  1,
		Thumbnail: "https://i.picsum.photos/id/290/457/300.jpg?hmac=vVVo8oVh4CrY7ddKOqBKNaTs1a9KJkHD2biSJKshv1g",
	},
}

func Load(log *zap.Logger, db *gorm.DB) {
	if err := db.Debug().DropTableIfExists(&types.Post{}).Error; err != nil {
		log.Fatal("Couldn't drop table: ", zap.Error(err))
	}

	if err := db.Debug().AutoMigrate(&types.Post{}).Error; err != nil {
		log.Fatal("Couldn't migrate table: ", zap.Error(err))
	}

	for i, _ := range posts {
		if err := db.Debug().Model(&types.Post{}).Create(&posts[i]).Error; err != nil {
			log.Fatal("Couldn't seed posts table: ", zap.Error(err))
		}
	}
}
