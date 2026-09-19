class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :stockKeepingUnits do |t|
      t.string :skuCode      
      t.string :name      
      t.decimal :weight      
      t.string :weightUnit      
      t.decimal :volume      
      t.string :volumeUnit      
      t.integer :shelfLifeDays      
      t.boolean :hazardousMaterial      
      t.integer :ItemType      
      t.integer :UnitOfMeasure      
      t.timestamps
    end
    create_table :warehouses do |t|
      t.string :name      
      t.string :code      
      t.string :address      
      t.string :timeZone      
      t.boolean :allowsOverAllocation      
      t.timestamps
    end
    create_table :storageLocations do |t|
      t.string :code      
      t.boolean :temperatureControlled      
      t.decimal :capacity      
      t.string :capacityUnit      
      t.integer :LocationType      
      t.timestamps
    end
    create_table :inventoryItems do |t|
      t.decimal :quantityOnHand      
      t.decimal :quantityAvailable      
      t.decimal :quantityReserved      
      t.string :unitCost      
      t.date :lastUpdated      
      t.integer :StockStatus      
      t.timestamps
    end
    create_table :lots do |t|
      t.string :batchNumber      
      t.date :manufactureDate      
      t.date :expirationDate      
      t.integer :LotStatus      
      t.timestamps
    end
    create_table :serialNumbers do |t|
      t.string :serial      
      t.date :activationDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :reservations do |t|
      t.string :referenceNumber      
      t.decimal :reservedQuantity      
      t.date :promisedDate      
      t.integer :ReservationStatus      
      t.integer :ReservationType      
      t.timestamps
    end
    create_table :demandSignals do |t|
      t.string :externalReference      
      t.date :requestedDate      
      t.decimal :quantity      
      t.integer :DemandType      
      t.timestamps
    end
    create_table :inventoryTransactions do |t|
      t.string :transactionNumber      
      t.decimal :quantity      
      t.string :unitCost      
      t.date :transactionDate      
      t.string :reasonCode      
      t.integer :TransactionType      
      t.integer :UnitOfMeasure      
      t.integer :Status      
      t.timestamps
    end
    create_table :transferOrders do |t|
      t.string :orderNumber      
      t.date :requestedShipDate      
      t.date :requestedReceiveDate      
      t.date :shippedDate      
      t.date :receivedDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :transferOrderLines do |t|
      t.integer :lineNumber      
      t.decimal :quantity      
      t.integer :UnitOfMeasure      
      t.integer :StockStatus      
      t.timestamps
    end
    create_table :stockAdjustments do |t|
      t.string :adjustmentNumber      
      t.string :reason      
      t.date :adjustmentDate      
      t.integer :AdjustmentType      
      t.integer :Status      
      t.timestamps
    end
    create_table :stockAdjustmentLines do |t|
      t.integer :lineNumber      
      t.decimal :quantity      
      t.integer :UnitOfMeasure      
      t.integer :StockStatus      
      t.timestamps
    end
    create_table :cycleCounts do |t|
      t.string :countNumber      
      t.date :scheduledDate      
      t.date :performedDate      
      t.string :approvedBy      
      t.integer :Status      
      t.timestamps
    end
    create_table :cycleCountEntrys do |t|
      t.integer :lineNumber      
      t.decimal :systemQuantity      
      t.decimal :countedQuantity      
      t.decimal :varianceQuantity      
      t.boolean :recountRequired      
      t.integer :StockStatus      
      t.timestamps
    end
    create_table :replenishmentPolicys do |t|
      t.decimal :minLevel      
      t.decimal :maxLevel      
      t.decimal :reorderPoint      
      t.decimal :reorderQuantity      
      t.integer :leadTimeDays      
      t.integer :reviewPeriodDays      
      t.integer :PolicyType      
      t.timestamps
    end
    create_table :uoMConversions do |t|
      t.decimal :factor      
      t.integer :precision      
      t.integer :FromUnit      
      t.integer :ToUnit      
      t.timestamps
    end
    create_table :inventoryThresholdAlerts do |t|
      t.string :alertNumber      
      t.date :detectedAt      
      t.string :message      
      t.integer :AlertType      
      t.integer :Status      
      t.timestamps
    end
    create_table :quarantines do |t|
      t.string :reason      
      t.date :startedAt      
      t.date :releasedAt      
      t.integer :Disposition      
      t.timestamps
    end
    create_table :expirationPolicys do |t|
      t.integer :rejectIfDaysToExpireLessThan      
      t.integer :autoQuarantineDaysToExpire      
      t.integer :RotationMethod      
      t.timestamps
    end
    create_table :inboundShipments do |t|
      t.string :shipmentNumber      
      t.date :expectedArrivalDate      
      t.date :arrivalDate      
      t.string :carrierName      
      t.integer :Status      
      t.timestamps
    end
    create_table :inboundShipmentLines do |t|
      t.integer :lineNumber      
      t.decimal :quantity      
      t.integer :UnitOfMeasure      
      t.integer :StockStatus      
      t.timestamps
    end
    create_table :outboundAllocations do |t|
      t.string :allocationNumber      
      t.decimal :allocatedQuantity      
      t.date :allocationDate      
      t.integer :Status      
      t.timestamps
    end
  end
end
