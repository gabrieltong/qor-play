package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

type Play struct {
}

func (play *Play) ConfigureApplication(app *config.App) {
	Play := app.Admin.AddResource(&models.Play{})

	for _, state := range models.PlayState {
		var state = state
		Play.Scope(&admin.Scope{
			Name:  state,
			Group: "State",
			Handler: func(db *gorm.DB, context *qor.Context) *gorm.DB {
				return db.Where("state = ?", state)
			},
		})
	}

	Play.Meta(&admin.Meta{Name: "Desc", Config: &admin.RichEditorConfig{Plugins: []admin.RedactorPlugin{
		{Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
		{Name: "table", Source: "/vendors/redactor_table.js"},
	},
		Settings: map[string]interface{}{
			"medialibraryUrl": "/admin/product_images",
		},
	}})

	Play.Meta(&admin.Meta{Name: "ClueTplType", Config: &admin.SelectOneConfig{Collection: models.ClueTplType, AllowBlank: true}})
	Play.Meta(&admin.Meta{Name: "Type", Config: &admin.SelectOneConfig{Collection: models.PlayType, AllowBlank: true}})
	Play.Meta(&admin.Meta{Name: "Level", Config: &admin.SelectOneConfig{Collection: models.PlayLevel, AllowBlank: true}})
	Play.Meta(&admin.Meta{Name: "RoundSize", Config: &admin.SelectOneConfig{Collection: models.PlayRoundSize, AllowBlank: true}})
	Play.Meta(&admin.Meta{Name: "Size", Config: &admin.SelectOneConfig{Collection: models.PlaySize, AllowBlank: true}})

	Play.IndexAttrs([]string{"Title", "Writer", "OnlineAt", "State", "Type", "Level", "Size", "GamesOverCount"})

	Play.NewAttrs(
		&admin.Section{
			Title: "Base Information",
			Rows: [][]string{
				{"Title", "Writer"},
				{"RongId", "Debug"},
				{"OnlineAt"},
			},
		},
		&admin.Section{
			Title: "Setting",
			Rows: [][]string{
				{"ClueTplType", "Type", "Size"},
				{"Level", "RoundSize"},
			},
		},
		&admin.Section{
			Title: "Time Limit",
			Rows: [][]string{
				{"TmRead", "TmReadLimit"},
				{"TmIntroPlayer", "TmIntroPlayerLimit"},
				{"TmIntroAll", "TmIntroAllLimit"},
				{"TmDiscussAll", "TmDiscussAllLimit"},
				{"TmRound", "TmRoundLimit"},
				{"TmPoll", "TmPollLimit"},
			},
		},
		"desc_pre_1",
		"desc_pre_2",
		"desc",
		"conclusion",
		"Actors",
	)

	// productPropertiesRes := product.Meta(&admin.Meta{Name: "ProductProperties"}).Resource
	Actor := Play.Meta(&admin.Meta{Name: "Actors"}).Resource
	Actor.NewAttrs(
		&admin.Section{
			Rows: [][]string{
				{"Name", "PublicClueLimit"},
				{"Show", "IsNpc", "IsKiller"},
				{"Point1", "Point2"},
				{"Point3", "Point4"},
				{"DefaultIntro"},
				{"ShortDesc"},
				{"Desc"},
			},
		},
		"TaskTpls",
	)

	TaskTpls := Actor.Meta(&admin.Meta{Name: "TaskTpls"}).Resource
	TaskTpls.NewAttrs("Title", "TaskTplAnswers")

	TaskTplAnswers := TaskTpls.Meta(&admin.Meta{Name: "TaskTplAnswers"}).Resource
	TaskTplAnswers.NewAttrs(&admin.Section{
		Rows: [][]string{
			{"Title", "IsAnswer"},
		},
	})
	// &admin.SelectOneConfig{Collection: [][]string{{"video", "Video"}, {"image", "Image"}, {"file", "File"}, {"video_link", "Video Link"}}},
	// Play.Meta(&admin.Meta{Name:})

}
