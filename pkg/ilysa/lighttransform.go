package ilysa

type LightTransformer func(Light) Light
type LightIDTransformer func(id LightID) LightIDSet

type LightIDTransformable interface {
	LightIDTransform(LightIDTransformer) Light
}
