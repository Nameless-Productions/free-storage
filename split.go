package main

var chunkSize = 9 * 1024 * 1024 // 9mb in bytes

func splitFile(data []byte) [][]byte {
	if len(data) <= chunkSize {
		return [][]byte{data}
	}

	var chunks [][]byte

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}

	return chunks
}