# NINE PRO

This is a proof of concept to rebuild the [NINE PRO](https://ninepro.webflow.io/) website using Go and Gin.

### Initialise Go for application
```
// via project directory
go mod init ninepro_go

// via GitHub
go mod init gtihub.com/PythonAkoto/ninepro_go
```

### Install required packages
```
go get -u github.com/gin-gonic/gin
go get gopkg.in/mail.v2
go get github.com/joho/godotenv
```

### When deploying the app
You need to build the Go application to produce an executable by using the following commands:
```
go build -o ninepro_go ./cmd/main.go
```

Now you can run the binary:
```
./ninepro_go    // On Unix/Linux/Mac
ninepro_go.exe  // On Windows
```

If you are deploying on Heroku you will need to create a `Procfile` and add the following:
```
web: ./ninepro_go
```

**By default, Heroku detects Go applications and can automatically set up a web process.**


**When deploying a Go application, Render can automatically detect it and build the executable.**
**You just need to specify the start command in Render's dashboard:**
```
./ninepro_go
```