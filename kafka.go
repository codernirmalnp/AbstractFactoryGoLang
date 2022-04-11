package main

type kafka struct {
}

func (n *kafka) makeChair() iChair {
	return &productA{
		chair: chair{
			name: "nike",
			size: 14,
		},
	}
}

func (n *kafka) makeSofa() iSofa {
	return &productB{
		sofa: sofa{
			name: "admin",
			size: 14,
		},
	}
}
