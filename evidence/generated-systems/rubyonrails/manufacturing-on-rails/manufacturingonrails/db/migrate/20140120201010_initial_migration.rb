class InitialMigration < ActiveRecord::Migration[6.1]
  def change
    create_table :enterprises do |t|
      t.string :name      
      t.string :legalName      
      t.string :registrationCountry      
      t.string :website      
      t.string :taxId      
      t.timestamps
    end
    create_table :businessUnits do |t|
      t.string :name      
      t.string :code      
      t.integer :Category      
      t.timestamps
    end
    create_table :plants do |t|
      t.string :name      
      t.string :plantCode      
      t.string :address      
      t.string :timeZone      
      t.timestamps
    end
    create_table :productionLines do |t|
      t.string :name      
      t.string :lineCode      
      t.integer :LineType      
      t.timestamps
    end
    create_table :workCenters do |t|
      t.string :name      
      t.string :code      
      t.integer :capacityPerHour      
      t.string :oeeTarget      
      t.integer :WorkCenterType      
      t.timestamps
    end
    create_table :items do |t|
      t.string :itemNumber      
      t.string :name      
      t.string :standardCost      
      t.string :weight      
      t.boolean :asSerialControlled      
      t.integer :ItemType      
      t.integer :ProcurementType      
      t.integer :UnitOfMeasure      
      t.integer :LifecycleStatus      
      t.timestamps
    end
    create_table :bOMs do |t|
      t.string :bomNumber      
      t.string :revision      
      t.date :effectivityStart      
      t.date :effectivityEnd      
      t.integer :Status      
      t.timestamps
    end
    create_table :bOMItems do |t|
      t.integer :lineNumber      
      t.string :quantity      
      t.string :scrapPercent      
      t.timestamps
    end
    create_table :routings do |t|
      t.string :routingNumber      
      t.string :revision      
      t.date :effectivityStart      
      t.date :effectivityEnd      
      t.integer :RoutingType      
      t.integer :Status      
      t.timestamps
    end
    create_table :operations do |t|
      t.string :operationNumber      
      t.string :name      
      t.string :setupTime      
      t.string :standardCycleTime      
      t.integer :OperationType      
      t.timestamps
    end
    create_table :workOrders do |t|
      t.string :workOrderNumber      
      t.datetime :plannedStart      
      t.datetime :plannedEnd      
      t.string :quantity      
      t.integer :priority      
      t.integer :Status      
      t.timestamps
    end
    create_table :productionSchedules do |t|
      t.string :scheduleNumber      
      t.date :horizonStart      
      t.date :horizonEnd      
      t.integer :Status      
      t.timestamps
    end
    create_table :suppliers do |t|
      t.string :name      
      t.string :supplierCode      
      t.string :address      
      t.integer :SupplierTier      
      t.integer :PaymentTerms      
      t.timestamps
    end
    create_table :purchaseOrders do |t|
      t.string :poNumber      
      t.date :orderDate      
      t.string :totalAmount      
      t.integer :Status      
      t.timestamps
    end
    create_table :purchaseOrderLines do |t|
      t.integer :lineNumber      
      t.string :quantity      
      t.string :unitPrice      
      t.date :dueDate      
      t.timestamps
    end
    create_table :goodsReceipts do |t|
      t.string :receiptNumber      
      t.date :receiptDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :goodsReceiptLines do |t|
      t.integer :lineNumber      
      t.string :receivedQuantity      
      t.string :acceptedQuantity      
      t.string :rejectedQuantity      
      t.string :lot      
      t.timestamps
    end
    create_table :warehouses do |t|
      t.string :name      
      t.string :warehouseCode      
      t.string :address      
      t.integer :WarehouseType      
      t.timestamps
    end
    create_table :locations do |t|
      t.string :locationCode      
      t.string :description      
      t.integer :LocationType      
      t.timestamps
    end
    create_table :inventoryItems do |t|
      t.string :quantityOnHand      
      t.string :quantityReserved      
      t.string :lotNumber      
      t.string :serialNumber      
      t.timestamps
    end
    create_table :inventoryTransactions do |t|
      t.string :transactionNumber      
      t.string :quantity      
      t.datetime :transactionDateTime      
      t.string :referenceDocument      
      t.integer :TransactionType      
      t.timestamps
    end
    create_table :customers do |t|
      t.string :name      
      t.string :customerCode      
      t.string :address      
      t.integer :CustomerType      
      t.timestamps
    end
    create_table :salesOrders do |t|
      t.string :orderNumber      
      t.date :orderDate      
      t.string :totalAmount      
      t.integer :Status      
      t.timestamps
    end
    create_table :salesOrderLines do |t|
      t.integer :lineNumber      
      t.string :quantity      
      t.string :unitPrice      
      t.date :dueDate      
      t.timestamps
    end
    create_table :qualitySpecifications do |t|
      t.string :specCode      
      t.string :name      
      t.string :version      
      t.timestamps
    end
    create_table :inspectionPlans do |t|
      t.string :planNumber      
      t.string :revision      
      t.integer :SamplingPlan      
      t.integer :Status      
      t.timestamps
    end
    create_table :inspectionCharacteristics do |t|
      t.string :characteristicCode      
      t.string :name      
      t.string :lowerSpecLimit      
      t.string :upperSpecLimit      
      t.string :target      
      t.integer :MeasurementType      
      t.timestamps
    end
    create_table :inspectionLots do |t|
      t.string :lotNumber      
      t.string :quantity      
      t.integer :sampleSize      
      t.datetime :createdOn      
      t.integer :InspectionType      
      t.integer :Status      
      t.timestamps
    end
    create_table :inspectionResults do |t|
      t.string :resultValue      
      t.datetime :recordedOn      
      t.string :notes      
      t.integer :ResultStatus      
      t.timestamps
    end
    create_table :nonconformances do |t|
      t.string :ncNumber      
      t.string :description      
      t.string :containmentAction      
      t.integer :NcType      
      t.integer :Severity      
      t.integer :Status      
      t.timestamps
    end
    create_table :correctiveActions do |t|
      t.string :capaNumber      
      t.string :rootCause      
      t.string :correctiveAction      
      t.date :verificationDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :assets do |t|
      t.string :assetTag      
      t.string :assetName      
      t.date :commissioningDate      
      t.integer :AssetStatus      
      t.timestamps
    end
    create_table :maintenancePlans do |t|
      t.string :planNumber      
      t.string :interval      
      t.date :lastServiceDate      
      t.integer :Strategy      
      t.timestamps
    end
    create_table :maintenanceOrders do |t|
      t.string :orderNumber      
      t.integer :priority      
      t.date :requestedDate      
      t.date :completionDate      
      t.integer :Status      
      t.timestamps
    end
    create_table :employees do |t|
      t.string :firstName      
      t.string :lastName      
      t.integer :Role      
      t.integer :SkillLevel      
      t.timestamps
    end
    create_table :shifts do |t|
      t.string :shiftName      
      t.string :startTime      
      t.string :endTime      
      t.integer :ShiftType      
      t.timestamps
    end
    create_table :shiftAssignments do |t|
      t.date :assignmentDate      
      t.timestamps
    end
    create_table :forecasts do |t|
      t.string :forecastNumber      
      t.date :forecastHorizonStart      
      t.date :forecastHorizonEnd      
      t.integer :Method      
      t.timestamps
    end
    create_table :forecastLines do |t|
      t.date :period      
      t.string :quantity      
      t.string :confidence      
      t.timestamps
    end
    create_table :mRPRuns do |t|
      t.string :runNumber      
      t.datetime :runDateTime      
      t.integer :planningHorizonDays      
      t.integer :Status      
      t.timestamps
    end
    create_table :plannedOrders do |t|
      t.string :plannedOrderNumber      
      t.string :quantity      
      t.date :dueDate      
      t.integer :OrderType      
      t.integer :Status      
      t.timestamps
    end
  end
end
