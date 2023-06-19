/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
   public:
    void reorderList(ListNode* head) {
        if (head == nullptr || head->next == nullptr) return;

        // Head of first list
        ListNode* l1 = head;
        // Head of second list
        ListNode* slow = head;
        // Tail of second list
        ListNode* fast = head;
        // Tail of first list
        ListNode* prev = nullptr;

        while (fast != nullptr && fast->next != nullptr) {
            prev = slow;
            slow = slow->next;
            fast = fast->next->next;
        }

        prev->next = nullptr;

        ListNode* l2 = reverse(slow);

        mergeTwoLists(l1, l2);
    }

    ListNode* reverse(ListNode* head) {
        ListNode* curr = head;
        ListNode* prev = nullptr;

        while (curr != nullptr) {
            ListNode* next = curr->next;
            curr->next = prev;
            prev = curr;
            curr = next;
        }

        return prev;
    }

    void mergeTwoLists(ListNode* l1, ListNode* l2) {
        while (l1 != nullptr) {
            ListNode* l1_next = l1->next;
            ListNode* l2_next = l2->next;

            l1->next = l2;
            if (l1_next == nullptr) {
                break;
            }
            l2->next = l1_next;

            l1 = l1_next;
            l2 = l2_next;
        }
    }
};
