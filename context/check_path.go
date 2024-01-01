package context

func CheckPath(path string) error {
	iter := NewPathUnitIterator(path)
	for iter.HasNext() {
		_, err := iter.Next()
		if err != nil {
			return err
		}

	}
	return nil
}
