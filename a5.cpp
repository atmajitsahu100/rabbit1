#include <iostream>
#include <vector>
using namespace std;

// Function to generate all subsequences
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
