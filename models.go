package main

type Schedule struct {
	ID                  int    `json:"id"`
	PersonID            int    `json:"person_id"`
	ScheduleCategoryID  int    `json:"schedule_cat_id"`
	ScheduleFrequencyID int    `json:"schedule_frequency_id"`
	ScheduleTime        string `json:"schedule_time"`
	StartDate           string `json:"start_date,omitempty"`
	EndDate             string `json:"end_date,omitempty"`
	Desc                string `json:"desc"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	Deleted             bool   `json:"deleted"`
}

type ScheduleCategory struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
}

type ScheduleFrequency struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
}

type PersonParentLink struct {
	ID        int    `json:"id"`
	PersonID  int    `json:"person_id"`
	ParentID  int    `json:"parent_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
}

type Person struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DOB       string `json:"dob"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
}

type Note struct {
	ID        int    `json:"id"`
	TableID   string `json:"table_id"`
	TableName string `json:"table_name"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
}

type Todo struct {
	ID        int    `json:"id"`
	PersonID  int    `json:"person_id"`
	Content   string `json:"content"`
	Priority  int    `json:"priority"`
	Complete  bool   `json:"complete"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
}
