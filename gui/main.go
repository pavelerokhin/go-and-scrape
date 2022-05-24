package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/pavelerokhin/go-and-scrape/gui/cli"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

func main() {
	logger := log.New(os.Stdout, "gui ", log.LstdFlags|log.Lshortfile)

	var mode string
	if len(os.Args) < 2 {
		//logger.Fatalln("no GUI mode has been set")
		mode = "cli"
	}
	//mode = os.Args[1]

	if mode == "cli" {

		news := []entities.ArticlePreview{
			{
				Title:    "Российское юрлицо Google собирается объявить о банкротстве",
				Subtitle: "",
			},
			{
				Title:    "Леонид Слуцкий возглавил фракцию ЛДПР в Госдуме",
				Subtitle: "",
			},
			{
				Title:    "Александр Роднянский снимет сериал про Путина.  Он будет основан на книге «Вся кремлевская рать» Михаила Зыгаря",
				Subtitle: "",
			},
			{
				Title:    "Лукашенко подписал закон о смертной казни за покушение на совершение теракта.  По этой статье в Беларуси обвиняют оппозиционеров",
				Subtitle: "",
			},
			{
				Title:    "ФСБ отчиталась о задержании «сторонника украинских нацистов»,  подозреваемого в повреждении ЛЭП в Кемерово",
				Subtitle: "",
			},
			{
				Title:    "Как читать ресурсы,  заблокированные российскими властями",
				Subtitle: "Простая инструкция «Медузы»",
			},
			{
				Title:    "Защитников «Азовстали» доставили в бывшую колонию под Донецком.  Всего с завода вывезли 959 украинских военных",
				Subtitle: "",
			},
			{
				Title:    "Швеция и Финляндия подали заявки на вступление в НАТО",
				Subtitle: "",
			},
			{
				Title:    "В России вновь попытались поджечь военкомат.  «Коктейлями Молотова» забросали здание Щелковского военного комиссариата в Подмосковье",
				Subtitle: "",
			},
			{
				Title:    "Bloomberg: США планируют после 25 мая заблокировать выплаты России по госдолгу",
				Subtitle: "",
			},
		}

		p := tea.NewProgram(cli.PopulateGeneralNewsModel(news))
		if err := p.Start(); err != nil {
			logger.Printf("error implementing cli mode: %v", err)
			os.Exit(1)
		}

		return
	} else if mode == "http" {

		return
	}

	logger.Fatalf("cannot implement %s mode", mode)
}
