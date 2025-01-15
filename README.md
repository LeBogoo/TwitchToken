# Twitch Token Generator

This project is a simple web application to generate OAuth tokens for Twitch. It includes a frontend to select the type of token and a backend to serve the necessary files and provide the client ID and redirect URI to the frontend.

## Features

- Generate tokens for different purposes/scopes (Chat Bot, Redemption Bot, Combined)
- Tokens are displayed securely and are not stored or processed by the server
- All token handling happens on the frontend (in the browser)
- Simple and minimalistic design

## Prerequisites

- Docker
- Go 1.23 or higher

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/lebogoo/twitchtoken.git
   cd twitchtoken
   ```

2. Set your Twitch Client ID and Redirect URI as environment variables and run the application:

   ```sh
   export CLIENT_ID=your_twitch_client_id
   export REDIRECT_URI=your_redirect_uri
   go run .
   ```

3. Open your browser and navigate to `http://localhost:8080`.

## Usage

1. Select the type of token you want to generate on the homepage.
2. Authorize the application on Twitch.
3. Copy the generated token displayed on the redirect page.

## File Structure

- `public/`: Contains the frontend HTML, CSS, and JavaScript files.
- `main.go`: The main Go application file.
- `go.mod`: Go module dependencies.
- `Dockerfile`: Dockerfile to build the application.

## License

This project is licensed under the GNU General Public License (GPL). See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.
