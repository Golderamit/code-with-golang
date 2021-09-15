package main

import(
	"encoding/json"
	"log"
	"net/http"
	"strings"
	
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)
var rnd*renderer.Render
var db*mgo.Database

const (
	hostName       string="localhost:27017"
	dbName         string="demo_todo"
	collectionName string="todo"
	port           string=":8080"
)
type(
	todoModel struct{
		 ID        bson.ObjectId`bson:"_id,omitempty"`
		 Title     string `bson:"title"`
		 Completed bool `bson:"completed"`
		 CreatedAt time.Time `bson:"createAt"`
	}
	todo struct {
         ID        string    `json:"id"`
		 Title     string    `json:"title"`
		 Completed bool      `json:"completed"`
		 CreatedAt time.Time `json:"created_at"`
	}
)
func init(){
	rnd=renderer.New()
	sess, err:=mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic,true)
	db=sess.DB(dbName)
}
func checkErr(err error){
         if err!=nil{
			 log.Fatal(err)
		 }
}
func homeHandler(w http.ResponseWriter,r *http.Request){
	err:=rnd.Template(w,http.StatusOK,[]string{"static/home.tpl"},nil)
	checkErr(err)
}
func fetchTodos(w http.ResponseWriter,r *http.Request){
     todos:=[]todoModel{}

	 if err:=db.C(collectionName).Find(bson.M{}).All(&todos);err!=nil{
		 rnd.JSON(w ,http.StatusProcessing,renderer.M{
			 "message":"Failed to fetch todo",
			 "error":err,
		 })
		 return
	 }
	 todolist :=[]todo{}

	 for _,t:=range todos{
		 todolist=append(todolist, todo{
			 ID:t.ID.Hex(),
			 Title: t.Title,
			 Completed: t.Completed,
			 CreatedAt: t.CreatedAt,
		 })
	 }
	 rnd.JSON(w,http.StatusOK,renderer.M{
		 "data":todolist,
	 })
}
func createTodo(w http.ResponseWriter,r *http.Request){
	var t todo 
	if err :=json.NewDecoder(r.Body).Decode(&t);err!=nil{
		rnd.JSON(w,http.StatusProcessing,err)
		return
	}
	// simple validation
	if t.Title == ""{
	rnd.JSON(w, http.StatusBadRequest,renderer.M{
		"message":"the title field is requried",
	})
	return
   }
 // if input is ok,create a todo
 tm := todoModel{
	 ID:       bson.NewObjectId(),
	 Title:    t.Title,
	 Completed: false,
	 CreatedAt: time.Now(),
 }
  if err := db.C(collectionName).Insert(&tm); err != nil{
	  rnd.JSON(w,http.StatusProcessing,renderer.M{
		  "message":"Failed to save todo ",
		  "error": err,
	  })
	  return
  }
   rnd.JSON(w, http.StatusCreated,renderer.M{
	   "message":"todo created successfully",
	   "todo_id":tm.ID.Hex(),
   })

}
func deleteTodo(w http.ResponseWriter,r *http.Request){
	id :=strings.TrimSpace(chi.URLParam(r,"id"))

	if !bson.IsObjectIdHex(id){
	  rnd.JSON(w, http.StatusBadRequest,renderer.M{
		  "message": "the id is invalid",
	  })
	  return	
	}

	if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil  {
          rnd.JSON(w, http.StatusProcessing, renderer.M{
			  "message": "failed to delete todo",
			  "error": err,
		  })
		  return
	}
	rnd.JSON(w, http.StatusOK,renderer.M{
		"message" : "Todo deleted successfully",
	})
}
func updateTodo(w http.ResponseWriter,r *http.Request){
	id:=strings.TrimSpace(chi.URLParam(r,"id"))

	if !bson.IsObjectIdHex(id){
		rnd.JSON(w, http.StatusBadRequest,renderer.M{
			"message": "The is is invalid",
		})
		return
	}
	var t todo
	if err:= json.NewDecoder(r.Body).Decode(&t);err != nil{
		rnd.JSON(w, http.StatusProcessing,err)
		return
	}
	if t.Title== ""{
	rnd.JSON(w,http.StatusBadRequest,renderer.M{
		"message": "the title field id required",
	})
	return
  }
  if err:= db.C(collectionName).
  Update(
	  bson.M{"_id":bson.ObjectIdHex(id)},
	  bson.M{"title":t.Title, "completed":t.Completed},
  );err != nil{
	rnd.JSON(w, http.StatusProcessing, renderer.M{
		"message": "Failed to update todo",
		"error":   err,
  })
  return
}
rnd.JSON(w, http.StatusOK, renderer.M{
	"message": "Todo updated successfully",
})
}
func main(){
	
	r :=chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/",homeHandler)
	r.Mount("/todo",todoHandlers())

	srv := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout: 15*time.Second,
		WriteTimeout: 15*time.Second,
		IdleTimeout: 15*time.Second,

	}
	go func ()  {
		log.Println("listening on port ", port)
		if err :=srv.ListenAndServe();err !=nil{
			log.Printf("listen:%s\n",err)
		}
		
	}()
	

}
func todoHandlers()http.Handler{
	rg:=chi.NewRouter()
	rg.Group(func(r chi.Router){
		r.Get("/",fetchTodos)
		r.Post("/",createTodo)
		r.Put("/{id}",updateTodo)
		r.Delete("/{id}",deleteTodo)
	})
	return rg
}