#include <iostream>
#include <vector>
using namespace std;

/**
 * @brief Generates all possible subsequences of a given string using recursive backtracking.
 *
 * This function recursively explores all combinations of characters in the input string,
 * creating subsequences by either including or excluding each character.
 *
 * @param str The original input string from which subsequences are generated.
 * @param current The current subsequence being constructed during recursion.
 * @param index The current index in the input string being processed.
 * @param subsequences A reference to a vector that stores all generated subsequences.
 *
 * @note Time complexity is O(2^n), where n is the length of the input string.
 * @note Space complexity is O(2^n) to store all possible subsequences.
 *
 * @details The function uses two recursive calls for each character:
 * 1. Include the current character in the subsequence
 * 2. Exclude the current character from the subsequence
 *
 * The base case is reached when the index equals the string's length,
 * at which point the current subsequence is added to the result vector.
 */
void generateSubsequences(string str, string current, int index, vector<string>& subsequences) {
    // Base case: If we've reached the end of the string
    if (index == str.size()) {
        subsequences.push_back(current);
        return;
    }

    // Include the current character in the subsequence
    generateSubsequences(str, current + str[index], index + 1, subsequences);

    // Exclude the current character from the subsequence
    generateSubsequences(str, current, index + 1, subsequences);
}

/**
 * @brief Main function to generate and display all subsequences of a user-input string.
 *
 * This function prompts the user to enter a string, generates all possible subsequences
 * using the generateSubsequences function, and then prints out these subsequences.
 *
 * @details The function performs the following steps:
 * 1. Prompts the user to input a string
 * 2. Creates an empty vector to store subsequences
 * 3. Calls generateSubsequences to populate the subsequences vector
 * 4. Prints all generated subsequences to the console
 *
 * @return 0 to indicate successful program execution
 *
 * @note Time complexity is O(2^n), where n is the length of the input string
 * @note Space complexity is also O(2^n) to store all subsequences
 */
int main() {
    string input;
    cout << "Enter a string: ";
    cin >> input;

    vector<string> subsequences;
    generateSubsequences(input, "", 0, subsequences);

    cout << "Subsequences of the string are:" << endl;
    for (const string& s : subsequences) {
        cout << s << endl;
    }

    return 0;
}
