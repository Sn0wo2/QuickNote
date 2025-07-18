package note

func (n *Note) Write() error {
	err := n.Encode()
	if err != nil {
		return err
	}

	return SetNote(*n)
}

func (n *Note) Read() error {
	no, err := GetNote(n.NID)
	if err != nil {
		return err
	}

	return n.Decode(no.Data)
}

func (n *Note) Delete() error {
	return DeleteNote(n.NID)
}
