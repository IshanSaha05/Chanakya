package scraper

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

/*
func GetHTMLContentAssemblyState(url string) (string, error) {
	var cancelFuncs []context.CancelFunc

	chromedpContext := context.Background()

	ctx, cancel := chromedp.NewExecAllocator(chromedpContext, chromedp.Flag("headless", false), chromedp.WindowSize(1920, 1080))
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = context.WithTimeout(ctx, time.Second*50)
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = chromedp.NewContext(ctx)
	cancelFuncs = append(cancelFuncs, cancel)

	cancelAll := func() {
		for _, cancel := range cancelFuncs {
			cancel()
		}
	}

	var htmlContent string

	var options []string

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady(`body`),
		chromedp.Click(`#stateSelect`, chromedp.ByID),
		chromedp.Evaluate(`Array.from(document.querySelectorAll('#state-select option'))
		.map(opt => {
			try {
				return JSON.parse(opt.value).stateName;
			} catch (e) {
				return null; // Handle the empty or invalid JSON option
			}
		})
		.filter(stateName => stateName !== null);`, &options),
	)

	cancelAll()

	if err != nil {
		return "", err
	}

	fmt.Println(options)

	chromedpContext = context.Background()

	ctx, cancel = chromedp.NewExecAllocator(chromedpContext, chromedp.Flag("headless", false), chromedp.WindowSize(1920, 1080))
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = context.WithTimeout(ctx, time.Second*50)
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = chromedp.NewContext(ctx)
	cancelFuncs = append(cancelFuncs, cancel)

	for _, stateName := range options {
		fmt.Println(stateName)
		err = chromedp.Run(
			ctx,
			chromedp.Navigate(url),
			chromedp.WaitReady(`body`),
			chromedp.Click(`#stateSelect`, chromedp.ByID),
			chromedp.SetValue(`#stateSelect`, fmt.Sprintf(`{"stateName":"%s"}`, stateName)),
			chromedp.Evaluate(`document.querySelector('#stateSelect').dispatchEvent(new Event('change'))`, nil),
			chromedp.Sleep(2*time.Second),
		)

		if err != nil {
			return "", err
		}
	}

	return htmlContent, nil
}
*/
/*
func GetHTMLContentAssemblyState(url string) (string, error) {
	var cancelFuncs []context.CancelFunc

	chromedpContext := context.Background()

	ctx, cancel := chromedp.NewExecAllocator(chromedpContext, chromedp.Flag("headless", false), chromedp.WindowSize(1920, 1080))
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = context.WithTimeout(ctx, time.Second*50)
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = chromedp.NewContext(ctx)
	cancelFuncs = append(cancelFuncs, cancel)

	cancelAll := func() {
		for _, cancel := range cancelFuncs {
			cancel()
		}
	}
	defer cancelAll()

	// Navigate to the page
	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		log.Fatal(err)
	}

	// Wait for the dropdown to load
	if err := chromedp.Run(ctx, chromedp.WaitReady(`body`)); err != nil {
		log.Fatal(err)
	}

	optionValues := []string{`{"stateName":"AndhraPradesh","stateDisplayName":"Andhra Pradesh","isNewData":true,"hasStrongWeakBoothInfo":true}`,
		`{"stateName":"Assam","stateDisplayName":"Assam","isNewData":true,"hasStrongWeakBoothInfo":true}`}

	var htmlContent string

	for _, value := range optionValues {

		if err := chromedp.Run(ctx,
			chromedp.Click(`#stateSelect`, chromedp.ByID),
			chromedp.Sleep(10*time.Second),
			chromedp.Click(`#state-select option[ng-repeat="stateData in allStates.states"][value='{"stateName":"AndhraPradesh","stateDisplayName":"Andhra Pradesh","isNewData":true,"hasStrongWealBoothInfo":true}']`),
			chromedp.Sleep(10*time.Second), // Wait for the page to load after selecting the option
		); err != nil {
			fmt.Println("Failed to select option:", err)
			return "", nil
		}

		// Here you can add code to process the loaded page after selecting the option
		// For example, print the HTML content
		var htmlContent string
		if err := chromedp.Run(ctx,
			chromedp.OuterHTML("html", &htmlContent),
		); err != nil {
			fmt.Println("Failed to get HTML content:", err)
			return "", nil
		}
		fmt.Println("HTML content after selecting", value, ":", htmlContent)
	}

	return htmlContent, nil
}
*/
func GetHTMLContentAssemblyState(url string) (string, error) {
	var cancelFuncs []context.CancelFunc

	chromedpContext := context.Background()

	ctx, cancel := chromedp.NewExecAllocator(chromedpContext, chromedp.Flag("headless", false), chromedp.WindowSize(1920, 1080))
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = context.WithTimeout(ctx, time.Second*50)
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = chromedp.NewContext(ctx)
	cancelFuncs = append(cancelFuncs, cancel)

	cancelAll := func() {
		for _, cancel := range cancelFuncs {
			cancel()
		}
	}
	defer cancelAll()

	// Navigate to the page
	if err := chromedp.Run(ctx, chromedp.Navigate(url)); err != nil {
		log.Fatal(err)
	}

	var htmlContent string
	if err := chromedp.Run(ctx,
		chromedp.WaitReady(`body`, chromedp.ByQuery),
		chromedp.WaitVisible(`#stateSelect`, chromedp.ByID),
		chromedp.Click(`#stateSelect`, chromedp.ByID),
		chromedp.WaitVisible(`#state-select > option:nth-child(2)`, chromedp.ByQuery),
		chromedp.Click(`#state-select > option:nth-child(2)`, chromedp.ByQuery),
		chromedp.Sleep(3*time.Second), // Wait for the page to load after selecting the option
	); err != nil {
		fmt.Println("Failed to select option:", err)
		return "", nil

	}

	return htmlContent, nil
}
