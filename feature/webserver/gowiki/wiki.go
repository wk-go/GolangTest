package main

import(
	"io/ioutil"
	"net/http"
	"html/template"
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	t,_ := template.ParseFiles(tmpl +".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil{
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p,err := loadPage(title)
	if err != nil {
		p = &Page{Title:title}
	}
	renderTemplate(w, "edit", p)
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
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}