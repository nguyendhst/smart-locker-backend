package api

import (
	"context"
	"net/http"

	"smart-locker/backend/db"
	"smart-locker/backend/db/sqlc"
	"smart-locker/backend/token"

	"github.com/antihax/optional"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	swagger "github.com/nguyendhst/adafruit-go-client-v2"
	log "github.com/rs/zerolog/log"
)

var (
	LOCKED   = "0"
	UNLOCKED = "1"
)

type (
	UnlockRequest struct {
		// NFC signature.
		NFCSig string `json:"nfc_sig"`
	}

	UnlockResponse struct {
		// Status of the locker.
		Status string `json:"status"`
	}

	LockRequest struct {
		// NFC signature.
		NFCSig string `json:"nfc_sig"`
	}

	LockResponse struct {
		// Status of the locker.
		Status string `json:"status"`
	}

	AllInfoRequest struct {
	}

	AllInfoResponse struct {
		LockersInfo []db.LockerInfo `json:"lockers_info"`
	}

	CreateLockerRequest struct {
		LockerNumber int32  `json:"locker_num"`
		Location     string `json:"location"`
		//Status       string `json:"status"`
		NFCSig string `json:"nfc_sig"`
	}

	CreateLockerResponse struct {
		Status string `json:"status"`
	}

	RemoveLockerRequest struct {
	}

	RemoveLockerResponse struct {
		Status string `json:"status"`
	}

	RegisterLockerRequest struct {
		NFCSig string `json:"nfc_sig"`
	}

	RegisterLockerResponse struct {
		Status string `json:"status"`
	}
)

// unlockLocker unlocks the locker.
func (s *Server) unlockLocker(c echo.Context) error {

	var unlockReq UnlockRequest
	if err := c.Bind(&unlockReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	params := db.GetFeedByNFCSigParams{
		NFCSig: unlockReq.NFCSig,
	}

	lock, err := s.Store.ExecGetFeedByNFCSigTx(
		context.Background(),
		params,
	)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// check if already unlocked
	resp, httpResp, err := s.AdafruitClient.DataApi.LastData(
		context.Background(),
		s.Config.AdafruitUsername,
		lock.Feedkey,
		&swagger.DataApiLastDataOpts{
			Include: optional.NewString("value,created_at"),
		},
	)
	if err != nil || httpResp.StatusCode != http.StatusOK {

		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)

	} else if resp.Value == UNLOCKED {

		if _, err = s.Store.ExecUpdateLockStatusTx(
			context.Background(),
			db.UpdateLockStatusParams{
				Status: sqlc.LockersLockStatusUnlocked,
				Id:     lock.LockerId,
			},
		); err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusAlreadyReported, UnlockResponse{
			Status: "Already unlocked",
		})

	} else if resp.Value == LOCKED {

		if _, err = s.Store.ExecUpdateLockStatusTx(
			context.Background(),
			db.UpdateLockStatusParams{
				Status: sqlc.LockersLockStatusUnlocked,
				Id:     lock.LockerId,
			},
		); err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		_, httpResp, err = s.AdafruitClient.DataApi.CreateData(
			context.Background(),
			s.Config.AdafruitUsername,
			lock.Feedkey,
			swagger.Datum{
				Value: LOCKED,
			},
		)

		if err != nil || httpResp.StatusCode != http.StatusOK {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, UnlockResponse{
			Status: "Unlocked",
		})

	} else {

		return c.JSON(http.StatusBadRequest, UnlockResponse{
			Status: "Bad request",
		})
	}
}

// lockLocker locks the locker.
func (s *Server) lockLocker(c echo.Context) error {

	var lockReq LockRequest
	if err := c.Bind(&lockReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	params := db.GetFeedByNFCSigParams{
		NFCSig: lockReq.NFCSig,
	}

	lock, err := s.Store.ExecGetFeedByNFCSigTx(
		context.Background(),
		params,
	)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// ping the adafruit server to check if already locked
	resp, httpResp, err := s.AdafruitClient.DataApi.LastData(
		context.Background(),
		s.Config.AdafruitUsername,
		lock.Feedkey,
		&swagger.DataApiLastDataOpts{
			Include: optional.NewString("value,created_at"),
		},
	)

	// compare to the registered state

	if err != nil || httpResp.StatusCode != http.StatusOK {

		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)

	} else if resp.Value == LOCKED {

		if _, err = s.Store.ExecUpdateLockStatusTx(
			context.Background(),
			db.UpdateLockStatusParams{
				Status: sqlc.LockersLockStatusLocked,
				Id:     lock.LockerId,
			},
		); err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusAlreadyReported, LockResponse{
			Status: "Already locked",
		})

	} else if resp.Value == UNLOCKED {

		// sync with DB
		if _, err = s.Store.ExecUpdateLockStatusTx(
			context.Background(),
			db.UpdateLockStatusParams{
				Status: sqlc.LockersLockStatusLocked,
				Id:     lock.LockerId,
			},
		); err != nil {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		// send update to broker
		_, httpResp, err = s.AdafruitClient.DataApi.CreateData(
			context.Background(),
			s.Config.AdafruitUsername,
			lock.Feedkey,
			swagger.Datum{
				Value: LOCKED,
			},
		)

		if err != nil || httpResp.StatusCode != http.StatusOK {
			log.Print(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, LockResponse{
			Status: "Locked",
		})

	} else {

		return c.JSON(http.StatusBadRequest, LockResponse{
			Status: "Bad request",
		})
	}
}

func (s *Server) createLocker(c echo.Context) error {

	var req CreateLockerRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	// Execute the query.
	// Get the email from the set context from jwt middleware
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.Payload)
	email := claims.Email

	params := db.CreateLockerParams{
		UserEmail:    email,
		LockerNumber: req.LockerNumber,
		Location:     req.Location,
		NfcSig:       req.NFCSig,
	}

	_, err := s.Store.ExecCreateLockerTx(
		context.Background(),
		params,
	)

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// create new feed on adafruit
	//_, httpResp, err := s.AdafruitClient.FeedsApi.CreateFeed(
	//	context.Background(),
	//	s.Config.AdafruitUsername,
	//	swagger.Feed{
	//		Name: "locker" + strconv.Itoa(int(req.LockerNumber)) + "-lock",
	//	},
	//	&swagger.FeedsApiCreateFeedOpts{},
	//)

	//if err != nil || httpResp.StatusCode != http.StatusOK {
	//	log.Print(err)
	//	return c.JSON(http.StatusInternalServerError, err)
	//}

	return c.JSON(http.StatusOK, CreateLockerResponse{
		Status: "Locker created",
	})

}

func (s *Server) registerLocker(c echo.Context) error {

	var req RegisterLockerRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	// Execute the query.
	// Get the email from the set context from jwt middleware
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.Payload)
	email := claims.Email

	params := db.RegisterLockerParams{
		NFCSig:    req.NFCSig,
		UserEmail: email,
	}

	_, err := s.Store.ExecRegisterLockerTx(
		context.Background(),
		params,
	)

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, RegisterLockerResponse{
		Status: "Locker registered",
	})

}

func (s *Server) removeLocker(c echo.Context) error {
	panic("TODO")
}

func (s *Server) getAllLockersInfo(c echo.Context) error {
	panic("TODO")
}
