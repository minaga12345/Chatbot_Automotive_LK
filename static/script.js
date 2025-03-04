// Function to load initial greeting message and buttons
async function loadInitialGreeting() {
    const response = await fetch('/initChat'); // Fetch initial chat data from the server
    const data = await response.json(); // Parse the JSON response
    renderInitialGreeting(data); // Render the greeting message and buttons
}

// Function to render initial greeting and buttons
function renderInitialGreeting(data) {
    const chatWindow = document.getElementById('chatWindow');

    // Display the greeting message
    const greetingBubble = document.createElement('div');
    greetingBubble.classList.add('message', 'chatbot-message');
    greetingBubble.innerHTML = `
        <img src="static/chatbot-avatar.jpg" alt="Chatbot Avatar" class="avatar chatbot-avatar">
        ${data.reply}
    `;
    chatWindow.appendChild(greetingBubble); // Append the greeting message to the chat window

    // Render buttons if available
    if (data.buttons && data.buttons.length > 0) {
        renderButtons(data.buttons); // Call renderButtons to create button elements
    }
}

// Function to send message
async function sendMessage() {
    const userMessage = document.getElementById('userMessage').value;
    if (userMessage.trim() === "") return;

    const chatWindow = document.getElementById('chatWindow');

    // Create and display the user's message bubble on the right
    const userMessageBubble = document.createElement('div');
    userMessageBubble.classList.add('message', 'user-message');
    userMessageBubble.innerHTML = `
        <img src="static/user-avatar.png" alt="User Avatar" class="avatar user-avatar">
        ${userMessage}
    `;
    chatWindow.appendChild(userMessageBubble);

    const response = await fetch('/chat', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ message: userMessage })
    });

    const data = await response.json();

    // Create and display the chatbot's response bubble on the left
    const chatbotMessageBubble = document.createElement('div');
    chatbotMessageBubble.classList.add('message', 'chatbot-message');
    chatbotMessageBubble.innerHTML = `
        <img src="static/chatbot-avatar.jpg" alt="Chatbot Avatar" class="avatar chatbot-avatar">
        ${data.reply}
    `;
    chatWindow.appendChild(chatbotMessageBubble);

    // Auto-scroll to the bottom of the chat window
    chatWindow.scrollTop = chatWindow.scrollHeight;

    // Clear the input field
    document.getElementById('userMessage').value = '';

    // Render buttons if needed
    if (data.buttons && data.buttons.length > 0) {
        renderButtons(data.buttons);
    }
}

// Handle Enter key press for sending message
function handleKeyDown(event) {
    if (event.key === "Enter") {
        event.preventDefault();
        sendMessage();
    }
}

// Function to render buttons for vehicle types or locations
function renderButtons(buttons) {
    const chatWindow = document.getElementById('chatWindow');
    const buttonContainer = document.createElement('div');
    buttonContainer.classList.add('button-container');

    buttons.forEach(button => {
        const btn = document.createElement('button');
        btn.innerText = button.text;
        btn.classList.add('button');
        btn.onclick = () => {
            document.getElementById('userMessage').value = button.value;
            sendMessage(); // Trigger message sending with the button value
        };
        buttonContainer.appendChild(btn);
    });

    chatWindow.appendChild(buttonContainer);
    chatWindow.scrollTop = chatWindow.scrollHeight; // Scroll to the bottom after adding buttons
}

// Function to handle button clicks
function buttonClicked(buttonValue) {
    document.getElementById('userMessage').value = buttonValue;
    sendMessage();
}

// Function to clear chat history
function clearChat() {
    const chatWindow = document.getElementById('chatWindow');
    chatWindow.innerHTML = ''; // Clear all messages
}

// Function to minimize or maximize chat window
function toggleChatWindow() {
    const chatbotContainer = document.getElementById('chatbotContainer');
    chatbotContainer.classList.toggle('minimized'); // Toggle the minimized state
}

// Call the function to load the greeting on page load
window.onload = loadInitialGreeting; // Load the initial greeting when the window is loaded
