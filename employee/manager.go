package employee

// Manager is a struct for manager
// It embeds Employee
// This means that Manager has all the fields and methods of Employee
type Manager struct {
	Employee
}

// NewManager creates a new manager
// It returns a pointer to the manager
// Creational method
func NewManager() *Manager {
	// Create a new manager
	// Set the name to "Manager"
	// Set the salary to 1000
	return &Manager{
		Employee: Employee{
			Name:   "Manager",
			Salary: 1000,
		},
	}
}
