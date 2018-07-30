package grpc_test

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/republicprotocol/republic-go/grpc"

	"github.com/republicprotocol/republic-go/crypto"
	"github.com/republicprotocol/republic-go/dispatch"
	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/leveldb"
	"github.com/republicprotocol/republic-go/swarm"
	"golang.org/x/net/context"
)

var _ = Describe("Swarming", func() {

	var server *Server
	var service SwarmService
	var serviceMultiAddr identity.MultiAddress
	var serviceClient swarm.Client
	var serviceClientDb swarm.MultiAddressStorer
	var client swarm.Client
	var clientDb swarm.MultiAddressStorer
	var swarmer swarm.Swarmer
	var ecdsaKey crypto.EcdsaKey

	BeforeEach(func() {
		var err error

		db, err := leveldb.NewStore("./tmp/swarm.1.out", 10*time.Hour)
		Expect(err).ShouldNot(HaveOccurred())
		serviceClientDb = db.SwarmMultiAddressStore()
		serviceClient, ecdsaKey, err = newSwarmClient(serviceClientDb)
		Expect(err).ShouldNot(HaveOccurred())

		swarmer = swarm.NewSwarmer(serviceClient, serviceClientDb, 10, &ecdsaKey)
		Expect(err).ShouldNot(HaveOccurred())
		service = NewSwarmService(swarm.NewServer(swarmer, serviceClientDb, 10), time.Microsecond)
		serviceMultiAddr = serviceClient.MultiAddress()
		server = NewServer()
		service.Register(server)

		db, err = leveldb.NewStore("./tmp/swarm.2.out", 10*time.Hour)
		Expect(err).ShouldNot(HaveOccurred())
		clientDb = db.SwarmMultiAddressStore()
		client, _, err = newSwarmClient(clientDb)
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll("./tmp")
		server.Stop()
	})

	Context("when pinging a service", func() {

		It("should return the multiaddress of the service", func(done Done) {
			defer close(done)

			go func() {
				defer GinkgoRecover()

				err := server.Start("0.0.0.0:18514")
				Expect(err).ShouldNot(HaveOccurred())
			}()
			time.Sleep(time.Millisecond)

			err := client.Ping(context.Background(), serviceMultiAddr, client.MultiAddress())
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should add the client to the service storer", func(done Done) {
			defer close(done)

			go func() {
				defer GinkgoRecover()

				err := server.Start("0.0.0.0:18514")
				Expect(err).ShouldNot(HaveOccurred())
			}()
			time.Sleep(time.Millisecond)

			err := client.Ping(context.Background(), serviceMultiAddr, client.MultiAddress())
			Expect(err).ShouldNot(HaveOccurred())
			multiAddrIter, err := serviceClientDb.MultiAddresses()
			Expect(err).ShouldNot(HaveOccurred())
			multiAddrs, err := multiAddrIter.Collect()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(multiAddrs).Should(HaveLen(2))
		})

	})

	Context("when querying a service", func() {

		It("should return the multiaddress of the service close to the query", func(done Done) {
			defer close(done)

			go func() {
				defer GinkgoRecover()

				err := server.Start("0.0.0.0:18514")
				Expect(err).ShouldNot(HaveOccurred())
			}()

			err := client.Ping(context.Background(), serviceMultiAddr, client.MultiAddress())
			Expect(err).ShouldNot(HaveOccurred())

			multiAddrs, err := client.Query(context.Background(), serviceMultiAddr, client.MultiAddress().Address())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(multiAddrs).Should(HaveLen(1))
		})

		It("should error when too many requests are sent to the server", func(done Done) {
			defer close(done)

			service = NewSwarmService(swarm.NewServer(swarmer, serviceClientDb, 10), time.Second)
			serviceMultiAddr = serviceClient.MultiAddress()
			server = NewServer()
			service.Register(server)

			go func() {
				defer GinkgoRecover()

				err := server.Start("0.0.0.0:18514")
				Expect(err).ShouldNot(HaveOccurred())
			}()

			err := client.Ping(context.Background(), serviceMultiAddr, client.MultiAddress())
			Expect(err).ShouldNot(HaveOccurred())

			dispatch.CoForAll(10, func(i int) {
				defer GinkgoRecover()

				ctx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
				defer cancel()

				_, err := client.Query(ctx, serviceMultiAddr, client.MultiAddress().Address())
				Expect(err).Should(HaveOccurred())
			})
		}, 30 /* 30s timeout */)

	})
})

func newSwarmClient(db swarm.MultiAddressStorer) (swarm.Client, crypto.EcdsaKey, error) {
	ecdsaKey, err := crypto.RandomEcdsaKey()
	if err != nil {
		return nil, crypto.EcdsaKey{}, err
	}
	addr := identity.Address(ecdsaKey.Address())
	multiAddr, err := identity.NewMultiAddressFromString(fmt.Sprintf("/ip4/0.0.0.0/tcp/18514/republic/%v", addr))
	if err != nil {
		return nil, crypto.EcdsaKey{}, err
	}
	multiAddr.Nonce = 1
	signature, err := ecdsaKey.Sign(multiAddr.Hash())
	if err != nil {
		return nil, crypto.EcdsaKey{}, err
	}
	multiAddr.Signature = signature
	db.PutMultiAddress(multiAddr)
	client := NewSwarmClient(db, multiAddr.Address())
	return client, ecdsaKey, nil
}
