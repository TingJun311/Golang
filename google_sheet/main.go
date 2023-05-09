package main

import (
    "context"
    "fmt"
    "log"

    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)

func main() {
    // Replace the sheet URL with your own
    sheetURL := "https://docs.google.com/spreadsheets/d/1I-lOoI0Qu7D0V5wxyLfA1vwUQYtAG7MGM6OYxtW9zf8/edit#gid=0"

    // Create a new Sheets API client
    ctx := context.Background()
    service, err := sheets.NewService(ctx, option.WithoutAuthentication())
    if err != nil {
        log.Fatalf("Unable to retrieve Sheets client: %v", err)
    }

    // Extract the spreadsheet ID from the URL
    spreadsheetID := extractSpreadsheetID(sheetURL)

    // Get the sheet data
    readRange := "Sheet1" // Change to the sheet name you want to retrieve data from
    response, err := service.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
    if err != nil {
        log.Fatalf("Unable to retrieve data from sheet: %v", err)
    }

    // Print the data
    if len(response.Values) == 0 {
        fmt.Println("No data found.")
    } else {
        for _, row := range response.Values {
            fmt.Printf("%s\n", row)
        }
    }
}

// extractSpreadsheetID extracts the spreadsheet ID from the sheet URL
func extractSpreadsheetID(sheetURL string) string {
    start := len("https://docs.google.com/spreadsheets/d/")
    end := len(sheetURL)
    for i := start; i < end; i++ {
        if sheetURL[i] == '/' {
            end = i
            break
        }
    }
    return sheetURL[start:end]
}