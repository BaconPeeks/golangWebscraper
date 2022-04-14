package main

//make imports for scraping data

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"github.com/TwiN/go-color"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: false,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

//make the main function for the scraper
func main() {
	//get the url
	url := getURL()
	//scrape the data
	data, err := scrapeData(url)
	if err != nil {
		log.Panic(err)
	}
	//store the data
	storeData(data)
}

//ask the user for the url
func getURL() string {
	//ask the user for the url
	fmt.Println(color.Bold + "Please enter the url: " + color.Reset)
	//make a scanner to read the input
	scanner := bufio.NewScanner(os.Stdin)
	//read the input
	scanner.Scan()
	//return the input
	return scanner.Text()
}

//make a function to scrape dB leaks from the web
func scrapeData(url string) (string, error) {
	//make a get request to the url
	resp, err := http.Get(url)
	fmt.Println(color.Blue+"Scraping data from:", url+color.Reset)
	if err != nil {
		return "", err
		log.Panic("Error getting data from url!")
	}
	//close the response
	defer resp.Body.Close()
	//read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
		log.Panic("Error reading response body!")
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(color.Yellow+"Scraped data from:", url+color.Reset)
	//return the body
	return string(body), nil
}

//cleanup the data before storing it
func cleanScrapedData(data string) string {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(color.Green + "Cleaning data from url!" + color.Reset)
	//make a regex to find the data
	re := regexp.MustCompile(`<td class="col-md-2">(.*?)</td>`)
	//make a slice to store the data
	var cleanData []string
	//find the data
	matches := re.FindAllStringSubmatch(data, -1)
	//loop through the data
	for _, match := range matches {
		//add the data to the slice
		cleanData = append(cleanData, match[1])
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Println(color.Purple + "Cleaned data from url!" + color.Reset)
	//return the data
	return strings.Join(cleanData, ",")
}

//store the data in a txt
func storeData(data string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(color.Cyan + "Storing data from scraped url!" + color.Reset)
	//open the file
	file, err := os.OpenFile("dbLeaks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		log.Error("Error opening file!")
	}

	time.Sleep(500 * time.Millisecond)
	fmt.Println(color.Bold + "Stored data from scraped url!" + color.Reset)
	//close the file
	defer file.Close()
	if _, err := file.WriteString(data); err != nil {
		fmt.Println(err)
		log.Panic("I couldnt write to the file!\n I will now exit!")
	}


    //wait till they press a key to exit
	time.Sleep(500 * time.Millisecond)
	fmt.Println(color.Green + "Press any key to exit!" + color.Reset)
	//make a scanner to read the input for enter
	scanner := bufio.NewScanner(os.Stdin)
	//read the input
	scanner.Scan()
	//add some sleep time
	time.Sleep(500 * time.Millisecond)
	//tell them Bye
	fmt.Println(color.Green + "Bye!" + color.Reset)
	//exit the program
	os.Exit(0)
}
