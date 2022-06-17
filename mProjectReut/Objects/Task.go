package Objects

type Task struct {
	Type        string `json:"type" binding:"required"`
	Status      Status `json:"status"`
	Description string `json:"description" binding:"required"`
	Size        Size   `json:"size" binding:"required"`
	Course      string `json:"course" binding:"required"`
	DueDate     string `json:"dueDate" binding:"required"`
	Details     string `json:"details" binding:"required"`
}

type Chore struct {
	Type        string `json:"type" binding:"required"`
	Status      Status `json:"status"`
	Description string `json:"description" binding:"required"`
	Size        Size   `json:"size" binding:"required"`
}

type HomeWork struct {
	Type    string `json:"type" binding:"required"`
	Status  Status `json:"status"`
	Course  string `json:"course" binding:"required"`
	DueDate string `json:"dueDate" binding:"required"`
	Details string `json:"details" binding:"required"`
}
