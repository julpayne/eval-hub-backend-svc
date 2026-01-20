package api

type Query map[string]string

type Storage interface {
	CreateEvaluationJOb(evaluation *EvaluationJobResource) error
	GetEvaluation(id string) (*EvaluationJobResource, error)
	GetEvaluations(query Query) (*EvaluationJobResourceList, error)
	UpdateEvaluation(evaluation *EvaluationJobResource) error
	DeleteEvaluation(id string) error

	CreateCollection(collection *CollectionResource) error
	GetCollection(id string) (*CollectionResource, error)
	GetCollections(query Query) (*CollectionResourceList, error)
	UpdateCollection(collection *CollectionResource) error
	DeleteCollection(id string) error
}
