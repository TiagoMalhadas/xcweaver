package antipode

import (
	"context"
	"fmt"
)

type Datastore_type interface {
	write(context.Context, string, AntiObj) error
	read(context.Context, string) (AntiObj, error)
	barrier(context.Context, []WriteIdentifier, string) error
}

type AntiObj struct {
	Version string
	Lineage []WriteIdentifier
}

type WriteIdentifier struct {
	Dtstid  string
	Key     string
	Version string
}

type contextKey string

func Write(ctx context.Context, datastoreType Datastore_type, datastore_ID string, key string, value string) (context.Context, error) {

	//span := trace.SpanFromContext(ctx)
	//if span.SpanContext().IsValid() {

	//	var lineage []string

	// Add the array attribute to the span
	//	span.SetAttributes(attribute.StringSlice("lineage", lineage))

	//}

	//extract lineage from ctx
	lineage := ctx.Value(contextKey("lineage")).([]WriteIdentifier)

	if lineage == nil {
		err := fmt.Errorf("Lineage not found inside context")
		return ctx, err
	}

	//update lineage
	lineage = append(lineage, WriteIdentifier{Dtstid: datastore_ID, Key: key, Version: value})

	//initialize AntiObj
	obj := AntiObj{value, lineage}

	err := datastoreType.write(ctx, key, obj)

	if err != nil {
		return ctx, err
	}

	//update ctx with the updated lineage
	ctx = context.WithValue(ctx, contextKey("lineage"), lineage)

	return ctx, nil
}

func Read(ctx context.Context, datastoreType Datastore_type, key string) (string, []WriteIdentifier, error) {

	obj, err := datastoreType.read(ctx, key)

	return obj.Version, obj.Lineage, err
}

func Barrier(ctx context.Context, datastoreType Datastore_type, datastore_ID string) error {
	//extract lineage from ctx
	lineage := ctx.Value(contextKey("lineage")).([]WriteIdentifier)

	if lineage == nil {
		err := fmt.Errorf("Lineage not found inside context")
		return err
	}

	return datastoreType.barrier(ctx, lineage, datastore_ID)
}

func Transfer(ctx context.Context, lineage []WriteIdentifier) (context.Context, error) {
	//extract lineage from ctx
	oldLineage := ctx.Value(contextKey("lineage")).([]WriteIdentifier)

	if oldLineage == nil {
		err := fmt.Errorf("Lineage not found inside context")
		return ctx, err
	}

	newLineage := append(oldLineage, lineage...)

	ctx = context.WithValue(ctx, contextKey("lineage"), newLineage)

	return ctx, nil
}

func GetLineage(ctx context.Context) ([]WriteIdentifier, error) {
	lineage := ctx.Value(contextKey("lineage"))

	if lineage == nil {
		err := fmt.Errorf("Lineage not found inside context")
		return []WriteIdentifier{}, err
	}

	return lineage.([]WriteIdentifier), nil
}

func InitCtx(ctx context.Context) context.Context {
	var lineage []WriteIdentifier = []WriteIdentifier{}

	ctx = context.WithValue(ctx, contextKey("lineage"), lineage)

	return ctx
}
