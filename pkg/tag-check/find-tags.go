package tag_check

import "bot-test/pkg/models"

type FindTagsResult struct {
	Tag     *models.Tag
	SubText string
}

func FindTags(text string, tags []*models.Tag) ([]*FindTagsResult, error) {
	result := make([]*FindTagsResult, 0)

	return result, nil
}
