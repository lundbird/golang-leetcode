/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getDecimalValue(head *ListNode) int {
    num := head.Val
    for ;head.Next!=nil;head=head.Next {
        num = 2*num + head.Next.Val
    }
    return num
}