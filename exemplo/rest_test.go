package exemplo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

var expectedArticles = []Article{
	{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	//{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

type MockArticle struct {
	mock.Mock
}

func (mock *MockArticle) getArticle() []Article {
	args := mock.Called()
	result := args.Get(0)
	return result.([]Article)
}

func TestRest(t *testing.T) {

	req, err := http.NewRequest("GET", "/article", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	mockArticle := new(MockArticle)
	mockArticle.On("getArticle").Return(expectedArticles)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllArticles)
	handler.ServeHTTP(rr, req)

	mockArticle.AssertExpectations(t)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	responseArticle := []Article{}
	json.NewDecoder(rr.Body).Decode(&responseArticle)

	fmt.Println(len(responseArticle))
	for counter, article := range responseArticle {
		if article != expectedArticles[counter] {
			t.Errorf("Return values not equal: got %v want %v",
				article, expectedArticles[counter])
		}
	}
}
