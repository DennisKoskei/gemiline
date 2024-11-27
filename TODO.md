# Gemiline - Breakdown

This is a simple python program that allows the user to query gemini directly from the terminal and get quick responses from there.
It is intended to act as a simple search "tool" which prevents the user from switching to a broswer to ask a simple question or "How to" steps.


# Purpose of this Project
1. Allow the user to access the power of Gemini directly from the terminal
2. Boost Efficiency: It is faster to open another terminal and type your query instad of starting up a browser just to search for the same and then go back to your terminal

# Benefits of this program
1. Save time when querying for information during a deep work session i.e coding.
2. Boost Productivity by preventing switching to a browser to search something you "might" have forgotten or want to know and unintentionally get sucked into distractions.


# Steps to Achieve this
## What: 
Write the initial program to allow the program to access the Gemini API and output information.
## How: 
1. Structure of the program
	1. The program should:
		- Take in input from the user
		- Take the input and request data from the gemini api 
		- Take the response and output it to the terminal

The program should allow the following mock example

```powershell
PS; C:\Users\User1\ > gemiline What is the height of the Eiffel Tower?
gemiline : The height of the eiffel tower is ...
```


# My Own Prerequisites
1. Learn how to write proper Industry standard python programs that will allow scalability and security
2. Learn how to interact with the Gemini api.
3. What is the industry standard used to store and access api keys for accessing Gemini api.


---

I still have a debate on how I want the user to interact with the program:

1. METHOD 1:


```Powershell
PS C:\Users\Dev001> gemiline What is the height of the Eiffel Tower?
----->
gemiline: "The height of the Eiffel Tower is 1010 meters"

# For this interface method, the user will type the program name and the query that he has.
# Then the program will take the input and process it through the API then output the results to the terminal.
```

2. METHOD 2:

```Powershell
PS C:\Users\Dev001> gemiline
----->
gemiline:> Hello, How may I help you?
You: What is the height of the Eiffel Tower?
gemiline: "The height of the Eiffel Tower is blah blah blah meters"

# For this Interface method, the user will invoke the program name.
# Then the program will start and the first output from gemiline will display itself "Hello, How may I help you?
# Then the cursor will place itself after the statement "You: " where you will type in your input
# Then after typing your query and pressing enter the program will take your input and process it before outputting the results.
```