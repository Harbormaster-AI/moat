package router

import (

    StockKeepingUnitController "inventory-on-golang/internal/controller"
    WarehouseController "inventory-on-golang/internal/controller"
    StorageLocationController "inventory-on-golang/internal/controller"
    InventoryItemController "inventory-on-golang/internal/controller"
    LotController "inventory-on-golang/internal/controller"
    SerialNumberController "inventory-on-golang/internal/controller"
    ReservationController "inventory-on-golang/internal/controller"
    DemandSignalController "inventory-on-golang/internal/controller"
    InventoryTransactionController "inventory-on-golang/internal/controller"
    TransferOrderController "inventory-on-golang/internal/controller"
    TransferOrderLineController "inventory-on-golang/internal/controller"
    StockAdjustmentController "inventory-on-golang/internal/controller"
    StockAdjustmentLineController "inventory-on-golang/internal/controller"
    CycleCountController "inventory-on-golang/internal/controller"
    CycleCountEntryController "inventory-on-golang/internal/controller"
    ReplenishmentPolicyController "inventory-on-golang/internal/controller"
    UoMConversionController "inventory-on-golang/internal/controller"
    InventoryThresholdAlertController "inventory-on-golang/internal/controller"
    QuarantineController "inventory-on-golang/internal/controller"
    ExpirationPolicyController "inventory-on-golang/internal/controller"
    InboundShipmentController "inventory-on-golang/internal/controller"
    InboundShipmentLineController "inventory-on-golang/internal/controller"
    OutboundAllocationController "inventory-on-golang/internal/controller"
    jsonResponseFormatter "inventory-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "inventory-on-golang/internal/controller"

)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    //----------------------------------------------------------------------------
    // default controllers for health and availability checking
    //----------------------------------------------------------------------------

    router.HandleFunc("/", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Default__)).Methods("GET", "OPTIONS")
    router.HandleFunc("/health", jsonResponseFormatter.FormatToJSON(PulseIndicatorController__.Health__)).Methods("GET", "OPTIONS")


    //----------------------------------------------------------------------------
    // StockKeepingUnit Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StockKeepingUnit/{id}", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.GetStockKeepingUnit)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/StockKeepingUnit", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.GetAllStockKeepingUnit)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewStockKeepingUnit", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.CreateStockKeepingUnit)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StockKeepingUnit/{id}", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.UpdateStockKeepingUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteStockKeepingUnit/{id}", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.DeleteStockKeepingUnit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInventoryItemsToStockKeepingUnit/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.AddInventoryItemsToStockKeepingUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromStockKeepingUnit/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.RemoveInventoryItemsFromStockKeepingUnit)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddUomConversionsToStockKeepingUnit/{parentId}/uomConversionsId", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.AddUomConversionsToStockKeepingUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveUomConversionsFromStockKeepingUnit/{parentId}/uomConversionsIds", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.RemoveUomConversionsFromStockKeepingUnit)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReplenishmentPoliciesToStockKeepingUnit/{parentId}/replenishmentPoliciesId", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.AddReplenishmentPoliciesToStockKeepingUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReplenishmentPoliciesFromStockKeepingUnit/{parentId}/replenishmentPoliciesIds", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.RemoveReplenishmentPoliciesFromStockKeepingUnit)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddLotsToStockKeepingUnit/{parentId}/lotsId", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.AddLotsToStockKeepingUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLotsFromStockKeepingUnit/{parentId}/lotsIds", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.RemoveLotsFromStockKeepingUnit)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSerialNumbersToStockKeepingUnit/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.AddSerialNumbersToStockKeepingUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromStockKeepingUnit/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(StockKeepingUnitController.RemoveSerialNumbersFromStockKeepingUnit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Warehouse Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Warehouse/{id}", jsonResponseFormatter.FormatToJSON(WarehouseController.GetWarehouse)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Warehouse", jsonResponseFormatter.FormatToJSON(WarehouseController.GetAllWarehouse)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWarehouse", jsonResponseFormatter.FormatToJSON(WarehouseController.CreateWarehouse)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Warehouse/{id}", jsonResponseFormatter.FormatToJSON(WarehouseController.UpdateWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWarehouse/{id}", jsonResponseFormatter.FormatToJSON(WarehouseController.DeleteWarehouse)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddStorageLocationsToWarehouse/{parentId}/storageLocationsId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddStorageLocationsToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveStorageLocationsFromWarehouse/{parentId}/storageLocationsIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveStorageLocationsFromWarehouse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInventoryItemsToWarehouse/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddInventoryItemsToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromWarehouse/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveInventoryItemsFromWarehouse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInboundShipmentsToWarehouse/{parentId}/inboundShipmentsId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddInboundShipmentsToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInboundShipmentsFromWarehouse/{parentId}/inboundShipmentsIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveInboundShipmentsFromWarehouse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOutboundAllocationsToWarehouse/{parentId}/outboundAllocationsId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddOutboundAllocationsToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOutboundAllocationsFromWarehouse/{parentId}/outboundAllocationsIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveOutboundAllocationsFromWarehouse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddOriginTransfersToWarehouse/{parentId}/originTransfersId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddOriginTransfersToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOriginTransfersFromWarehouse/{parentId}/originTransfersIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveOriginTransfersFromWarehouse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddDestinationTransfersToWarehouse/{parentId}/destinationTransfersId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddDestinationTransfersToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveDestinationTransfersFromWarehouse/{parentId}/destinationTransfersIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveDestinationTransfersFromWarehouse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCycleCountsToWarehouse/{parentId}/cycleCountsId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddCycleCountsToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCycleCountsFromWarehouse/{parentId}/cycleCountsIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveCycleCountsFromWarehouse)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // StorageLocation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StorageLocation/{id}", jsonResponseFormatter.FormatToJSON(StorageLocationController.GetStorageLocation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/StorageLocation", jsonResponseFormatter.FormatToJSON(StorageLocationController.GetAllStorageLocation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewStorageLocation", jsonResponseFormatter.FormatToJSON(StorageLocationController.CreateStorageLocation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StorageLocation/{id}", jsonResponseFormatter.FormatToJSON(StorageLocationController.UpdateStorageLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteStorageLocation/{id}", jsonResponseFormatter.FormatToJSON(StorageLocationController.DeleteStorageLocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWarehouseToStorageLocation/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(StorageLocationController.AssignWarehouseToStorageLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromStorageLocation/{parentId}", jsonResponseFormatter.FormatToJSON(StorageLocationController.UnassignWarehouseFromStorageLocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignParentLocationToStorageLocation/{parentId}/parentLocationId", jsonResponseFormatter.FormatToJSON(StorageLocationController.AssignParentLocationToStorageLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignParentLocationFromStorageLocation/{parentId}", jsonResponseFormatter.FormatToJSON(StorageLocationController.UnassignParentLocationFromStorageLocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddChildLocationsToStorageLocation/{parentId}/childLocationsId", jsonResponseFormatter.FormatToJSON(StorageLocationController.AddChildLocationsToStorageLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveChildLocationsFromStorageLocation/{parentId}/childLocationsIds", jsonResponseFormatter.FormatToJSON(StorageLocationController.RemoveChildLocationsFromStorageLocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInventoryItemsToStorageLocation/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(StorageLocationController.AddInventoryItemsToStorageLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromStorageLocation/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(StorageLocationController.RemoveInventoryItemsFromStorageLocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InventoryItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.GetInventoryItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InventoryItem", jsonResponseFormatter.FormatToJSON(InventoryItemController.GetAllInventoryItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInventoryItem", jsonResponseFormatter.FormatToJSON(InventoryItemController.CreateInventoryItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UpdateInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInventoryItem/{id}", jsonResponseFormatter.FormatToJSON(InventoryItemController.DeleteInventoryItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToInventoryItem/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignSkuToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignSkuFromInventoryItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarehouseToInventoryItem/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignWarehouseToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignWarehouseFromInventoryItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToInventoryItem/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignLocationToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignLocationFromInventoryItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToInventoryItem/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignLotToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignLotFromInventoryItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSerialNumbersToInventoryItem/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AddSerialNumbersToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromInventoryItem/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(InventoryItemController.RemoveSerialNumbersFromInventoryItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTransactionsToInventoryItem/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AddTransactionsToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromInventoryItem/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(InventoryItemController.RemoveTransactionsFromInventoryItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddReservationsToInventoryItem/{parentId}/reservationsId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AddReservationsToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReservationsFromInventoryItem/{parentId}/reservationsIds", jsonResponseFormatter.FormatToJSON(InventoryItemController.RemoveReservationsFromInventoryItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Lot Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Lot/{id}", jsonResponseFormatter.FormatToJSON(LotController.GetLot)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Lot", jsonResponseFormatter.FormatToJSON(LotController.GetAllLot)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLot", jsonResponseFormatter.FormatToJSON(LotController.CreateLot)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Lot/{id}", jsonResponseFormatter.FormatToJSON(LotController.UpdateLot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLot/{id}", jsonResponseFormatter.FormatToJSON(LotController.DeleteLot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToLot/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(LotController.AssignSkuToLot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromLot/{parentId}", jsonResponseFormatter.FormatToJSON(LotController.UnassignSkuFromLot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInventoryItemsToLot/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(LotController.AddInventoryItemsToLot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromLot/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(LotController.RemoveInventoryItemsFromLot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SerialNumber Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SerialNumber/{id}", jsonResponseFormatter.FormatToJSON(SerialNumberController.GetSerialNumber)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SerialNumber", jsonResponseFormatter.FormatToJSON(SerialNumberController.GetAllSerialNumber)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSerialNumber", jsonResponseFormatter.FormatToJSON(SerialNumberController.CreateSerialNumber)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SerialNumber/{id}", jsonResponseFormatter.FormatToJSON(SerialNumberController.UpdateSerialNumber)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSerialNumber/{id}", jsonResponseFormatter.FormatToJSON(SerialNumberController.DeleteSerialNumber)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToSerialNumber/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(SerialNumberController.AssignSkuToSerialNumber)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromSerialNumber/{parentId}", jsonResponseFormatter.FormatToJSON(SerialNumberController.UnassignSkuFromSerialNumber)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCurrentInventoryItemToSerialNumber/{parentId}/currentInventoryItemId", jsonResponseFormatter.FormatToJSON(SerialNumberController.AssignCurrentInventoryItemToSerialNumber)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCurrentInventoryItemFromSerialNumber/{parentId}", jsonResponseFormatter.FormatToJSON(SerialNumberController.UnassignCurrentInventoryItemFromSerialNumber)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToSerialNumber/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(SerialNumberController.AssignLotToSerialNumber)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromSerialNumber/{parentId}", jsonResponseFormatter.FormatToJSON(SerialNumberController.UnassignLotFromSerialNumber)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Reservation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Reservation/{id}", jsonResponseFormatter.FormatToJSON(ReservationController.GetReservation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Reservation", jsonResponseFormatter.FormatToJSON(ReservationController.GetAllReservation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewReservation", jsonResponseFormatter.FormatToJSON(ReservationController.CreateReservation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Reservation/{id}", jsonResponseFormatter.FormatToJSON(ReservationController.UpdateReservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteReservation/{id}", jsonResponseFormatter.FormatToJSON(ReservationController.DeleteReservation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToReservation/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(ReservationController.AssignSkuToReservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromReservation/{parentId}", jsonResponseFormatter.FormatToJSON(ReservationController.UnassignSkuFromReservation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarehouseToReservation/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(ReservationController.AssignWarehouseToReservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromReservation/{parentId}", jsonResponseFormatter.FormatToJSON(ReservationController.UnassignWarehouseFromReservation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToReservation/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(ReservationController.AssignLocationToReservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromReservation/{parentId}", jsonResponseFormatter.FormatToJSON(ReservationController.UnassignLocationFromReservation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInventoryItemToReservation/{parentId}/inventoryItemId", jsonResponseFormatter.FormatToJSON(ReservationController.AssignInventoryItemToReservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInventoryItemFromReservation/{parentId}", jsonResponseFormatter.FormatToJSON(ReservationController.UnassignInventoryItemFromReservation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToReservation/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(ReservationController.AssignLotToReservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromReservation/{parentId}", jsonResponseFormatter.FormatToJSON(ReservationController.UnassignLotFromReservation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDemandSignalToReservation/{parentId}/demandSignalId", jsonResponseFormatter.FormatToJSON(ReservationController.AssignDemandSignalToReservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDemandSignalFromReservation/{parentId}", jsonResponseFormatter.FormatToJSON(ReservationController.UnassignDemandSignalFromReservation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSerialNumbersToReservation/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(ReservationController.AddSerialNumbersToReservation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromReservation/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(ReservationController.RemoveSerialNumbersFromReservation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // DemandSignal Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/DemandSignal/{id}", jsonResponseFormatter.FormatToJSON(DemandSignalController.GetDemandSignal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/DemandSignal", jsonResponseFormatter.FormatToJSON(DemandSignalController.GetAllDemandSignal)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewDemandSignal", jsonResponseFormatter.FormatToJSON(DemandSignalController.CreateDemandSignal)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/DemandSignal/{id}", jsonResponseFormatter.FormatToJSON(DemandSignalController.UpdateDemandSignal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteDemandSignal/{id}", jsonResponseFormatter.FormatToJSON(DemandSignalController.DeleteDemandSignal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToDemandSignal/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(DemandSignalController.AssignSkuToDemandSignal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromDemandSignal/{parentId}", jsonResponseFormatter.FormatToJSON(DemandSignalController.UnassignSkuFromDemandSignal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddReservationsToDemandSignal/{parentId}/reservationsId", jsonResponseFormatter.FormatToJSON(DemandSignalController.AddReservationsToDemandSignal)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveReservationsFromDemandSignal/{parentId}/reservationsIds", jsonResponseFormatter.FormatToJSON(DemandSignalController.RemoveReservationsFromDemandSignal)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InventoryTransaction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InventoryTransaction/{id}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.GetInventoryTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InventoryTransaction", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.GetAllInventoryTransaction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInventoryTransaction", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.CreateInventoryTransaction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InventoryTransaction/{id}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UpdateInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInventoryTransaction/{id}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.DeleteInventoryTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToInventoryTransaction/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignSkuToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignSkuFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarehouseToInventoryTransaction/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignWarehouseToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignWarehouseFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToInventoryTransaction/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignLocationToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignLocationFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToInventoryTransaction/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignLotToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignLotFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRelatedReservationToInventoryTransaction/{parentId}/relatedReservationId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignRelatedReservationToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRelatedReservationFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignRelatedReservationFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignTransferOrderToInventoryTransaction/{parentId}/transferOrderId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignTransferOrderToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransferOrderFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignTransferOrderFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignAdjustmentToInventoryTransaction/{parentId}/adjustmentId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignAdjustmentToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdjustmentFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignAdjustmentFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCycleCountToInventoryTransaction/{parentId}/cycleCountId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignCycleCountToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCycleCountFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignCycleCountFromInventoryTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSerialNumbersToInventoryTransaction/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AddSerialNumbersToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromInventoryTransaction/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.RemoveSerialNumbersFromInventoryTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // TransferOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TransferOrder/{id}", jsonResponseFormatter.FormatToJSON(TransferOrderController.GetTransferOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TransferOrder", jsonResponseFormatter.FormatToJSON(TransferOrderController.GetAllTransferOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTransferOrder", jsonResponseFormatter.FormatToJSON(TransferOrderController.CreateTransferOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TransferOrder/{id}", jsonResponseFormatter.FormatToJSON(TransferOrderController.UpdateTransferOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTransferOrder/{id}", jsonResponseFormatter.FormatToJSON(TransferOrderController.DeleteTransferOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignOriginWarehouseToTransferOrder/{parentId}/originWarehouseId", jsonResponseFormatter.FormatToJSON(TransferOrderController.AssignOriginWarehouseToTransferOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOriginWarehouseFromTransferOrder/{parentId}", jsonResponseFormatter.FormatToJSON(TransferOrderController.UnassignOriginWarehouseFromTransferOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDestinationWarehouseToTransferOrder/{parentId}/destinationWarehouseId", jsonResponseFormatter.FormatToJSON(TransferOrderController.AssignDestinationWarehouseToTransferOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDestinationWarehouseFromTransferOrder/{parentId}", jsonResponseFormatter.FormatToJSON(TransferOrderController.UnassignDestinationWarehouseFromTransferOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLinesToTransferOrder/{parentId}/linesId", jsonResponseFormatter.FormatToJSON(TransferOrderController.AddLinesToTransferOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLinesFromTransferOrder/{parentId}/linesIds", jsonResponseFormatter.FormatToJSON(TransferOrderController.RemoveLinesFromTransferOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTransactionsToTransferOrder/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(TransferOrderController.AddTransactionsToTransferOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromTransferOrder/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(TransferOrderController.RemoveTransactionsFromTransferOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // TransferOrderLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/TransferOrderLine/{id}", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.GetTransferOrderLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/TransferOrderLine", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.GetAllTransferOrderLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewTransferOrderLine", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.CreateTransferOrderLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/TransferOrderLine/{id}", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.UpdateTransferOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteTransferOrderLine/{id}", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.DeleteTransferOrderLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignTransferOrderToTransferOrderLine/{parentId}/transferOrderId", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.AssignTransferOrderToTransferOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignTransferOrderFromTransferOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.UnassignTransferOrderFromTransferOrderLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSkuToTransferOrderLine/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.AssignSkuToTransferOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromTransferOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.UnassignSkuFromTransferOrderLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToTransferOrderLine/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.AssignLotToTransferOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromTransferOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.UnassignLotFromTransferOrderLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignFromLocationToTransferOrderLine/{parentId}/fromLocationId", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.AssignFromLocationToTransferOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignFromLocationFromTransferOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.UnassignFromLocationFromTransferOrderLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignToLocationToTransferOrderLine/{parentId}/toLocationId", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.AssignToLocationToTransferOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignToLocationFromTransferOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.UnassignToLocationFromTransferOrderLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSerialNumbersToTransferOrderLine/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.AddSerialNumbersToTransferOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromTransferOrderLine/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(TransferOrderLineController.RemoveSerialNumbersFromTransferOrderLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // StockAdjustment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StockAdjustment/{id}", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.GetStockAdjustment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/StockAdjustment", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.GetAllStockAdjustment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewStockAdjustment", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.CreateStockAdjustment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StockAdjustment/{id}", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.UpdateStockAdjustment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteStockAdjustment/{id}", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.DeleteStockAdjustment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWarehouseToStockAdjustment/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.AssignWarehouseToStockAdjustment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromStockAdjustment/{parentId}", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.UnassignWarehouseFromStockAdjustment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLinesToStockAdjustment/{parentId}/linesId", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.AddLinesToStockAdjustment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLinesFromStockAdjustment/{parentId}/linesIds", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.RemoveLinesFromStockAdjustment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTransactionsToStockAdjustment/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.AddTransactionsToStockAdjustment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromStockAdjustment/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(StockAdjustmentController.RemoveTransactionsFromStockAdjustment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // StockAdjustmentLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/StockAdjustmentLine/{id}", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.GetStockAdjustmentLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/StockAdjustmentLine", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.GetAllStockAdjustmentLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewStockAdjustmentLine", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.CreateStockAdjustmentLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/StockAdjustmentLine/{id}", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.UpdateStockAdjustmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteStockAdjustmentLine/{id}", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.DeleteStockAdjustmentLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAdjustmentToStockAdjustmentLine/{parentId}/adjustmentId", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.AssignAdjustmentToStockAdjustmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAdjustmentFromStockAdjustmentLine/{parentId}", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.UnassignAdjustmentFromStockAdjustmentLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSkuToStockAdjustmentLine/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.AssignSkuToStockAdjustmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromStockAdjustmentLine/{parentId}", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.UnassignSkuFromStockAdjustmentLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToStockAdjustmentLine/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.AssignLotToStockAdjustmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromStockAdjustmentLine/{parentId}", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.UnassignLotFromStockAdjustmentLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToStockAdjustmentLine/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.AssignLocationToStockAdjustmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromStockAdjustmentLine/{parentId}", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.UnassignLocationFromStockAdjustmentLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSerialNumbersToStockAdjustmentLine/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.AddSerialNumbersToStockAdjustmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromStockAdjustmentLine/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(StockAdjustmentLineController.RemoveSerialNumbersFromStockAdjustmentLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CycleCount Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CycleCount/{id}", jsonResponseFormatter.FormatToJSON(CycleCountController.GetCycleCount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CycleCount", jsonResponseFormatter.FormatToJSON(CycleCountController.GetAllCycleCount)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCycleCount", jsonResponseFormatter.FormatToJSON(CycleCountController.CreateCycleCount)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CycleCount/{id}", jsonResponseFormatter.FormatToJSON(CycleCountController.UpdateCycleCount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCycleCount/{id}", jsonResponseFormatter.FormatToJSON(CycleCountController.DeleteCycleCount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWarehouseToCycleCount/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(CycleCountController.AssignWarehouseToCycleCount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromCycleCount/{parentId}", jsonResponseFormatter.FormatToJSON(CycleCountController.UnassignWarehouseFromCycleCount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLocationsToCycleCount/{parentId}/locationsId", jsonResponseFormatter.FormatToJSON(CycleCountController.AddLocationsToCycleCount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLocationsFromCycleCount/{parentId}/locationsIds", jsonResponseFormatter.FormatToJSON(CycleCountController.RemoveLocationsFromCycleCount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddEntriesToCycleCount/{parentId}/entriesId", jsonResponseFormatter.FormatToJSON(CycleCountController.AddEntriesToCycleCount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEntriesFromCycleCount/{parentId}/entriesIds", jsonResponseFormatter.FormatToJSON(CycleCountController.RemoveEntriesFromCycleCount)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTransactionsToCycleCount/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(CycleCountController.AddTransactionsToCycleCount)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromCycleCount/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(CycleCountController.RemoveTransactionsFromCycleCount)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // CycleCountEntry Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CycleCountEntry/{id}", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.GetCycleCountEntry)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CycleCountEntry", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.GetAllCycleCountEntry)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCycleCountEntry", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.CreateCycleCountEntry)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CycleCountEntry/{id}", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.UpdateCycleCountEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCycleCountEntry/{id}", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.DeleteCycleCountEntry)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCycleCountToCycleCountEntry/{parentId}/cycleCountId", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.AssignCycleCountToCycleCountEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCycleCountFromCycleCountEntry/{parentId}", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.UnassignCycleCountFromCycleCountEntry)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSkuToCycleCountEntry/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.AssignSkuToCycleCountEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromCycleCountEntry/{parentId}", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.UnassignSkuFromCycleCountEntry)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToCycleCountEntry/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.AssignLotToCycleCountEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromCycleCountEntry/{parentId}", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.UnassignLotFromCycleCountEntry)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToCycleCountEntry/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.AssignLocationToCycleCountEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromCycleCountEntry/{parentId}", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.UnassignLocationFromCycleCountEntry)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSerialNumbersToCycleCountEntry/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.AddSerialNumbersToCycleCountEntry)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromCycleCountEntry/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(CycleCountEntryController.RemoveSerialNumbersFromCycleCountEntry)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ReplenishmentPolicy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ReplenishmentPolicy/{id}", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.GetReplenishmentPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ReplenishmentPolicy", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.GetAllReplenishmentPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewReplenishmentPolicy", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.CreateReplenishmentPolicy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ReplenishmentPolicy/{id}", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.UpdateReplenishmentPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteReplenishmentPolicy/{id}", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.DeleteReplenishmentPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToReplenishmentPolicy/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.AssignSkuToReplenishmentPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromReplenishmentPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.UnassignSkuFromReplenishmentPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarehouseToReplenishmentPolicy/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.AssignWarehouseToReplenishmentPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromReplenishmentPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.UnassignWarehouseFromReplenishmentPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToReplenishmentPolicy/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.AssignLocationToReplenishmentPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromReplenishmentPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(ReplenishmentPolicyController.UnassignLocationFromReplenishmentPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // UoMConversion Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/UoMConversion/{id}", jsonResponseFormatter.FormatToJSON(UoMConversionController.GetUoMConversion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/UoMConversion", jsonResponseFormatter.FormatToJSON(UoMConversionController.GetAllUoMConversion)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewUoMConversion", jsonResponseFormatter.FormatToJSON(UoMConversionController.CreateUoMConversion)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/UoMConversion/{id}", jsonResponseFormatter.FormatToJSON(UoMConversionController.UpdateUoMConversion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteUoMConversion/{id}", jsonResponseFormatter.FormatToJSON(UoMConversionController.DeleteUoMConversion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToUoMConversion/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(UoMConversionController.AssignSkuToUoMConversion)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromUoMConversion/{parentId}", jsonResponseFormatter.FormatToJSON(UoMConversionController.UnassignSkuFromUoMConversion)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // InventoryThresholdAlert Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InventoryThresholdAlert/{id}", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.GetInventoryThresholdAlert)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InventoryThresholdAlert", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.GetAllInventoryThresholdAlert)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInventoryThresholdAlert", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.CreateInventoryThresholdAlert)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InventoryThresholdAlert/{id}", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.UpdateInventoryThresholdAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInventoryThresholdAlert/{id}", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.DeleteInventoryThresholdAlert)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToInventoryThresholdAlert/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.AssignSkuToInventoryThresholdAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromInventoryThresholdAlert/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.UnassignSkuFromInventoryThresholdAlert)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarehouseToInventoryThresholdAlert/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.AssignWarehouseToInventoryThresholdAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromInventoryThresholdAlert/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.UnassignWarehouseFromInventoryThresholdAlert)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToInventoryThresholdAlert/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.AssignLocationToInventoryThresholdAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromInventoryThresholdAlert/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.UnassignLocationFromInventoryThresholdAlert)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRelatedPolicyToInventoryThresholdAlert/{parentId}/relatedPolicyId", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.AssignRelatedPolicyToInventoryThresholdAlert)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRelatedPolicyFromInventoryThresholdAlert/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryThresholdAlertController.UnassignRelatedPolicyFromInventoryThresholdAlert)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Quarantine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Quarantine/{id}", jsonResponseFormatter.FormatToJSON(QuarantineController.GetQuarantine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Quarantine", jsonResponseFormatter.FormatToJSON(QuarantineController.GetAllQuarantine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewQuarantine", jsonResponseFormatter.FormatToJSON(QuarantineController.CreateQuarantine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Quarantine/{id}", jsonResponseFormatter.FormatToJSON(QuarantineController.UpdateQuarantine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteQuarantine/{id}", jsonResponseFormatter.FormatToJSON(QuarantineController.DeleteQuarantine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWarehouseToQuarantine/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(QuarantineController.AssignWarehouseToQuarantine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromQuarantine/{parentId}", jsonResponseFormatter.FormatToJSON(QuarantineController.UnassignWarehouseFromQuarantine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToQuarantine/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(QuarantineController.AssignLotToQuarantine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromQuarantine/{parentId}", jsonResponseFormatter.FormatToJSON(QuarantineController.UnassignLotFromQuarantine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddItemsToQuarantine/{parentId}/itemsId", jsonResponseFormatter.FormatToJSON(QuarantineController.AddItemsToQuarantine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveItemsFromQuarantine/{parentId}/itemsIds", jsonResponseFormatter.FormatToJSON(QuarantineController.RemoveItemsFromQuarantine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSerialNumbersToQuarantine/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(QuarantineController.AddSerialNumbersToQuarantine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromQuarantine/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(QuarantineController.RemoveSerialNumbersFromQuarantine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ExpirationPolicy Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ExpirationPolicy/{id}", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.GetExpirationPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ExpirationPolicy", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.GetAllExpirationPolicy)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewExpirationPolicy", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.CreateExpirationPolicy)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ExpirationPolicy/{id}", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.UpdateExpirationPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteExpirationPolicy/{id}", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.DeleteExpirationPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSkuToExpirationPolicy/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.AssignSkuToExpirationPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromExpirationPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.UnassignSkuFromExpirationPolicy)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarehouseToExpirationPolicy/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.AssignWarehouseToExpirationPolicy)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromExpirationPolicy/{parentId}", jsonResponseFormatter.FormatToJSON(ExpirationPolicyController.UnassignWarehouseFromExpirationPolicy)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // InboundShipment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InboundShipment/{id}", jsonResponseFormatter.FormatToJSON(InboundShipmentController.GetInboundShipment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InboundShipment", jsonResponseFormatter.FormatToJSON(InboundShipmentController.GetAllInboundShipment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInboundShipment", jsonResponseFormatter.FormatToJSON(InboundShipmentController.CreateInboundShipment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InboundShipment/{id}", jsonResponseFormatter.FormatToJSON(InboundShipmentController.UpdateInboundShipment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInboundShipment/{id}", jsonResponseFormatter.FormatToJSON(InboundShipmentController.DeleteInboundShipment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWarehouseToInboundShipment/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(InboundShipmentController.AssignWarehouseToInboundShipment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromInboundShipment/{parentId}", jsonResponseFormatter.FormatToJSON(InboundShipmentController.UnassignWarehouseFromInboundShipment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLinesToInboundShipment/{parentId}/linesId", jsonResponseFormatter.FormatToJSON(InboundShipmentController.AddLinesToInboundShipment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLinesFromInboundShipment/{parentId}/linesIds", jsonResponseFormatter.FormatToJSON(InboundShipmentController.RemoveLinesFromInboundShipment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddTransactionsToInboundShipment/{parentId}/transactionsId", jsonResponseFormatter.FormatToJSON(InboundShipmentController.AddTransactionsToInboundShipment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveTransactionsFromInboundShipment/{parentId}/transactionsIds", jsonResponseFormatter.FormatToJSON(InboundShipmentController.RemoveTransactionsFromInboundShipment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InboundShipmentLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InboundShipmentLine/{id}", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.GetInboundShipmentLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InboundShipmentLine", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.GetAllInboundShipmentLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInboundShipmentLine", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.CreateInboundShipmentLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InboundShipmentLine/{id}", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.UpdateInboundShipmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInboundShipmentLine/{id}", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.DeleteInboundShipmentLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInboundShipmentToInboundShipmentLine/{parentId}/inboundShipmentId", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.AssignInboundShipmentToInboundShipmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInboundShipmentFromInboundShipmentLine/{parentId}", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.UnassignInboundShipmentFromInboundShipmentLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSkuToInboundShipmentLine/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.AssignSkuToInboundShipmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromInboundShipmentLine/{parentId}", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.UnassignSkuFromInboundShipmentLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToInboundShipmentLine/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.AssignLotToInboundShipmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromInboundShipmentLine/{parentId}", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.UnassignLotFromInboundShipmentLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignDestinationLocationToInboundShipmentLine/{parentId}/destinationLocationId", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.AssignDestinationLocationToInboundShipmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignDestinationLocationFromInboundShipmentLine/{parentId}", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.UnassignDestinationLocationFromInboundShipmentLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSerialNumbersToInboundShipmentLine/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.AddSerialNumbersToInboundShipmentLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromInboundShipmentLine/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(InboundShipmentLineController.RemoveSerialNumbersFromInboundShipmentLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // OutboundAllocation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/OutboundAllocation/{id}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.GetOutboundAllocation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/OutboundAllocation", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.GetAllOutboundAllocation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOutboundAllocation", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.CreateOutboundAllocation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/OutboundAllocation/{id}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.UpdateOutboundAllocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOutboundAllocation/{id}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.DeleteOutboundAllocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWarehouseToOutboundAllocation/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.AssignWarehouseToOutboundAllocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromOutboundAllocation/{parentId}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.UnassignWarehouseFromOutboundAllocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSkuToOutboundAllocation/{parentId}/skuId", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.AssignSkuToOutboundAllocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSkuFromOutboundAllocation/{parentId}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.UnassignSkuFromOutboundAllocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInventoryItemToOutboundAllocation/{parentId}/inventoryItemId", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.AssignInventoryItemToOutboundAllocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInventoryItemFromOutboundAllocation/{parentId}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.UnassignInventoryItemFromOutboundAllocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignReservationToOutboundAllocation/{parentId}/reservationId", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.AssignReservationToOutboundAllocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignReservationFromOutboundAllocation/{parentId}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.UnassignReservationFromOutboundAllocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLotToOutboundAllocation/{parentId}/lotId", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.AssignLotToOutboundAllocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLotFromOutboundAllocation/{parentId}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.UnassignLotFromOutboundAllocation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSourceLocationToOutboundAllocation/{parentId}/sourceLocationId", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.AssignSourceLocationToOutboundAllocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSourceLocationFromOutboundAllocation/{parentId}", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.UnassignSourceLocationFromOutboundAllocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddSerialNumbersToOutboundAllocation/{parentId}/serialNumbersId", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.AddSerialNumbersToOutboundAllocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSerialNumbersFromOutboundAllocation/{parentId}/serialNumbersIds", jsonResponseFormatter.FormatToJSON(OutboundAllocationController.RemoveSerialNumbersFromOutboundAllocation)).Methods("DELETE", "OPTIONS")

    return router
}
