package middleware

import (
	"context"
	"net/http"
	"time"
)

func TimeoutMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestTimeout := 10 * time.Second
        ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
        defer cancel() 

        r = r.WithContext(ctx)

        done := make(chan struct{})
        defer close(done) 

        go func() {
          
            defer cancel() 

            select {
            
            case <-time.After(requestTimeout):

                w.WriteHeader(http.StatusRequestTimeout)
              
                w.Write([]byte("Request timed out"))
           
            case <-done:
               
                return
            }
        }()

        next.ServeHTTP(w, r)
    })
}
