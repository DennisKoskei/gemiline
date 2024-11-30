#!/usr/bin/env python 3

# 0. Ask for the user api and store it in .env // bypass this initially
# 1. Take in input from the user
# 2. Take the input and request data from the gemini api
# 3. Take the response and output it to the terminal


import google.generativeai as genai 

user_query = input("Hello, How may I help you today? ")

genai.configure(api_key=".env:API_KEY")
model = genai.GenerativeModel("gemini-1.5-flash")
response = model.generate_content(user_query)
print(response.text)j