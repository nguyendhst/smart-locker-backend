package api

import (
	"context"
	"net/http"

	"smart-locker/backend/db"
	"smart-locker/backend/db/sqlc"

	"github.com/antihax/optional"
	"github.com/labstack/echo/v4"
	swagger "github.com/nguyendhst/adafruit-go-client-v2"
	log "github.com/rs/zerolog/log"
)

var ()

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
	}

	log.Print("lock stat", lock.RegisteredLockStatus)

	//if (resp.Value == "0" && lock.RegisteredLockStatus != "unlocked") || (resp.Value == "1" && lock.RegisteredLockStatus != "locked") {
	//	defer syncLockerState(s, lock, resp.Value)
	//	return c.JSON(http.StatusConflict, UnlockResponse{
	//		Status: "State conflict",
	//	})
	//}

	if resp.Value == "0" {
		return c.JSON(http.StatusAlreadyReported, UnlockResponse{
			Status: "Already unlocked",
		})
	}

	// update the DB

	_, err = s.Store.ExecUpdateLockStatusTx(
		context.Background(),
		db.UpdateLockStatusParams{
			Status: sqlc.LockersLockStatusUnlocked,
			Id:     lock.LockerId,
		},
	)

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	_, httpResp, err = s.AdafruitClient.DataApi.CreateData(
		context.Background(),
		s.Config.AdafruitUsername,
		lock.Feedkey,
		swagger.Datum{
			Value: "0",
		},
	)

	if err != nil || httpResp.StatusCode != http.StatusOK {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, UnlockResponse{
		Status: "Unlocked",
	})
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
	}

	log.Print("Lock stat:", lock.RegisteredLockStatus)

	//if (resp.Value == "1" && lock.RegisteredLockStatus != "locked") || (resp.Value == "0" && lock.RegisteredLockStatus != "unlocked") {
	//	defer syncLockerState(s, lock, resp.Value)
	//	return c.JSON(http.StatusConflict, LockResponse{
	//		Status: "State conflict",
	//	})
	//}

	if resp.Value == "1" {
		return c.JSON(http.StatusAlreadyReported, LockResponse{
			Status: "Already locked",
		})
	}

	_, err = s.Store.ExecUpdateLockStatusTx(
		context.Background(),
		db.UpdateLockStatusParams{
			Status: sqlc.LockersLockStatusLocked,
			Id:     lock.LockerId,
		},
	)

	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	_, httpResp, err = s.AdafruitClient.DataApi.CreateData(
		context.Background(),
		s.Config.AdafruitUsername,
		lock.Feedkey,
		swagger.Datum{
			Value: "1",
		},
	)

	if err != nil || httpResp.StatusCode != http.StatusOK {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, LockResponse{
		Status: "Locked",
	})
}

func syncLockerState(s *Server, lock db.GetFeedByNFCSigResult, state string) {
	_, httpResp, err := s.AdafruitClient.DataApi.CreateData(
		context.Background(),
		s.Config.AdafruitUsername,
		lock.Feedkey,
		swagger.Datum{
			Value: state,
		},
	)

	if err != nil || httpResp.StatusCode != http.StatusOK {
		log.Print(err)
	}
}
