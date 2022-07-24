package gun

func getGun(gunType string) IGun {
	switch gunType {
	case "ak47":
		return newAK47()
	case "musket":
		return newMusket()
	default:
		return newAK47()
	}

}
