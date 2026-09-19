class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :financialInstitutions do |t|
      t.string :name      
      t.string :legalName      
      t.string :countryOfIncorporation      
      t.string :bic      
      t.string :website      
      t.timestamps
    end
    create_table :branchs do |t|
      t.string :name      
      t.string :branchCode      
      t.string :address      
      t.timestamps
    end
    create_table :productOfferings do |t|
      t.string :name      
      t.string :productCode      
      t.integer :Category      
      t.timestamps
    end
    create_table :pricingPlans do |t|
      t.string :name      
      t.string :planCode      
      t.string :baseCurrency      
      t.integer :Status      
      t.timestamps
    end
    create_table :feeSchedules do |t|
      t.string :name      
      t.string :amount      
      t.decimal :percentage      
      t.string :minimum      
      t.string :maximum      
      t.integer :FeeType      
      t.integer :CalculationMethod      
      t.timestamps
    end
    create_table :usageLimits do |t|
      t.string :name      
      t.string :amount      
      t.integer :count      
      t.integer :Scope      
      t.integer :Period      
      t.timestamps
    end
    create_table :customers do |t|
      t.string :firstName      
      t.string :lastName      
      t.date :dateOfBirth      
      t.string :email      
      t.string :phone      
      t.string :address      
      t.string :taxId      
      t.string :riskScore      
      t.integer :CustomerType      
      t.timestamps
    end
    create_table :kYCProfiles do |t|
      t.string :profileId      
      t.datetime :createdAt      
      t.integer :Status      
      t.integer :VerificationLevel      
      t.timestamps
    end
    create_table :kYCDocuments do |t|
      t.string :reference      
      t.string :issuedCountry      
      t.date :expirationDate      
      t.integer :DocumentType      
      t.integer :Status      
      t.timestamps
    end
    create_table :screenings do |t|
      t.string :score      
      t.datetime :screenedAt      
      t.integer :ScreeningType      
      t.integer :Status      
      t.timestamps
    end
    create_table :verifiedAddresss do |t|
      t.string :address      
      t.datetime :verifiedAt      
      t.integer :VerificationStatus      
      t.timestamps
    end
    create_table :compliancePolicys do |t|
      t.string :name      
      t.string :policyCode      
      t.string :description      
      t.integer :Status      
      t.timestamps
    end
    create_table :complianceAlerts do |t|
      t.string :alertCode      
      t.datetime :raisedAt      
      t.string :notes      
      t.integer :Severity      
      t.integer :Status      
      t.timestamps
    end
    create_table :consents do |t|
      t.datetime :grantedAt      
      t.datetime :expiresAt      
      t.string :scope      
      t.integer :ConsentType      
      t.integer :Status      
      t.timestamps
    end
    create_table :aPIClients do |t|
      t.string :name      
      t.string :clientId      
      t.string :redirectUri      
      t.integer :ClientType      
      t.timestamps
    end
    create_table :agreements do |t|
      t.string :agreementNumber      
      t.date :effectiveDate      
      t.integer :AgreementType      
      t.integer :Status      
      t.timestamps
    end
    create_table :accounts do |t|
      t.string :accountNumber      
      t.string :iban      
      t.string :bic      
      t.date :openedDate      
      t.string :currency      
      t.string :balance      
      t.string :availableBalance      
      t.integer :AccountType      
      t.integer :Status      
      t.timestamps
    end
    create_table :wallets do |t|
      t.string :currency      
      t.string :balance      
      t.integer :Status      
      t.timestamps
    end
    create_table :paymentCards do |t|
      t.string :cardToken      
      t.string :maskedPan      
      t.integer :expiryMonth      
      t.integer :expiryYear      
      t.string :cardholderName      
      t.integer :Scheme      
      t.integer :Status      
      t.timestamps
    end
    create_table :cardTokenizations do |t|
      t.string :tokenReference      
      t.datetime :createdAt      
      t.integer :WalletProvider      
      t.integer :Status      
      t.timestamps
    end
    create_table :merchants do |t|
      t.string :name      
      t.string :mcc      
      t.string :url      
      t.string :country      
      t.string :settlementCurrency      
      t.timestamps
    end
    create_table :terminals do |t|
      t.string :location      
      t.integer :Type      
      t.integer :Status      
      t.timestamps
    end
    create_table :paymentContracts do |t|
      t.string :contractNumber      
      t.string :pricingPlanCode      
      t.integer :Status      
      t.timestamps
    end
    create_table :paymentProcessors do |t|
      t.string :name      
      t.string :processorCode      
      t.string :networkSupport      
      t.timestamps
    end
    create_table :transactions do |t|
      t.string :amount      
      t.string :fee      
      t.decimal :exchangeRate      
      t.datetime :createdAt      
      t.datetime :completedAt      
      t.string :narrative      
      t.integer :TransactionType      
      t.integer :Status      
      t.timestamps
    end
    create_table :paymentOrders do |t|
      t.string :orderReference      
      t.date :requestedExecutionDate      
      t.string :purpose      
      t.integer :PaymentMethod      
      t.integer :Status      
      t.integer :Priority      
      t.timestamps
    end
    create_table :beneficiarys do |t|
      t.string :name      
      t.string :accountIdentifier      
      t.string :iban      
      t.string :bic      
      t.string :address      
      t.timestamps
    end
    create_table :appliedFees do |t|
      t.string :amount      
      t.string :description      
      t.integer :FeeType      
      t.timestamps
    end
    create_table :fXQuotes do |t|
      t.string :baseCurrency      
      t.string :quoteCurrency      
      t.decimal :rate      
      t.datetime :quotedAt      
      t.datetime :expiresAt      
      t.integer :PriceType      
      t.timestamps
    end
    create_table :fXDeals do |t|
      t.string :dealReference      
      t.string :baseCurrency      
      t.string :quoteCurrency      
      t.decimal :rate      
      t.string :amount      
      t.date :settlementDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :settlementBatchs do |t|
      t.string :batchId      
      t.datetime :periodStart      
      t.datetime :periodEnd      
      t.string :totalVolume      
      t.integer :totalCount      
      t.integer :Status      
      t.timestamps
    end
    create_table :payouts do |t|
      t.string :payoutReference      
      t.string :amount      
      t.string :currency      
      t.date :scheduledDate      
      t.date :paidDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :disputes do |t|
      t.string :disputeReference      
      t.datetime :openedAt      
      t.datetime :closedAt      
      t.integer :Reason      
      t.integer :Status      
      t.timestamps
    end
    create_table :chargebacks do |t|
      t.string :chargebackReference      
      t.string :amount      
      t.datetime :postedAt      
      t.integer :Stage      
      t.integer :Status      
      t.timestamps
    end
    create_table :invoices do |t|
      t.string :invoiceNumber      
      t.date :issueDate      
      t.date :dueDate      
      t.string :total      
      t.string :currency      
      t.integer :Status      
      t.timestamps
    end
    create_table :accountStatements do |t|
      t.string :statementNumber      
      t.date :periodStart      
      t.date :periodEnd      
      t.string :openingBalance      
      t.string :closingBalance      
      t.datetime :generatedAt      
      t.timestamps
    end
    create_table :directDebitMandates do |t|
      t.string :mandateId      
      t.datetime :signedAt      
      t.integer :Scheme      
      t.integer :Status      
      t.timestamps
    end
    create_table :creditors do |t|
      t.string :name      
      t.string :bic      
      t.string :address      
      t.timestamps
    end
    create_table :loanApplications do |t|
      t.string :applicationNumber      
      t.string :amountRequested      
      t.integer :termMonths      
      t.datetime :submittedAt      
      t.integer :Product      
      t.integer :Purpose      
      t.integer :Status      
      t.timestamps
    end
    create_table :riskAssessments do |t|
      t.string :score      
      t.datetime :assessedAt      
      t.string :modelVersion      
      t.string :notes      
      t.integer :Decision      
      t.timestamps
    end
    create_table :loans do |t|
      t.string :loanNumber      
      t.string :principal      
      t.decimal :interestRate      
      t.date :originationDate      
      t.date :maturityDate      
      t.integer :RateType      
      t.integer :Status      
      t.timestamps
    end
    create_table :repaymentSchedules do |t|
      t.integer :installmentNumber      
      t.date :dueDate      
      t.string :amountDue      
      t.string :principalDue      
      t.string :interestDue      
      t.integer :Status      
      t.timestamps
    end
    create_table :collaterals do |t|
      t.string :description      
      t.string :value      
      t.integer :CollateralType      
      t.timestamps
    end
    create_table :loanTransactions do |t|
      t.string :transactionId      
      t.string :amount      
      t.date :postingDate      
      t.integer :Type      
      t.integer :Status      
      t.timestamps
    end
    create_table :investmentPortfolios do |t|
      t.string :portfolioCode      
      t.string :baseCurrency      
      t.datetime :createdAt      
      t.integer :Status      
      t.timestamps
    end
    create_table :investmentAccounts do |t|
      t.string :accountNumber      
      t.string :baseCurrency      
      t.string :balance      
      t.integer :AccountType      
      t.integer :Status      
      t.timestamps
    end
    create_table :securitys do |t|
      t.string :symbol      
      t.string :isin      
      t.string :cusip      
      t.string :currency      
      t.integer :SecurityType      
      t.timestamps
    end
    create_table :positions do |t|
      t.decimal :quantity      
      t.string :averageCost      
      t.string :marketValue      
      t.timestamps
    end
    create_table :tradeOrders do |t|
      t.string :orderId      
      t.decimal :quantity      
      t.string :limitPrice      
      t.datetime :placedAt      
      t.integer :Side      
      t.integer :Type      
      t.integer :Status      
      t.integer :TimeInForce      
      t.timestamps
    end
    create_table :trades do |t|
      t.datetime :executedAt      
      t.decimal :quantity      
      t.string :price      
      t.string :fees      
      t.date :settlementDate      
      t.timestamps
    end
    create_table :exchangeRates do |t|
      t.string :baseCurrency      
      t.string :quoteCurrency      
      t.decimal :rate      
      t.datetime :asOf      
      t.string :source      
      t.timestamps
    end
  end
end
