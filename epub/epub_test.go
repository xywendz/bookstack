package epub

import (
	"log"
	"os"
	"path"
	"testing"
	"time"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/TruthHun/BookStack/commands"
	"github.com/TruthHun/BookStack/models"
	_ "github.com/go-sql-driver/mysql"
)

func TestImportEpub(t *testing.T) {
	commands.ConfigurationFile = "../conf/app.conf"
	commands.ResolveCommand(os.Args)
	bookName := "从零开始读懂经济学、金融学、投资理财学"
	bookPath := "C:\\Users\\Administrator.DESKTOP-O1GNEVF\\Downloads\\《从零开始读懂经济学、金融学、投资理财学》\\从零开始读懂经济学、金融学、投资理财学.epub"
	bk, err := Open(bookPath)
	if err != nil {
		t.Fatal(err)
	}
	defer bk.Close()
	c, err := bk.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	book := models.NewBook()
	book.BookName = bookName
	book.ParentId = 0
	book.Identify = "jinrong-jingji"
	book.Author = "管理员"
	book.PrivatelyOwned = 1
	book.MemberId = 1
	book.Editor = "markdown"
	book.Theme = "default"
	book.Version = time.Now().Unix()
	book.ReleaseTime = time.Now()
	book.GenerateTime = time.Now()
	book.LastClickGenerate = time.Now()

	book.Insert()

	bookResult, err := models.NewBookResult().FindByIdentify(book.Identify, 1)
	if err != nil {
		t.Fatal(err)
	}
	bookId := bookResult.BookId
	for _, v := range c {
		markdown, err := htmltomarkdown.ConvertString(v[1])
		if err != nil {
			log.Fatal(err)
		}
		ch := v[0]
		doc := models.NewDocument()
		doc.BookId = bookId
		doc.MemberId = 1
		doc.ParentId = 0
		doc.DocumentName = ch
		doc.Markdown = markdown
		doc.Identify = time.Now().Format("20060102150405.000")
		doc.Version = time.Now().Unix()
		htmlContent := MarkdownToHtml([]byte(markdown))
		doc.Release = string(htmlContent)
		docIdInt64, _ := doc.InsertOrUpdate()
		ds := models.NewDocumentStore()
		ds.DocumentId = int(docIdInt64)
		ds.Markdown = markdown
		ds.Content = string(htmlContent)
		ds.InsertOrUpdate(*ds, "markdown")
	}
}

func TestImportBookDir(t *testing.T) {
	commands.ConfigurationFile = "../conf/app.conf"
	commands.ResolveCommand(os.Args)
	bookName := "A-B测试从0到1"
	bookPath := "C:\\workspace\\geektime-docs\\AI-大数据\\A-B测试从0到1\\docs"

	book := models.NewBook()
	book.BookName = bookName
	book.ParentId = 0
	book.Identify = "ai-test-0-1"
	book.Author = "管理员"
	book.PrivatelyOwned = 1
	book.MemberId = 1
	book.Editor = "markdown"
	book.Theme = "default"
	book.Version = time.Now().Unix()
	book.ReleaseTime = time.Now()
	book.GenerateTime = time.Now()
	book.LastClickGenerate = time.Now()

	book.Insert()

	bookResult, err := models.NewBookResult().FindByIdentify(book.Identify, 1)
	if err != nil {
		t.Fatal(err)
	}
	bookId := bookResult.BookId
	bk, _ := os.Open(bookPath)
	chs, _ := bk.Readdir(0)
	for _, v := range chs {
		markdown, _ := os.ReadFile(path.Join(bookPath, v.Name()))
		doc := models.NewDocument()
		doc.BookId = bookId
		doc.MemberId = 1
		doc.ParentId = 0
		doc.DocumentName = v.Name()
		doc.Markdown = string(markdown)
		doc.Identify = time.Now().Format("20060102150405.000")
		doc.Version = time.Now().Unix()
		htmlContent := MarkdownToHtml([]byte(markdown))
		doc.Release = string(htmlContent)
		docIdInt64, _ := doc.InsertOrUpdate()
		ds := models.NewDocumentStore()
		ds.DocumentId = int(docIdInt64)
		ds.Markdown = string(markdown)
		ds.Content = string(htmlContent)
		ds.InsertOrUpdate(*ds, "markdown")
	}
}

func TestToMarkdown(t *testing.T) {
	name := "2"
	bk, err := Open("C:\\Users\\Administrator.DESKTOP-O1GNEVF\\Downloads\\2.epub")
	if err != nil {
		t.Fatal(err)
	}
	defer bk.Close()
	targetDir := "d:\\" + name
	os.Mkdir(targetDir, os.ModePerm)
	c, err := bk.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range c {
		markdown, err := htmltomarkdown.ConvertString(v[1])
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(targetDir+"\\"+v[0]+".md", []byte(markdown), os.ModePerm)
	}
}

func TestEpub(t *testing.T) {
	bk, err := Open("C:\\Users\\Administrator.DESKTOP-O1GNEVF\\Downloads\\2.epub")
	if err != nil {
		t.Fatal(err)
	}
	defer bk.Close()
	c, err := bk.ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range c {
		markdown, err := htmltomarkdown.ConvertString(v[1])
		if err != nil {
			log.Fatal(err)
		}
		t.Log("title:" + v[0])
		t.Log("-----")
		t.Log(markdown)
	}
}
