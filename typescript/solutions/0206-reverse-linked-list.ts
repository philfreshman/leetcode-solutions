export class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val
    this.next = next === undefined ? null : next
  }
}


// recursive solution
// time: O(n)
// space: O(n)

function reverseList0(head: ListNode | null): ListNode | null {
  if (head === null || head.next === null) {
    return head
  }
  const reversedRest = reverseList0(head.next)
  head.next.next = head
  head.next = null
  return reversedRest
}


// in-place solution
// time: O(n)
// space: O(1)

function reverseList1(head: ListNode | null): ListNode | null {
  let prev: ListNode | null = null
  let next: ListNode | null = null

  while (head !== null) {
    next = head.next
    head.next = prev
    prev = head
    head = next
  }

  return prev
}

// new-copy solution
// time: O(n)
// space: O(n)

function reverseList2(head: ListNode | null): ListNode | null {
  let newHead: ListNode | null = null

  while (head !== null) {
    newHead = new ListNode(head.val, newHead)
    head = head.next
  }

  return newHead
}
