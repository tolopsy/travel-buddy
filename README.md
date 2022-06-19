# ✈️🛍 TRAVEL BUDDY

Travel Buddy is a simple API that suggests places to visit when you travel to a new place.

## Setting things up ⚙️

To set things up for use in Travel Buddy, Be sure to have Go already configured on your local machine. Also, you need to create an API Key in Google Console with access to the **Google Places API**.
After which you do the following👇🏾

 - Clone repo to your local machine
 - Create a .env file in the project directory and in your .env file assign your API Key to GOOGLE_PLACES_API_KEY variable.
	 `GOOGLE_PLACES_API_KEY=<YourAPIKey>` 
	 
	 Note: You should assign your API key without quotation marks around it. Let's assume your API key is V5tYuYOei9U
	 
	 Don't do `GOOGLE_PLACES_API_KEY="V5tYuYOei9U"` ❌
     Do `GOOGLE_PLACES_API_KEY=V5tYuYOei9U` ✅

## Running travel-buddy
To run travel-buddy, 

 - Confirm in your terminal that the current working directory is travel-buddy project directory.
 - Run `go mod tidy`
 - Run 👉🏾 `go build -o travel-app ./cmd/travel`
This will create an executable with the name "travel-app" in your project directory.
 - Run the travel-app executable with `./travel-app` in your terminal.
 - Travel Buddy API is made available at port 9000 and provides two endpoints - `journeys` and `recommendations`
 
	 ### The `journeys` endpoint
	 This provides the kind of places you can visit on different occasions. For example, If you're going on a romantic trip, then you can visit a movie theatre, restaurant, park etc. This is pretty much what this endpoint makes you know🤷🏾‍♂️.
	 You can access this by visiting `http://localhost:9000/journeys` in your browser.


	### The `recommendations` endpoint
	This provides suggestions for the kind of place a user wants to visit in a location.
	A typical example of API call to this endpoint is `http://localhost:9000/recommendations?lat=40.7127840&lng=-74.0059410&radius=5000&journey=cafe|bar|jewelry_store|restaurant&cost=$...$$$$`
	Take a look at the query parameters
	
	 - lat and lng: represents latitude and longitude (respectively) of the place user is located or wants to visit
	 - radius: represents the radius in meters around the location represented by `lng` and `lat` within the user can visit
	 - journey: represents the different kinds of places that user is open to visit. In the case of the given example, the user would be getting suggestions of cafes, bars, jewellery stores or restaurants within 5000m radius of their given location.
	 - cost: represents the expense level of the visit. `$...$$$$` represents the price range. Maximum cost level (i.e most expensive) is represented by five dollar signs `$$$$$` while most affordable would have one dollar sign $.


## Other things you should know

 - There is a dedicated log directory named `logs`. In an event when something goes wrong or no suggestion is provided for a particular kind of place, it will be logged into `logs/main.log` and also in your terminal.

	 
