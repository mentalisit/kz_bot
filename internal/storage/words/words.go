package words

type Words struct {
	ru map[string]string
	ua map[string]string
	en map[string]string
}

func NewWords() *Words {
	w := &Words{
		ru: make(map[string]string),
		ua: make(map[string]string),
		en: make(map[string]string),
	}
	w.setWords()

	return w
}

func (w *Words) setWords() {
	w.setWordsUa()
	w.setWordsRu()
	w.setWordsEn()

}
func (w *Words) GetWords(lang string, key string) string {
	if lang == "ru" {
		return w.ru[key]
	} else if lang == "ua" {
		return w.ua[key]
	} else if lang == "en" {
		return w.en[key]
	} else if lang == "" {
		return w.ru[key]
	}
	return "ошибка слов"
}
