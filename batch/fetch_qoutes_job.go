package batch

import (
	"log"
	"strings"

	"github.com/goodgoodjm/peter/models"
	"gorm.io/gorm"
)

const SLICE_SIZE = 300

type fetchQoutesJob struct {
	db *gorm.DB
}

func (job fetchQoutesJob) Run() {
	var registrationsGroups []models.RegistrationGroup
	exec := job.db.
		Model(&models.Registration{}).
		Select("symbol, count(*) as count").
		Group("symbol").
		Order("count DESC").
		Find(&registrationsGroups)

	if err := exec.Error; err != nil {
		log.Fatal(err.Error())
		return
	}

	totalLength := len(registrationsGroups)
	for i := 0; i <= totalLength/SLICE_SIZE; i++ {
		start := i * SLICE_SIZE
		end := (i + 1) * SLICE_SIZE
		if end > totalLength {
			end = totalLength
		}
		slice := registrationsGroups[start:end]
		go fetchQuotes(job.db, slice)

	}
}

func transformQuotesToPrices(quotes []Quote) []models.Price {
	prices := make([]models.Price, 0)
	for _, quote := range quotes {
		prices = append(prices, models.Price{
			Symbol: quote.Symbol,
			Price:  quote.Price,
		})
	}
	return prices

}

func fetchQuotes(db *gorm.DB, slice []models.RegistrationGroup) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln("Fail to fetch quotes.", r)
		}
	}()

	tickers := []string{}
	for _, registration := range slice {
		tickers = append(tickers, registration.Symbol)
	}
	symbols := strings.Join(tickers, ",")

	quotes, err := FetchQuotes(symbols)
	if err != nil {
		panic(err)
	}

	prices := transformQuotesToPrices(quotes)
	if err := db.Create(prices).Error; err != nil {
		panic(err)
	}

	log.Println("Success to fetch quotes. ", symbols)
}
