package main

import (
	"context"
	"log"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()

	if err := chromedp.Run(ctx,
		page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorDeny).WithDownloadPath("."),
		chromedp.Navigate(`http://www.mersenne.org/ftp_root/gimps/p95v287.MacOSX.noGUI.tar.gz`),
	); err != nil {
		log.Fatal(err)
	}
}