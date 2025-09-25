#include <iostream>
#include <bits/stdc++.h>

using namespace std;

int main() {
    string input_string;
    string output_string;
    cin >> input_string;

    for (char s: input_string) {
        s = tolower(s);
        if (s == 'a' || s == 'o' || s == 'y' || s == 'e' || s == 'u' || s == 'i') {
            continue;
        } else {
            output_string += '.';
            output_string += s;
        }
    }
    cout << output_string << endl;
    return 0;
}