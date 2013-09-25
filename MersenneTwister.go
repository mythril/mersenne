package mersenne

type mersenneTwister struct {
	seed uint32
	iteration uint64
	mt [624]uint32
}

func New(s uint32) *mersenneTwister {
	t := mersenneTwister{seed: s}
	t.mt[0] = s
	
	for i := 1; i < 624; i += 1 {
		prev := t.mt[i - 1]
		shifted := prev >> 30
		xored := prev ^ shifted
		t.mt[i] = 1812433253 * xored + uint32(i)
	}
	t.generateMore()
	return &t
}

func (t *mersenneTwister) generateMore() {
	for i := 0; i < 624; i += 1 {
		// bit 31 (32nd bit) of mt[i]
		lastBit := t.mt[i] & 0x80000000
		// bits 0-30 (first 31 bits) of mt[...]
		firstBitsNext := t.mt[(i + 1) % 624] & 0x7fffffff
		y := lastBit + firstBitsNext
		div2 := y >> 1
		three97later := t.mt[(i + 397) % 624]
		t.mt[i] = three97later ^ div2
		
		if (y % 2) != 0 { // y is odd
			t.mt[i] = t.mt[i] ^ (2567483615); // 0x9908b0df
		}
	}
}

func (t *mersenneTwister) Next() {
	if (t.iteration % 624) == 0 && t.iteration > 0 {
		t.generateMore()
	}
	t.iteration += 1
}

func (t *mersenneTwister) Get() uint32 {
	y := t.mt[t.iteration % 624]
	y = y ^ (y >> 11)
	y = y ^ ((y << 7) & (2636928640))
	y = y ^ ((y << 15) & (4022730752))
	y = y ^ (y >> 18)
	
	return y
}

func (t *mersenneTwister) GetIteration() uint64 {
	return t.iteration
}

func (t *mersenneTwister) GetSeed() uint32 {
	return t.seed
}


