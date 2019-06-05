package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// Ebook Ebook entity
type Ebook struct {
	ID     int64    `dapper:"id,primarykey,autoincrement,table=ebooks"`
	Year   string   `dapper:"year"`
	Class  string   `dapper:"class"`
	Name   string   `dapper:"name"`
	Date   string   `dapper:"date"`
	Hash   string   `dapper:"hash"`
	Images []string `dapper:"-"`
	HTML   string   `dapper:"-"`
	CSS    string   `dapper:"-"`
}

// ResolveHash resolve content md5 hash
func (e *Ebook) ResolveHash() {
	var sb strings.Builder
	for _, img := range e.Images {
		sb.WriteString(img)
	}

	sb.WriteString(e.HTML)
	sb.WriteString(e.CSS)
	str := sb.String()

	hash := md5.Sum([]byte(str))
	e.Hash = hex.EncodeToString(hash[:])
}

// ResolveCloudCSS replace image link with oss image
func (e *Ebook) ResolveCloudCSS() string {
	return strings.Replace(e.CSS, "../img/", "https://dong-feng.oss-cn-shanghai.aliyuncs.com/ebook/img/", -1)
}

// ResolveCloudHTML replace style link with oss css
func (e *Ebook) ResolveCloudHTML() string {
	return strings.Replace(e.HTML, "./css/", fmt.Sprintf("https://dong-feng.oss-cn-shanghai.aliyuncs.com/ebook/css/%s/%s/%s/%s/", e.Year, e.Class, e.Name, e.Date), -1)
}
