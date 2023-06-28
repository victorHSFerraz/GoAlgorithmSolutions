package main

func main() {
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	addTwoNumbers(&l1, &l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var l3 *ListNode
	var l3Head *ListNode
	var carry int

	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.Val
		}
		if l2 != nil {
			sum += l2.Val
		}

		sum += carry

		if sum >= 10 {
			carry = 1
			sum = sum % 10
		} else {
			carry = 0
		}

		if l3 == nil {
			l3 = &ListNode{Val: sum}
			l3Head = l3
		} else {
			l3.Next = &ListNode{Val: sum}
			l3 = l3.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry == 1 {
		l3.Next = &ListNode{Val: carry}
	}

	return l3Head

}
