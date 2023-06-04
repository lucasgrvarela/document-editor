package main

import "fmt"

type Document struct {
	Content string
}

type DocumentEditor struct {
	document          Document
	undoHistory       *LinkedList
	redoHistory       *LinkedList
	currentUndoAction *Node
}

type Node struct {
	Value    Document
	Previous *Node
	Next     *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (ll *LinkedList) Add(value Document) {
	node := &Node{Value: value}

	if ll.Head == nil {
		// If the linked list is empty, set the new node as both the head and the tail
		ll.Head = node
		ll.Tail = node
	} else {
		// If the linked list is not empty, append the new node to the tail
		ll.Tail.Next = node
		node.Previous = ll.Tail
		ll.Tail = node
	}
}

func (ll *LinkedList) RemoveLast() {
	if ll.Tail == nil {
		return
	}

	if ll.Tail.Previous != nil {
		// If there is a previous node, update the tail to the previous node
		ll.Tail = ll.Tail.Previous
		ll.Tail.Next = nil
	} else {
		// If there is no previous node, the linked list becomes empty
		ll.Head = nil
		ll.Tail = nil
	}
}

func (de *DocumentEditor) AddContent(content string) {
	// Append the content to the existing document content
	de.document.Content += content
}

func (de *DocumentEditor) Save() {
	// Add the current document state to the undo history
	de.undoHistory.Add(de.document)
	// Create a new instance of LinkedList for redo history
	de.redoHistory = &LinkedList{}
	de.currentUndoAction = nil
}

func (de *DocumentEditor) Undo() {
	if de.undoHistory.Head == nil {
		return
	}

	if de.currentUndoAction == nil {
		// If there is no current undo action, set it to the last action in the undo history
		de.currentUndoAction = de.undoHistory.Tail
	} else if de.currentUndoAction.Previous != nil {
		// If there is a previous undo action, update the current undo action to the previous one
		de.currentUndoAction = de.currentUndoAction.Previous
	}

	de.document = de.currentUndoAction.Previous.Value
	// Add the last undo action to the redo history
	de.redoHistory.Add(de.undoHistory.Tail.Value)
	// Remove the last action from the undo history
	de.undoHistory.RemoveLast()
}

func (de *DocumentEditor) Redo() {
	if de.redoHistory.Head == nil {
		return
	}

	de.document = de.redoHistory.Tail.Value
	// Add the last redo action to the undo history
	de.undoHistory.Add(de.redoHistory.Tail.Value)
	// Remove the last action from the redo history
	de.redoHistory.RemoveLast()

	// Update the currentUndoAction to the last action added to the undo history
	de.currentUndoAction = de.undoHistory.Tail
}

func main() {
	documentEditor := DocumentEditor{
		undoHistory: &LinkedList{},
		redoHistory: &LinkedList{},
	}

	// Initial document state
	documentEditor.AddContent("Hello, ")
	documentEditor.Save()

	// Perform some changes
	documentEditor.AddContent("world!")
	documentEditor.Save()
	documentEditor.AddContent("world!")
	documentEditor.Save()
	documentEditor.AddContent("world!")
	documentEditor.Save()
	fmt.Println(documentEditor.document.Content) // "Hello, world!world!world!"

	// Undo the changes
	documentEditor.Undo()
	fmt.Println(documentEditor.document.Content) // "Hello, world!world!"

	// Undo the changes
	documentEditor.Undo()
	fmt.Println(documentEditor.document.Content) // "Hello, world!"

	// Undo the changes
	documentEditor.Undo()
	fmt.Println(documentEditor.document.Content) // "Hello, " -> initial document state

	// Redo the changes
	documentEditor.Redo()
	fmt.Println(documentEditor.document.Content) // "Hello, world!"

	// Redo the changes
	documentEditor.Redo()
	fmt.Println(documentEditor.document.Content) // "Hello, world!world!"

	// Redo the changes
	documentEditor.Redo()
	fmt.Println(documentEditor.document.Content) // "Hello, world!world!world!"
}
