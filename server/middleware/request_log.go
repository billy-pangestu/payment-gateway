package middleware

import (
	"net/http"
)

// RequestLoggerMiddleware ...
func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// var (
		// 	err     error
		// 	payload []byte
		// )

		if r.Method == "GET" || r.Method == "HEAD" {
			next.ServeHTTP(w, r)
			return
		}

		if r.Method != "GET" {
			// type requestLoggerEntity struct {
			// 	id        bson.ObjectId `bson:"_id,omitempty"`
			// 	Method    string        `bson:"method"`
			// 	Payload   string        `bson:"payload"`
			// 	CreatedAt time.Time     `bson:"created_at"`
			// }

			// payload, err = ioutil.ReadAll(r.Body)
			// if err != nil {
			// 	apiHandler.SendBadRequest(w, err.Error())
			// 	return
			// }

			// rlE := requestLoggerEntity{
			// 	Method:    r.Method,
			// 	Payload:   string(payload),
			// 	CreatedAt: time.Now(),
			// }

			// ds := mw.DB.Copy()
			// defer ds.Close()
			// err = ds.DB(mw.Cfg.GetString("database.mongo.db")).C("request_logger").Insert(&rlE)
			// if err != nil {
			// 	apiHandler.SendBadRequest(w, err.Error())
			// 	return
			// }

			// r.Body = ioutil.NopCloser(bytes.NewBuffer(payload))

			// // Create a response wrapper:
			// mrw := &rw{
			// 	ResponseWriter: w,
			// 	buf:            &bytes.Buffer{},
			// }

			// Call next handler, passing the response wrapper:
			// next.ServeHTTP(mrw, r)
			next.ServeHTTP(w, r)
		}
	})
}
