package main

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()
	var statusCode int64
	var responseHeaders network.Headers
	url := `http://www.mersenne.org/ftp_root/gimps/p95v287.MacOSX.noGUI.tar.gz`
	chromedp.ListenTarget(ctx, func(event interface{}) {
		log.Println("ListenTarget")
		switch responseReceivedEvent := event.(type) {
		case *network.EventResponseReceived:
			response := responseReceivedEvent.Response
			if response.URL == url {
				statusCode = response.Status
				responseHeaders = response.Headers
			}
			log.Println("EventResponseReceived")
		case *network.EventRequestWillBeSent:
			request := responseReceivedEvent.Request
			log.Println("chromedp is requesting url (could be in background): %s\n", request.URL)
			if responseReceivedEvent.RedirectResponse != nil {
				url = request.URL
				log.Println(" got redirect: %s\n", responseReceivedEvent.RedirectResponse.URL)
			}
		case *page.EventDownloadProgress:
			log.Println("page.EventDownloadProgress: TotalBytes: %+v, State: %+v",
				responseReceivedEvent.TotalBytes, responseReceivedEvent.State)
		}
	})
	if err := chromedp.Run(ctx,
		network.Enable(),
		network.SetExtraHTTPHeaders(map[string]interface{}{"User-Agent": "Mozilla/5.0"}),
		page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorDeny).WithDownloadPath("."),
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("ActionFunc url: ", url)
			return nil
		}),
	); err != nil {
		log.Fatal(err)
	}
}