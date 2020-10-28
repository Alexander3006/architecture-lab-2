package lab2

import (
  "fmt"
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
  cases := []struct {
    in, want string
  }{
    {"17", "17"},
    {"a b *", "* a b"},
    {"і ї + г /", "/ + і ї г"},
    {"3 8 + 11 /", "/ + 3 8 11"},
    {"lab1 lab2 +", "+ lab1 lab2"},
    {"4 2 - 3 * 5 +", "+ * - 4 2 3 5"},
    {"15 7 8 3 ^ / - 12 +", "+ - 15 / 7 ^ 8 3 12"},
    {"10 45 + 88 77 - 2 ^ /", "/ + 10 45 ^ - 88 77 2"},
    {"5 6 4 10 8 9 8 1 + ^ / / - * +", "+ 5 * 6 - 4 / 10 / 8 ^ 9 + 8 1"},
    {"5210 1000 + 2922 ^ 9200 1934 * 6666 ^ - 8971 /", "/ - ^ + 5210 1000 2922 ^ * 9200 1934 6666 8971"},
    {"9227 1002 ^ 2929 / 2688 1000 + - 15208 8000 ^ / 9999 +", "+ / - / ^ 9227 1002 2929 + 2688 1000 ^ 15208 8000 9999"},
  }

  for _, c := range cases {
    res, err := PostfixToPrefix(c.in)
    if assert.Nil(t, err) {
      assert.Equal(t, c.want, res)
    }
  }

  casesWrong := []string{
    "",                            //empty
    "+",                           //unary operator
    "# $ @ & ?",                   //wrong characters
    "28*7+(7-4)^10",               //infix
    "+ * - 4 2 3 5",               //postfix
    "\xe1\xe0\xe1\xe0",            //unicode
    "23 83 1221 4444 93332 + * /", //wrong prefix
    "https://www.youtube.com/",    //link
  }

  for _, c := range casesWrong {
    _, err := PostfixToPrefix(c)
    if !assert.NotNil(t, err) {
      t.Fatal("Not Nil")
    }
  }
}

func ExamplePostfixToPrefix() {
  res, _ := PostfixToPrefix("2 2 +")
  fmt.Println(res)
}
