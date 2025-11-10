package types

type MetaData map[string]string

// FileInfo contains information about a single upload resource.
type FileInfo struct {
	// ID is the unique identifier of the upload resource.
	ID string
	// Total file size in bytes specified in the NewUpload call
	Size int64
	// Indicates whether the total file size is deferred until later
	SizeIsDeferred bool
	// Offset in bytes (zero-based)
	Offset   int64
	MetaData MetaData
	// Indicates that this is a partial upload which will later be used to form
	// a final upload by concatenation. Partial uploads should not be processed
	// when they are finished since they are only incomplete chunks of files.
	IsPartial bool
	// Indicates that this is a final upload
	IsFinal bool
	// If the upload is a final one (see IsFinal) this will be a non-empty
	// ordered slice containing the ids of the uploads of which the final upload
	// will consist after concatenation.
	PartialUploads []string
	// Storage contains information about where the data storage saves the upload,
	// for example a file path. The available values vary depending on what data
	// store is used. This map may also be nil.
	Storage map[string]string
}
