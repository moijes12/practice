#include <iostream>
#include <unordered_set>
#include <vector>

using namespace std;

bool has_duplicate(vector<int> &nums) {
    return nums.size() > unordered_set<int>(nums.begin(), nums.end()).size(); 
}

int main() {
    vector<int> c = {1,2,3,4,1};
    cout << has_duplicate(c) << endl;
    return 0;
}