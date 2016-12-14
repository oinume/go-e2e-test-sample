package e2e

import (
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/oinume/go-e2e-test-sample/app"
	"github.com/sclevine/agouti"
)

var (
	serverURL string
	webDriver *agouti.WebDriver
)

func TestMain(m *testing.M) {
	if err := testMain(m); err != nil {
		log.Fatalf("err = %v\n", err)
	}
}

func testMain(m *testing.M) error {
	// Setup HTTP server
	mux := app.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()
	serverURL = server.URL

	// Setup WebDriver
	webDriver = agouti.ChromeDriver()
	if err := webDriver.Start(); err != nil {
		return err
	}
	defer webDriver.Stop()

	// Run tests
	if status := m.Run(); status != 0 {
		return fmt.Errorf("Run status = %v", status)
	}

	return nil
}

func TestIndex(t *testing.T) {
	fmt.Printf("serverURL = %v\n", serverURL)
	page, err := webDriver.NewPage()
	if err != nil {
		t.Error(err)
	}
	if err := page.Navigate(serverURL); err != nil {
		t.Error(err)
	}
}

func newWebDriver() *agouti.WebDriver {
	envWebDriver := os.Getenv("WEB_DRIVER")
	var driver *agouti.WebDriver
	switch strings.ToLower(envWebDriver) {
	case "phantomjs":
		driver = agouti.PhantomJS()
	default:
		driver = agouti.ChromeDriver()
	}
	//driver.HTTPClient = client
	return driver
}
