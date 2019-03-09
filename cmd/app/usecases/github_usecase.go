package usecases

type (
	gitHubUseCase struct {
	}
)

func NewGitHubUseCase() GitHubInteractor {
	return &gitHubUseCase{}
}

func (gu *gitHubUseCase) SignUp() {

}
