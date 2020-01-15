package models

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Ebook Ebook entity
type Ebook struct {
	ID        int64    `dapper:"id,primarykey,autoincrement,table=ebooks"`
	Year      string   `dapper:"year"`
	Class     string   `dapper:"class"`
	Name      string   `dapper:"name"`
	Date      string   `dapper:"date"`
	Hash      string   `dapper:"hash"`
	Converted bool     `dapper:"converted"`
	Images    []string `dapper:"-"`
	HTML      string   `dapper:"-"`
	CSS       string   `dapper:"-"`
	Dates     []string `dapper:"-"`
}

// ResolveHash resolve content md5 hash
func (e *Ebook) ResolveHash() {
	var sb strings.Builder

	sb.WriteString(e.Year)
	sb.WriteString(e.Class)
	sb.WriteString(e.Name)
	sb.WriteString(e.Date)

	for _, img := range e.Images {
		sb.WriteString(img)
	}

	sb.WriteString(e.HTML)
	sb.WriteString(e.CSS)
	str := sb.String()

	hash := md5.Sum([]byte(str))
	e.Hash = hex.EncodeToString(hash[:])
}

// ResolveCloudCSS replace image link
func (e *Ebook) ResolveCloudCSS() string {
	return strings.Replace(e.CSS, "../img/", "../../../../../../img/", -1)
}

// ResolveCloudHTML replace style link with oss css
func (e *Ebook) ResolveCloudHTML() string {
	return strings.Replace(e.HTML, "./img/", "../../../../../img/", -1)
}
