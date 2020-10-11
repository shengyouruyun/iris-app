package main

import (
	// "google.golang.org/appengine/log"
	// "os"
	//"log"
  // "github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12"
)

// func main() {
// 	app := iris.New()
	
//    // l :=app.Logger()
// 	// Logger middleware by-default will write the logs using app.Logger().Infof.
//     //app.UseRouter(logger.New())
// 	app.Logger().AddOutput(os.Stdout)
//     // Per route middleware, you can add as many as you desire.
//    // app.Get("/benchmark", MyBenchLogger(), benchEndpoint)

//     booksAPI := app.Party("/books")
//     {
//         booksAPI.Use(iris.Compression)

//         // GET: http://localhost:8080/books
//         booksAPI.Get("/", list)
//         // POST: http://localhost:8080/books
//         booksAPI.Post("/", create)
//     }

//     app.Listen(":8080")
// }

// // Book example.
// type Book struct {
//     Title string `json:"title"`
// }

// func list(ctx iris.Context) {
//     books := []Book{
//         {"Mastering Concurrency in Go"},
//         {"Go Design Patterns"},
//         {"Black Hat Go"},
// 	}


//      log.Println(" list of book request")
     
//     ctx.Logger().Info()
	
	 

   
// 	ctx.JSON(books)

//     // TIP: negotiate the response between server's prioritizes
//     // and client's requirements, instead of ctx.JSON:
//     // ctx.Negotiation().JSON().MsgPack().Protobuf()
//     // ctx.Negotiate(books)
// }

// func create(ctx iris.Context) {
//     var b Book
//     err := ctx.ReadJSON(&b)
//     // TIP: use ctx.ReadBody(&b) to bind
//     // any type of incoming data instead.
//     if err != nil {
//         ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
//             Title("Book creation failure").DetailErr(err))
//         // TIP: use ctx.StopWithError(code, err) when only
//         // plain text responses are expected on errors.
//         return
//     }

//     println("Received Book: " + b.Title)

// 	ctx.StatusCode(iris.StatusCreated)
// 	ctx.Writef("Hello %s \n", b.Title)
// }



// func main() {
//     app := iris.Default()
//     // Logging to a file.
//     // Colors are automatically disabled when writing to a file.
//     f, _ := os.Create("iris.log")
//     app.Logger().SetOutput(f)

//     // Use the following code if you need to write the logs
//     // to file and console at the same time.
//     app.Logger().AddOutput(os.Stdout)

//     app.Get("/ping", func(ctx iris.Context) {
// 		ctx.WriteString("pong")
	
// 	})
	
	

//    app.Listen(":8080")
// }


// func main() {
// 	app := iris.New()
	
	
// 	f, _ := os.Create("iris.log")
//      app.Logger().SetOutput(f)
//     reqLogger := logger.New(logger.Config{
//         Status:           true,
//         IP:               true,
//         Method:           true,
//         Path:             true,
//         PathAfterHandler: false,
//         Query:            false,
//         TraceRoute:       true,
//         //Columns:          false,
//        // LogFunc:          customLogFunc,
//         LogFuncCtx:       nil,
//         Skippers:         nil,
//     })
// 	app.UseRouter(reqLogger)

//     app.Get("/ping", func(ctx iris.Context) {
//         ctx.WriteString("pong")
//     })

//     app.Listen(":8080")
// }


func main() {
   app := iris.New()
   //app.Use(recover.New())
   app.Use(logger.New())
   //输出html
   // 请求方式: GET
   // 访问地址: http://localhost:8080/welcome
   app.Handle("GET", "/welcome", func(ctx iris.Context) {
       ctx.HTML("<h1>Welcome</h1>")
   })
   //输出字符串
   // 类似于 app.Handle("GET", "/ping", [...])
   // 请求方式: GET
   // 请求地址: http://localhost:8080/ping
   app.Get("/ping", func(ctx iris.Context) {
       ctx.WriteString("pong")
       app.Logger().Info("get request")
   })
   //输出json
   // 请求方式: GET
   // 请求地址: http://localhost:8080/hello
   app.Get("/hello",Hellofunc)
   app.Run(iris.Addr(":8080"))//8080 监听端口
}


func Hellofunc(ctx iris.Context) {

    app := iris.New()
    ctx.JSON(iris.Map{"message": "Hello Iris!"})

    app.Logger().Info("send hello!")
}