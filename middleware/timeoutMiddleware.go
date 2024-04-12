package middleware

import (
	"context"
	"net/http"
	"time"
)

func TimeoutMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestTimeout := 5 * time.Second
        ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
        defer cancel() 

        r = r.WithContext(ctx)


        done := make(chan struct{})
        go func ()  {
            defer close(done)
            next.ServeHTTP(w, r) 
        }()

        
          

            select {
            
            case <-ctx.Done():
                if ctx.Err() == context.DeadlineExceeded {
                    w.WriteHeader(http.StatusRequestTimeout)
              
                    w.Write([]byte("Request timed out"))
                }

                
           
            case <-done:
               
                return
            }
        

        
    })
}
