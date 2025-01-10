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
     * Initializes a TrieNode with no word ending marker. 
     * The isEndOfWord flag is set to false by default, indicating 
     * that this node does not represent the end of a complete word.
     * 
     * @note This constructor is typically used when creating new nodes 
     * during Trie insertion or traversal.
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
     * @note The root node is initially empty and does not represent a complete word.
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
     * @return bool True if the entire word exists in the Trie as a complete word, false otherwise.
     *
     * @details This method traverses the Trie by following the characters of the input word.
     * It checks each character's existence in the Trie and ensures the word is marked as a complete word.
     *
     * Time Complexity: O(m), where m is the length of the word.
     * Space Complexity: O(1), as it uses a constant amount of extra space.
     *
     * @note Returns false if:
     * - Any character in the word is not found in the Trie
     * - The word is not marked as a complete word (even if its prefix exists)
     *
     * @example
     * Trie trie;
     * trie.insert("apple");
     * trie.search("apple");   // Returns true
     * trie.search("app");     // Returns false
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
     * @brief Checks if any word in the Trie starts with the given prefix.
     *
     * @param prefix The string prefix to search for in the Trie.
     * @return bool True if the prefix exists in the Trie, false otherwise.
     *
     * @details This method traverses the Trie by following the characters of the prefix.
     * If all characters in the prefix are found in the Trie, it returns true.
     * If any character is not found, it returns false.
     *
     * @complexity Time complexity: O(m), where m is the length of the prefix.
     * @complexity Space complexity: O(1), as it uses a constant amount of extra space.
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
 * @brief Main function demonstrating Trie data structure functionality
 *
 * This function showcases the basic operations of a Trie:
 * - Inserting words
 * - Searching for complete words
 * - Checking word prefixes
 *
 * Demonstrates the following scenarios:
 * 1. Inserting "apple" and verifying its presence
 * 2. Attempting to search for a partial word "app" before insertion
 * 3. Checking if "app" is a valid prefix
 * 4. Inserting "app" and then searching for it
 *
 * @return int Exit status of the program (0 indicates successful execution)
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
