class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :insurers do |t|
      t.string :name      
      t.string :legalName      
      t.string :domicileCountry      
      t.string :naicNumber      
      t.string :website      
      t.timestamps
    end
    create_table :insuranceProducts do |t|
      t.string :name      
      t.string :productCode      
      t.integer :LineOfBusiness      
      t.timestamps
    end
    create_table :coverageDefinitions do |t|
      t.string :name      
      t.string :defaultLimit      
      t.string :defaultDeductible      
      t.boolean :asMandatory      
      t.integer :CoverageType      
      t.timestamps
    end
    create_table :distributors do |t|
      t.string :name      
      t.string :licenseNumber      
      t.string :region      
      t.integer :DistributorType      
      t.timestamps
    end
    create_table :agents do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :licenseId      
      t.integer :Status      
      t.timestamps
    end
    create_table :customers do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :organizationName      
      t.string :taxId      
      t.date :dateOfBirth      
      t.string :primaryAddress      
      t.integer :CustomerType      
      t.timestamps
    end
    create_table :applications do |t|
      t.string :applicationNumber      
      t.date :submissionDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :quotes do |t|
      t.string :quoteNumber      
      t.string :totalPremium      
      t.date :ratingDate      
      t.boolean :asBound      
      t.timestamps
    end
    create_table :underwritingDecisions do |t|
      t.string :notes      
      t.date :decisionDate      
      t.integer :Decision      
      t.timestamps
    end
    create_table :underwriters do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :employeeId      
      t.string :authorityLimit      
      t.timestamps
    end
    create_table :policys do |t|
      t.string :policyNumber      
      t.string :effectivePeriod      
      t.string :totalPremium      
      t.integer :Status      
      t.integer :PaymentPlan      
      t.timestamps
    end
    create_table :endorsements do |t|
      t.string :endorsementNumber      
      t.date :effectiveDate      
      t.string :description      
      t.timestamps
    end
    create_table :policyCoverages do |t|
      t.string :limit      
      t.string :deductible      
      t.string :premium      
      t.integer :CoverageType      
      t.timestamps
    end
    create_table :insuredObjects do |t|
      t.string :description      
      t.string :serialOrId      
      t.string :primaryAddress      
      t.integer :ObjectType      
      t.timestamps
    end
    create_table :beneficiarys do |t|
      t.string :name      
      t.string :share      
      t.integer :Relationship      
      t.timestamps
    end
    create_table :billingAccounts do |t|
      t.string :accountNumber      
      t.string :balance      
      t.integer :Status      
      t.timestamps
    end
    create_table :invoices do |t|
      t.string :invoiceNumber      
      t.date :dueDate      
      t.string :totalDue      
      t.integer :Status      
      t.timestamps
    end
    create_table :payments do |t|
      t.string :paymentReference      
      t.string :amount      
      t.date :paymentDate      
      t.integer :Method      
      t.integer :Status      
      t.timestamps
    end
    create_table :claims do |t|
      t.string :claimNumber      
      t.date :noticeDate      
      t.date :lossDate      
      t.string :reportedBy      
      t.integer :Status      
      t.integer :LossCause      
      t.timestamps
    end
    create_table :incidents do |t|
      t.string :location      
      t.string :description      
      t.integer :IncidentType      
      t.timestamps
    end
    create_table :exposures do |t|
      t.integer :ExposureType      
      t.integer :Status      
      t.timestamps
    end
    create_table :adjusters do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :licenseNumber      
      t.integer :AdjusterType      
      t.timestamps
    end
    create_table :claimReserves do |t|
      t.string :amount      
      t.date :setDate      
      t.integer :ReserveType      
      t.integer :Status      
      t.timestamps
    end
    create_table :claimPayments do |t|
      t.string :paymentNumber      
      t.string :amount      
      t.date :paymentDate      
      t.integer :PayeeType      
      t.integer :Method      
      t.integer :Status      
      t.timestamps
    end
    create_table :serviceProviders do |t|
      t.string :name      
      t.string :taxId      
      t.integer :ProviderType      
      t.integer :NetworkStatus      
      t.timestamps
    end
    create_table :reinsuranceAgreements do |t|
      t.string :agreementNumber      
      t.string :effectivePeriod      
      t.string :retention      
      t.string :limit      
      t.string :cessionPercentage      
      t.integer :ReinsuranceType      
      t.integer :TreatyType      
      t.timestamps
    end
    create_table :subrogationRecoverys do |t|
      t.string :recoveryReference      
      t.string :amount      
      t.date :recoveryDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :thirdPartys do |t|
      t.string :name      
      t.string :taxId      
      t.string :address      
      t.integer :PartyType      
      t.timestamps
    end
    create_table :documents do |t|
      t.string :fileName      
      t.date :uploadedDate      
      t.integer :DocumentType      
      t.timestamps
    end
  end
end
