package analyzer

import "fmt"

type FileNotFoundError struct {
	FilePath string
	Err      error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable ou inaccessible: %s", e.FilePath)
}

func (e *FileNotFoundError) Unwrap() error {
	return e.Err
}

type ParseError struct {
	Message string
	Err     error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("erreur de parsing: %s", e.Message)
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

func NewFileNotFoundError(filePath string, err error) *FileNotFoundError {
	return &FileNotFoundError{
		FilePath: filePath,
		Err:      err,
	}
}

func NewParseError(message string, err error) *ParseError {
	return &ParseError{
		Message: message,
		Err:     err,
	}
}
