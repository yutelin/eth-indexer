package store

import (
	"math/big"
	"os"
	"testing"

	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
	"github.com/maichain/eth-indexer/common"
	"github.com/maichain/eth-indexer/model"
	"github.com/maichain/eth-indexer/store/account"
	"github.com/maichain/eth-indexer/store/block_header"
	"github.com/maichain/eth-indexer/store/transaction"
	"github.com/maichain/eth-indexer/store/transaction_receipt"
	"github.com/maichain/mapi/base/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager Test", func() {
	var (
		mysql *test.MySQLContainer
		db    *gorm.DB
	)
	BeforeSuite(func() {
		var err error
		mysql, err = test.NewMySQLContainer("quay.io/amis/eth-indexer-db-migration")
		Expect(mysql).ShouldNot(BeNil())
		Expect(err).Should(Succeed())
		Expect(mysql.Start()).Should(Succeed())

		db, err = gorm.Open("mysql", mysql.URL)
		Expect(err).Should(Succeed())
		Expect(db).ShouldNot(BeNil())

		db.LogMode(os.Getenv("ENABLE_DB_LOG_IN_TEST") != "")
	})

	AfterSuite(func() {
		mysql.Stop()
	})

	BeforeEach(func() {
		db.Table(block_header.TableName).Delete(&model.Header{})
		db.Table(transaction.TableName).Delete(&model.Transaction{})
		db.Table(transaction_receipt.TableName).Delete(&model.Receipt{})
		db.Table(account.NameAccounts).Delete(&model.Account{})
		db.Table(account.NameERC20).Delete(&model.ERC20{})
		db.Table(account.NameStateBlocks).Delete(&model.StateBlock{})
	})

	Context("InsertBlock()", func() {
		It("should be ok", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			header := &types.Header{
				Number: big.NewInt(10),
			}
			block := types.NewBlock(header, nil, []*types.Header{
				header,
			}, []*types.Receipt{
				types.NewReceipt([]byte{}, false, 0),
			})

			err = manager.InsertBlock(block, nil)
			Expect(err).Should(Succeed())

			By("insert the same block again, should be ok")
			err = manager.InsertBlock(block, nil)
			Expect(err).Should(Succeed())
		})

		It("saves TD", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			header := &types.Header{
				Number: big.NewInt(100),
			}
			block := types.NewBlock(header, nil, []*types.Header{
				header,
			}, []*types.Receipt{
				types.NewReceipt([]byte{}, false, 0),
			})

			err = manager.InsertTd(block, new(big.Int).SetInt64(123456789))
			Expect(err).Should(Succeed())

			err = manager.InsertTd(block, new(big.Int).SetInt64(123456789))
			Expect(common.DuplicateError(err)).Should(BeTrue())
		})

		It("failed due to wrong signer", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			header := &types.Header{
				Number: big.NewInt(11),
			}
			block := types.NewBlock(header, []*types.Transaction{
				types.NewTransaction(0, gethCommon.Address{}, gethCommon.Big0, 0, gethCommon.Big0, []byte{}),
			}, []*types.Header{
				header,
			}, []*types.Receipt{
				types.NewReceipt([]byte{}, false, 0),
			})

			err = manager.InsertBlock(block, nil)
			Expect(err).Should(Equal(common.ErrWrongSigner))
		})
	})

	Context("GetHeaderByNumber()", func() {
		It("gets the right header", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			block1 := types.NewBlockWithHeader(&types.Header{
				Number: big.NewInt(100),
			})
			block2 := types.NewBlockWithHeader(&types.Header{
				Number: big.NewInt(99),
			})
			err = manager.InsertBlock(block1, nil)
			Expect(err).Should(Succeed())
			err = manager.InsertBlock(block2, nil)
			Expect(err).Should(Succeed())

			header, err := manager.GetHeaderByNumber(100)
			Expect(err).Should(Succeed())
			Expect(header).Should(Equal(common.Header(block1)))

			header, err = manager.GetHeaderByNumber(99)
			Expect(err).Should(Succeed())
			Expect(header).Should(Equal(common.Header(block2)))

			header, err = manager.GetHeaderByNumber(199)
			Expect(common.NotFoundError(err)).Should(BeTrue())
		})
	})

	Context("GetTd()", func() {
		It("gets the block TD", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			block1 := types.NewBlockWithHeader(&types.Header{
				Number: big.NewInt(100),
			})
			block2 := types.NewBlockWithHeader(&types.Header{
				Number: big.NewInt(99),
			})
			err = manager.InsertTd(block1, new(big.Int).SetInt64(123456789))
			Expect(err).Should(Succeed())
			err = manager.InsertTd(block2, new(big.Int).SetInt64(987654321))
			Expect(err).Should(Succeed())

			td, err := manager.GetTd(block1.Hash().Bytes())
			Expect(err).Should(Succeed())
			Expect(td).Should(Equal(&model.TotalDifficulty{
				Block: block1.Number().Int64(),
				Hash:  block1.Hash().Bytes(),
				Td:    "123456789",
			}))

			td, err = manager.GetTd(block2.Hash().Bytes())
			Expect(err).Should(Succeed())
			Expect(td).Should(Equal(&model.TotalDifficulty{
				Block: block2.Number().Int64(),
				Hash:  block2.Hash().Bytes(),
				Td:    "987654321",
			}))
		})
	})

	Context("LatestHeader()", func() {
		It("gets the latest header", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			block1 := types.NewBlockWithHeader(&types.Header{
				Number: big.NewInt(100),
			})
			block2 := types.NewBlockWithHeader(&types.Header{
				Number: big.NewInt(99),
			})
			err = manager.InsertBlock(block1, nil)
			Expect(err).Should(Succeed())
			err = manager.InsertBlock(block2, nil)
			Expect(err).Should(Succeed())

			header, err := manager.LatestHeader()
			Expect(err).Should(Succeed())
			Expect(header).Should(Equal(common.Header(block1)))
		})
	})

	Context("UpdateState()", func() {
		It("should be ok", func() {
			accountStore := account.NewWithDB(db)
			erc20 := &model.ERC20{
				Address:     gethCommon.HexToAddress("0x01").Bytes(),
				Code:        common.HexToBytes("0x02"),
				TotalSupply: "1000000",
				Name:        "erc20",
				Decimals:    18,
			}
			err := accountStore.InsertERC20(erc20)
			Expect(err).Should(Succeed())
			tmp := model.ERC20Storage{
				Address: erc20.Address,
			}
			Expect(db.HasTable(tmp)).Should(BeTrue())

			manager, err := NewManager(db)
			Expect(err).Should(BeNil())
			block1 := types.NewBlockWithHeader(&types.Header{
				Number: big.NewInt(100),
				Root:   gethCommon.StringToHash("1234567890"),
			})

			value := gethCommon.HexToHash("0x0a")
			key := gethCommon.HexToHash("0x0b")
			dump := map[string]state.DumpDirtyAccount{
				"fffffffffff": {
					Nonce:   100,
					Balance: "101",
				},
				// contract
				common.BytesToHex(erc20.Address): {
					Nonce:   900,
					Balance: "901",
					Storage: map[string]string{
						key.Hex(): value.Hex(),
					},
				},
			}
			err = manager.UpdateState(block1, dump)
			Expect(err).Should(Succeed())

			By("update the same state again, should be ok")
			err = manager.UpdateState(block1, dump)
			Expect(err).Should(Succeed())

			s, err := account.NewWithDB(db).FindERC20Storage(gethCommon.BytesToAddress(erc20.Address), key, block1.Number().Int64())
			Expect(err).Should(Succeed())
			Expect(s.Address).Should(Equal(erc20.Address))
			Expect(s.Key).Should(Equal(key.Bytes()))
			Expect(s.Value).Should(Equal(value.Bytes()))
		})
	})

	Context("LatestStateBlock()", func() {
		It("gets the latest state block", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			block1 := types.NewBlockWithHeader(&types.Header{
				Number: big.NewInt(100),
				Root:   gethCommon.StringToHash("1234567890"),
			})

			err = manager.UpdateState(block1, nil)
			Expect(err).Should(Succeed())

			block, err := manager.LatestStateBlock()
			Expect(err).Should(Succeed())
			Expect(block.Number).Should(Equal(block1.Number().Int64()))
		})
	})

	Context("DeleteDataFromBlock()", func() {
		It("deletes data from a block number", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			for i := int64(100); i < 120; i++ {
				block := types.NewBlockWithHeader(&types.Header{
					Number: big.NewInt(i),
					Root:   gethCommon.StringToHash("1234567890"),
				})
				err := manager.InsertBlock(block, nil)
				Expect(err).Should(Succeed())

				err = manager.UpdateState(block, nil)
				Expect(err).Should(Succeed())
			}
			state, err := manager.LatestStateBlock()
			Expect(err).Should(Succeed())
			Expect(state.Number).Should(Equal(int64(119)))

			manager.DeleteStateFromBlock(int64(110))
			state, err = manager.LatestStateBlock()
			Expect(err).Should(Succeed())
			Expect(state.Number).Should(Equal(int64(109)))
		})
	})

	Context("UpdateBlock()", func() {
		It("changes data for a block number", func() {
			manager, err := NewManager(db)
			Expect(err).Should(BeNil())

			blockNum := int64(654321)
			block := types.NewBlock(
				&types.Header{
					Number: big.NewInt(blockNum),
					Root:   gethCommon.StringToHash("1234567890"),
				}, []*types.Transaction{}, nil, []*types.Receipt{})

			newBlock := types.NewBlock(
				&types.Header{
					Number: big.NewInt(blockNum),
					Root:   gethCommon.StringToHash("9876543210"),
				}, []*types.Transaction{}, nil, []*types.Receipt{})

			Expect(block.Hash()).ShouldNot(Equal(newBlock.Hash()))

			err = manager.InsertBlock(block, nil)
			Expect(err).Should(Succeed())

			hdr, err := manager.GetHeaderByNumber(blockNum)
			Expect(err).Should(Succeed())
			Expect(hdr.Hash).Should(Equal(block.Hash().Bytes()))

			err = manager.UpdateBlock(newBlock, nil, nil)
			Expect(err).Should(Succeed())

			hdr, err = manager.GetHeaderByNumber(blockNum)
			Expect(err).Should(Succeed())
			Expect(hdr.Hash).Should(Equal(newBlock.Hash().Bytes()))
		})
	})
})

func TestStore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Store Test")
}
