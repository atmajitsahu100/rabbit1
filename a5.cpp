#include <iostream>
#include <vector>
using namespace std;

/**
 * @brief Generates all possible subsequences of a given string using recursive backtracking.
 *
 * This function recursively explores all possible combinations of characters
 * from the input string, creating subsequences by either including or excluding
 * each character at the current index.
 *
 * @param str The original input string from which subsequences are generated.
 * @param current The current subsequence being constructed during recursion.
 * @param index The current index in the input string being processed.
 * @param subsequences A reference to a vector that stores all generated subsequences.
 *
 * @note Time complexity is O(2^n), where n is the length of the input string.
 * @note Space complexity is O(2^n) to store all possible subsequences.
 *
 * @warning The function modifies the subsequences vector in-place.
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
 * @brief Main function to generate and display all subsequences of a given string.
 *
 * This function prompts the user to enter a string, generates all possible subsequences
 * using the generateSubsequences function, and then prints out these subsequences.
 *
 * @details The function performs the following steps:
 * 1. Prompt user to input a string
 * 2. Create an empty vector to store subsequences
 * 3. Call generateSubsequences to populate the subsequences vector
 * 4. Print all generated subsequences
 *
 * @return 0 to indicate successful program execution
 *
 * @note Time complexity is O(2^n), where n is the length of the input string
 * @note Space complexity is O(2^n) to store all subsequences
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
