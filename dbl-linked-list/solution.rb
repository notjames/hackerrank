#!/usr/bin/env ruby

require 'json'
require 'stringio'

class DoublyLinkedListNode
    attr_accessor :data, :next, :prev

    def initialize node_data
        @data = node_data
        @next = nil
        @prev = nil
    end
end

class DoublyLinkedList
    attr_accessor :head, :tail

    def initialize
        @head = nil
        @tail = nil
    end

    def insert_node node_data
        node = DoublyLinkedListNode.new node_data

        if not self.head
            self.head = node
        else
            self.tail.next = node
            node.prev = self.tail
        end

        self.tail = node
    end
end

def print_doubly_linked_list node, sep, fptr
  while node != nil
    fptr.write node.data

    node = node.next

    fptr.write sep if node != nil
  end
end

# Complete the reverse function below.

#
# For your reference:
#
# DoublyLinkedListNode
#     int data
#     DoublyLinkedListNode next
#     DoublyLinkedListNode prev
#
#
def reverse(head)

end

fptr = File.open(ENV['OUTPUT_PATH'], 'w')

t = gets.to_i

t.times do |t_itr|
  llist = DoublyLinkedList.new

  gets.to_i.times do
      llist_item = gets.to_i
      llist.insert_node llist_item
  end
  print_doubly_linked_list llist, " ", fptr

  #llist1 = reverse llist.head

  #print_doubly_linked_list llist1, " ", fptr
  fptr.write "\n"
end

fptr.close()
