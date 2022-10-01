package types

type File struct {
	FileId       string  `json:"file_id"`
	FileUniqueId string  `json:"file_unique_id"`
	FileSize     *int    `json:"file_size"`
	FilePath     *string `json:"file_path"`
}
