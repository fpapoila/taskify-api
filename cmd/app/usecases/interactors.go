package usecases

type (
	InteractorFactory interface {
		CreateGitHubInteractor() GitHubInteractor
	}

	interactorFactory struct {

	}

	GitHubInteractor interface {
		SignUp()
	}
)

func (f *interactorFactory) CreateGitHubInteractor() GitHubInteractor {
	return NewGitHubUseCase()
}

func NewFactory() InteractorFactory {
	return &interactorFactory{}
}
