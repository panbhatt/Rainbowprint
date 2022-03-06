package rainbowprint

tyep RGB struct {
	Red, Green, Blue int 	
}

func NewRGB(red int, green int, blue int) RGB {

	return RGB {
		int(math.Sin(0.1*i)*127 + 128 ),
		int(math.Sin(0.1*i+2*math.Pi/3)*127 + 128 ),
		int(math.Sin(0.1*i+4*math.Pi/3)*127 + 128 ),
	}

}