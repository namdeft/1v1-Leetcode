#include <iostream>
#include <vector>
using namespace std;

int main() {
    vector<int> nums = {-1, 0, 3, 5, 9, 12};
    int target = 9;
    int size = nums.size();
    int left = 0;
    int right = size - 1;

    while (left < right) {
        int mid = (left + right) / 2;

        if (mid < target) {
        }
    }

    return 0;
}