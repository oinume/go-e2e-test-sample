package e2e

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Index)
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
	page, err := webDriver.NewPage()
	if err != nil {
		t.Error(err)
	}
	if err := page.Navigate(serverURL); err != nil {
		t.Error(err)
	}
	if err := page.FindByXPath("//input[@name='name']").Fill("test"); err != nil {
		t.Error(err)
	}
	if err := page.FindByXPath("//input[@type='submit']").Submit(); err != nil {
		t.Error(err)
	}

	text, _ := page.FindByXPath("//p").Text()
	if expected, actual := "Hello test", text; expected != actual {
		t.Errorf("expected = %v, actual = %v", expected, actual)
	}
}
