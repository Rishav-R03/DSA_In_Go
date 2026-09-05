class Node:
    def __init__(self,data:int):
        self.data = data 
        self.next = None

    def __repr__(self):
        """Developer-Friendly string representation"""
        return f"Node:{self.data}"
    
class LinkedList:
    def __init__(self):
        self.head = None 

    def append(self,data:int):
        """Add node to the end of the list""" 
        newNode = Node(data)
        # if head is not there
    
        if not self.head:
            self.head = newNode
            return 
        # if head is there
        curr = self.head 
        while curr.next:
            curr = curr.next 
        curr.next = newNode

    def prepend(self,data):
        """Add a node to the beginning of the list: O(1) time."""
        new_node = Node(data)
        new_node.next = self.head 
        self.head = new_node 

    def delete_by_value(self,target:int):
        """Remove the first node containing the target value: O(N) time."""
        if not self.head:
            return 
        if self.head.data == target:
            self.head = self.head.next 
            return 

        curr = self.head 
        while curr.next and curr.next.data != target:
            curr = curr.next 

        if curr.next:
            curr.next = curr.next.next

    def search(self,target:int):
        """Search for a value in the list: O(N) time."""
        curr = self.head 
        while curr:
            if curr.data == target:
                return True
            curr = curr.next
        return False 

    def reverse(self):
        """Reverse the linked list in-place using three pointter: O(N)"""
        prev = None 
        curr = self.head 

        while curr:
            next_node = curr.next 
            curr.next = prev 
            prev = curr 
            curr = next_node 

        self.head = prev 

    def __str__(self):
        """Return readable string representation"""
        nodes = []
        curr = self.head 
        while curr:
            nodes.append(str(curr.data))
            curr = curr.next 
        return " -> ".join(nodes) + " -> None"

# --- Example Usage ---
if __name__ == "__main__":
    ll = LinkedList()

    # Append items
    ll.append(10)
    ll.append(20)
    ll.append(30)
    print("Initial List:", ll)
    # Output: 10 -> 20 -> 30 -> None

    # Prepend item
    ll.prepend(5)
    print("After prepending 5:", ll)
    # Output: 5 -> 10 -> 20 -> 30 -> None

    # Search items
    print("Is 20 in list?", ll.search(20))  # Output: True
    print("Is 99 in list?", ll.search(99))  # Output: False

    # Delete an item
    ll.delete_by_value(20)
    print("After deleting 20:", ll)
    # Output: 5 -> 10 -> 30 -> None

    # Reverse list
    ll.reverse()
    print("Reversed List:", ll)
    # Output: 30 -> 10 -> 5 -> None