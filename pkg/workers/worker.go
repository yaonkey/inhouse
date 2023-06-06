package workers

import (
	"log"
	"net/http"
	"sync"
	"time"
	"yaonkey/inhouse/pkg/models"
)

// Функция, запускающая воркер,
// который каждую минуту будет проверять доступность сайтов
// с получением времени запроса
func run() {
	// переменные time chan int64, wg waitgroup
	timer := make(chan int64, 1)
	var wg sync.WaitGroup
	// получение списка всех сайтов
	sites, err := models.GetAllSites()
	if err != nil {
		log.Println("Error with get site's list")
	}
	// цикл for с перебором всех сайтов
	for _, site := range sites {
		wg.Add(1)
		// вызов асинхронной функции getSiteTime()
		go func(site models.Site) {
			defer wg.Done()
			start := time.Now()
			// запрос к сайту
			_, err := http.NewRequest(http.MethodGet, site.Url, nil)
			if err != nil {
				log.Printf("Some error with url: %s\n", site.Url)
			}
			// получение времени и запись в канал времени в наносекундах
			timer <- int64(time.Since(start).Nanoseconds())
			// вызов метода site UpdateTime(передаем новое время)
			go site.UpdateTime(<-timer)
		}(site)
	}
	wg.Wait()
}

// Фоновый запуск воркера
func Background() {
	for {
		run()
		time.Sleep(60 * time.Second)
	}
}
