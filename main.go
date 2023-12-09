package main

import (
	"fmt"
)

func fillSeatsWithS(rows, seatsAtEachRow int, seats [][]string) {
	for i := 0; i < rows; i++ {
		for j := 0; j < seatsAtEachRow; j++ {
			seats[i][j] = "S"
		}
	}
}

func printSeats(rows int, seatsAtEachRow int, seats [][]string) {
	fmt.Println("Cinema:")
	for i := 0; i < rows; i++ {
		if i == 0 {
			fmt.Print(" ")
			for j := 1; j <= seatsAtEachRow; j++ {
				fmt.Printf("%2d", j)
			}
			fmt.Print("\n")
		}
		fmt.Printf("%d", i+1)
		for j := 0; j < seatsAtEachRow; j++ {
			fmt.Printf("%2s", seats[i][j])
		}
		fmt.Println()
	}
}

func boughtSeatWithB(purchasedTicketCount *int, ticketPrice, rowNumber, seatNumberInRow int, seats [][]string, isBought *bool) int {
	var currentIncome = new(int)
	if seats[rowNumber-1][seatNumberInRow-1] == "B" {
		fmt.Println("\nThat ticket has already been purchased!")
		*isBought = false
		return *currentIncome
	} else if seats[rowNumber-1][seatNumberInRow-1] == "S" {
		seats[rowNumber-1][seatNumberInRow-1] = "B"
		fmt.Printf("\nTicket price: $%d\n", ticketPrice)
		*purchasedTicketCount++
		*isBought = true
		*currentIncome += ticketPrice
		return *currentIncome
	} else if rowNumber > len(seats) || seatNumberInRow > len(seats[0][0]) {
		fmt.Println("Wrong input!")
		*isBought = false
	}
	return *currentIncome
}

func calcStatistics(purchasedTicketCount *int, currentIncome, totalIncome, rows, seatsAtEachRow int) {
	fmt.Printf("\nNumber of purchased tickets: %d\n", *purchasedTicketCount)
	var percentage float32
	percentage = float32(*purchasedTicketCount) / float32(rows*seatsAtEachRow) * 100
	fmt.Printf("Percentage: %.2f%%\n", percentage)
	fmt.Printf("Current income: $%d\n", currentIncome)
	fmt.Printf("Total income: $%d\n", totalIncome)

}

func takeNumberOfRowsAndSeats() (int, int) {
	var rows, seatsAtEachRow int
	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)
	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seatsAtEachRow)
	return rows, seatsAtEachRow
}

func buyTicket() (int, int) {
	var rowNumber, seatNumberInRow int
	fmt.Println("\nEnter a row number:")
	fmt.Scan(&rowNumber)
	fmt.Println("Enter a seat number in that row:")
	fmt.Scan(&seatNumberInRow)

	return rowNumber, seatNumberInRow
}

func calcTotalIncome(rows, seatsAtEachRow int) int {
	totalNumberOfSeats := rows * seatsAtEachRow
	var totalIncome int
	var ticketPrice int
	if totalNumberOfSeats <= 60 {
		ticketPrice = 10
		totalIncome = ticketPrice * totalNumberOfSeats
	} else {
		totalIncome = (rows/2)*10*seatsAtEachRow + 8*(totalNumberOfSeats-(rows/2*seatsAtEachRow))
	}
	return totalIncome
}

func calculateTicketPrice(rows, seatsAtEachRow, rowNumber, seatNumberInRow int) int {
	var ticketPrice int
	if rows*seatsAtEachRow <= 60 {
		ticketPrice = 10
	} else {
		if rowNumber <= rows/2 {
			ticketPrice = 10
		} else {
			ticketPrice = 8
		}
	}
	return ticketPrice
}

func main() {
	var rows, seatsAtEachRow = takeNumberOfRowsAndSeats()
	var purchasedTicketCount *int = new(int)
	var isBought *bool = new(bool)
	*isBought = false
	*purchasedTicketCount = 0
	var currentIncome, totalIncome int
	seats := make([][]string, rows)
	for i := range seats {
		seats[i] = make([]string, seatsAtEachRow)
	}
	fillSeatsWithS(rows, seatsAtEachRow, seats)

	for {
		fmt.Print("\n1. Show the seats\n2. Buy a ticket\n3. Statistics\n0. Exit\n")
		var answer int
		fmt.Scan(&answer)
		totalIncome = calcTotalIncome(rows, seatsAtEachRow)

		if answer == 1 {
			printSeats(rows, seatsAtEachRow, seats)
		} else if answer == 2 {
			for *isBought == false {
				var rowNumber, seatNumberInRow = buyTicket()
				if rowNumber > len(seats) || seatNumberInRow > len(seats[0]) {
					fmt.Println("\nWrong input!")
					*isBought = false
					continue
				} else {
					var ticketPrice int = calculateTicketPrice(rows, seatsAtEachRow, rowNumber, seatNumberInRow)
					currentIncome += boughtSeatWithB(purchasedTicketCount, ticketPrice, rowNumber, seatNumberInRow, seats, isBought)
				}

			}
			*isBought = false
		} else if answer == 0 {
			return
		} else if answer == 3 {
			calcStatistics(purchasedTicketCount, currentIncome, totalIncome, rows, seatsAtEachRow)
		}
	}

}
