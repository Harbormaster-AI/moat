class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :organizations do |t|
      t.string :name      
      t.string :defaultCurrency      
      t.string :defaultLocale      
      t.string :website      
      t.timestamps
    end
    create_table :users do |t|
      t.string :username      
      t.string :fullName      
      t.string :email      
      t.string :locale      
      t.integer :Role      
      t.integer :Status      
      t.timestamps
    end
    create_table :teams do |t|
      t.string :name      
      t.integer :TeamType      
      t.timestamps
    end
    create_table :territorys do |t|
      t.string :name      
      t.string :region      
      t.integer :TerritoryType      
      t.timestamps
    end
    create_table :accounts do |t|
      t.string :name      
      t.string :accountNumber      
      t.string :industry      
      t.string :billingAddress      
      t.string :shippingAddress      
      t.string :website      
      t.string :phone      
      t.boolean :asActive      
      t.integer :AccountType      
      t.integer :LifecycleStage      
      t.timestamps
    end
    create_table :contacts do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :title      
      t.string :email      
      t.string :phone      
      t.string :mobile      
      t.string :mailingAddress      
      t.integer :PreferredContactMethod      
      t.timestamps
    end
    create_table :leads do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :company      
      t.string :email      
      t.string :phone      
      t.boolean :converted      
      t.integer :Status      
      t.integer :Source      
      t.integer :Rating      
      t.timestamps
    end
    create_table :opportunitys do |t|
      t.string :name      
      t.string :amount      
      t.date :closeDate      
      t.decimal :probability      
      t.string :description      
      t.integer :Stage      
      t.integer :Type      
      t.integer :ForecastCategory      
      t.timestamps
    end
    create_table :opportunityLineItems do |t|
      t.decimal :quantity      
      t.string :unitPrice      
      t.decimal :discountPercent      
      t.string :totalPrice      
      t.timestamps
    end
    create_table :opportunityStageHistorys do |t|
      t.datetime :changedAt      
      t.string :comment      
      t.integer :FromStage      
      t.integer :ToStage      
      t.timestamps
    end
    create_table :products do |t|
      t.string :sku      
      t.string :name      
      t.boolean :asActive      
      t.string :standardPrice      
      t.string :description      
      t.integer :ProductType      
      t.integer :Uom      
      t.timestamps
    end
    create_table :priceBooks do |t|
      t.string :name      
      t.boolean :asActive      
      t.string :description      
      t.timestamps
    end
    create_table :priceBookEntrys do |t|
      t.string :unitPrice      
      t.date :effectiveDate      
      t.date :expirationDate      
      t.boolean :asActive      
      t.timestamps
    end
    create_table :quotes do |t|
      t.string :quoteNumber      
      t.date :validityStart      
      t.date :validityEnd      
      t.string :totalAmount      
      t.decimal :discountPercent      
      t.string :taxAmount      
      t.string :shippingAmount      
      t.integer :Status      
      t.timestamps
    end
    create_table :quoteLineItems do |t|
      t.decimal :quantity      
      t.string :unitPrice      
      t.string :discountAmount      
      t.string :taxAmount      
      t.string :totalAmount      
      t.timestamps
    end
    create_table :orders do |t|
      t.string :orderNumber      
      t.date :orderDate      
      t.string :totalAmount      
      t.string :taxAmount      
      t.string :shippingAmount      
      t.integer :Status      
      t.timestamps
    end
    create_table :orderItems do |t|
      t.decimal :quantity      
      t.string :unitPrice      
      t.string :discountAmount      
      t.string :taxAmount      
      t.string :totalAmount      
      t.timestamps
    end
    create_table :contracts do |t|
      t.string :contractNumber      
      t.date :startDate      
      t.date :endDate      
      t.integer :renewalTermMonths      
      t.boolean :autoRenew      
      t.integer :Status      
      t.timestamps
    end
    create_table :case_s do |t|
      t.string :caseNumber      
      t.string :subject      
      t.string :description      
      t.datetime :slaDue      
      t.integer :Status      
      t.integer :Priority      
      t.integer :Origin      
      t.integer :Severity      
      t.timestamps
    end
    create_table :activitys do |t|
      t.string :subject      
      t.date :dueDate      
      t.datetime :startAt      
      t.datetime :endAt      
      t.string :location      
      t.integer :ActivityType      
      t.integer :Status      
      t.integer :Priority      
      t.timestamps
    end
    create_table :campaigns do |t|
      t.string :name      
      t.date :startDate      
      t.date :endDate      
      t.string :budget      
      t.string :actualCost      
      t.string :expectedRevenue      
      t.integer :Status      
      t.integer :Type      
      t.timestamps
    end
    create_table :campaignMembers do |t|
      t.boolean :responded      
      t.integer :Status      
      t.integer :MemberType      
      t.timestamps
    end
    create_table :notes do |t|
      t.string :title      
      t.string :content      
      t.datetime :createdAt      
      t.datetime :updatedAt      
      t.timestamps
    end
    create_table :emailMessages do |t|
      t.string :subject      
      t.string :body      
      t.datetime :sentAt      
      t.string :messageId      
      t.integer :Direction      
      t.integer :Status      
      t.timestamps
    end
  end
end
