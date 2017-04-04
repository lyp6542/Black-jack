package main

import "fmt"
import "math/rand"
import "time"

var (

	playerResult int //player result
	dealerResult int//dealer result
	playerCard [10] string // save player cards
	dealerCard [10] string// save dealer player card


)

//
type Card struct{

	Name string
	Value int
}

//
type Deck []Card 


//
type Dealer struct {

	DealerDeck Deck


}



func GenerateRandomCard(Dd Deck) Card {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	key:=r.Intn(52)
	return Dd[key]



}




func DeckToDealer() Deck {

	Dealer := new(Dealer)
	Dealer.DealerDeck=InitializeDeck()
	return Dealer.DealerDeck

}

func main() {
	
	//fmt.Println(Dealer.DealerDeck)
	//fmt.Println(GenerateRandomCard(Dealer.DealerDeck))

	for true {

		playBlackJack()

		time.Sleep(300 * time.Millisecond)
		fmt.Println("\n continue? Y/N")
		var input string
		fmt.Scan(&input)
		if(input != "Y"){
			break
		}


	}

}


func welcomeShow(){

	fmt.Println("***********Welecome to play Black Jack ******************\n")

}

func playBlackJack(){

	Dealer := new(Dealer)
	Dealer.DealerDeck = InitializeDeck()

	welcomeShow()

	fmt.Println(" Let's start!!!!!! ")

	var input string

	fmt.Println("\n Your turn\n Take your first card and second card")

	var playcard1 = GenerateRandomCard(Dealer.DealerDeck)
	
	playerCard[0] = playcard1.Name
	playerResult = playcard1.Value
	fmt.Println("\n Your first card is ", playcard1.Name)

	time.Sleep(300 * time.Millisecond)
	var playcard2 = GenerateRandomCard(Dealer.DealerDeck)
	playerCard[1] = playcard2.Name
	playerResult += playcard2.Value
	fmt.Println("\n Your Second card is ", playcard2.Name)

	if(playerResult == 20){
		fmt.Println("\n You Win!!!!!")
		return
	}
	
	time.Sleep(300 * time.Millisecond)
	fmt.Println("\n Now it is time for dealer")
	time.Sleep(300 * time.Millisecond)
	var dealercard1 = GenerateRandomCard(Dealer.DealerDeck)
	dealerCard[0] = dealercard1.Name
	dealerResult = dealercard1.Value
	fmt.Println("\n Dealer has got its hidden card")

	time.Sleep(300 * time.Millisecond)
	var dealercard2 = GenerateRandomCard(Dealer.DealerDeck)
	dealerCard[1] = dealercard2.Name
	dealerResult += dealercard2.Value
	fmt.Println("\n Dealer's Second card is ", dealercard2.Name)
	
	time.Sleep(300 * time.Millisecond)
	fmt.Println("\n Now, your point is ", playerResult)
	time.Sleep(300 * time.Millisecond)
	fmt.Println("\n Now, dealer's point you can see is ", dealercard2.Value)


	for true{
		time.Sleep(300 * time.Millisecond)
		fmt.Println("\n continue? Y/N")
		fmt.Scan(&input)
		if(input == "Y"){
			var index = 2
			var playcardAfter = GenerateRandomCard(Dealer.DealerDeck)
			playerCard[index] = playcardAfter.Name
			playerResult += playcardAfter.Value
			fmt.Println("\n Your  card is ", playcardAfter.Name)
			time.Sleep(300 * time.Millisecond)
			fmt.Println("\n Now, your point is ", playerResult)
			if(playerResult > 21){
				break
			}
			index ++	

		}else{
			break
		}
	}

	time.Sleep(300 * time.Millisecond)
	if(playerResult > 21){
		fmt.Println("\n Busting, You Lose!!!!")
		return

	}else{
		for(dealerResult < 17){
			time.Sleep(300 * time.Millisecond)
			fmt.Println(" Now, turn to dealer")
			var index = 2
			var dealcardAfter = GenerateRandomCard(Dealer.DealerDeck)
			dealerCard[index] = dealcardAfter.Name
			dealerResult += dealcardAfter.Value
			fmt.Println("\n  deal get a card is ", dealcardAfter.Name)
			
			if(dealerResult > 21){
				break
			}
			index ++
		}
	}

	time.Sleep(300 * time.Millisecond)
	fmt.Println("\n  dealer hiden card is ", dealerCard[0])
	time.Sleep(300 * time.Millisecond)
	fmt.Println("\n Now, dealer point is ", dealerResult)
	if(dealerResult > 21){
		fmt.Println("\n Busting, You Win!!!!!")
		return
	}else{
	
		if(dealerResult > playerResult){
			fmt.Println("\n Sorry, You Lose!!!!")

		}else if(dealerResult < playerResult){
			fmt.Println("\n You Win!!!!!")
			
		}else{
			fmt.Println("\n This is a even")
		}
	}

}

/*
* 
*/
func InitializeDeck() Deck {
 	cards := make([]Card,0,0)

	//Ace

	for i:=0;i<4;i++{
		cards = append(cards, Card{Name:"A",Value:11})
	}

	// for face value cards
	for i:=2;i<=10;i++{
		for j:=0;j<4;j++{
			cards = append(cards, Card{Name:fmt.Sprintf("%d",i),Value:i})
		}
	}

	aboveFaceValueCards := []string{"J","Q","K"}
	for _,item := range aboveFaceValueCards{
		for i:=0;i<4;i++{
			cards = append(cards, Card{Name:item,Value:10})
		}
	}

	return cards
}