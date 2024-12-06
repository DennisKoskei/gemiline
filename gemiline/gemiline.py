#!/usr/bin/env python3

# This program takes user input, queries the Gemini API..
# and output the result to the terminal


import os
import argparse
import google.generativeai as genai
from dotenv import load_dotenv

# Load the .env file to access the Gemini API key
load_dotenv(".env")
API_KEY = os.getenv("GEMINI_API_KEY")


def is_api_key_available():
    if API_KEY is None:
        print("Error: GEMINI_API_KEY not set")
        return False
    return True


def main():
    # Verify the API key is set
    if not is_api_key_available():
        return

    # Configure the Gemini API KEY
    genai.configure(api_key=API_KEY)

    parser = argparse.ArgumentParser(
        description="Get the user's input and query it through the Gemini api"
    )

    args = parser.add_argument(
        "userInput",
        nargs="+",
        help="Provide input to be queried through the Gemini API",
    )

    args = parser.parse_args()
    # Combine user input into a single string
    user_query = " ".join(args.userInput)

    # Query the Gemini API
    try:
        model = genai.GenerativeModel("gemini-1.5-flash")
        response = model.generate_content(user_query)
        print(response.text)
    except IOError as e:
        print(f"An error occured: {e}")


# Entry point of the script
if __name__ == "__main__":
    main()
