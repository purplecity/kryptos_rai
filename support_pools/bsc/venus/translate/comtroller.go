// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package translate

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ComtrollerABI is the input ABI used to generate the binding from.
const ComtrollerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"ActionProtocolPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"venusVAIMintIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIMinterVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedVAIVaultVenus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"}],\"name\":\"NewTreasuryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTreasuryGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"}],\"name\":\"NewTreasuryGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTreasuryPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"NewTreasuryPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"oldVAIController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractVAIControllerInterface\",\"name\":\"newVAIController\",\"type\":\"address\"}],\"name\":\"NewVAIController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVAIMintRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"NewVAIMintRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseInterval_\",\"type\":\"uint256\"}],\"name\":\"NewVAIVaultInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVenusVAIVaultRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVenusVAIVaultRate\",\"type\":\"uint256\"}],\"name\":\"NewVenusVAIVaultRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"VenusSpeedUpdated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setProtocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasuryGuardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newTreasuryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"_setTreasuryData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"vaiController_\",\"type\":\"address\"}],\"name\":\"_setVAIController\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVAIMintRate\",\"type\":\"uint256\"}],\"name\":\"_setVAIMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"releaseStartBlock_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReleaseAmount_\",\"type\":\"uint256\"}],\"name\":\"_setVAIVaultInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"venusSpeed\",\"type\":\"uint256\"}],\"name\":\"_setVenusSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"venusVAIRate_\",\"type\":\"uint256\"}],\"name\":\"_setVenusVAIRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"venusVAIVaultRate_\",\"type\":\"uint256\"}],\"name\":\"_setVenusVAIVaultRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractVToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractVToken\",\"name\":\"vToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractVToken[]\",\"name\":\"vTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaiMinter\",\"type\":\"address\"}],\"name\":\"distributeVAIMinterVenus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractVToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getXVSAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateVAICalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isVenus\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minReleaseAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintedVAIs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseStartBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"repayVAIGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMintedVAIOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiController\",\"outputs\":[{\"internalType\":\"contractVAIControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiMintRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaiVaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"venusSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"venusVAIVaultRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ComtrollerBin is the compiled bytecode used for deploying new contracts.
var ComtrollerBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556105e4806100326000396000f3fe60806040526004361061007b5760003560e01c8063dcfbc0c71161004e578063dcfbc0c71461019e578063e992a041146101b3578063e9c714f2146101e6578063f851a440146101fb5761007b565b806326782247146100fe578063b71d1a0c1461012f578063bb82aa5e14610174578063c1e8033414610189575b6002546040516000916001600160a01b031690829036908083838082843760405192019450600093509091505080830381855af49150503d80600081146100de576040519150601f19603f3d011682016040523d82523d6000602084013e6100e3565b606091505b505090506040513d6000823e8180156100fa573d82f35b3d82fd5b34801561010a57600080fd5b50610113610210565b604080516001600160a01b039092168252519081900360200190f35b34801561013b57600080fd5b506101626004803603602081101561015257600080fd5b50356001600160a01b031661021f565b60408051918252519081900360200190f35b34801561018057600080fd5b506101136102b0565b34801561019557600080fd5b506101626102bf565b3480156101aa57600080fd5b506101136103ba565b3480156101bf57600080fd5b50610162600480360360208110156101d657600080fd5b50356001600160a01b03166103c9565b3480156101f257600080fd5b5061016261044d565b34801561020757600080fd5b50610113610533565b6001546001600160a01b031681565b600080546001600160a01b031633146102455761023e6001600e610542565b90506102ab565b600180546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9929181900390910190a160005b9150505b919050565b6002546001600160a01b031681565b6003546000906001600160a01b0316331415806102e557506003546001600160a01b0316155b156102fc576102f5600180610542565b90506103b7565b60028054600380546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927fd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a92908290030190a1600354604080516001600160a01b038085168252909216602083015280517fe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d8159281900390910190a160005b925050505b90565b6003546001600160a01b031681565b600080546001600160a01b031633146103e85761023e6001600f610542565b600380546001600160a01b038481166001600160a01b0319831617928390556040805192821680845293909116602083015280517fe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d8159281900390910190a160006102a7565b6001546000906001600160a01b031633141580610468575033155b15610479576102f560016000610542565b60008054600180546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc92908290030190a1600154604080516001600160a01b038085168252909216602083015280517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a99281900390910190a160006103b2565b6000546001600160a01b031681565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083601381111561057157fe5b83601681111561057d57fe5b604080519283526020830191909152600082820152519081900360600190a18260138111156105a857fe5b939250505056fea265627a7a7231582032652b6a6ec48d80ae67bc555f3a9a1af782093d9b3fc91eb2e08ac56798924964736f6c63430005110032"

// DeployComtroller deploys a new Ethereum contract, binding an instance of Comtroller to it.
func DeployComtroller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Comtroller, error) {
	parsed, err := abi.JSON(strings.NewReader(ComtrollerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ComtrollerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Comtroller{ComtrollerCaller: ComtrollerCaller{contract: contract}, ComtrollerTransactor: ComtrollerTransactor{contract: contract}, ComtrollerFilterer: ComtrollerFilterer{contract: contract}}, nil
}

// Comtroller is an auto generated Go binding around an Ethereum contract.
type Comtroller struct {
	ComtrollerCaller     // Read-only binding to the contract
	ComtrollerTransactor // Write-only binding to the contract
	ComtrollerFilterer   // Log filterer for contract events
}

// ComtrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComtrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComtrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComtrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComtrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComtrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComtrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComtrollerSession struct {
	Contract     *Comtroller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComtrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComtrollerCallerSession struct {
	Contract *ComtrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ComtrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComtrollerTransactorSession struct {
	Contract     *ComtrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ComtrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComtrollerRaw struct {
	Contract *Comtroller // Generic contract binding to access the raw methods on
}

// ComtrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComtrollerCallerRaw struct {
	Contract *ComtrollerCaller // Generic read-only contract binding to access the raw methods on
}

// ComtrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComtrollerTransactorRaw struct {
	Contract *ComtrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComtroller creates a new instance of Comtroller, bound to a specific deployed contract.
func NewComtroller(address common.Address, backend bind.ContractBackend) (*Comtroller, error) {
	contract, err := bindComtroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Comtroller{ComtrollerCaller: ComtrollerCaller{contract: contract}, ComtrollerTransactor: ComtrollerTransactor{contract: contract}, ComtrollerFilterer: ComtrollerFilterer{contract: contract}}, nil
}

// NewComtrollerCaller creates a new read-only instance of Comtroller, bound to a specific deployed contract.
func NewComtrollerCaller(address common.Address, caller bind.ContractCaller) (*ComtrollerCaller, error) {
	contract, err := bindComtroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComtrollerCaller{contract: contract}, nil
}

// NewComtrollerTransactor creates a new write-only instance of Comtroller, bound to a specific deployed contract.
func NewComtrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*ComtrollerTransactor, error) {
	contract, err := bindComtroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComtrollerTransactor{contract: contract}, nil
}

// NewComtrollerFilterer creates a new log filterer instance of Comtroller, bound to a specific deployed contract.
func NewComtrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*ComtrollerFilterer, error) {
	contract, err := bindComtroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComtrollerFilterer{contract: contract}, nil
}

// bindComtroller binds a generic wrapper to an already deployed contract.
func bindComtroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComtrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comtroller *ComtrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comtroller.Contract.ComtrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comtroller *ComtrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comtroller.Contract.ComtrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comtroller *ComtrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comtroller.Contract.ComtrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comtroller *ComtrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comtroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comtroller *ComtrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comtroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comtroller *ComtrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comtroller.Contract.contract.Transact(opts, method, params...)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "_borrowGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerSession) BorrowGuardianPaused() (bool, error) {
	return _Comtroller.Contract.BorrowGuardianPaused(&_Comtroller.CallOpts)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCallerSession) BorrowGuardianPaused() (bool, error) {
	return _Comtroller.Contract.BorrowGuardianPaused(&_Comtroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCaller) MintGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "_mintGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerSession) MintGuardianPaused() (bool, error) {
	return _Comtroller.Contract.MintGuardianPaused(&_Comtroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCallerSession) MintGuardianPaused() (bool, error) {
	return _Comtroller.Contract.MintGuardianPaused(&_Comtroller.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comtroller *ComtrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comtroller *ComtrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Comtroller.Contract.AccountAssets(&_Comtroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comtroller *ComtrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Comtroller.Contract.AccountAssets(&_Comtroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comtroller *ComtrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comtroller *ComtrollerSession) Admin() (common.Address, error) {
	return _Comtroller.Contract.Admin(&_Comtroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comtroller *ComtrollerCallerSession) Admin() (common.Address, error) {
	return _Comtroller.Contract.Admin(&_Comtroller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comtroller *ComtrollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comtroller *ComtrollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Comtroller.Contract.AllMarkets(&_Comtroller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comtroller *ComtrollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Comtroller.Contract.AllMarkets(&_Comtroller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comtroller *ComtrollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comtroller *ComtrollerSession) BorrowCapGuardian() (common.Address, error) {
	return _Comtroller.Contract.BorrowCapGuardian(&_Comtroller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comtroller *ComtrollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _Comtroller.Contract.BorrowCapGuardian(&_Comtroller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comtroller *ComtrollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comtroller *ComtrollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.BorrowCaps(&_Comtroller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.BorrowCaps(&_Comtroller.CallOpts, arg0)
}

// TestborrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comtroller *ComtrollerCaller) TestborrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestborrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comtroller *ComtrollerSession) TestborrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comtroller.Contract.TestborrowGuardianPaused(&_Comtroller.CallOpts, arg0)
}

// TestborrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comtroller *ComtrollerCallerSession) TestborrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comtroller.Contract.TestborrowGuardianPaused(&_Comtroller.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Comtroller *ComtrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, vToken common.Address) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "checkMembership", account, vToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Comtroller *ComtrollerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _Comtroller.Contract.CheckMembership(&_Comtroller.CallOpts, account, vToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address vToken) view returns(bool)
func (_Comtroller *ComtrollerCallerSession) CheckMembership(account common.Address, vToken common.Address) (bool, error) {
	return _Comtroller.Contract.CheckMembership(&_Comtroller.CallOpts, account, vToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comtroller *ComtrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comtroller *ComtrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Comtroller.Contract.CloseFactorMantissa(&_Comtroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Comtroller.Contract.CloseFactorMantissa(&_Comtroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comtroller *ComtrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comtroller *ComtrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Comtroller.Contract.ComptrollerImplementation(&_Comtroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comtroller *ComtrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Comtroller.Contract.ComptrollerImplementation(&_Comtroller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comtroller *ComtrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comtroller *ComtrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Comtroller.Contract.GetAccountLiquidity(&_Comtroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comtroller *ComtrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Comtroller.Contract.GetAccountLiquidity(&_Comtroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comtroller *ComtrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comtroller *ComtrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _Comtroller.Contract.GetAllMarkets(&_Comtroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comtroller *ComtrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _Comtroller.Contract.GetAllMarkets(&_Comtroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comtroller *ComtrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comtroller *ComtrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Comtroller.Contract.GetAssetsIn(&_Comtroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comtroller *ComtrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Comtroller.Contract.GetAssetsIn(&_Comtroller.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comtroller *ComtrollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comtroller *ComtrollerSession) GetBlockNumber() (*big.Int, error) {
	return _Comtroller.Contract.GetBlockNumber(&_Comtroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Comtroller.Contract.GetBlockNumber(&_Comtroller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comtroller *ComtrollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, vTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comtroller *ComtrollerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Comtroller.Contract.GetHypotheticalAccountLiquidity(&_Comtroller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address vTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comtroller *ComtrollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, vTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Comtroller.Contract.GetHypotheticalAccountLiquidity(&_Comtroller.CallOpts, account, vTokenModify, redeemTokens, borrowAmount)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Comtroller *ComtrollerCaller) GetXVSAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "getXVSAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Comtroller *ComtrollerSession) GetXVSAddress() (common.Address, error) {
	return _Comtroller.Contract.GetXVSAddress(&_Comtroller.CallOpts)
}

// GetXVSAddress is a free data retrieval call binding the contract method 0xbf32442d.
//
// Solidity: function getXVSAddress() view returns(address)
func (_Comtroller *ComtrollerCallerSession) GetXVSAddress() (common.Address, error) {
	return _Comtroller.Contract.GetXVSAddress(&_Comtroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comtroller *ComtrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comtroller *ComtrollerSession) IsComptroller() (bool, error) {
	return _Comtroller.Contract.IsComptroller(&_Comtroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comtroller *ComtrollerCallerSession) IsComptroller() (bool, error) {
	return _Comtroller.Contract.IsComptroller(&_Comtroller.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comtroller *ComtrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", vTokenBorrowed, vTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comtroller *ComtrollerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comtroller.Contract.LiquidateCalculateSeizeTokens(&_Comtroller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address vTokenBorrowed, address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comtroller *ComtrollerCallerSession) LiquidateCalculateSeizeTokens(vTokenBorrowed common.Address, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comtroller.Contract.LiquidateCalculateSeizeTokens(&_Comtroller.CallOpts, vTokenBorrowed, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comtroller *ComtrollerCaller) LiquidateVAICalculateSeizeTokens(opts *bind.CallOpts, vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "liquidateVAICalculateSeizeTokens", vTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comtroller *ComtrollerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comtroller.Contract.LiquidateVAICalculateSeizeTokens(&_Comtroller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidateVAICalculateSeizeTokens is a free data retrieval call binding the contract method 0xa78dc775.
//
// Solidity: function liquidateVAICalculateSeizeTokens(address vTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comtroller *ComtrollerCallerSession) LiquidateVAICalculateSeizeTokens(vTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comtroller.Contract.LiquidateVAICalculateSeizeTokens(&_Comtroller.CallOpts, vTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comtroller *ComtrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comtroller *ComtrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Comtroller.Contract.LiquidationIncentiveMantissa(&_Comtroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Comtroller.Contract.LiquidationIncentiveMantissa(&_Comtroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Comtroller *ComtrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsVenus                  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsListed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsVenus = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Comtroller *ComtrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _Comtroller.Contract.Markets(&_Comtroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isVenus)
func (_Comtroller *ComtrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsVenus                  bool
}, error) {
	return _Comtroller.Contract.Markets(&_Comtroller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Comtroller *ComtrollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Comtroller *ComtrollerSession) MaxAssets() (*big.Int, error) {
	return _Comtroller.Contract.MaxAssets(&_Comtroller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) MaxAssets() (*big.Int, error) {
	return _Comtroller.Contract.MaxAssets(&_Comtroller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Comtroller *ComtrollerCaller) MinReleaseAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "minReleaseAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Comtroller *ComtrollerSession) MinReleaseAmount() (*big.Int, error) {
	return _Comtroller.Contract.MinReleaseAmount(&_Comtroller.CallOpts)
}

// MinReleaseAmount is a free data retrieval call binding the contract method 0x0db4b4e5.
//
// Solidity: function minReleaseAmount() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) MinReleaseAmount() (*big.Int, error) {
	return _Comtroller.Contract.MinReleaseAmount(&_Comtroller.CallOpts)
}

// TestmintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comtroller *ComtrollerCaller) TestmintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestmintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comtroller *ComtrollerSession) TestmintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comtroller.Contract.TestmintGuardianPaused(&_Comtroller.CallOpts, arg0)
}

// TestmintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comtroller *ComtrollerCallerSession) TestmintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comtroller.Contract.TestmintGuardianPaused(&_Comtroller.CallOpts, arg0)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCaller) MintVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "mintVAIGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerSession) MintVAIGuardianPaused() (bool, error) {
	return _Comtroller.Contract.MintVAIGuardianPaused(&_Comtroller.CallOpts)
}

// MintVAIGuardianPaused is a free data retrieval call binding the contract method 0x4088c73e.
//
// Solidity: function mintVAIGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCallerSession) MintVAIGuardianPaused() (bool, error) {
	return _Comtroller.Contract.MintVAIGuardianPaused(&_Comtroller.CallOpts)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Comtroller *ComtrollerCaller) MintedVAIs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "mintedVAIs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Comtroller *ComtrollerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.MintedVAIs(&_Comtroller.CallOpts, arg0)
}

// MintedVAIs is a free data retrieval call binding the contract method 0x2bc7e29e.
//
// Solidity: function mintedVAIs(address ) view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) MintedVAIs(arg0 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.MintedVAIs(&_Comtroller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comtroller *ComtrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comtroller *ComtrollerSession) Oracle() (common.Address, error) {
	return _Comtroller.Contract.Oracle(&_Comtroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comtroller *ComtrollerCallerSession) Oracle() (common.Address, error) {
	return _Comtroller.Contract.Oracle(&_Comtroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comtroller *ComtrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comtroller *ComtrollerSession) PauseGuardian() (common.Address, error) {
	return _Comtroller.Contract.PauseGuardian(&_Comtroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comtroller *ComtrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _Comtroller.Contract.PauseGuardian(&_Comtroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comtroller *ComtrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comtroller *ComtrollerSession) PendingAdmin() (common.Address, error) {
	return _Comtroller.Contract.PendingAdmin(&_Comtroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comtroller *ComtrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Comtroller.Contract.PendingAdmin(&_Comtroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comtroller *ComtrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comtroller *ComtrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Comtroller.Contract.PendingComptrollerImplementation(&_Comtroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comtroller *ComtrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Comtroller.Contract.PendingComptrollerImplementation(&_Comtroller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Comtroller *ComtrollerCaller) ProtocolPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "protocolPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Comtroller *ComtrollerSession) ProtocolPaused() (bool, error) {
	return _Comtroller.Contract.ProtocolPaused(&_Comtroller.CallOpts)
}

// ProtocolPaused is a free data retrieval call binding the contract method 0x425fad58.
//
// Solidity: function protocolPaused() view returns(bool)
func (_Comtroller *ComtrollerCallerSession) ProtocolPaused() (bool, error) {
	return _Comtroller.Contract.ProtocolPaused(&_Comtroller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Comtroller *ComtrollerCaller) ReleaseStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "releaseStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Comtroller *ComtrollerSession) ReleaseStartBlock() (*big.Int, error) {
	return _Comtroller.Contract.ReleaseStartBlock(&_Comtroller.CallOpts)
}

// ReleaseStartBlock is a free data retrieval call binding the contract method 0x719f701b.
//
// Solidity: function releaseStartBlock() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) ReleaseStartBlock() (*big.Int, error) {
	return _Comtroller.Contract.ReleaseStartBlock(&_Comtroller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCaller) RepayVAIGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "repayVAIGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerSession) RepayVAIGuardianPaused() (bool, error) {
	return _Comtroller.Contract.RepayVAIGuardianPaused(&_Comtroller.CallOpts)
}

// RepayVAIGuardianPaused is a free data retrieval call binding the contract method 0x76551383.
//
// Solidity: function repayVAIGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCallerSession) RepayVAIGuardianPaused() (bool, error) {
	return _Comtroller.Contract.RepayVAIGuardianPaused(&_Comtroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerSession) SeizeGuardianPaused() (bool, error) {
	return _Comtroller.Contract.SeizeGuardianPaused(&_Comtroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _Comtroller.Contract.SeizeGuardianPaused(&_Comtroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerSession) TransferGuardianPaused() (bool, error) {
	return _Comtroller.Contract.TransferGuardianPaused(&_Comtroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comtroller *ComtrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _Comtroller.Contract.TransferGuardianPaused(&_Comtroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Comtroller *ComtrollerCaller) TreasuryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "treasuryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Comtroller *ComtrollerSession) TreasuryAddress() (common.Address, error) {
	return _Comtroller.Contract.TreasuryAddress(&_Comtroller.CallOpts)
}

// TreasuryAddress is a free data retrieval call binding the contract method 0xc5f956af.
//
// Solidity: function treasuryAddress() view returns(address)
func (_Comtroller *ComtrollerCallerSession) TreasuryAddress() (common.Address, error) {
	return _Comtroller.Contract.TreasuryAddress(&_Comtroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Comtroller *ComtrollerCaller) TreasuryGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "treasuryGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Comtroller *ComtrollerSession) TreasuryGuardian() (common.Address, error) {
	return _Comtroller.Contract.TreasuryGuardian(&_Comtroller.CallOpts)
}

// TreasuryGuardian is a free data retrieval call binding the contract method 0xb2eafc39.
//
// Solidity: function treasuryGuardian() view returns(address)
func (_Comtroller *ComtrollerCallerSession) TreasuryGuardian() (common.Address, error) {
	return _Comtroller.Contract.TreasuryGuardian(&_Comtroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Comtroller *ComtrollerCaller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "treasuryPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Comtroller *ComtrollerSession) TreasuryPercent() (*big.Int, error) {
	return _Comtroller.Contract.TreasuryPercent(&_Comtroller.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) TreasuryPercent() (*big.Int, error) {
	return _Comtroller.Contract.TreasuryPercent(&_Comtroller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Comtroller *ComtrollerCaller) VaiController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "vaiController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Comtroller *ComtrollerSession) VaiController() (common.Address, error) {
	return _Comtroller.Contract.VaiController(&_Comtroller.CallOpts)
}

// VaiController is a free data retrieval call binding the contract method 0x9254f5e5.
//
// Solidity: function vaiController() view returns(address)
func (_Comtroller *ComtrollerCallerSession) VaiController() (common.Address, error) {
	return _Comtroller.Contract.VaiController(&_Comtroller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Comtroller *ComtrollerCaller) VaiMintRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "vaiMintRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Comtroller *ComtrollerSession) VaiMintRate() (*big.Int, error) {
	return _Comtroller.Contract.VaiMintRate(&_Comtroller.CallOpts)
}

// VaiMintRate is a free data retrieval call binding the contract method 0xbec04f72.
//
// Solidity: function vaiMintRate() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) VaiMintRate() (*big.Int, error) {
	return _Comtroller.Contract.VaiMintRate(&_Comtroller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Comtroller *ComtrollerCaller) VaiVaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "vaiVaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Comtroller *ComtrollerSession) VaiVaultAddress() (common.Address, error) {
	return _Comtroller.Contract.VaiVaultAddress(&_Comtroller.CallOpts)
}

// VaiVaultAddress is a free data retrieval call binding the contract method 0x7d172bd5.
//
// Solidity: function vaiVaultAddress() view returns(address)
func (_Comtroller *ComtrollerCallerSession) VaiVaultAddress() (common.Address, error) {
	return _Comtroller.Contract.VaiVaultAddress(&_Comtroller.CallOpts)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Comtroller *ComtrollerCaller) VenusAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Comtroller *ComtrollerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.VenusAccrued(&_Comtroller.CallOpts, arg0)
}

// VenusAccrued is a free data retrieval call binding the contract method 0x8a7dc165.
//
// Solidity: function venusAccrued(address ) view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) VenusAccrued(arg0 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.VenusAccrued(&_Comtroller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comtroller *ComtrollerCaller) VenusBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comtroller *ComtrollerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comtroller.Contract.VenusBorrowState(&_Comtroller.CallOpts, arg0)
}

// VenusBorrowState is a free data retrieval call binding the contract method 0xe37d4b79.
//
// Solidity: function venusBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comtroller *ComtrollerCallerSession) VenusBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comtroller.Contract.VenusBorrowState(&_Comtroller.CallOpts, arg0)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Comtroller *ComtrollerCaller) VenusBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Comtroller *ComtrollerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.VenusBorrowerIndex(&_Comtroller.CallOpts, arg0, arg1)
}

// VenusBorrowerIndex is a free data retrieval call binding the contract method 0x08e0225c.
//
// Solidity: function venusBorrowerIndex(address , address ) view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) VenusBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.VenusBorrowerIndex(&_Comtroller.CallOpts, arg0, arg1)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Comtroller *ComtrollerCaller) VenusInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Comtroller *ComtrollerSession) VenusInitialIndex() (*big.Int, error) {
	return _Comtroller.Contract.VenusInitialIndex(&_Comtroller.CallOpts)
}

// VenusInitialIndex is a free data retrieval call binding the contract method 0xc5b4db55.
//
// Solidity: function venusInitialIndex() view returns(uint224)
func (_Comtroller *ComtrollerCallerSession) VenusInitialIndex() (*big.Int, error) {
	return _Comtroller.Contract.VenusInitialIndex(&_Comtroller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Comtroller *ComtrollerCaller) VenusRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Comtroller *ComtrollerSession) VenusRate() (*big.Int, error) {
	return _Comtroller.Contract.VenusRate(&_Comtroller.CallOpts)
}

// VenusRate is a free data retrieval call binding the contract method 0x879c2e1d.
//
// Solidity: function venusRate() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) VenusRate() (*big.Int, error) {
	return _Comtroller.Contract.VenusRate(&_Comtroller.CallOpts)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Comtroller *ComtrollerCaller) VenusSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Comtroller *ComtrollerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.VenusSpeeds(&_Comtroller.CallOpts, arg0)
}

// VenusSpeeds is a free data retrieval call binding the contract method 0x1abcaa77.
//
// Solidity: function venusSpeeds(address ) view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) VenusSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.VenusSpeeds(&_Comtroller.CallOpts, arg0)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Comtroller *ComtrollerCaller) VenusSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Comtroller *ComtrollerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.VenusSupplierIndex(&_Comtroller.CallOpts, arg0, arg1)
}

// VenusSupplierIndex is a free data retrieval call binding the contract method 0x41a18d2c.
//
// Solidity: function venusSupplierIndex(address , address ) view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) VenusSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comtroller.Contract.VenusSupplierIndex(&_Comtroller.CallOpts, arg0, arg1)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comtroller *ComtrollerCaller) VenusSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comtroller *ComtrollerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comtroller.Contract.VenusSupplyState(&_Comtroller.CallOpts, arg0)
}

// VenusSupplyState is a free data retrieval call binding the contract method 0xb8324c7c.
//
// Solidity: function venusSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comtroller *ComtrollerCallerSession) VenusSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comtroller.Contract.VenusSupplyState(&_Comtroller.CallOpts, arg0)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Comtroller *ComtrollerCaller) VenusVAIRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusVAIRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Comtroller *ComtrollerSession) VenusVAIRate() (*big.Int, error) {
	return _Comtroller.Contract.VenusVAIRate(&_Comtroller.CallOpts)
}

// VenusVAIRate is a free data retrieval call binding the contract method 0x399cc80c.
//
// Solidity: function venusVAIRate() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) VenusVAIRate() (*big.Int, error) {
	return _Comtroller.Contract.VenusVAIRate(&_Comtroller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Comtroller *ComtrollerCaller) VenusVAIVaultRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comtroller.contract.Call(opts, &out, "venusVAIVaultRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Comtroller *ComtrollerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _Comtroller.Contract.VenusVAIVaultRate(&_Comtroller.CallOpts)
}

// VenusVAIVaultRate is a free data retrieval call binding the contract method 0xfa6331d8.
//
// Solidity: function venusVAIVaultRate() view returns(uint256)
func (_Comtroller *ComtrollerCallerSession) VenusVAIVaultRate() (*big.Int, error) {
	return _Comtroller.Contract.VenusVAIVaultRate(&_Comtroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comtroller *ComtrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comtroller *ComtrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.Become(&_Comtroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comtroller *ComtrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.Become(&_Comtroller.TransactOpts, unitroller)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comtroller *ComtrollerTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comtroller *ComtrollerSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SetBorrowCapGuardian(&_Comtroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comtroller *ComtrollerTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SetBorrowCapGuardian(&_Comtroller.TransactOpts, newBorrowCapGuardian)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comtroller *ComtrollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetCloseFactor(&_Comtroller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetCloseFactor(&_Comtroller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setCollateralFactor", vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comtroller *ComtrollerSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetCollateralFactor(&_Comtroller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address vToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetCollateralFactor(vToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetCollateralFactor(&_Comtroller.TransactOpts, vToken, newCollateralFactorMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comtroller *ComtrollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetLiquidationIncentive(&_Comtroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetLiquidationIncentive(&_Comtroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Comtroller *ComtrollerTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setMarketBorrowCaps", vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Comtroller *ComtrollerSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetMarketBorrowCaps(&_Comtroller.TransactOpts, vTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] vTokens, uint256[] newBorrowCaps) returns()
func (_Comtroller *ComtrollerTransactorSession) SetMarketBorrowCaps(vTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetMarketBorrowCaps(&_Comtroller.TransactOpts, vTokens, newBorrowCaps)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comtroller *ComtrollerSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SetPauseGuardian(&_Comtroller.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SetPauseGuardian(&_Comtroller.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comtroller *ComtrollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SetPriceOracle(&_Comtroller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SetPriceOracle(&_Comtroller.TransactOpts, newOracle)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Comtroller *ComtrollerTransactor) SetProtocolPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setProtocolPaused", state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Comtroller *ComtrollerSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _Comtroller.Contract.SetProtocolPaused(&_Comtroller.TransactOpts, state)
}

// SetProtocolPaused is a paid mutator transaction binding the contract method 0x2a6a6065.
//
// Solidity: function _setProtocolPaused(bool state) returns(bool)
func (_Comtroller *ComtrollerTransactorSession) SetProtocolPaused(state bool) (*types.Transaction, error) {
	return _Comtroller.Contract.SetProtocolPaused(&_Comtroller.TransactOpts, state)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetTreasuryData(opts *bind.TransactOpts, newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setTreasuryData", newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Comtroller *ComtrollerSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetTreasuryData(&_Comtroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetTreasuryData is a paid mutator transaction binding the contract method 0xd24febad.
//
// Solidity: function _setTreasuryData(address newTreasuryGuardian, address newTreasuryAddress, uint256 newTreasuryPercent) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetTreasuryData(newTreasuryGuardian common.Address, newTreasuryAddress common.Address, newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetTreasuryData(&_Comtroller.TransactOpts, newTreasuryGuardian, newTreasuryAddress, newTreasuryPercent)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetVAIController(opts *bind.TransactOpts, vaiController_ common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setVAIController", vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Comtroller *ComtrollerSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVAIController(&_Comtroller.TransactOpts, vaiController_)
}

// SetVAIController is a paid mutator transaction binding the contract method 0x9cfdd9e6.
//
// Solidity: function _setVAIController(address vaiController_) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetVAIController(vaiController_ common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVAIController(&_Comtroller.TransactOpts, vaiController_)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetVAIMintRate(opts *bind.TransactOpts, newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setVAIMintRate", newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Comtroller *ComtrollerSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVAIMintRate(&_Comtroller.TransactOpts, newVAIMintRate)
}

// SetVAIMintRate is a paid mutator transaction binding the contract method 0x2ec04124.
//
// Solidity: function _setVAIMintRate(uint256 newVAIMintRate) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetVAIMintRate(newVAIMintRate *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVAIMintRate(&_Comtroller.TransactOpts, newVAIMintRate)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Comtroller *ComtrollerTransactor) SetVAIVaultInfo(opts *bind.TransactOpts, vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setVAIVaultInfo", vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Comtroller *ComtrollerSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVAIVaultInfo(&_Comtroller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVAIVaultInfo is a paid mutator transaction binding the contract method 0x4e0853db.
//
// Solidity: function _setVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 minReleaseAmount_) returns()
func (_Comtroller *ComtrollerTransactorSession) SetVAIVaultInfo(vault_ common.Address, releaseStartBlock_ *big.Int, minReleaseAmount_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVAIVaultInfo(&_Comtroller.TransactOpts, vault_, releaseStartBlock_, minReleaseAmount_)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Comtroller *ComtrollerTransactor) SetVenusSpeed(opts *bind.TransactOpts, vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setVenusSpeed", vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Comtroller *ComtrollerSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVenusSpeed(&_Comtroller.TransactOpts, vToken, venusSpeed)
}

// SetVenusSpeed is a paid mutator transaction binding the contract method 0xa06a87f1.
//
// Solidity: function _setVenusSpeed(address vToken, uint256 venusSpeed) returns()
func (_Comtroller *ComtrollerTransactorSession) SetVenusSpeed(vToken common.Address, venusSpeed *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVenusSpeed(&_Comtroller.TransactOpts, vToken, venusSpeed)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_Comtroller *ComtrollerTransactor) SetVenusVAIRate(opts *bind.TransactOpts, venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setVenusVAIRate", venusVAIRate_)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_Comtroller *ComtrollerSession) SetVenusVAIRate(venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVenusVAIRate(&_Comtroller.TransactOpts, venusVAIRate_)
}

// SetVenusVAIRate is a paid mutator transaction binding the contract method 0xe6de6528.
//
// Solidity: function _setVenusVAIRate(uint256 venusVAIRate_) returns()
func (_Comtroller *ComtrollerTransactorSession) SetVenusVAIRate(venusVAIRate_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVenusVAIRate(&_Comtroller.TransactOpts, venusVAIRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Comtroller *ComtrollerTransactor) SetVenusVAIVaultRate(opts *bind.TransactOpts, venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_setVenusVAIVaultRate", venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Comtroller *ComtrollerSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVenusVAIVaultRate(&_Comtroller.TransactOpts, venusVAIVaultRate_)
}

// SetVenusVAIVaultRate is a paid mutator transaction binding the contract method 0x6662c7c9.
//
// Solidity: function _setVenusVAIVaultRate(uint256 venusVAIVaultRate_) returns()
func (_Comtroller *ComtrollerTransactorSession) SetVenusVAIVaultRate(venusVAIVaultRate_ *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetVenusVAIVaultRate(&_Comtroller.TransactOpts, venusVAIVaultRate_)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SupportMarket(opts *bind.TransactOpts, vToken common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "_supportMarket", vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Comtroller *ComtrollerSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SupportMarket(&_Comtroller.TransactOpts, vToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address vToken) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SupportMarket(vToken common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.SupportMarket(&_Comtroller.TransactOpts, vToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comtroller *ComtrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "borrowAllowed", vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comtroller *ComtrollerSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.BorrowAllowed(&_Comtroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address vToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) BorrowAllowed(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.BorrowAllowed(&_Comtroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Comtroller *ComtrollerTransactor) BorrowVerify(opts *bind.TransactOpts, vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "borrowVerify", vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Comtroller *ComtrollerSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.BorrowVerify(&_Comtroller.TransactOpts, vToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address vToken, address borrower, uint256 borrowAmount) returns()
func (_Comtroller *ComtrollerTransactorSession) BorrowVerify(vToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.BorrowVerify(&_Comtroller.TransactOpts, vToken, borrower, borrowAmount)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Comtroller *ComtrollerTransactor) ClaimVenus(opts *bind.TransactOpts, holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "claimVenus", holder, vTokens)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Comtroller *ComtrollerSession) ClaimVenus(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.ClaimVenus(&_Comtroller.TransactOpts, holder, vTokens)
}

// ClaimVenus is a paid mutator transaction binding the contract method 0x86df31ee.
//
// Solidity: function claimVenus(address holder, address[] vTokens) returns()
func (_Comtroller *ComtrollerTransactorSession) ClaimVenus(holder common.Address, vTokens []common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.ClaimVenus(&_Comtroller.TransactOpts, holder, vTokens)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Comtroller *ComtrollerTransactor) ClaimVenus0(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "claimVenus0", holder)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Comtroller *ComtrollerSession) ClaimVenus0(holder common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.ClaimVenus0(&_Comtroller.TransactOpts, holder)
}

// ClaimVenus0 is a paid mutator transaction binding the contract method 0xadcd5fb9.
//
// Solidity: function claimVenus(address holder) returns()
func (_Comtroller *ComtrollerTransactorSession) ClaimVenus0(holder common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.ClaimVenus0(&_Comtroller.TransactOpts, holder)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Comtroller *ComtrollerTransactor) ClaimVenus1(opts *bind.TransactOpts, holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "claimVenus1", holders, vTokens, borrowers, suppliers)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Comtroller *ComtrollerSession) ClaimVenus1(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comtroller.Contract.ClaimVenus1(&_Comtroller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// ClaimVenus1 is a paid mutator transaction binding the contract method 0xd09c54ba.
//
// Solidity: function claimVenus(address[] holders, address[] vTokens, bool borrowers, bool suppliers) returns()
func (_Comtroller *ComtrollerTransactorSession) ClaimVenus1(holders []common.Address, vTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comtroller.Contract.ClaimVenus1(&_Comtroller.TransactOpts, holders, vTokens, borrowers, suppliers)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Comtroller *ComtrollerTransactor) DistributeVAIMinterVenus(opts *bind.TransactOpts, vaiMinter common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "distributeVAIMinterVenus", vaiMinter)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Comtroller *ComtrollerSession) DistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.DistributeVAIMinterVenus(&_Comtroller.TransactOpts, vaiMinter)
}

// DistributeVAIMinterVenus is a paid mutator transaction binding the contract method 0xf4b8d5fe.
//
// Solidity: function distributeVAIMinterVenus(address vaiMinter) returns()
func (_Comtroller *ComtrollerTransactorSession) DistributeVAIMinterVenus(vaiMinter common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.DistributeVAIMinterVenus(&_Comtroller.TransactOpts, vaiMinter)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Comtroller *ComtrollerTransactor) EnterMarkets(opts *bind.TransactOpts, vTokens []common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "enterMarkets", vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Comtroller *ComtrollerSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.EnterMarkets(&_Comtroller.TransactOpts, vTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] vTokens) returns(uint256[])
func (_Comtroller *ComtrollerTransactorSession) EnterMarkets(vTokens []common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.EnterMarkets(&_Comtroller.TransactOpts, vTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Comtroller *ComtrollerTransactor) ExitMarket(opts *bind.TransactOpts, vTokenAddress common.Address) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "exitMarket", vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Comtroller *ComtrollerSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.ExitMarket(&_Comtroller.TransactOpts, vTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address vTokenAddress) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) ExitMarket(vTokenAddress common.Address) (*types.Transaction, error) {
	return _Comtroller.Contract.ExitMarket(&_Comtroller.TransactOpts, vTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comtroller *ComtrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "liquidateBorrowAllowed", vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comtroller *ComtrollerSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.LiquidateBorrowAllowed(&_Comtroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) LiquidateBorrowAllowed(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.LiquidateBorrowAllowed(&_Comtroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comtroller *ComtrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "liquidateBorrowVerify", vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comtroller *ComtrollerSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.LiquidateBorrowVerify(&_Comtroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address vTokenBorrowed, address vTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comtroller *ComtrollerTransactorSession) LiquidateBorrowVerify(vTokenBorrowed common.Address, vTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.LiquidateBorrowVerify(&_Comtroller.TransactOpts, vTokenBorrowed, vTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comtroller *ComtrollerTransactor) MintAllowed(opts *bind.TransactOpts, vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "mintAllowed", vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comtroller *ComtrollerSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.MintAllowed(&_Comtroller.TransactOpts, vToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address vToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) MintAllowed(vToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.MintAllowed(&_Comtroller.TransactOpts, vToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comtroller *ComtrollerTransactor) MintVerify(opts *bind.TransactOpts, vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "mintVerify", vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comtroller *ComtrollerSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.MintVerify(&_Comtroller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address vToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comtroller *ComtrollerTransactorSession) MintVerify(vToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.MintVerify(&_Comtroller.TransactOpts, vToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comtroller *ComtrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "redeemAllowed", vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comtroller *ComtrollerSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.RedeemAllowed(&_Comtroller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address vToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) RedeemAllowed(vToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.RedeemAllowed(&_Comtroller.TransactOpts, vToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comtroller *ComtrollerTransactor) RedeemVerify(opts *bind.TransactOpts, vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "redeemVerify", vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comtroller *ComtrollerSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.RedeemVerify(&_Comtroller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address vToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comtroller *ComtrollerTransactorSession) RedeemVerify(vToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.RedeemVerify(&_Comtroller.TransactOpts, vToken, redeemer, redeemAmount, redeemTokens)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Comtroller *ComtrollerTransactor) ReleaseToVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "releaseToVault")
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Comtroller *ComtrollerSession) ReleaseToVault() (*types.Transaction, error) {
	return _Comtroller.Contract.ReleaseToVault(&_Comtroller.TransactOpts)
}

// ReleaseToVault is a paid mutator transaction binding the contract method 0xddfd287e.
//
// Solidity: function releaseToVault() returns()
func (_Comtroller *ComtrollerTransactorSession) ReleaseToVault() (*types.Transaction, error) {
	return _Comtroller.Contract.ReleaseToVault(&_Comtroller.TransactOpts)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comtroller *ComtrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "repayBorrowAllowed", vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comtroller *ComtrollerSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.RepayBorrowAllowed(&_Comtroller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address vToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) RepayBorrowAllowed(vToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.RepayBorrowAllowed(&_Comtroller.TransactOpts, vToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comtroller *ComtrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "repayBorrowVerify", vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comtroller *ComtrollerSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.RepayBorrowVerify(&_Comtroller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address vToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comtroller *ComtrollerTransactorSession) RepayBorrowVerify(vToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.RepayBorrowVerify(&_Comtroller.TransactOpts, vToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "seizeAllowed", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comtroller *ComtrollerSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SeizeAllowed(&_Comtroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SeizeAllowed(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SeizeAllowed(&_Comtroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comtroller *ComtrollerTransactor) SeizeVerify(opts *bind.TransactOpts, vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "seizeVerify", vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comtroller *ComtrollerSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SeizeVerify(&_Comtroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address vTokenCollateral, address vTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comtroller *ComtrollerTransactorSession) SeizeVerify(vTokenCollateral common.Address, vTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SeizeVerify(&_Comtroller.TransactOpts, vTokenCollateral, vTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Comtroller *ComtrollerTransactor) SetMintedVAIOf(opts *bind.TransactOpts, owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "setMintedVAIOf", owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Comtroller *ComtrollerSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetMintedVAIOf(&_Comtroller.TransactOpts, owner, amount)
}

// SetMintedVAIOf is a paid mutator transaction binding the contract method 0xfd51a3ad.
//
// Solidity: function setMintedVAIOf(address owner, uint256 amount) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) SetMintedVAIOf(owner common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.SetMintedVAIOf(&_Comtroller.TransactOpts, owner, amount)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comtroller *ComtrollerTransactor) TransferAllowed(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "transferAllowed", vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comtroller *ComtrollerSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.TransferAllowed(&_Comtroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address vToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comtroller *ComtrollerTransactorSession) TransferAllowed(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.TransferAllowed(&_Comtroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Comtroller *ComtrollerTransactor) TransferVerify(opts *bind.TransactOpts, vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.contract.Transact(opts, "transferVerify", vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Comtroller *ComtrollerSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.TransferVerify(&_Comtroller.TransactOpts, vToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address vToken, address src, address dst, uint256 transferTokens) returns()
func (_Comtroller *ComtrollerTransactorSession) TransferVerify(vToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comtroller.Contract.TransferVerify(&_Comtroller.TransactOpts, vToken, src, dst, transferTokens)
}

// ComtrollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the Comtroller contract.
type ComtrollerActionPausedIterator struct {
	Event *ComtrollerActionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerActionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerActionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerActionPaused represents a ActionPaused event raised by the Comtroller contract.
type ComtrollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comtroller *ComtrollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*ComtrollerActionPausedIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &ComtrollerActionPausedIterator{contract: _Comtroller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comtroller *ComtrollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *ComtrollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerActionPaused)
				if err := _Comtroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comtroller *ComtrollerFilterer) ParseActionPaused(log types.Log) (*ComtrollerActionPaused, error) {
	event := new(ComtrollerActionPaused)
	if err := _Comtroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerActionPaused0Iterator is returned from FilterActionPaused0 and is used to iterate over the raw logs and unpacked data for ActionPaused0 events raised by the Comtroller contract.
type ComtrollerActionPaused0Iterator struct {
	Event *ComtrollerActionPaused0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerActionPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerActionPaused0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerActionPaused0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerActionPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerActionPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerActionPaused0 represents a ActionPaused0 event raised by the Comtroller contract.
type ComtrollerActionPaused0 struct {
	VToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused0 is a free log retrieval operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Comtroller *ComtrollerFilterer) FilterActionPaused0(opts *bind.FilterOpts) (*ComtrollerActionPaused0Iterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return &ComtrollerActionPaused0Iterator{contract: _Comtroller.contract, event: "ActionPaused0", logs: logs, sub: sub}, nil
}

// WatchActionPaused0 is a free log subscription operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Comtroller *ComtrollerFilterer) WatchActionPaused0(opts *bind.WatchOpts, sink chan<- *ComtrollerActionPaused0) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerActionPaused0)
				if err := _Comtroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused0 is a log parse operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address vToken, string action, bool pauseState)
func (_Comtroller *ComtrollerFilterer) ParseActionPaused0(log types.Log) (*ComtrollerActionPaused0, error) {
	event := new(ComtrollerActionPaused0)
	if err := _Comtroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerActionProtocolPausedIterator is returned from FilterActionProtocolPaused and is used to iterate over the raw logs and unpacked data for ActionProtocolPaused events raised by the Comtroller contract.
type ComtrollerActionProtocolPausedIterator struct {
	Event *ComtrollerActionProtocolPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerActionProtocolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerActionProtocolPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerActionProtocolPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerActionProtocolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerActionProtocolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerActionProtocolPaused represents a ActionProtocolPaused event raised by the Comtroller contract.
type ComtrollerActionProtocolPaused struct {
	State bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterActionProtocolPaused is a free log retrieval operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Comtroller *ComtrollerFilterer) FilterActionProtocolPaused(opts *bind.FilterOpts) (*ComtrollerActionProtocolPausedIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return &ComtrollerActionProtocolPausedIterator{contract: _Comtroller.contract, event: "ActionProtocolPaused", logs: logs, sub: sub}, nil
}

// WatchActionProtocolPaused is a free log subscription operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Comtroller *ComtrollerFilterer) WatchActionProtocolPaused(opts *bind.WatchOpts, sink chan<- *ComtrollerActionProtocolPaused) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "ActionProtocolPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerActionProtocolPaused)
				if err := _Comtroller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionProtocolPaused is a log parse operation binding the contract event 0xd7500633dd3ddd74daa7af62f8c8404c7fe4a81da179998db851696bed004b38.
//
// Solidity: event ActionProtocolPaused(bool state)
func (_Comtroller *ComtrollerFilterer) ParseActionProtocolPaused(log types.Log) (*ComtrollerActionProtocolPaused, error) {
	event := new(ComtrollerActionProtocolPaused)
	if err := _Comtroller.contract.UnpackLog(event, "ActionProtocolPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerDistributedBorrowerVenusIterator is returned from FilterDistributedBorrowerVenus and is used to iterate over the raw logs and unpacked data for DistributedBorrowerVenus events raised by the Comtroller contract.
type ComtrollerDistributedBorrowerVenusIterator struct {
	Event *ComtrollerDistributedBorrowerVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerDistributedBorrowerVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerDistributedBorrowerVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerDistributedBorrowerVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerDistributedBorrowerVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerDistributedBorrowerVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerDistributedBorrowerVenus represents a DistributedBorrowerVenus event raised by the Comtroller contract.
type ComtrollerDistributedBorrowerVenus struct {
	VToken           common.Address
	Borrower         common.Address
	VenusDelta       *big.Int
	VenusBorrowIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerVenus is a free log retrieval operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Comtroller *ComtrollerFilterer) FilterDistributedBorrowerVenus(opts *bind.FilterOpts, vToken []common.Address, borrower []common.Address) (*ComtrollerDistributedBorrowerVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &ComtrollerDistributedBorrowerVenusIterator{contract: _Comtroller.contract, event: "DistributedBorrowerVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerVenus is a free log subscription operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Comtroller *ComtrollerFilterer) WatchDistributedBorrowerVenus(opts *bind.WatchOpts, sink chan<- *ComtrollerDistributedBorrowerVenus, vToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "DistributedBorrowerVenus", vTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerDistributedBorrowerVenus)
				if err := _Comtroller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedBorrowerVenus is a log parse operation binding the contract event 0x837bdc11fca9f17ce44167944475225a205279b17e88c791c3b1f66f354668fb.
//
// Solidity: event DistributedBorrowerVenus(address indexed vToken, address indexed borrower, uint256 venusDelta, uint256 venusBorrowIndex)
func (_Comtroller *ComtrollerFilterer) ParseDistributedBorrowerVenus(log types.Log) (*ComtrollerDistributedBorrowerVenus, error) {
	event := new(ComtrollerDistributedBorrowerVenus)
	if err := _Comtroller.contract.UnpackLog(event, "DistributedBorrowerVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerDistributedSupplierVenusIterator is returned from FilterDistributedSupplierVenus and is used to iterate over the raw logs and unpacked data for DistributedSupplierVenus events raised by the Comtroller contract.
type ComtrollerDistributedSupplierVenusIterator struct {
	Event *ComtrollerDistributedSupplierVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerDistributedSupplierVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerDistributedSupplierVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerDistributedSupplierVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerDistributedSupplierVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerDistributedSupplierVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerDistributedSupplierVenus represents a DistributedSupplierVenus event raised by the Comtroller contract.
type ComtrollerDistributedSupplierVenus struct {
	VToken           common.Address
	Supplier         common.Address
	VenusDelta       *big.Int
	VenusSupplyIndex *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierVenus is a free log retrieval operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Comtroller *ComtrollerFilterer) FilterDistributedSupplierVenus(opts *bind.FilterOpts, vToken []common.Address, supplier []common.Address) (*ComtrollerDistributedSupplierVenusIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &ComtrollerDistributedSupplierVenusIterator{contract: _Comtroller.contract, event: "DistributedSupplierVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierVenus is a free log subscription operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Comtroller *ComtrollerFilterer) WatchDistributedSupplierVenus(opts *bind.WatchOpts, sink chan<- *ComtrollerDistributedSupplierVenus, vToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "DistributedSupplierVenus", vTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerDistributedSupplierVenus)
				if err := _Comtroller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedSupplierVenus is a log parse operation binding the contract event 0xfa9d964d891991c113b49e3db1932abd6c67263d387119707aafdd6c4010a3a9.
//
// Solidity: event DistributedSupplierVenus(address indexed vToken, address indexed supplier, uint256 venusDelta, uint256 venusSupplyIndex)
func (_Comtroller *ComtrollerFilterer) ParseDistributedSupplierVenus(log types.Log) (*ComtrollerDistributedSupplierVenus, error) {
	event := new(ComtrollerDistributedSupplierVenus)
	if err := _Comtroller.contract.UnpackLog(event, "DistributedSupplierVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerDistributedVAIMinterVenusIterator is returned from FilterDistributedVAIMinterVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIMinterVenus events raised by the Comtroller contract.
type ComtrollerDistributedVAIMinterVenusIterator struct {
	Event *ComtrollerDistributedVAIMinterVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerDistributedVAIMinterVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerDistributedVAIMinterVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerDistributedVAIMinterVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerDistributedVAIMinterVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerDistributedVAIMinterVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerDistributedVAIMinterVenus represents a DistributedVAIMinterVenus event raised by the Comtroller contract.
type ComtrollerDistributedVAIMinterVenus struct {
	VaiMinter         common.Address
	VenusDelta        *big.Int
	VenusVAIMintIndex *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIMinterVenus is a free log retrieval operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Comtroller *ComtrollerFilterer) FilterDistributedVAIMinterVenus(opts *bind.FilterOpts, vaiMinter []common.Address) (*ComtrollerDistributedVAIMinterVenusIterator, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return &ComtrollerDistributedVAIMinterVenusIterator{contract: _Comtroller.contract, event: "DistributedVAIMinterVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIMinterVenus is a free log subscription operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Comtroller *ComtrollerFilterer) WatchDistributedVAIMinterVenus(opts *bind.WatchOpts, sink chan<- *ComtrollerDistributedVAIMinterVenus, vaiMinter []common.Address) (event.Subscription, error) {

	var vaiMinterRule []interface{}
	for _, vaiMinterItem := range vaiMinter {
		vaiMinterRule = append(vaiMinterRule, vaiMinterItem)
	}

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "DistributedVAIMinterVenus", vaiMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerDistributedVAIMinterVenus)
				if err := _Comtroller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedVAIMinterVenus is a log parse operation binding the contract event 0x2fb3baf25f3d9fc9f9eb9dfd7da8567731a91413437d91bc1b8a839d0a1ba88f.
//
// Solidity: event DistributedVAIMinterVenus(address indexed vaiMinter, uint256 venusDelta, uint256 venusVAIMintIndex)
func (_Comtroller *ComtrollerFilterer) ParseDistributedVAIMinterVenus(log types.Log) (*ComtrollerDistributedVAIMinterVenus, error) {
	event := new(ComtrollerDistributedVAIMinterVenus)
	if err := _Comtroller.contract.UnpackLog(event, "DistributedVAIMinterVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerDistributedVAIVaultVenusIterator is returned from FilterDistributedVAIVaultVenus and is used to iterate over the raw logs and unpacked data for DistributedVAIVaultVenus events raised by the Comtroller contract.
type ComtrollerDistributedVAIVaultVenusIterator struct {
	Event *ComtrollerDistributedVAIVaultVenus // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerDistributedVAIVaultVenusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerDistributedVAIVaultVenus)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerDistributedVAIVaultVenus)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerDistributedVAIVaultVenusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerDistributedVAIVaultVenusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerDistributedVAIVaultVenus represents a DistributedVAIVaultVenus event raised by the Comtroller contract.
type ComtrollerDistributedVAIVaultVenus struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedVAIVaultVenus is a free log retrieval operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Comtroller *ComtrollerFilterer) FilterDistributedVAIVaultVenus(opts *bind.FilterOpts) (*ComtrollerDistributedVAIVaultVenusIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return &ComtrollerDistributedVAIVaultVenusIterator{contract: _Comtroller.contract, event: "DistributedVAIVaultVenus", logs: logs, sub: sub}, nil
}

// WatchDistributedVAIVaultVenus is a free log subscription operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Comtroller *ComtrollerFilterer) WatchDistributedVAIVaultVenus(opts *bind.WatchOpts, sink chan<- *ComtrollerDistributedVAIVaultVenus) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "DistributedVAIVaultVenus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerDistributedVAIVaultVenus)
				if err := _Comtroller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedVAIVaultVenus is a log parse operation binding the contract event 0xf6d4b8f74d85a6e2d7a50225957b8a6cfec69ad92f5905627260541aa0a5565d.
//
// Solidity: event DistributedVAIVaultVenus(uint256 amount)
func (_Comtroller *ComtrollerFilterer) ParseDistributedVAIVaultVenus(log types.Log) (*ComtrollerDistributedVAIVaultVenus, error) {
	event := new(ComtrollerDistributedVAIVaultVenus)
	if err := _Comtroller.contract.UnpackLog(event, "DistributedVAIVaultVenus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Comtroller contract.
type ComtrollerFailureIterator struct {
	Event *ComtrollerFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerFailure represents a Failure event raised by the Comtroller contract.
type ComtrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comtroller *ComtrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*ComtrollerFailureIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &ComtrollerFailureIterator{contract: _Comtroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comtroller *ComtrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *ComtrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerFailure)
				if err := _Comtroller.contract.UnpackLog(event, "Failure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comtroller *ComtrollerFilterer) ParseFailure(log types.Log) (*ComtrollerFailure, error) {
	event := new(ComtrollerFailure)
	if err := _Comtroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Comtroller contract.
type ComtrollerMarketEnteredIterator struct {
	Event *ComtrollerMarketEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerMarketEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerMarketEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerMarketEntered represents a MarketEntered event raised by the Comtroller contract.
type ComtrollerMarketEntered struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Comtroller *ComtrollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*ComtrollerMarketEnteredIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &ComtrollerMarketEnteredIterator{contract: _Comtroller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Comtroller *ComtrollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *ComtrollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerMarketEntered)
				if err := _Comtroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address vToken, address account)
func (_Comtroller *ComtrollerFilterer) ParseMarketEntered(log types.Log) (*ComtrollerMarketEntered, error) {
	event := new(ComtrollerMarketEntered)
	if err := _Comtroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Comtroller contract.
type ComtrollerMarketExitedIterator struct {
	Event *ComtrollerMarketExited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerMarketExited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerMarketExited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerMarketExited represents a MarketExited event raised by the Comtroller contract.
type ComtrollerMarketExited struct {
	VToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Comtroller *ComtrollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*ComtrollerMarketExitedIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &ComtrollerMarketExitedIterator{contract: _Comtroller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Comtroller *ComtrollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *ComtrollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerMarketExited)
				if err := _Comtroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address vToken, address account)
func (_Comtroller *ComtrollerFilterer) ParseMarketExited(log types.Log) (*ComtrollerMarketExited, error) {
	event := new(ComtrollerMarketExited)
	if err := _Comtroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the Comtroller contract.
type ComtrollerMarketListedIterator struct {
	Event *ComtrollerMarketListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerMarketListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerMarketListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerMarketListed represents a MarketListed event raised by the Comtroller contract.
type ComtrollerMarketListed struct {
	VToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Comtroller *ComtrollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*ComtrollerMarketListedIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &ComtrollerMarketListedIterator{contract: _Comtroller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Comtroller *ComtrollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *ComtrollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerMarketListed)
				if err := _Comtroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address vToken)
func (_Comtroller *ComtrollerFilterer) ParseMarketListed(log types.Log) (*ComtrollerMarketListed, error) {
	event := new(ComtrollerMarketListed)
	if err := _Comtroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the Comtroller contract.
type ComtrollerNewBorrowCapIterator struct {
	Event *ComtrollerNewBorrowCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewBorrowCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewBorrowCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewBorrowCap represents a NewBorrowCap event raised by the Comtroller contract.
type ComtrollerNewBorrowCap struct {
	VToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Comtroller *ComtrollerFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, vToken []common.Address) (*ComtrollerNewBorrowCapIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewBorrowCapIterator{contract: _Comtroller.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Comtroller *ComtrollerFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *ComtrollerNewBorrowCap, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewBorrowCap", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewBorrowCap)
				if err := _Comtroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed vToken, uint256 newBorrowCap)
func (_Comtroller *ComtrollerFilterer) ParseNewBorrowCap(log types.Log) (*ComtrollerNewBorrowCap, error) {
	event := new(ComtrollerNewBorrowCap)
	if err := _Comtroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the Comtroller contract.
type ComtrollerNewBorrowCapGuardianIterator struct {
	Event *ComtrollerNewBorrowCapGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewBorrowCapGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewBorrowCapGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the Comtroller contract.
type ComtrollerNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comtroller *ComtrollerFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*ComtrollerNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewBorrowCapGuardianIterator{contract: _Comtroller.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comtroller *ComtrollerFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *ComtrollerNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewBorrowCapGuardian)
				if err := _Comtroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comtroller *ComtrollerFilterer) ParseNewBorrowCapGuardian(log types.Log) (*ComtrollerNewBorrowCapGuardian, error) {
	event := new(ComtrollerNewBorrowCapGuardian)
	if err := _Comtroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Comtroller contract.
type ComtrollerNewCloseFactorIterator struct {
	Event *ComtrollerNewCloseFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewCloseFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewCloseFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewCloseFactor represents a NewCloseFactor event raised by the Comtroller contract.
type ComtrollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comtroller *ComtrollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*ComtrollerNewCloseFactorIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewCloseFactorIterator{contract: _Comtroller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comtroller *ComtrollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *ComtrollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewCloseFactor)
				if err := _Comtroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comtroller *ComtrollerFilterer) ParseNewCloseFactor(log types.Log) (*ComtrollerNewCloseFactor, error) {
	event := new(ComtrollerNewCloseFactor)
	if err := _Comtroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Comtroller contract.
type ComtrollerNewCollateralFactorIterator struct {
	Event *ComtrollerNewCollateralFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewCollateralFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewCollateralFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewCollateralFactor represents a NewCollateralFactor event raised by the Comtroller contract.
type ComtrollerNewCollateralFactor struct {
	VToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comtroller *ComtrollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*ComtrollerNewCollateralFactorIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewCollateralFactorIterator{contract: _Comtroller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comtroller *ComtrollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *ComtrollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewCollateralFactor)
				if err := _Comtroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address vToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comtroller *ComtrollerFilterer) ParseNewCollateralFactor(log types.Log) (*ComtrollerNewCollateralFactor, error) {
	event := new(ComtrollerNewCollateralFactor)
	if err := _Comtroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Comtroller contract.
type ComtrollerNewLiquidationIncentiveIterator struct {
	Event *ComtrollerNewLiquidationIncentive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewLiquidationIncentive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewLiquidationIncentive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Comtroller contract.
type ComtrollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comtroller *ComtrollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*ComtrollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewLiquidationIncentiveIterator{contract: _Comtroller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comtroller *ComtrollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *ComtrollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewLiquidationIncentive)
				if err := _Comtroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comtroller *ComtrollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*ComtrollerNewLiquidationIncentive, error) {
	event := new(ComtrollerNewLiquidationIncentive)
	if err := _Comtroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Comtroller contract.
type ComtrollerNewPauseGuardianIterator struct {
	Event *ComtrollerNewPauseGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewPauseGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewPauseGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewPauseGuardian represents a NewPauseGuardian event raised by the Comtroller contract.
type ComtrollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comtroller *ComtrollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*ComtrollerNewPauseGuardianIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewPauseGuardianIterator{contract: _Comtroller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comtroller *ComtrollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *ComtrollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewPauseGuardian)
				if err := _Comtroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comtroller *ComtrollerFilterer) ParseNewPauseGuardian(log types.Log) (*ComtrollerNewPauseGuardian, error) {
	event := new(ComtrollerNewPauseGuardian)
	if err := _Comtroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Comtroller contract.
type ComtrollerNewPriceOracleIterator struct {
	Event *ComtrollerNewPriceOracle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewPriceOracle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewPriceOracle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewPriceOracle represents a NewPriceOracle event raised by the Comtroller contract.
type ComtrollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comtroller *ComtrollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*ComtrollerNewPriceOracleIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewPriceOracleIterator{contract: _Comtroller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comtroller *ComtrollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *ComtrollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewPriceOracle)
				if err := _Comtroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comtroller *ComtrollerFilterer) ParseNewPriceOracle(log types.Log) (*ComtrollerNewPriceOracle, error) {
	event := new(ComtrollerNewPriceOracle)
	if err := _Comtroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewTreasuryAddressIterator is returned from FilterNewTreasuryAddress and is used to iterate over the raw logs and unpacked data for NewTreasuryAddress events raised by the Comtroller contract.
type ComtrollerNewTreasuryAddressIterator struct {
	Event *ComtrollerNewTreasuryAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewTreasuryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewTreasuryAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewTreasuryAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewTreasuryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewTreasuryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewTreasuryAddress represents a NewTreasuryAddress event raised by the Comtroller contract.
type ComtrollerNewTreasuryAddress struct {
	OldTreasuryAddress common.Address
	NewTreasuryAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryAddress is a free log retrieval operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Comtroller *ComtrollerFilterer) FilterNewTreasuryAddress(opts *bind.FilterOpts) (*ComtrollerNewTreasuryAddressIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewTreasuryAddressIterator{contract: _Comtroller.contract, event: "NewTreasuryAddress", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryAddress is a free log subscription operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Comtroller *ComtrollerFilterer) WatchNewTreasuryAddress(opts *bind.WatchOpts, sink chan<- *ComtrollerNewTreasuryAddress) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewTreasuryAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewTreasuryAddress)
				if err := _Comtroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryAddress is a log parse operation binding the contract event 0x8de763046d7b8f08b6c3d03543de1d615309417842bb5d2d62f110f65809ddac.
//
// Solidity: event NewTreasuryAddress(address oldTreasuryAddress, address newTreasuryAddress)
func (_Comtroller *ComtrollerFilterer) ParseNewTreasuryAddress(log types.Log) (*ComtrollerNewTreasuryAddress, error) {
	event := new(ComtrollerNewTreasuryAddress)
	if err := _Comtroller.contract.UnpackLog(event, "NewTreasuryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewTreasuryGuardianIterator is returned from FilterNewTreasuryGuardian and is used to iterate over the raw logs and unpacked data for NewTreasuryGuardian events raised by the Comtroller contract.
type ComtrollerNewTreasuryGuardianIterator struct {
	Event *ComtrollerNewTreasuryGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewTreasuryGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewTreasuryGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewTreasuryGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewTreasuryGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewTreasuryGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewTreasuryGuardian represents a NewTreasuryGuardian event raised by the Comtroller contract.
type ComtrollerNewTreasuryGuardian struct {
	OldTreasuryGuardian common.Address
	NewTreasuryGuardian common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryGuardian is a free log retrieval operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Comtroller *ComtrollerFilterer) FilterNewTreasuryGuardian(opts *bind.FilterOpts) (*ComtrollerNewTreasuryGuardianIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewTreasuryGuardianIterator{contract: _Comtroller.contract, event: "NewTreasuryGuardian", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryGuardian is a free log subscription operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Comtroller *ComtrollerFilterer) WatchNewTreasuryGuardian(opts *bind.WatchOpts, sink chan<- *ComtrollerNewTreasuryGuardian) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewTreasuryGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewTreasuryGuardian)
				if err := _Comtroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryGuardian is a log parse operation binding the contract event 0x29f06ea15931797ebaed313d81d100963dc22cb213cb4ce2737b5a62b1a8b1e8.
//
// Solidity: event NewTreasuryGuardian(address oldTreasuryGuardian, address newTreasuryGuardian)
func (_Comtroller *ComtrollerFilterer) ParseNewTreasuryGuardian(log types.Log) (*ComtrollerNewTreasuryGuardian, error) {
	event := new(ComtrollerNewTreasuryGuardian)
	if err := _Comtroller.contract.UnpackLog(event, "NewTreasuryGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewTreasuryPercentIterator is returned from FilterNewTreasuryPercent and is used to iterate over the raw logs and unpacked data for NewTreasuryPercent events raised by the Comtroller contract.
type ComtrollerNewTreasuryPercentIterator struct {
	Event *ComtrollerNewTreasuryPercent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewTreasuryPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewTreasuryPercent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewTreasuryPercent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewTreasuryPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewTreasuryPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewTreasuryPercent represents a NewTreasuryPercent event raised by the Comtroller contract.
type ComtrollerNewTreasuryPercent struct {
	OldTreasuryPercent *big.Int
	NewTreasuryPercent *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewTreasuryPercent is a free log retrieval operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Comtroller *ComtrollerFilterer) FilterNewTreasuryPercent(opts *bind.FilterOpts) (*ComtrollerNewTreasuryPercentIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewTreasuryPercentIterator{contract: _Comtroller.contract, event: "NewTreasuryPercent", logs: logs, sub: sub}, nil
}

// WatchNewTreasuryPercent is a free log subscription operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Comtroller *ComtrollerFilterer) WatchNewTreasuryPercent(opts *bind.WatchOpts, sink chan<- *ComtrollerNewTreasuryPercent) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewTreasuryPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewTreasuryPercent)
				if err := _Comtroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewTreasuryPercent is a log parse operation binding the contract event 0x0893f8f4101baaabbeb513f96761e7a36eb837403c82cc651c292a4abdc94ed7.
//
// Solidity: event NewTreasuryPercent(uint256 oldTreasuryPercent, uint256 newTreasuryPercent)
func (_Comtroller *ComtrollerFilterer) ParseNewTreasuryPercent(log types.Log) (*ComtrollerNewTreasuryPercent, error) {
	event := new(ComtrollerNewTreasuryPercent)
	if err := _Comtroller.contract.UnpackLog(event, "NewTreasuryPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewVAIControllerIterator is returned from FilterNewVAIController and is used to iterate over the raw logs and unpacked data for NewVAIController events raised by the Comtroller contract.
type ComtrollerNewVAIControllerIterator struct {
	Event *ComtrollerNewVAIController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewVAIControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewVAIController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewVAIController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewVAIControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewVAIControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewVAIController represents a NewVAIController event raised by the Comtroller contract.
type ComtrollerNewVAIController struct {
	OldVAIController common.Address
	NewVAIController common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewVAIController is a free log retrieval operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Comtroller *ComtrollerFilterer) FilterNewVAIController(opts *bind.FilterOpts) (*ComtrollerNewVAIControllerIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewVAIControllerIterator{contract: _Comtroller.contract, event: "NewVAIController", logs: logs, sub: sub}, nil
}

// WatchNewVAIController is a free log subscription operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Comtroller *ComtrollerFilterer) WatchNewVAIController(opts *bind.WatchOpts, sink chan<- *ComtrollerNewVAIController) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewVAIController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewVAIController)
				if err := _Comtroller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIController is a log parse operation binding the contract event 0xe1ddcb2dab8e5b03cfc8c67a0d5861d91d16f7bd2612fd381faf4541d212c9b2.
//
// Solidity: event NewVAIController(address oldVAIController, address newVAIController)
func (_Comtroller *ComtrollerFilterer) ParseNewVAIController(log types.Log) (*ComtrollerNewVAIController, error) {
	event := new(ComtrollerNewVAIController)
	if err := _Comtroller.contract.UnpackLog(event, "NewVAIController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewVAIMintRateIterator is returned from FilterNewVAIMintRate and is used to iterate over the raw logs and unpacked data for NewVAIMintRate events raised by the Comtroller contract.
type ComtrollerNewVAIMintRateIterator struct {
	Event *ComtrollerNewVAIMintRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewVAIMintRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewVAIMintRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewVAIMintRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewVAIMintRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewVAIMintRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewVAIMintRate represents a NewVAIMintRate event raised by the Comtroller contract.
type ComtrollerNewVAIMintRate struct {
	OldVAIMintRate *big.Int
	NewVAIMintRate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewVAIMintRate is a free log retrieval operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Comtroller *ComtrollerFilterer) FilterNewVAIMintRate(opts *bind.FilterOpts) (*ComtrollerNewVAIMintRateIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewVAIMintRateIterator{contract: _Comtroller.contract, event: "NewVAIMintRate", logs: logs, sub: sub}, nil
}

// WatchNewVAIMintRate is a free log subscription operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Comtroller *ComtrollerFilterer) WatchNewVAIMintRate(opts *bind.WatchOpts, sink chan<- *ComtrollerNewVAIMintRate) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewVAIMintRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewVAIMintRate)
				if err := _Comtroller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIMintRate is a log parse operation binding the contract event 0x73747d68b346dce1e932bcd238282e7ac84c01569e1f8d0469c222fdc6e9d5a4.
//
// Solidity: event NewVAIMintRate(uint256 oldVAIMintRate, uint256 newVAIMintRate)
func (_Comtroller *ComtrollerFilterer) ParseNewVAIMintRate(log types.Log) (*ComtrollerNewVAIMintRate, error) {
	event := new(ComtrollerNewVAIMintRate)
	if err := _Comtroller.contract.UnpackLog(event, "NewVAIMintRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewVAIVaultInfoIterator is returned from FilterNewVAIVaultInfo and is used to iterate over the raw logs and unpacked data for NewVAIVaultInfo events raised by the Comtroller contract.
type ComtrollerNewVAIVaultInfoIterator struct {
	Event *ComtrollerNewVAIVaultInfo // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewVAIVaultInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewVAIVaultInfo)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewVAIVaultInfo)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewVAIVaultInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewVAIVaultInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewVAIVaultInfo represents a NewVAIVaultInfo event raised by the Comtroller contract.
type ComtrollerNewVAIVaultInfo struct {
	Vault             common.Address
	ReleaseStartBlock *big.Int
	ReleaseInterval   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewVAIVaultInfo is a free log retrieval operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Comtroller *ComtrollerFilterer) FilterNewVAIVaultInfo(opts *bind.FilterOpts) (*ComtrollerNewVAIVaultInfoIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewVAIVaultInfoIterator{contract: _Comtroller.contract, event: "NewVAIVaultInfo", logs: logs, sub: sub}, nil
}

// WatchNewVAIVaultInfo is a free log subscription operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Comtroller *ComtrollerFilterer) WatchNewVAIVaultInfo(opts *bind.WatchOpts, sink chan<- *ComtrollerNewVAIVaultInfo) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewVAIVaultInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewVAIVaultInfo)
				if err := _Comtroller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVAIVaultInfo is a log parse operation binding the contract event 0x7059037d74ee16b0fb06a4a30f3348dd2635f301f92e373c92899a25a522ef6e.
//
// Solidity: event NewVAIVaultInfo(address vault_, uint256 releaseStartBlock_, uint256 releaseInterval_)
func (_Comtroller *ComtrollerFilterer) ParseNewVAIVaultInfo(log types.Log) (*ComtrollerNewVAIVaultInfo, error) {
	event := new(ComtrollerNewVAIVaultInfo)
	if err := _Comtroller.contract.UnpackLog(event, "NewVAIVaultInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewVenusVAIRateIterator is returned from FilterNewVenusVAIRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIRate events raised by the Comtroller contract.
type ComtrollerNewVenusVAIRateIterator struct {
	Event *ComtrollerNewVenusVAIRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewVenusVAIRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewVenusVAIRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewVenusVAIRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewVenusVAIRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewVenusVAIRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewVenusVAIRate represents a NewVenusVAIRate event raised by the Comtroller contract.
type ComtrollerNewVenusVAIRate struct {
	OldVenusVAIRate *big.Int
	NewVenusVAIRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIRate is a free log retrieval operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Comtroller *ComtrollerFilterer) FilterNewVenusVAIRate(opts *bind.FilterOpts) (*ComtrollerNewVenusVAIRateIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewVenusVAIRateIterator{contract: _Comtroller.contract, event: "NewVenusVAIRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIRate is a free log subscription operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Comtroller *ComtrollerFilterer) WatchNewVenusVAIRate(opts *bind.WatchOpts, sink chan<- *ComtrollerNewVenusVAIRate) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewVenusVAIRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewVenusVAIRate)
				if err := _Comtroller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVenusVAIRate is a log parse operation binding the contract event 0x75c84862cb29e997a2ed3ab3c3db0f5af24a181e6bf58897c5ea676668511c19.
//
// Solidity: event NewVenusVAIRate(uint256 oldVenusVAIRate, uint256 newVenusVAIRate)
func (_Comtroller *ComtrollerFilterer) ParseNewVenusVAIRate(log types.Log) (*ComtrollerNewVenusVAIRate, error) {
	event := new(ComtrollerNewVenusVAIRate)
	if err := _Comtroller.contract.UnpackLog(event, "NewVenusVAIRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerNewVenusVAIVaultRateIterator is returned from FilterNewVenusVAIVaultRate and is used to iterate over the raw logs and unpacked data for NewVenusVAIVaultRate events raised by the Comtroller contract.
type ComtrollerNewVenusVAIVaultRateIterator struct {
	Event *ComtrollerNewVenusVAIVaultRate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerNewVenusVAIVaultRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerNewVenusVAIVaultRate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerNewVenusVAIVaultRate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerNewVenusVAIVaultRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerNewVenusVAIVaultRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerNewVenusVAIVaultRate represents a NewVenusVAIVaultRate event raised by the Comtroller contract.
type ComtrollerNewVenusVAIVaultRate struct {
	OldVenusVAIVaultRate *big.Int
	NewVenusVAIVaultRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewVenusVAIVaultRate is a free log retrieval operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Comtroller *ComtrollerFilterer) FilterNewVenusVAIVaultRate(opts *bind.FilterOpts) (*ComtrollerNewVenusVAIVaultRateIterator, error) {

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return &ComtrollerNewVenusVAIVaultRateIterator{contract: _Comtroller.contract, event: "NewVenusVAIVaultRate", logs: logs, sub: sub}, nil
}

// WatchNewVenusVAIVaultRate is a free log subscription operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Comtroller *ComtrollerFilterer) WatchNewVenusVAIVaultRate(opts *bind.WatchOpts, sink chan<- *ComtrollerNewVenusVAIVaultRate) (event.Subscription, error) {

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "NewVenusVAIVaultRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerNewVenusVAIVaultRate)
				if err := _Comtroller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewVenusVAIVaultRate is a log parse operation binding the contract event 0xe81d4ac15e5afa1e708e66664eddc697177423d950d133bda8262d8885e6da3b.
//
// Solidity: event NewVenusVAIVaultRate(uint256 oldVenusVAIVaultRate, uint256 newVenusVAIVaultRate)
func (_Comtroller *ComtrollerFilterer) ParseNewVenusVAIVaultRate(log types.Log) (*ComtrollerNewVenusVAIVaultRate, error) {
	event := new(ComtrollerNewVenusVAIVaultRate)
	if err := _Comtroller.contract.UnpackLog(event, "NewVenusVAIVaultRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComtrollerVenusSpeedUpdatedIterator is returned from FilterVenusSpeedUpdated and is used to iterate over the raw logs and unpacked data for VenusSpeedUpdated events raised by the Comtroller contract.
type ComtrollerVenusSpeedUpdatedIterator struct {
	Event *ComtrollerVenusSpeedUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ComtrollerVenusSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComtrollerVenusSpeedUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ComtrollerVenusSpeedUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ComtrollerVenusSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComtrollerVenusSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComtrollerVenusSpeedUpdated represents a VenusSpeedUpdated event raised by the Comtroller contract.
type ComtrollerVenusSpeedUpdated struct {
	VToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVenusSpeedUpdated is a free log retrieval operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Comtroller *ComtrollerFilterer) FilterVenusSpeedUpdated(opts *bind.FilterOpts, vToken []common.Address) (*ComtrollerVenusSpeedUpdatedIterator, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Comtroller.contract.FilterLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return &ComtrollerVenusSpeedUpdatedIterator{contract: _Comtroller.contract, event: "VenusSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchVenusSpeedUpdated is a free log subscription operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Comtroller *ComtrollerFilterer) WatchVenusSpeedUpdated(opts *bind.WatchOpts, sink chan<- *ComtrollerVenusSpeedUpdated, vToken []common.Address) (event.Subscription, error) {

	var vTokenRule []interface{}
	for _, vTokenItem := range vToken {
		vTokenRule = append(vTokenRule, vTokenItem)
	}

	logs, sub, err := _Comtroller.contract.WatchLogs(opts, "VenusSpeedUpdated", vTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComtrollerVenusSpeedUpdated)
				if err := _Comtroller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVenusSpeedUpdated is a log parse operation binding the contract event 0x2a0ce45ba05a83e75ba21e1a10d6b48a8395028cc6e1ae66f6c313648379d548.
//
// Solidity: event VenusSpeedUpdated(address indexed vToken, uint256 newSpeed)
func (_Comtroller *ComtrollerFilterer) ParseVenusSpeedUpdated(log types.Log) (*ComtrollerVenusSpeedUpdated, error) {
	event := new(ComtrollerVenusSpeedUpdated)
	if err := _Comtroller.contract.UnpackLog(event, "VenusSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
