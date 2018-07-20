package leveldb_test

import (
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/republicprotocol/republic-go/leveldb"

	"github.com/republicprotocol/republic-go/identity"
	"github.com/republicprotocol/republic-go/swarm"
	"github.com/republicprotocol/republic-go/testutils"
)

var _ = Describe("Swarm storage", func() {

	multiAddresses := make([]identity.MultiAddress, 100)
	dbFolder := "./tmp/"
	dbFile := dbFolder + "db"

	BeforeEach(func() {
		for i := 0; i < 100; i++ {
			var err error
			multiAddresses[i], err = testutils.RandomMultiAddress()
			Expect(err).ShouldNot(HaveOccurred())
		}
	})

	AfterEach(func() {
		os.RemoveAll(dbFolder)
	})

	Context("when pruning data", func() {
		It("should not retrieve expired data", func() {
			db := newDB(dbFile)
			swarmMultiAddressTable := NewSwarmMultiAddressTable(db, 2*time.Second)

			// Put the multiAddresses into the table and attempt to retrieve
			for i := 0; i < len(multiAddresses); i++ {
				_, err := swarmMultiAddressTable.PutMultiAddress(multiAddresses[i], 1)
				Expect(err).ShouldNot(HaveOccurred())
			}
			for i := 0; i < len(multiAddresses); i++ {
				multiAddr, nonce, err := swarmMultiAddressTable.MultiAddress(multiAddresses[i].Address())
				Expect(err).ShouldNot(HaveOccurred())
				Expect(multiAddr.String()).Should(Equal(multiAddresses[i].String()))
				Expect(nonce).Should(Equal(uint64(1)))
			}

			// Sleep and then prune to expire the data
			time.Sleep(2 * time.Second)
			swarmMultiAddressTable.Prune()

			// All data should have expired so we should not get any data back
			multiAddrsIter, err := swarmMultiAddressTable.MultiAddresses()
			Expect(err).ShouldNot(HaveOccurred())
			defer multiAddrsIter.Release()
			multiAddrs, nonces, err := multiAddrsIter.Collect()
			Expect(err).ShouldNot(HaveOccurred())

			Expect(multiAddrs).Should(HaveLen(0))
			Expect(nonces).Should(HaveLen(0))
		})
	})

	Context("when iterating through out of range data", func() {
		It("should trigger an out of range error", func() {
			db := newDB(dbFile)
			swarmMultiAddressTable := NewSwarmMultiAddressTable(db, 2*time.Second)

			// Put the multiAddresses into the table and attempt to retrieve
			for i := 0; i < len(multiAddresses); i++ {
				_, err := swarmMultiAddressTable.PutMultiAddress(multiAddresses[i], 1)
				Expect(err).ShouldNot(HaveOccurred())
			}
			for i := 0; i < len(multiAddresses); i++ {
				multiAddr, nonce, err := swarmMultiAddressTable.MultiAddress(multiAddresses[i].Address())
				Expect(err).ShouldNot(HaveOccurred())
				Expect(multiAddr.String()).Should(Equal(multiAddresses[i].String()))
				Expect(nonce).Should(Equal(uint64(1)))
			}

			multiAddrsIter, err := swarmMultiAddressTable.MultiAddresses()
			defer multiAddrsIter.Release()
			for multiAddrsIter.Next() {
				_, _, err := multiAddrsIter.Cursor()
				Expect(err).ShouldNot(HaveOccurred())
			}

			// This is out of range so we should expect an error
			_, _, err = multiAddrsIter.Cursor()
			Expect(err).Should(Equal(swarm.ErrCursorOutOfRange))
		})
	})

	Context("when updating details about self", func() {
		It("should not return own multiaddress", func() {
			db := newDB(dbFile)
			swarmMultiAddressTable := NewSwarmMultiAddressTable(db, 2*time.Second)

			myMultiAddr, err := testutils.RandomMultiAddress()
			Expect(err).ShouldNot(HaveOccurred())

			nonce, err := swarmMultiAddressTable.PutSelf(myMultiAddr, 0)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(nonce).Should(Equal(uint64(1)))

			// Put other multiAddresses into the table
			for i := 0; i < len(multiAddresses); i++ {
				_, err := swarmMultiAddressTable.PutMultiAddress(multiAddresses[i], 1)
				Expect(err).ShouldNot(HaveOccurred())
			}

			// Attempt to retrieve my multiaddress using my address.
			multi, nonce, err := swarmMultiAddressTable.MultiAddress(myMultiAddr.Address())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(multi.String()).Should(Equal(myMultiAddr.String()))
			Expect(nonce).Should(Equal(uint64(1)))

			multiAddrsIter, err := swarmMultiAddressTable.MultiAddresses()
			defer multiAddrsIter.Release()
			multiAddrs, _, err := multiAddrsIter.Collect()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(multiAddrs)).Should(Equal(len(multiAddresses)))
		})
	})
})
