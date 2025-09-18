package domain

type ModelHasPermission struct {
	Permission *Permission
	ModelType  string
	ModelId    int64
}

type ModelHasRole struct {
	Role      *Role
	ModelType string
	ModelId   int64
}
