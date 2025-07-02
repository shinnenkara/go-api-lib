package dtos

type EntityReference struct {
	ID string `uri:"id" json:"id" binding:"required,uuid"`
}
