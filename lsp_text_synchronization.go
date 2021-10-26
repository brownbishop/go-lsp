package lsp

type DidOpenTextDocumentParams struct {
	// The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument,required"`
}

type DidCloseTextDocumentParams struct {
	// The document that was closed.
	TextDocument TextDocumentIdentifier `json:"textDocument,required"`
}

type DidChangeTextDocumentParams struct {
	// The document that did change. The version number points
	// to the version after all provided content changes have
	// been applied.
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument,required"`

	// The actual content changes. The content changes describe single state
	// changes to the document. So if there are two content changes c1 (at
	// array index 0) and c2 (at array index 1) for a document in state S then
	// c1 moves the document from S to S' and c2 from S' to S''. So c1 is
	// computed on the state S and c2 is computed on the state S'.
	//
	// To mirror the content of a document using change events use the following
	// approach:
	// - start with the same initial content
	// - apply the 'textDocument/didChange' notifications in the order you
	//   receive them.
	// - apply the `TextDocumentContentChangeEvent`s in a single notification
	//   in the order you receive them.
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges,required"`
}

type TextDocumentIdentifier struct {
	// The text document's URI.
	URI DocumentURI `json:"uri,required"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	// The version number of this document.
	//
	// The version number of a document will increase after each change,
	// including undo/redo. The number doesn't need to be consecutive.
	Version int `json:"version.required"`
}

// An event describing a change to a text document. If range and rangeLength are
// omitted the new text is considered to be the full content of the document.
type TextDocumentContentChangeEvent struct {
	// The range of the document that changed.
	Range Range `json:"range,omitempty"`

	// The optional length of the range that got replaced.
	//
	// @deprecated use range instead.
	RangeLength *int `json:"rangeLength,omitempty"`

	// The new text for the provided range or the new text of the whole document if range and rangeLength are omitted.
	Text string `json:"text,required"`
}

type TextDocumentItem struct {
	// The text document's URI.
	URI DocumentURI `json:"uri,required"`

	// The text document's language identifier.
	LanguageID string `json:"languageId,required"`

	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int `json:"version,required"`

	// The content of the opened text document.
	Text string `json:"text,required"`
}

type DidSaveTextDocumentParams struct {
	// The document that was saved.
	TextDocument TextDocumentIdentifier `json:"textDocument,required"`

	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text string `json:"text,omitempty"`
}

type RenameParams struct {
	*TextDocumentPositionParams
	*WorkDoneProgressParams

	// The new name of the symbol. If the given name is not valid the
	// request must return a [ResponseError](#ResponseError) with an
	// appropriate message set.
	NewName string `json:"newName,required"`
}