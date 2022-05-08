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
	Latitude         float32 `json:"latitude" bson:"latitude,omitempty"  faker:"lat"`
	Longitude        float32 `faker:"long"`
	CreditCardNumber string  `faker:"cc_number"`
	CreditCardType   string  `faker:"cc_type"`
	Email            string  `json:"email"  bson:"email,omitempty"  faker:"email"`
	IPV4             string  `faker:"ipv4"`
	IPV6             string  `faker:"ipv6"`
	Password         string  `faker:"password"`
	PhoneNumber      string  `faker:"phone_number"`
	MacAddress       string  `faker:"mac_address"`
	URL              string  `faker:"url"`
	UserName         string  `faker:"username"`
	//ToolFreeNumber   string  `faker:"tool_free_number"` // @fixme
	E164PhoneNumber    string             `faker:"e_164_phone_number"`
	TitleMale          string             `faker:"title_male"`
	TitleFemale        string             `faker:"title_female"`
	FirstName          string             `faker:"first_name"`
	FirstNameMale      string             `faker:"first_name_male"`
	FirstNameFemale    string             `faker:"first_name_female"`
	LastName           string             `faker:"last_name"`
	Name               string             `faker:"name"`
	UnixTime           int64              `faker:"unix_time"`
	Date               string             `faker:"date"`
	Time               string             `faker:"time"`
	MonthName          string             `faker:"month_name"`
	Year               string             `faker:"year"`
	DayOfWeek          string             `faker:"day_of_week"`
	DayOfMonth         string             `faker:"day_of_month"`
	Timestamp          string             `faker:"timestamp"`
	Century            string             `faker:"century"`
	TimeZone           string             `faker:"timezone"`
	TimePeriod         string             `faker:"time_period"`
	Word               string             `faker:"word"`
	Sentence           string             `faker:"sentence"`
	Paragraph          string             `faker:"paragraph"`
	Currency           string             `faker:"currency"`
	Amount             float64            `faker:"amount"`
	AmountWithCurrency string             `faker:"amount_with_currency"`
	UUIDHypenated      string             `faker:"uuid_hyphenated"`
	UUID               string             `faker:"uuid_digit"`
	ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty" faker:"customIdFaker"`
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
