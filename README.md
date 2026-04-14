Description

ASCII Art Web is a Go-based web application that converts user input text into ASCII art using predefined banner templates. It provides a simple web interface where users can generate styled text using different banner formats.

The project demonstrates HTTP server handling, form processing, file-based data loading, and HTML template rendering in Go.

## Features
- Convert text into ASCII art via web interface
- Supports three banner styles: standard, shadow, thinkertoy
- Optional "all banners" mode to display all outputs together
- Simple error handling with user feedback
- Preserves ASCII formatting using HTML <pre> rendering

## Project Structure
- main.go – Initializes server and routes
- handlers/ – HTTP request handling logic
- ascii/ – ASCII generation logic and banner processing
- printArt/ – Core ASCII rendering engine
- banners/ – ASCII character templates
- templates/ – Frontend HTML template

## How It Works
- User submits text and selects a banner via web form
- Server receives a POST request at /ascii
- Input is validated (text and banner required)
- ASCII generator loads corresponding banner file
- Text is converted into ASCII art
- Result is rendered back into the HTML page

If “all banners” is selected, the system generates output for all available banner styles and groups them in a single response.

## How to Run
go run main.go

Then open:

http://localhost:8080

## HTTP Endpoints

GET /

Renders the homepage with input form

POST /ascii

Processes text and returns generated ASCII art

Form fields:

text (string)
banner (standard | shadow | thinkertoy | all)

## Error Handling
400 Bad Request: Missing input or invalid form submission
500 Internal Server Error: Banner loading or processing failure

Error messages are displayed directly on the UI.

## Implementation Notes
- Built using Go’s standard net/http package
- HTML rendering handled via Go html/template
- ASCII generation uses file-based character maps
- Output formatting preserved using <pre> tags

## Requirements Compliance (Zone01)
- Uses only standard Go libraries
- Implements required HTTP endpoints
- Handles form input correctly via POST
- Displays results dynamically using templates
- Includes proper error handling with HTTP status codes
- Maintains modular and clean project structure

## Author

Developed as part of a Go web programming exercise focused on backend fundamentals, HTTP handling, and text processing.
