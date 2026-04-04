# Gemiline - A CLI based gemini query program

## Elevator Pitch

Imagine if you could use **Gemini AI on the command line** for **quick queries.**

No need of opening up a browser, typing in the AI website and typing in your query - and getting distracted while you are at it.

**Target Audience**: Programmers who spend most time on the terminal and need a **_quick query_** of something **_without exiting the terminal_**.

## Usage

```powershell
$ gel
```

```powershell
PS C:\Users\User_1> gel What is the height of the Eiffel tower?
Gemiline:
---
| The height of the Eiffel Tower can vary slightly depending on whether its antenna is included.
|
| *   **Including the antenna:** The current official height of the Eiffel Tower, including the broadcast antenna at its very top, is **330 meters (1,083 feet)**. This height was updated in March 2022 when a new digital radio antenna was installed.
| *   **Without the antenna (to the tip of the flagpole):** The height to the tip of the flagpole (without the antenna) is **324 meters (1,063 feet)**.
|
| The antenna has been changed several times throughout its history, which accounts for the variations in its reported height over the years.
---
PS C:\Users\User_1>

```

---

## Features Included

Flags you can use with gemiline:

- `-s` or `--short` - Get a quick short one line **Answer**
- `-p` or `--persistent` - Have a **persistent interactive session.**

### IN INTERACTIVE MODE

When in interactive mode, after you have gotten your first result, you have the following choices:

- `E` - (E)xit the program
- `T` - Enter (T)ext to continue with query thread.
  - After clicking this, you'll get an interactive input. After entering text, just click enter to get result.
- `M` - Go to main menu to see all threads

When in main menu you will see a history of previous threads with a incremental number beside it.

When you click the number (e.g. 1, 2, 3) on your keyboard or navigate with the h,j,k,l buttons to the desired thread title and press enter, you will see the conversation on that thread.

It will give you a change to go through that thread's history using your scroll mouse or the `j` or `h` keys to navigate down or up respectively.

You will also have the option to press the `E`, `T` or `M` buttons to Exit, Enter new query text or go back to Main Menu respectively.

---

- Example of the `-s` flag being used

```powershell
PS C:\Users\User_1> gel -s What is the height of the Eiffel tower?
Gemiline:
---
|The height of the Eiffel Tower can vary slightly depending on whether its antenna is included.
|
---
PS C:\Users\User_1>

```

- Example of the `-p` flag being used and the `E` Exit button being used

```powershell
PS C:\Users\User_1> gel -p What is the height of the Eiffel tower?
Gemiline:
---
| The height of the Eiffel Tower can vary slightly depending on whether its antenna is included.
|
| *   **Including the antenna:** The current official height of the Eiffel Tower, including the broadcast antenna at its very top, is **330 meters (1,083 feet)**. This height was updated in March 2022 when a new digital radio antenna was installed.
| *   **Without the antenna (to the tip of the flagpole):** The height to the tip of the flagpole (without the antenna) is **324 meters (1,063 feet)**.
|
| The antenna has been changed several times throughout its history, which accounts for the variations in its reported height over the years.
---
|
|---| T - Enter Text  | M - Main Menu |  E - Exit |---|
:> E
|> Exiting program. Have a nice day!
PS C:\Users\User_1>

```

- Example of the `-p` flag being used and the `T` Enter Text button being used

```powershell
PS C:\Users\User_1> gel -p What is the height of the Eiffel tower?
Gemiline:
---
| The height of the Eiffel Tower can vary slightly depending on whether its antenna is included.
|
| *   **Including the antenna:** The current official height of the Eiffel Tower, including the broadcast antenna at its very top, is **330 meters (1,083 feet)**. This height was updated in March 2022 when a new digital radio antenna was installed.
| *   **Without the antenna (to the tip of the flagpole):** The height to the tip of the flagpole (without the antenna) is **324 meters (1,063 feet)**.
|
| The antenna has been changed several times throughout its history, which accounts for the variations in its reported height over the years.
---
|
|---| T - Enter Text  | M - Main Menu |  E - Exit |---|
:> E
| Exiting program. Have a nice day!
PS C:\Users\User_1>

```

- Example of the `-p` flag being used and the `M` Main Menu button being used

```powershell
PS C:\Users\User_1> gel -p What is the height of the Eiffel tower?
Gemiline:
---
| The height of the Eiffel Tower can vary slightly depending on whether its antenna is included.
|
| *   **Including the antenna:** The current official height of the Eiffel Tower, including the broadcast antenna at its very top, is **330 meters (1,083 feet)**. This height was updated in March 2022 when a new digital radio antenna was installed.
| *   **Without the antenna (to the tip of the flagpole):** The height to the tip of the flagpole (without the antenna) is **324 meters (1,063 feet)**.
|
| The antenna has been changed several times throughout its history, which accounts for the variations in its reported height over the years.
---
|
|---| T - Enter Text  | M - Main Menu |  E - Exit |---|
:> M
|
|------------------- GEMILINE ------------------------|
|                    MAIN MENU                        |
|                                                     |
| 1 - Random thing you searched                       |
| 2 - some other random                               |
| 3 - another random thing you searched               |
| 4 - yet another thing you searched                  |
| 5 - the first thing you searched                    |
| 6 - ...cont list of things searched                 |
|
|-----------------------------------------------------|
|
|---| T - Enter Text  | M - Main Menu |  E - Exit |---|
:> E
|> Exiting program. Have a nice day!
PS C:\Users\User_1>

```
