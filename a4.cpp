#include <iostream>
#include <unordered_map>
using namespace std;

// Trie Node class
class TrieNode {
public:
    unordered_map<char, TrieNode*> children;
    bool isEndOfWord;

    /**
     * @brief Constructs a new TrieNode object.
     *
     * Initializes a TrieNode with no children and marks it as not the end of a word.
     * By default, a newly created TrieNode does not represent a complete word.
     *
     * @note The node's children map will be empty upon construction.
     * @note The isEndOfWord flag is set to false by default.
     */
    TrieNode() {
        isEndOfWord = false;
    }
};

// Trie class
class Trie {
private:
    TrieNode* root;

public:
    /**
     * @brief Constructs a new Trie object.
     * 
     * Initializes the Trie by creating a new root TrieNode.
     * The root node serves as the starting point for all Trie operations,
     * allowing insertion, search, and prefix matching of words.
     * 
     * @note The root node is initially empty and does not represent any character.
     * @post A new Trie is created with an empty root node ready for word insertion.
     */
    Trie() {
        root = new TrieNode();
    }

    /**
     * @brief Inserts a word into the Trie data structure.
     *
     * This method adds a word to the Trie by creating nodes for each character
     * if they do not already exist. The last node of the word is marked as
     * the end of a word.
     *
     * @param word The string to be inserted into the Trie
     *
     * @note Time complexity: O(m), where m is the length of the word
     * @note Space complexity: O(m) in the worst case when no characters exist
     *
     * @see search()
     * @see startsWith()
     */
    void insert(string word) {
        TrieNode* current = root;
        for (char c : word) {
            if (current->children.find(c) == current->children.end()) {
                current->children[c] = new TrieNode();
            }
            current = current->children[c];
        }
        current->isEndOfWord = true;
    }

    /**
     * @brief Searches for a complete word in the Trie data structure.
     *
     * @param word The string to search for in the Trie.
     * @return bool True if the entire word exists in the Trie, false otherwise.
     *
     * @details This method traverses the Trie by following the characters of the input word.
     * It checks if each character exists as a child node in the current path.
     * If any character is not found, the method immediately returns false.
     * If all characters are found, it checks if the last node represents the end of a complete word.
     *
     * @note Time complexity: O(m), where m is the length of the word.
     * @note Space complexity: O(1), as it uses a constant amount of extra space.
     *
     * @example
     * Trie trie;
     * trie.insert("apple");
     * bool exists = trie.search("apple");  // Returns true
     * bool notExists = trie.search("app"); // Returns false
     */
    bool search(string word) {
        TrieNode* current = root;
        for (char c : word) {
            if (current->children.find(c) == current->children.end()) {
                return false;
            }
            current = current->children[c];
        }
        return current->isEndOfWord;
    }

    /**
     * @brief Checks if any word in the Trie begins with the specified prefix.
     *
     * @param prefix The character sequence to search for at the beginning of words.
     * @return bool True if the prefix exists in the Trie, false otherwise.
     *
     * @details This method traverses the Trie by following the characters of the prefix.
     * If all characters in the prefix are found in sequence, it returns true, indicating
     * that at least one word in the Trie starts with the given prefix.
     *
     * @note Time complexity: O(m), where m is the length of the prefix.
     * @note Space complexity: O(1), as it only uses a constant amount of extra space.
     *
     * @example
     * Trie trie;
     * trie.insert("apple");
     * bool result = trie.startsWith("app"); // Returns true
     * bool result2 = trie.startsWith("ban"); // Returns false
     */
    bool startsWith(string prefix) {
        TrieNode* current = root;
        for (char c : prefix) {
            if (current->children.find(c) == current->children.end()) {
                return false;
            }
            current = current->children[c];
        }
        return true;
    }
};

/**
 * @brief Main function demonstrating Trie data structure operations.
 *
 * This function showcases the basic functionality of the Trie class by:
 * 1. Creating a Trie instance
 * 2. Inserting words "apple" and "app"
 * 3. Performing search and prefix search operations
 * 4. Printing the results of these operations
 *
 * @return int Exit status of the program (0 indicates successful execution)
 *
 * @note Demonstrates insert(), search(), and startsWith() methods of Trie
 * @example
 * Expected console output:
 * Search 'apple': 1
 * Search 'app': 0
 * StartsWith 'app': 1
 * Search 'app' after inserting: 1
 */
int main() {
    Trie trie;
    trie.insert("apple");
    cout << "Search 'apple': " << trie.search("apple") << endl;   // true
    cout << "Search 'app': " << trie.search("app") << endl;       // false
    cout << "StartsWith 'app': " << trie.startsWith("app") << endl; // true
    trie.insert("app"); 
    cout << "Search 'app' after inserting: " << trie.search("app") << endl; // true

    return 0;
}
