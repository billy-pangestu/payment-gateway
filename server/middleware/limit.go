package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"payment-gateway-backend/usecase"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// LimitOTPInit ...
type LimitOTPInit struct {
	*usecase.ContractUC
	MaxLimit       float64
	Duration       string
	MaxLimitIDUser float64
}

// Create a map to hold the rate limiters for each visitor and a mutex.
var visitors = make(map[string]*rate.Limiter)
var mtx sync.Mutex

// Create a new rate limiter and add it to the visitors map, using the
// IP address as the key.
func addVisitor(ip string) *rate.Limiter {
	limiter := rate.NewLimiter(1, 1)
	mtx.Lock()
	visitors[ip] = limiter
	mtx.Unlock()
	return limiter
}

// Retrieve and return the rate limiter for the current visitor if it
// already exists. Otherwise call the addVisitor function to add a
// new entry to the map.
func getVisitor(ip string) *rate.Limiter {
	fmt.Println(ip)
	mtx.Lock()
	limiter, exists := visitors[ip]
	mtx.Unlock()
	if !exists {
		return addVisitor(ip)
	}
	return limiter
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the getVisitor function to retreive the rate limiter for
		// the current user.
		fmt.Println(r.Method)
		fmt.Println(r.URL.Path)
		limiter := getVisitor(r.RemoteAddr)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// LimitByIP ...
func (li LimitOTPInit) LimitByIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dur, err := time.ParseDuration(li.Duration)
		if err != nil {
			http.Error(w, http.StatusText(422), http.StatusInternalServerError)
			return
		}

		key := r.RemoteAddr + "-" + r.Method + "-" + r.URL.Path
		res, err := li.Redis.Get(key).Result()
		if err != nil {
			li.Redis.Set(key, 1, dur)
			next.ServeHTTP(w, r)
			return
		}
		if res == "" {
			li.Redis.Set(key, 1, dur)
			next.ServeHTTP(w, r)
			return
		}
		var cb interface{}
		err = json.Unmarshal([]byte(res), &cb)
		if err != nil {
			li.Redis.Set(key, 1, dur)
			next.ServeHTTP(w, r)
			return
		}
		counter := cb.(float64)
		if counter >= li.MaxLimit {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		li.Redis.Set(key, counter+1, dur)

		next.ServeHTTP(w, r)
	})
}

func requestIDFromContextInterface(ctx context.Context, key string) map[string]interface{} {
	return ctx.Value(key).(map[string]interface{})
}
