package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// data format
// binary: 불안정
// string 권장
// - XML: tag, 사이즈가 커짐
// - JSON(JavaScript Object Notation): key-value
// annotation: json 형태를 알려주어 decode 시에 활용
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

// fooHandler의 메소드
// Handler 인터페이스의 ServeHTTP 메소드 구현
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)

	// JSON 문자열을 GO value(User type data)로 변경
	// url이 아닌 body로 넘기는 방식
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}

	// user 생성 시간
	user.CreatedAt = time.Now()

	// json으로 encoding
	data, _ := json.Marshal(user)
	// content-type을 지정하여 json 형태로 출력되도록 함
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// 웹 브라우저가 웹 서버에 요청을 보낼 때 url을 통해 input을 넘길 수 있다
	// URL query string으로부터 파라미터 값 가져오기
	// query string: url 주소에 사전에 협의된 data를 파라미터를 통해 넘기는 것
	// 정해진 엔드포인트 주소 이후에 ?를 쓰는 것
	// =로 key, value 구분, parameter=value
	// &로 여러 파라미터 넘길 수 있다
	// ex:) /about?name=takityaki >> Hello, takityaki!
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// http.handler는 interface로 SeverHttp function을 가짐
// mux가 해당 interface를 상속함
func NewHttpHandler() http.Handler {
	// router
	// 경로에 따라 다르게 분배해주는 기능
	// mux(multiplexer): 입력에 따라 다른 모듈로 분기해주는 것
	// 웹 request에 따라 다른 핸들러로 연결해주는 역할
	// NewServeMux: HTTP 요청 멀티플렉서 인스턴스 생성
	// http에 직접 넘기지 않고 router를 통해 넘기기
	mux := http.NewServeMux()

	// 1
	// url별로 요청을 처리할 핸들러 함수 등록
	// "/" 경로(index page)로 접속했을 때 처리할 핸들러 함수 지정
	// http.ResponseWriter 타입에 값을 쓰면 HTTP 응답으로 전송
	// http.Request: 사용자가 요청한 정보
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// "Hello, Wolrd!"를 화면에 출력
		// Fprint: 지정한 스트림에 출력
		fmt.Fprintln(w, "Hello, World!")
	})

	// 1-1
	// 함수를 따로 정의하고 인자로 넘기기
	// 경로("/about")에 따른 핸들러
	mux.HandleFunc("/about", AboutHandler)

	// 2
	// function이 아닌 instance를 등록
	mux.Handle("/foo", &fooHandler{})

	return mux

}
