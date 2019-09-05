package bootstrap

import (
	"context"
	konvoy_cp "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/config/app/konvoy-cp"
	"github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/apis/mesh"
	core_model "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/model"
	core_store "github.com/Kong/konvoy/components/konvoy-control-plane/pkg/core/resources/store"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bootstrap", func() {

	It("should create default mesh", func() {
		// given
		cfg := konvoy_cp.DefaultConfig()

		// when control plane is started
		rt, err := Bootstrap(cfg)
		ch := make(chan struct{})
		go func() {
			defer GinkgoRecover()
			err := rt.Start(ch)
			Expect(err).ToNot(HaveOccurred())
		}()
		Expect(err).ToNot(HaveOccurred())

		// then wait until resource is created
		resManager := rt.ResourceManager()
		Eventually(func() error {
			getOpts := core_store.GetByKey(core_model.DefaultNamespace, core_model.DefaultMesh, core_model.DefaultMesh)
			return resManager.Get(context.Background(), &mesh.MeshResource{}, getOpts)
		}).Should(Succeed())

		// when
		getOpts := core_store.GetByKey(core_model.DefaultNamespace, core_model.DefaultMesh, core_model.DefaultMesh)
		defaultMesh := mesh.MeshResource{}
		err = resManager.Get(context.Background(), &defaultMesh, getOpts)
		Expect(err).ToNot(HaveOccurred())

		// then
		meshMeta := defaultMesh.GetMeta()
		Expect(meshMeta.GetName()).To(Equal("default"))
		Expect(meshMeta.GetMesh()).To(Equal("default"))
		Expect(meshMeta.GetNamespace()).To(Equal("default"))
	})

	It("should skip creating mesh if one already exist", func() {
		// given
		cfg := konvoy_cp.DefaultConfig()
		runtime, err := buildRuntime(cfg)
		Expect(err).ToNot(HaveOccurred())

		// when
		Expect(createDefaultMesh(runtime)).To(Succeed())

		// then mesh exists
		getOpts := core_store.GetByKey(core_model.DefaultNamespace, core_model.DefaultMesh, core_model.DefaultMesh)
		err = runtime.ResourceManager().Get(context.Background(), &mesh.MeshResource{}, getOpts)
		Expect(err).ToNot(HaveOccurred())

		// when createDefaultMesh is called once mesh already exist
		err = createDefaultMesh(runtime)

		// then
		Expect(err).ToNot(HaveOccurred())
	})

})