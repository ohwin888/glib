package transcategory

// TransactionCategory represents transaction main category types
type TransactionCategory int

var TransactionCategories = []TransactionCategory{
	TransactionCategoryCapitalSwitch,
	TransactionCategoryUserRecharge,
	TransactionCategoryUserWithdraw,
	TransactionCategoryBankMerchantSettlement,
	TransactionCategoryCapitalRevise,
	TransactionCategoryActivity,
	TransactionCategoryGoldReturn,
	TransactionCategoryRebate,
	TransactionCategoryInterest,
	TransactionCategoryTask,
	TransactionCategoryVipReward,
	TransactionCategoryRechargeBonus,
	TransactionCategoryClub,
	TransactionCategoryReward,
	TransactionCategoryGuaranteeClaim,
	TransactionCategoryProvidentFund,
	TransactionCategoryBlindBoxLottery,
	TransactionCategoryGameBet,
}

const (
	_                                         TransactionCategory = iota
	TransactionCategoryCapitalSwitch                              // Fund Switching
	TransactionCategoryUserRecharge                               // User Recharge
	TransactionCategoryUserWithdraw                               // User Withdrawal
	TransactionCategoryBankMerchantSettlement                     // Bank Merchant Settlement
	TransactionCategoryCapitalRevise                              // Capital Revision
	TransactionCategoryActivity                                   // Activity
	TransactionCategoryGoldReturn                                 // Gold Return
	TransactionCategoryRebate                                     // Rebate
	TransactionCategoryInterest                                   // Interest
	TransactionCategoryTask                                       // Task
	TransactionCategoryVipReward                                  // VIP Reward
	TransactionCategoryRechargeBonus                              // Recharge Bonus
	TransactionCategoryClub                                       // Club
	TransactionCategoryReward                                     // Reward
	TransactionCategoryGuaranteeClaim                             // Guarantee Claim
	TransactionCategoryProvidentFund                              // Provident Fund
	TransactionCategoryBlindBoxLottery                            // Blind Box Lottery
	TransactionCategoryGameBet                                    // Game Bet
)

func (t TransactionCategory) Int() int {
	return int(t)
}

func (t TransactionCategory) Int64() int64 {
	return int64(t)
}

func (t TransactionCategory) String() string {
	switch t {
	case TransactionCategoryCapitalSwitch:
		return "Fund Switching"
	case TransactionCategoryUserRecharge:
		return "User Recharge"
	case TransactionCategoryUserWithdraw:
		return "User Withdrawal"
	case TransactionCategoryBankMerchantSettlement:
		return "Bank Merchant Settlement"
	case TransactionCategoryCapitalRevise:
		return "Capital Revision"
	case TransactionCategoryActivity:
		return "Activity"
	case TransactionCategoryGoldReturn:
		return "Gold Return"
	case TransactionCategoryRebate:
		return "Rebate"
	case TransactionCategoryInterest:
		return "Interest"
	case TransactionCategoryTask:
		return "Task"
	case TransactionCategoryVipReward:
		return "VIP Reward"
	case TransactionCategoryRechargeBonus:
		return "Recharge Bonus"
	case TransactionCategoryClub:
		return "Club"
	case TransactionCategoryReward:
		return "Reward"
	case TransactionCategoryGuaranteeClaim:
		return "Guarantee Claim"
	case TransactionCategoryProvidentFund:
		return "Provident Fund"
	case TransactionCategoryBlindBoxLottery:
		return "Blind Box Lottery"
	case TransactionCategoryGameBet:
		return "Game Bet"
	default:
		return ""
	}
}

// SubCategory 获取子类
func (t TransactionCategory) SubCategory() []TransactionSubCategory {
	switch t {
	case TransactionCategoryCapitalSwitch:
		return []TransactionSubCategory{
			TransactionSubCategoryWithdrawal, TransactionSubCategoryDeposits,
		}
	case TransactionCategoryUserRecharge:
		return []TransactionSubCategory{
			TransactionSubCategoryUalaTransfer, TransactionSubCategoryMercadoPagoTransfer, TransactionSubCategoryUPDAYTransfer, TransactionSubCategoryRechargeOnline,
		}
	case TransactionCategoryUserWithdraw:
		return []TransactionSubCategory{
			TransactionSubCategoryWithdrawFrozen, TransactionSubCategoryWithdrawReject,
			TransactionSubCategoryWithdrawDefrost, TransactionSubCategoryWithdrawSucceed,
		}
	case TransactionCategoryBankMerchantSettlement:
		return []TransactionSubCategory{
			TransactionSubCategoryBankMerchantTransfer, TransactionSubCategoryBankMerchantAddAmount, TransactionSubCategoryBankMerchantRecharge, // 银商充值
			TransactionSubCategoryBankMerchantGiveUser, TransactionSubCategoryBankMerchantSubtractAmount,
		}
	case TransactionCategoryCapitalRevise:
		return []TransactionSubCategory{
			TransactionSubCategoryManualAddAmount, TransactionSubCategoryManualAddOrder, TransactionSubCategoryBalanceRevise,
			TransactionSubCategoryManualAddRewardAmount, TransactionSubCategoryManualSubtractAmount, TransactionSubCategoryDeductAll, // 扣除全部资产
			TransactionSubCategorySurDeduct, TransactionSubCategoryManualPullBack, TransactionSubCategoryDeductExcessProfit,
		}
	case TransactionCategoryActivity:
		return []TransactionSubCategory{
			TransactionSubCategoryAgentActivity, TransactionSubCategoryLuckyBetActivity, TransactionSubCategoryInvestActivity,
			TransactionSubCategoryNewcomerRewardActivity, TransactionSubCategoryBenefitActivity, TransactionSubCategoryPromoteActivity,
			TransactionSubCategoryFeedbackRewardActivity, TransactionSubCategoryRedPacketActivity, TransactionSubCategoryBetActivity,
			TransactionSubCategoryLotteryAssistanceActivity, TransactionSubCategoryRankActivity, TransactionSubCategoryCustomizeActivity,
			TransactionSubCategoryBargainActivity, TransactionSubCategorySpinActivity, TransactionSubCategoryChannelRewardActivity,
			TransactionSubCategoryWordCollectionActivity, TransactionSubCategoryQuizActivity,
		}
	case TransactionCategoryGoldReturn:
		return []TransactionSubCategory{
			TransactionSubCategoryServiceChargeReceive,
		}
	case TransactionCategoryRebate:
		return []TransactionSubCategory{
			TransactionSubCategoryRebateSend, TransactionSubCategoryRebateReceive,
		}
	case TransactionCategoryInterest:
		return []TransactionSubCategory{
			TransactionSubCategoryInterestProfit, TransactionSubCategoryHallToInterest,
			TransactionSubCategoryInterestManualPullback, TransactionSubCategoryInterestToHall,
		}
	case TransactionCategoryTask:
		return []TransactionSubCategory{
			TransactionSubCategoryDailyTask, TransactionSubCategoryWeekLyTask, TransactionSubCategoryNewcomerProfitTask,
			TransactionSubCategoryAlivenessBox, TransactionSubCategorySecretTask,
		}
	case TransactionCategoryVipReward:
		return []TransactionSubCategory{
			TransactionSubCategoryVipMonthlyReward, TransactionSubCategoryVipDailyReward,
			TransactionSubCategoryVipWeeklyReward, TransactionSubCategoryVipUpgradeReward,
		}
	case TransactionCategoryRechargeBonus:
		return []TransactionSubCategory{
			TransactionSubCategoryRechargeBonus,
		}
	case TransactionCategoryClub:
		return []TransactionSubCategory{
			TransactionSubCategoryManualPullbackClub, TransactionSubCategoryHallToClub, TransactionSubCategoryClubToHall,
		}
	case TransactionCategoryReward:
		return []TransactionSubCategory{
			TransactionSubCategoryGiveUpReward, TransactionSubCategoryTransferOut, TransactionSubCategoryTransferIn,
		}
	case TransactionCategoryGuaranteeClaim:
		return []TransactionSubCategory{
			TransactionSubCategoryClaimFreeze, TransactionSubCategoryClaimDefrost, TransactionSubCategoryClaimFee,
			TransactionSubCategoryClaimScoresIncrease, TransactionSubCategoryClaimScoresDecrease,
		}
	case TransactionCategoryProvidentFund:
		return []TransactionSubCategory{
			TransactionSubCategoryRechargeBonus,
		}
	case TransactionCategoryBlindBoxLottery:
		return []TransactionSubCategory{
			TransactionSubCategoryBlindBoxLotteryDeduct, TransactionSubCategoryBlindBoxLotteryReward,
		}
	case TransactionCategoryGameBet:
		return []TransactionSubCategory{
			TransactionSubCategoryGameBet,
			TransactionSubCategoryGameBetSettlement,
			TransactionSubCategoryGameBetCancel,
		}
	default:
		return []TransactionSubCategory{}
	}
}

// TransactionSubCategory represents transaction sub-category types
type TransactionSubCategory int

// Capital Switch Subcategories
const (
	SubCategoryCapitalSwitch         TransactionSubCategory = 1000
	TransactionSubCategoryWithdrawal                        = iota + SubCategoryCapitalSwitch // Transfer Out
	TransactionSubCategoryDeposits                                                            // Transfer In
)

// User Recharge Subcategories
const (
	SubCategoryUserRecharge                   TransactionSubCategory = 2000
	TransactionSubCategoryUalaTransfer                               = iota + SubCategoryUserRecharge // Uala Transfer
	TransactionSubCategoryMercadoPagoTransfer                                                         // Mercado Pago Transfer
	TransactionSubCategoryUPDAYTransfer                                                               // UPDAY Transfer
	TransactionSubCategoryRechargeOnline                                                              // Online Recharge
)

// User Withdrawal Subcategories
const (
	SubCategoryUserWithdraw               TransactionSubCategory = 3000
	TransactionSubCategoryWithdrawFrozen                         = iota + SubCategoryUserWithdraw // Withdrawal Freeze
	TransactionSubCategoryWithdrawReject                                                          // Withdrawal Rejection
	TransactionSubCategoryWithdrawDefrost                                                         // Withdrawal Unfreeze
	TransactionSubCategoryWithdrawSucceed                                                         // Withdrawal Success
)

// Bank Merchant Settlement Subcategories
const (
	SubCategoryBankMerchantSettlement                TransactionSubCategory = 4000
	TransactionSubCategoryBankMerchantTransfer                              = iota + SubCategoryBankMerchantSettlement // Transfer to Others
	TransactionSubCategoryBankMerchantAddAmount                                                                        // Merchant Addition
	TransactionSubCategoryBankMerchantRecharge                                                                         // Merchant Recharge
	TransactionSubCategoryBankMerchantGiveUser                                                                         // Merchant User Gift
	TransactionSubCategoryBankMerchantSubtractAmount                                                                   // Merchant Deduction
)

// Capital Revision Subcategories
const (
	SubCategoryCapitalRevise                    TransactionSubCategory = 5000
	TransactionSubCategoryManualAddAmount                              = iota + SubCategoryCapitalRevise // Manual Addition
	TransactionSubCategoryManualAddOrder                                                                 // Manual Order Adjustment
	TransactionSubCategoryBalanceRevise                                                                  // Negative Balance Correction
	TransactionSubCategoryManualAddRewardAmount                                                          // Reward Manual Addition
	TransactionSubCategoryManualSubtractAmount                                                           // Manual Deduction
	TransactionSubCategoryDeductAll                                                                      // Full Asset Deduction
	TransactionSubCategorySurDeduct                                                                      // Additional Deduction
	TransactionSubCategoryManualPullBack                                                                 // Manual Recall
	TransactionSubCategoryDeductExcessProfit                                                             // Excess Profit Deduction
)

// Activity Subcategories
const (
	SubCategoryActivity                             TransactionSubCategory = 6000
	TransactionSubCategoryAgentActivity                                    = iota + SubCategoryActivity // Agent Activity
	TransactionSubCategoryLuckyBetActivity                                                              // Lucky Bet Activity
	TransactionSubCategoryInvestActivity                                                                // Investment Activity
	TransactionSubCategoryNewcomerRewardActivity                                                        // Newcomer Bonus
	TransactionSubCategoryBenefitActivity                                                               // Relief Fund Activity
	TransactionSubCategoryPromoteActivity                                                               // Promotion Activity
	TransactionSubCategoryFeedbackRewardActivity                                                        // Feedback Reward Activity
	TransactionSubCategoryRedPacketActivity                                                             // Red Packet Activity
	TransactionSubCategoryBetActivity                                                                   // Betting Activity
	TransactionSubCategoryLotteryAssistanceActivity                                                     // Lottery Assistance
	TransactionSubCategoryRankActivity                                                                  // Ranking Activity
	TransactionSubCategoryCustomizeActivity                                                             // Custom Activity
	TransactionSubCategoryBargainActivity                                                               // Bargain Activity
	TransactionSubCategorySpinActivity                                                                  // Spin Activity
	TransactionSubCategoryChannelRewardActivity                                                         // Channel Reward
	TransactionSubCategoryWordCollectionActivity                                                        // Word Collection
	TransactionSubCategoryQuizActivity                                                                  // Quiz
	TransactionSubCategoryRechargeActivity                                                              // 充值活动
	TransactionSubCategorySignInActivity                                                                // 签到活动// Activity
)

// Gold Return Subcategories
const (
	SubCategoryServiceCharge                   TransactionSubCategory = 7000
	TransactionSubCategoryServiceChargeReceive                        = iota + SubCategoryServiceCharge // Service Charge Receipt
)

// Rebate Subcategories
const (
	SubCategoryRebate                   TransactionSubCategory = 8000
	TransactionSubCategoryRebateSend                           = iota + SubCategoryRebate // Rebate Distribution
	TransactionSubCategoryRebateReceive                                                   // Rebate Receipt
)

// Interest Subcategories
const (
	SubCategoryInterest                          TransactionSubCategory = 9000
	TransactionSubCategoryInterestProfit                                = iota + SubCategoryInterest // Interest Earnings
	TransactionSubCategoryHallToInterest                                                             // Hall to Interest Transfer
	TransactionSubCategoryInterestManualPullback                                                     // Manual Recall - Interest
	TransactionSubCategoryInterestToHall                                                             // Interest to Hall Transfer
)

// Task Subcategories
const (
	SubCategoryTask                          TransactionSubCategory = 10000
	TransactionSubCategoryDailyTask                                 = iota + SubCategoryTask // Daily Task
	TransactionSubCategoryWeekLyTask                                                         // Weekly Task
	TransactionSubCategoryNewcomerProfitTask                                                 // Newcomer Benefits
	TransactionSubCategoryAlivenessBox                                                       // Activity Box
	TransactionSubCategorySecretTask                                                         // Secret Task
)

// VIP Reward Subcategories
const (
	SubCategoryVipReward                   TransactionSubCategory = 11000
	TransactionSubCategoryVipMonthlyReward                        = iota + SubCategoryVipReward // VIP Monthly Reward
	TransactionSubCategoryVipDailyReward                                                        // VIP Daily Reward
	TransactionSubCategoryVipWeeklyReward                                                       // VIP Weekly Reward
	TransactionSubCategoryVipUpgradeReward                                                      // VIP Upgrade Reward
)

// Recharge Bonus Subcategories
const (
	SubCategoryRechargeBonus            TransactionSubCategory = 12000
	TransactionSubCategoryRechargeBonus                        = iota + SubCategoryRechargeBonus // Bonus
)

// Club Subcategories
const (
	SubCategoryClub                          TransactionSubCategory = 13000
	TransactionSubCategoryManualPullbackClub                        = iota + SubCategoryClub // Manual Recall - Club
	TransactionSubCategoryHallToClub                                                         // Hall to Club Transfer
	TransactionSubCategoryClubToHall                                                         // Club to Hall Transfer
)

// Reward Subcategories
const (
	SubCategoryReward                  TransactionSubCategory = 14000
	TransactionSubCategoryGiveUpReward                        = iota + SubCategoryReward // Reward Waiver
	TransactionSubCategoryTransferOut                                                    // Transfer Out
	TransactionSubCategoryTransferIn                                                     // Transfer In
)

// Guarantee Claim Subcategories
const (
	SubCategoryGuaranteeClaim                 TransactionSubCategory = 15000
	TransactionSubCategoryClaimFreeze                                = iota + SubCategoryGuaranteeClaim // Claim Freeze
	TransactionSubCategoryClaimDefrost                                                                  // Claim Unfreeze
	TransactionSubCategoryClaimFee                                                                      // Claim Fee
	TransactionSubCategoryClaimScoresIncrease                                                           // Score Increase
	TransactionSubCategoryClaimScoresDecrease                                                           // Score Decrease
)

// Provident Fund Subcategories
const (
	SubCategoryProvidentFund                   TransactionSubCategory = 16000
	TransactionSubCategoryProvidentFundReceive                        = iota + SubCategoryProvidentFund // Provident Fund Receipt
)

// Blind Box Lottery Subcategories
const (
	SubCategoryBlindBoxLottery                  TransactionSubCategory = 17000
	TransactionSubCategoryBlindBoxLotteryDeduct                        = iota + SubCategoryBlindBoxLottery // Blind Box Deduction
	TransactionSubCategoryBlindBoxLotteryReward                                                            // Blind Box Reward
)

// Game Bet Subcategories
const (
	SubCategoryGameBetSettlement            TransactionSubCategory = 18000
	TransactionSubCategoryGameBet                                  = iota + SubCategoryGameBetSettlement // Game Bet
	TransactionSubCategoryGameBetSettlement                                                              // Bet Settlement
	TransactionSubCategoryGameBetCancel                                                                  // Bet Cancellation
)

func (t TransactionSubCategory) Int() int {
	return int(t)
}

func (t TransactionSubCategory) Int64() int64 {
	return int64(t)
}

func (t TransactionSubCategory) String() string {
	switch t {
	case TransactionSubCategoryWithdrawal:
		return "Transfer Out"
	case TransactionSubCategoryDeposits:
		return "Transfer In"
	case TransactionSubCategoryUalaTransfer:
		return "Uala Transfer"
	case TransactionSubCategoryMercadoPagoTransfer:
		return "Mercado Pago Transfer"
	case TransactionSubCategoryUPDAYTransfer:
		return "UPDAY Transfer"
	case TransactionSubCategoryWithdrawFrozen:
		return "Withdrawal Freeze"
	case TransactionSubCategoryWithdrawReject:
		return "Withdrawal Rejection"
	case TransactionSubCategoryWithdrawDefrost:
		return "Withdrawal Unfreeze"
	case TransactionSubCategoryWithdrawSucceed:
		return "Withdrawal Success"
	case TransactionSubCategoryBankMerchantTransfer:
		return "Transfer to Others"
	case TransactionSubCategoryBankMerchantAddAmount:
		return "Merchant Addition"
	case TransactionSubCategoryBankMerchantRecharge:
		return "Merchant Recharge"
	case TransactionSubCategoryBankMerchantGiveUser:
		return "Merchant User Gift"
	case TransactionSubCategoryBankMerchantSubtractAmount:
		return "Merchant Deduction"
	case TransactionSubCategoryManualAddAmount:
		return "Manual Addition"
	case TransactionSubCategoryManualAddOrder:
		return "Manual Order Adjustment"
	case TransactionSubCategoryBalanceRevise:
		return "Negative Balance Correction"
	case TransactionSubCategoryManualAddRewardAmount:
		return "Reward Manual Addition"
	case TransactionSubCategoryManualSubtractAmount:
		return "Manual Deduction"
	case TransactionSubCategoryDeductAll:
		return "Full Asset Deduction"
	case TransactionSubCategorySurDeduct:
		return "Additional Deduction"
	case TransactionSubCategoryManualPullBack:
		return "Manual Recall"
	case TransactionSubCategoryDeductExcessProfit:
		return "Excess Profit Deduction"
	case TransactionSubCategoryAgentActivity:
		return "Agent Activity"
	case TransactionSubCategoryLuckyBetActivity:
		return "Lucky Bet Activity"
	case TransactionSubCategoryInvestActivity:
		return "Investment Activity"
	case TransactionSubCategoryNewcomerRewardActivity:
		return "Newcomer Bonus"
	case TransactionSubCategoryBenefitActivity:
		return "Relief Fund Activity"
	case TransactionSubCategoryPromoteActivity:
		return "Promotion Activity"
	case TransactionSubCategoryFeedbackRewardActivity:
		return "Feedback Reward Activity"
	case TransactionSubCategoryRedPacketActivity:
		return "Red Packet Activity"
	case TransactionSubCategoryBetActivity:
		return "Betting Activity"
	case TransactionSubCategoryLotteryAssistanceActivity:
		return "Lottery Assistance"
	case TransactionSubCategoryRankActivity:
		return "Ranking Activity"
	case TransactionSubCategoryCustomizeActivity:
		return "Custom Activity"
	case TransactionSubCategoryBargainActivity:
		return "Bargain Activity"
	case TransactionSubCategorySpinActivity:
		return "Spin Activity"
	case TransactionSubCategoryChannelRewardActivity:
		return "Channel Reward"
	case TransactionSubCategoryWordCollectionActivity:
		return "Word Collection"
	case TransactionSubCategoryQuizActivity:
		return "Quiz Activity"
	case TransactionSubCategoryServiceChargeReceive:
		return "Service Charge Receipt"
	case TransactionSubCategoryRebateSend:
		return "Rebate Distribution"
	case TransactionSubCategoryRebateReceive:
		return "Rebate Receipt"
	case TransactionSubCategoryInterestProfit:
		return "Interest Earnings"
	case TransactionSubCategoryHallToInterest:
		return "Hall to Interest Transfer"
	case TransactionSubCategoryInterestManualPullback:
		return "Manual Recall - Interest"
	case TransactionSubCategoryInterestToHall:
		return "Interest to Hall Transfer"
	case TransactionSubCategoryDailyTask:
		return "Daily Task"
	case TransactionSubCategoryWeekLyTask:
		return "Weekly Task"
	case TransactionSubCategoryNewcomerProfitTask:
		return "Newcomer Benefits"
	case TransactionSubCategoryAlivenessBox:
		return "Activity Box"
	case TransactionSubCategorySecretTask:
		return "Secret Task"
	case TransactionSubCategoryVipMonthlyReward:
		return "VIP Monthly Reward"
	case TransactionSubCategoryVipDailyReward:
		return "VIP Daily Reward"
	case TransactionSubCategoryVipWeeklyReward:
		return "VIP Weekly Reward"
	case TransactionSubCategoryVipUpgradeReward:
		return "VIP Upgrade Reward"
	case TransactionSubCategoryRechargeBonus:
		return "Bonus"
	case TransactionSubCategoryManualPullbackClub:
		return "Manual Recall - Club"
	case TransactionSubCategoryHallToClub:
		return "Hall to Club Transfer"
	case TransactionSubCategoryClubToHall:
		return "Club to Hall Transfer"
	case TransactionSubCategoryGiveUpReward:
		return "Reward Waiver"
	case TransactionSubCategoryTransferOut:
		return "Transfer Out"
	case TransactionSubCategoryTransferIn:
		return "Transfer In"
	case TransactionSubCategoryClaimFreeze:
		return "Claim Freeze"
	case TransactionSubCategoryClaimDefrost:
		return "Claim Unfreeze"
	case TransactionSubCategoryClaimFee:
		return "Claim Fee"
	case TransactionSubCategoryClaimScoresIncrease:
		return "Score Increase"
	case TransactionSubCategoryClaimScoresDecrease:
		return "Score Decrease"
	case TransactionSubCategoryProvidentFundReceive:
		return "Provident Fund Receipt"
	case TransactionSubCategoryBlindBoxLotteryDeduct:
		return "Blind Box Deduction"
	case TransactionSubCategoryBlindBoxLotteryReward:
		return "Blind Box Reward"
	case TransactionSubCategoryGameBetSettlement:
		return "Bet Settlement"
	case TransactionSubCategoryRechargeOnline:
		return "Online Recharge"
	default:
		return ""
	}
}
