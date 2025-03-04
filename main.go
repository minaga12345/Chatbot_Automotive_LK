package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Reply   string   `json:"reply"`
	Buttons []Button `json:"buttons,omitempty"`
}

type Button struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

// Default response when the user first visits
var defaultResponse = Response{
	Reply: "Hi, I'm Automotive LK's Virtual Assistant. \n\n" +
		"I can assist you with your queries with regard to our vehicle products and services. \n" +
		"You may select from the following options or type your specific requirement.",
	Buttons: []Button{
		{Text: "View Car Models", Value: "car models"},
		{Text: "Locate a Branch", Value: "branches"},
		{Text: "Current Promotions", Value: "promotions"},
		{Text: "Services & Maintenance", Value: "service"},
		{Text: "Make a Payment", Value: "online payment"},
	},
}

// Map of predefined responses
var responses = map[string]Response{
	"hi": {
		Reply: "Hi! How can I assist you today?",
	},
	"hello": {
		Reply: "Hello! How can I assist you today?",
	},
	"hey": {
		Reply: "Hey! How can I assist you today?",
	},
	"how are you": {
		Reply: "I'm Automotive LK's Virtual Assistant, here to assist you with Automotive-related queries!",
	},
	"what is your name, who are you, name": {
		Reply: "I'm the Automotive LK's Virtual Assistant! How can I assist you?",
	},
	"contact, helpline, contact number, contact information, how can I contact you": {
		Reply: `
        <div style='line-height: 1.5;'>
            <strong>Tel:</strong> <a href='tel:+94112939000'>+9411 293 9000 - 6</a><br>
            <strong>Fax:</strong> <a href='tel:+94112939005'>+9411 293 9005</a><br>
            <strong>Email:</strong> <a href='mailto:info@automotive.lk'>info@automotive.lk</a>
        </div>
    `,
	},
	"Email": {
		Reply: "Email Us: <a href='mailto:info@automotive.lk'>info@automotive.lk</a>",
	},
	"opening hours, what are your opening hours, showrooms open time": {
		Reply: "Our showrooms are open from 8:30 AM to 5:30 PM, Monday to Saturday.",
	},
	"main office, where is your main office located, headquarters location, address": {
		Reply: "Our main office is located at 337, Negombo Road, Wattala 11300, Sri Lanka.",
	},
	"about": {
		Reply: "Automotive LK is committed to quality, innovation, and customer satisfaction. More details can be found on our website: <a href='https://www.toyota.lk/about' target='_blank'>About Automotive LK</a>.",
	},
	"promotions, promotion": {
		Reply: "Automotive LK offers several promotions and discounts for members and staff at several organizations. You can check out for more promotion details on our website: <a href='https://www.toyota.lk/promotions' target='_blank'>Promotions Automotive LK</a>.",
	},
	"news, newsroom, news room, article, articles, latest": {
		Reply: "Check out the Latest Events and Occasions at the News Room. You can check out for news on our website: <a href='https://www.toyota.lk/newsroom' target='_blank'>News Room Automotive LK</a>.",
	},
	"branch, showroom, service center, dealerships, dealership, branches": {
		Reply: "Locate your nearest branch using our website. You can locate your nearest branch on our website: <a href='https://www.toyota.lk/locate-a-branch' target='_blank'>Locate a branch</a>.",
	},
	"online payment, payment": {
		Reply: "Do your payments using our online payment facility through the website: <a href='https://payments.toyota.lk/' target='_blank'>Payment Portal</a>.",
	},
	"service, services, maintenance, repair": {
		Reply: "Maintenance in servicing and repairs through Toyota Lanka. Check out for General Repairs and Periodic Maintenance: <a href='https://www.toyota.lk/maintenance-services/' target='_blank'>Repair and Maintenance</a>.",
	},
	"safety, safety sense": {
		Reply: "Automotive LK has a long history of creating advancements and innovations, striving to produce vehicles that are safe and reliable: <a href='https://www.toyota.lk/toyota-safety-sense-2/' target='_blank'>Automotive LK Safety Sense</a>.",
	},

	"car models, models, cars, car, suv, pickup, bus, sedan, hatchback, vehicles, vehicle": {
		Reply: "Please select a car model:",
		Buttons: []Button{
			{Text: "Compact SUV", Value: "compact suv"},
			{Text: "SUV", Value: "suv"},
			{Text: "Pickup & Bus", Value: "pickup & bus"},
			{Text: "Sedan", Value: "sedan"},
			{Text: "Hatchback", Value: "hatchback"},
		},
	},
	"locations, location": {
		Reply: "Please select a location:",
		Buttons: []Button{
			{Text: "Automotive LK Colombo", Value: "colombo"},
			{Text: "Automotive LK Kandy", Value: "kandy"},
			{Text: "Automotive LK Galle", Value: "galle"},
		},
	},
}

// Map of vehicle details
var vehicleDetails = map[string]Response{
	"compact suv": {
		Reply: "The Compact SUV is perfect for city driving and offers great fuel efficiency. Would you like more details? Please select a model:",
		Buttons: []Button{
			{Text: "Toyota RAV4", Value: "rav4"},
			{Text: "Toyota Rush", Value: "rush"},
		},
	},
	"suv": {
		Reply: "Our SUVs are designed for adventure and comfort. Interested in a specific model?:",
		Buttons: []Button{
			{Text: "Toyota Fortuner", Value: "fortuner"},
			{Text: "Toyota Land Cruiser", Value: "land cruiser"},
		},
	},
	"pickup & bus": {
		Reply: "You have selected Pickup & Bus. Please select a model:",
		Buttons: []Button{
			{Text: "Toyota Hilux", Value: "hilux"},
			{Text: "Toyota Coaster", Value: "coaster"},
		},
	},
	"sedan": {
		Reply: "You have selected Sedan. Please select a model:",
		Buttons: []Button{
			{Text: "Toyota Camry", Value: "camry"},
			{Text: "Toyota Corolla", Value: "corolla"},
		},
	},
	"hatchback": {
		Reply: "You have selected Hatchback. Please select a model:",
		Buttons: []Button{
			{Text: "Toyota Yaris", Value: "yaris"},
			{Text: "Toyota Prius", Value: "prius"},
		},
	},
}

// Map of vehicle descriptions
var vehicleDescriptions = map[string]Response{
	"rav4": {
		Reply: "The Toyota RAV4 is a versatile compact SUV that offers an adventurous spirit with comfort and style. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/rav4' target='_blank'>Toyota RAV4</a>.",
	},
	"rush": {
		Reply: "The Toyota Rush is a dynamic compact family SUV with 7 seats that combines bold styling with excellent fuel efficiency. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/rush' target='_blank'>Toyota Rush</a>.",
	},
	"fortuner": {
		Reply: "The Toyota Fortuner is a tough and reliable SUV built for both urban and off-road adventures. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/fortuner' target='_blank'>Toyota Fortuner</a>.",
	},
	"land cruiser": {
		Reply: "The Toyota Land Cruiser is a legendary SUV known for its off-road capabilities and luxury features. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/land-cruiser' target='_blank'>Toyota Land Cruiser</a>.",
	},
	"hilux": {
		Reply: "The Toyota Hilux is a rugged pickup designed for heavy-duty performance and durability. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/hilux' target='_blank'>Toyota Hilux</a>.",
	},
	"coaster": {
		Reply: "The Toyota Coaster is an ideal bus for transporting passengers with comfort and reliability. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/coaster' target='_blank'>Toyota Coaster</a>.",
	},
	"camry": {
		Reply: "The Toyota Camry is a premium sedan that offers an upscale experience with advanced safety features. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/camry' target='_blank'>Toyota Camry</a>.",
	},
	"corolla": {
		Reply: "The Toyota Corolla is a compact sedan that stands out for its reliability and efficiency. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/corolla' target='_blank'>Toyota Corolla</a>.",
	},
	"yaris": {
		Reply: "The Toyota Yaris is a compact hatchback that offers great maneuverability and efficiency. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/yaris' target='_blank'>Toyota Yaris</a>.",
	},
	"prius": {
		Reply: "The Toyota Prius is a hybrid hatchback that prioritizes eco-friendliness and efficiency. For more information, visit: <a href='https://www.toyota.lk/vehicle-model/prius' target='_blank'>Toyota Prius</a>.",
	},
}

// Connect to the database
func connectDB() (*sql.DB, error) {
	dsn := "root:mysql@tcp(localhost:3306)/toyota_chatbot"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Error while opening the database connection: %v\n", err)
		return nil, err
	}

	// Check if the database connection is alive
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging the database: %v\n", err)
		return nil, err
	}

	fmt.Println("Database connection successful!")
	return db, nil
}

func handleUserQuery(userMessage string, db *sql.DB) Response {
	// Normalize user input
	userMessage = strings.ToLower(strings.TrimSpace(userMessage))
	log.Println("Handling user query:", userMessage)

	// Check if the userMessage matches a vehicle category first
	if detailsResponse, exists := vehicleDetails[userMessage]; exists {
		log.Println("Found vehicle details for:", userMessage)
		return detailsResponse
	}

	// Check if the userMessage matches a specific vehicle model
	if descriptionResponse, exists := vehicleDescriptions[userMessage]; exists {
		log.Println("Found vehicle description for:", userMessage)
		return descriptionResponse
	}

	// Check predefined responses
	predefinedResponse := getChatResponse(userMessage, db)
	if predefinedResponse.Reply != "I'm sorry, I didn't understand that. Can you please ask me something else?" {
		log.Println("Found predefined response for:", userMessage)
		return predefinedResponse
	}

	// Query the database for a response based on the keyword
	dbResponse, err := getResponseByKeyword(db, userMessage)
	if err == nil && dbResponse.Reply != "" {
		log.Println("Database response found for:", userMessage)
		return dbResponse
	}

	log.Println("No response found for:", userMessage) // Log unmatched query for future analysis
	return Response{Reply: "I'm sorry, I didn't understand that. Can you please ask me something else?"}
}

func getResponseByKeyword(db *sql.DB, keyword string) (Response, error) {
	var reply string
	var buttons sql.NullString // Use sql.NullString to handle possible NULL values

	log.Println("Attempting to retrieve response for keyword:", keyword)

	// Match against the keyword in the 'responses' table first
	query := "SELECT reply, buttons FROM responses WHERE keywords LIKE ?"
	err := db.QueryRow(query, "%"+keyword+"%").Scan(&reply, &buttons)
	if err == nil {
		// If a match is found in the 'responses' table, process the buttons if present
		buttonList := []Button{}
		if buttons.Valid && buttons.String != "" {
			err = json.Unmarshal([]byte(buttons.String), &buttonList)
			if err != nil {
				log.Println("JSON unmarshal error:", err)
				return Response{}, err
			}
		}
		log.Println("Database reply found in 'responses':", reply)
		return Response{Reply: reply, Buttons: buttonList}, nil
	}

	// If no match in 'responses', check in the 'vehicle_models' table
	var modelName, description, link string
	vehicleQuery := "SELECT model_name, description, link FROM vehicle_models WHERE model_name LIKE ?"
	err = db.QueryRow(vehicleQuery, "%"+keyword+"%").Scan(&modelName, &description, &link)
	if err == nil {
		// If a vehicle model is found, format the response
		log.Println("Database vehicle model found:", modelName)
		vehicleReply := fmt.Sprintf("%s For more information, visit: <a href='%s' target='_blank'>%s</a>.", description, link, modelName)
		return Response{Reply: vehicleReply}, nil
	}

	// Log if no match is found in either table
	log.Println("No response found in 'responses' or 'vehicle_models' for keyword:", keyword)
	return Response{}, err
}

func matchesKeyword(userInput, keyword string) bool {
	words := strings.Fields(userInput) // Split user input into words
	for _, word := range words {
		// Check for exact match
		if word == keyword {
			return true
		}
		// Check plural/singular match
		if word == keyword+"s" || (strings.HasSuffix(keyword, "s") && word == keyword[:len(keyword)-1]) {
			return true
		}
	}
	return false
}

func getChatResponse(userInput string, db *sql.DB) Response {
	userInput = strings.ToLower(strings.TrimSpace(userInput))
	log.Println("Checking predefined responses for:", userInput)

	// Split user input into words for better matching
	userWords := strings.Fields(userInput)

	// Iterate over predefined responses and look for a match
	for keywords, response := range responses {
		// Split the keywords string into individual keywords
		for _, keyword := range strings.Split(keywords, ", ") {
			for _, word := range userWords {
				// If any word in the user input matches a keyword, return the response
				if strings.EqualFold(word, keyword) {
					log.Println("Keyword match found for:", keyword)
					return response
				}
			}
		}
	}

	log.Println("No predefined response matched for:", userInput)

	// Check vehicle categories based on keywords in the user's message
	for category, details := range vehicleDetails {
		if strings.Contains(userInput, category) {
			log.Println("Vehicle category detected:", userInput)
			return details
		}
	}

	// Check vehicle models based on keywords in the user's message
	for model, description := range vehicleDescriptions {
		if strings.Contains(userInput, model) {
			log.Println("Vehicle model detected:", userInput)
			return description
		}
	}

	// Check for location responses
	switch userInput {
	case "colombo":
		log.Println("Location 'colombo' detected")
		return Response{Reply: "Automotive LK Colombo is located here: <a href='https://maps.app.goo.gl/tQfaZPHEuykjMoCAA' target='_blank'>Google Map Link</a>."}
	case "kandy":
		log.Println("Location 'kandy' detected")
		return Response{Reply: "Automotive LK Kandy is located here: <a href='https://maps.app.goo.gl/2UnfFbQEaEzMnCrn6' target='_blank'>Google Map Link</a>."}
	case "galle":
		log.Println("Location 'galle' detected")
		return Response{Reply: "Automotive LK Galle is located here: <a href='https://maps.app.goo.gl/8NUrpvjeCzf48fhx7' target='_blank'>Google Map Link</a>."}
	}

	log.Println("No location matched for:", userInput)

	// If no response was found, return a default message
	return Response{Reply: "I'm sorry, I didn't understand that. Can you please ask me something else?"}
}

// Chat handler for POST requests
func chatbotHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var input map[string]string
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userMessage := strings.TrimSpace(strings.ToLower(input["message"]))
		fmt.Println("Received message:", userMessage)

		db, err := connectDB()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// Now use handleUserQuery to process the user's message
		botResponse := handleUserQuery(userMessage, db)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(botResponse)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// New handler for GET requests to initialize the chat with a greeting
func initChatHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(defaultResponse)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/chat", chatbotHandler)
	http.HandleFunc("/initChat", initChatHandler)
	fmt.Println("Server is running on port 8081...")
	http.ListenAndServe(":8081", nil)
}
