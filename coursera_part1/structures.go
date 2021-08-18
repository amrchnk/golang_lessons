package main

import "fmt"

type Payer interface{
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet)Pay(amount int)error{
	if w.Cash<amount{
		return fmt.Errorf("Не хватает денюжек")
	}
	w.Cash-=amount
	return nil
}

////////////////////////////////////////////////

type Card struct {
	Balance int
	ValidUntil string
	CardHolder string
	CVV string
	Number string
}

func (c *Card) Pay(amount int)error{
	if c.Balance<amount{
		return fmt.Errorf("Не хватает денюжек")
	}
	c.Balance-=amount
	return nil
}

//////////////////////////////////////////

type ApplePay struct {
	Money int
	AppleID string
}

func (a *ApplePay) Pay(amount int)error{
	if a.Money<amount{
		return fmt.Errorf("Не хватает денюжек")
	}
	a.Money-=amount
	return nil
}

///////////////////////////////////////////

func Buy(p Payer){
	switch p.(type) {
	case *Wallet:
		fmt.Println("Oplata nalichnymi?")
	case *Card:
		plastikCard,ok:=p.(*Card)
		if!ok{
			fmt.Println("Не удалось преобразовать к типу Card")
		}
		fmt.Println("Вставляйте карту",plastikCard.CardHolder)
	default:
		println("Something new!")
	}
	err:=p.Pay(200)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("spasibo za pokupku cherez %T\n\n",p)
}

func main(){
	myWallet:=&Wallet{Cash:100}
	Buy(myWallet)

	var myMoney Payer

	myMoney=&Card{Balance: 100,CardHolder: "anek"}
	Buy(myMoney)

	myMoney=&ApplePay{Money: 230,AppleID: "1212"}
	Buy(myMoney)
}