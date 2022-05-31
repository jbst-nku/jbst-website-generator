package articles

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Article struct{
	Title string
	Authors []string
	Part string
	Type string
	Year string
	Volume string
	Issue string
	Pages string
	Date string
	Abstract string
	Keywords string
	Correspondence string
	CorrespondenceEmail string
	ReceiveDate string
	AcceptDate string
	PublishDate string
	EISSN string
	DOI string
	Cite string
	PDFURL string
	DetailsURL string
}


func ParseCSV(path string) []Article{
	f, err := os.Open(path)
	if err != nil{
		log.Fatalln(err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil{
		log.Fatalln(err)
	}

	articles := make([]Article, len(records))
	for i, v := range records{
		title := strings.TrimSpace(v[1])
		part := strings.TrimSpace(v[2])
		articleType := strings.TrimSpace(v[3])
		year := strings.TrimSpace(v[4])
		volume := strings.TrimSpace(v[5])
		issue := strings.TrimSpace(v[6])
		pages := strings.TrimSpace(v[7])
		date := strings.TrimSpace(v[8])
		abstract := strings.TrimSpace(v[14])
		keywords := strings.TrimSpace(v[15])
		correspondence := strings.TrimSpace(v[16])
		correspondenceEmail := strings.TrimSpace(v[17])
		received := strings.TrimSpace(v[18])
		accepted := strings.TrimSpace(v[19])
		published := strings.TrimSpace(v[20])
		eissn := strings.TrimSpace(v[21])
		doi := strings.TrimSpace(v[22])
		cite := strings.TrimSpace(v[23])

		whiteSpace := regexp.MustCompile(`\r|\n|\t`)
		cite = whiteSpace.ReplaceAllString(cite, "")
		abstract = whiteSpace.ReplaceAllString(abstract, "")
		keywords = whiteSpace.ReplaceAllString(keywords, "")

		authors := make([]string,0)
		for i := 9; i < 14; i++{
			author := strings.TrimSpace(v[i])
			if author != ""{
				authors = append(authors, author)
			}else{
				break
			}
		}

		item := Article{
			Title:               title,
			Authors:             authors,
			Part:                strings.ToTitle(part),
			Type:                articleType,
			Year:                year,
			Volume:              volume,
			Issue:               issue,
			Pages:               pages,
			Date:                date,
			Abstract:            abstract,
			Keywords:            keywords,
			Correspondence:      correspondence,
			CorrespondenceEmail: correspondenceEmail,
			ReceiveDate:         received,
			AcceptDate:          accepted,
			PublishDate:         published,
			EISSN:               eissn,
			DOI:                 doi,
			Cite:                cite,
			PDFURL: fmt.Sprintf("articles/%v_%v_%v_%v.pdf", year, volume,issue, i + 1),
			DetailsURL: fmt.Sprintf("details/%v_%v_%v_%v.html", year, volume,issue, i + 1),
		}
		articles[i] = item
	}
	return articles
}
