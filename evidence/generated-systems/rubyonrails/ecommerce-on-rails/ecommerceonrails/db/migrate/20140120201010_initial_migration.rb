class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :merchants do |t|
      t.string :name      
      t.string :legalName      
      t.string :website      
      t.string :defaultCurrency      
      t.string :defaultLocale      
      t.string :supportEmail      
      t.timestamps
    end
    create_table :channels do |t|
      t.string :name      
      t.string :channelCode      
      t.string :locale      
      t.string :domain      
      t.boolean :asActive      
      t.string :defaultCurrency      
      t.integer :ChannelType      
      t.timestamps
    end
    create_table :brands do |t|
      t.string :name      
      t.string :description      
      t.string :website      
      t.timestamps
    end
    create_table :catalogs do |t|
      t.string :name      
      t.string :catalogCode      
      t.boolean :asActive      
      t.timestamps
    end
    create_table :categorys do |t|
      t.string :name      
      t.string :slug      
      t.integer :position      
      t.boolean :asActive      
      t.timestamps
    end
    create_table :products do |t|
      t.string :name      
      t.string :slug      
      t.boolean :asActive      
      t.integer :ProductType      
      t.integer :DefaultTaxClass      
      t.timestamps
    end
    create_table :productVariants do |t|
      t.string :sku      
      t.string :barcode      
      t.string :title      
      t.decimal :weight      
      t.boolean :requiresShipping      
      t.integer :WeightUnit      
      t.timestamps
    end
    create_table :productPricings do |t|
      t.string :listPrice      
      t.string :salePrice      
      t.date :validFrom      
      t.date :validTo      
      t.timestamps
    end
    create_table :mediaAssets do |t|
      t.string :url      
      t.string :altText      
      t.integer :position      
      t.integer :MediaType      
      t.timestamps
    end
    create_table :fulfillmentCenters do |t|
      t.string :name      
      t.string :centerCode      
      t.string :address      
      t.string :timezone      
      t.boolean :asActive      
      t.timestamps
    end
    create_table :inventoryItems do |t|
      t.integer :quantityOnHand      
      t.integer :quantityReserved      
      t.integer :safetyStock      
      t.integer :Status      
      t.timestamps
    end
    create_table :suppliers do |t|
      t.string :name      
      t.string :contactEmail      
      t.string :website      
      t.integer :Status      
      t.timestamps
    end
    create_table :sellers do |t|
      t.string :name      
      t.string :sellerCode      
      t.string :contactEmail      
      t.integer :Status      
      t.timestamps
    end
    create_table :customers do |t|
      t.string :firstName      
      t.string :lastName      
      t.string :email      
      t.string :phone      
      t.boolean :marketingOptIn      
      t.integer :CustomerGroup      
      t.timestamps
    end
    create_table :customerAddresss do |t|
      t.string :label      
      t.string :address      
      t.boolean :asDefaultShipping      
      t.boolean :asDefaultBilling      
      t.timestamps
    end
    create_table :wishlists do |t|
      t.string :name      
      t.boolean :asPublic      
      t.date :createdAt      
      t.timestamps
    end
    create_table :wishlistItems do |t|
      t.date :addedDate      
      t.timestamps
    end
    create_table :carts do |t|
      t.string :cartNumber      
      t.date :createdAt      
      t.string :currency      
      t.string :shippingAddress      
      t.string :billingAddress      
      t.integer :Status      
      t.timestamps
    end
    create_table :cartItems do |t|
      t.integer :quantity      
      t.string :unitPrice      
      t.string :totalPrice      
      t.timestamps
    end
    create_table :orders do |t|
      t.string :orderNumber      
      t.date :placedDate      
      t.string :subtotal      
      t.string :discountTotal      
      t.string :shippingTotal      
      t.string :taxTotal      
      t.string :grandTotal      
      t.string :shippingAddress      
      t.string :billingAddress      
      t.integer :Status      
      t.timestamps
    end
    create_table :orderLines do |t|
      t.integer :quantity      
      t.string :unitPrice      
      t.string :totalPrice      
      t.string :taxRate      
      t.integer :LineStatus      
      t.timestamps
    end
    create_table :payments do |t|
      t.string :paymentNumber      
      t.string :amount      
      t.string :transactionId      
      t.date :authorizedAt      
      t.date :capturedAt      
      t.integer :Status      
      t.integer :PaymentMethod      
      t.timestamps
    end
    create_table :refunds do |t|
      t.string :refundNumber      
      t.string :amount      
      t.string :reason      
      t.date :createdAt      
      t.integer :Status      
      t.timestamps
    end
    create_table :shipments do |t|
      t.string :shipmentNumber      
      t.date :shippedDate      
      t.date :deliveredDate      
      t.string :trackingNumber      
      t.string :shippingAddress      
      t.integer :Status      
      t.integer :Carrier      
      t.timestamps
    end
    create_table :shipmentItems do |t|
      t.integer :quantity      
      t.timestamps
    end
    create_table :returnRequests do |t|
      t.string :returnNumber      
      t.date :createdAt      
      t.string :refundAmount      
      t.integer :Status      
      t.timestamps
    end
    create_table :returnItems do |t|
      t.integer :quantity      
      t.integer :Reason      
      t.integer :Condition      
      t.timestamps
    end
    create_table :promotions do |t|
      t.string :name      
      t.string :code      
      t.decimal :value      
      t.date :startDate      
      t.date :endDate      
      t.boolean :asStackable      
      t.integer :maxRedemptions      
      t.integer :PromotionType      
      t.integer :DiscountType      
      t.timestamps
    end
    create_table :coupons do |t|
      t.string :code      
      t.integer :usageLimit      
      t.integer :perCustomerLimit      
      t.date :expirationDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :couponRedemptions do |t|
      t.date :redeemedAt      
      t.timestamps
    end
    create_table :taxRules do |t|
      t.string :name      
      t.string :country      
      t.string :region      
      t.string :rate      
      t.boolean :taxInclusive      
      t.integer :TaxClass      
      t.timestamps
    end
    create_table :shippingMethods do |t|
      t.string :name      
      t.string :flatRate      
      t.integer :estimatedDays      
      t.boolean :asActive      
      t.integer :MethodType      
      t.timestamps
    end
    create_table :carrierServices do |t|
      t.string :name      
      t.string :code      
      t.integer :Carrier      
      t.integer :ServiceLevel      
      t.timestamps
    end
    create_table :reviews do |t|
      t.integer :rating      
      t.string :title      
      t.string :content      
      t.date :createdAt      
      t.integer :Status      
      t.timestamps
    end
    create_table :subscriptions do |t|
      t.string :subscriptionNumber      
      t.date :nextBillingDate      
      t.date :startDate      
      t.date :endDate      
      t.integer :Status      
      t.integer :Interval      
      t.timestamps
    end
    create_table :paymentProviders do |t|
      t.string :name      
      t.boolean :enabled      
      t.string :merchantAccountId      
      t.integer :ProviderType      
      t.timestamps
    end
    create_table :invoices do |t|
      t.string :invoiceNumber      
      t.date :issuedDate      
      t.date :dueDate      
      t.string :total      
      t.integer :Status      
      t.timestamps
    end
    create_table :giftCards do |t|
      t.string :code      
      t.string :balance      
      t.date :expirationDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :giftCardRedemptions do |t|
      t.date :redeemedAt      
      t.string :amount      
      t.timestamps
    end
    create_table :payouts do |t|
      t.string :payoutNumber      
      t.string :amount      
      t.date :scheduledDate      
      t.date :paidDate      
      t.integer :Status      
      t.timestamps
    end
  end
end
