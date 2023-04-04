package api

import (
	"context"
	"net/http"

	"smart-locker/backend/db"

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

	feed, err := s.Store.ExecGetFeedByNFCSigTx(
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
		feed.Feedkey,
		&swagger.DataApiLastDataOpts{
			Include: optional.NewString("value,created_at"),
		},
	)
	if err != nil || httpResp.StatusCode != http.StatusOK {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	} else if resp.Value == "0" {
		return c.JSON(http.StatusAlreadyReported, UnlockResponse{
			Status: "Already unlocked",
		})
	}

	_, httpResp, err = s.AdafruitClient.DataApi.CreateData(
		context.Background(),
		s.Config.AdafruitUsername,
		feed.Feedkey,
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

	feed, err := s.Store.ExecGetFeedByNFCSigTx(
		context.Background(),
		params,
	)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// check if already locked
	resp, httpResp, err := s.AdafruitClient.DataApi.LastData(
		context.Background(),
		s.Config.AdafruitUsername,
		feed.Feedkey,
		&swagger.DataApiLastDataOpts{
			Include: optional.NewString("value,created_at"),
		},
	)
	if err != nil || httpResp.StatusCode != http.StatusOK {
		log.Print(err)
		return c.JSON(http.StatusInternalServerError, err)
	} else if resp.Value == "1" {
		return c.JSON(http.StatusAlreadyReported, LockResponse{
			Status: "Already locked",
		})
	}

	_, httpResp, err = s.AdafruitClient.DataApi.CreateData(
		context.Background(),
		s.Config.AdafruitUsername,
		feed.Feedkey,
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
