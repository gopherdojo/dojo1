package typing

type Manager struct {
	questions       questions
	currentQuestion string
}

func (m *Manager) AddQuestions(qs []string) {
	for _, q := range qs {
		m.questions.addQuestion(q)
	}
}

func (m *Manager) AddQuestion(q string) {
	m.questions.addQuestion(q)
}

func (m *Manager) SetNewQuestion() bool {
	if q, ok := m.questions.shiftQuestion(); ok {
		m.currentQuestion = q
		return true
	}
	return false
}

func (m *Manager) GetCurrentQuestion() string {
	return m.currentQuestion
}

func (m *Manager) ValidateAnswer(a string) bool {
	return a == m.currentQuestion
}

func (m *Manager) HasQuestion() bool {
	return m.questions.hasQuestion()
}

func NewManager() *Manager {
	return &Manager{
		questions:       []string{},
		currentQuestion: "",
	}
}

type questions []string

func (qs *questions) addQuestion(q string) {
	*qs = append(*qs, q)
}

func (qs *questions) shiftQuestion() (string, bool) {
	if len(*qs) <= 0 {
		return "", false
	}
	q := (*qs)[0]
	*qs = (*qs)[1:]
	return q, true
}

func (qs *questions) hasQuestion() bool {
	return len(*qs) > 0
}
