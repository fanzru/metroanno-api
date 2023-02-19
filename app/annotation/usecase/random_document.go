package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"metroanno-api/app/annotation/domain/models"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (a *AnnotationsApp) RandomDocuments(ctx echo.Context) (*models.Document, error) {
	ID, err := strconv.ParseInt(fmt.Sprintf("%v", ctx.Get("user_id")), 10, 64)
	if err != nil {
		return nil, err
	}

	user, err := a.AnnotationsRepo.GetUserByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	documentDoneUser, err := a.AnnotationsRepo.GetAllDocumentDoneUser(ctx, ID)
	if err != nil {
		return nil, err
	}
	documents, err := a.AnnotationsRepo.GetAllDocuments(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("-------------------------------- TEST")
	b, _ := json.MarshalIndent(documents, "", "  ")
	fmt.Println(string(b))

	// remove done document
	newArrDocuments := []models.Document{}
	for _, doc := range documents {
		if doc.DoneNumberOfAnnotators != doc.MinNumberOfAnnotators {
			log.Println("-------------------------------- TRUE")
			found := false
			for _, document := range *documentDoneUser {
				if document.DocumentID == doc.Id {
					found = true
					break
				}
			}
			if !found {
				newArrDocuments = append(newArrDocuments, doc)
			}
		}
	}

	log.Println("-------------------------------- TEST : ", len(newArrDocuments))

	if len(newArrDocuments) == 0 {
		_, err = a.AnnotationsRepo.UpdateUsersById(ctx, 0, user.Id)
		if err != nil {
			return nil, err
		}
		return nil, ErrNotHaveDocuments
	}
	log.Println("-------------------------------- TEST")

	// sort by done number of documents
	sort.SliceStable(newArrDocuments, func(i, j int) bool {
		return newArrDocuments[i].DoneNumberOfAnnotators < newArrDocuments[j].DoneNumberOfAnnotators
	})

	_, err = a.AnnotationsRepo.UpdateUsersById(ctx, newArrDocuments[0].Id, user.Id)
	if err != nil {
		return nil, err
	}

	log.Println("-------------------------------- TEST")

	return &newArrDocuments[0], nil
}
