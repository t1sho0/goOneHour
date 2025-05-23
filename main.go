package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Title       string
	Description template.HTML
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "Муминов Абдурашид — разработчик",
		Description: template.HTML(`
<p>Привет! Меня зовут Абдурашид. Я программист, пока не заточенный под одно направление.<br>
Мне интересно развиваться в мире IT, постоянно узнавая что-то новое и необычное.<br>
На данный момент я изучаю язык программирования Go, и он мне очень нравится — вдохновляет делать что-то новое.</p>
<p><strong>Технологии, с которыми я знаком:</strong></p>
<ul style="text-align: left; max-width: 600px; margin: auto;">
<li>Python</li>
<li>Django</li>
<li>Aiogram, Telebot</li>
<li>SQLite3 (интеграция с Python)</li>
<li>Terminal (Bash) — немного</li>
<li>BeautifulSoup и другие библиотеки для парсинга</li>
<li>Pandas (для работы с Excel)</li>
<li>Go — в процессе изучения</li>
<li>Docker (теоретически)</li>
<li>Git, GitHub — базовый контроль версий</li>
</ul>
<p>Желаю хорошего настроения! Если вы оставите email — возможно, мы поработаем вместе в будущем 🙂</p>`),
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("static/images"))))

	fmt.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}