package services

// Storage struct holds password slices
// Used Ref from slides to use slices
type Storage struct {
	Checked   []string
	Generated []string
}

// SavePassword adds a generated password
func (s *Storage) SavePassword(pwd string) {
	s.Generated = append(s.Generated, pwd)
}

// GetGenerated returns all generated passwords
func (s *Storage) GetGenerated() []string {
	return s.Generated
}
