# Trie Substring Finder

## Objective

The goal of this program is to find all substrings in a **target word** that match any word from a provided list. Given a main word (the target word) and a set of other words, the program needs to identifie which strings are in part of the main word. It then outputs all matching substrings found in the main word.

## Approach

### 1. Input Handling
The program accepts command-line arguments. The first argument is the **target word** in which we will search for matching substrings. All subsequent arguments are the **words** that will need to be found in the main string.

### 2. Using a Trie for Word Storage
We use a **Trie (prefix tree)** to store the provided words. The Trie is an efficient data structure for storing and searching strings. Each character of a word is stored as a node in the Trie, and the end of each word is marked. By breaking down each word into characters and adding them to the Trie, we can quickly traverse and check whether a substring exists.

### 3. Substring Search in the Target Word
Once the Trie is built, the program iterates over each character of the **target word** and checks for any matching substrings. Starting from each position in the target word, the program attempts to find a sequence of characters that form a valid word stored in the Trie. If a match is found, the substring is stored in the solution array.

### 4. Output
The program prints all the substrings of the target word that are found in the Trie. These are the words that match part of the target word.

### Example Usage

```bash
go run main.go <target_word> <word1> <word2> ... <wordN>
```

## Example
### Input
```bash
go run main.go iamlearninggo learn go happy coding
```
### Output
```bash
Printing all the strings found in the trie:
        - learn
        - go
```