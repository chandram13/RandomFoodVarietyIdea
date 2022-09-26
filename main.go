// Marvish Chandra

func desiredCuisine(foodChoice):
	IndianFood = []
	JapaneseFood = []
	Seafood = []
	ChineseFood = []
	ItalianFood = []
	AmericanFood = []
	MexicanFood = []
	if foodChoice == "spiciness":
		indianFood += 1
		ChineseFood += 1
		MexicanFood += 1
	if foodChoice == "peppers":
		indianFood += 1
	if foodChoice == "complexity":
		indianFood += 1
	if foodChoice == "enjoy rice":
		JapaneseFood += 1
	if foodChoice == "earthy":
		JapaneseFood += 1
	if foodChoice == "wasabi":
		JapaneseFood += 1
		Seafood += 1
	if foodChoice == "noodles":
		JapaneseFood += 1
		ChineseFood += 1
	if foodChoice == "broth":
		JapaneseFood += 1
	if foodChoice == "mild":
		Seafood += 1
	if foodChoice = "dry":
		Seafood += 1
	if foodChoice == "sweet":
		Seafood += 1
		ChineseFood += 1
	if foodChoice == "salty":
		ChineseFood += 1
	if foodChoice == "sour":
		ChineseFood += 1
	if foodChoice == "bitter":
		ChineseFood += 1
	if foodChoice == "dry meat":
		ItalianFood += 1
		AmericanFood += 1
	if foodChoice == "cheese":
		ItalianFood += 1
		AmericanFood += 1
	if foodChoice == "tomato":
		ItalianFood += 1
	if foodChoice == "basil":
		ItalianFood += 1
	if foodChoice == "olive oil":
		ItalianFood += 1
	if foodChoice == "pasta":
		ItalianFood += 1
	if foodChoice == "fried":
		AmericanFood += 1
	if foodChoice == "pasta":
		AmericanFood += 1
	if foodChoice == "cheese melt":
		AmericanFood += 1
	if foodChoice == "street type":
		MexicanFood += 1
	if foodChoice == "guacamole":
		MexicanFood += 1
	if foodChoice == "corn dough":
		MexicanFood += 1
	else: fmt.printIn("Please describe your food choice based on flavor or specific food.")

func main(){
	fmt.PrintIn(IndianFood)
	fmt.PrintIn(JapaneseFood)
	fmt.PrintIn(Seafood)
	fmt.PrintIn(ChineseFood)
	fmt.PrintIn(ItalianFood)
	fmt.PrintIn(AmericanFood)
	fmt.PrintIn(MexicanFood)
}