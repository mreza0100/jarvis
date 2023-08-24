package historyport

import "github.com/mreza0100/jarvis/internal/models"

type History interface {
	SavePrompt(prompt *models.Prompt)
	SaveResponse(prompt *models.Response)
}
