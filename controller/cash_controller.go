package controller

import (
	"fidtest_golang/entity"
	"fidtest_golang/model"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

//According to the assignment, in the case of setting conditions for limiting existing banknotes
func Cashier(w http.ResponseWriter, r *http.Request) {

	product_price, err1 := strconv.ParseFloat(r.FormValue("product_price"), 64)
	if err1 != nil { // if product price is not number type ,it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "กรุณาตั้งราคาเป็นตัวเลขเท่านั้น!")
		return
	}
	cash, err2 := strconv.ParseFloat(r.FormValue("cash"), 64)
	if err2 != nil { // if cash is not number type ,it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "กรุณากรอกเงินที่จ่ายเป็นตัวเลขเท่านั้น!")
		return
	}
	if math.Mod(product_price, 0.25) != 0 { // if product price didn't end with 0.25,0.75,0.50 ,it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "กรุณาตั้งราคาสินค้าใหม่,เนื่องจากเหรียญที่สามารถใช้ได้ 10,5,1,25สตางค์")
		return
	}
	if math.Mod(cash, 0.25) != 0 { // if cash didn't end with 0.25,0.75,0.50 ,it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "กรุณากรอกเงินที่จ่ายให้ถูกต้อง,เหรียญที่สามารถใช้ได้ 10,5,1,25สตางค์")
		return
	}
	bank_note := map[float64]int{1000: 10, 500: 20, 100: 15, 50: 20, 20: 30, 10: 20, 5: 20, 1: 20, 0.25: 50}
	var total_bank_note float64
	var change_1000 int
	var change_500 int
	var change_100 int
	var change_50 int
	var change_20 int
	var change_10 int
	var change_5 int
	var change_1 int
	var change_025 int
	var remainder float64
	_ = remainder
	//find total of bank note
	for k := range bank_note {
		total_bank_note += k * float64(bank_note[k])
	}

	//if cash is equal to product price, we don't need to return change.
	if cash == product_price {
		model.ModelResponseText(w, http.StatusOK, "ขอบคุณที่ใช้บริการ")
		return
	} else if cash < product_price { // if cash is less than product price, it means that you can't buy this product and return response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "ยอดเงินชำระไม่เพียงพอ กรุณาชำระใหม่")
		return
	} else if (cash - product_price) > total_bank_note { //if change is too much more than total_banknote is available on deck, it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "ขออภัย ขณะนี้ธนบัตรไม่เพียงพอในการทอนเงิน,โปรดชำระอีกครั้ง")
		return
	} else if cash > product_price {
		change := math.Round((cash-product_price)*100) / 100
		fmt.Println("change : ", change)
		//if change is > 0 we will calculate the best way to return change to customer as condition limit of bank note
		if change > 0 {
			if int(change/1000) > 10 {
				change_1000 = 10
				change -= float64(change_1000 * 1000)
				remainder = math.Mod(change, 1000)
				total_bank_note = total_bank_note - (float64(change_1000) * 1000)
			} else {
				change_1000 = int(change / 1000)
				change -= float64(change_1000 * 1000)
				remainder = math.Mod(change, 1000)
				total_bank_note = total_bank_note - (float64(change_1000) * 1000)
			}
			if int(change/500) > 20 {
				change_500 = 20
				change -= float64(change_500 * 500)
				remainder = math.Mod(change, 500)
				total_bank_note = total_bank_note - (float64(change_500) * 500)
			} else {
				change_500 = int(change / 500)
				change -= float64(change_500 * 500)
				remainder = math.Mod(change, 500)
				total_bank_note = total_bank_note - (float64(change_500) * 500)
			}
			if int(change/100) > 15 {
				change_100 = 15
				change -= float64(change_100 * 100)
				remainder = math.Mod(change, 100)
				total_bank_note = total_bank_note - (float64(change_500) * 100)
			} else {
				change_100 = int(change / 100)
				change -= float64(change_100 * 100)
				remainder = math.Mod(change, 100)
				total_bank_note = total_bank_note - (float64(change_500) * 100)
			}
			if int(change/50) > 20 {
				change_50 = 20
				change -= float64(change_50 * 50)
				remainder = math.Mod(change, 50)
				total_bank_note = total_bank_note - (float64(change_50) * 500)
			} else {
				change_50 = int(change / 50)
				change -= float64(change_50 * 50)
				remainder = math.Mod(change, 50)
				total_bank_note = total_bank_note - (float64(change_50) * 500)
			}
			if int(change/20) > 30 {
				change_20 = 30
				change -= float64(change_20 * 20)
				remainder = math.Mod(change, 20)
				total_bank_note = total_bank_note - (float64(change_20) * 20)
			} else {
				change_20 = int(change / 20)
				change -= float64(change_20 * 20)
				remainder = math.Mod(change, 20)
				total_bank_note = total_bank_note - (float64(change_20) * 20)
			}
			if int(change/10) > 20 {
				change_10 = 20
				change -= float64(change_10 * 10)
				remainder = math.Mod(change, 10)
				total_bank_note = total_bank_note - (float64(change_10) * 10)
			} else {
				change_10 = int(change / 10)
				change -= float64(change_10 * 10)
				remainder = math.Mod(change, 10)
				total_bank_note = total_bank_note - (float64(change_10) * 10)
			}
			if int(change/5) > 20 {
				change_5 = 20
				change -= float64(change_5 * 5)
				remainder = math.Mod(change, 5)
				total_bank_note = total_bank_note - (float64(change_5) * 5)
			} else {
				change_5 = int(change / 5)
				change -= float64(change_5 * 5)
				remainder = math.Mod(change, 5)
				total_bank_note = total_bank_note - (float64(change_5) * 5)
			}
			if int(change/1) > 20 {
				change_1 = 20
				change -= float64(change_1)
				remainder = math.Mod(change, 1)
				total_bank_note = total_bank_note - float64(change_1)
			} else {
				change_1 = int(change / 1)
				change -= float64(change_1)
				remainder = math.Mod(change, 1)
				total_bank_note = total_bank_note - float64(change_1)
			}
			if int(change/0.25) > 50 {
				change_025 = 50
				change -= float64(change_025)
				remainder = math.Mod(change, 0.25)
				total_bank_note = total_bank_note - (float64(change_025) * 0.25)
			} else {
				change_025 = int(change / 0.25)
				change -= float64(change_025)
				remainder = math.Mod(change, 0.25)
				total_bank_note = total_bank_note - (float64(change_025) * 0.25)
			}
			payload := entity.CashInfo{
				Cash_1000: change_1000,
				Cash_500:  change_500,
				Cash_100:  change_100,
				Cash_50:   change_50,
				Cash_20:   change_20,
				Coin_10:   change_10,
				Coin_5:    change_5,
				Coin_1:    change_1,
				Coin_025:  change_025,
			}
			model.ModelResponseData(w, http.StatusOK, http.StatusText(http.StatusOK), payload)
		}
	}

}

// this function just in case for no codition for limiting existing banknotes
func Cashier_infinity(w http.ResponseWriter, r *http.Request) {
	product_price, err1 := strconv.ParseFloat(r.FormValue("product_price"), 64)
	if err1 != nil { // if product price is not number type ,it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "กรุณาตั้งราคาเป็นตัวเลขเท่านั้น!")
		return
	}
	cash, err2 := strconv.ParseFloat(r.FormValue("cash"), 64)
	if err2 != nil { // if cash is not number type ,it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "กรุณากรอกเงินที่จ่ายเป็นตัวเลขเท่านั้น!")
		return
	}
	if math.Mod(product_price, 0.25) != 0 { // if product price didn't end with 0.25,0.75,0.50 ,it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "กรุณาตั้งราคาสินค้าใหม่,เนื่องจากเหรียญที่สามารถใช้ได้ 10,5,1,25สตางค์")
		return
	}
	if math.Mod(cash, 0.25) != 0 { // if cash didn't end with 0.25,0.75,0.50 ,it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "กรุณากรอกเงินที่จ่ายให้ถูกต้อง,เหรียญที่สามารถใช้ได้ 10,5,1,25สตางค์")
		return
	}
	bank_note := map[float64]int{1000: 10, 500: 20, 100: 15, 50: 20, 20: 30, 10: 20, 5: 20, 1: 20, 0.25: 50}
	var total_bank_note float64
	var change_1000 int
	var change_500 int
	var change_100 int
	var change_50 int
	var change_20 int
	var change_10 int
	var change_5 int
	var change_1 int
	var change_025 int
	var remainder float64
	_ = remainder
	//find total of bank note
	for k := range bank_note {
		total_bank_note += k * float64(bank_note[k])
	}

	//if cash is equal to product price, we don't need to return change.
	if cash == product_price {
		model.ModelResponseText(w, http.StatusOK, "ขอบคุณที่ใช้บริการ")
		return
	} else if cash < product_price { // if cash is less than product price, it means that you can't buy this product and return response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "ยอดเงินชำระไม่เพียงพอ กรุณาชำระใหม่")
		return
	} else if (cash - product_price) > total_bank_note { //if change is too much more than total_banknote is available on deck, it will throw the response as 400
		model.ModelResponseText(w, http.StatusBadRequest, "ขออภัย ขณะนี้ธนบัตรไม่เพียงพอในการทอนเงิน,โปรดชำระอีกครั้ง")
		return
	} else if cash > product_price {
		change := math.Round((cash-product_price)*100) / 100
		fmt.Println("change : ", change)
		if change > 0 {
			//calculate for 1000 banknotes
			change_1000 = int(change / 1000)
			change -= float64(change_1000 * 1000)
			remainder = math.Mod(change, 1000)
			total_bank_note = total_bank_note - (float64(change_1000) * 1000)
			//calculate for 500 banknotes
			change_500 = int(change / 500)
			change -= float64(change_500 * 500)
			remainder = math.Mod(change, 500)
			total_bank_note = total_bank_note - (float64(change_500) * 500)
			//calculate for 100 banknotes
			change_100 = int(change / 100)
			change -= float64(change_100 * 100)
			remainder = math.Mod(change, 100)
			total_bank_note = total_bank_note - (float64(change_500) * 100)
			//calculate for 50 banknotes
			change_50 = int(change / 50)
			change -= float64(change_50 * 50)
			remainder = math.Mod(change, 50)
			total_bank_note = total_bank_note - (float64(change_50) * 50)
			//calculate for 20 banknotes
			change_20 = int(change / 20)
			change -= float64(change_20 * 20)
			remainder = math.Mod(change, 20)
			total_bank_note = total_bank_note - (float64(change_20) * 20)
			//calculate for 10 banknotes
			change_10 = int(change / 10)
			change -= float64(change_10 * 10)
			remainder = math.Mod(change, 10)
			total_bank_note = total_bank_note - (float64(change_10) * 10)
			//calculate for coin 5
			change_5 = int(change / 5)
			change -= float64(change_5 * 5)
			remainder = math.Mod(change, 5)
			total_bank_note = total_bank_note - (float64(change_5) * 5)
			//calculate for coin 1
			change_1 = int(change / 1)
			change -= float64(change_1)
			remainder = math.Mod(change, 1)
			total_bank_note = total_bank_note - float64(change_1)
			//calculate for coin 0.25
			change_025 = int(change / 0.25)
			change -= float64(change_025)
			remainder = math.Mod(change, 0.25)
			total_bank_note = total_bank_note - (float64(change_025) * 0.25)
			payload := entity.CashInfo{
				Cash_1000: change_1000,
				Cash_500:  change_500,
				Cash_100:  change_100,
				Cash_50:   change_50,
				Cash_20:   change_20,
				Coin_10:   change_10,
				Coin_5:    change_5,
				Coin_1:    change_1,
				Coin_025:  change_025,
			}
			model.ModelResponseData(w, http.StatusOK, http.StatusText(http.StatusOK), payload)
		}
	}
}
