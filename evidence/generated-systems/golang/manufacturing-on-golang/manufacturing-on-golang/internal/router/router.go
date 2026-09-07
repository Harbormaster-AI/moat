package router

import (

    EnterpriseController "manufacturing-on-golang/internal/controller"
    BusinessUnitController "manufacturing-on-golang/internal/controller"
    PlantController "manufacturing-on-golang/internal/controller"
    ProductionLineController "manufacturing-on-golang/internal/controller"
    WorkCenterController "manufacturing-on-golang/internal/controller"
    ItemController "manufacturing-on-golang/internal/controller"
    BOMController "manufacturing-on-golang/internal/controller"
    BOMItemController "manufacturing-on-golang/internal/controller"
    RoutingController "manufacturing-on-golang/internal/controller"
    OperationController "manufacturing-on-golang/internal/controller"
    WorkOrderController "manufacturing-on-golang/internal/controller"
    ProductionScheduleController "manufacturing-on-golang/internal/controller"
    SupplierController "manufacturing-on-golang/internal/controller"
    PurchaseOrderController "manufacturing-on-golang/internal/controller"
    PurchaseOrderLineController "manufacturing-on-golang/internal/controller"
    GoodsReceiptController "manufacturing-on-golang/internal/controller"
    GoodsReceiptLineController "manufacturing-on-golang/internal/controller"
    WarehouseController "manufacturing-on-golang/internal/controller"
    LocationController "manufacturing-on-golang/internal/controller"
    InventoryItemController "manufacturing-on-golang/internal/controller"
    InventoryTransactionController "manufacturing-on-golang/internal/controller"
    CustomerController "manufacturing-on-golang/internal/controller"
    SalesOrderController "manufacturing-on-golang/internal/controller"
    SalesOrderLineController "manufacturing-on-golang/internal/controller"
    QualitySpecificationController "manufacturing-on-golang/internal/controller"
    InspectionPlanController "manufacturing-on-golang/internal/controller"
    InspectionCharacteristicController "manufacturing-on-golang/internal/controller"
    InspectionLotController "manufacturing-on-golang/internal/controller"
    InspectionResultController "manufacturing-on-golang/internal/controller"
    NonconformanceController "manufacturing-on-golang/internal/controller"
    CorrectiveActionController "manufacturing-on-golang/internal/controller"
    AssetController "manufacturing-on-golang/internal/controller"
    MaintenancePlanController "manufacturing-on-golang/internal/controller"
    MaintenanceOrderController "manufacturing-on-golang/internal/controller"
    EmployeeController "manufacturing-on-golang/internal/controller"
    ShiftController "manufacturing-on-golang/internal/controller"
    ShiftAssignmentController "manufacturing-on-golang/internal/controller"
    ForecastController "manufacturing-on-golang/internal/controller"
    ForecastLineController "manufacturing-on-golang/internal/controller"
    MRPRunController "manufacturing-on-golang/internal/controller"
    PlannedOrderController "manufacturing-on-golang/internal/controller"
    jsonResponseFormatter "manufacturing-on-golang/internal/response"
    "github.com/gorilla/mux"

    PulseIndicatorController__ "manufacturing-on-golang/internal/controller"

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
    // Enterprise Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Enterprise/{id}", jsonResponseFormatter.FormatToJSON(EnterpriseController.GetEnterprise)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Enterprise", jsonResponseFormatter.FormatToJSON(EnterpriseController.GetAllEnterprise)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEnterprise", jsonResponseFormatter.FormatToJSON(EnterpriseController.CreateEnterprise)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Enterprise/{id}", jsonResponseFormatter.FormatToJSON(EnterpriseController.UpdateEnterprise)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEnterprise/{id}", jsonResponseFormatter.FormatToJSON(EnterpriseController.DeleteEnterprise)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddBusinessUnitsToEnterprise/{parentId}/businessUnitsId", jsonResponseFormatter.FormatToJSON(EnterpriseController.AddBusinessUnitsToEnterprise)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBusinessUnitsFromEnterprise/{parentId}/businessUnitsIds", jsonResponseFormatter.FormatToJSON(EnterpriseController.RemoveBusinessUnitsFromEnterprise)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPlantsToEnterprise/{parentId}/plantsId", jsonResponseFormatter.FormatToJSON(EnterpriseController.AddPlantsToEnterprise)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlantsFromEnterprise/{parentId}/plantsIds", jsonResponseFormatter.FormatToJSON(EnterpriseController.RemovePlantsFromEnterprise)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSuppliersToEnterprise/{parentId}/suppliersId", jsonResponseFormatter.FormatToJSON(EnterpriseController.AddSuppliersToEnterprise)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSuppliersFromEnterprise/{parentId}/suppliersIds", jsonResponseFormatter.FormatToJSON(EnterpriseController.RemoveSuppliersFromEnterprise)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCustomersToEnterprise/{parentId}/customersId", jsonResponseFormatter.FormatToJSON(EnterpriseController.AddCustomersToEnterprise)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCustomersFromEnterprise/{parentId}/customersIds", jsonResponseFormatter.FormatToJSON(EnterpriseController.RemoveCustomersFromEnterprise)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BusinessUnit Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BusinessUnit/{id}", jsonResponseFormatter.FormatToJSON(BusinessUnitController.GetBusinessUnit)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BusinessUnit", jsonResponseFormatter.FormatToJSON(BusinessUnitController.GetAllBusinessUnit)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBusinessUnit", jsonResponseFormatter.FormatToJSON(BusinessUnitController.CreateBusinessUnit)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BusinessUnit/{id}", jsonResponseFormatter.FormatToJSON(BusinessUnitController.UpdateBusinessUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBusinessUnit/{id}", jsonResponseFormatter.FormatToJSON(BusinessUnitController.DeleteBusinessUnit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEnterpriseToBusinessUnit/{parentId}/enterpriseId", jsonResponseFormatter.FormatToJSON(BusinessUnitController.AssignEnterpriseToBusinessUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEnterpriseFromBusinessUnit/{parentId}", jsonResponseFormatter.FormatToJSON(BusinessUnitController.UnassignEnterpriseFromBusinessUnit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddItemsToBusinessUnit/{parentId}/itemsId", jsonResponseFormatter.FormatToJSON(BusinessUnitController.AddItemsToBusinessUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveItemsFromBusinessUnit/{parentId}/itemsIds", jsonResponseFormatter.FormatToJSON(BusinessUnitController.RemoveItemsFromBusinessUnit)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPlantsToBusinessUnit/{parentId}/plantsId", jsonResponseFormatter.FormatToJSON(BusinessUnitController.AddPlantsToBusinessUnit)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlantsFromBusinessUnit/{parentId}/plantsIds", jsonResponseFormatter.FormatToJSON(BusinessUnitController.RemovePlantsFromBusinessUnit)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Plant Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Plant/{id}", jsonResponseFormatter.FormatToJSON(PlantController.GetPlant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Plant", jsonResponseFormatter.FormatToJSON(PlantController.GetAllPlant)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPlant", jsonResponseFormatter.FormatToJSON(PlantController.CreatePlant)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Plant/{id}", jsonResponseFormatter.FormatToJSON(PlantController.UpdatePlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePlant/{id}", jsonResponseFormatter.FormatToJSON(PlantController.DeletePlant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignEnterpriseToPlant/{parentId}/enterpriseId", jsonResponseFormatter.FormatToJSON(PlantController.AssignEnterpriseToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEnterpriseFromPlant/{parentId}", jsonResponseFormatter.FormatToJSON(PlantController.UnassignEnterpriseFromPlant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddProductionLinesToPlant/{parentId}/productionLinesId", jsonResponseFormatter.FormatToJSON(PlantController.AddProductionLinesToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductionLinesFromPlant/{parentId}/productionLinesIds", jsonResponseFormatter.FormatToJSON(PlantController.RemoveProductionLinesFromPlant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddWorkCentersToPlant/{parentId}/workCentersId", jsonResponseFormatter.FormatToJSON(PlantController.AddWorkCentersToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkCentersFromPlant/{parentId}/workCentersIds", jsonResponseFormatter.FormatToJSON(PlantController.RemoveWorkCentersFromPlant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddWarehousesToPlant/{parentId}/warehousesId", jsonResponseFormatter.FormatToJSON(PlantController.AddWarehousesToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWarehousesFromPlant/{parentId}/warehousesIds", jsonResponseFormatter.FormatToJSON(PlantController.RemoveWarehousesFromPlant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddAssetsToPlant/{parentId}/assetsId", jsonResponseFormatter.FormatToJSON(PlantController.AddAssetsToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAssetsFromPlant/{parentId}/assetsIds", jsonResponseFormatter.FormatToJSON(PlantController.RemoveAssetsFromPlant)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddProductionSchedulesToPlant/{parentId}/productionSchedulesId", jsonResponseFormatter.FormatToJSON(PlantController.AddProductionSchedulesToPlant)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveProductionSchedulesFromPlant/{parentId}/productionSchedulesIds", jsonResponseFormatter.FormatToJSON(PlantController.RemoveProductionSchedulesFromPlant)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ProductionLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ProductionLine/{id}", jsonResponseFormatter.FormatToJSON(ProductionLineController.GetProductionLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ProductionLine", jsonResponseFormatter.FormatToJSON(ProductionLineController.GetAllProductionLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProductionLine", jsonResponseFormatter.FormatToJSON(ProductionLineController.CreateProductionLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ProductionLine/{id}", jsonResponseFormatter.FormatToJSON(ProductionLineController.UpdateProductionLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProductionLine/{id}", jsonResponseFormatter.FormatToJSON(ProductionLineController.DeleteProductionLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPlantToProductionLine/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(ProductionLineController.AssignPlantToProductionLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromProductionLine/{parentId}", jsonResponseFormatter.FormatToJSON(ProductionLineController.UnassignPlantFromProductionLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddWorkCentersToProductionLine/{parentId}/workCentersId", jsonResponseFormatter.FormatToJSON(ProductionLineController.AddWorkCentersToProductionLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkCentersFromProductionLine/{parentId}/workCentersIds", jsonResponseFormatter.FormatToJSON(ProductionLineController.RemoveWorkCentersFromProductionLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // WorkCenter Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/WorkCenter/{id}", jsonResponseFormatter.FormatToJSON(WorkCenterController.GetWorkCenter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/WorkCenter", jsonResponseFormatter.FormatToJSON(WorkCenterController.GetAllWorkCenter)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWorkCenter", jsonResponseFormatter.FormatToJSON(WorkCenterController.CreateWorkCenter)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/WorkCenter/{id}", jsonResponseFormatter.FormatToJSON(WorkCenterController.UpdateWorkCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWorkCenter/{id}", jsonResponseFormatter.FormatToJSON(WorkCenterController.DeleteWorkCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignProductionLineToWorkCenter/{parentId}/productionLineId", jsonResponseFormatter.FormatToJSON(WorkCenterController.AssignProductionLineToWorkCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductionLineFromWorkCenter/{parentId}", jsonResponseFormatter.FormatToJSON(WorkCenterController.UnassignProductionLineFromWorkCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAssetsToWorkCenter/{parentId}/assetsId", jsonResponseFormatter.FormatToJSON(WorkCenterController.AddAssetsToWorkCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAssetsFromWorkCenter/{parentId}/assetsIds", jsonResponseFormatter.FormatToJSON(WorkCenterController.RemoveAssetsFromWorkCenter)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMaintenanceOrdersToWorkCenter/{parentId}/maintenanceOrdersId", jsonResponseFormatter.FormatToJSON(WorkCenterController.AddMaintenanceOrdersToWorkCenter)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMaintenanceOrdersFromWorkCenter/{parentId}/maintenanceOrdersIds", jsonResponseFormatter.FormatToJSON(WorkCenterController.RemoveMaintenanceOrdersFromWorkCenter)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Item Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Item/{id}", jsonResponseFormatter.FormatToJSON(ItemController.GetItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Item", jsonResponseFormatter.FormatToJSON(ItemController.GetAllItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewItem", jsonResponseFormatter.FormatToJSON(ItemController.CreateItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Item/{id}", jsonResponseFormatter.FormatToJSON(ItemController.UpdateItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteItem/{id}", jsonResponseFormatter.FormatToJSON(ItemController.DeleteItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBusinessUnitToItem/{parentId}/businessUnitId", jsonResponseFormatter.FormatToJSON(ItemController.AssignBusinessUnitToItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBusinessUnitFromItem/{parentId}", jsonResponseFormatter.FormatToJSON(ItemController.UnassignBusinessUnitFromItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddBomsToItem/{parentId}/bomsId", jsonResponseFormatter.FormatToJSON(ItemController.AddBomsToItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBomsFromItem/{parentId}/bomsIds", jsonResponseFormatter.FormatToJSON(ItemController.RemoveBomsFromItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddRoutingsToItem/{parentId}/routingsId", jsonResponseFormatter.FormatToJSON(ItemController.AddRoutingsToItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveRoutingsFromItem/{parentId}/routingsIds", jsonResponseFormatter.FormatToJSON(ItemController.RemoveRoutingsFromItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSuppliersToItem/{parentId}/suppliersId", jsonResponseFormatter.FormatToJSON(ItemController.AddSuppliersToItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSuppliersFromItem/{parentId}/suppliersIds", jsonResponseFormatter.FormatToJSON(ItemController.RemoveSuppliersFromItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddQualitySpecificationsToItem/{parentId}/qualitySpecificationsId", jsonResponseFormatter.FormatToJSON(ItemController.AddQualitySpecificationsToItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveQualitySpecificationsFromItem/{parentId}/qualitySpecificationsIds", jsonResponseFormatter.FormatToJSON(ItemController.RemoveQualitySpecificationsFromItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInventoryItemsToItem/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(ItemController.AddInventoryItemsToItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromItem/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(ItemController.RemoveInventoryItemsFromItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BOM Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BOM/{id}", jsonResponseFormatter.FormatToJSON(BOMController.GetBOM)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BOM", jsonResponseFormatter.FormatToJSON(BOMController.GetAllBOM)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBOM", jsonResponseFormatter.FormatToJSON(BOMController.CreateBOM)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BOM/{id}", jsonResponseFormatter.FormatToJSON(BOMController.UpdateBOM)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBOM/{id}", jsonResponseFormatter.FormatToJSON(BOMController.DeleteBOM)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignParentItemToBOM/{parentId}/parentItemId", jsonResponseFormatter.FormatToJSON(BOMController.AssignParentItemToBOM)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignParentItemFromBOM/{parentId}", jsonResponseFormatter.FormatToJSON(BOMController.UnassignParentItemFromBOM)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddBomItemsToBOM/{parentId}/bomItemsId", jsonResponseFormatter.FormatToJSON(BOMController.AddBomItemsToBOM)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveBomItemsFromBOM/{parentId}/bomItemsIds", jsonResponseFormatter.FormatToJSON(BOMController.RemoveBomItemsFromBOM)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // BOMItem Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/BOMItem/{id}", jsonResponseFormatter.FormatToJSON(BOMItemController.GetBOMItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/BOMItem", jsonResponseFormatter.FormatToJSON(BOMItemController.GetAllBOMItem)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewBOMItem", jsonResponseFormatter.FormatToJSON(BOMItemController.CreateBOMItem)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/BOMItem/{id}", jsonResponseFormatter.FormatToJSON(BOMItemController.UpdateBOMItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteBOMItem/{id}", jsonResponseFormatter.FormatToJSON(BOMItemController.DeleteBOMItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignBomToBOMItem/{parentId}/bomId", jsonResponseFormatter.FormatToJSON(BOMItemController.AssignBomToBOMItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBomFromBOMItem/{parentId}", jsonResponseFormatter.FormatToJSON(BOMItemController.UnassignBomFromBOMItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignComponentToBOMItem/{parentId}/componentId", jsonResponseFormatter.FormatToJSON(BOMItemController.AssignComponentToBOMItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignComponentFromBOMItem/{parentId}", jsonResponseFormatter.FormatToJSON(BOMItemController.UnassignComponentFromBOMItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Routing Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Routing/{id}", jsonResponseFormatter.FormatToJSON(RoutingController.GetRouting)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Routing", jsonResponseFormatter.FormatToJSON(RoutingController.GetAllRouting)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewRouting", jsonResponseFormatter.FormatToJSON(RoutingController.CreateRouting)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Routing/{id}", jsonResponseFormatter.FormatToJSON(RoutingController.UpdateRouting)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteRouting/{id}", jsonResponseFormatter.FormatToJSON(RoutingController.DeleteRouting)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignItemToRouting/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(RoutingController.AssignItemToRouting)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromRouting/{parentId}", jsonResponseFormatter.FormatToJSON(RoutingController.UnassignItemFromRouting)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddOperationsToRouting/{parentId}/operationsId", jsonResponseFormatter.FormatToJSON(RoutingController.AddOperationsToRouting)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveOperationsFromRouting/{parentId}/operationsIds", jsonResponseFormatter.FormatToJSON(RoutingController.RemoveOperationsFromRouting)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Operation Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Operation/{id}", jsonResponseFormatter.FormatToJSON(OperationController.GetOperation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Operation", jsonResponseFormatter.FormatToJSON(OperationController.GetAllOperation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewOperation", jsonResponseFormatter.FormatToJSON(OperationController.CreateOperation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Operation/{id}", jsonResponseFormatter.FormatToJSON(OperationController.UpdateOperation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteOperation/{id}", jsonResponseFormatter.FormatToJSON(OperationController.DeleteOperation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignRoutingToOperation/{parentId}/routingId", jsonResponseFormatter.FormatToJSON(OperationController.AssignRoutingToOperation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRoutingFromOperation/{parentId}", jsonResponseFormatter.FormatToJSON(OperationController.UnassignRoutingFromOperation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkCenterToOperation/{parentId}/workCenterId", jsonResponseFormatter.FormatToJSON(OperationController.AssignWorkCenterToOperation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkCenterFromOperation/{parentId}", jsonResponseFormatter.FormatToJSON(OperationController.UnassignWorkCenterFromOperation)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInspectionPlanToOperation/{parentId}/inspectionPlanId", jsonResponseFormatter.FormatToJSON(OperationController.AssignInspectionPlanToOperation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInspectionPlanFromOperation/{parentId}", jsonResponseFormatter.FormatToJSON(OperationController.UnassignInspectionPlanFromOperation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // WorkOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/WorkOrder/{id}", jsonResponseFormatter.FormatToJSON(WorkOrderController.GetWorkOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/WorkOrder", jsonResponseFormatter.FormatToJSON(WorkOrderController.GetAllWorkOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewWorkOrder", jsonResponseFormatter.FormatToJSON(WorkOrderController.CreateWorkOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/WorkOrder/{id}", jsonResponseFormatter.FormatToJSON(WorkOrderController.UpdateWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteWorkOrder/{id}", jsonResponseFormatter.FormatToJSON(WorkOrderController.DeleteWorkOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignItemToWorkOrder/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(WorkOrderController.AssignItemToWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(WorkOrderController.UnassignItemFromWorkOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlantToWorkOrder/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(WorkOrderController.AssignPlantToWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(WorkOrderController.UnassignPlantFromWorkOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignRoutingToWorkOrder/{parentId}/routingId", jsonResponseFormatter.FormatToJSON(WorkOrderController.AssignRoutingToWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignRoutingFromWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(WorkOrderController.UnassignRoutingFromWorkOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignBomToWorkOrder/{parentId}/bomId", jsonResponseFormatter.FormatToJSON(WorkOrderController.AssignBomToWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignBomFromWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(WorkOrderController.UnassignBomFromWorkOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignProductionScheduleToWorkOrder/{parentId}/productionScheduleId", jsonResponseFormatter.FormatToJSON(WorkOrderController.AssignProductionScheduleToWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignProductionScheduleFromWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(WorkOrderController.UnassignProductionScheduleFromWorkOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSalesOrderToWorkOrder/{parentId}/salesOrderId", jsonResponseFormatter.FormatToJSON(WorkOrderController.AssignSalesOrderToWorkOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSalesOrderFromWorkOrder/{parentId}", jsonResponseFormatter.FormatToJSON(WorkOrderController.UnassignSalesOrderFromWorkOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // ProductionSchedule Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ProductionSchedule/{id}", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.GetProductionSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ProductionSchedule", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.GetAllProductionSchedule)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewProductionSchedule", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.CreateProductionSchedule)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ProductionSchedule/{id}", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.UpdateProductionSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteProductionSchedule/{id}", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.DeleteProductionSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPlantToProductionSchedule/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.AssignPlantToProductionSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromProductionSchedule/{parentId}", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.UnassignPlantFromProductionSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddWorkOrdersToProductionSchedule/{parentId}/workOrdersId", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.AddWorkOrdersToProductionSchedule)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkOrdersFromProductionSchedule/{parentId}/workOrdersIds", jsonResponseFormatter.FormatToJSON(ProductionScheduleController.RemoveWorkOrdersFromProductionSchedule)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Supplier Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Supplier/{id}", jsonResponseFormatter.FormatToJSON(SupplierController.GetSupplier)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Supplier", jsonResponseFormatter.FormatToJSON(SupplierController.GetAllSupplier)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSupplier", jsonResponseFormatter.FormatToJSON(SupplierController.CreateSupplier)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Supplier/{id}", jsonResponseFormatter.FormatToJSON(SupplierController.UpdateSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSupplier/{id}", jsonResponseFormatter.FormatToJSON(SupplierController.DeleteSupplier)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEnterprisesToSupplier/{parentId}/enterprisesId", jsonResponseFormatter.FormatToJSON(SupplierController.AddEnterprisesToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEnterprisesFromSupplier/{parentId}/enterprisesIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemoveEnterprisesFromSupplier)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddItemsToSupplier/{parentId}/itemsId", jsonResponseFormatter.FormatToJSON(SupplierController.AddItemsToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveItemsFromSupplier/{parentId}/itemsIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemoveItemsFromSupplier)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddPurchaseOrdersToSupplier/{parentId}/purchaseOrdersId", jsonResponseFormatter.FormatToJSON(SupplierController.AddPurchaseOrdersToSupplier)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePurchaseOrdersFromSupplier/{parentId}/purchaseOrdersIds", jsonResponseFormatter.FormatToJSON(SupplierController.RemovePurchaseOrdersFromSupplier)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PurchaseOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PurchaseOrder/{id}", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.GetPurchaseOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PurchaseOrder", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.GetAllPurchaseOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPurchaseOrder", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.CreatePurchaseOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PurchaseOrder/{id}", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.UpdatePurchaseOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePurchaseOrder/{id}", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.DeletePurchaseOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSupplierToPurchaseOrder/{parentId}/supplierId", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.AssignSupplierToPurchaseOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSupplierFromPurchaseOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.UnassignSupplierFromPurchaseOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlantToPurchaseOrder/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.AssignPlantToPurchaseOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromPurchaseOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.UnassignPlantFromPurchaseOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLinesToPurchaseOrder/{parentId}/linesId", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.AddLinesToPurchaseOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLinesFromPurchaseOrder/{parentId}/linesIds", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.RemoveLinesFromPurchaseOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddGoodsReceiptsToPurchaseOrder/{parentId}/goodsReceiptsId", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.AddGoodsReceiptsToPurchaseOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveGoodsReceiptsFromPurchaseOrder/{parentId}/goodsReceiptsIds", jsonResponseFormatter.FormatToJSON(PurchaseOrderController.RemoveGoodsReceiptsFromPurchaseOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PurchaseOrderLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PurchaseOrderLine/{id}", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.GetPurchaseOrderLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PurchaseOrderLine", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.GetAllPurchaseOrderLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPurchaseOrderLine", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.CreatePurchaseOrderLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PurchaseOrderLine/{id}", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.UpdatePurchaseOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePurchaseOrderLine/{id}", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.DeletePurchaseOrderLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPurchaseOrderToPurchaseOrderLine/{parentId}/purchaseOrderId", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.AssignPurchaseOrderToPurchaseOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPurchaseOrderFromPurchaseOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.UnassignPurchaseOrderFromPurchaseOrderLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignItemToPurchaseOrderLine/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.AssignItemToPurchaseOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromPurchaseOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(PurchaseOrderLineController.UnassignItemFromPurchaseOrderLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // GoodsReceipt Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/GoodsReceipt/{id}", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.GetGoodsReceipt)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/GoodsReceipt", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.GetAllGoodsReceipt)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewGoodsReceipt", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.CreateGoodsReceipt)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/GoodsReceipt/{id}", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.UpdateGoodsReceipt)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteGoodsReceipt/{id}", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.DeleteGoodsReceipt)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPurchaseOrderToGoodsReceipt/{parentId}/purchaseOrderId", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.AssignPurchaseOrderToGoodsReceipt)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPurchaseOrderFromGoodsReceipt/{parentId}", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.UnassignPurchaseOrderFromGoodsReceipt)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWarehouseToGoodsReceipt/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.AssignWarehouseToGoodsReceipt)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromGoodsReceipt/{parentId}", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.UnassignWarehouseFromGoodsReceipt)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLinesToGoodsReceipt/{parentId}/linesId", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.AddLinesToGoodsReceipt)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLinesFromGoodsReceipt/{parentId}/linesIds", jsonResponseFormatter.FormatToJSON(GoodsReceiptController.RemoveLinesFromGoodsReceipt)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // GoodsReceiptLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/GoodsReceiptLine/{id}", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.GetGoodsReceiptLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/GoodsReceiptLine", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.GetAllGoodsReceiptLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewGoodsReceiptLine", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.CreateGoodsReceiptLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/GoodsReceiptLine/{id}", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.UpdateGoodsReceiptLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteGoodsReceiptLine/{id}", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.DeleteGoodsReceiptLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignGoodsReceiptToGoodsReceiptLine/{parentId}/goodsReceiptId", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.AssignGoodsReceiptToGoodsReceiptLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignGoodsReceiptFromGoodsReceiptLine/{parentId}", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.UnassignGoodsReceiptFromGoodsReceiptLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignItemToGoodsReceiptLine/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.AssignItemToGoodsReceiptLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromGoodsReceiptLine/{parentId}", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.UnassignItemFromGoodsReceiptLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInventoryTransactionToGoodsReceiptLine/{parentId}/inventoryTransactionId", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.AssignInventoryTransactionToGoodsReceiptLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInventoryTransactionFromGoodsReceiptLine/{parentId}", jsonResponseFormatter.FormatToJSON(GoodsReceiptLineController.UnassignInventoryTransactionFromGoodsReceiptLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

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
    router.HandleFunc("/api/AssignPlantToWarehouse/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(WarehouseController.AssignPlantToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromWarehouse/{parentId}", jsonResponseFormatter.FormatToJSON(WarehouseController.UnassignPlantFromWarehouse)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLocationsToWarehouse/{parentId}/locationsId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddLocationsToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLocationsFromWarehouse/{parentId}/locationsIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveLocationsFromWarehouse)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddInventoryItemsToWarehouse/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(WarehouseController.AddInventoryItemsToWarehouse)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromWarehouse/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(WarehouseController.RemoveInventoryItemsFromWarehouse)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Location Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Location/{id}", jsonResponseFormatter.FormatToJSON(LocationController.GetLocation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Location", jsonResponseFormatter.FormatToJSON(LocationController.GetAllLocation)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewLocation", jsonResponseFormatter.FormatToJSON(LocationController.CreateLocation)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Location/{id}", jsonResponseFormatter.FormatToJSON(LocationController.UpdateLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteLocation/{id}", jsonResponseFormatter.FormatToJSON(LocationController.DeleteLocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWarehouseToLocation/{parentId}/warehouseId", jsonResponseFormatter.FormatToJSON(LocationController.AssignWarehouseToLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWarehouseFromLocation/{parentId}", jsonResponseFormatter.FormatToJSON(LocationController.UnassignWarehouseFromLocation)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddInventoryItemsToLocation/{parentId}/inventoryItemsId", jsonResponseFormatter.FormatToJSON(LocationController.AddInventoryItemsToLocation)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveInventoryItemsFromLocation/{parentId}/inventoryItemsIds", jsonResponseFormatter.FormatToJSON(LocationController.RemoveInventoryItemsFromLocation)).Methods("DELETE", "OPTIONS")

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
    router.HandleFunc("/api/AssignItemToInventoryItem/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignItemToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignItemFromInventoryItem)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToInventoryItem/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(InventoryItemController.AssignLocationToInventoryItem)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromInventoryItem/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryItemController.UnassignLocationFromInventoryItem)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

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
    router.HandleFunc("/api/AssignItemToInventoryTransaction/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignItemToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignItemFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignLocationToInventoryTransaction/{parentId}/locationId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignLocationToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignLocationFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignLocationFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkOrderToInventoryTransaction/{parentId}/workOrderId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignWorkOrderToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkOrderFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignWorkOrderFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPurchaseOrderToInventoryTransaction/{parentId}/purchaseOrderId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignPurchaseOrderToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPurchaseOrderFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignPurchaseOrderFromInventoryTransaction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignSalesOrderToInventoryTransaction/{parentId}/salesOrderId", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.AssignSalesOrderToInventoryTransaction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSalesOrderFromInventoryTransaction/{parentId}", jsonResponseFormatter.FormatToJSON(InventoryTransactionController.UnassignSalesOrderFromInventoryTransaction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Customer Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Customer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.GetCustomer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Customer", jsonResponseFormatter.FormatToJSON(CustomerController.GetAllCustomer)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCustomer", jsonResponseFormatter.FormatToJSON(CustomerController.CreateCustomer)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Customer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.UpdateCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCustomer/{id}", jsonResponseFormatter.FormatToJSON(CustomerController.DeleteCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddEnterprisesToCustomer/{parentId}/enterprisesId", jsonResponseFormatter.FormatToJSON(CustomerController.AddEnterprisesToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveEnterprisesFromCustomer/{parentId}/enterprisesIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveEnterprisesFromCustomer)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddSalesOrdersToCustomer/{parentId}/salesOrdersId", jsonResponseFormatter.FormatToJSON(CustomerController.AddSalesOrdersToCustomer)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveSalesOrdersFromCustomer/{parentId}/salesOrdersIds", jsonResponseFormatter.FormatToJSON(CustomerController.RemoveSalesOrdersFromCustomer)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SalesOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SalesOrder/{id}", jsonResponseFormatter.FormatToJSON(SalesOrderController.GetSalesOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SalesOrder", jsonResponseFormatter.FormatToJSON(SalesOrderController.GetAllSalesOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSalesOrder", jsonResponseFormatter.FormatToJSON(SalesOrderController.CreateSalesOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SalesOrder/{id}", jsonResponseFormatter.FormatToJSON(SalesOrderController.UpdateSalesOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSalesOrder/{id}", jsonResponseFormatter.FormatToJSON(SalesOrderController.DeleteSalesOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignCustomerToSalesOrder/{parentId}/customerId", jsonResponseFormatter.FormatToJSON(SalesOrderController.AssignCustomerToSalesOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCustomerFromSalesOrder/{parentId}", jsonResponseFormatter.FormatToJSON(SalesOrderController.UnassignCustomerFromSalesOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlantToSalesOrder/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(SalesOrderController.AssignPlantToSalesOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromSalesOrder/{parentId}", jsonResponseFormatter.FormatToJSON(SalesOrderController.UnassignPlantFromSalesOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLinesToSalesOrder/{parentId}/linesId", jsonResponseFormatter.FormatToJSON(SalesOrderController.AddLinesToSalesOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLinesFromSalesOrder/{parentId}/linesIds", jsonResponseFormatter.FormatToJSON(SalesOrderController.RemoveLinesFromSalesOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddWorkOrdersToSalesOrder/{parentId}/workOrdersId", jsonResponseFormatter.FormatToJSON(SalesOrderController.AddWorkOrdersToSalesOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveWorkOrdersFromSalesOrder/{parentId}/workOrdersIds", jsonResponseFormatter.FormatToJSON(SalesOrderController.RemoveWorkOrdersFromSalesOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // SalesOrderLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/SalesOrderLine/{id}", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.GetSalesOrderLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/SalesOrderLine", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.GetAllSalesOrderLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewSalesOrderLine", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.CreateSalesOrderLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/SalesOrderLine/{id}", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.UpdateSalesOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteSalesOrderLine/{id}", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.DeleteSalesOrderLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignSalesOrderToSalesOrderLine/{parentId}/salesOrderId", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.AssignSalesOrderToSalesOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignSalesOrderFromSalesOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.UnassignSalesOrderFromSalesOrderLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignItemToSalesOrderLine/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.AssignItemToSalesOrderLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromSalesOrderLine/{parentId}", jsonResponseFormatter.FormatToJSON(SalesOrderLineController.UnassignItemFromSalesOrderLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // QualitySpecification Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/QualitySpecification/{id}", jsonResponseFormatter.FormatToJSON(QualitySpecificationController.GetQualitySpecification)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/QualitySpecification", jsonResponseFormatter.FormatToJSON(QualitySpecificationController.GetAllQualitySpecification)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewQualitySpecification", jsonResponseFormatter.FormatToJSON(QualitySpecificationController.CreateQualitySpecification)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/QualitySpecification/{id}", jsonResponseFormatter.FormatToJSON(QualitySpecificationController.UpdateQualitySpecification)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteQualitySpecification/{id}", jsonResponseFormatter.FormatToJSON(QualitySpecificationController.DeleteQualitySpecification)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignItemToQualitySpecification/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(QualitySpecificationController.AssignItemToQualitySpecification)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromQualitySpecification/{parentId}", jsonResponseFormatter.FormatToJSON(QualitySpecificationController.UnassignItemFromQualitySpecification)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // InspectionPlan Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InspectionPlan/{id}", jsonResponseFormatter.FormatToJSON(InspectionPlanController.GetInspectionPlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InspectionPlan", jsonResponseFormatter.FormatToJSON(InspectionPlanController.GetAllInspectionPlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInspectionPlan", jsonResponseFormatter.FormatToJSON(InspectionPlanController.CreateInspectionPlan)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InspectionPlan/{id}", jsonResponseFormatter.FormatToJSON(InspectionPlanController.UpdateInspectionPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInspectionPlan/{id}", jsonResponseFormatter.FormatToJSON(InspectionPlanController.DeleteInspectionPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignItemToInspectionPlan/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(InspectionPlanController.AssignItemToInspectionPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromInspectionPlan/{parentId}", jsonResponseFormatter.FormatToJSON(InspectionPlanController.UnassignItemFromInspectionPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddCharacteristicsToInspectionPlan/{parentId}/characteristicsId", jsonResponseFormatter.FormatToJSON(InspectionPlanController.AddCharacteristicsToInspectionPlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCharacteristicsFromInspectionPlan/{parentId}/characteristicsIds", jsonResponseFormatter.FormatToJSON(InspectionPlanController.RemoveCharacteristicsFromInspectionPlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InspectionCharacteristic Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InspectionCharacteristic/{id}", jsonResponseFormatter.FormatToJSON(InspectionCharacteristicController.GetInspectionCharacteristic)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InspectionCharacteristic", jsonResponseFormatter.FormatToJSON(InspectionCharacteristicController.GetAllInspectionCharacteristic)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInspectionCharacteristic", jsonResponseFormatter.FormatToJSON(InspectionCharacteristicController.CreateInspectionCharacteristic)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InspectionCharacteristic/{id}", jsonResponseFormatter.FormatToJSON(InspectionCharacteristicController.UpdateInspectionCharacteristic)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInspectionCharacteristic/{id}", jsonResponseFormatter.FormatToJSON(InspectionCharacteristicController.DeleteInspectionCharacteristic)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInspectionPlanToInspectionCharacteristic/{parentId}/inspectionPlanId", jsonResponseFormatter.FormatToJSON(InspectionCharacteristicController.AssignInspectionPlanToInspectionCharacteristic)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInspectionPlanFromInspectionCharacteristic/{parentId}", jsonResponseFormatter.FormatToJSON(InspectionCharacteristicController.UnassignInspectionPlanFromInspectionCharacteristic)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // InspectionLot Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InspectionLot/{id}", jsonResponseFormatter.FormatToJSON(InspectionLotController.GetInspectionLot)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InspectionLot", jsonResponseFormatter.FormatToJSON(InspectionLotController.GetAllInspectionLot)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInspectionLot", jsonResponseFormatter.FormatToJSON(InspectionLotController.CreateInspectionLot)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InspectionLot/{id}", jsonResponseFormatter.FormatToJSON(InspectionLotController.UpdateInspectionLot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInspectionLot/{id}", jsonResponseFormatter.FormatToJSON(InspectionLotController.DeleteInspectionLot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignItemToInspectionLot/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(InspectionLotController.AssignItemToInspectionLot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromInspectionLot/{parentId}", jsonResponseFormatter.FormatToJSON(InspectionLotController.UnassignItemFromInspectionLot)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkOrderToInspectionLot/{parentId}/workOrderId", jsonResponseFormatter.FormatToJSON(InspectionLotController.AssignWorkOrderToInspectionLot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkOrderFromInspectionLot/{parentId}", jsonResponseFormatter.FormatToJSON(InspectionLotController.UnassignWorkOrderFromInspectionLot)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignGoodsReceiptToInspectionLot/{parentId}/goodsReceiptId", jsonResponseFormatter.FormatToJSON(InspectionLotController.AssignGoodsReceiptToInspectionLot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignGoodsReceiptFromInspectionLot/{parentId}", jsonResponseFormatter.FormatToJSON(InspectionLotController.UnassignGoodsReceiptFromInspectionLot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddResultsToInspectionLot/{parentId}/resultsId", jsonResponseFormatter.FormatToJSON(InspectionLotController.AddResultsToInspectionLot)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveResultsFromInspectionLot/{parentId}/resultsIds", jsonResponseFormatter.FormatToJSON(InspectionLotController.RemoveResultsFromInspectionLot)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // InspectionResult Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/InspectionResult/{id}", jsonResponseFormatter.FormatToJSON(InspectionResultController.GetInspectionResult)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/InspectionResult", jsonResponseFormatter.FormatToJSON(InspectionResultController.GetAllInspectionResult)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewInspectionResult", jsonResponseFormatter.FormatToJSON(InspectionResultController.CreateInspectionResult)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/InspectionResult/{id}", jsonResponseFormatter.FormatToJSON(InspectionResultController.UpdateInspectionResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteInspectionResult/{id}", jsonResponseFormatter.FormatToJSON(InspectionResultController.DeleteInspectionResult)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignInspectionLotToInspectionResult/{parentId}/inspectionLotId", jsonResponseFormatter.FormatToJSON(InspectionResultController.AssignInspectionLotToInspectionResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInspectionLotFromInspectionResult/{parentId}", jsonResponseFormatter.FormatToJSON(InspectionResultController.UnassignInspectionLotFromInspectionResult)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCharacteristicToInspectionResult/{parentId}/characteristicId", jsonResponseFormatter.FormatToJSON(InspectionResultController.AssignCharacteristicToInspectionResult)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCharacteristicFromInspectionResult/{parentId}", jsonResponseFormatter.FormatToJSON(InspectionResultController.UnassignCharacteristicFromInspectionResult)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Nonconformance Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Nonconformance/{id}", jsonResponseFormatter.FormatToJSON(NonconformanceController.GetNonconformance)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Nonconformance", jsonResponseFormatter.FormatToJSON(NonconformanceController.GetAllNonconformance)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewNonconformance", jsonResponseFormatter.FormatToJSON(NonconformanceController.CreateNonconformance)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Nonconformance/{id}", jsonResponseFormatter.FormatToJSON(NonconformanceController.UpdateNonconformance)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteNonconformance/{id}", jsonResponseFormatter.FormatToJSON(NonconformanceController.DeleteNonconformance)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignItemToNonconformance/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(NonconformanceController.AssignItemToNonconformance)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromNonconformance/{parentId}", jsonResponseFormatter.FormatToJSON(NonconformanceController.UnassignItemFromNonconformance)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkOrderToNonconformance/{parentId}/workOrderId", jsonResponseFormatter.FormatToJSON(NonconformanceController.AssignWorkOrderToNonconformance)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkOrderFromNonconformance/{parentId}", jsonResponseFormatter.FormatToJSON(NonconformanceController.UnassignWorkOrderFromNonconformance)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignInspectionLotToNonconformance/{parentId}/inspectionLotId", jsonResponseFormatter.FormatToJSON(NonconformanceController.AssignInspectionLotToNonconformance)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignInspectionLotFromNonconformance/{parentId}", jsonResponseFormatter.FormatToJSON(NonconformanceController.UnassignInspectionLotFromNonconformance)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignCorrectiveActionToNonconformance/{parentId}/correctiveActionId", jsonResponseFormatter.FormatToJSON(NonconformanceController.AssignCorrectiveActionToNonconformance)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignCorrectiveActionFromNonconformance/{parentId}", jsonResponseFormatter.FormatToJSON(NonconformanceController.UnassignCorrectiveActionFromNonconformance)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // CorrectiveAction Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/CorrectiveAction/{id}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.GetCorrectiveAction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/CorrectiveAction", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.GetAllCorrectiveAction)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewCorrectiveAction", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.CreateCorrectiveAction)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/CorrectiveAction/{id}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.UpdateCorrectiveAction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteCorrectiveAction/{id}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.DeleteCorrectiveAction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignNonconformanceToCorrectiveAction/{parentId}/nonconformanceId", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.AssignNonconformanceToCorrectiveAction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignNonconformanceFromCorrectiveAction/{parentId}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.UnassignNonconformanceFromCorrectiveAction)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignOwnerToCorrectiveAction/{parentId}/ownerId", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.AssignOwnerToCorrectiveAction)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignOwnerFromCorrectiveAction/{parentId}", jsonResponseFormatter.FormatToJSON(CorrectiveActionController.UnassignOwnerFromCorrectiveAction)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Asset Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Asset/{id}", jsonResponseFormatter.FormatToJSON(AssetController.GetAsset)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Asset", jsonResponseFormatter.FormatToJSON(AssetController.GetAllAsset)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewAsset", jsonResponseFormatter.FormatToJSON(AssetController.CreateAsset)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Asset/{id}", jsonResponseFormatter.FormatToJSON(AssetController.UpdateAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteAsset/{id}", jsonResponseFormatter.FormatToJSON(AssetController.DeleteAsset)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPlantToAsset/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(AssetController.AssignPlantToAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromAsset/{parentId}", jsonResponseFormatter.FormatToJSON(AssetController.UnassignPlantFromAsset)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkCenterToAsset/{parentId}/workCenterId", jsonResponseFormatter.FormatToJSON(AssetController.AssignWorkCenterToAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkCenterFromAsset/{parentId}", jsonResponseFormatter.FormatToJSON(AssetController.UnassignWorkCenterFromAsset)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddMaintenanceOrdersToAsset/{parentId}/maintenanceOrdersId", jsonResponseFormatter.FormatToJSON(AssetController.AddMaintenanceOrdersToAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMaintenanceOrdersFromAsset/{parentId}/maintenanceOrdersIds", jsonResponseFormatter.FormatToJSON(AssetController.RemoveMaintenanceOrdersFromAsset)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddMaintenancePlansToAsset/{parentId}/maintenancePlansId", jsonResponseFormatter.FormatToJSON(AssetController.AddMaintenancePlansToAsset)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMaintenancePlansFromAsset/{parentId}/maintenancePlansIds", jsonResponseFormatter.FormatToJSON(AssetController.RemoveMaintenancePlansFromAsset)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // MaintenancePlan Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MaintenancePlan/{id}", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.GetMaintenancePlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MaintenancePlan", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.GetAllMaintenancePlan)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMaintenancePlan", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.CreateMaintenancePlan)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MaintenancePlan/{id}", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.UpdateMaintenancePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMaintenancePlan/{id}", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.DeleteMaintenancePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAssetToMaintenancePlan/{parentId}/assetId", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.AssignAssetToMaintenancePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAssetFromMaintenancePlan/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.UnassignAssetFromMaintenancePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddMaintenanceOrdersToMaintenancePlan/{parentId}/maintenanceOrdersId", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.AddMaintenanceOrdersToMaintenancePlan)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveMaintenanceOrdersFromMaintenancePlan/{parentId}/maintenanceOrdersIds", jsonResponseFormatter.FormatToJSON(MaintenancePlanController.RemoveMaintenanceOrdersFromMaintenancePlan)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // MaintenanceOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MaintenanceOrder/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.GetMaintenanceOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MaintenanceOrder", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.GetAllMaintenanceOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMaintenanceOrder", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.CreateMaintenanceOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MaintenanceOrder/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.UpdateMaintenanceOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMaintenanceOrder/{id}", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.DeleteMaintenanceOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignAssetToMaintenanceOrder/{parentId}/assetId", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.AssignAssetToMaintenanceOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignAssetFromMaintenanceOrder/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.UnassignAssetFromMaintenanceOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlanToMaintenanceOrder/{parentId}/planId", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.AssignPlanToMaintenanceOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlanFromMaintenanceOrder/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.UnassignPlanFromMaintenanceOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkCenterToMaintenanceOrder/{parentId}/workCenterId", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.AssignWorkCenterToMaintenanceOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkCenterFromMaintenanceOrder/{parentId}", jsonResponseFormatter.FormatToJSON(MaintenanceOrderController.UnassignWorkCenterFromMaintenanceOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Employee Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Employee/{id}", jsonResponseFormatter.FormatToJSON(EmployeeController.GetEmployee)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Employee", jsonResponseFormatter.FormatToJSON(EmployeeController.GetAllEmployee)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewEmployee", jsonResponseFormatter.FormatToJSON(EmployeeController.CreateEmployee)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Employee/{id}", jsonResponseFormatter.FormatToJSON(EmployeeController.UpdateEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteEmployee/{id}", jsonResponseFormatter.FormatToJSON(EmployeeController.DeleteEmployee)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignWorkCenterToEmployee/{parentId}/workCenterId", jsonResponseFormatter.FormatToJSON(EmployeeController.AssignWorkCenterToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkCenterFromEmployee/{parentId}", jsonResponseFormatter.FormatToJSON(EmployeeController.UnassignWorkCenterFromEmployee)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddShiftAssignmentsToEmployee/{parentId}/shiftAssignmentsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddShiftAssignmentsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveShiftAssignmentsFromEmployee/{parentId}/shiftAssignmentsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveShiftAssignmentsFromEmployee)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AddCorrectiveActionsToEmployee/{parentId}/correctiveActionsId", jsonResponseFormatter.FormatToJSON(EmployeeController.AddCorrectiveActionsToEmployee)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveCorrectiveActionsFromEmployee/{parentId}/correctiveActionsIds", jsonResponseFormatter.FormatToJSON(EmployeeController.RemoveCorrectiveActionsFromEmployee)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Shift Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Shift/{id}", jsonResponseFormatter.FormatToJSON(ShiftController.GetShift)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Shift", jsonResponseFormatter.FormatToJSON(ShiftController.GetAllShift)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewShift", jsonResponseFormatter.FormatToJSON(ShiftController.CreateShift)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Shift/{id}", jsonResponseFormatter.FormatToJSON(ShiftController.UpdateShift)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteShift/{id}", jsonResponseFormatter.FormatToJSON(ShiftController.DeleteShift)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPlantToShift/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(ShiftController.AssignPlantToShift)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromShift/{parentId}", jsonResponseFormatter.FormatToJSON(ShiftController.UnassignPlantFromShift)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddAssignmentsToShift/{parentId}/assignmentsId", jsonResponseFormatter.FormatToJSON(ShiftController.AddAssignmentsToShift)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveAssignmentsFromShift/{parentId}/assignmentsIds", jsonResponseFormatter.FormatToJSON(ShiftController.RemoveAssignmentsFromShift)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ShiftAssignment Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ShiftAssignment/{id}", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.GetShiftAssignment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ShiftAssignment", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.GetAllShiftAssignment)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewShiftAssignment", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.CreateShiftAssignment)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ShiftAssignment/{id}", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.UpdateShiftAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteShiftAssignment/{id}", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.DeleteShiftAssignment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignShiftToShiftAssignment/{parentId}/shiftId", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.AssignShiftToShiftAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignShiftFromShiftAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.UnassignShiftFromShiftAssignment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignEmployeeToShiftAssignment/{parentId}/employeeId", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.AssignEmployeeToShiftAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignEmployeeFromShiftAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.UnassignEmployeeFromShiftAssignment)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignWorkCenterToShiftAssignment/{parentId}/workCenterId", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.AssignWorkCenterToShiftAssignment)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignWorkCenterFromShiftAssignment/{parentId}", jsonResponseFormatter.FormatToJSON(ShiftAssignmentController.UnassignWorkCenterFromShiftAssignment)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Forecast Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/Forecast/{id}", jsonResponseFormatter.FormatToJSON(ForecastController.GetForecast)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/Forecast", jsonResponseFormatter.FormatToJSON(ForecastController.GetAllForecast)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewForecast", jsonResponseFormatter.FormatToJSON(ForecastController.CreateForecast)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/Forecast/{id}", jsonResponseFormatter.FormatToJSON(ForecastController.UpdateForecast)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteForecast/{id}", jsonResponseFormatter.FormatToJSON(ForecastController.DeleteForecast)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddLinesToForecast/{parentId}/linesId", jsonResponseFormatter.FormatToJSON(ForecastController.AddLinesToForecast)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemoveLinesFromForecast/{parentId}/linesIds", jsonResponseFormatter.FormatToJSON(ForecastController.RemoveLinesFromForecast)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // ForecastLine Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/ForecastLine/{id}", jsonResponseFormatter.FormatToJSON(ForecastLineController.GetForecastLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/ForecastLine", jsonResponseFormatter.FormatToJSON(ForecastLineController.GetAllForecastLine)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewForecastLine", jsonResponseFormatter.FormatToJSON(ForecastLineController.CreateForecastLine)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/ForecastLine/{id}", jsonResponseFormatter.FormatToJSON(ForecastLineController.UpdateForecastLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteForecastLine/{id}", jsonResponseFormatter.FormatToJSON(ForecastLineController.DeleteForecastLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignForecastToForecastLine/{parentId}/forecastId", jsonResponseFormatter.FormatToJSON(ForecastLineController.AssignForecastToForecastLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignForecastFromForecastLine/{parentId}", jsonResponseFormatter.FormatToJSON(ForecastLineController.UnassignForecastFromForecastLine)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignItemToForecastLine/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(ForecastLineController.AssignItemToForecastLine)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromForecastLine/{parentId}", jsonResponseFormatter.FormatToJSON(ForecastLineController.UnassignItemFromForecastLine)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // MRPRun Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/MRPRun/{id}", jsonResponseFormatter.FormatToJSON(MRPRunController.GetMRPRun)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/MRPRun", jsonResponseFormatter.FormatToJSON(MRPRunController.GetAllMRPRun)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewMRPRun", jsonResponseFormatter.FormatToJSON(MRPRunController.CreateMRPRun)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/MRPRun/{id}", jsonResponseFormatter.FormatToJSON(MRPRunController.UpdateMRPRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeleteMRPRun/{id}", jsonResponseFormatter.FormatToJSON(MRPRunController.DeleteMRPRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignPlantToMRPRun/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(MRPRunController.AssignPlantToMRPRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromMRPRun/{parentId}", jsonResponseFormatter.FormatToJSON(MRPRunController.UnassignPlantFromMRPRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AddPlannedOrdersToMRPRun/{parentId}/plannedOrdersId", jsonResponseFormatter.FormatToJSON(MRPRunController.AddPlannedOrdersToMRPRun)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/RemovePlannedOrdersFromMRPRun/{parentId}/plannedOrdersIds", jsonResponseFormatter.FormatToJSON(MRPRunController.RemovePlannedOrdersFromMRPRun)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // PlannedOrder Routes to JSON response formatter first
    // then to the correct Controller function
    //----------------------------------------------------------------------------

    //----------------------------------------------------------------------------
    // Standard Lifecycle Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/PlannedOrder/{id}", jsonResponseFormatter.FormatToJSON(PlannedOrderController.GetPlannedOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/PlannedOrder", jsonResponseFormatter.FormatToJSON(PlannedOrderController.GetAllPlannedOrder)).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/NewPlannedOrder", jsonResponseFormatter.FormatToJSON(PlannedOrderController.CreatePlannedOrder)).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/PlannedOrder/{id}", jsonResponseFormatter.FormatToJSON(PlannedOrderController.UpdatePlannedOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/DeletePlannedOrder/{id}", jsonResponseFormatter.FormatToJSON(PlannedOrderController.DeletePlannedOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Single Association Routers
    //----------------------------------------------------------------------------
    router.HandleFunc("/api/AssignMrpRunToPlannedOrder/{parentId}/mrpRunId", jsonResponseFormatter.FormatToJSON(PlannedOrderController.AssignMrpRunToPlannedOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignMrpRunFromPlannedOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PlannedOrderController.UnassignMrpRunFromPlannedOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignItemToPlannedOrder/{parentId}/itemId", jsonResponseFormatter.FormatToJSON(PlannedOrderController.AssignItemToPlannedOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignItemFromPlannedOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PlannedOrderController.UnassignItemFromPlannedOrder)).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/AssignPlantToPlannedOrder/{parentId}/plantId", jsonResponseFormatter.FormatToJSON(PlannedOrderController.AssignPlantToPlannedOrder)).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/UnassignPlantFromPlannedOrder/{parentId}", jsonResponseFormatter.FormatToJSON(PlannedOrderController.UnassignPlantFromPlannedOrder)).Methods("DELETE", "OPTIONS")

    //----------------------------------------------------------------------------
    // Multiple Association Routers
    //----------------------------------------------------------------------------

    return router
}
