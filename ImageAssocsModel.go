package files

import (
	"time"

	"github.com/go-catupiry/catu/models"
)

// ImageAssocsModel - Model to associate a image with a record of any type
type ImageAssocsModel struct {
	models.Base

	ModelName string `gorm:"column:modelName;type:varchar(255);" json:"modelName"`
	ModelID   int64  `gorm:"column:modelId;type:bigint(20);" json:"modelId"`
	Field     string `gorm:"column:field;type:varchar(255);" json:"field"`
	Order     int    `gorm:"column:order;type:int(11);" json:"order"`

	ImageID int64      `gorm:"column:imageId;type:int(11);" json:"imageId"`
	Image   ImageModel `gorm:"foreignKey:imageId" json:"image"`

	CreatedAt time.Time `gorm:"column:createdAt;" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt;" json:"updatedAt"`
}

// TableName ...
func (m *ImageAssocsModel) TableName() string {
	return "imageassocs"
}
