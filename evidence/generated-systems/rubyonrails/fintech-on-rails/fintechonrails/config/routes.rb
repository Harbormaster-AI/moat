Rails.application.routes.draw do
  root "application#health"
  resources :financialinstitutions do
    resources :branches
    resources :customers
    resources :productofferings
    resources :paymentprocessors
    resources :compliancepolicies
  end
  resources :branchs do
    resource :institution
  end
  resources :productofferings do
    resource :institution
    resources :pricingplans
  end
  resources :pricingplans do
    resource :productoffering
    resources :feeschedules
    resources :limits
  end
  resources :feeschedules do
    resource :pricingplan
  end
  resources :usagelimits do
    resource :pricingplan
  end
  resources :customers do
    resource :institution
    resources :accounts
    resources :wallets
    resources :cards
    resources :kycprofiles
    resources :consents
    resources :agreements
    resources :loanapplications
    resources :loans
    resources :portfolios
    resources :disputes
  end
  resources :kycprofiles do
    resource :customer
    resources :documents
    resources :screenings
    resources :addresses
  end
  resources :kycdocuments do
    resource :kycprofile
  end
  resources :screenings do
    resource :kycprofile
    resources :alerts
  end
  resources :verifiedaddresss do
    resource :kycprofile
  end
  resources :compliancepolicys do
    resource :institution
  end
  resources :compliancealerts do
    resource :screening
    resource :transaction
  end
  resources :consents do
    resource :customer
    resource :apiclient
  end
  resources :apiclients do
    resources :consents
  end
  resources :agreements do
    resource :customer
    resource :productoffering
  end
  resources :accounts do
    resource :customer
    resource :institution
    resources :transactions
    resources :cards
    resources :statements
    resources :mandates
  end
  resources :wallets do
    resource :customer
    resources :transactions
  end
  resources :paymentcards do
    resource :customer
    resource :account
    resources :tokenizations
    resources :disputes
  end
  resources :cardtokenizations do
    resource :card
  end
  resources :merchants do
    resources :terminals
    resources :paymentcontracts
    resources :payouts
    resources :settlements
    resources :disputes
    resources :invoices
  end
  resources :terminals do
    resource :merchant
  end
  resources :paymentcontracts do
    resource :merchant
    resource :acquirer
  end
  resources :paymentprocessors do
    resources :institutions
    resources :contracts
    resources :settlements
  end
  resources :transactions do
    resource :account
    resource :wallet
    resource :paymentorder
    resource :merchant
    resource :card
    resources :relatedtransactions
    resources :alerts
  end
  resources :paymentorders do
    resource :sourceaccount
    resource :destinationaccount
    resource :beneficiary
    resources :transactions
    resource :fxdeal
    resources :fees
  end
  resources :beneficiarys do
    resource :customer
  end
  resources :appliedfees do
    resource :paymentorder
    resource :transaction
  end
  resources :fxquotes do
    resource :requestedby
  end
  resources :fxdeals do
    resource :quote
    resources :paymentorders
  end
  resources :settlementbatchs do
    resource :processor
    resource :merchant
    resources :payouts
    resources :transactions
  end
  resources :payouts do
    resource :merchant
    resource :settlementbatch
    resource :destinationaccount
  end
  resources :disputes do
    resource :transaction
    resource :card
    resource :merchant
    resources :chargebacks
  end
  resources :chargebacks do
    resource :dispute
    resource :transaction
  end
  resources :invoices do
    resource :merchant
    resources :payments
  end
  resources :accountstatements do
    resource :account
  end
  resources :directdebitmandates do
    resource :account
    resource :creditor
  end
  resources :creditors do
    resources :mandates
  end
  resources :loanapplications do
    resource :customer
    resource :riskassessment
    resource :loan
  end
  resources :riskassessments do
    resource :application
  end
  resources :loans do
    resource :customer
    resources :schedule
    resources :collateral
    resources :transactions
  end
  resources :repaymentschedules do
    resource :loan
    resources :payments
  end
  resources :collaterals do
    resource :loan
  end
  resources :loantransactions do
    resource :loan
  end
  resources :investmentportfolios do
    resource :customer
    resources :accounts
    resources :orders
    resources :holdings
  end
  resources :investmentaccounts do
    resource :portfolio
    resources :trades
    resources :orders
  end
  resources :securitys do
    resources :positions
    resources :trades
    resources :orders
  end
  resources :positions do
    resource :portfolio
    resource :security
  end
  resources :tradeorders do
    resource :portfolio
    resource :security
    resources :trades
  end
  resources :trades do
    resource :order
    resource :security
    resource :investmentaccount
  end
  resources :exchangerates do
    resources :usedbyquotes
  end
end
