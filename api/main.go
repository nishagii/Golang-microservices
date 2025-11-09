package main

import(
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/dgrijalva/jwt-go"
)

var MySignInKey = []byte(os.Getenv("SECRET_KEY"))

