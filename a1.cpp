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
     * such as inserting, searching, and checking word prefixes.
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
     * @return bool True if the entire word exists in the Trie, false otherwise.
     *
     * @details This method traverses the Trie character by character. 
     * It checks if each character exists as a child node in the current path. 
     * If any character is not found, the method immediately returns false. 
     * At the end of traversal, it verifies if the last node represents a complete word.
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
     * @param prefix The prefix to search for in the Trie.
     * @return bool True if at least one word in the Trie starts with the given prefix, false otherwise.
     *
     * @details This method traverses the Trie by following the characters of the prefix.
     * If all characters in the prefix are found in the Trie, it returns true, indicating
     * that there is at least one word that starts with the given prefix.
     *
     * @note Time complexity: O(m), where m is the length of the prefix.
     * @note Space complexity: O(1), as it only uses a constant amount of extra space.
     *
     * @example
     * Trie trie;
     * trie.insert("apple");
     * bool result = trie.startsWith("app"); // Returns true
     * result = trie.startsWith("ban");      // Returns false
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
 * The example demonstrates:
 * 1. Inserting "apple" into the Trie
 * 2. Searching for "apple" (expected: true)
 * 3. Searching for "app" (expected: false)
 * 4. Checking prefix "app" (expected: true)
 * 5. Inserting "app"
 * 6. Searching for "app" after insertion (expected: true)
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
