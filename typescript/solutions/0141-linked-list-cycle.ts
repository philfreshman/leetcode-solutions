import { ListNode } from "./0206-reverse-linked-list"

// time: O(n)
// space: O(1)

function hasCycle(head: ListNode | null): boolean {
  let slow = head
  let fast = head

  while (fast?.next?.next) {
    if (slow === fast) return true
    slow = slow?.next!
    fast = fast?.next.next
  }

  return false
}
