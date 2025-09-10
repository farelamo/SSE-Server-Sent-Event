package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Adjust as needed
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))

	r.GET("/progress", func(c *gin.Context) {
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		c.Writer.Flush()

		progressChan := make(chan int)
		closeNotify := c.Writer.CloseNotify()

		go func() {
			for i := 0; i <= 100; i++ {
				time.Sleep(100 * time.Millisecond)
				fmt.Println("progress:", i)
				progressChan <- i
			}
			/*
				ni anak kalo ga ada, si looping progressChan (baris bawah nya)
				akan nunggu forever ampe di close cuy.. ngabisin memory blokk

			*/
			close(progressChan)
		}()

		// 	//kirim2 lah cuyy ke client nichh
		// 	for progress := range progressChan {
		// 		fmt.Fprintf(c.Writer, "data: %d\n\n", progress)

		// 		// flush ni dia ngirim apapun itu data buffer langsung ga pake babibubebo ke client
		// 		c.Writer.Flush()
		// 	}

		for {
			select {
			case progress, ok := <-progressChan:
				if !ok {
					return // Channel closed, exit the loop
				}
				fmt.Fprintf(c.Writer, "data: %d\n\n", progress)
				c.Writer.Flush()
			case <-closeNotify:
				fmt.Println("Client disconnected")
				return
			}
		}
	})

	r.Run(":8080")
}
