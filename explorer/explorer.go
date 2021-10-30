package explorer

import (
	"fmt"
	"github.com/pudgekim/gocoin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const (
	// main.go는 explorer 밖에 있으므로 explorer/ 를 붙여줘야함
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	// 대문자,소문자로 인한 public/private은 template에도 영향을 끼침
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(w http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.BlockChain().AllBlocks()}
	templates.ExecuteTemplate(w, "home", data)
}

func add(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates.ExecuteTemplate(w, "add", nil)
	case http.MethodPost:
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.BlockChain().AddBlock(data)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
	templates.ExecuteTemplate(w, "add", nil)
}

func Start(port int) {
	handler := http.NewServeMux()
	// 여기서 pages 하위에 있는 것들을 glob함
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	// 위에서 만들어진 templates에 추가로 partials 하위 파일들을 glob
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)
	fmt.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
