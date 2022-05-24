# go-and-scrape
Golang parallel web scrapper of news portals. It scraps and persists in a SQLite DB file
article previews and entire articles from different media sources. 

## configuration
App can be configured in `config.yaml` file. The file has two sections: `mediums`
and `persistence`. The first one has the following structure:
```
mediums:
...
  - medium:                            #list of media
    name: ABC                          #name of the medium
    medium-url: https://xxx.com/       #base url for scrapping (where article previews are placed)
    html-preview-tags:                 #CSS-selectors for article's previews:
      article: .css-selector-4           #article DOM-element CSS-selector
      tag: .css-selector-5               #news tag DOM-element CSS-selector
      title: .css-selector-6             #title DOM-element CSS-selector
      subtitle: .css-selector-7          #subtitle DOM-element CSS-selector
      article-url: .css-selector-8       #the DOM-element which has href attribute of link to the entire article
    html-article-tags:                 #CSS-selectors for the entire articles:
      author: .css-selector-1            #author DOM-element CSS-selector
      date: .css-selector-2              #date DOM-element CSS-selector
      text: .css-selector-3              #text DOM-element CSS-selector
...
```
`persistence` section can be configured as well:
  * `filename` - name of SQLite DB file (without extension),
  * `interval` - scrapper's int parameter of scheduling interval in seconds. If `0`, app executes just one time and finishes. Default: `0`.

## run scrapper
```
docker-compose up
go run main.go
```

