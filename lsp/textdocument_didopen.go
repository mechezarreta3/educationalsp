package lsp

type DidOpenTextDocumentNotification struct {
	Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
	/**
	 * The document that was opened
	 */
	TextDocument TextDocumentItem `json:"textdocument"`
}
