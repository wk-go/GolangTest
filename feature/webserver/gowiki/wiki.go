package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title	string
	Body 	string
}

func (p *Page)save() error{
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func loadPage(title string)(*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title:title, Body:string(body)}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w,"<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p,err := loadPage(title)
	if err != nil {
		p = &Page{Title:title}
	}
	fmt.Fprintf(w,"<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/save/"):]
	body :=r.FormValue("body")
	p := &Page{Title: title, Body: body}
	err := p.save()
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func main(){
	p1 := &Page{Title: "test", Body:"This is a sample Page"}
	p1.save()

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}