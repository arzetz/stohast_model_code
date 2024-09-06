package main

import (
	"fmt"
	"math"
)

type stohModel struct {
	podiumRate       float32
	topFiveRate      float32
	randomDamageRate float32
	crashRate        float32
	airTemp          float32
}

type stohPointsModel struct {
	pilotName   string
	podiumRate  float32
	topFiveRate float32
	crashRate   float32
	efficiency  float32
}

func efficiencyRate(pilots ...*stohPointsModel) {
	for _, dude := range pilots {
		dude.efficiency = dude.podiumRate + dude.topFiveRate*(1-dude.crashRate)
		fmt.Printf("\nЭффективность пилота %s равна %d%%", dude.pilotName, int(dude.efficiency*100))
	}

}

func main() {
	model := stohModel{0.1, 0.3, 0.2, 0.27, 0.3}
	winChance := model.podiumRate * model.topFiveRate * (1 - model.randomDamageRate) * (1 - model.crashRate) * (1 - model.airTemp)
	myWinChance := float64(winChance)
	anyCarWinChance := float64(winChance * 3)
	fmt.Printf("Шанс победы любой из трёх машин: %f%% \nШанс моей победы: %f%%", (anyCarWinChance+(1-math.Pow((1-anyCarWinChance), 5)))*100, (myWinChance+(1-math.Pow((1-myWinChance), 5)))*100)
	AMelnikov := &stohPointsModel{"Артемий", 0.04, 0.3, 0.3, 0}
	MArkhangelskiy := &stohPointsModel{"Макс", 0.06, 0.2, 0.4, 0}
	IOvsienko := &stohPointsModel{"Иван", 0.2, 0.3, 0.2, 0}
	efficiencyRate(AMelnikov, MArkhangelskiy, IOvsienko)
	//+podiumRate + topFiveRate
}
