package db

type Url struct {
	ID    string `json:"id" gorm:"column:id;primary key"`
	Short string `json:"short" gorm:"column:short"`
	Long  string `json:"long" gorm:"column:long"`
}
