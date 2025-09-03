#include <iostream>
#include <vector>

using namespace std;

vector<int> get_product_of_array_except_self(vector<int> &nums) {
    int n = nums.size();
    vector<int> result(n, 1);
    int left = 1;
    int right = 1;
    for (int i = 0; i < n; i++) {
        result[i] *= left;
        left *= nums[i];
        result[n-1-i] *= right;
        right *= nums[i];
    }
    return result;
}