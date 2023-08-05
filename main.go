package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Availability struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type BusinessHour struct {
	Id         string `json:"id"`
	ResourceId string `json:"resource_id"`
	Quantity   int64  `json:"quantity"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type BlockHour struct {
	Id         string `json:"id"`
	ResourceId string `json:"resource_id"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type Appointment struct {
	Id         string `json:"id"`
	ResourceId string `json:"resource_id"`
	Quantity   int64  `json:"quantity"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type ListBusinessHoursRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListBlockHoursRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListAppointmentRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

// API Call for given endpoints

func getDurations(c *gin.Context) {
	url := "http://api.internship.appointy.com:8000/v1/durations"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjMwMDF9.8pZMhoqZdBLqOKT0V7perD4vkoA347idSHVLaCcdefs")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func getResources(c *gin.Context) {
	url := "http://api.internship.appointy.com:8000/v1/resources"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjMwMDF9.8pZMhoqZdBLqOKT0V7perD4vkoA347idSHVLaCcdefs")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func getBusinessHours(resourceId string, startTime string, endTime string) (string, error) {
	url := "http://api.internship.appointy.com:8000/v1/business-hours?resourceId=" + resourceId + "&startTime=" + startTime + "&endTime=" + endTime
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjMwMDF9.8pZMhoqZdBLqOKT0V7perD4vkoA347idSHVLaCcdefs")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body), nil
}
func getBlockHours(resourceId string, startTime string, endTime string) (string, error) {
	url := "http://api.internship.appointy.com:8000/v1/block-hours?resourceId=" + resourceId + "&startTime=" + startTime + "&endTime=" + endTime
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjMwMDF9.8pZMhoqZdBLqOKT0V7perD4vkoA347idSHVLaCcdefs")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body), nil
}

func getAppointments(resourceId string, startTime string, endTime string) (string, error) {
	url := "http://api.internship.appointy.com:8000/v1/appointments?resourceId=" + resourceId + "&startTime=" + startTime + "&endTime=" + endTime
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjMwMDF9.8pZMhoqZdBLqOKT0V7perD4vkoA347idSHVLaCcdefs")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body), nil
}

// -------------------------------------------

// Main API calculation

func calculation(businessHours string, blockHours string, appointments string, date string, duration string, quantity string) (string, error) {
	var businessHour []BusinessHour
	var blockHour []BlockHour
	var appointment []Appointment
	quantityInt, err := strconv.ParseInt(quantity, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	var availability []Availability

	errr := json.Unmarshal([]byte(businessHours), &businessHour)
	if errr != nil {
		fmt.Println(errr)
		return "", errr
	}
	err = json.Unmarshal([]byte(blockHours), &blockHour)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = json.Unmarshal([]byte(appointments), &appointment)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for i := 0; i < len(businessHour); i++ {
		businessStartTime := businessHour[i].StartTime
		businessEndTime := businessHour[i].EndTime
		availableQuantity := businessHour[i].Quantity

		for j := 0; j < len(blockHour); j++ {
			blockStartTime := blockHour[j].StartTime
			blockEndTime := blockHour[j].EndTime

			// avoiding block timings

			if businessStartTime <= blockStartTime && businessEndTime >= blockEndTime {
				if businessStartTime != blockStartTime && availableQuantity >= quantityInt {
					availability = append(availability, Availability{businessStartTime, blockStartTime})
				}
				if blockEndTime != businessEndTime && availableQuantity >= quantityInt {
					availability = append(availability, Availability{blockEndTime, businessEndTime})
				}
			}
			if businessStartTime <= blockStartTime && businessEndTime >= blockStartTime && businessEndTime <= blockEndTime {
				if businessStartTime != blockStartTime && availableQuantity >= quantityInt {
					availability = append(availability, Availability{businessStartTime, blockStartTime})
				}
			}
			if businessStartTime >= blockStartTime && businessStartTime <= blockEndTime && businessEndTime >= blockEndTime {
				if blockEndTime != businessEndTime && availableQuantity >= quantityInt {
					availability = append(availability, Availability{blockEndTime, businessEndTime})
				}
			}
		}

		for j := 0; j < len(appointment); j++ {
			appointmentStartTime := appointment[j].StartTime
			appointmentEndTime := appointment[j].EndTime

			// avoiding appointment timings

			if businessStartTime <= appointmentStartTime && businessEndTime >= appointmentEndTime {
				if businessStartTime != appointmentStartTime && availableQuantity >= quantityInt {
					availability = append(availability, Availability{businessStartTime, appointmentStartTime})
				}
				if appointmentEndTime != businessEndTime && availableQuantity >= quantityInt {
					availability = append(availability, Availability{appointmentEndTime, businessEndTime})
				}
			}
			if businessStartTime <= appointmentStartTime && businessEndTime >= appointmentStartTime && businessEndTime <= appointmentEndTime {
				if businessStartTime != appointmentStartTime && availableQuantity >= quantityInt {
					availability = append(availability, Availability{businessStartTime, appointmentStartTime})
				}
			}
			if businessStartTime >= appointmentStartTime && businessStartTime <= appointmentEndTime && businessEndTime >= appointmentEndTime {
				if appointmentEndTime != businessEndTime && availableQuantity >= quantityInt {
					availability = append(availability, Availability{appointmentEndTime, businessEndTime})
				}
			}
		}
	}

	availabilityJson, err := json.Marshal(availability)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(availabilityJson), nil
}

func getAvailability(resourceId string, duration string, date string, quantity string) (string, error) {
	startTime := date + "T00:00:00Z"
	endTime := date + "T23:59:59Z"
	businessHours, err := getBusinessHours(resourceId, startTime, endTime)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	blockHours, err := getBlockHours(resourceId, startTime, endTime)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	appointments, err := getAppointments(resourceId, startTime, endTime)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	Availability, err := calculation(businessHours, blockHours, appointments, date, duration, quantity)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return Availability, nil
}

func getAvailabilityApi(c *gin.Context) {
	resourceId := c.Query("resourceId")
	date := c.Query("date")
	duration := c.Query("duration")
	quantity := c.Query("quantity")

	Availability, err := getAvailability(resourceId, duration, date, quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(Availability)
	c.JSON(http.StatusOK, Availability)
}

func main() {
	router := gin.Default()
	router.GET("/v1/availability", getAvailabilityApi)
	router.Run("localhost:8080")
}