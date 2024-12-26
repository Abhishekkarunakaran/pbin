package service

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	log "log/slog"
	"time"

	"github.com/Abhishekkarunakaran/pbin/src/core/constants"
	"github.com/Abhishekkarunakaran/pbin/src/core/domain"
	"github.com/Abhishekkarunakaran/pbin/src/core/ports"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

type service struct {
	repository ports.Repository
}

func NewPbinService(repo ports.Repository) ports.Service {
	return &service{
		repository: repo,
	}
}

func (s *service) SaveContent(ctx context.Context, payload *domain.Payload) (uuid.UUID, error) {

	//1. generate a uuid

	id, err := uuid.NewV4()
	if err != nil {
		log.Error(err.Error())
		return uuid.Nil, ErrGenUUID
	}

	//2. hash password
	password := []byte(payload.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
		return uuid.Nil, ErrHashPass
	}

	//3a. pad the key to 32
	key := pbkdf2.Key(password, []byte(constants.Env.Salt), 1024, 32, sha256.New)

	//3b. encrypt the content
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return uuid.Nil, ErrGenBlock
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Error(err.Error())
		return uuid.Nil, ErrGenGCM
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		log.Error(err.Error())
		return uuid.Nil, ErrGenNonce
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(payload.Content), nil)

	//4. save the data
	data := domain.Data{
		Password:  string(hashedPassword),
		Content:   string(cipherText),
		CreatedAt: int(time.Now().Unix()),
	}

	if err = s.repository.AddData(ctx, id, data); err != nil {
		log.Error(err.Error())
		return uuid.Nil, ErrSaveData
	}

	// 5. return the generated uuid
	return id, nil
}

func (s *service) GetContent(ctx context.Context, dataRequest *domain.DataRequest) (*domain.Content, error) {
	//1. get data from the db
	data, err := s.repository.GetData(ctx, dataRequest.Id)
	if err != nil {
		log.Error(err.Error())
		return nil, ErrGetData
	}
	if stringIsEmpty(data.Password) {
		log.Error("Data is absent")
		return nil, ErrGetDataAbsent
	}
	//2. compare the password throw error if not matched
	if err = bcrypt.CompareHashAndPassword([]byte(data.Password),[]byte(dataRequest.Password)); err != nil{
		log.Error(err.Error())
		return nil, ErrIncorrectPassword
	}
	//3. Decrypt the content using the password
	key := pbkdf2.Key([]byte(dataRequest.Password), []byte(constants.Env.Salt), 1024, 32, sha256.New)


	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err.Error())
		return nil, ErrGenBlock
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Error(err.Error())
		return nil, ErrGenGCM
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := data.Content[:nonceSize],data.Content[nonceSize:]

	plainText, err := gcm.Open(nil, []byte(nonce),[]byte(cipherText),nil)
	if err != nil {
		log.Error(err.Error())
		return nil, ErrDecrypting
	}
	// 4. remove the data from the db
	err = s.repository.RemoveData(ctx,dataRequest.Id)
	if err != nil {
		log.Error(err.Error())
	}
	
	// 5. return the content
	content := domain.Content(string(plainText))
	return &content, nil
}

func (s *service) IsContentPresent(ctx context.Context, id uuid.UUID) bool {
	return s.repository.IsContentPresent(ctx,id)
}