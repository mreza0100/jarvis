package historyport

import "github.com/mreza0100/gptjarvis/internal/models"

type History interface {
	SavePrompt(prompt *models.Prompt)
	SaveResponse(prompt *models.Response)
}
