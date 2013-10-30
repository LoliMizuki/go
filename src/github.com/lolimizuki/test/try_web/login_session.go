package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// 	"time"
// )

// var globalSessions *Manager
// var providers = make(map[string]Provider)

// func init() {
// 	globalSessions, err := NewManager("memmory", "gosessionid", 3600)
// 	go globalSessions.GC()

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func loginWithSession(w http.ResponseWriter, r *http.Request) {
// 	if globalSessions == nil {
// 		fmt.Println("globalSessions is nil")
// 	}

// 	session := globalSessions.SessionStart(w, r)
// 	r.ParseForm()
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("login_session.gtpl")
// 		w.Header().Set("Content-Type", "text/html")
// 		t.Execute(w, session)
// 	} else {
// 		session.Set("username", r.Form["username"])
// 		http.Redirect(w, r, "/", 302)
// 	}
// }

// func count(w http.ResponseWriter, r *http.Request) {
// 	sess := globalSessions.SessionStart(w, r)
// 	createtime := sess.Get("createtime")
// 	if createtime == nil {
// 		sess.Set("createtime", time.Now().Unix())
// 	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
// 		globalSessions.SessionDestroy(w, r)
// 		sess = globalSessions.SessionStart(w, r)
// 	}
// 	ct := sess.Get("countnum")
// 	if ct == nil {
// 		sess.Set("countnum", 1)
// 	} else {
// 		sess.Set("countnum", (ct.(int) + 1))
// 	}
// 	t, _ := template.ParseFiles("count.gtpl")
// 	w.Header().Set("Content-Type", "text/html")
// 	t.Execute(w, sess.Get("countnum"))
// }

// func Register(name string, provider Provider) {
// 	if provider == nil {
// 		panic("session: Register provide is nil")
// 	}
// 	if _, dup := providers[name]; dup {
// 		panic("session: Register called twice for provide " + name)
// 	}
// 	providers[name] = provider
// }

// type Session interface {
// 	Set(key, value interface{}) error // set session value
// 	Get(key interface{}) interface{}  // get session value
// 	Delete(key interface{}) error     // delete session value
// 	SessionID() string                // back current sessionID
// }

// type Provider interface {
// 	SessionInit(sid string) (Session, error)
// 	SessionRead(sid string) (Session, error)
// 	SessionDestroy(sid string) error
// 	SessionGC(maxLifeTime int64)
// }
