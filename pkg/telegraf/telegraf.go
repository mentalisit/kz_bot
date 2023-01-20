package telegraf

import (
	"github.com/anonyindian/telegraph-go"
	"github.com/sirupsen/logrus"
)

type Telegraf struct {
	log *logrus.Logger
	a   *telegraph.Account
}

func (t *Telegraf) InitTelegraf(log *logrus.Logger) {
	t.log = log

	//Use this method to create account
	a, err := telegraph.CreateAccount("kzbot", &telegraph.CreateAccountOpts{
		AuthorName: "Кз Бот, Автор Mentalisit",
	})
	if err != nil {
		t.log.Println(err.Error())
		return
	}
	t.a = a
}
func (t *Telegraf) CreatePage() {
	// The Telegraph API uses a DOM-based format to represent the content of the page.
	// https://telegra.ph/api#Content-format
	_, err := t.a.CreatePage("SampleDimpl", `<h3>SampleDimpl Page #1</h3> <p>Hello world Dimpl! This telegraph page is created using telegraph-go package.</p><br><a href="https://github.com/anonyindian/telegraph-go">Click here to open package</a>`, &telegraph.PageOpts{
		AuthorName: "kzbot",
	})
	if err != nil {
		t.log.Println(err.Error())
	}

	_, err = t.a.CreatePage("Sample", `<h3>Sample Page #2</h3> <p>Hello world! This telegraph page is created using telegraph-go package.</p>`, &telegraph.PageOpts{
		AuthorName: "User1",
	})
	if err != nil {
		t.log.Println(err.Error())
	}

}
func (t *Telegraf) CreatePageUserStatistic(title, contents string) string {
	//content := fmt.Sprintf(`<p> %s </p>`, contents)

	_, err := t.a.CreatePage(title, contents, &telegraph.PageOpts{AuthorName: "kzbot"})
	if err != nil {
		t.log.Println("Ошибка создания страницы ", err.Error(), err)
	}
	url := ""
	plist, _ := t.a.GetPageList(nil)
	for _, page := range plist.Pages {
		// you can print all pages with the help of loop
		url = page.Url
	}
	return url
}
func (t *Telegraf) ReturtPageList() {
	// Get a list of pages in your current account with this method
	plist, _ := t.a.GetPageList(nil)
	for _, page := range plist.Pages {
		// you can print all pages with the help of loop
		t.log.Println(page.Url)
	}

	// Get total pages count in this way
	pcount := plist.TotalCount
	t.log.Println(pcount)
}
func (t *Telegraf) TestingFunc() {
	t.CreatePageUserStatistic("Статистика игрока Mentalisit", "Контент статистики ")
	//t.ReturtPageList()
}
