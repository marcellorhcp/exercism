package strand

func ToRNA(dna string) string {
	var rna string
	RNATranscription := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	for _, v := range dna {
		rna += string(RNATranscription[v])
	}
	return rna
}
