package handlers

import (
	dto "dumbmerch/dto/result"
	transactiondto "dumbmerch/dto/transaction"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	// import midtrans-go package here ...
	// import midtrans-go/coreapi package here ...
	// import midtrans-go/snap package here ...
)

// Declare Coreapi Client here ...

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) FindTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	transactions, err := h.TransactionRepository.FindTransactions(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	var responseTransaction []transactiondto.TransactionResponse
	for _, t := range transactions {
		responseTransaction = append(responseTransaction, convertResponseTransaction(t))
	}

	for i, t := range responseTransaction {
		imagePath := os.Getenv("PATH_FILE") + t.Product.Image
		responseTransaction[i].Product.Image = imagePath
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: responseTransaction}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTransaction) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	var request transactiondto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Unique Transaction Id here ...

	transaction := models.Transaction{
		// ID: TransactionId,
		ProductID: request.ProductId,
		BuyerID: userId,
		SellerID: request.SellerId,
		Price: request.Price,
		Status: "pending",
	}

	log.Print(transaction)

	// newTransaction, err := h.TransactionRepository.CreateTransaction(transaction)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(err.Error())
	// 	return
	// }

	// dataTransactions, err := h.TransactionRepository.GetTransaction(newTransaction.ID)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(err.Error())
	// 	return
	// }
	
	// Request payment token from midtrans here ...
}

// Notification method ...

func convertResponseTransaction(t models.Transaction) transactiondto.TransactionResponse {
	return transactiondto.TransactionResponse{
		ID:      	t.ID,
		Product:   	t.Product,
		Buyer:  	t.Buyer,
		Seller: 	t.Seller,
		Price:  	t.Price,
		Status:    	t.Status,
	}
}