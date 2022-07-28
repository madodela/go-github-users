package repository

import "errors"

var (
	ErrOpeningCSV        = errors.New("opening-csv-file")
	ErrInvalidFileStruct = errors.New("invalid-csv-file-struct")
	ErrReadingLineCSV    = errors.New("reading-line-csv-file")
)
