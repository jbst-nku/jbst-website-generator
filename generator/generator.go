package generator

import (
	"fmt"
	"github.com/gusanmaz/jbst-processor/articles"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func IssueHTML(indexHTMLPath string, info articles.HTMLInfo) string{
	prefix := info.Issues.Prefix
	f, _ := os.Open(indexHTMLPath)
	content, _ := ioutil.ReadAll(f)
	contentS := string(content)

	contentS = strings.ReplaceAll(contentS, "{{.Issues.Prefix}}", prefix)
	t := template.Must(template.New("issue.html").
		Funcs(template.FuncMap{"lower": strings.ToLower}).
		Parse(contentS))


	/*t := template.Must(template.New("issue.html").
		Funcs(template.FuncMap{"lower": strings.ToLower}).
		ParseFiles(indexHTMLPath))*/

	b := new(strings.Builder)
	t.Execute(b, info)

	return b.String()
}

func DetailHTML(article articles.Article, detailPath string) string{
	t := template.Must(template.New("details.html").
		ParseFiles(detailPath))

	b := new(strings.Builder)
	t.Execute(b, article)
	return b.String()
}

func GenerateWebPage(csvPaths []string, wwwPath, staticPath string){
	CopyDirectory(staticPath, wwwPath)

	issuesPath := path.Join(wwwPath, "issues")
	os.Mkdir(issuesPath, 0777)

	detailsPath := path.Join(wwwPath, "details")
	os.Mkdir(detailsPath, 0777)

	issuesInfo := articles.Issues{
		Prefix: "../",
		List:   make([]articles.Issue, len(csvPaths)),
	}

	indicesInfo := articles.Issues{
		Prefix: "",
		List:   make([]articles.Issue, len(csvPaths)),
	}

	for i, v := range csvPaths{
		val := strings.TrimSpace(v)
		val = strings.ReplaceAll(val, ".csv", "")
		parts := strings.Split(val, "_")
		year  := parts[0]
		month := parts[1]
		indicesInfo.List[i].Name = fmt.Sprintf("%v %v", year, strings.Title(month))
		indicesInfo.List[i].FileName = fmt.Sprintf("%v_%v.html", year, strings.ToLower(month))
		indicesInfo.List[i].URL = fmt.Sprintf("issues/%v_%v.html", year, strings.ToLower(month))

		issuesInfo.List[i].Name = fmt.Sprintf("%v %v", year, strings.Title(month))
		issuesInfo.List[i].URL = fmt.Sprintf("%v_%v.html", year, strings.ToLower(month))
		issuesInfo.List[i].FileName = fmt.Sprintf("%v_%v.html", year, strings.ToLower(month))

	}

	for i, csvPath := range csvPaths{
		arts := articles.ParseCSV(csvPath)
		infos := articles.HTMLInfo{
			Articles:     arts,
		}


		if i == 0{
			infos.Issues = indicesInfo
			curIssue := articles.CurrentIssue{
				Name:        issuesInfo.List[0].Name,
				CurrentPast: "Current Issue",
			}
			infos.CurrentIssue = curIssue
			indexStr := IssueHTML("template_files/issue.html", infos)
			indexPath := path.Join(wwwPath, "index.html")
			f, _ := os.Create(indexPath)
			f.WriteString(indexStr)
			f.Close()

			indexCSSPath := filepath.Join("www_static", "index.css")
			f, _ = os.Open(indexCSSPath)
			css, _ := ioutil.ReadAll(f)
			infos.CSS = string(css)
			indexStr = IssueHTML("template_files/index_nku.html", infos)
			indexPath = path.Join(wwwPath, "index_nku.html")
			f, _ = os.Create(indexPath)
			f.WriteString(indexStr)
			f.Close()
		}

		infos.Issues = issuesInfo
		curIssue := articles.CurrentIssue{
			Name:        issuesInfo.List[i].Name,
			CurrentPast: "Past Issue",
		}

		if i == 0{
			curIssue.CurrentPast = "Current Issue"
		}

		infos.CurrentIssue = curIssue

		issueStr := IssueHTML("template_files/issue.html", infos)
		issuePath := path.Join(issuesPath, issuesInfo.List[i].FileName)
		f, _ := os.Create(issuePath)
		f.WriteString(issueStr)
		f.Close()

		for _, article := range arts {
			detailStr := DetailHTML(article, "template_files/details.html")
			detailPath := path.Join(wwwPath, article.DetailsURL)
			f, _ := os.Create(detailPath)
			f.WriteString(detailStr)
			f.Close()
		}
	}

}


