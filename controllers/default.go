package controllers

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/fixwa/go-faker-api/fakery"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type FakeStruct struct {
	Latitude         float32 `json:"latitude" bson:"latitude,omitempty" faker:"lat"`
	Longitude        float32 `json:"longitude" faker:"long"`
	CreditCardNumber string  `json:"creditCardNumber" faker:"cc_number"`
	CreditCardType   string  `json:"creditCardType" faker:"cc_type"`
	Email            string  `json:"email"  bson:"email,omitempty"  faker:"email"`
	IPV4             string  `json:"IPV4" faker:"ipv4"`
	IPV6             string  `json:"IPV6" faker:"ipv6"`
	Password         string  `json:"password" faker:"password"`
	PhoneNumber      string  `json:"phoneNumber" faker:"phone_number"`
	MacAddress       string  `json:"macAddress" faker:"mac_address"`
	URL              string  `json:"url" faker:"url"`
	UserName         string  `json:"userName" faker:"username"`
	//ToolFreeNumber   string  `faker:"tool_free_number"` // @fixme
	E164PhoneNumber    string             `json:"e164PhoneNumber" faker:"e_164_phone_number"`
	TitleMale          string             `json:"titleMale" faker:"title_male"`
	TitleFemale        string             `json:"titleFemale" faker:"title_female"`
	FirstName          string             `json:"firstName" faker:"first_name"`
	FirstNameMale      string             `json:"firstNameMale" faker:"first_name_male"`
	FirstNameFemale    string             `json:"firstNameFemale" faker:"first_name_female"`
	LastName           string             `json:"lastName" faker:"last_name"`
	Name               string             `json:"name" faker:"name"`
	UnixTime           int64              `json:"unixTime" faker:"unix_time"`
	Date               string             `json:"date" faker:"date"`
	Time               string             `json:"time" faker:"time"`
	MonthName          string             `json:"monthName" faker:"month_name"`
	Year               string             `json:"year" faker:"year"`
	DayOfWeek          string             `json:"dayOfWeek" faker:"day_of_week"`
	DayOfMonth         string             `json:"dayOfMonth" faker:"day_of_month"`
	Timestamp          string             `json:"timestamp" faker:"timestamp"`
	Century            string             `json:"century" faker:"century"`
	TimeZone           string             `json:"timeZone" faker:"timezone"`
	TimePeriod         string             `json:"timePeriod" faker:"time_period"`
	Word               string             `json:"word" faker:"word"`
	Sentence           string             `json:"sentence" faker:"sentence"`
	Paragraph          string             `json:"paragraph" faker:"paragraph"`
	Currency           string             `json:"currency" faker:"currency"`
	Amount             float64            `json:"amount" faker:"amount"`
	AmountWithCurrency string             `json:"amountWithCurrency" faker:"amount_with_currency"`
	UUIDHypenated      string             `json:"uuidHypenated" faker:"uuid_hyphenated"`
	UUID               string             `json:"uuid" faker:"uuid_digit"`
	ID                 primitive.ObjectID `json:"id" json:"id" bson:"_id,omitempty" faker:"customIdFaker"`
}

func GenerateAllFaked(c *gin.Context) {
	var a FakeStruct
	fakery.Generator()
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, a)
}
