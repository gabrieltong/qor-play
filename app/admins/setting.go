package admins

import (
	"github.com/gabrieltong/qor-play/app/models"
	"github.com/gabrieltong/qor-play/config"
	"github.com/qor/admin"
)

type Setting struct {
}

func (setting *Setting) ConfigureApplication(app *config.App) {
	Setting := app.Admin.AddResource(&models.Setting{}, &admin.Config{Singleton: true})
	Setting.EditAttrs(
		&admin.Section{
			Title: "Base",
			Rows: [][]string{
				{"Title", "Kefu"},
				{"SExp", "AExp", "BExp"},
				{"CExp", "LevelExp", "ExpLast"},
				{"CheckExp", "HighLevel", "AutoCancelOnwerOffline"},
			},
		},
		&admin.Section{
			Title: "Version",
			Rows: [][]string{
				{"IosVersion", "IosUrl"},
				{"AndroidVersion", "AndroidUrl"},
			},
		},
		&admin.Section{
			Title: "File",
			Rows: [][]string{
				{"AgreementFile", "AboutFile"},
				{"QrcodeFile", "GuideFile"},
			},
		},
		"PlayType",
		"PlayLevel",
		"PlaySize",
		"PlaySizeRange",
	)

	for _, Col := range []string{"PlayType", "PlayLevel", "PlaySize", "PlaySizeRange"} {
		Res := Setting.Meta(&admin.Meta{Label: "", Name: Col}).Resource
		Res.NewAttrs(&admin.Section{
			Rows: [][]string{{"Key", "V"}},
		})
		Res.EditAttrs(&admin.Section{
			Rows: [][]string{{"Key", "V"}},
		})
	}

}
