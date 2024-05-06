class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val
    this.next = next === undefined ? null : next
  }
}

// in-place solution
// time: O(n)
// space: O(1)

function reverseList(head: ListNode | null): ListNode | null {
  let prev: ListNode | null = null
  let current: ListNode | null = head
  let next: ListNode | null = null

  while (current !== null) {
    next = current.next
    current.next = prev
    prev = current
    current = next
  }

  return prev
}

// new-copy solution
// time: O(n)
// space: O(n)

// function reverseList(head: ListNode | null): ListNode | null {
//   let newHead: ListNode | null = null
//
//   while (head !== null) {
//     newHead = new ListNode(head.val, newHead)
//     head = head.next
//   }
//
//   return newHead
// }
