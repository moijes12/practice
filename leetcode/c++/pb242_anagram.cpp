#include <iostream>
#include <map>

using namespace std;

bool is_anagram(string s, string t) {
    map<char, int> smap;
    map<char, int> tmap;

    if (s.size() != t.size()) return false;
    for (char c: s) {
        smap[c]++;
    } 
    for (char c: t) {
        tmap[c]++;
    }

    for (auto it: smap) {
        if (it.second != tmap[it.first]) return false;
    }
    
    return true;
}

int main() {
    string s = "aacc";
    string t = "ccac";
    cout << is_anagram(s, t);
    return 0;
}