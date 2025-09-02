#include <iostream>
#include <vector>

using namespace std;

int missingNumber(vector<int> &nums) {
    int result = 0;
    for (auto it: nums) {
        result ^= it;
    }
    return result;
}

int main() {
    vector<int> v{2,2,1};
    missingNumber(v);
}