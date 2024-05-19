package models

type Product struct {
	Id          int64  `gorm: "primaryKey", json: "id"`
	NamaProduct string `gorm: "type: varchar(300)", json: "nama_product"`
	Deskripsi   int64  `gorm: "type: text", json: "deskripsi"`
}
