package test

import ( 
	"testing"
    dao "manufacturing-on-golang/internal/dao"
	"manufacturing-on-golang/internal/model"
	"manufacturing-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestEnterpriseCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Enterprise
	//----------------------------------------------------------------------------
	EnterpriseObj := model.Enterprise                                                                                                                                                            {Name:"test value for Name",LegalName:"test value for LegalName",RegistrationCountry:"test value for RegistrationCountry",Website:"test value for Website",TaxId:"test value for TaxId"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEnterpriseRequestResult := dao.CreateEnterprise( EnterpriseObj )
	
	if createEnterpriseRequestResult.Success == false {
		t.Errorf(createEnterpriseRequestResult.Msg)
	} else {
		fmt.Println("Check Create Enterprise success...")
	}
	
	createEnterpriseObj,_ := createEnterpriseRequestResult.Data. (model.Enterprise)

	// --------------------------------------------------------------
	// Check Enterprise Obj ID
	// --------------------------------------------------------------	
	if createEnterpriseObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Enterprise" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEnterpriseRequestResult := dao.GetEnterprise( uint64(createEnterpriseObj.ID) )
	
	if getEnterpriseRequestResult.Success == false {
		t.Errorf(getEnterpriseRequestResult.Msg)
	} else {
		fmt.Println("Check Get Enterprise success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEnterpriseObj,_ := getEnterpriseRequestResult.Data. (model.Enterprise)
	compareEnterprise := cmp.Equal(createEnterpriseObj.ID, getEnterpriseObj.ID)
	
	if  compareEnterprise == false	{
		t.Errorf( "Created Enterprise object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEnterpriseRequestResult := dao.GetAllEnterprise()

	if getAllEnterpriseRequestResult.Success == false {
			t.Errorf(getAllEnterpriseRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Enterprise success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEnterpriseObj []model.Enterprise = getAllEnterpriseRequestResult.Data. ([]model.Enterprise)
		
	equalEnterprise := cmp.Equal(createEnterpriseObj.ID, getAllEnterpriseObj[len(getAllEnterpriseObj)-1].ID)
		
	if equalEnterprise == false {
		t.Errorf( "Created object is not equal to the last entry in Enterprise[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Enterprise
	// --------------------------------------------------------------	
	deleteEnterpriseRequestResult := dao.DeleteEnterprise(uint64(createEnterpriseObj.ID))

	if deleteEnterpriseRequestResult.Success == false {
			t.Errorf(deleteEnterpriseRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Enterprise success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEnterpriseRequestResult = dao.GetEnterprise( uint64(createEnterpriseObj.ID) )
	
	if getEnterpriseRequestResult.Success == true {
		t.Errorf(getEnterpriseRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBusinessUnitCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BusinessUnit
	//----------------------------------------------------------------------------
	BusinessUnitObj := model.BusinessUnit                                                                            {Name:"test value for Name",Code:"test value for Code",Category:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBusinessUnitRequestResult := dao.CreateBusinessUnit( BusinessUnitObj )
	
	if createBusinessUnitRequestResult.Success == false {
		t.Errorf(createBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Create BusinessUnit success...")
	}
	
	createBusinessUnitObj,_ := createBusinessUnitRequestResult.Data. (model.BusinessUnit)

	// --------------------------------------------------------------
	// Check BusinessUnit Obj ID
	// --------------------------------------------------------------	
	if createBusinessUnitObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BusinessUnit" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBusinessUnitRequestResult := dao.GetBusinessUnit( uint64(createBusinessUnitObj.ID) )
	
	if getBusinessUnitRequestResult.Success == false {
		t.Errorf(getBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Get BusinessUnit success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBusinessUnitObj,_ := getBusinessUnitRequestResult.Data. (model.BusinessUnit)
	compareBusinessUnit := cmp.Equal(createBusinessUnitObj.ID, getBusinessUnitObj.ID)
	
	if  compareBusinessUnit == false	{
		t.Errorf( "Created BusinessUnit object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBusinessUnitRequestResult := dao.GetAllBusinessUnit()

	if getAllBusinessUnitRequestResult.Success == false {
			t.Errorf(getAllBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BusinessUnit success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBusinessUnitObj []model.BusinessUnit = getAllBusinessUnitRequestResult.Data. ([]model.BusinessUnit)
		
	equalBusinessUnit := cmp.Equal(createBusinessUnitObj.ID, getAllBusinessUnitObj[len(getAllBusinessUnitObj)-1].ID)
		
	if equalBusinessUnit == false {
		t.Errorf( "Created object is not equal to the last entry in BusinessUnit[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BusinessUnit
	// --------------------------------------------------------------	
	deleteBusinessUnitRequestResult := dao.DeleteBusinessUnit(uint64(createBusinessUnitObj.ID))

	if deleteBusinessUnitRequestResult.Success == false {
			t.Errorf(deleteBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BusinessUnit success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBusinessUnitRequestResult = dao.GetBusinessUnit( uint64(createBusinessUnitObj.ID) )
	
	if getBusinessUnitRequestResult.Success == true {
		t.Errorf(getBusinessUnitRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPlantCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Plant
	//----------------------------------------------------------------------------
	PlantObj := model.Plant                                                                                                            {Name:"test value for Name",PlantCode:"test value for PlantCode",Address:new Address(),TimeZone:"test value for TimeZone"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPlantRequestResult := dao.CreatePlant( PlantObj )
	
	if createPlantRequestResult.Success == false {
		t.Errorf(createPlantRequestResult.Msg)
	} else {
		fmt.Println("Check Create Plant success...")
	}
	
	createPlantObj,_ := createPlantRequestResult.Data. (model.Plant)

	// --------------------------------------------------------------
	// Check Plant Obj ID
	// --------------------------------------------------------------	
	if createPlantObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Plant" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPlantRequestResult := dao.GetPlant( uint64(createPlantObj.ID) )
	
	if getPlantRequestResult.Success == false {
		t.Errorf(getPlantRequestResult.Msg)
	} else {
		fmt.Println("Check Get Plant success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPlantObj,_ := getPlantRequestResult.Data. (model.Plant)
	comparePlant := cmp.Equal(createPlantObj.ID, getPlantObj.ID)
	
	if  comparePlant == false	{
		t.Errorf( "Created Plant object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPlantRequestResult := dao.GetAllPlant()

	if getAllPlantRequestResult.Success == false {
			t.Errorf(getAllPlantRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Plant success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPlantObj []model.Plant = getAllPlantRequestResult.Data. ([]model.Plant)
		
	equalPlant := cmp.Equal(createPlantObj.ID, getAllPlantObj[len(getAllPlantObj)-1].ID)
		
	if equalPlant == false {
		t.Errorf( "Created object is not equal to the last entry in Plant[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Plant
	// --------------------------------------------------------------	
	deletePlantRequestResult := dao.DeletePlant(uint64(createPlantObj.ID))

	if deletePlantRequestResult.Success == false {
			t.Errorf(deletePlantRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Plant success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPlantRequestResult = dao.GetPlant( uint64(createPlantObj.ID) )
	
	if getPlantRequestResult.Success == true {
		t.Errorf(getPlantRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProductionLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ProductionLine
	//----------------------------------------------------------------------------
	ProductionLineObj := model.ProductionLine                                                                            {Name:"test value for Name",LineCode:"test value for LineCode",LineType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProductionLineRequestResult := dao.CreateProductionLine( ProductionLineObj )
	
	if createProductionLineRequestResult.Success == false {
		t.Errorf(createProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create ProductionLine success...")
	}
	
	createProductionLineObj,_ := createProductionLineRequestResult.Data. (model.ProductionLine)

	// --------------------------------------------------------------
	// Check ProductionLine Obj ID
	// --------------------------------------------------------------	
	if createProductionLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ProductionLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProductionLineRequestResult := dao.GetProductionLine( uint64(createProductionLineObj.ID) )
	
	if getProductionLineRequestResult.Success == false {
		t.Errorf(getProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get ProductionLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProductionLineObj,_ := getProductionLineRequestResult.Data. (model.ProductionLine)
	compareProductionLine := cmp.Equal(createProductionLineObj.ID, getProductionLineObj.ID)
	
	if  compareProductionLine == false	{
		t.Errorf( "Created ProductionLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProductionLineRequestResult := dao.GetAllProductionLine()

	if getAllProductionLineRequestResult.Success == false {
			t.Errorf(getAllProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ProductionLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProductionLineObj []model.ProductionLine = getAllProductionLineRequestResult.Data. ([]model.ProductionLine)
		
	equalProductionLine := cmp.Equal(createProductionLineObj.ID, getAllProductionLineObj[len(getAllProductionLineObj)-1].ID)
		
	if equalProductionLine == false {
		t.Errorf( "Created object is not equal to the last entry in ProductionLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ProductionLine
	// --------------------------------------------------------------	
	deleteProductionLineRequestResult := dao.DeleteProductionLine(uint64(createProductionLineObj.ID))

	if deleteProductionLineRequestResult.Success == false {
			t.Errorf(deleteProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ProductionLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProductionLineRequestResult = dao.GetProductionLine( uint64(createProductionLineObj.ID) )
	
	if getProductionLineRequestResult.Success == true {
		t.Errorf(getProductionLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWorkCenterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for WorkCenter
	//----------------------------------------------------------------------------
	WorkCenterObj := model.WorkCenter                                                                                                                            {Name:"test value for Name",Code:"test value for Code",CapacityPerHour:100,OeeTarget:new Percentage(),WorkCenterType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWorkCenterRequestResult := dao.CreateWorkCenter( WorkCenterObj )
	
	if createWorkCenterRequestResult.Success == false {
		t.Errorf(createWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Create WorkCenter success...")
	}
	
	createWorkCenterObj,_ := createWorkCenterRequestResult.Data. (model.WorkCenter)

	// --------------------------------------------------------------
	// Check WorkCenter Obj ID
	// --------------------------------------------------------------	
	if createWorkCenterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for WorkCenter" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWorkCenterRequestResult := dao.GetWorkCenter( uint64(createWorkCenterObj.ID) )
	
	if getWorkCenterRequestResult.Success == false {
		t.Errorf(getWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Get WorkCenter success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWorkCenterObj,_ := getWorkCenterRequestResult.Data. (model.WorkCenter)
	compareWorkCenter := cmp.Equal(createWorkCenterObj.ID, getWorkCenterObj.ID)
	
	if  compareWorkCenter == false	{
		t.Errorf( "Created WorkCenter object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWorkCenterRequestResult := dao.GetAllWorkCenter()

	if getAllWorkCenterRequestResult.Success == false {
			t.Errorf(getAllWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll WorkCenter success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWorkCenterObj []model.WorkCenter = getAllWorkCenterRequestResult.Data. ([]model.WorkCenter)
		
	equalWorkCenter := cmp.Equal(createWorkCenterObj.ID, getAllWorkCenterObj[len(getAllWorkCenterObj)-1].ID)
		
	if equalWorkCenter == false {
		t.Errorf( "Created object is not equal to the last entry in WorkCenter[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for WorkCenter
	// --------------------------------------------------------------	
	deleteWorkCenterRequestResult := dao.DeleteWorkCenter(uint64(createWorkCenterObj.ID))

	if deleteWorkCenterRequestResult.Success == false {
			t.Errorf(deleteWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion WorkCenter success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWorkCenterRequestResult = dao.GetWorkCenter( uint64(createWorkCenterObj.ID) )
	
	if getWorkCenterRequestResult.Success == true {
		t.Errorf(getWorkCenterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Item
	//----------------------------------------------------------------------------
	ItemObj := model.Item                                                                                                                                                                                            {ItemNumber:"test value for ItemNumber",Name:"test value for Name",StandardCost:new Money(),Weight:new Measurement(),AsSerialControlled:true,ItemType:0,ProcurementType:0,UnitOfMeasure:0,LifecycleStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createItemRequestResult := dao.CreateItem( ItemObj )
	
	if createItemRequestResult.Success == false {
		t.Errorf(createItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create Item success...")
	}
	
	createItemObj,_ := createItemRequestResult.Data. (model.Item)

	// --------------------------------------------------------------
	// Check Item Obj ID
	// --------------------------------------------------------------	
	if createItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Item" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getItemRequestResult := dao.GetItem( uint64(createItemObj.ID) )
	
	if getItemRequestResult.Success == false {
		t.Errorf(getItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get Item success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getItemObj,_ := getItemRequestResult.Data. (model.Item)
	compareItem := cmp.Equal(createItemObj.ID, getItemObj.ID)
	
	if  compareItem == false	{
		t.Errorf( "Created Item object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllItemRequestResult := dao.GetAllItem()

	if getAllItemRequestResult.Success == false {
			t.Errorf(getAllItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Item success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllItemObj []model.Item = getAllItemRequestResult.Data. ([]model.Item)
		
	equalItem := cmp.Equal(createItemObj.ID, getAllItemObj[len(getAllItemObj)-1].ID)
		
	if equalItem == false {
		t.Errorf( "Created object is not equal to the last entry in Item[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Item
	// --------------------------------------------------------------	
	deleteItemRequestResult := dao.DeleteItem(uint64(createItemObj.ID))

	if deleteItemRequestResult.Success == false {
			t.Errorf(deleteItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Item success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getItemRequestResult = dao.GetItem( uint64(createItemObj.ID) )
	
	if getItemRequestResult.Success == true {
		t.Errorf(getItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBOMCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BOM
	//----------------------------------------------------------------------------
	BOMObj := model.BOM                                                                                                                                                                                            {BomNumber:"test value for BomNumber",Revision:"test value for Revision",EffectivityStart:time.Now(),EffectivityEnd:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBOMRequestResult := dao.CreateBOM( BOMObj )
	
	if createBOMRequestResult.Success == false {
		t.Errorf(createBOMRequestResult.Msg)
	} else {
		fmt.Println("Check Create BOM success...")
	}
	
	createBOMObj,_ := createBOMRequestResult.Data. (model.BOM)

	// --------------------------------------------------------------
	// Check BOM Obj ID
	// --------------------------------------------------------------	
	if createBOMObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BOM" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBOMRequestResult := dao.GetBOM( uint64(createBOMObj.ID) )
	
	if getBOMRequestResult.Success == false {
		t.Errorf(getBOMRequestResult.Msg)
	} else {
		fmt.Println("Check Get BOM success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBOMObj,_ := getBOMRequestResult.Data. (model.BOM)
	compareBOM := cmp.Equal(createBOMObj.ID, getBOMObj.ID)
	
	if  compareBOM == false	{
		t.Errorf( "Created BOM object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBOMRequestResult := dao.GetAllBOM()

	if getAllBOMRequestResult.Success == false {
			t.Errorf(getAllBOMRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BOM success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBOMObj []model.BOM = getAllBOMRequestResult.Data. ([]model.BOM)
		
	equalBOM := cmp.Equal(createBOMObj.ID, getAllBOMObj[len(getAllBOMObj)-1].ID)
		
	if equalBOM == false {
		t.Errorf( "Created object is not equal to the last entry in BOM[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BOM
	// --------------------------------------------------------------	
	deleteBOMRequestResult := dao.DeleteBOM(uint64(createBOMObj.ID))

	if deleteBOMRequestResult.Success == false {
			t.Errorf(deleteBOMRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BOM success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBOMRequestResult = dao.GetBOM( uint64(createBOMObj.ID) )
	
	if getBOMRequestResult.Success == true {
		t.Errorf(getBOMRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBOMItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BOMItem
	//----------------------------------------------------------------------------
	BOMItemObj := model.BOMItem                                                            {LineNumber:100,Quantity:new Quantity(),ScrapPercent:new Percentage()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBOMItemRequestResult := dao.CreateBOMItem( BOMItemObj )
	
	if createBOMItemRequestResult.Success == false {
		t.Errorf(createBOMItemRequestResult.Msg)
	} else {
		fmt.Println("Check Create BOMItem success...")
	}
	
	createBOMItemObj,_ := createBOMItemRequestResult.Data. (model.BOMItem)

	// --------------------------------------------------------------
	// Check BOMItem Obj ID
	// --------------------------------------------------------------	
	if createBOMItemObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BOMItem" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBOMItemRequestResult := dao.GetBOMItem( uint64(createBOMItemObj.ID) )
	
	if getBOMItemRequestResult.Success == false {
		t.Errorf(getBOMItemRequestResult.Msg)
	} else {
		fmt.Println("Check Get BOMItem success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBOMItemObj,_ := getBOMItemRequestResult.Data. (model.BOMItem)
	compareBOMItem := cmp.Equal(createBOMItemObj.ID, getBOMItemObj.ID)
	
	if  compareBOMItem == false	{
		t.Errorf( "Created BOMItem object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBOMItemRequestResult := dao.GetAllBOMItem()

	if getAllBOMItemRequestResult.Success == false {
			t.Errorf(getAllBOMItemRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BOMItem success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBOMItemObj []model.BOMItem = getAllBOMItemRequestResult.Data. ([]model.BOMItem)
		
	equalBOMItem := cmp.Equal(createBOMItemObj.ID, getAllBOMItemObj[len(getAllBOMItemObj)-1].ID)
		
	if equalBOMItem == false {
		t.Errorf( "Created object is not equal to the last entry in BOMItem[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BOMItem
	// --------------------------------------------------------------	
	deleteBOMItemRequestResult := dao.DeleteBOMItem(uint64(createBOMItemObj.ID))

	if deleteBOMItemRequestResult.Success == false {
			t.Errorf(deleteBOMItemRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BOMItem success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBOMItemRequestResult = dao.GetBOMItem( uint64(createBOMItemObj.ID) )
	
	if getBOMItemRequestResult.Success == true {
		t.Errorf(getBOMItemRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRoutingCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Routing
	//----------------------------------------------------------------------------
	RoutingObj := model.Routing                                                                                                                                                                                                            {RoutingNumber:"test value for RoutingNumber",Revision:"test value for Revision",EffectivityStart:time.Now(),EffectivityEnd:time.Now(),RoutingType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRoutingRequestResult := dao.CreateRouting( RoutingObj )
	
	if createRoutingRequestResult.Success == false {
		t.Errorf(createRoutingRequestResult.Msg)
	} else {
		fmt.Println("Check Create Routing success...")
	}
	
	createRoutingObj,_ := createRoutingRequestResult.Data. (model.Routing)

	// --------------------------------------------------------------
	// Check Routing Obj ID
	// --------------------------------------------------------------	
	if createRoutingObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Routing" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRoutingRequestResult := dao.GetRouting( uint64(createRoutingObj.ID) )
	
	if getRoutingRequestResult.Success == false {
		t.Errorf(getRoutingRequestResult.Msg)
	} else {
		fmt.Println("Check Get Routing success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRoutingObj,_ := getRoutingRequestResult.Data. (model.Routing)
	compareRouting := cmp.Equal(createRoutingObj.ID, getRoutingObj.ID)
	
	if  compareRouting == false	{
		t.Errorf( "Created Routing object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRoutingRequestResult := dao.GetAllRouting()

	if getAllRoutingRequestResult.Success == false {
			t.Errorf(getAllRoutingRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Routing success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRoutingObj []model.Routing = getAllRoutingRequestResult.Data. ([]model.Routing)
		
	equalRouting := cmp.Equal(createRoutingObj.ID, getAllRoutingObj[len(getAllRoutingObj)-1].ID)
		
	if equalRouting == false {
		t.Errorf( "Created object is not equal to the last entry in Routing[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Routing
	// --------------------------------------------------------------	
	deleteRoutingRequestResult := dao.DeleteRouting(uint64(createRoutingObj.ID))

	if deleteRoutingRequestResult.Success == false {
			t.Errorf(deleteRoutingRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Routing success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRoutingRequestResult = dao.GetRouting( uint64(createRoutingObj.ID) )
	
	if getRoutingRequestResult.Success == true {
		t.Errorf(getRoutingRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestOperationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Operation
	//----------------------------------------------------------------------------
	OperationObj := model.Operation                                                                                                            {OperationNumber:"test value for OperationNumber",Name:"test value for Name",SetupTime:new TimeDuration(),StandardCycleTime:new TimeDuration(),OperationType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createOperationRequestResult := dao.CreateOperation( OperationObj )
	
	if createOperationRequestResult.Success == false {
		t.Errorf(createOperationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Operation success...")
	}
	
	createOperationObj,_ := createOperationRequestResult.Data. (model.Operation)

	// --------------------------------------------------------------
	// Check Operation Obj ID
	// --------------------------------------------------------------	
	if createOperationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Operation" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getOperationRequestResult := dao.GetOperation( uint64(createOperationObj.ID) )
	
	if getOperationRequestResult.Success == false {
		t.Errorf(getOperationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Operation success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getOperationObj,_ := getOperationRequestResult.Data. (model.Operation)
	compareOperation := cmp.Equal(createOperationObj.ID, getOperationObj.ID)
	
	if  compareOperation == false	{
		t.Errorf( "Created Operation object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllOperationRequestResult := dao.GetAllOperation()

	if getAllOperationRequestResult.Success == false {
			t.Errorf(getAllOperationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Operation success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllOperationObj []model.Operation = getAllOperationRequestResult.Data. ([]model.Operation)
		
	equalOperation := cmp.Equal(createOperationObj.ID, getAllOperationObj[len(getAllOperationObj)-1].ID)
		
	if equalOperation == false {
		t.Errorf( "Created object is not equal to the last entry in Operation[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Operation
	// --------------------------------------------------------------	
	deleteOperationRequestResult := dao.DeleteOperation(uint64(createOperationObj.ID))

	if deleteOperationRequestResult.Success == false {
			t.Errorf(deleteOperationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Operation success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getOperationRequestResult = dao.GetOperation( uint64(createOperationObj.ID) )
	
	if getOperationRequestResult.Success == true {
		t.Errorf(getOperationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWorkOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for WorkOrder
	//----------------------------------------------------------------------------
	WorkOrderObj := model.WorkOrder                                                                                                                                                                                                            {WorkOrderNumber:"test value for WorkOrderNumber",PlannedStart:time.Now(),PlannedEnd:time.Now(),Quantity:new Quantity(),Priority:100,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createWorkOrderRequestResult := dao.CreateWorkOrder( WorkOrderObj )
	
	if createWorkOrderRequestResult.Success == false {
		t.Errorf(createWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create WorkOrder success...")
	}
	
	createWorkOrderObj,_ := createWorkOrderRequestResult.Data. (model.WorkOrder)

	// --------------------------------------------------------------
	// Check WorkOrder Obj ID
	// --------------------------------------------------------------	
	if createWorkOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for WorkOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getWorkOrderRequestResult := dao.GetWorkOrder( uint64(createWorkOrderObj.ID) )
	
	if getWorkOrderRequestResult.Success == false {
		t.Errorf(getWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get WorkOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getWorkOrderObj,_ := getWorkOrderRequestResult.Data. (model.WorkOrder)
	compareWorkOrder := cmp.Equal(createWorkOrderObj.ID, getWorkOrderObj.ID)
	
	if  compareWorkOrder == false	{
		t.Errorf( "Created WorkOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllWorkOrderRequestResult := dao.GetAllWorkOrder()

	if getAllWorkOrderRequestResult.Success == false {
			t.Errorf(getAllWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll WorkOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllWorkOrderObj []model.WorkOrder = getAllWorkOrderRequestResult.Data. ([]model.WorkOrder)
		
	equalWorkOrder := cmp.Equal(createWorkOrderObj.ID, getAllWorkOrderObj[len(getAllWorkOrderObj)-1].ID)
		
	if equalWorkOrder == false {
		t.Errorf( "Created object is not equal to the last entry in WorkOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for WorkOrder
	// --------------------------------------------------------------	
	deleteWorkOrderRequestResult := dao.DeleteWorkOrder(uint64(createWorkOrderObj.ID))

	if deleteWorkOrderRequestResult.Success == false {
			t.Errorf(deleteWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion WorkOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getWorkOrderRequestResult = dao.GetWorkOrder( uint64(createWorkOrderObj.ID) )
	
	if getWorkOrderRequestResult.Success == true {
		t.Errorf(getWorkOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestProductionScheduleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ProductionSchedule
	//----------------------------------------------------------------------------
	ProductionScheduleObj := model.ProductionSchedule                                                                                                                                                            {ScheduleNumber:"test value for ScheduleNumber",HorizonStart:time.Now(),HorizonEnd:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createProductionScheduleRequestResult := dao.CreateProductionSchedule( ProductionScheduleObj )
	
	if createProductionScheduleRequestResult.Success == false {
		t.Errorf(createProductionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Create ProductionSchedule success...")
	}
	
	createProductionScheduleObj,_ := createProductionScheduleRequestResult.Data. (model.ProductionSchedule)

	// --------------------------------------------------------------
	// Check ProductionSchedule Obj ID
	// --------------------------------------------------------------	
	if createProductionScheduleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ProductionSchedule" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getProductionScheduleRequestResult := dao.GetProductionSchedule( uint64(createProductionScheduleObj.ID) )
	
	if getProductionScheduleRequestResult.Success == false {
		t.Errorf(getProductionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Get ProductionSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getProductionScheduleObj,_ := getProductionScheduleRequestResult.Data. (model.ProductionSchedule)
	compareProductionSchedule := cmp.Equal(createProductionScheduleObj.ID, getProductionScheduleObj.ID)
	
	if  compareProductionSchedule == false	{
		t.Errorf( "Created ProductionSchedule object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllProductionScheduleRequestResult := dao.GetAllProductionSchedule()

	if getAllProductionScheduleRequestResult.Success == false {
			t.Errorf(getAllProductionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ProductionSchedule success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllProductionScheduleObj []model.ProductionSchedule = getAllProductionScheduleRequestResult.Data. ([]model.ProductionSchedule)
		
	equalProductionSchedule := cmp.Equal(createProductionScheduleObj.ID, getAllProductionScheduleObj[len(getAllProductionScheduleObj)-1].ID)
		
	if equalProductionSchedule == false {
		t.Errorf( "Created object is not equal to the last entry in ProductionSchedule[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ProductionSchedule
	// --------------------------------------------------------------	
	deleteProductionScheduleRequestResult := dao.DeleteProductionSchedule(uint64(createProductionScheduleObj.ID))

	if deleteProductionScheduleRequestResult.Success == false {
			t.Errorf(deleteProductionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ProductionSchedule success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getProductionScheduleRequestResult = dao.GetProductionSchedule( uint64(createProductionScheduleObj.ID) )
	
	if getProductionScheduleRequestResult.Success == true {
		t.Errorf(getProductionScheduleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSupplierCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Supplier
	//----------------------------------------------------------------------------
	SupplierObj := model.Supplier                                                                                                            {Name:"test value for Name",SupplierCode:"test value for SupplierCode",Address:new Address(),SupplierTier:0,PaymentTerms:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSupplierRequestResult := dao.CreateSupplier( SupplierObj )
	
	if createSupplierRequestResult.Success == false {
		t.Errorf(createSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Create Supplier success...")
	}
	
	createSupplierObj,_ := createSupplierRequestResult.Data. (model.Supplier)

	// --------------------------------------------------------------
	// Check Supplier Obj ID
	// --------------------------------------------------------------	
	if createSupplierObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Supplier" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSupplierRequestResult := dao.GetSupplier( uint64(createSupplierObj.ID) )
	
	if getSupplierRequestResult.Success == false {
		t.Errorf(getSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Get Supplier success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSupplierObj,_ := getSupplierRequestResult.Data. (model.Supplier)
	compareSupplier := cmp.Equal(createSupplierObj.ID, getSupplierObj.ID)
	
	if  compareSupplier == false	{
		t.Errorf( "Created Supplier object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSupplierRequestResult := dao.GetAllSupplier()

	if getAllSupplierRequestResult.Success == false {
			t.Errorf(getAllSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Supplier success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSupplierObj []model.Supplier = getAllSupplierRequestResult.Data. ([]model.Supplier)
		
	equalSupplier := cmp.Equal(createSupplierObj.ID, getAllSupplierObj[len(getAllSupplierObj)-1].ID)
		
	if equalSupplier == false {
		t.Errorf( "Created object is not equal to the last entry in Supplier[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Supplier
	// --------------------------------------------------------------	
	deleteSupplierRequestResult := dao.DeleteSupplier(uint64(createSupplierObj.ID))

	if deleteSupplierRequestResult.Success == false {
			t.Errorf(deleteSupplierRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Supplier success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSupplierRequestResult = dao.GetSupplier( uint64(createSupplierObj.ID) )
	
	if getSupplierRequestResult.Success == true {
		t.Errorf(getSupplierRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPurchaseOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PurchaseOrder
	//----------------------------------------------------------------------------
	PurchaseOrderObj := model.PurchaseOrder                                                                                                                    {PoNumber:"test value for PoNumber",OrderDate:time.Now(),TotalAmount:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPurchaseOrderRequestResult := dao.CreatePurchaseOrder( PurchaseOrderObj )
	
	if createPurchaseOrderRequestResult.Success == false {
		t.Errorf(createPurchaseOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create PurchaseOrder success...")
	}
	
	createPurchaseOrderObj,_ := createPurchaseOrderRequestResult.Data. (model.PurchaseOrder)

	// --------------------------------------------------------------
	// Check PurchaseOrder Obj ID
	// --------------------------------------------------------------	
	if createPurchaseOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PurchaseOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPurchaseOrderRequestResult := dao.GetPurchaseOrder( uint64(createPurchaseOrderObj.ID) )
	
	if getPurchaseOrderRequestResult.Success == false {
		t.Errorf(getPurchaseOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get PurchaseOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPurchaseOrderObj,_ := getPurchaseOrderRequestResult.Data. (model.PurchaseOrder)
	comparePurchaseOrder := cmp.Equal(createPurchaseOrderObj.ID, getPurchaseOrderObj.ID)
	
	if  comparePurchaseOrder == false	{
		t.Errorf( "Created PurchaseOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPurchaseOrderRequestResult := dao.GetAllPurchaseOrder()

	if getAllPurchaseOrderRequestResult.Success == false {
			t.Errorf(getAllPurchaseOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PurchaseOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPurchaseOrderObj []model.PurchaseOrder = getAllPurchaseOrderRequestResult.Data. ([]model.PurchaseOrder)
		
	equalPurchaseOrder := cmp.Equal(createPurchaseOrderObj.ID, getAllPurchaseOrderObj[len(getAllPurchaseOrderObj)-1].ID)
		
	if equalPurchaseOrder == false {
		t.Errorf( "Created object is not equal to the last entry in PurchaseOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PurchaseOrder
	// --------------------------------------------------------------	
	deletePurchaseOrderRequestResult := dao.DeletePurchaseOrder(uint64(createPurchaseOrderObj.ID))

	if deletePurchaseOrderRequestResult.Success == false {
			t.Errorf(deletePurchaseOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PurchaseOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPurchaseOrderRequestResult = dao.GetPurchaseOrder( uint64(createPurchaseOrderObj.ID) )
	
	if getPurchaseOrderRequestResult.Success == true {
		t.Errorf(getPurchaseOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPurchaseOrderLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PurchaseOrderLine
	//----------------------------------------------------------------------------
	PurchaseOrderLineObj := model.PurchaseOrderLine                                                                                                                    {LineNumber:100,Quantity:new Quantity(),UnitPrice:new Money(),DueDate:time.Now()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPurchaseOrderLineRequestResult := dao.CreatePurchaseOrderLine( PurchaseOrderLineObj )
	
	if createPurchaseOrderLineRequestResult.Success == false {
		t.Errorf(createPurchaseOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create PurchaseOrderLine success...")
	}
	
	createPurchaseOrderLineObj,_ := createPurchaseOrderLineRequestResult.Data. (model.PurchaseOrderLine)

	// --------------------------------------------------------------
	// Check PurchaseOrderLine Obj ID
	// --------------------------------------------------------------	
	if createPurchaseOrderLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PurchaseOrderLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPurchaseOrderLineRequestResult := dao.GetPurchaseOrderLine( uint64(createPurchaseOrderLineObj.ID) )
	
	if getPurchaseOrderLineRequestResult.Success == false {
		t.Errorf(getPurchaseOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get PurchaseOrderLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPurchaseOrderLineObj,_ := getPurchaseOrderLineRequestResult.Data. (model.PurchaseOrderLine)
	comparePurchaseOrderLine := cmp.Equal(createPurchaseOrderLineObj.ID, getPurchaseOrderLineObj.ID)
	
	if  comparePurchaseOrderLine == false	{
		t.Errorf( "Created PurchaseOrderLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPurchaseOrderLineRequestResult := dao.GetAllPurchaseOrderLine()

	if getAllPurchaseOrderLineRequestResult.Success == false {
			t.Errorf(getAllPurchaseOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PurchaseOrderLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPurchaseOrderLineObj []model.PurchaseOrderLine = getAllPurchaseOrderLineRequestResult.Data. ([]model.PurchaseOrderLine)
		
	equalPurchaseOrderLine := cmp.Equal(createPurchaseOrderLineObj.ID, getAllPurchaseOrderLineObj[len(getAllPurchaseOrderLineObj)-1].ID)
		
	if equalPurchaseOrderLine == false {
		t.Errorf( "Created object is not equal to the last entry in PurchaseOrderLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PurchaseOrderLine
	// --------------------------------------------------------------	
	deletePurchaseOrderLineRequestResult := dao.DeletePurchaseOrderLine(uint64(createPurchaseOrderLineObj.ID))

	if deletePurchaseOrderLineRequestResult.Success == false {
			t.Errorf(deletePurchaseOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PurchaseOrderLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPurchaseOrderLineRequestResult = dao.GetPurchaseOrderLine( uint64(createPurchaseOrderLineObj.ID) )
	
	if getPurchaseOrderLineRequestResult.Success == true {
		t.Errorf(getPurchaseOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestGoodsReceiptCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for GoodsReceipt
	//----------------------------------------------------------------------------
	GoodsReceiptObj := model.GoodsReceipt                                                                                                    {ReceiptNumber:"test value for ReceiptNumber",ReceiptDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createGoodsReceiptRequestResult := dao.CreateGoodsReceipt( GoodsReceiptObj )
	
	if createGoodsReceiptRequestResult.Success == false {
		t.Errorf(createGoodsReceiptRequestResult.Msg)
	} else {
		fmt.Println("Check Create GoodsReceipt success...")
	}
	
	createGoodsReceiptObj,_ := createGoodsReceiptRequestResult.Data. (model.GoodsReceipt)

	// --------------------------------------------------------------
	// Check GoodsReceipt Obj ID
	// --------------------------------------------------------------	
	if createGoodsReceiptObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for GoodsReceipt" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getGoodsReceiptRequestResult := dao.GetGoodsReceipt( uint64(createGoodsReceiptObj.ID) )
	
	if getGoodsReceiptRequestResult.Success == false {
		t.Errorf(getGoodsReceiptRequestResult.Msg)
	} else {
		fmt.Println("Check Get GoodsReceipt success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getGoodsReceiptObj,_ := getGoodsReceiptRequestResult.Data. (model.GoodsReceipt)
	compareGoodsReceipt := cmp.Equal(createGoodsReceiptObj.ID, getGoodsReceiptObj.ID)
	
	if  compareGoodsReceipt == false	{
		t.Errorf( "Created GoodsReceipt object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllGoodsReceiptRequestResult := dao.GetAllGoodsReceipt()

	if getAllGoodsReceiptRequestResult.Success == false {
			t.Errorf(getAllGoodsReceiptRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll GoodsReceipt success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllGoodsReceiptObj []model.GoodsReceipt = getAllGoodsReceiptRequestResult.Data. ([]model.GoodsReceipt)
		
	equalGoodsReceipt := cmp.Equal(createGoodsReceiptObj.ID, getAllGoodsReceiptObj[len(getAllGoodsReceiptObj)-1].ID)
		
	if equalGoodsReceipt == false {
		t.Errorf( "Created object is not equal to the last entry in GoodsReceipt[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for GoodsReceipt
	// --------------------------------------------------------------	
	deleteGoodsReceiptRequestResult := dao.DeleteGoodsReceipt(uint64(createGoodsReceiptObj.ID))

	if deleteGoodsReceiptRequestResult.Success == false {
			t.Errorf(deleteGoodsReceiptRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion GoodsReceipt success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getGoodsReceiptRequestResult = dao.GetGoodsReceipt( uint64(createGoodsReceiptObj.ID) )
	
	if getGoodsReceiptRequestResult.Success == true {
		t.Errorf(getGoodsReceiptRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestGoodsReceiptLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for GoodsReceiptLine
	//----------------------------------------------------------------------------
	GoodsReceiptLineObj := model.GoodsReceiptLine                                                                                            {LineNumber:100,ReceivedQuantity:new Quantity(),AcceptedQuantity:new Quantity(),RejectedQuantity:new Quantity(),Lot:new LotId()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createGoodsReceiptLineRequestResult := dao.CreateGoodsReceiptLine( GoodsReceiptLineObj )
	
	if createGoodsReceiptLineRequestResult.Success == false {
		t.Errorf(createGoodsReceiptLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create GoodsReceiptLine success...")
	}
	
	createGoodsReceiptLineObj,_ := createGoodsReceiptLineRequestResult.Data. (model.GoodsReceiptLine)

	// --------------------------------------------------------------
	// Check GoodsReceiptLine Obj ID
	// --------------------------------------------------------------	
	if createGoodsReceiptLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for GoodsReceiptLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getGoodsReceiptLineRequestResult := dao.GetGoodsReceiptLine( uint64(createGoodsReceiptLineObj.ID) )
	
	if getGoodsReceiptLineRequestResult.Success == false {
		t.Errorf(getGoodsReceiptLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get GoodsReceiptLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getGoodsReceiptLineObj,_ := getGoodsReceiptLineRequestResult.Data. (model.GoodsReceiptLine)
	compareGoodsReceiptLine := cmp.Equal(createGoodsReceiptLineObj.ID, getGoodsReceiptLineObj.ID)
	
	if  compareGoodsReceiptLine == false	{
		t.Errorf( "Created GoodsReceiptLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllGoodsReceiptLineRequestResult := dao.GetAllGoodsReceiptLine()

	if getAllGoodsReceiptLineRequestResult.Success == false {
			t.Errorf(getAllGoodsReceiptLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll GoodsReceiptLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllGoodsReceiptLineObj []model.GoodsReceiptLine = getAllGoodsReceiptLineRequestResult.Data. ([]model.GoodsReceiptLine)
		
	equalGoodsReceiptLine := cmp.Equal(createGoodsReceiptLineObj.ID, getAllGoodsReceiptLineObj[len(getAllGoodsReceiptLineObj)-1].ID)
		
	if equalGoodsReceiptLine == false {
		t.Errorf( "Created object is not equal to the last entry in GoodsReceiptLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for GoodsReceiptLine
	// --------------------------------------------------------------	
	deleteGoodsReceiptLineRequestResult := dao.DeleteGoodsReceiptLine(uint64(createGoodsReceiptLineObj.ID))

	if deleteGoodsReceiptLineRequestResult.Success == false {
			t.Errorf(deleteGoodsReceiptLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion GoodsReceiptLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getGoodsReceiptLineRequestResult = dao.GetGoodsReceiptLine( uint64(createGoodsReceiptLineObj.ID) )
	
	if getGoodsReceiptLineRequestResult.Success == true {
		t.Errorf(getGoodsReceiptLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestWarehouseCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Warehouse
	//----------------------------------------------------------------------------
	WarehouseObj := model.Warehouse                                                                                            {Name:"test value for Name",WarehouseCode:"test value for WarehouseCode",Address:new Address(),WarehouseType:0}

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


func TestLocationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Location
	//----------------------------------------------------------------------------
	LocationObj := model.Location                                                                            {LocationCode:"test value for LocationCode",Description:"test value for Description",LocationType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLocationRequestResult := dao.CreateLocation( LocationObj )
	
	if createLocationRequestResult.Success == false {
		t.Errorf(createLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Location success...")
	}
	
	createLocationObj,_ := createLocationRequestResult.Data. (model.Location)

	// --------------------------------------------------------------
	// Check Location Obj ID
	// --------------------------------------------------------------	
	if createLocationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Location" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLocationRequestResult := dao.GetLocation( uint64(createLocationObj.ID) )
	
	if getLocationRequestResult.Success == false {
		t.Errorf(getLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Location success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLocationObj,_ := getLocationRequestResult.Data. (model.Location)
	compareLocation := cmp.Equal(createLocationObj.ID, getLocationObj.ID)
	
	if  compareLocation == false	{
		t.Errorf( "Created Location object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLocationRequestResult := dao.GetAllLocation()

	if getAllLocationRequestResult.Success == false {
			t.Errorf(getAllLocationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Location success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLocationObj []model.Location = getAllLocationRequestResult.Data. ([]model.Location)
		
	equalLocation := cmp.Equal(createLocationObj.ID, getAllLocationObj[len(getAllLocationObj)-1].ID)
		
	if equalLocation == false {
		t.Errorf( "Created object is not equal to the last entry in Location[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Location
	// --------------------------------------------------------------	
	deleteLocationRequestResult := dao.DeleteLocation(uint64(createLocationObj.ID))

	if deleteLocationRequestResult.Success == false {
			t.Errorf(deleteLocationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Location success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLocationRequestResult = dao.GetLocation( uint64(createLocationObj.ID) )
	
	if getLocationRequestResult.Success == true {
		t.Errorf(getLocationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInventoryItemCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InventoryItem
	//----------------------------------------------------------------------------
	InventoryItemObj := model.InventoryItem                                                            {QuantityOnHand:new Quantity(),QuantityReserved:new Quantity(),LotNumber:new LotId(),SerialNumber:new SerialId()}

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


func TestInventoryTransactionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InventoryTransaction
	//----------------------------------------------------------------------------
	InventoryTransactionObj := model.InventoryTransaction                                                                                                                                                    {TransactionNumber:"test value for TransactionNumber",Quantity:new Quantity(),TransactionDateTime:time.Now(),ReferenceDocument:"test value for ReferenceDocument",TransactionType:0}

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


func TestCustomerCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Customer
	//----------------------------------------------------------------------------
	CustomerObj := model.Customer                                                                                            {Name:"test value for Name",CustomerCode:"test value for CustomerCode",Address:new Address(),CustomerType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCustomerRequestResult := dao.CreateCustomer( CustomerObj )
	
	if createCustomerRequestResult.Success == false {
		t.Errorf(createCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Create Customer success...")
	}
	
	createCustomerObj,_ := createCustomerRequestResult.Data. (model.Customer)

	// --------------------------------------------------------------
	// Check Customer Obj ID
	// --------------------------------------------------------------	
	if createCustomerObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Customer" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCustomerRequestResult := dao.GetCustomer( uint64(createCustomerObj.ID) )
	
	if getCustomerRequestResult.Success == false {
		t.Errorf(getCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Get Customer success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCustomerObj,_ := getCustomerRequestResult.Data. (model.Customer)
	compareCustomer := cmp.Equal(createCustomerObj.ID, getCustomerObj.ID)
	
	if  compareCustomer == false	{
		t.Errorf( "Created Customer object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCustomerRequestResult := dao.GetAllCustomer()

	if getAllCustomerRequestResult.Success == false {
			t.Errorf(getAllCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Customer success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCustomerObj []model.Customer = getAllCustomerRequestResult.Data. ([]model.Customer)
		
	equalCustomer := cmp.Equal(createCustomerObj.ID, getAllCustomerObj[len(getAllCustomerObj)-1].ID)
		
	if equalCustomer == false {
		t.Errorf( "Created object is not equal to the last entry in Customer[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Customer
	// --------------------------------------------------------------	
	deleteCustomerRequestResult := dao.DeleteCustomer(uint64(createCustomerObj.ID))

	if deleteCustomerRequestResult.Success == false {
			t.Errorf(deleteCustomerRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Customer success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCustomerRequestResult = dao.GetCustomer( uint64(createCustomerObj.ID) )
	
	if getCustomerRequestResult.Success == true {
		t.Errorf(getCustomerRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSalesOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SalesOrder
	//----------------------------------------------------------------------------
	SalesOrderObj := model.SalesOrder                                                                                                                    {OrderNumber:"test value for OrderNumber",OrderDate:time.Now(),TotalAmount:new Money(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSalesOrderRequestResult := dao.CreateSalesOrder( SalesOrderObj )
	
	if createSalesOrderRequestResult.Success == false {
		t.Errorf(createSalesOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create SalesOrder success...")
	}
	
	createSalesOrderObj,_ := createSalesOrderRequestResult.Data. (model.SalesOrder)

	// --------------------------------------------------------------
	// Check SalesOrder Obj ID
	// --------------------------------------------------------------	
	if createSalesOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SalesOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSalesOrderRequestResult := dao.GetSalesOrder( uint64(createSalesOrderObj.ID) )
	
	if getSalesOrderRequestResult.Success == false {
		t.Errorf(getSalesOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get SalesOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSalesOrderObj,_ := getSalesOrderRequestResult.Data. (model.SalesOrder)
	compareSalesOrder := cmp.Equal(createSalesOrderObj.ID, getSalesOrderObj.ID)
	
	if  compareSalesOrder == false	{
		t.Errorf( "Created SalesOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSalesOrderRequestResult := dao.GetAllSalesOrder()

	if getAllSalesOrderRequestResult.Success == false {
			t.Errorf(getAllSalesOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SalesOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSalesOrderObj []model.SalesOrder = getAllSalesOrderRequestResult.Data. ([]model.SalesOrder)
		
	equalSalesOrder := cmp.Equal(createSalesOrderObj.ID, getAllSalesOrderObj[len(getAllSalesOrderObj)-1].ID)
		
	if equalSalesOrder == false {
		t.Errorf( "Created object is not equal to the last entry in SalesOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SalesOrder
	// --------------------------------------------------------------	
	deleteSalesOrderRequestResult := dao.DeleteSalesOrder(uint64(createSalesOrderObj.ID))

	if deleteSalesOrderRequestResult.Success == false {
			t.Errorf(deleteSalesOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SalesOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSalesOrderRequestResult = dao.GetSalesOrder( uint64(createSalesOrderObj.ID) )
	
	if getSalesOrderRequestResult.Success == true {
		t.Errorf(getSalesOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSalesOrderLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SalesOrderLine
	//----------------------------------------------------------------------------
	SalesOrderLineObj := model.SalesOrderLine                                                                                                                    {LineNumber:100,Quantity:new Quantity(),UnitPrice:new Money(),DueDate:time.Now()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSalesOrderLineRequestResult := dao.CreateSalesOrderLine( SalesOrderLineObj )
	
	if createSalesOrderLineRequestResult.Success == false {
		t.Errorf(createSalesOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create SalesOrderLine success...")
	}
	
	createSalesOrderLineObj,_ := createSalesOrderLineRequestResult.Data. (model.SalesOrderLine)

	// --------------------------------------------------------------
	// Check SalesOrderLine Obj ID
	// --------------------------------------------------------------	
	if createSalesOrderLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SalesOrderLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSalesOrderLineRequestResult := dao.GetSalesOrderLine( uint64(createSalesOrderLineObj.ID) )
	
	if getSalesOrderLineRequestResult.Success == false {
		t.Errorf(getSalesOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get SalesOrderLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSalesOrderLineObj,_ := getSalesOrderLineRequestResult.Data. (model.SalesOrderLine)
	compareSalesOrderLine := cmp.Equal(createSalesOrderLineObj.ID, getSalesOrderLineObj.ID)
	
	if  compareSalesOrderLine == false	{
		t.Errorf( "Created SalesOrderLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSalesOrderLineRequestResult := dao.GetAllSalesOrderLine()

	if getAllSalesOrderLineRequestResult.Success == false {
			t.Errorf(getAllSalesOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SalesOrderLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSalesOrderLineObj []model.SalesOrderLine = getAllSalesOrderLineRequestResult.Data. ([]model.SalesOrderLine)
		
	equalSalesOrderLine := cmp.Equal(createSalesOrderLineObj.ID, getAllSalesOrderLineObj[len(getAllSalesOrderLineObj)-1].ID)
		
	if equalSalesOrderLine == false {
		t.Errorf( "Created object is not equal to the last entry in SalesOrderLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SalesOrderLine
	// --------------------------------------------------------------	
	deleteSalesOrderLineRequestResult := dao.DeleteSalesOrderLine(uint64(createSalesOrderLineObj.ID))

	if deleteSalesOrderLineRequestResult.Success == false {
			t.Errorf(deleteSalesOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SalesOrderLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSalesOrderLineRequestResult = dao.GetSalesOrderLine( uint64(createSalesOrderLineObj.ID) )
	
	if getSalesOrderLineRequestResult.Success == true {
		t.Errorf(getSalesOrderLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestQualitySpecificationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for QualitySpecification
	//----------------------------------------------------------------------------
	QualitySpecificationObj := model.QualitySpecification                                                                                            {SpecCode:"test value for SpecCode",Name:"test value for Name",Version:"test value for Version"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createQualitySpecificationRequestResult := dao.CreateQualitySpecification( QualitySpecificationObj )
	
	if createQualitySpecificationRequestResult.Success == false {
		t.Errorf(createQualitySpecificationRequestResult.Msg)
	} else {
		fmt.Println("Check Create QualitySpecification success...")
	}
	
	createQualitySpecificationObj,_ := createQualitySpecificationRequestResult.Data. (model.QualitySpecification)

	// --------------------------------------------------------------
	// Check QualitySpecification Obj ID
	// --------------------------------------------------------------	
	if createQualitySpecificationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for QualitySpecification" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getQualitySpecificationRequestResult := dao.GetQualitySpecification( uint64(createQualitySpecificationObj.ID) )
	
	if getQualitySpecificationRequestResult.Success == false {
		t.Errorf(getQualitySpecificationRequestResult.Msg)
	} else {
		fmt.Println("Check Get QualitySpecification success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getQualitySpecificationObj,_ := getQualitySpecificationRequestResult.Data. (model.QualitySpecification)
	compareQualitySpecification := cmp.Equal(createQualitySpecificationObj.ID, getQualitySpecificationObj.ID)
	
	if  compareQualitySpecification == false	{
		t.Errorf( "Created QualitySpecification object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllQualitySpecificationRequestResult := dao.GetAllQualitySpecification()

	if getAllQualitySpecificationRequestResult.Success == false {
			t.Errorf(getAllQualitySpecificationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll QualitySpecification success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllQualitySpecificationObj []model.QualitySpecification = getAllQualitySpecificationRequestResult.Data. ([]model.QualitySpecification)
		
	equalQualitySpecification := cmp.Equal(createQualitySpecificationObj.ID, getAllQualitySpecificationObj[len(getAllQualitySpecificationObj)-1].ID)
		
	if equalQualitySpecification == false {
		t.Errorf( "Created object is not equal to the last entry in QualitySpecification[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for QualitySpecification
	// --------------------------------------------------------------	
	deleteQualitySpecificationRequestResult := dao.DeleteQualitySpecification(uint64(createQualitySpecificationObj.ID))

	if deleteQualitySpecificationRequestResult.Success == false {
			t.Errorf(deleteQualitySpecificationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion QualitySpecification success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getQualitySpecificationRequestResult = dao.GetQualitySpecification( uint64(createQualitySpecificationObj.ID) )
	
	if getQualitySpecificationRequestResult.Success == true {
		t.Errorf(getQualitySpecificationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInspectionPlanCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InspectionPlan
	//----------------------------------------------------------------------------
	InspectionPlanObj := model.InspectionPlan                                                                                            {PlanNumber:"test value for PlanNumber",Revision:"test value for Revision",SamplingPlan:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInspectionPlanRequestResult := dao.CreateInspectionPlan( InspectionPlanObj )
	
	if createInspectionPlanRequestResult.Success == false {
		t.Errorf(createInspectionPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Create InspectionPlan success...")
	}
	
	createInspectionPlanObj,_ := createInspectionPlanRequestResult.Data. (model.InspectionPlan)

	// --------------------------------------------------------------
	// Check InspectionPlan Obj ID
	// --------------------------------------------------------------	
	if createInspectionPlanObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InspectionPlan" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInspectionPlanRequestResult := dao.GetInspectionPlan( uint64(createInspectionPlanObj.ID) )
	
	if getInspectionPlanRequestResult.Success == false {
		t.Errorf(getInspectionPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Get InspectionPlan success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInspectionPlanObj,_ := getInspectionPlanRequestResult.Data. (model.InspectionPlan)
	compareInspectionPlan := cmp.Equal(createInspectionPlanObj.ID, getInspectionPlanObj.ID)
	
	if  compareInspectionPlan == false	{
		t.Errorf( "Created InspectionPlan object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInspectionPlanRequestResult := dao.GetAllInspectionPlan()

	if getAllInspectionPlanRequestResult.Success == false {
			t.Errorf(getAllInspectionPlanRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InspectionPlan success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInspectionPlanObj []model.InspectionPlan = getAllInspectionPlanRequestResult.Data. ([]model.InspectionPlan)
		
	equalInspectionPlan := cmp.Equal(createInspectionPlanObj.ID, getAllInspectionPlanObj[len(getAllInspectionPlanObj)-1].ID)
		
	if equalInspectionPlan == false {
		t.Errorf( "Created object is not equal to the last entry in InspectionPlan[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InspectionPlan
	// --------------------------------------------------------------	
	deleteInspectionPlanRequestResult := dao.DeleteInspectionPlan(uint64(createInspectionPlanObj.ID))

	if deleteInspectionPlanRequestResult.Success == false {
			t.Errorf(deleteInspectionPlanRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InspectionPlan success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInspectionPlanRequestResult = dao.GetInspectionPlan( uint64(createInspectionPlanObj.ID) )
	
	if getInspectionPlanRequestResult.Success == true {
		t.Errorf(getInspectionPlanRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInspectionCharacteristicCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InspectionCharacteristic
	//----------------------------------------------------------------------------
	InspectionCharacteristicObj := model.InspectionCharacteristic                                                                                                                            {CharacteristicCode:"test value for CharacteristicCode",Name:"test value for Name",LowerSpecLimit:new Measurement(),UpperSpecLimit:new Measurement(),Target:new Measurement(),MeasurementType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInspectionCharacteristicRequestResult := dao.CreateInspectionCharacteristic( InspectionCharacteristicObj )
	
	if createInspectionCharacteristicRequestResult.Success == false {
		t.Errorf(createInspectionCharacteristicRequestResult.Msg)
	} else {
		fmt.Println("Check Create InspectionCharacteristic success...")
	}
	
	createInspectionCharacteristicObj,_ := createInspectionCharacteristicRequestResult.Data. (model.InspectionCharacteristic)

	// --------------------------------------------------------------
	// Check InspectionCharacteristic Obj ID
	// --------------------------------------------------------------	
	if createInspectionCharacteristicObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InspectionCharacteristic" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInspectionCharacteristicRequestResult := dao.GetInspectionCharacteristic( uint64(createInspectionCharacteristicObj.ID) )
	
	if getInspectionCharacteristicRequestResult.Success == false {
		t.Errorf(getInspectionCharacteristicRequestResult.Msg)
	} else {
		fmt.Println("Check Get InspectionCharacteristic success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInspectionCharacteristicObj,_ := getInspectionCharacteristicRequestResult.Data. (model.InspectionCharacteristic)
	compareInspectionCharacteristic := cmp.Equal(createInspectionCharacteristicObj.ID, getInspectionCharacteristicObj.ID)
	
	if  compareInspectionCharacteristic == false	{
		t.Errorf( "Created InspectionCharacteristic object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInspectionCharacteristicRequestResult := dao.GetAllInspectionCharacteristic()

	if getAllInspectionCharacteristicRequestResult.Success == false {
			t.Errorf(getAllInspectionCharacteristicRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InspectionCharacteristic success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInspectionCharacteristicObj []model.InspectionCharacteristic = getAllInspectionCharacteristicRequestResult.Data. ([]model.InspectionCharacteristic)
		
	equalInspectionCharacteristic := cmp.Equal(createInspectionCharacteristicObj.ID, getAllInspectionCharacteristicObj[len(getAllInspectionCharacteristicObj)-1].ID)
		
	if equalInspectionCharacteristic == false {
		t.Errorf( "Created object is not equal to the last entry in InspectionCharacteristic[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InspectionCharacteristic
	// --------------------------------------------------------------	
	deleteInspectionCharacteristicRequestResult := dao.DeleteInspectionCharacteristic(uint64(createInspectionCharacteristicObj.ID))

	if deleteInspectionCharacteristicRequestResult.Success == false {
			t.Errorf(deleteInspectionCharacteristicRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InspectionCharacteristic success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInspectionCharacteristicRequestResult = dao.GetInspectionCharacteristic( uint64(createInspectionCharacteristicObj.ID) )
	
	if getInspectionCharacteristicRequestResult.Success == true {
		t.Errorf(getInspectionCharacteristicRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInspectionLotCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InspectionLot
	//----------------------------------------------------------------------------
	InspectionLotObj := model.InspectionLot                                                                                                                                                                    {LotNumber:"test value for LotNumber",Quantity:new Quantity(),SampleSize:100,CreatedOn:time.Now(),InspectionType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInspectionLotRequestResult := dao.CreateInspectionLot( InspectionLotObj )
	
	if createInspectionLotRequestResult.Success == false {
		t.Errorf(createInspectionLotRequestResult.Msg)
	} else {
		fmt.Println("Check Create InspectionLot success...")
	}
	
	createInspectionLotObj,_ := createInspectionLotRequestResult.Data. (model.InspectionLot)

	// --------------------------------------------------------------
	// Check InspectionLot Obj ID
	// --------------------------------------------------------------	
	if createInspectionLotObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InspectionLot" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInspectionLotRequestResult := dao.GetInspectionLot( uint64(createInspectionLotObj.ID) )
	
	if getInspectionLotRequestResult.Success == false {
		t.Errorf(getInspectionLotRequestResult.Msg)
	} else {
		fmt.Println("Check Get InspectionLot success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInspectionLotObj,_ := getInspectionLotRequestResult.Data. (model.InspectionLot)
	compareInspectionLot := cmp.Equal(createInspectionLotObj.ID, getInspectionLotObj.ID)
	
	if  compareInspectionLot == false	{
		t.Errorf( "Created InspectionLot object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInspectionLotRequestResult := dao.GetAllInspectionLot()

	if getAllInspectionLotRequestResult.Success == false {
			t.Errorf(getAllInspectionLotRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InspectionLot success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInspectionLotObj []model.InspectionLot = getAllInspectionLotRequestResult.Data. ([]model.InspectionLot)
		
	equalInspectionLot := cmp.Equal(createInspectionLotObj.ID, getAllInspectionLotObj[len(getAllInspectionLotObj)-1].ID)
		
	if equalInspectionLot == false {
		t.Errorf( "Created object is not equal to the last entry in InspectionLot[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InspectionLot
	// --------------------------------------------------------------	
	deleteInspectionLotRequestResult := dao.DeleteInspectionLot(uint64(createInspectionLotObj.ID))

	if deleteInspectionLotRequestResult.Success == false {
			t.Errorf(deleteInspectionLotRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InspectionLot success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInspectionLotRequestResult = dao.GetInspectionLot( uint64(createInspectionLotObj.ID) )
	
	if getInspectionLotRequestResult.Success == true {
		t.Errorf(getInspectionLotRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInspectionResultCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InspectionResult
	//----------------------------------------------------------------------------
	InspectionResultObj := model.InspectionResult                                                                                                                    {ResultValue:new Measurement(),RecordedOn:time.Now(),Notes:"test value for Notes",ResultStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInspectionResultRequestResult := dao.CreateInspectionResult( InspectionResultObj )
	
	if createInspectionResultRequestResult.Success == false {
		t.Errorf(createInspectionResultRequestResult.Msg)
	} else {
		fmt.Println("Check Create InspectionResult success...")
	}
	
	createInspectionResultObj,_ := createInspectionResultRequestResult.Data. (model.InspectionResult)

	// --------------------------------------------------------------
	// Check InspectionResult Obj ID
	// --------------------------------------------------------------	
	if createInspectionResultObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InspectionResult" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInspectionResultRequestResult := dao.GetInspectionResult( uint64(createInspectionResultObj.ID) )
	
	if getInspectionResultRequestResult.Success == false {
		t.Errorf(getInspectionResultRequestResult.Msg)
	} else {
		fmt.Println("Check Get InspectionResult success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInspectionResultObj,_ := getInspectionResultRequestResult.Data. (model.InspectionResult)
	compareInspectionResult := cmp.Equal(createInspectionResultObj.ID, getInspectionResultObj.ID)
	
	if  compareInspectionResult == false	{
		t.Errorf( "Created InspectionResult object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInspectionResultRequestResult := dao.GetAllInspectionResult()

	if getAllInspectionResultRequestResult.Success == false {
			t.Errorf(getAllInspectionResultRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InspectionResult success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInspectionResultObj []model.InspectionResult = getAllInspectionResultRequestResult.Data. ([]model.InspectionResult)
		
	equalInspectionResult := cmp.Equal(createInspectionResultObj.ID, getAllInspectionResultObj[len(getAllInspectionResultObj)-1].ID)
		
	if equalInspectionResult == false {
		t.Errorf( "Created object is not equal to the last entry in InspectionResult[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InspectionResult
	// --------------------------------------------------------------	
	deleteInspectionResultRequestResult := dao.DeleteInspectionResult(uint64(createInspectionResultObj.ID))

	if deleteInspectionResultRequestResult.Success == false {
			t.Errorf(deleteInspectionResultRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InspectionResult success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInspectionResultRequestResult = dao.GetInspectionResult( uint64(createInspectionResultObj.ID) )
	
	if getInspectionResultRequestResult.Success == true {
		t.Errorf(getInspectionResultRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestNonconformanceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Nonconformance
	//----------------------------------------------------------------------------
	NonconformanceObj := model.Nonconformance                                                                                                                                            {NcNumber:"test value for NcNumber",Description:"test value for Description",ContainmentAction:"test value for ContainmentAction",NcType:0,Severity:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createNonconformanceRequestResult := dao.CreateNonconformance( NonconformanceObj )
	
	if createNonconformanceRequestResult.Success == false {
		t.Errorf(createNonconformanceRequestResult.Msg)
	} else {
		fmt.Println("Check Create Nonconformance success...")
	}
	
	createNonconformanceObj,_ := createNonconformanceRequestResult.Data. (model.Nonconformance)

	// --------------------------------------------------------------
	// Check Nonconformance Obj ID
	// --------------------------------------------------------------	
	if createNonconformanceObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Nonconformance" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getNonconformanceRequestResult := dao.GetNonconformance( uint64(createNonconformanceObj.ID) )
	
	if getNonconformanceRequestResult.Success == false {
		t.Errorf(getNonconformanceRequestResult.Msg)
	} else {
		fmt.Println("Check Get Nonconformance success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getNonconformanceObj,_ := getNonconformanceRequestResult.Data. (model.Nonconformance)
	compareNonconformance := cmp.Equal(createNonconformanceObj.ID, getNonconformanceObj.ID)
	
	if  compareNonconformance == false	{
		t.Errorf( "Created Nonconformance object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllNonconformanceRequestResult := dao.GetAllNonconformance()

	if getAllNonconformanceRequestResult.Success == false {
			t.Errorf(getAllNonconformanceRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Nonconformance success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllNonconformanceObj []model.Nonconformance = getAllNonconformanceRequestResult.Data. ([]model.Nonconformance)
		
	equalNonconformance := cmp.Equal(createNonconformanceObj.ID, getAllNonconformanceObj[len(getAllNonconformanceObj)-1].ID)
		
	if equalNonconformance == false {
		t.Errorf( "Created object is not equal to the last entry in Nonconformance[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Nonconformance
	// --------------------------------------------------------------	
	deleteNonconformanceRequestResult := dao.DeleteNonconformance(uint64(createNonconformanceObj.ID))

	if deleteNonconformanceRequestResult.Success == false {
			t.Errorf(deleteNonconformanceRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Nonconformance success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getNonconformanceRequestResult = dao.GetNonconformance( uint64(createNonconformanceObj.ID) )
	
	if getNonconformanceRequestResult.Success == true {
		t.Errorf(getNonconformanceRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestCorrectiveActionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for CorrectiveAction
	//----------------------------------------------------------------------------
	CorrectiveActionObj := model.CorrectiveAction                                                                                                                                                                    {CapaNumber:"test value for CapaNumber",RootCause:"test value for RootCause",CorrectiveAction:"test value for CorrectiveAction",VerificationDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createCorrectiveActionRequestResult := dao.CreateCorrectiveAction( CorrectiveActionObj )
	
	if createCorrectiveActionRequestResult.Success == false {
		t.Errorf(createCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Check Create CorrectiveAction success...")
	}
	
	createCorrectiveActionObj,_ := createCorrectiveActionRequestResult.Data. (model.CorrectiveAction)

	// --------------------------------------------------------------
	// Check CorrectiveAction Obj ID
	// --------------------------------------------------------------	
	if createCorrectiveActionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for CorrectiveAction" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getCorrectiveActionRequestResult := dao.GetCorrectiveAction( uint64(createCorrectiveActionObj.ID) )
	
	if getCorrectiveActionRequestResult.Success == false {
		t.Errorf(getCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Check Get CorrectiveAction success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getCorrectiveActionObj,_ := getCorrectiveActionRequestResult.Data. (model.CorrectiveAction)
	compareCorrectiveAction := cmp.Equal(createCorrectiveActionObj.ID, getCorrectiveActionObj.ID)
	
	if  compareCorrectiveAction == false	{
		t.Errorf( "Created CorrectiveAction object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllCorrectiveActionRequestResult := dao.GetAllCorrectiveAction()

	if getAllCorrectiveActionRequestResult.Success == false {
			t.Errorf(getAllCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll CorrectiveAction success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllCorrectiveActionObj []model.CorrectiveAction = getAllCorrectiveActionRequestResult.Data. ([]model.CorrectiveAction)
		
	equalCorrectiveAction := cmp.Equal(createCorrectiveActionObj.ID, getAllCorrectiveActionObj[len(getAllCorrectiveActionObj)-1].ID)
		
	if equalCorrectiveAction == false {
		t.Errorf( "Created object is not equal to the last entry in CorrectiveAction[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for CorrectiveAction
	// --------------------------------------------------------------	
	deleteCorrectiveActionRequestResult := dao.DeleteCorrectiveAction(uint64(createCorrectiveActionObj.ID))

	if deleteCorrectiveActionRequestResult.Success == false {
			t.Errorf(deleteCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion CorrectiveAction success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getCorrectiveActionRequestResult = dao.GetCorrectiveAction( uint64(createCorrectiveActionObj.ID) )
	
	if getCorrectiveActionRequestResult.Success == true {
		t.Errorf(getCorrectiveActionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAssetCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Asset
	//----------------------------------------------------------------------------
	AssetObj := model.Asset                                                                                                                                    {AssetTag:"test value for AssetTag",AssetName:"test value for AssetName",CommissioningDate:time.Now(),AssetStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAssetRequestResult := dao.CreateAsset( AssetObj )
	
	if createAssetRequestResult.Success == false {
		t.Errorf(createAssetRequestResult.Msg)
	} else {
		fmt.Println("Check Create Asset success...")
	}
	
	createAssetObj,_ := createAssetRequestResult.Data. (model.Asset)

	// --------------------------------------------------------------
	// Check Asset Obj ID
	// --------------------------------------------------------------	
	if createAssetObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Asset" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAssetRequestResult := dao.GetAsset( uint64(createAssetObj.ID) )
	
	if getAssetRequestResult.Success == false {
		t.Errorf(getAssetRequestResult.Msg)
	} else {
		fmt.Println("Check Get Asset success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAssetObj,_ := getAssetRequestResult.Data. (model.Asset)
	compareAsset := cmp.Equal(createAssetObj.ID, getAssetObj.ID)
	
	if  compareAsset == false	{
		t.Errorf( "Created Asset object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAssetRequestResult := dao.GetAllAsset()

	if getAllAssetRequestResult.Success == false {
			t.Errorf(getAllAssetRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Asset success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAssetObj []model.Asset = getAllAssetRequestResult.Data. ([]model.Asset)
		
	equalAsset := cmp.Equal(createAssetObj.ID, getAllAssetObj[len(getAllAssetObj)-1].ID)
		
	if equalAsset == false {
		t.Errorf( "Created object is not equal to the last entry in Asset[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Asset
	// --------------------------------------------------------------	
	deleteAssetRequestResult := dao.DeleteAsset(uint64(createAssetObj.ID))

	if deleteAssetRequestResult.Success == false {
			t.Errorf(deleteAssetRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Asset success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAssetRequestResult = dao.GetAsset( uint64(createAssetObj.ID) )
	
	if getAssetRequestResult.Success == true {
		t.Errorf(getAssetRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMaintenancePlanCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MaintenancePlan
	//----------------------------------------------------------------------------
	MaintenancePlanObj := model.MaintenancePlan                                                                                                                    {PlanNumber:"test value for PlanNumber",Interval:new TimeDuration(),LastServiceDate:time.Now(),Strategy:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMaintenancePlanRequestResult := dao.CreateMaintenancePlan( MaintenancePlanObj )
	
	if createMaintenancePlanRequestResult.Success == false {
		t.Errorf(createMaintenancePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Create MaintenancePlan success...")
	}
	
	createMaintenancePlanObj,_ := createMaintenancePlanRequestResult.Data. (model.MaintenancePlan)

	// --------------------------------------------------------------
	// Check MaintenancePlan Obj ID
	// --------------------------------------------------------------	
	if createMaintenancePlanObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MaintenancePlan" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMaintenancePlanRequestResult := dao.GetMaintenancePlan( uint64(createMaintenancePlanObj.ID) )
	
	if getMaintenancePlanRequestResult.Success == false {
		t.Errorf(getMaintenancePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Get MaintenancePlan success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMaintenancePlanObj,_ := getMaintenancePlanRequestResult.Data. (model.MaintenancePlan)
	compareMaintenancePlan := cmp.Equal(createMaintenancePlanObj.ID, getMaintenancePlanObj.ID)
	
	if  compareMaintenancePlan == false	{
		t.Errorf( "Created MaintenancePlan object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMaintenancePlanRequestResult := dao.GetAllMaintenancePlan()

	if getAllMaintenancePlanRequestResult.Success == false {
			t.Errorf(getAllMaintenancePlanRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MaintenancePlan success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMaintenancePlanObj []model.MaintenancePlan = getAllMaintenancePlanRequestResult.Data. ([]model.MaintenancePlan)
		
	equalMaintenancePlan := cmp.Equal(createMaintenancePlanObj.ID, getAllMaintenancePlanObj[len(getAllMaintenancePlanObj)-1].ID)
		
	if equalMaintenancePlan == false {
		t.Errorf( "Created object is not equal to the last entry in MaintenancePlan[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MaintenancePlan
	// --------------------------------------------------------------	
	deleteMaintenancePlanRequestResult := dao.DeleteMaintenancePlan(uint64(createMaintenancePlanObj.ID))

	if deleteMaintenancePlanRequestResult.Success == false {
			t.Errorf(deleteMaintenancePlanRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MaintenancePlan success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMaintenancePlanRequestResult = dao.GetMaintenancePlan( uint64(createMaintenancePlanObj.ID) )
	
	if getMaintenancePlanRequestResult.Success == true {
		t.Errorf(getMaintenancePlanRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMaintenanceOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MaintenanceOrder
	//----------------------------------------------------------------------------
	MaintenanceOrderObj := model.MaintenanceOrder                                                                                                                                                                                            {OrderNumber:"test value for OrderNumber",Priority:100,RequestedDate:time.Now(),CompletionDate:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMaintenanceOrderRequestResult := dao.CreateMaintenanceOrder( MaintenanceOrderObj )
	
	if createMaintenanceOrderRequestResult.Success == false {
		t.Errorf(createMaintenanceOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create MaintenanceOrder success...")
	}
	
	createMaintenanceOrderObj,_ := createMaintenanceOrderRequestResult.Data. (model.MaintenanceOrder)

	// --------------------------------------------------------------
	// Check MaintenanceOrder Obj ID
	// --------------------------------------------------------------	
	if createMaintenanceOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MaintenanceOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMaintenanceOrderRequestResult := dao.GetMaintenanceOrder( uint64(createMaintenanceOrderObj.ID) )
	
	if getMaintenanceOrderRequestResult.Success == false {
		t.Errorf(getMaintenanceOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get MaintenanceOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMaintenanceOrderObj,_ := getMaintenanceOrderRequestResult.Data. (model.MaintenanceOrder)
	compareMaintenanceOrder := cmp.Equal(createMaintenanceOrderObj.ID, getMaintenanceOrderObj.ID)
	
	if  compareMaintenanceOrder == false	{
		t.Errorf( "Created MaintenanceOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMaintenanceOrderRequestResult := dao.GetAllMaintenanceOrder()

	if getAllMaintenanceOrderRequestResult.Success == false {
			t.Errorf(getAllMaintenanceOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MaintenanceOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMaintenanceOrderObj []model.MaintenanceOrder = getAllMaintenanceOrderRequestResult.Data. ([]model.MaintenanceOrder)
		
	equalMaintenanceOrder := cmp.Equal(createMaintenanceOrderObj.ID, getAllMaintenanceOrderObj[len(getAllMaintenanceOrderObj)-1].ID)
		
	if equalMaintenanceOrder == false {
		t.Errorf( "Created object is not equal to the last entry in MaintenanceOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MaintenanceOrder
	// --------------------------------------------------------------	
	deleteMaintenanceOrderRequestResult := dao.DeleteMaintenanceOrder(uint64(createMaintenanceOrderObj.ID))

	if deleteMaintenanceOrderRequestResult.Success == false {
			t.Errorf(deleteMaintenanceOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MaintenanceOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMaintenanceOrderRequestResult = dao.GetMaintenanceOrder( uint64(createMaintenanceOrderObj.ID) )
	
	if getMaintenanceOrderRequestResult.Success == true {
		t.Errorf(getMaintenanceOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEmployeeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Employee
	//----------------------------------------------------------------------------
	EmployeeObj := model.Employee                                                                                            {FirstName:"test value for FirstName",LastName:"test value for LastName",Role:0,SkillLevel:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEmployeeRequestResult := dao.CreateEmployee( EmployeeObj )
	
	if createEmployeeRequestResult.Success == false {
		t.Errorf(createEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Check Create Employee success...")
	}
	
	createEmployeeObj,_ := createEmployeeRequestResult.Data. (model.Employee)

	// --------------------------------------------------------------
	// Check Employee Obj ID
	// --------------------------------------------------------------	
	if createEmployeeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Employee" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEmployeeRequestResult := dao.GetEmployee( uint64(createEmployeeObj.ID) )
	
	if getEmployeeRequestResult.Success == false {
		t.Errorf(getEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Check Get Employee success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEmployeeObj,_ := getEmployeeRequestResult.Data. (model.Employee)
	compareEmployee := cmp.Equal(createEmployeeObj.ID, getEmployeeObj.ID)
	
	if  compareEmployee == false	{
		t.Errorf( "Created Employee object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEmployeeRequestResult := dao.GetAllEmployee()

	if getAllEmployeeRequestResult.Success == false {
			t.Errorf(getAllEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Employee success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEmployeeObj []model.Employee = getAllEmployeeRequestResult.Data. ([]model.Employee)
		
	equalEmployee := cmp.Equal(createEmployeeObj.ID, getAllEmployeeObj[len(getAllEmployeeObj)-1].ID)
		
	if equalEmployee == false {
		t.Errorf( "Created object is not equal to the last entry in Employee[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Employee
	// --------------------------------------------------------------	
	deleteEmployeeRequestResult := dao.DeleteEmployee(uint64(createEmployeeObj.ID))

	if deleteEmployeeRequestResult.Success == false {
			t.Errorf(deleteEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Employee success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEmployeeRequestResult = dao.GetEmployee( uint64(createEmployeeObj.ID) )
	
	if getEmployeeRequestResult.Success == true {
		t.Errorf(getEmployeeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestShiftCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Shift
	//----------------------------------------------------------------------------
	ShiftObj := model.Shift                                                                                                            {ShiftName:"test value for ShiftName",StartTime:"test value for StartTime",EndTime:"test value for EndTime",ShiftType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createShiftRequestResult := dao.CreateShift( ShiftObj )
	
	if createShiftRequestResult.Success == false {
		t.Errorf(createShiftRequestResult.Msg)
	} else {
		fmt.Println("Check Create Shift success...")
	}
	
	createShiftObj,_ := createShiftRequestResult.Data. (model.Shift)

	// --------------------------------------------------------------
	// Check Shift Obj ID
	// --------------------------------------------------------------	
	if createShiftObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Shift" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getShiftRequestResult := dao.GetShift( uint64(createShiftObj.ID) )
	
	if getShiftRequestResult.Success == false {
		t.Errorf(getShiftRequestResult.Msg)
	} else {
		fmt.Println("Check Get Shift success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getShiftObj,_ := getShiftRequestResult.Data. (model.Shift)
	compareShift := cmp.Equal(createShiftObj.ID, getShiftObj.ID)
	
	if  compareShift == false	{
		t.Errorf( "Created Shift object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllShiftRequestResult := dao.GetAllShift()

	if getAllShiftRequestResult.Success == false {
			t.Errorf(getAllShiftRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Shift success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllShiftObj []model.Shift = getAllShiftRequestResult.Data. ([]model.Shift)
		
	equalShift := cmp.Equal(createShiftObj.ID, getAllShiftObj[len(getAllShiftObj)-1].ID)
		
	if equalShift == false {
		t.Errorf( "Created object is not equal to the last entry in Shift[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Shift
	// --------------------------------------------------------------	
	deleteShiftRequestResult := dao.DeleteShift(uint64(createShiftObj.ID))

	if deleteShiftRequestResult.Success == false {
			t.Errorf(deleteShiftRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Shift success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getShiftRequestResult = dao.GetShift( uint64(createShiftObj.ID) )
	
	if getShiftRequestResult.Success == true {
		t.Errorf(getShiftRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestShiftAssignmentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ShiftAssignment
	//----------------------------------------------------------------------------
	ShiftAssignmentObj := model.ShiftAssignment                                                    {AssignmentDate:time.Now()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createShiftAssignmentRequestResult := dao.CreateShiftAssignment( ShiftAssignmentObj )
	
	if createShiftAssignmentRequestResult.Success == false {
		t.Errorf(createShiftAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Create ShiftAssignment success...")
	}
	
	createShiftAssignmentObj,_ := createShiftAssignmentRequestResult.Data. (model.ShiftAssignment)

	// --------------------------------------------------------------
	// Check ShiftAssignment Obj ID
	// --------------------------------------------------------------	
	if createShiftAssignmentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ShiftAssignment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getShiftAssignmentRequestResult := dao.GetShiftAssignment( uint64(createShiftAssignmentObj.ID) )
	
	if getShiftAssignmentRequestResult.Success == false {
		t.Errorf(getShiftAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Get ShiftAssignment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getShiftAssignmentObj,_ := getShiftAssignmentRequestResult.Data. (model.ShiftAssignment)
	compareShiftAssignment := cmp.Equal(createShiftAssignmentObj.ID, getShiftAssignmentObj.ID)
	
	if  compareShiftAssignment == false	{
		t.Errorf( "Created ShiftAssignment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllShiftAssignmentRequestResult := dao.GetAllShiftAssignment()

	if getAllShiftAssignmentRequestResult.Success == false {
			t.Errorf(getAllShiftAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ShiftAssignment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllShiftAssignmentObj []model.ShiftAssignment = getAllShiftAssignmentRequestResult.Data. ([]model.ShiftAssignment)
		
	equalShiftAssignment := cmp.Equal(createShiftAssignmentObj.ID, getAllShiftAssignmentObj[len(getAllShiftAssignmentObj)-1].ID)
		
	if equalShiftAssignment == false {
		t.Errorf( "Created object is not equal to the last entry in ShiftAssignment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ShiftAssignment
	// --------------------------------------------------------------	
	deleteShiftAssignmentRequestResult := dao.DeleteShiftAssignment(uint64(createShiftAssignmentObj.ID))

	if deleteShiftAssignmentRequestResult.Success == false {
			t.Errorf(deleteShiftAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ShiftAssignment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getShiftAssignmentRequestResult = dao.GetShiftAssignment( uint64(createShiftAssignmentObj.ID) )
	
	if getShiftAssignmentRequestResult.Success == true {
		t.Errorf(getShiftAssignmentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestForecastCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Forecast
	//----------------------------------------------------------------------------
	ForecastObj := model.Forecast                                                                                                                                                            {ForecastNumber:"test value for ForecastNumber",ForecastHorizonStart:time.Now(),ForecastHorizonEnd:time.Now(),Method:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createForecastRequestResult := dao.CreateForecast( ForecastObj )
	
	if createForecastRequestResult.Success == false {
		t.Errorf(createForecastRequestResult.Msg)
	} else {
		fmt.Println("Check Create Forecast success...")
	}
	
	createForecastObj,_ := createForecastRequestResult.Data. (model.Forecast)

	// --------------------------------------------------------------
	// Check Forecast Obj ID
	// --------------------------------------------------------------	
	if createForecastObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Forecast" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getForecastRequestResult := dao.GetForecast( uint64(createForecastObj.ID) )
	
	if getForecastRequestResult.Success == false {
		t.Errorf(getForecastRequestResult.Msg)
	} else {
		fmt.Println("Check Get Forecast success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getForecastObj,_ := getForecastRequestResult.Data. (model.Forecast)
	compareForecast := cmp.Equal(createForecastObj.ID, getForecastObj.ID)
	
	if  compareForecast == false	{
		t.Errorf( "Created Forecast object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllForecastRequestResult := dao.GetAllForecast()

	if getAllForecastRequestResult.Success == false {
			t.Errorf(getAllForecastRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Forecast success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllForecastObj []model.Forecast = getAllForecastRequestResult.Data. ([]model.Forecast)
		
	equalForecast := cmp.Equal(createForecastObj.ID, getAllForecastObj[len(getAllForecastObj)-1].ID)
		
	if equalForecast == false {
		t.Errorf( "Created object is not equal to the last entry in Forecast[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Forecast
	// --------------------------------------------------------------	
	deleteForecastRequestResult := dao.DeleteForecast(uint64(createForecastObj.ID))

	if deleteForecastRequestResult.Success == false {
			t.Errorf(deleteForecastRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Forecast success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getForecastRequestResult = dao.GetForecast( uint64(createForecastObj.ID) )
	
	if getForecastRequestResult.Success == true {
		t.Errorf(getForecastRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestForecastLineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ForecastLine
	//----------------------------------------------------------------------------
	ForecastLineObj := model.ForecastLine                                                                                    {Period:time.Now(),Quantity:new Quantity(),Confidence:new Percentage()}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createForecastLineRequestResult := dao.CreateForecastLine( ForecastLineObj )
	
	if createForecastLineRequestResult.Success == false {
		t.Errorf(createForecastLineRequestResult.Msg)
	} else {
		fmt.Println("Check Create ForecastLine success...")
	}
	
	createForecastLineObj,_ := createForecastLineRequestResult.Data. (model.ForecastLine)

	// --------------------------------------------------------------
	// Check ForecastLine Obj ID
	// --------------------------------------------------------------	
	if createForecastLineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ForecastLine" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getForecastLineRequestResult := dao.GetForecastLine( uint64(createForecastLineObj.ID) )
	
	if getForecastLineRequestResult.Success == false {
		t.Errorf(getForecastLineRequestResult.Msg)
	} else {
		fmt.Println("Check Get ForecastLine success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getForecastLineObj,_ := getForecastLineRequestResult.Data. (model.ForecastLine)
	compareForecastLine := cmp.Equal(createForecastLineObj.ID, getForecastLineObj.ID)
	
	if  compareForecastLine == false	{
		t.Errorf( "Created ForecastLine object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllForecastLineRequestResult := dao.GetAllForecastLine()

	if getAllForecastLineRequestResult.Success == false {
			t.Errorf(getAllForecastLineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ForecastLine success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllForecastLineObj []model.ForecastLine = getAllForecastLineRequestResult.Data. ([]model.ForecastLine)
		
	equalForecastLine := cmp.Equal(createForecastLineObj.ID, getAllForecastLineObj[len(getAllForecastLineObj)-1].ID)
		
	if equalForecastLine == false {
		t.Errorf( "Created object is not equal to the last entry in ForecastLine[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ForecastLine
	// --------------------------------------------------------------	
	deleteForecastLineRequestResult := dao.DeleteForecastLine(uint64(createForecastLineObj.ID))

	if deleteForecastLineRequestResult.Success == false {
			t.Errorf(deleteForecastLineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ForecastLine success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getForecastLineRequestResult = dao.GetForecastLine( uint64(createForecastLineObj.ID) )
	
	if getForecastLineRequestResult.Success == true {
		t.Errorf(getForecastLineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMRPRunCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for MRPRun
	//----------------------------------------------------------------------------
	MRPRunObj := model.MRPRun                                                                                                                                    {RunNumber:"test value for RunNumber",RunDateTime:time.Now(),PlanningHorizonDays:100,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMRPRunRequestResult := dao.CreateMRPRun( MRPRunObj )
	
	if createMRPRunRequestResult.Success == false {
		t.Errorf(createMRPRunRequestResult.Msg)
	} else {
		fmt.Println("Check Create MRPRun success...")
	}
	
	createMRPRunObj,_ := createMRPRunRequestResult.Data. (model.MRPRun)

	// --------------------------------------------------------------
	// Check MRPRun Obj ID
	// --------------------------------------------------------------	
	if createMRPRunObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for MRPRun" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMRPRunRequestResult := dao.GetMRPRun( uint64(createMRPRunObj.ID) )
	
	if getMRPRunRequestResult.Success == false {
		t.Errorf(getMRPRunRequestResult.Msg)
	} else {
		fmt.Println("Check Get MRPRun success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMRPRunObj,_ := getMRPRunRequestResult.Data. (model.MRPRun)
	compareMRPRun := cmp.Equal(createMRPRunObj.ID, getMRPRunObj.ID)
	
	if  compareMRPRun == false	{
		t.Errorf( "Created MRPRun object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMRPRunRequestResult := dao.GetAllMRPRun()

	if getAllMRPRunRequestResult.Success == false {
			t.Errorf(getAllMRPRunRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll MRPRun success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMRPRunObj []model.MRPRun = getAllMRPRunRequestResult.Data. ([]model.MRPRun)
		
	equalMRPRun := cmp.Equal(createMRPRunObj.ID, getAllMRPRunObj[len(getAllMRPRunObj)-1].ID)
		
	if equalMRPRun == false {
		t.Errorf( "Created object is not equal to the last entry in MRPRun[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for MRPRun
	// --------------------------------------------------------------	
	deleteMRPRunRequestResult := dao.DeleteMRPRun(uint64(createMRPRunObj.ID))

	if deleteMRPRunRequestResult.Success == false {
			t.Errorf(deleteMRPRunRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion MRPRun success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMRPRunRequestResult = dao.GetMRPRun( uint64(createMRPRunObj.ID) )
	
	if getMRPRunRequestResult.Success == true {
		t.Errorf(getMRPRunRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPlannedOrderCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for PlannedOrder
	//----------------------------------------------------------------------------
	PlannedOrderObj := model.PlannedOrder                                                                                                                                    {PlannedOrderNumber:"test value for PlannedOrderNumber",Quantity:new Quantity(),DueDate:time.Now(),OrderType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPlannedOrderRequestResult := dao.CreatePlannedOrder( PlannedOrderObj )
	
	if createPlannedOrderRequestResult.Success == false {
		t.Errorf(createPlannedOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Create PlannedOrder success...")
	}
	
	createPlannedOrderObj,_ := createPlannedOrderRequestResult.Data. (model.PlannedOrder)

	// --------------------------------------------------------------
	// Check PlannedOrder Obj ID
	// --------------------------------------------------------------	
	if createPlannedOrderObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for PlannedOrder" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPlannedOrderRequestResult := dao.GetPlannedOrder( uint64(createPlannedOrderObj.ID) )
	
	if getPlannedOrderRequestResult.Success == false {
		t.Errorf(getPlannedOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Get PlannedOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPlannedOrderObj,_ := getPlannedOrderRequestResult.Data. (model.PlannedOrder)
	comparePlannedOrder := cmp.Equal(createPlannedOrderObj.ID, getPlannedOrderObj.ID)
	
	if  comparePlannedOrder == false	{
		t.Errorf( "Created PlannedOrder object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPlannedOrderRequestResult := dao.GetAllPlannedOrder()

	if getAllPlannedOrderRequestResult.Success == false {
			t.Errorf(getAllPlannedOrderRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll PlannedOrder success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPlannedOrderObj []model.PlannedOrder = getAllPlannedOrderRequestResult.Data. ([]model.PlannedOrder)
		
	equalPlannedOrder := cmp.Equal(createPlannedOrderObj.ID, getAllPlannedOrderObj[len(getAllPlannedOrderObj)-1].ID)
		
	if equalPlannedOrder == false {
		t.Errorf( "Created object is not equal to the last entry in PlannedOrder[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for PlannedOrder
	// --------------------------------------------------------------	
	deletePlannedOrderRequestResult := dao.DeletePlannedOrder(uint64(createPlannedOrderObj.ID))

	if deletePlannedOrderRequestResult.Success == false {
			t.Errorf(deletePlannedOrderRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion PlannedOrder success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPlannedOrderRequestResult = dao.GetPlannedOrder( uint64(createPlannedOrderObj.ID) )
	
	if getPlannedOrderRequestResult.Success == true {
		t.Errorf(getPlannedOrderRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

