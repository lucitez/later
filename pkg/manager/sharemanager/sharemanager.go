package sharemanager

import (
	"later.co/pkg/body"
	"later.co/pkg/later/share"
	"later.co/pkg/later/usercontent"
	"later.co/pkg/repository/sharerepo"
	"later.co/pkg/repository/usercontentrepo"
)

// CreateMultiple creates multiple shares from multiple bodies
func CreateMultiple(createBodies []body.ShareCreateBody) ([]share.Share, error) {
	shares := []share.Share{}

	for _, createBody := range createBodies {
		share, err := Create(createBody)

		if err != nil {
			return nil, err
		}

		shares = append(shares, *share)
	}

	return shares, nil
}

// Create creates a share and usercontent
// Should probably do the notification stuff here
// TODO Two Goroutines:
// Update _body.Content.shares_ total by getting count(shares distinct on user_id with this content_id)
// Send Push notification if user has signed up <-- maybe move this to usercontent
func Create(body body.ShareCreateBody) (*share.Share, error) {
	share, err := share.New(
		body.Content.ID,
		body.SenderUserID,
		body.RecipientUserID)

	if err != nil {
		return nil, err
	}

	share, err = sharerepo.Insert(share)

	if err != nil {
		return nil, err
	}

	usercontent, err := usercontent.New(
		share.ID,
		body.Content.ID,
		body.Content.ContentType,
		body.RecipientUserID,
		body.SenderUserID)

	if err != nil {
		return share, err
	}

	_, err = usercontentrepo.Insert(usercontent)

	if err != nil {
		return nil, err
	}

	return share, nil
}
