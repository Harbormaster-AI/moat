package test

import ( 
	"testing"
    dao "inventory-on-golang/internal/dao"
	"inventory-on-golang/internal/model"
	"inventory-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestStockKeepingUnitCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for StockKeepingUnit
	//----------------------------------------------------------------------------
	StockKeepingUnitObj := model.StockKeepingUnit                                                                                                                                                                                                                                                                                                                            {SkuCode:new SKU(),Name:"test value for Name",Weight:"test value",WeightUnit:"test value for WeightUnit",Volume:"test value",VolumeUnit:"test value for VolumeUnit",ShelfLifeDays:100,HazardousMaterial:true,ItemType:0,UnitOfMeasure:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createStockKeepingUnitRequestResult := dao.CreateStockKeepingUnit( StockKeepingUnitObj )
	
	if createStockKeepingUnitRequestResult.Success == false {
		t.Errorf(createStockKeepingUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Create StockKeepingUnit success...")
	}
	
	createStockKeepingUnitObj,_ := createStockKeepingUnitRequestResult.Data. (model.StockKeepingUnit)

	// --------------------------------------------------------------
	// Check StockKeepingUnit Obj ID
	// --------------------------------------------------------------	
	if createStockKeepingUnitObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for StockKeepingUnit" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getStockKeepingUnitRequestResult := dao.GetStockKeepingUnit( uint64(createStockKeepingUnitObj.ID) )
	
	if getStockKeepingUnitRequestResult.Success == false {
		t.Errorf(getStockKeepingUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Get StockKeepingUnit success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getStockKeepingUnitObj,_ := getStockKeepingUnitRequestResult.Data. (model.StockKeepingUnit)
	compareStockKeepingUnit := cmp.Equal(createStockKeepingUnitObj.ID, getStockKeepingUnitObj.ID)
	
	if  compareStockKeepingUnit == false	{
		t.Errorf( "Created StockKeepingUnit object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllStockKeepingUnitRequestResult := dao.GetAllStockKeepingUnit()

	if getAllStockKeepingUnitRequestResult.Success == false {
			t.Errorf(getAllStockKeepingUnitRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll StockKeepingUnit success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllStockKeepingUnitObj []model.StockKeepingUnit = getAllStockKeepingUnitRequestResult.Data. ([]model.StockKeepingUnit)
		
	equalStockKeepingUnit := cmp.Equal(createStockKeepingUnitObj.ID, getAllStockKeepingUnitObj[len(getAllStockKeepingUnitObj)-1].ID)
		
	if equalStockKeepingUnit == false {
		t.Errorf( "Created object is not equal to the last entry in StockKeepingUnit[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for StockKeepingUnit
	// --------------------------------------------------------------	
	deleteStockKeepingUnitRequestResult := dao.DeleteStockKeepingUnit(uint64(createStockKeepingUnitObj.ID))

	if deleteStockKeepingUnitRequestResult.Success == false {
			t.Errorf(deleteStockKeepingUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion StockKeepingUnit success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getStockKeepingUnitRequestResult = dao.GetStockKeepingUnit( uint64(createStockKeepingUnitObj.ID) )
	
	if getStockKeepingUnitRequestResult.Success == true {
		t.Errorf(getStockKeepingUnitRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWarehouseCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Warehouse
	//----------------------------------------------------------------------------
	WarehouseObj := model.Warehouse                                                                                                                                            {Name:"test value for Name",Code:"test value for Code",Address:new Address(),TimeZone:"test value for TimeZone",AllowsOverAllocation:true}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWarehouseRequestResult := dao.CreateWarehouse( WarehouseObj )
	
	if createWarehouseRequestResult.Success == false {
		t.Errorf(createWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Check Create Warehouse success...")
	}
	
	createWarehouseObj,_ := createWarehouseRequestResult.Data. (model.Warehouse)

	// --------------------------------------------------------------
	// Check Warehouse Obj ID
	// --------------------------------------------------------------	
	if createWarehouseObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Warehouse" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWarehouseRequestResult := dao.GetWarehouse( uint64(createWarehouseObj.ID) )
	
	if getWarehouseRequestResult.Success == false {
		t.Errorf(getWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Check Get Warehouse success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWarehouseObj,_ := getWarehouseRequestResult.Data. (model.Warehouse)
	compareWarehouse := cmp.Equal(createWarehouseObj.ID, getWarehouseObj.ID)
	
	if  compareWarehouse == false	{
		t.Errorf( "Created Warehouse object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWarehouseRequestResult := dao.GetAllWarehouse()

	if getAllWarehouseRequestResult.Success == false {
			t.Errorf(getAllWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Warehouse success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWarehouseObj []model.Warehouse = getAllWarehouseRequestResult.Data. ([]model.Warehouse)
		
	equalWarehouse := cmp.Equal(createWarehouseObj.ID, getAllWarehouseObj[len(getAllWarehouseObj)-1].ID)
		
	if equalWarehouse == false {
		t.Errorf( "Created object is not equal to the last entry in Warehouse[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Warehouse
	// --------------------------------------------------------------	
	deleteWarehouseRequestResult := dao.DeleteWarehouse(uint64(createWarehouseObj.ID))

	if deleteWarehouseRequestResult.Success == false {
			t.Errorf(deleteWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Warehouse success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWarehouseRequestResult = dao.GetWarehouse( uint64(createWarehouseObj.ID) )
	
	if getWarehouseRequestResult.Success == true {
		t.Errorf(getWarehouseRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestStorageLocationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for StorageLocation
	//----------------------------------------------------------------------------
	StorageLocationObj := model.StorageLocation                                                                                                                                                                    {Code:"test value for Code",TemperatureControlled:true,Capacity:"test value",CapacityUnit:"test value for CapacityUnit",LocationType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createStorageLocationRequestResult := dao.CreateStorageLocation( StorageLocationObj )
	
	if createStorageLocationRequestResult.Success == false {
		t.Errorf(createStorageLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Create StorageLocation success...")
	}
	
	createStorageLocationObj,_ := createStorageLocationRequestResult.Data. (model.StorageLocation)

	// --------------------------------------------------------------
	// Check StorageLocation Obj ID
	// --------------------------------------------------------------	
	if createStorageLocationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for StorageLocation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getStorageLocationRequestResult := dao.GetStorageLocation( uint64(createStorageLocationObj.ID) )
	
	if getStorageLocationRequestResult.Success == false {
		t.Errorf(getStorageLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Get StorageLocation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getStorageLocationObj,_ := getStorageLocationRequestResult.Data. (model.StorageLocation)
	compareStorageLocation := cmp.Equal(createStorageLocationObj.ID, getStorageLocationObj.ID)
	
	if  compareStorageLocation == false	{
		t.Errorf( "Created StorageLocation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllStorageLocationRequestResult := dao.GetAllStorageLocation()

	if getAllStorageLocationRequestResult.Success == false {
			t.Errorf(getAllStorageLocationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll StorageLocation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllStorageLocationObj []model.StorageLocation = getAllStorageLocationRequestResult.Data. ([]model.StorageLocation)
		
	equalStorageLocation := cmp.Equal(createStorageLocationObj.ID, getAllStorageLocationObj[len(getAllStorageLocationObj)-1].ID)
		
	if equalStorageLocation == false {
		t.Errorf( "Created object is not equal to the last entry in StorageLocation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for StorageLocation
	// --------------------------------------------------------------	
	deleteStorageLocationRequestResult := dao.DeleteStorageLocation(uint64(createStorageLocationObj.ID))

	if deleteStorageLocationRequestResult.Success == false {
			t.Errorf(deleteStorageLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion StorageLocation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getStorageLocationRequestResult = dao.GetStorageLocation( uint64(createStorageLocationObj.ID) )
	
	if getStorageLocationRequestResult.Success == true {
		t.Errorf(getStorageLocationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInventoryItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InventoryItem
	//----------------------------------------------------------------------------
	InventoryItemObj := model.InventoryItem                                                                                                                                                                                                                                                            {QuantityOnHand:"test value",QuantityAvailable:"test value",QuantityReserved:"test value",UnitCost:new Money(),LastUpdated:time.Now(),StockStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInventoryItemRequestResult := dao.CreateInventoryItem( InventoryItemObj )
	
	if createInventoryItemRequestResult.Success == false {
		t.Errorf(createInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create InventoryItem success...")
	}
	
	createInventoryItemObj,_ := createInventoryItemRequestResult.Data. (model.InventoryItem)

	// --------------------------------------------------------------
	// Check InventoryItem Obj ID
	// --------------------------------------------------------------	
	if createInventoryItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InventoryItem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInventoryItemRequestResult := dao.GetInventoryItem( uint64(createInventoryItemObj.ID) )
	
	if getInventoryItemRequestResult.Success == false {
		t.Errorf(getInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get InventoryItem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInventoryItemObj,_ := getInventoryItemRequestResult.Data. (model.InventoryItem)
	compareInventoryItem := cmp.Equal(createInventoryItemObj.ID, getInventoryItemObj.ID)
	
	if  compareInventoryItem == false	{
		t.Errorf( "Created InventoryItem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInventoryItemRequestResult := dao.GetAllInventoryItem()

	if getAllInventoryItemRequestResult.Success == false {
			t.Errorf(getAllInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InventoryItem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInventoryItemObj []model.InventoryItem = getAllInventoryItemRequestResult.Data. ([]model.InventoryItem)
		
	equalInventoryItem := cmp.Equal(createInventoryItemObj.ID, getAllInventoryItemObj[len(getAllInventoryItemObj)-1].ID)
		
	if equalInventoryItem == false {
		t.Errorf( "Created object is not equal to the last entry in InventoryItem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InventoryItem
	// --------------------------------------------------------------	
	deleteInventoryItemRequestResult := dao.DeleteInventoryItem(uint64(createInventoryItemObj.ID))

	if deleteInventoryItemRequestResult.Success == false {
			t.Errorf(deleteInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InventoryItem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInventoryItemRequestResult = dao.GetInventoryItem( uint64(createInventoryItemObj.ID) )
	
	if getInventoryItemRequestResult.Success == true {
		t.Errorf(getInventoryItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLotCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Lot
	//----------------------------------------------------------------------------
	LotObj := model.Lot                                                                                                                                            {BatchNumber:new BatchNumber(),ManufactureDate:time.Now(),ExpirationDate:time.Now(),LotStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLotRequestResult := dao.CreateLot( LotObj )
	
	if createLotRequestResult.Success == false {
		t.Errorf(createLotRequestResult.Msg)
	} else {
		fmt.Println("Check Create Lot success...")
	}
	
	createLotObj,_ := createLotRequestResult.Data. (model.Lot)

	// --------------------------------------------------------------
	// Check Lot Obj ID
	// --------------------------------------------------------------	
	if createLotObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Lot" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLotRequestResult := dao.GetLot( uint64(createLotObj.ID) )
	
	if getLotRequestResult.Success == false {
		t.Errorf(getLotRequestResult.Msg)
	} else {
		fmt.Println("Check Get Lot success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLotObj,_ := getLotRequestResult.Data. (model.Lot)
	compareLot := cmp.Equal(createLotObj.ID, getLotObj.ID)
	
	if  compareLot == false	{
		t.Errorf( "Created Lot object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLotRequestResult := dao.GetAllLot()

	if getAllLotRequestResult.Success == false {
			t.Errorf(getAllLotRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Lot success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLotObj []model.Lot = getAllLotRequestResult.Data. ([]model.Lot)
		
	equalLot := cmp.Equal(createLotObj.ID, getAllLotObj[len(getAllLotObj)-1].ID)
		
	if equalLot == false {
		t.Errorf( "Created object is not equal to the last entry in Lot[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Lot
	// --------------------------------------------------------------	
	deleteLotRequestResult := dao.DeleteLot(uint64(createLotObj.ID))

	if deleteLotRequestResult.Success == false {
			t.Errorf(deleteLotRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Lot success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLotRequestResult = dao.GetLot( uint64(createLotObj.ID) )
	
	if getLotRequestResult.Success == true {
		t.Errorf(getLotRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSerialNumberCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SerialNumber
	//----------------------------------------------------------------------------
	SerialNumberObj := model.SerialNumber                                                                                    {Serial:new SerialCode(),ActivationDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSerialNumberRequestResult := dao.CreateSerialNumber( SerialNumberObj )
	
	if createSerialNumberRequestResult.Success == false {
		t.Errorf(createSerialNumberRequestResult.Msg)
	} else {
		fmt.Println("Check Create SerialNumber success...")
	}
	
	createSerialNumberObj,_ := createSerialNumberRequestResult.Data. (model.SerialNumber)

	// --------------------------------------------------------------
	// Check SerialNumber Obj ID
	// --------------------------------------------------------------	
	if createSerialNumberObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SerialNumber" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSerialNumberRequestResult := dao.GetSerialNumber( uint64(createSerialNumberObj.ID) )
	
	if getSerialNumberRequestResult.Success == false {
		t.Errorf(getSerialNumberRequestResult.Msg)
	} else {
		fmt.Println("Check Get SerialNumber success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSerialNumberObj,_ := getSerialNumberRequestResult.Data. (model.SerialNumber)
	compareSerialNumber := cmp.Equal(createSerialNumberObj.ID, getSerialNumberObj.ID)
	
	if  compareSerialNumber == false	{
		t.Errorf( "Created SerialNumber object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSerialNumberRequestResult := dao.GetAllSerialNumber()

	if getAllSerialNumberRequestResult.Success == false {
			t.Errorf(getAllSerialNumberRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SerialNumber success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSerialNumberObj []model.SerialNumber = getAllSerialNumberRequestResult.Data. ([]model.SerialNumber)
		
	equalSerialNumber := cmp.Equal(createSerialNumberObj.ID, getAllSerialNumberObj[len(getAllSerialNumberObj)-1].ID)
		
	if equalSerialNumber == false {
		t.Errorf( "Created object is not equal to the last entry in SerialNumber[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SerialNumber
	// --------------------------------------------------------------	
	deleteSerialNumberRequestResult := dao.DeleteSerialNumber(uint64(createSerialNumberObj.ID))

	if deleteSerialNumberRequestResult.Success == false {
			t.Errorf(deleteSerialNumberRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SerialNumber success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSerialNumberRequestResult = dao.GetSerialNumber( uint64(createSerialNumberObj.ID) )
	
	if getSerialNumberRequestResult.Success == true {
		t.Errorf(getSerialNumberRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestReservationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Reservation
	//----------------------------------------------------------------------------
	ReservationObj := model.Reservation                                                                                                                                                                            {ReferenceNumber:"test value for ReferenceNumber",ReservedQuantity:"test value",PromisedDate:time.Now(),ReservationStatus:0,ReservationType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createReservationRequestResult := dao.CreateReservation( ReservationObj )
	
	if createReservationRequestResult.Success == false {
		t.Errorf(createReservationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Reservation success...")
	}
	
	createReservationObj,_ := createReservationRequestResult.Data. (model.Reservation)

	// --------------------------------------------------------------
	// Check Reservation Obj ID
	// --------------------------------------------------------------	
	if createReservationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Reservation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getReservationRequestResult := dao.GetReservation( uint64(createReservationObj.ID) )
	
	if getReservationRequestResult.Success == false {
		t.Errorf(getReservationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Reservation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getReservationObj,_ := getReservationRequestResult.Data. (model.Reservation)
	compareReservation := cmp.Equal(createReservationObj.ID, getReservationObj.ID)
	
	if  compareReservation == false	{
		t.Errorf( "Created Reservation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllReservationRequestResult := dao.GetAllReservation()

	if getAllReservationRequestResult.Success == false {
			t.Errorf(getAllReservationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Reservation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllReservationObj []model.Reservation = getAllReservationRequestResult.Data. ([]model.Reservation)
		
	equalReservation := cmp.Equal(createReservationObj.ID, getAllReservationObj[len(getAllReservationObj)-1].ID)
		
	if equalReservation == false {
		t.Errorf( "Created object is not equal to the last entry in Reservation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Reservation
	// --------------------------------------------------------------	
	deleteReservationRequestResult := dao.DeleteReservation(uint64(createReservationObj.ID))

	if deleteReservationRequestResult.Success == false {
			t.Errorf(deleteReservationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Reservation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getReservationRequestResult = dao.GetReservation( uint64(createReservationObj.ID) )
	
	if getReservationRequestResult.Success == true {
		t.Errorf(getReservationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDemandSignalCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DemandSignal
	//----------------------------------------------------------------------------
	DemandSignalObj := model.DemandSignal                                                                                                                                                            {ExternalReference:"test value for ExternalReference",RequestedDate:time.Now(),Quantity:"test value",DemandType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDemandSignalRequestResult := dao.CreateDemandSignal( DemandSignalObj )
	
	if createDemandSignalRequestResult.Success == false {
		t.Errorf(createDemandSignalRequestResult.Msg)
	} else {
		fmt.Println("Check Create DemandSignal success...")
	}
	
	createDemandSignalObj,_ := createDemandSignalRequestResult.Data. (model.DemandSignal)

	// --------------------------------------------------------------
	// Check DemandSignal Obj ID
	// --------------------------------------------------------------	
	if createDemandSignalObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DemandSignal" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDemandSignalRequestResult := dao.GetDemandSignal( uint64(createDemandSignalObj.ID) )
	
	if getDemandSignalRequestResult.Success == false {
		t.Errorf(getDemandSignalRequestResult.Msg)
	} else {
		fmt.Println("Check Get DemandSignal success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDemandSignalObj,_ := getDemandSignalRequestResult.Data. (model.DemandSignal)
	compareDemandSignal := cmp.Equal(createDemandSignalObj.ID, getDemandSignalObj.ID)
	
	if  compareDemandSignal == false	{
		t.Errorf( "Created DemandSignal object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDemandSignalRequestResult := dao.GetAllDemandSignal()

	if getAllDemandSignalRequestResult.Success == false {
			t.Errorf(getAllDemandSignalRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DemandSignal success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDemandSignalObj []model.DemandSignal = getAllDemandSignalRequestResult.Data. ([]model.DemandSignal)
		
	equalDemandSignal := cmp.Equal(createDemandSignalObj.ID, getAllDemandSignalObj[len(getAllDemandSignalObj)-1].ID)
		
	if equalDemandSignal == false {
		t.Errorf( "Created object is not equal to the last entry in DemandSignal[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DemandSignal
	// --------------------------------------------------------------	
	deleteDemandSignalRequestResult := dao.DeleteDemandSignal(uint64(createDemandSignalObj.ID))

	if deleteDemandSignalRequestResult.Success == false {
			t.Errorf(deleteDemandSignalRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DemandSignal success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDemandSignalRequestResult = dao.GetDemandSignal( uint64(createDemandSignalObj.ID) )
	
	if getDemandSignalRequestResult.Success == true {
		t.Errorf(getDemandSignalRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInventoryTransactionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InventoryTransaction
	//----------------------------------------------------------------------------
	InventoryTransactionObj := model.InventoryTransaction                                                                                                                                                                                                                                            {TransactionNumber:"test value for TransactionNumber",Quantity:"test value",UnitCost:new Money(),TransactionDate:time.Now(),ReasonCode:"test value for ReasonCode",TransactionType:0,UnitOfMeasure:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInventoryTransactionRequestResult := dao.CreateInventoryTransaction( InventoryTransactionObj )
	
	if createInventoryTransactionRequestResult.Success == false {
		t.Errorf(createInventoryTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Create InventoryTransaction success...")
	}
	
	createInventoryTransactionObj,_ := createInventoryTransactionRequestResult.Data. (model.InventoryTransaction)

	// --------------------------------------------------------------
	// Check InventoryTransaction Obj ID
	// --------------------------------------------------------------	
	if createInventoryTransactionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InventoryTransaction" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInventoryTransactionRequestResult := dao.GetInventoryTransaction( uint64(createInventoryTransactionObj.ID) )
	
	if getInventoryTransactionRequestResult.Success == false {
		t.Errorf(getInventoryTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Get InventoryTransaction success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInventoryTransactionObj,_ := getInventoryTransactionRequestResult.Data. (model.InventoryTransaction)
	compareInventoryTransaction := cmp.Equal(createInventoryTransactionObj.ID, getInventoryTransactionObj.ID)
	
	if  compareInventoryTransaction == false	{
		t.Errorf( "Created InventoryTransaction object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInventoryTransactionRequestResult := dao.GetAllInventoryTransaction()

	if getAllInventoryTransactionRequestResult.Success == false {
			t.Errorf(getAllInventoryTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InventoryTransaction success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInventoryTransactionObj []model.InventoryTransaction = getAllInventoryTransactionRequestResult.Data. ([]model.InventoryTransaction)
		
	equalInventoryTransaction := cmp.Equal(createInventoryTransactionObj.ID, getAllInventoryTransactionObj[len(getAllInventoryTransactionObj)-1].ID)
		
	if equalInventoryTransaction == false {
		t.Errorf( "Created object is not equal to the last entry in InventoryTransaction[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InventoryTransaction
	// --------------------------------------------------------------	
	deleteInventoryTransactionRequestResult := dao.DeleteInventoryTransaction(uint64(createInventoryTransactionObj.ID))

	if deleteInventoryTransactionRequestResult.Success == false {
			t.Errorf(deleteInventoryTransactionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InventoryTransaction success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInventoryTransactionRequestResult = dao.GetInventoryTransaction( uint64(createInventoryTransactionObj.ID) )
	
	if getInventoryTransactionRequestResult.Success == true {
		t.Errorf(getInventoryTransactionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTransferOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TransferOrder
	//----------------------------------------------------------------------------
	TransferOrderObj := model.TransferOrder                                                                                                                                                                                                                                                                            {OrderNumber:"test value for OrderNumber",RequestedShipDate:time.Now(),RequestedReceiveDate:time.Now(),ShippedDate:time.Now(),ReceivedDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTransferOrderRequestResult := dao.CreateTransferOrder( TransferOrderObj )
	
	if createTransferOrderRequestResult.Success == false {
		t.Errorf(createTransferOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create TransferOrder success...")
	}
	
	createTransferOrderObj,_ := createTransferOrderRequestResult.Data. (model.TransferOrder)

	// --------------------------------------------------------------
	// Check TransferOrder Obj ID
	// --------------------------------------------------------------	
	if createTransferOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TransferOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTransferOrderRequestResult := dao.GetTransferOrder( uint64(createTransferOrderObj.ID) )
	
	if getTransferOrderRequestResult.Success == false {
		t.Errorf(getTransferOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get TransferOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTransferOrderObj,_ := getTransferOrderRequestResult.Data. (model.TransferOrder)
	compareTransferOrder := cmp.Equal(createTransferOrderObj.ID, getTransferOrderObj.ID)
	
	if  compareTransferOrder == false	{
		t.Errorf( "Created TransferOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTransferOrderRequestResult := dao.GetAllTransferOrder()

	if getAllTransferOrderRequestResult.Success == false {
			t.Errorf(getAllTransferOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TransferOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTransferOrderObj []model.TransferOrder = getAllTransferOrderRequestResult.Data. ([]model.TransferOrder)
		
	equalTransferOrder := cmp.Equal(createTransferOrderObj.ID, getAllTransferOrderObj[len(getAllTransferOrderObj)-1].ID)
		
	if equalTransferOrder == false {
		t.Errorf( "Created object is not equal to the last entry in TransferOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TransferOrder
	// --------------------------------------------------------------	
	deleteTransferOrderRequestResult := dao.DeleteTransferOrder(uint64(createTransferOrderObj.ID))

	if deleteTransferOrderRequestResult.Success == false {
			t.Errorf(deleteTransferOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TransferOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTransferOrderRequestResult = dao.GetTransferOrder( uint64(createTransferOrderObj.ID) )
	
	if getTransferOrderRequestResult.Success == true {
		t.Errorf(getTransferOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTransferOrderLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TransferOrderLine
	//----------------------------------------------------------------------------
	TransferOrderLineObj := model.TransferOrderLine                                                                                                                    {LineNumber:100,Quantity:"test value",UnitOfMeasure:0,StockStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTransferOrderLineRequestResult := dao.CreateTransferOrderLine( TransferOrderLineObj )
	
	if createTransferOrderLineRequestResult.Success == false {
		t.Errorf(createTransferOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create TransferOrderLine success...")
	}
	
	createTransferOrderLineObj,_ := createTransferOrderLineRequestResult.Data. (model.TransferOrderLine)

	// --------------------------------------------------------------
	// Check TransferOrderLine Obj ID
	// --------------------------------------------------------------	
	if createTransferOrderLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TransferOrderLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTransferOrderLineRequestResult := dao.GetTransferOrderLine( uint64(createTransferOrderLineObj.ID) )
	
	if getTransferOrderLineRequestResult.Success == false {
		t.Errorf(getTransferOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get TransferOrderLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTransferOrderLineObj,_ := getTransferOrderLineRequestResult.Data. (model.TransferOrderLine)
	compareTransferOrderLine := cmp.Equal(createTransferOrderLineObj.ID, getTransferOrderLineObj.ID)
	
	if  compareTransferOrderLine == false	{
		t.Errorf( "Created TransferOrderLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTransferOrderLineRequestResult := dao.GetAllTransferOrderLine()

	if getAllTransferOrderLineRequestResult.Success == false {
			t.Errorf(getAllTransferOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TransferOrderLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTransferOrderLineObj []model.TransferOrderLine = getAllTransferOrderLineRequestResult.Data. ([]model.TransferOrderLine)
		
	equalTransferOrderLine := cmp.Equal(createTransferOrderLineObj.ID, getAllTransferOrderLineObj[len(getAllTransferOrderLineObj)-1].ID)
		
	if equalTransferOrderLine == false {
		t.Errorf( "Created object is not equal to the last entry in TransferOrderLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TransferOrderLine
	// --------------------------------------------------------------	
	deleteTransferOrderLineRequestResult := dao.DeleteTransferOrderLine(uint64(createTransferOrderLineObj.ID))

	if deleteTransferOrderLineRequestResult.Success == false {
			t.Errorf(deleteTransferOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TransferOrderLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTransferOrderLineRequestResult = dao.GetTransferOrderLine( uint64(createTransferOrderLineObj.ID) )
	
	if getTransferOrderLineRequestResult.Success == true {
		t.Errorf(getTransferOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestStockAdjustmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for StockAdjustment
	//----------------------------------------------------------------------------
	StockAdjustmentObj := model.StockAdjustment                                                                                                                                                    {AdjustmentNumber:"test value for AdjustmentNumber",Reason:"test value for Reason",AdjustmentDate:time.Now(),AdjustmentType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createStockAdjustmentRequestResult := dao.CreateStockAdjustment( StockAdjustmentObj )
	
	if createStockAdjustmentRequestResult.Success == false {
		t.Errorf(createStockAdjustmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create StockAdjustment success...")
	}
	
	createStockAdjustmentObj,_ := createStockAdjustmentRequestResult.Data. (model.StockAdjustment)

	// --------------------------------------------------------------
	// Check StockAdjustment Obj ID
	// --------------------------------------------------------------	
	if createStockAdjustmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for StockAdjustment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getStockAdjustmentRequestResult := dao.GetStockAdjustment( uint64(createStockAdjustmentObj.ID) )
	
	if getStockAdjustmentRequestResult.Success == false {
		t.Errorf(getStockAdjustmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get StockAdjustment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getStockAdjustmentObj,_ := getStockAdjustmentRequestResult.Data. (model.StockAdjustment)
	compareStockAdjustment := cmp.Equal(createStockAdjustmentObj.ID, getStockAdjustmentObj.ID)
	
	if  compareStockAdjustment == false	{
		t.Errorf( "Created StockAdjustment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllStockAdjustmentRequestResult := dao.GetAllStockAdjustment()

	if getAllStockAdjustmentRequestResult.Success == false {
			t.Errorf(getAllStockAdjustmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll StockAdjustment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllStockAdjustmentObj []model.StockAdjustment = getAllStockAdjustmentRequestResult.Data. ([]model.StockAdjustment)
		
	equalStockAdjustment := cmp.Equal(createStockAdjustmentObj.ID, getAllStockAdjustmentObj[len(getAllStockAdjustmentObj)-1].ID)
		
	if equalStockAdjustment == false {
		t.Errorf( "Created object is not equal to the last entry in StockAdjustment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for StockAdjustment
	// --------------------------------------------------------------	
	deleteStockAdjustmentRequestResult := dao.DeleteStockAdjustment(uint64(createStockAdjustmentObj.ID))

	if deleteStockAdjustmentRequestResult.Success == false {
			t.Errorf(deleteStockAdjustmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion StockAdjustment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getStockAdjustmentRequestResult = dao.GetStockAdjustment( uint64(createStockAdjustmentObj.ID) )
	
	if getStockAdjustmentRequestResult.Success == true {
		t.Errorf(getStockAdjustmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestStockAdjustmentLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for StockAdjustmentLine
	//----------------------------------------------------------------------------
	StockAdjustmentLineObj := model.StockAdjustmentLine                                                                                                                    {LineNumber:100,Quantity:"test value",UnitOfMeasure:0,StockStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createStockAdjustmentLineRequestResult := dao.CreateStockAdjustmentLine( StockAdjustmentLineObj )
	
	if createStockAdjustmentLineRequestResult.Success == false {
		t.Errorf(createStockAdjustmentLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create StockAdjustmentLine success...")
	}
	
	createStockAdjustmentLineObj,_ := createStockAdjustmentLineRequestResult.Data. (model.StockAdjustmentLine)

	// --------------------------------------------------------------
	// Check StockAdjustmentLine Obj ID
	// --------------------------------------------------------------	
	if createStockAdjustmentLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for StockAdjustmentLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getStockAdjustmentLineRequestResult := dao.GetStockAdjustmentLine( uint64(createStockAdjustmentLineObj.ID) )
	
	if getStockAdjustmentLineRequestResult.Success == false {
		t.Errorf(getStockAdjustmentLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get StockAdjustmentLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getStockAdjustmentLineObj,_ := getStockAdjustmentLineRequestResult.Data. (model.StockAdjustmentLine)
	compareStockAdjustmentLine := cmp.Equal(createStockAdjustmentLineObj.ID, getStockAdjustmentLineObj.ID)
	
	if  compareStockAdjustmentLine == false	{
		t.Errorf( "Created StockAdjustmentLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllStockAdjustmentLineRequestResult := dao.GetAllStockAdjustmentLine()

	if getAllStockAdjustmentLineRequestResult.Success == false {
			t.Errorf(getAllStockAdjustmentLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll StockAdjustmentLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllStockAdjustmentLineObj []model.StockAdjustmentLine = getAllStockAdjustmentLineRequestResult.Data. ([]model.StockAdjustmentLine)
		
	equalStockAdjustmentLine := cmp.Equal(createStockAdjustmentLineObj.ID, getAllStockAdjustmentLineObj[len(getAllStockAdjustmentLineObj)-1].ID)
		
	if equalStockAdjustmentLine == false {
		t.Errorf( "Created object is not equal to the last entry in StockAdjustmentLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for StockAdjustmentLine
	// --------------------------------------------------------------	
	deleteStockAdjustmentLineRequestResult := dao.DeleteStockAdjustmentLine(uint64(createStockAdjustmentLineObj.ID))

	if deleteStockAdjustmentLineRequestResult.Success == false {
			t.Errorf(deleteStockAdjustmentLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion StockAdjustmentLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getStockAdjustmentLineRequestResult = dao.GetStockAdjustmentLine( uint64(createStockAdjustmentLineObj.ID) )
	
	if getStockAdjustmentLineRequestResult.Success == true {
		t.Errorf(getStockAdjustmentLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCycleCountCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CycleCount
	//----------------------------------------------------------------------------
	CycleCountObj := model.CycleCount                                                                                                                                                                                            {CountNumber:"test value for CountNumber",ScheduledDate:time.Now(),PerformedDate:time.Now(),ApprovedBy:"test value for ApprovedBy",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCycleCountRequestResult := dao.CreateCycleCount( CycleCountObj )
	
	if createCycleCountRequestResult.Success == false {
		t.Errorf(createCycleCountRequestResult.Msg)
	} else {
		fmt.Println("Check Create CycleCount success...")
	}
	
	createCycleCountObj,_ := createCycleCountRequestResult.Data. (model.CycleCount)

	// --------------------------------------------------------------
	// Check CycleCount Obj ID
	// --------------------------------------------------------------	
	if createCycleCountObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CycleCount" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCycleCountRequestResult := dao.GetCycleCount( uint64(createCycleCountObj.ID) )
	
	if getCycleCountRequestResult.Success == false {
		t.Errorf(getCycleCountRequestResult.Msg)
	} else {
		fmt.Println("Check Get CycleCount success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCycleCountObj,_ := getCycleCountRequestResult.Data. (model.CycleCount)
	compareCycleCount := cmp.Equal(createCycleCountObj.ID, getCycleCountObj.ID)
	
	if  compareCycleCount == false	{
		t.Errorf( "Created CycleCount object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCycleCountRequestResult := dao.GetAllCycleCount()

	if getAllCycleCountRequestResult.Success == false {
			t.Errorf(getAllCycleCountRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CycleCount success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCycleCountObj []model.CycleCount = getAllCycleCountRequestResult.Data. ([]model.CycleCount)
		
	equalCycleCount := cmp.Equal(createCycleCountObj.ID, getAllCycleCountObj[len(getAllCycleCountObj)-1].ID)
		
	if equalCycleCount == false {
		t.Errorf( "Created object is not equal to the last entry in CycleCount[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CycleCount
	// --------------------------------------------------------------	
	deleteCycleCountRequestResult := dao.DeleteCycleCount(uint64(createCycleCountObj.ID))

	if deleteCycleCountRequestResult.Success == false {
			t.Errorf(deleteCycleCountRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CycleCount success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCycleCountRequestResult = dao.GetCycleCount( uint64(createCycleCountObj.ID) )
	
	if getCycleCountRequestResult.Success == true {
		t.Errorf(getCycleCountRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCycleCountEntryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CycleCountEntry
	//----------------------------------------------------------------------------
	CycleCountEntryObj := model.CycleCountEntry                                                                                                                                                                                                                                                    {LineNumber:100,SystemQuantity:"test value",CountedQuantity:"test value",VarianceQuantity:"test value",RecountRequired:true,StockStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCycleCountEntryRequestResult := dao.CreateCycleCountEntry( CycleCountEntryObj )
	
	if createCycleCountEntryRequestResult.Success == false {
		t.Errorf(createCycleCountEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Create CycleCountEntry success...")
	}
	
	createCycleCountEntryObj,_ := createCycleCountEntryRequestResult.Data. (model.CycleCountEntry)

	// --------------------------------------------------------------
	// Check CycleCountEntry Obj ID
	// --------------------------------------------------------------	
	if createCycleCountEntryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CycleCountEntry" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCycleCountEntryRequestResult := dao.GetCycleCountEntry( uint64(createCycleCountEntryObj.ID) )
	
	if getCycleCountEntryRequestResult.Success == false {
		t.Errorf(getCycleCountEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Get CycleCountEntry success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCycleCountEntryObj,_ := getCycleCountEntryRequestResult.Data. (model.CycleCountEntry)
	compareCycleCountEntry := cmp.Equal(createCycleCountEntryObj.ID, getCycleCountEntryObj.ID)
	
	if  compareCycleCountEntry == false	{
		t.Errorf( "Created CycleCountEntry object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCycleCountEntryRequestResult := dao.GetAllCycleCountEntry()

	if getAllCycleCountEntryRequestResult.Success == false {
			t.Errorf(getAllCycleCountEntryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CycleCountEntry success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCycleCountEntryObj []model.CycleCountEntry = getAllCycleCountEntryRequestResult.Data. ([]model.CycleCountEntry)
		
	equalCycleCountEntry := cmp.Equal(createCycleCountEntryObj.ID, getAllCycleCountEntryObj[len(getAllCycleCountEntryObj)-1].ID)
		
	if equalCycleCountEntry == false {
		t.Errorf( "Created object is not equal to the last entry in CycleCountEntry[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CycleCountEntry
	// --------------------------------------------------------------	
	deleteCycleCountEntryRequestResult := dao.DeleteCycleCountEntry(uint64(createCycleCountEntryObj.ID))

	if deleteCycleCountEntryRequestResult.Success == false {
			t.Errorf(deleteCycleCountEntryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CycleCountEntry success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCycleCountEntryRequestResult = dao.GetCycleCountEntry( uint64(createCycleCountEntryObj.ID) )
	
	if getCycleCountEntryRequestResult.Success == true {
		t.Errorf(getCycleCountEntryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestReplenishmentPolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ReplenishmentPolicy
	//----------------------------------------------------------------------------
	ReplenishmentPolicyObj := model.ReplenishmentPolicy                                                                                                                                                                                                                                                                                                            {MinLevel:"test value",MaxLevel:"test value",ReorderPoint:"test value",ReorderQuantity:"test value",LeadTimeDays:100,ReviewPeriodDays:100,PolicyType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createReplenishmentPolicyRequestResult := dao.CreateReplenishmentPolicy( ReplenishmentPolicyObj )
	
	if createReplenishmentPolicyRequestResult.Success == false {
		t.Errorf(createReplenishmentPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Create ReplenishmentPolicy success...")
	}
	
	createReplenishmentPolicyObj,_ := createReplenishmentPolicyRequestResult.Data. (model.ReplenishmentPolicy)

	// --------------------------------------------------------------
	// Check ReplenishmentPolicy Obj ID
	// --------------------------------------------------------------	
	if createReplenishmentPolicyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ReplenishmentPolicy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getReplenishmentPolicyRequestResult := dao.GetReplenishmentPolicy( uint64(createReplenishmentPolicyObj.ID) )
	
	if getReplenishmentPolicyRequestResult.Success == false {
		t.Errorf(getReplenishmentPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Get ReplenishmentPolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getReplenishmentPolicyObj,_ := getReplenishmentPolicyRequestResult.Data. (model.ReplenishmentPolicy)
	compareReplenishmentPolicy := cmp.Equal(createReplenishmentPolicyObj.ID, getReplenishmentPolicyObj.ID)
	
	if  compareReplenishmentPolicy == false	{
		t.Errorf( "Created ReplenishmentPolicy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllReplenishmentPolicyRequestResult := dao.GetAllReplenishmentPolicy()

	if getAllReplenishmentPolicyRequestResult.Success == false {
			t.Errorf(getAllReplenishmentPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ReplenishmentPolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllReplenishmentPolicyObj []model.ReplenishmentPolicy = getAllReplenishmentPolicyRequestResult.Data. ([]model.ReplenishmentPolicy)
		
	equalReplenishmentPolicy := cmp.Equal(createReplenishmentPolicyObj.ID, getAllReplenishmentPolicyObj[len(getAllReplenishmentPolicyObj)-1].ID)
		
	if equalReplenishmentPolicy == false {
		t.Errorf( "Created object is not equal to the last entry in ReplenishmentPolicy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ReplenishmentPolicy
	// --------------------------------------------------------------	
	deleteReplenishmentPolicyRequestResult := dao.DeleteReplenishmentPolicy(uint64(createReplenishmentPolicyObj.ID))

	if deleteReplenishmentPolicyRequestResult.Success == false {
			t.Errorf(deleteReplenishmentPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ReplenishmentPolicy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getReplenishmentPolicyRequestResult = dao.GetReplenishmentPolicy( uint64(createReplenishmentPolicyObj.ID) )
	
	if getReplenishmentPolicyRequestResult.Success == true {
		t.Errorf(getReplenishmentPolicyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestUoMConversionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for UoMConversion
	//----------------------------------------------------------------------------
	UoMConversionObj := model.UoMConversion                                                                                                                    {Factor:"test value",Precision:100,FromUnit:0,ToUnit:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createUoMConversionRequestResult := dao.CreateUoMConversion( UoMConversionObj )
	
	if createUoMConversionRequestResult.Success == false {
		t.Errorf(createUoMConversionRequestResult.Msg)
	} else {
		fmt.Println("Check Create UoMConversion success...")
	}
	
	createUoMConversionObj,_ := createUoMConversionRequestResult.Data. (model.UoMConversion)

	// --------------------------------------------------------------
	// Check UoMConversion Obj ID
	// --------------------------------------------------------------	
	if createUoMConversionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for UoMConversion" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getUoMConversionRequestResult := dao.GetUoMConversion( uint64(createUoMConversionObj.ID) )
	
	if getUoMConversionRequestResult.Success == false {
		t.Errorf(getUoMConversionRequestResult.Msg)
	} else {
		fmt.Println("Check Get UoMConversion success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getUoMConversionObj,_ := getUoMConversionRequestResult.Data. (model.UoMConversion)
	compareUoMConversion := cmp.Equal(createUoMConversionObj.ID, getUoMConversionObj.ID)
	
	if  compareUoMConversion == false	{
		t.Errorf( "Created UoMConversion object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllUoMConversionRequestResult := dao.GetAllUoMConversion()

	if getAllUoMConversionRequestResult.Success == false {
			t.Errorf(getAllUoMConversionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll UoMConversion success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllUoMConversionObj []model.UoMConversion = getAllUoMConversionRequestResult.Data. ([]model.UoMConversion)
		
	equalUoMConversion := cmp.Equal(createUoMConversionObj.ID, getAllUoMConversionObj[len(getAllUoMConversionObj)-1].ID)
		
	if equalUoMConversion == false {
		t.Errorf( "Created object is not equal to the last entry in UoMConversion[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for UoMConversion
	// --------------------------------------------------------------	
	deleteUoMConversionRequestResult := dao.DeleteUoMConversion(uint64(createUoMConversionObj.ID))

	if deleteUoMConversionRequestResult.Success == false {
			t.Errorf(deleteUoMConversionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion UoMConversion success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getUoMConversionRequestResult = dao.GetUoMConversion( uint64(createUoMConversionObj.ID) )
	
	if getUoMConversionRequestResult.Success == true {
		t.Errorf(getUoMConversionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInventoryThresholdAlertCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InventoryThresholdAlert
	//----------------------------------------------------------------------------
	InventoryThresholdAlertObj := model.InventoryThresholdAlert                                                                                                                                                    {AlertNumber:"test value for AlertNumber",DetectedAt:time.Now(),Message:"test value for Message",AlertType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInventoryThresholdAlertRequestResult := dao.CreateInventoryThresholdAlert( InventoryThresholdAlertObj )
	
	if createInventoryThresholdAlertRequestResult.Success == false {
		t.Errorf(createInventoryThresholdAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Create InventoryThresholdAlert success...")
	}
	
	createInventoryThresholdAlertObj,_ := createInventoryThresholdAlertRequestResult.Data. (model.InventoryThresholdAlert)

	// --------------------------------------------------------------
	// Check InventoryThresholdAlert Obj ID
	// --------------------------------------------------------------	
	if createInventoryThresholdAlertObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InventoryThresholdAlert" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInventoryThresholdAlertRequestResult := dao.GetInventoryThresholdAlert( uint64(createInventoryThresholdAlertObj.ID) )
	
	if getInventoryThresholdAlertRequestResult.Success == false {
		t.Errorf(getInventoryThresholdAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Get InventoryThresholdAlert success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInventoryThresholdAlertObj,_ := getInventoryThresholdAlertRequestResult.Data. (model.InventoryThresholdAlert)
	compareInventoryThresholdAlert := cmp.Equal(createInventoryThresholdAlertObj.ID, getInventoryThresholdAlertObj.ID)
	
	if  compareInventoryThresholdAlert == false	{
		t.Errorf( "Created InventoryThresholdAlert object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInventoryThresholdAlertRequestResult := dao.GetAllInventoryThresholdAlert()

	if getAllInventoryThresholdAlertRequestResult.Success == false {
			t.Errorf(getAllInventoryThresholdAlertRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InventoryThresholdAlert success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInventoryThresholdAlertObj []model.InventoryThresholdAlert = getAllInventoryThresholdAlertRequestResult.Data. ([]model.InventoryThresholdAlert)
		
	equalInventoryThresholdAlert := cmp.Equal(createInventoryThresholdAlertObj.ID, getAllInventoryThresholdAlertObj[len(getAllInventoryThresholdAlertObj)-1].ID)
		
	if equalInventoryThresholdAlert == false {
		t.Errorf( "Created object is not equal to the last entry in InventoryThresholdAlert[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InventoryThresholdAlert
	// --------------------------------------------------------------	
	deleteInventoryThresholdAlertRequestResult := dao.DeleteInventoryThresholdAlert(uint64(createInventoryThresholdAlertObj.ID))

	if deleteInventoryThresholdAlertRequestResult.Success == false {
			t.Errorf(deleteInventoryThresholdAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InventoryThresholdAlert success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInventoryThresholdAlertRequestResult = dao.GetInventoryThresholdAlert( uint64(createInventoryThresholdAlertObj.ID) )
	
	if getInventoryThresholdAlertRequestResult.Success == true {
		t.Errorf(getInventoryThresholdAlertRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestQuarantineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Quarantine
	//----------------------------------------------------------------------------
	QuarantineObj := model.Quarantine                                                                                                                                                            {Reason:"test value for Reason",StartedAt:time.Now(),ReleasedAt:time.Now(),Disposition:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createQuarantineRequestResult := dao.CreateQuarantine( QuarantineObj )
	
	if createQuarantineRequestResult.Success == false {
		t.Errorf(createQuarantineRequestResult.Msg)
	} else {
		fmt.Println("Check Create Quarantine success...")
	}
	
	createQuarantineObj,_ := createQuarantineRequestResult.Data. (model.Quarantine)

	// --------------------------------------------------------------
	// Check Quarantine Obj ID
	// --------------------------------------------------------------	
	if createQuarantineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Quarantine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getQuarantineRequestResult := dao.GetQuarantine( uint64(createQuarantineObj.ID) )
	
	if getQuarantineRequestResult.Success == false {
		t.Errorf(getQuarantineRequestResult.Msg)
	} else {
		fmt.Println("Check Get Quarantine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getQuarantineObj,_ := getQuarantineRequestResult.Data. (model.Quarantine)
	compareQuarantine := cmp.Equal(createQuarantineObj.ID, getQuarantineObj.ID)
	
	if  compareQuarantine == false	{
		t.Errorf( "Created Quarantine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllQuarantineRequestResult := dao.GetAllQuarantine()

	if getAllQuarantineRequestResult.Success == false {
			t.Errorf(getAllQuarantineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Quarantine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllQuarantineObj []model.Quarantine = getAllQuarantineRequestResult.Data. ([]model.Quarantine)
		
	equalQuarantine := cmp.Equal(createQuarantineObj.ID, getAllQuarantineObj[len(getAllQuarantineObj)-1].ID)
		
	if equalQuarantine == false {
		t.Errorf( "Created object is not equal to the last entry in Quarantine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Quarantine
	// --------------------------------------------------------------	
	deleteQuarantineRequestResult := dao.DeleteQuarantine(uint64(createQuarantineObj.ID))

	if deleteQuarantineRequestResult.Success == false {
			t.Errorf(deleteQuarantineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Quarantine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getQuarantineRequestResult = dao.GetQuarantine( uint64(createQuarantineObj.ID) )
	
	if getQuarantineRequestResult.Success == true {
		t.Errorf(getQuarantineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestExpirationPolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ExpirationPolicy
	//----------------------------------------------------------------------------
	ExpirationPolicyObj := model.ExpirationPolicy                                                                            {RejectIfDaysToExpireLessThan:100,AutoQuarantineDaysToExpire:100,RotationMethod:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createExpirationPolicyRequestResult := dao.CreateExpirationPolicy( ExpirationPolicyObj )
	
	if createExpirationPolicyRequestResult.Success == false {
		t.Errorf(createExpirationPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Create ExpirationPolicy success...")
	}
	
	createExpirationPolicyObj,_ := createExpirationPolicyRequestResult.Data. (model.ExpirationPolicy)

	// --------------------------------------------------------------
	// Check ExpirationPolicy Obj ID
	// --------------------------------------------------------------	
	if createExpirationPolicyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ExpirationPolicy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getExpirationPolicyRequestResult := dao.GetExpirationPolicy( uint64(createExpirationPolicyObj.ID) )
	
	if getExpirationPolicyRequestResult.Success == false {
		t.Errorf(getExpirationPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Get ExpirationPolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getExpirationPolicyObj,_ := getExpirationPolicyRequestResult.Data. (model.ExpirationPolicy)
	compareExpirationPolicy := cmp.Equal(createExpirationPolicyObj.ID, getExpirationPolicyObj.ID)
	
	if  compareExpirationPolicy == false	{
		t.Errorf( "Created ExpirationPolicy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllExpirationPolicyRequestResult := dao.GetAllExpirationPolicy()

	if getAllExpirationPolicyRequestResult.Success == false {
			t.Errorf(getAllExpirationPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ExpirationPolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllExpirationPolicyObj []model.ExpirationPolicy = getAllExpirationPolicyRequestResult.Data. ([]model.ExpirationPolicy)
		
	equalExpirationPolicy := cmp.Equal(createExpirationPolicyObj.ID, getAllExpirationPolicyObj[len(getAllExpirationPolicyObj)-1].ID)
		
	if equalExpirationPolicy == false {
		t.Errorf( "Created object is not equal to the last entry in ExpirationPolicy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ExpirationPolicy
	// --------------------------------------------------------------	
	deleteExpirationPolicyRequestResult := dao.DeleteExpirationPolicy(uint64(createExpirationPolicyObj.ID))

	if deleteExpirationPolicyRequestResult.Success == false {
			t.Errorf(deleteExpirationPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ExpirationPolicy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getExpirationPolicyRequestResult = dao.GetExpirationPolicy( uint64(createExpirationPolicyObj.ID) )
	
	if getExpirationPolicyRequestResult.Success == true {
		t.Errorf(getExpirationPolicyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInboundShipmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InboundShipment
	//----------------------------------------------------------------------------
	InboundShipmentObj := model.InboundShipment                                                                                                                                                                                            {ShipmentNumber:"test value for ShipmentNumber",ExpectedArrivalDate:time.Now(),ArrivalDate:time.Now(),CarrierName:"test value for CarrierName",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInboundShipmentRequestResult := dao.CreateInboundShipment( InboundShipmentObj )
	
	if createInboundShipmentRequestResult.Success == false {
		t.Errorf(createInboundShipmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create InboundShipment success...")
	}
	
	createInboundShipmentObj,_ := createInboundShipmentRequestResult.Data. (model.InboundShipment)

	// --------------------------------------------------------------
	// Check InboundShipment Obj ID
	// --------------------------------------------------------------	
	if createInboundShipmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InboundShipment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInboundShipmentRequestResult := dao.GetInboundShipment( uint64(createInboundShipmentObj.ID) )
	
	if getInboundShipmentRequestResult.Success == false {
		t.Errorf(getInboundShipmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get InboundShipment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInboundShipmentObj,_ := getInboundShipmentRequestResult.Data. (model.InboundShipment)
	compareInboundShipment := cmp.Equal(createInboundShipmentObj.ID, getInboundShipmentObj.ID)
	
	if  compareInboundShipment == false	{
		t.Errorf( "Created InboundShipment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInboundShipmentRequestResult := dao.GetAllInboundShipment()

	if getAllInboundShipmentRequestResult.Success == false {
			t.Errorf(getAllInboundShipmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InboundShipment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInboundShipmentObj []model.InboundShipment = getAllInboundShipmentRequestResult.Data. ([]model.InboundShipment)
		
	equalInboundShipment := cmp.Equal(createInboundShipmentObj.ID, getAllInboundShipmentObj[len(getAllInboundShipmentObj)-1].ID)
		
	if equalInboundShipment == false {
		t.Errorf( "Created object is not equal to the last entry in InboundShipment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InboundShipment
	// --------------------------------------------------------------	
	deleteInboundShipmentRequestResult := dao.DeleteInboundShipment(uint64(createInboundShipmentObj.ID))

	if deleteInboundShipmentRequestResult.Success == false {
			t.Errorf(deleteInboundShipmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InboundShipment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInboundShipmentRequestResult = dao.GetInboundShipment( uint64(createInboundShipmentObj.ID) )
	
	if getInboundShipmentRequestResult.Success == true {
		t.Errorf(getInboundShipmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInboundShipmentLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InboundShipmentLine
	//----------------------------------------------------------------------------
	InboundShipmentLineObj := model.InboundShipmentLine                                                                                                                    {LineNumber:100,Quantity:"test value",UnitOfMeasure:0,StockStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInboundShipmentLineRequestResult := dao.CreateInboundShipmentLine( InboundShipmentLineObj )
	
	if createInboundShipmentLineRequestResult.Success == false {
		t.Errorf(createInboundShipmentLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create InboundShipmentLine success...")
	}
	
	createInboundShipmentLineObj,_ := createInboundShipmentLineRequestResult.Data. (model.InboundShipmentLine)

	// --------------------------------------------------------------
	// Check InboundShipmentLine Obj ID
	// --------------------------------------------------------------	
	if createInboundShipmentLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InboundShipmentLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInboundShipmentLineRequestResult := dao.GetInboundShipmentLine( uint64(createInboundShipmentLineObj.ID) )
	
	if getInboundShipmentLineRequestResult.Success == false {
		t.Errorf(getInboundShipmentLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get InboundShipmentLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInboundShipmentLineObj,_ := getInboundShipmentLineRequestResult.Data. (model.InboundShipmentLine)
	compareInboundShipmentLine := cmp.Equal(createInboundShipmentLineObj.ID, getInboundShipmentLineObj.ID)
	
	if  compareInboundShipmentLine == false	{
		t.Errorf( "Created InboundShipmentLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInboundShipmentLineRequestResult := dao.GetAllInboundShipmentLine()

	if getAllInboundShipmentLineRequestResult.Success == false {
			t.Errorf(getAllInboundShipmentLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InboundShipmentLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInboundShipmentLineObj []model.InboundShipmentLine = getAllInboundShipmentLineRequestResult.Data. ([]model.InboundShipmentLine)
		
	equalInboundShipmentLine := cmp.Equal(createInboundShipmentLineObj.ID, getAllInboundShipmentLineObj[len(getAllInboundShipmentLineObj)-1].ID)
		
	if equalInboundShipmentLine == false {
		t.Errorf( "Created object is not equal to the last entry in InboundShipmentLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InboundShipmentLine
	// --------------------------------------------------------------	
	deleteInboundShipmentLineRequestResult := dao.DeleteInboundShipmentLine(uint64(createInboundShipmentLineObj.ID))

	if deleteInboundShipmentLineRequestResult.Success == false {
			t.Errorf(deleteInboundShipmentLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InboundShipmentLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInboundShipmentLineRequestResult = dao.GetInboundShipmentLine( uint64(createInboundShipmentLineObj.ID) )
	
	if getInboundShipmentLineRequestResult.Success == true {
		t.Errorf(getInboundShipmentLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOutboundAllocationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for OutboundAllocation
	//----------------------------------------------------------------------------
	OutboundAllocationObj := model.OutboundAllocation                                                                                                                                                            {AllocationNumber:"test value for AllocationNumber",AllocatedQuantity:"test value",AllocationDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOutboundAllocationRequestResult := dao.CreateOutboundAllocation( OutboundAllocationObj )
	
	if createOutboundAllocationRequestResult.Success == false {
		t.Errorf(createOutboundAllocationRequestResult.Msg)
	} else {
		fmt.Println("Check Create OutboundAllocation success...")
	}
	
	createOutboundAllocationObj,_ := createOutboundAllocationRequestResult.Data. (model.OutboundAllocation)

	// --------------------------------------------------------------
	// Check OutboundAllocation Obj ID
	// --------------------------------------------------------------	
	if createOutboundAllocationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for OutboundAllocation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOutboundAllocationRequestResult := dao.GetOutboundAllocation( uint64(createOutboundAllocationObj.ID) )
	
	if getOutboundAllocationRequestResult.Success == false {
		t.Errorf(getOutboundAllocationRequestResult.Msg)
	} else {
		fmt.Println("Check Get OutboundAllocation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOutboundAllocationObj,_ := getOutboundAllocationRequestResult.Data. (model.OutboundAllocation)
	compareOutboundAllocation := cmp.Equal(createOutboundAllocationObj.ID, getOutboundAllocationObj.ID)
	
	if  compareOutboundAllocation == false	{
		t.Errorf( "Created OutboundAllocation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOutboundAllocationRequestResult := dao.GetAllOutboundAllocation()

	if getAllOutboundAllocationRequestResult.Success == false {
			t.Errorf(getAllOutboundAllocationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll OutboundAllocation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOutboundAllocationObj []model.OutboundAllocation = getAllOutboundAllocationRequestResult.Data. ([]model.OutboundAllocation)
		
	equalOutboundAllocation := cmp.Equal(createOutboundAllocationObj.ID, getAllOutboundAllocationObj[len(getAllOutboundAllocationObj)-1].ID)
		
	if equalOutboundAllocation == false {
		t.Errorf( "Created object is not equal to the last entry in OutboundAllocation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for OutboundAllocation
	// --------------------------------------------------------------	
	deleteOutboundAllocationRequestResult := dao.DeleteOutboundAllocation(uint64(createOutboundAllocationObj.ID))

	if deleteOutboundAllocationRequestResult.Success == false {
			t.Errorf(deleteOutboundAllocationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion OutboundAllocation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOutboundAllocationRequestResult = dao.GetOutboundAllocation( uint64(createOutboundAllocationObj.ID) )
	
	if getOutboundAllocationRequestResult.Success == true {
		t.Errorf(getOutboundAllocationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

