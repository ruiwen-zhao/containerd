package server

import (
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/containers"
	"github.com/containerd/containerd/oci"
	prototypes "github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type fakeContainer struct {
	spec *oci.Spec
}

func (c *fakeContainer) WithSpec(spec *oci.Spec) {
	c.spec = spec
}

func (*fakeContainer) ID() string {
	return "fakeContainerID"
}

func (c *fakeContainer) Info(ctx context.Context, opts ...containerd.InfoOpts) (containers.Container, error) {
	return containers.Container{}, errors.New("Not implemented")
}

func (*fakeContainer) Delete(context.Context, ...containerd.DeleteOpts) error {
	return errors.New("Not implemented")
}

func (*fakeContainer) NewTask(context.Context, cio.Creator, ...containerd.NewTaskOpts) (containerd.Task, error) {
	return nil, errors.New("Not implemented")
}

func (c *fakeContainer) Spec(context.Context) (*oci.Spec, error) {
	return c.spec, nil
}

func (*fakeContainer) Task(context.Context, cio.Attach) (containerd.Task, error) {
	return nil, errors.New("Not implemented")
}

func (*fakeContainer) Image(context.Context) (containerd.Image, error) {
	return nil, errors.New("Not implemented")
}

func (*fakeContainer) Labels(context.Context) (map[string]string, error) {
	return nil, errors.New("Not implemented")
}

func (*fakeContainer) SetLabels(context.Context, map[string]string) (map[string]string, error) {
	return nil, errors.New("Not implemented")
}

func (*fakeContainer) Extensions(context.Context) (map[string]prototypes.Any, error) {
	return nil, errors.New("Not implemented")
}

func (*fakeContainer) Update(context.Context, ...containerd.UpdateContainerOpts) error {
	return errors.New("Not implemented")
}

func (*fakeContainer) Checkpoint(context.Context, string, ...containerd.CheckpointOpts) (containerd.Image, error) {
	return nil, errors.New("Not implemented")
}
