# Document Editor

This is a simple document editor program implemented in Go. It allows you to perform actions such as adding content to a document, saving the document state, undoing the changes, and redoing the changes.

## Usage

To use the document editor, follow these steps:

1. Initialize a new `DocumentEditor` instance.

```go
documentEditor := DocumentEditor{
    undoHistory: &LinkedList{},
    redoHistory: &LinkedList{},
}
```

2. Add content to the document using the AddContent() method.
```go
documentEditor.AddContent("Hello, ")
```

3. Save the current state of the document using the Save() method.
```go
documentEditor.Save()
```

4. Perform changes to the document by adding more content and saving after each change.
```go
documentEditor.AddContent("world!")
documentEditor.Save()
```

5. Undo the changes using the Undo() method.
```go
documentEditor.Undo()
```

6. Redo the changes using the Redo() method.
```go
documentEditor.Redo()
```
