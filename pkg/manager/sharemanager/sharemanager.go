package sharemanager

import (
	"later.co/pkg/body"
	"later.co/pkg/later/share"
	"later.co/pkg/later/usercontent"
	"later.co/pkg/repository/sharerepo"
	"later.co/pkg/repository/usercontentrepo"
)

// Create creates a share and usercontent
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
