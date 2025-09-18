package domain

type ModelHasPermission struct {
	Permission *Permission
	ModelType  string
	ModelID    int64
}

type ModelHasRole struct {
	Role      *Role
	ModelType string
	ModelID   int64
}
