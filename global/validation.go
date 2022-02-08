package global

import (
	 "github.com/astaxie/beego/validation"
)

func init() {
    validation.SetDefaultMessage(map[string]string{
        "Required":     "Tidak boleh kosong",
        "Min":          "Nilai minimal yang di izinkan %d",
        "Max":          "Nilai maksimal yang di izinkan %d",
        "Range":        "Harus dalam kisaran dari %d sampai %d",
        "MinSize":      "Panjang minimal yang di izinkan %d",
        "MaxSize":      "Panjang maksimal yang di izinkan %d",
        "Length":       "Panjangnya harus sama dengan %d",
        "Alpha":        "Harus terdiri dari huruf",
        "Numeric":      "Harus berupa angka",
        "AlphaNumeric": "Harus berupa huruf atau angka",
        "Match":        "Harus cocok %s",
        "NoMatch":      "Seharusnya tidak cocok %s",
        "AlphaDash":    "Harus berupa huruf, angka, atau simbol (-_)",
        "Email":        "Email harus dalam format yang benar",
        "IP":           "Harus memiliki alamat IP yang benar",
        "Base64":       "Harus dalam format base64 yang benar",
        "Mobile":       "Harus nomor ponsel yang benar",
        "Tel":          "Harus nomor telepon yang benar",
        "Phone":        "Harus nomor telepon atau ponsel yang benar",
        "ZipCode":      "Kode pos harus benar",
    })
}	