package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post("http://www.tmubei.com", map[string]string{"contents":"another fakeed immoc.com",})
	return s.Get("http://www.tmubei.com")
}

func download(r real.Retriever) string {
	return r.Get("http://www.tmubei.com")
}

func main(){
	fmt.Println(("----------"))
	var r Retriever

	retriever := &mock.Retriever{"this is a fake imooc.com"}
	r = retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "mozilla/5.0",
		Timeout: time.Minute,
	}
	inspect(r)

	//Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	}else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))

	fmt.Println("try a seesion")
	fmt.Println(session(retriever))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Content:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}