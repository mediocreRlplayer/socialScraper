# Social Scraper

Social Scraper is a Go program that extracts social media links from a list of websites and updates an Airtable database with the extracted links.

## Prerequisites

Before running the program, make sure you have the following installed:

- Go (version 1.16 or higher)
- Airtable account with an API key

## Installation

1. Clone the repository:

```shell
git clone https://github.com/mediocreRLplayer/social-scraper.git
```

2. Navigate to project directory
  `cd social-scraper`

3. Install required dependencies
  `go mod download`

## Configuration

1. Set up your Airtable database by creating a table with the necessary fields to store website information and social links. Note down the Base ID and Table Name.

2. Create a .env file in the project directory and set the following environment variables: 
```
AIRTABLE_API_KEY=<your_airtable_api_key>
AIRTABLE_BASE_ID=<your_airtable_base_id>
AIRTABLE_TABLE_NAME=<your_airtable_table_name>
```
  - Your base id and table name can be found in the url of the database page
  - Your base id starts with app
  - Table starts with tbl

## Usage

1. Update the main.go file with the desired logic for scraping and processing social links. You can customize the scrapeSocialLinks and isSocialLink functions according to your needs.

2. Build and run the project:
` go run main.go`

  - The program will fetch websites from Airtable, scrape social links from each website, and update the Airtable database with the extracted links.

### Contributing

Contributions are welcome! If you have any suggestions, improvements, or bug fixes, please open an issue or submit a pull request.

### License

* This project is licensed under the MIT License 
