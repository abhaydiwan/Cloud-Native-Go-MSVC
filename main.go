package main
import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main(){
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context){
		c.String(http.StatusOK, "pong")
	})

	engine.GET("/hello", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message":"Hello Gin framework."})
	})

    //get all books
	engine.GET("/api/books", func(c *gin.Context){
		c.JSON(http.StatusOK, AllBooks())
	})

	//create a new book
	engine.POST("/api/books", func(c *gin.Context){
		var book Book
		if c.BindJSON(&book) == nil{
			isbn,created:= CreateBook(book)
			if created {
				c.Header("Location","/api/books"+isbn)
				c.Status(http.StatusCreated)
			}else{
				c.Status(http.StatusConflict)
			   }
			}
		})

	//create a new book
	engine.GET("/api/books/:isbn", func(c *gin.Context){
			isbn := c.Params.ByName("isbn")
			book,found := GetBook(isbn)
			if found {
				c.JSON(http.StatusOK, book)
			}else{
				c.AbortWithStatus(http.StatusNotFound)
			}
		})
//update existing book
engine.PUT("/api/books/:isbn", func(c *gin.Context){
	isbn := c.Params.ByName("isbn")

	var book Book
	if c.BindJSON(&book) == nil{
	   exists := UpdateBook(isbn,book)
	   if exists {
            c.Status(http.StatusOK)
	   }else{
            c.Status(http.StatusNotFound)
	   }
	}
})

//delete book
engine.DELETE("/api/books/:isbn", func(c *gin.Context){
	isbn := c.Params.ByName("isbn")
     DeleteBook(isbn)
	 c.Status(http.StatusOK)
})

//run on port
	engine.Run(port())
}

func port() string{
  port:= os.Getenv("port")
  if len(port)==0{
	  port ="8090"
  }
  return ":"+port

}