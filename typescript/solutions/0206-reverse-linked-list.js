"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.reverseList = exports.ListNode = void 0;
class ListNode {
    constructor(val, next) {
        this.val = val === undefined ? 0 : val;
        this.next = next === undefined ? null : next;
    }
}
exports.ListNode = ListNode;
function reverseList(head) {
    let newHead = null;
    while (head !== null) {
        newHead = new ListNode(head.val, newHead);
        head = head.next;
    }
    return newHead;
}
exports.reverseList = reverseList;
