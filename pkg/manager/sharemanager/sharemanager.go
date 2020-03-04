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
// TODO how do we want to increment total # shares? maybe just unique user_ids on shares table per content_id?
// TODO how do we want to send notifications? get user after creating the user content, if they have signed up, send notif, else send sms
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
