# This file should contain all the record creation needed to seed the database with its default values.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).
#
# Examples:
#
#   movies = Movie.create([{ name: 'Star Wars' }, { name: 'Lord of the Rings' }])
#   Character.create(name: 'Luke', movie: movies.first)


5.times do |i|
  StockKeepingUnit.create( skuCode:"test value", name:"test string for name", weight:"test value", weightUnit:"test string for weightUnit", volume:"test value", volumeUnit:"test string for volumeUnit", shelfLifeDays:100, hazardousMaterial:true, ItemType:0, UnitOfMeasure:0 )
  Warehouse.create( name:"test string for name", code:"test string for code", address:"test value", timeZone:"test string for timeZone", allowsOverAllocation:true )
  StorageLocation.create( code:"test string for code", temperatureControlled:true, capacity:"test value", capacityUnit:"test string for capacityUnit", LocationType:0 )
  InventoryItem.create( quantityOnHand:"test value", quantityAvailable:"test value", quantityReserved:"test value", unitCost:"test value", lastUpdated:1.week.ago, StockStatus:0 )
  Lot.create( batchNumber:"test value", manufactureDate:1.week.ago, expirationDate:1.week.ago, LotStatus:0 )
  SerialNumber.create( serial:"test value", activationDate:1.week.ago, Status:0 )
  Reservation.create( referenceNumber:"test string for referenceNumber", reservedQuantity:"test value", promisedDate:1.week.ago, ReservationStatus:0, ReservationType:0 )
  DemandSignal.create( externalReference:"test string for externalReference", requestedDate:1.week.ago, quantity:"test value", DemandType:0 )
  InventoryTransaction.create( transactionNumber:"test string for transactionNumber", quantity:"test value", unitCost:"test value", transactionDate:1.week.ago, reasonCode:"test string for reasonCode", TransactionType:0, UnitOfMeasure:0, Status:0 )
  TransferOrder.create( orderNumber:"test string for orderNumber", requestedShipDate:1.week.ago, requestedReceiveDate:1.week.ago, shippedDate:1.week.ago, receivedDate:1.week.ago, Status:0 )
  TransferOrderLine.create( lineNumber:100, quantity:"test value", UnitOfMeasure:0, StockStatus:0 )
  StockAdjustment.create( adjustmentNumber:"test string for adjustmentNumber", reason:"test string for reason", adjustmentDate:1.week.ago, AdjustmentType:0, Status:0 )
  StockAdjustmentLine.create( lineNumber:100, quantity:"test value", UnitOfMeasure:0, StockStatus:0 )
  CycleCount.create( countNumber:"test string for countNumber", scheduledDate:1.week.ago, performedDate:1.week.ago, approvedBy:"test string for approvedBy", Status:0 )
  CycleCountEntry.create( lineNumber:100, systemQuantity:"test value", countedQuantity:"test value", varianceQuantity:"test value", recountRequired:true, StockStatus:0 )
  ReplenishmentPolicy.create( minLevel:"test value", maxLevel:"test value", reorderPoint:"test value", reorderQuantity:"test value", leadTimeDays:100, reviewPeriodDays:100, PolicyType:0 )
  UoMConversion.create( factor:"test value", precision:100, FromUnit:0, ToUnit:0 )
  InventoryThresholdAlert.create( alertNumber:"test string for alertNumber", detectedAt:1.week.ago, message:"test string for message", AlertType:0, Status:0 )
  Quarantine.create( reason:"test string for reason", startedAt:1.week.ago, releasedAt:1.week.ago, Disposition:0 )
  ExpirationPolicy.create( rejectIfDaysToExpireLessThan:100, autoQuarantineDaysToExpire:100, RotationMethod:0 )
  InboundShipment.create( shipmentNumber:"test string for shipmentNumber", expectedArrivalDate:1.week.ago, arrivalDate:1.week.ago, carrierName:"test string for carrierName", Status:0 )
  InboundShipmentLine.create( lineNumber:100, quantity:"test value", UnitOfMeasure:0, StockStatus:0 )
  OutboundAllocation.create( allocationNumber:"test string for allocationNumber", allocatedQuantity:"test value", allocationDate:1.week.ago, Status:0 )
end
