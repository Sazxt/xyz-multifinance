package models

import "gorm.io/gorm"

type Transaction struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	ContractNo  string  `gorm:"unique;not null"`
	OTR         float64 `gorm:"not null"`
	AdminFee    float64 // Angka admin transaksi barang baik itu  White Godds, Motor atau Mobil konsumen.
	Installment float64 // Angka jumlah cicilan transaksi barang baik itu White Godds, Motor atau Mobil konsumen
	Interest    float64 // Angka bunga yang ditagihkan setiap transaksi barang baik itu White Godds, Motor atau Mobil konsumen
	AssetName   string  // Nama barang yang dibeli konsumen
}

func MigrateTransaction(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}
