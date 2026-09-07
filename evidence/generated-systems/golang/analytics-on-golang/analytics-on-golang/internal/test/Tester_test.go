package test

import ( 
	"testing"
    dao "analytics-on-golang/internal/dao"
	"analytics-on-golang/internal/model"
	"analytics-on-golang/internal/utils"
	"github.com/google/go-cmp/cmp"
	"fmt"
)

func init() {
	utils.InitializeEnvironment()
}


func TestAnalyticsWorkspaceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AnalyticsWorkspace
	//----------------------------------------------------------------------------
	AnalyticsWorkspaceObj := model.AnalyticsWorkspace                                                                                                            {Name:"test value for Name",BusinessDomain:"test value for BusinessDomain",OwnerTeam:"test value for OwnerTeam",GovernanceTier:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAnalyticsWorkspaceRequestResult := dao.CreateAnalyticsWorkspace( AnalyticsWorkspaceObj )
	
	if createAnalyticsWorkspaceRequestResult.Success == false {
		t.Errorf(createAnalyticsWorkspaceRequestResult.Msg)
	} else {
		fmt.Println("Check Create AnalyticsWorkspace success...")
	}
	
	createAnalyticsWorkspaceObj,_ := createAnalyticsWorkspaceRequestResult.Data. (model.AnalyticsWorkspace)

	// --------------------------------------------------------------
	// Check AnalyticsWorkspace Obj ID
	// --------------------------------------------------------------	
	if createAnalyticsWorkspaceObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AnalyticsWorkspace" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAnalyticsWorkspaceRequestResult := dao.GetAnalyticsWorkspace( uint64(createAnalyticsWorkspaceObj.ID) )
	
	if getAnalyticsWorkspaceRequestResult.Success == false {
		t.Errorf(getAnalyticsWorkspaceRequestResult.Msg)
	} else {
		fmt.Println("Check Get AnalyticsWorkspace success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAnalyticsWorkspaceObj,_ := getAnalyticsWorkspaceRequestResult.Data. (model.AnalyticsWorkspace)
	compareAnalyticsWorkspace := cmp.Equal(createAnalyticsWorkspaceObj.ID, getAnalyticsWorkspaceObj.ID)
	
	if  compareAnalyticsWorkspace == false	{
		t.Errorf( "Created AnalyticsWorkspace object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAnalyticsWorkspaceRequestResult := dao.GetAllAnalyticsWorkspace()

	if getAllAnalyticsWorkspaceRequestResult.Success == false {
			t.Errorf(getAllAnalyticsWorkspaceRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AnalyticsWorkspace success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAnalyticsWorkspaceObj []model.AnalyticsWorkspace = getAllAnalyticsWorkspaceRequestResult.Data. ([]model.AnalyticsWorkspace)
		
	equalAnalyticsWorkspace := cmp.Equal(createAnalyticsWorkspaceObj.ID, getAllAnalyticsWorkspaceObj[len(getAllAnalyticsWorkspaceObj)-1].ID)
		
	if equalAnalyticsWorkspace == false {
		t.Errorf( "Created object is not equal to the last entry in AnalyticsWorkspace[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AnalyticsWorkspace
	// --------------------------------------------------------------	
	deleteAnalyticsWorkspaceRequestResult := dao.DeleteAnalyticsWorkspace(uint64(createAnalyticsWorkspaceObj.ID))

	if deleteAnalyticsWorkspaceRequestResult.Success == false {
			t.Errorf(deleteAnalyticsWorkspaceRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AnalyticsWorkspace success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAnalyticsWorkspaceRequestResult = dao.GetAnalyticsWorkspace( uint64(createAnalyticsWorkspaceObj.ID) )
	
	if getAnalyticsWorkspaceRequestResult.Success == true {
		t.Errorf(getAnalyticsWorkspaceRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDataSourceCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataSource
	//----------------------------------------------------------------------------
	DataSourceObj := model.DataSource                                                                                                            {Name:"test value for Name",Connection:new ConnectionInfo(),Streaming:true,SourceType:0,Format:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataSourceRequestResult := dao.CreateDataSource( DataSourceObj )
	
	if createDataSourceRequestResult.Success == false {
		t.Errorf(createDataSourceRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataSource success...")
	}
	
	createDataSourceObj,_ := createDataSourceRequestResult.Data. (model.DataSource)

	// --------------------------------------------------------------
	// Check DataSource Obj ID
	// --------------------------------------------------------------	
	if createDataSourceObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataSource" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataSourceRequestResult := dao.GetDataSource( uint64(createDataSourceObj.ID) )
	
	if getDataSourceRequestResult.Success == false {
		t.Errorf(getDataSourceRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataSource success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataSourceObj,_ := getDataSourceRequestResult.Data. (model.DataSource)
	compareDataSource := cmp.Equal(createDataSourceObj.ID, getDataSourceObj.ID)
	
	if  compareDataSource == false	{
		t.Errorf( "Created DataSource object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataSourceRequestResult := dao.GetAllDataSource()

	if getAllDataSourceRequestResult.Success == false {
			t.Errorf(getAllDataSourceRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataSource success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataSourceObj []model.DataSource = getAllDataSourceRequestResult.Data. ([]model.DataSource)
		
	equalDataSource := cmp.Equal(createDataSourceObj.ID, getAllDataSourceObj[len(getAllDataSourceObj)-1].ID)
		
	if equalDataSource == false {
		t.Errorf( "Created object is not equal to the last entry in DataSource[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataSource
	// --------------------------------------------------------------	
	deleteDataSourceRequestResult := dao.DeleteDataSource(uint64(createDataSourceObj.ID))

	if deleteDataSourceRequestResult.Success == false {
			t.Errorf(deleteDataSourceRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataSource success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataSourceRequestResult = dao.GetDataSource( uint64(createDataSourceObj.ID) )
	
	if getDataSourceRequestResult.Success == true {
		t.Errorf(getDataSourceRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDataSetCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataSet
	//----------------------------------------------------------------------------
	DataSetObj := model.DataSet                                                                                                                            {Name:"test value for Name",SchemaVersion:"test value for SchemaVersion",RefreshSchedule:new CronSchedule(),Sensitive:true,DataFormat:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataSetRequestResult := dao.CreateDataSet( DataSetObj )
	
	if createDataSetRequestResult.Success == false {
		t.Errorf(createDataSetRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataSet success...")
	}
	
	createDataSetObj,_ := createDataSetRequestResult.Data. (model.DataSet)

	// --------------------------------------------------------------
	// Check DataSet Obj ID
	// --------------------------------------------------------------	
	if createDataSetObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataSet" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataSetRequestResult := dao.GetDataSet( uint64(createDataSetObj.ID) )
	
	if getDataSetRequestResult.Success == false {
		t.Errorf(getDataSetRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataSet success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataSetObj,_ := getDataSetRequestResult.Data. (model.DataSet)
	compareDataSet := cmp.Equal(createDataSetObj.ID, getDataSetObj.ID)
	
	if  compareDataSet == false	{
		t.Errorf( "Created DataSet object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataSetRequestResult := dao.GetAllDataSet()

	if getAllDataSetRequestResult.Success == false {
			t.Errorf(getAllDataSetRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataSet success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataSetObj []model.DataSet = getAllDataSetRequestResult.Data. ([]model.DataSet)
		
	equalDataSet := cmp.Equal(createDataSetObj.ID, getAllDataSetObj[len(getAllDataSetObj)-1].ID)
		
	if equalDataSet == false {
		t.Errorf( "Created object is not equal to the last entry in DataSet[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataSet
	// --------------------------------------------------------------	
	deleteDataSetRequestResult := dao.DeleteDataSet(uint64(createDataSetObj.ID))

	if deleteDataSetRequestResult.Success == false {
			t.Errorf(deleteDataSetRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataSet success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataSetRequestResult = dao.GetDataSet( uint64(createDataSetObj.ID) )
	
	if getDataSetRequestResult.Success == true {
		t.Errorf(getDataSetRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDataPipelineCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataPipeline
	//----------------------------------------------------------------------------
	DataPipelineObj := model.DataPipeline                                                                            {Name:"test value for Name",Schedule:new CronSchedule(),TriggerType:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataPipelineRequestResult := dao.CreateDataPipeline( DataPipelineObj )
	
	if createDataPipelineRequestResult.Success == false {
		t.Errorf(createDataPipelineRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataPipeline success...")
	}
	
	createDataPipelineObj,_ := createDataPipelineRequestResult.Data. (model.DataPipeline)

	// --------------------------------------------------------------
	// Check DataPipeline Obj ID
	// --------------------------------------------------------------	
	if createDataPipelineObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataPipeline" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataPipelineRequestResult := dao.GetDataPipeline( uint64(createDataPipelineObj.ID) )
	
	if getDataPipelineRequestResult.Success == false {
		t.Errorf(getDataPipelineRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataPipeline success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataPipelineObj,_ := getDataPipelineRequestResult.Data. (model.DataPipeline)
	compareDataPipeline := cmp.Equal(createDataPipelineObj.ID, getDataPipelineObj.ID)
	
	if  compareDataPipeline == false	{
		t.Errorf( "Created DataPipeline object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataPipelineRequestResult := dao.GetAllDataPipeline()

	if getAllDataPipelineRequestResult.Success == false {
			t.Errorf(getAllDataPipelineRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataPipeline success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataPipelineObj []model.DataPipeline = getAllDataPipelineRequestResult.Data. ([]model.DataPipeline)
		
	equalDataPipeline := cmp.Equal(createDataPipelineObj.ID, getAllDataPipelineObj[len(getAllDataPipelineObj)-1].ID)
		
	if equalDataPipeline == false {
		t.Errorf( "Created object is not equal to the last entry in DataPipeline[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataPipeline
	// --------------------------------------------------------------	
	deleteDataPipelineRequestResult := dao.DeleteDataPipeline(uint64(createDataPipelineObj.ID))

	if deleteDataPipelineRequestResult.Success == false {
			t.Errorf(deleteDataPipelineRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataPipeline success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataPipelineRequestResult = dao.GetDataPipeline( uint64(createDataPipelineObj.ID) )
	
	if getDataPipelineRequestResult.Success == true {
		t.Errorf(getDataPipelineRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDataTaskCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for DataTask
	//----------------------------------------------------------------------------
	DataTaskObj := model.DataTask                                                                                                            {Name:"test value for Name",Command:"test value for Command",Retries:100,TaskType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDataTaskRequestResult := dao.CreateDataTask( DataTaskObj )
	
	if createDataTaskRequestResult.Success == false {
		t.Errorf(createDataTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Create DataTask success...")
	}
	
	createDataTaskObj,_ := createDataTaskRequestResult.Data. (model.DataTask)

	// --------------------------------------------------------------
	// Check DataTask Obj ID
	// --------------------------------------------------------------	
	if createDataTaskObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for DataTask" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDataTaskRequestResult := dao.GetDataTask( uint64(createDataTaskObj.ID) )
	
	if getDataTaskRequestResult.Success == false {
		t.Errorf(getDataTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Get DataTask success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDataTaskObj,_ := getDataTaskRequestResult.Data. (model.DataTask)
	compareDataTask := cmp.Equal(createDataTaskObj.ID, getDataTaskObj.ID)
	
	if  compareDataTask == false	{
		t.Errorf( "Created DataTask object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDataTaskRequestResult := dao.GetAllDataTask()

	if getAllDataTaskRequestResult.Success == false {
			t.Errorf(getAllDataTaskRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll DataTask success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDataTaskObj []model.DataTask = getAllDataTaskRequestResult.Data. ([]model.DataTask)
		
	equalDataTask := cmp.Equal(createDataTaskObj.ID, getAllDataTaskObj[len(getAllDataTaskObj)-1].ID)
		
	if equalDataTask == false {
		t.Errorf( "Created object is not equal to the last entry in DataTask[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for DataTask
	// --------------------------------------------------------------	
	deleteDataTaskRequestResult := dao.DeleteDataTask(uint64(createDataTaskObj.ID))

	if deleteDataTaskRequestResult.Success == false {
			t.Errorf(deleteDataTaskRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion DataTask success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDataTaskRequestResult = dao.GetDataTask( uint64(createDataTaskObj.ID) )
	
	if getDataTaskRequestResult.Success == true {
		t.Errorf(getDataTaskRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSemanticModelCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for SemanticModel
	//----------------------------------------------------------------------------
	SemanticModelObj := model.SemanticModel                                                                                            {Name:"test value for Name",Version:"test value for Version",Grain:"test value for Grain"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSemanticModelRequestResult := dao.CreateSemanticModel( SemanticModelObj )
	
	if createSemanticModelRequestResult.Success == false {
		t.Errorf(createSemanticModelRequestResult.Msg)
	} else {
		fmt.Println("Check Create SemanticModel success...")
	}
	
	createSemanticModelObj,_ := createSemanticModelRequestResult.Data. (model.SemanticModel)

	// --------------------------------------------------------------
	// Check SemanticModel Obj ID
	// --------------------------------------------------------------	
	if createSemanticModelObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for SemanticModel" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSemanticModelRequestResult := dao.GetSemanticModel( uint64(createSemanticModelObj.ID) )
	
	if getSemanticModelRequestResult.Success == false {
		t.Errorf(getSemanticModelRequestResult.Msg)
	} else {
		fmt.Println("Check Get SemanticModel success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSemanticModelObj,_ := getSemanticModelRequestResult.Data. (model.SemanticModel)
	compareSemanticModel := cmp.Equal(createSemanticModelObj.ID, getSemanticModelObj.ID)
	
	if  compareSemanticModel == false	{
		t.Errorf( "Created SemanticModel object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSemanticModelRequestResult := dao.GetAllSemanticModel()

	if getAllSemanticModelRequestResult.Success == false {
			t.Errorf(getAllSemanticModelRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll SemanticModel success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSemanticModelObj []model.SemanticModel = getAllSemanticModelRequestResult.Data. ([]model.SemanticModel)
		
	equalSemanticModel := cmp.Equal(createSemanticModelObj.ID, getAllSemanticModelObj[len(getAllSemanticModelObj)-1].ID)
		
	if equalSemanticModel == false {
		t.Errorf( "Created object is not equal to the last entry in SemanticModel[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for SemanticModel
	// --------------------------------------------------------------	
	deleteSemanticModelRequestResult := dao.DeleteSemanticModel(uint64(createSemanticModelObj.ID))

	if deleteSemanticModelRequestResult.Success == false {
			t.Errorf(deleteSemanticModelRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion SemanticModel success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSemanticModelRequestResult = dao.GetSemanticModel( uint64(createSemanticModelObj.ID) )
	
	if getSemanticModelRequestResult.Success == true {
		t.Errorf(getSemanticModelRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDimensionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Dimension
	//----------------------------------------------------------------------------
	DimensionObj := model.Dimension                                                                            {Name:"test value for Name",TypeTime:true,DimensionType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDimensionRequestResult := dao.CreateDimension( DimensionObj )
	
	if createDimensionRequestResult.Success == false {
		t.Errorf(createDimensionRequestResult.Msg)
	} else {
		fmt.Println("Check Create Dimension success...")
	}
	
	createDimensionObj,_ := createDimensionRequestResult.Data. (model.Dimension)

	// --------------------------------------------------------------
	// Check Dimension Obj ID
	// --------------------------------------------------------------	
	if createDimensionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Dimension" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDimensionRequestResult := dao.GetDimension( uint64(createDimensionObj.ID) )
	
	if getDimensionRequestResult.Success == false {
		t.Errorf(getDimensionRequestResult.Msg)
	} else {
		fmt.Println("Check Get Dimension success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDimensionObj,_ := getDimensionRequestResult.Data. (model.Dimension)
	compareDimension := cmp.Equal(createDimensionObj.ID, getDimensionObj.ID)
	
	if  compareDimension == false	{
		t.Errorf( "Created Dimension object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDimensionRequestResult := dao.GetAllDimension()

	if getAllDimensionRequestResult.Success == false {
			t.Errorf(getAllDimensionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Dimension success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDimensionObj []model.Dimension = getAllDimensionRequestResult.Data. ([]model.Dimension)
		
	equalDimension := cmp.Equal(createDimensionObj.ID, getAllDimensionObj[len(getAllDimensionObj)-1].ID)
		
	if equalDimension == false {
		t.Errorf( "Created object is not equal to the last entry in Dimension[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Dimension
	// --------------------------------------------------------------	
	deleteDimensionRequestResult := dao.DeleteDimension(uint64(createDimensionObj.ID))

	if deleteDimensionRequestResult.Success == false {
			t.Errorf(deleteDimensionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Dimension success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDimensionRequestResult = dao.GetDimension( uint64(createDimensionObj.ID) )
	
	if getDimensionRequestResult.Success == true {
		t.Errorf(getDimensionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMeasureCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Measure
	//----------------------------------------------------------------------------
	MeasureObj := model.Measure                                                                            {Name:"test value for Name",Format:"test value for Format",Aggregation:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMeasureRequestResult := dao.CreateMeasure( MeasureObj )
	
	if createMeasureRequestResult.Success == false {
		t.Errorf(createMeasureRequestResult.Msg)
	} else {
		fmt.Println("Check Create Measure success...")
	}
	
	createMeasureObj,_ := createMeasureRequestResult.Data. (model.Measure)

	// --------------------------------------------------------------
	// Check Measure Obj ID
	// --------------------------------------------------------------	
	if createMeasureObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Measure" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMeasureRequestResult := dao.GetMeasure( uint64(createMeasureObj.ID) )
	
	if getMeasureRequestResult.Success == false {
		t.Errorf(getMeasureRequestResult.Msg)
	} else {
		fmt.Println("Check Get Measure success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMeasureObj,_ := getMeasureRequestResult.Data. (model.Measure)
	compareMeasure := cmp.Equal(createMeasureObj.ID, getMeasureObj.ID)
	
	if  compareMeasure == false	{
		t.Errorf( "Created Measure object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMeasureRequestResult := dao.GetAllMeasure()

	if getAllMeasureRequestResult.Success == false {
			t.Errorf(getAllMeasureRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Measure success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMeasureObj []model.Measure = getAllMeasureRequestResult.Data. ([]model.Measure)
		
	equalMeasure := cmp.Equal(createMeasureObj.ID, getAllMeasureObj[len(getAllMeasureObj)-1].ID)
		
	if equalMeasure == false {
		t.Errorf( "Created object is not equal to the last entry in Measure[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Measure
	// --------------------------------------------------------------	
	deleteMeasureRequestResult := dao.DeleteMeasure(uint64(createMeasureObj.ID))

	if deleteMeasureRequestResult.Success == false {
			t.Errorf(deleteMeasureRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Measure success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMeasureRequestResult = dao.GetMeasure( uint64(createMeasureObj.ID) )
	
	if getMeasureRequestResult.Success == true {
		t.Errorf(getMeasureRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestMetricCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Metric
	//----------------------------------------------------------------------------
	MetricObj := model.Metric                                                                                                            {Name:"test value for Name",Expression:"test value for Expression",Unit:"test value for Unit",MetricType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createMetricRequestResult := dao.CreateMetric( MetricObj )
	
	if createMetricRequestResult.Success == false {
		t.Errorf(createMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Create Metric success...")
	}
	
	createMetricObj,_ := createMetricRequestResult.Data. (model.Metric)

	// --------------------------------------------------------------
	// Check Metric Obj ID
	// --------------------------------------------------------------	
	if createMetricObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Metric" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getMetricRequestResult := dao.GetMetric( uint64(createMetricObj.ID) )
	
	if getMetricRequestResult.Success == false {
		t.Errorf(getMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Get Metric success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getMetricObj,_ := getMetricRequestResult.Data. (model.Metric)
	compareMetric := cmp.Equal(createMetricObj.ID, getMetricObj.ID)
	
	if  compareMetric == false	{
		t.Errorf( "Created Metric object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllMetricRequestResult := dao.GetAllMetric()

	if getAllMetricRequestResult.Success == false {
			t.Errorf(getAllMetricRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Metric success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllMetricObj []model.Metric = getAllMetricRequestResult.Data. ([]model.Metric)
		
	equalMetric := cmp.Equal(createMetricObj.ID, getAllMetricObj[len(getAllMetricObj)-1].ID)
		
	if equalMetric == false {
		t.Errorf( "Created object is not equal to the last entry in Metric[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Metric
	// --------------------------------------------------------------	
	deleteMetricRequestResult := dao.DeleteMetric(uint64(createMetricObj.ID))

	if deleteMetricRequestResult.Success == false {
			t.Errorf(deleteMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Metric success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getMetricRequestResult = dao.GetMetric( uint64(createMetricObj.ID) )
	
	if getMetricRequestResult.Success == true {
		t.Errorf(getMetricRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestReportCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Report
	//----------------------------------------------------------------------------
	ReportObj := model.Report                                                                            {Title:"test value for Title",Audience:"test value for Audience",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createReportRequestResult := dao.CreateReport( ReportObj )
	
	if createReportRequestResult.Success == false {
		t.Errorf(createReportRequestResult.Msg)
	} else {
		fmt.Println("Check Create Report success...")
	}
	
	createReportObj,_ := createReportRequestResult.Data. (model.Report)

	// --------------------------------------------------------------
	// Check Report Obj ID
	// --------------------------------------------------------------	
	if createReportObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Report" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getReportRequestResult := dao.GetReport( uint64(createReportObj.ID) )
	
	if getReportRequestResult.Success == false {
		t.Errorf(getReportRequestResult.Msg)
	} else {
		fmt.Println("Check Get Report success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getReportObj,_ := getReportRequestResult.Data. (model.Report)
	compareReport := cmp.Equal(createReportObj.ID, getReportObj.ID)
	
	if  compareReport == false	{
		t.Errorf( "Created Report object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllReportRequestResult := dao.GetAllReport()

	if getAllReportRequestResult.Success == false {
			t.Errorf(getAllReportRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Report success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllReportObj []model.Report = getAllReportRequestResult.Data. ([]model.Report)
		
	equalReport := cmp.Equal(createReportObj.ID, getAllReportObj[len(getAllReportObj)-1].ID)
		
	if equalReport == false {
		t.Errorf( "Created object is not equal to the last entry in Report[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Report
	// --------------------------------------------------------------	
	deleteReportRequestResult := dao.DeleteReport(uint64(createReportObj.ID))

	if deleteReportRequestResult.Success == false {
			t.Errorf(deleteReportRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Report success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getReportRequestResult = dao.GetReport( uint64(createReportObj.ID) )
	
	if getReportRequestResult.Success == true {
		t.Errorf(getReportRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestDashboardCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Dashboard
	//----------------------------------------------------------------------------
	DashboardObj := model.Dashboard                                                                            {Title:"test value for Title",Theme:"test value for Theme",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createDashboardRequestResult := dao.CreateDashboard( DashboardObj )
	
	if createDashboardRequestResult.Success == false {
		t.Errorf(createDashboardRequestResult.Msg)
	} else {
		fmt.Println("Check Create Dashboard success...")
	}
	
	createDashboardObj,_ := createDashboardRequestResult.Data. (model.Dashboard)

	// --------------------------------------------------------------
	// Check Dashboard Obj ID
	// --------------------------------------------------------------	
	if createDashboardObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Dashboard" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getDashboardRequestResult := dao.GetDashboard( uint64(createDashboardObj.ID) )
	
	if getDashboardRequestResult.Success == false {
		t.Errorf(getDashboardRequestResult.Msg)
	} else {
		fmt.Println("Check Get Dashboard success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getDashboardObj,_ := getDashboardRequestResult.Data. (model.Dashboard)
	compareDashboard := cmp.Equal(createDashboardObj.ID, getDashboardObj.ID)
	
	if  compareDashboard == false	{
		t.Errorf( "Created Dashboard object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllDashboardRequestResult := dao.GetAllDashboard()

	if getAllDashboardRequestResult.Success == false {
			t.Errorf(getAllDashboardRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Dashboard success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllDashboardObj []model.Dashboard = getAllDashboardRequestResult.Data. ([]model.Dashboard)
		
	equalDashboard := cmp.Equal(createDashboardObj.ID, getAllDashboardObj[len(getAllDashboardObj)-1].ID)
		
	if equalDashboard == false {
		t.Errorf( "Created object is not equal to the last entry in Dashboard[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Dashboard
	// --------------------------------------------------------------	
	deleteDashboardRequestResult := dao.DeleteDashboard(uint64(createDashboardObj.ID))

	if deleteDashboardRequestResult.Success == false {
			t.Errorf(deleteDashboardRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Dashboard success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getDashboardRequestResult = dao.GetDashboard( uint64(createDashboardObj.ID) )
	
	if getDashboardRequestResult.Success == true {
		t.Errorf(getDashboardRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestVisualizationCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Visualization
	//----------------------------------------------------------------------------
	VisualizationObj := model.Visualization                                                            {Title:"test value for Title",Options:new ChartOptions(),ChartType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createVisualizationRequestResult := dao.CreateVisualization( VisualizationObj )
	
	if createVisualizationRequestResult.Success == false {
		t.Errorf(createVisualizationRequestResult.Msg)
	} else {
		fmt.Println("Check Create Visualization success...")
	}
	
	createVisualizationObj,_ := createVisualizationRequestResult.Data. (model.Visualization)

	// --------------------------------------------------------------
	// Check Visualization Obj ID
	// --------------------------------------------------------------	
	if createVisualizationObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Visualization" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getVisualizationRequestResult := dao.GetVisualization( uint64(createVisualizationObj.ID) )
	
	if getVisualizationRequestResult.Success == false {
		t.Errorf(getVisualizationRequestResult.Msg)
	} else {
		fmt.Println("Check Get Visualization success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getVisualizationObj,_ := getVisualizationRequestResult.Data. (model.Visualization)
	compareVisualization := cmp.Equal(createVisualizationObj.ID, getVisualizationObj.ID)
	
	if  compareVisualization == false	{
		t.Errorf( "Created Visualization object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllVisualizationRequestResult := dao.GetAllVisualization()

	if getAllVisualizationRequestResult.Success == false {
			t.Errorf(getAllVisualizationRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Visualization success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllVisualizationObj []model.Visualization = getAllVisualizationRequestResult.Data. ([]model.Visualization)
		
	equalVisualization := cmp.Equal(createVisualizationObj.ID, getAllVisualizationObj[len(getAllVisualizationObj)-1].ID)
		
	if equalVisualization == false {
		t.Errorf( "Created object is not equal to the last entry in Visualization[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Visualization
	// --------------------------------------------------------------	
	deleteVisualizationRequestResult := dao.DeleteVisualization(uint64(createVisualizationObj.ID))

	if deleteVisualizationRequestResult.Success == false {
			t.Errorf(deleteVisualizationRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Visualization success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getVisualizationRequestResult = dao.GetVisualization( uint64(createVisualizationObj.ID) )
	
	if getVisualizationRequestResult.Success == true {
		t.Errorf(getVisualizationRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestNotebookCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Notebook
	//----------------------------------------------------------------------------
	NotebookObj := model.Notebook                                                            {Title:"test value for Title",Repository:new RepositoryRef(),Language:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createNotebookRequestResult := dao.CreateNotebook( NotebookObj )
	
	if createNotebookRequestResult.Success == false {
		t.Errorf(createNotebookRequestResult.Msg)
	} else {
		fmt.Println("Check Create Notebook success...")
	}
	
	createNotebookObj,_ := createNotebookRequestResult.Data. (model.Notebook)

	// --------------------------------------------------------------
	// Check Notebook Obj ID
	// --------------------------------------------------------------	
	if createNotebookObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Notebook" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getNotebookRequestResult := dao.GetNotebook( uint64(createNotebookObj.ID) )
	
	if getNotebookRequestResult.Success == false {
		t.Errorf(getNotebookRequestResult.Msg)
	} else {
		fmt.Println("Check Get Notebook success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getNotebookObj,_ := getNotebookRequestResult.Data. (model.Notebook)
	compareNotebook := cmp.Equal(createNotebookObj.ID, getNotebookObj.ID)
	
	if  compareNotebook == false	{
		t.Errorf( "Created Notebook object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllNotebookRequestResult := dao.GetAllNotebook()

	if getAllNotebookRequestResult.Success == false {
			t.Errorf(getAllNotebookRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Notebook success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllNotebookObj []model.Notebook = getAllNotebookRequestResult.Data. ([]model.Notebook)
		
	equalNotebook := cmp.Equal(createNotebookObj.ID, getAllNotebookObj[len(getAllNotebookObj)-1].ID)
		
	if equalNotebook == false {
		t.Errorf( "Created object is not equal to the last entry in Notebook[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Notebook
	// --------------------------------------------------------------	
	deleteNotebookRequestResult := dao.DeleteNotebook(uint64(createNotebookObj.ID))

	if deleteNotebookRequestResult.Success == false {
			t.Errorf(deleteNotebookRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Notebook success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getNotebookRequestResult = dao.GetNotebook( uint64(createNotebookObj.ID) )
	
	if getNotebookRequestResult.Success == true {
		t.Errorf(getNotebookRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBIQueryCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BIQuery
	//----------------------------------------------------------------------------
	BIQueryObj := model.BIQuery                                                                            {Name:"test value for Name",Text:"test value for Text",Dialect:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBIQueryRequestResult := dao.CreateBIQuery( BIQueryObj )
	
	if createBIQueryRequestResult.Success == false {
		t.Errorf(createBIQueryRequestResult.Msg)
	} else {
		fmt.Println("Check Create BIQuery success...")
	}
	
	createBIQueryObj,_ := createBIQueryRequestResult.Data. (model.BIQuery)

	// --------------------------------------------------------------
	// Check BIQuery Obj ID
	// --------------------------------------------------------------	
	if createBIQueryObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BIQuery" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBIQueryRequestResult := dao.GetBIQuery( uint64(createBIQueryObj.ID) )
	
	if getBIQueryRequestResult.Success == false {
		t.Errorf(getBIQueryRequestResult.Msg)
	} else {
		fmt.Println("Check Get BIQuery success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBIQueryObj,_ := getBIQueryRequestResult.Data. (model.BIQuery)
	compareBIQuery := cmp.Equal(createBIQueryObj.ID, getBIQueryObj.ID)
	
	if  compareBIQuery == false	{
		t.Errorf( "Created BIQuery object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBIQueryRequestResult := dao.GetAllBIQuery()

	if getAllBIQueryRequestResult.Success == false {
			t.Errorf(getAllBIQueryRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BIQuery success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBIQueryObj []model.BIQuery = getAllBIQueryRequestResult.Data. ([]model.BIQuery)
		
	equalBIQuery := cmp.Equal(createBIQueryObj.ID, getAllBIQueryObj[len(getAllBIQueryObj)-1].ID)
		
	if equalBIQuery == false {
		t.Errorf( "Created object is not equal to the last entry in BIQuery[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BIQuery
	// --------------------------------------------------------------	
	deleteBIQueryRequestResult := dao.DeleteBIQuery(uint64(createBIQueryObj.ID))

	if deleteBIQueryRequestResult.Success == false {
			t.Errorf(deleteBIQueryRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BIQuery success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBIQueryRequestResult = dao.GetBIQuery( uint64(createBIQueryObj.ID) )
	
	if getBIQueryRequestResult.Success == true {
		t.Errorf(getBIQueryRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestExperimentCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Experiment
	//----------------------------------------------------------------------------
	ExperimentObj := model.Experiment                                                                            {Name:"test value for Name",Objective:"test value for Objective",Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createExperimentRequestResult := dao.CreateExperiment( ExperimentObj )
	
	if createExperimentRequestResult.Success == false {
		t.Errorf(createExperimentRequestResult.Msg)
	} else {
		fmt.Println("Check Create Experiment success...")
	}
	
	createExperimentObj,_ := createExperimentRequestResult.Data. (model.Experiment)

	// --------------------------------------------------------------
	// Check Experiment Obj ID
	// --------------------------------------------------------------	
	if createExperimentObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Experiment" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getExperimentRequestResult := dao.GetExperiment( uint64(createExperimentObj.ID) )
	
	if getExperimentRequestResult.Success == false {
		t.Errorf(getExperimentRequestResult.Msg)
	} else {
		fmt.Println("Check Get Experiment success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getExperimentObj,_ := getExperimentRequestResult.Data. (model.Experiment)
	compareExperiment := cmp.Equal(createExperimentObj.ID, getExperimentObj.ID)
	
	if  compareExperiment == false	{
		t.Errorf( "Created Experiment object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllExperimentRequestResult := dao.GetAllExperiment()

	if getAllExperimentRequestResult.Success == false {
			t.Errorf(getAllExperimentRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Experiment success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllExperimentObj []model.Experiment = getAllExperimentRequestResult.Data. ([]model.Experiment)
		
	equalExperiment := cmp.Equal(createExperimentObj.ID, getAllExperimentObj[len(getAllExperimentObj)-1].ID)
		
	if equalExperiment == false {
		t.Errorf( "Created object is not equal to the last entry in Experiment[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Experiment
	// --------------------------------------------------------------	
	deleteExperimentRequestResult := dao.DeleteExperiment(uint64(createExperimentObj.ID))

	if deleteExperimentRequestResult.Success == false {
			t.Errorf(deleteExperimentRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Experiment success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getExperimentRequestResult = dao.GetExperiment( uint64(createExperimentObj.ID) )
	
	if getExperimentRequestResult.Success == true {
		t.Errorf(getExperimentRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTrainingRunCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TrainingRun
	//----------------------------------------------------------------------------
	TrainingRunObj := model.TrainingRun                                                                                                                                                            {RunLabel:"test value for RunLabel",StartedAt:time.Now(),CompletedAt:time.Now(),Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTrainingRunRequestResult := dao.CreateTrainingRun( TrainingRunObj )
	
	if createTrainingRunRequestResult.Success == false {
		t.Errorf(createTrainingRunRequestResult.Msg)
	} else {
		fmt.Println("Check Create TrainingRun success...")
	}
	
	createTrainingRunObj,_ := createTrainingRunRequestResult.Data. (model.TrainingRun)

	// --------------------------------------------------------------
	// Check TrainingRun Obj ID
	// --------------------------------------------------------------	
	if createTrainingRunObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TrainingRun" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTrainingRunRequestResult := dao.GetTrainingRun( uint64(createTrainingRunObj.ID) )
	
	if getTrainingRunRequestResult.Success == false {
		t.Errorf(getTrainingRunRequestResult.Msg)
	} else {
		fmt.Println("Check Get TrainingRun success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTrainingRunObj,_ := getTrainingRunRequestResult.Data. (model.TrainingRun)
	compareTrainingRun := cmp.Equal(createTrainingRunObj.ID, getTrainingRunObj.ID)
	
	if  compareTrainingRun == false	{
		t.Errorf( "Created TrainingRun object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTrainingRunRequestResult := dao.GetAllTrainingRun()

	if getAllTrainingRunRequestResult.Success == false {
			t.Errorf(getAllTrainingRunRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TrainingRun success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTrainingRunObj []model.TrainingRun = getAllTrainingRunRequestResult.Data. ([]model.TrainingRun)
		
	equalTrainingRun := cmp.Equal(createTrainingRunObj.ID, getAllTrainingRunObj[len(getAllTrainingRunObj)-1].ID)
		
	if equalTrainingRun == false {
		t.Errorf( "Created object is not equal to the last entry in TrainingRun[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TrainingRun
	// --------------------------------------------------------------	
	deleteTrainingRunRequestResult := dao.DeleteTrainingRun(uint64(createTrainingRunObj.ID))

	if deleteTrainingRunRequestResult.Success == false {
			t.Errorf(deleteTrainingRunRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TrainingRun success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTrainingRunRequestResult = dao.GetTrainingRun( uint64(createTrainingRunObj.ID) )
	
	if getTrainingRunRequestResult.Success == true {
		t.Errorf(getTrainingRunRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRunMetricCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RunMetric
	//----------------------------------------------------------------------------
	RunMetricObj := model.RunMetric                                                                                    {Name:"test value for Name",Value:"test value"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRunMetricRequestResult := dao.CreateRunMetric( RunMetricObj )
	
	if createRunMetricRequestResult.Success == false {
		t.Errorf(createRunMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Create RunMetric success...")
	}
	
	createRunMetricObj,_ := createRunMetricRequestResult.Data. (model.RunMetric)

	// --------------------------------------------------------------
	// Check RunMetric Obj ID
	// --------------------------------------------------------------	
	if createRunMetricObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RunMetric" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRunMetricRequestResult := dao.GetRunMetric( uint64(createRunMetricObj.ID) )
	
	if getRunMetricRequestResult.Success == false {
		t.Errorf(getRunMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Get RunMetric success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRunMetricObj,_ := getRunMetricRequestResult.Data. (model.RunMetric)
	compareRunMetric := cmp.Equal(createRunMetricObj.ID, getRunMetricObj.ID)
	
	if  compareRunMetric == false	{
		t.Errorf( "Created RunMetric object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRunMetricRequestResult := dao.GetAllRunMetric()

	if getAllRunMetricRequestResult.Success == false {
			t.Errorf(getAllRunMetricRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RunMetric success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRunMetricObj []model.RunMetric = getAllRunMetricRequestResult.Data. ([]model.RunMetric)
		
	equalRunMetric := cmp.Equal(createRunMetricObj.ID, getAllRunMetricObj[len(getAllRunMetricObj)-1].ID)
		
	if equalRunMetric == false {
		t.Errorf( "Created object is not equal to the last entry in RunMetric[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RunMetric
	// --------------------------------------------------------------	
	deleteRunMetricRequestResult := dao.DeleteRunMetric(uint64(createRunMetricObj.ID))

	if deleteRunMetricRequestResult.Success == false {
			t.Errorf(deleteRunMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RunMetric success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRunMetricRequestResult = dao.GetRunMetric( uint64(createRunMetricObj.ID) )
	
	if getRunMetricRequestResult.Success == true {
		t.Errorf(getRunMetricRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRunParameterCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RunParameter
	//----------------------------------------------------------------------------
	RunParameterObj := model.RunParameter                                                            {Name:"test value for Name",Value:"test value for Value"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRunParameterRequestResult := dao.CreateRunParameter( RunParameterObj )
	
	if createRunParameterRequestResult.Success == false {
		t.Errorf(createRunParameterRequestResult.Msg)
	} else {
		fmt.Println("Check Create RunParameter success...")
	}
	
	createRunParameterObj,_ := createRunParameterRequestResult.Data. (model.RunParameter)

	// --------------------------------------------------------------
	// Check RunParameter Obj ID
	// --------------------------------------------------------------	
	if createRunParameterObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RunParameter" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRunParameterRequestResult := dao.GetRunParameter( uint64(createRunParameterObj.ID) )
	
	if getRunParameterRequestResult.Success == false {
		t.Errorf(getRunParameterRequestResult.Msg)
	} else {
		fmt.Println("Check Get RunParameter success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRunParameterObj,_ := getRunParameterRequestResult.Data. (model.RunParameter)
	compareRunParameter := cmp.Equal(createRunParameterObj.ID, getRunParameterObj.ID)
	
	if  compareRunParameter == false	{
		t.Errorf( "Created RunParameter object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRunParameterRequestResult := dao.GetAllRunParameter()

	if getAllRunParameterRequestResult.Success == false {
			t.Errorf(getAllRunParameterRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RunParameter success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRunParameterObj []model.RunParameter = getAllRunParameterRequestResult.Data. ([]model.RunParameter)
		
	equalRunParameter := cmp.Equal(createRunParameterObj.ID, getAllRunParameterObj[len(getAllRunParameterObj)-1].ID)
		
	if equalRunParameter == false {
		t.Errorf( "Created object is not equal to the last entry in RunParameter[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RunParameter
	// --------------------------------------------------------------	
	deleteRunParameterRequestResult := dao.DeleteRunParameter(uint64(createRunParameterObj.ID))

	if deleteRunParameterRequestResult.Success == false {
			t.Errorf(deleteRunParameterRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RunParameter success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRunParameterRequestResult = dao.GetRunParameter( uint64(createRunParameterObj.ID) )
	
	if getRunParameterRequestResult.Success == true {
		t.Errorf(getRunParameterRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestModel_CRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Model_
	//----------------------------------------------------------------------------
	Model_Obj := model.Model_                                                                            {Name:"test value for Name",TaskDescription:"test value for TaskDescription",ModelType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createModel_RequestResult := dao.CreateModel_( Model_Obj )
	
	if createModel_RequestResult.Success == false {
		t.Errorf(createModel_RequestResult.Msg)
	} else {
		fmt.Println("Check Create Model_ success...")
	}
	
	createModel_Obj,_ := createModel_RequestResult.Data. (model.Model_)

	// --------------------------------------------------------------
	// Check Model_ Obj ID
	// --------------------------------------------------------------	
	if createModel_Obj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Model_" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getModel_RequestResult := dao.GetModel_( uint64(createModel_Obj.ID) )
	
	if getModel_RequestResult.Success == false {
		t.Errorf(getModel_RequestResult.Msg)
	} else {
		fmt.Println("Check Get Model_ success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getModel_Obj,_ := getModel_RequestResult.Data. (model.Model_)
	compareModel_ := cmp.Equal(createModel_Obj.ID, getModel_Obj.ID)
	
	if  compareModel_ == false	{
		t.Errorf( "Created Model_ object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllModel_RequestResult := dao.GetAllModel_()

	if getAllModel_RequestResult.Success == false {
			t.Errorf(getAllModel_RequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Model_ success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllModel_Obj []model.Model_ = getAllModel_RequestResult.Data. ([]model.Model_)
		
	equalModel_ := cmp.Equal(createModel_Obj.ID, getAllModel_Obj[len(getAllModel_Obj)-1].ID)
		
	if equalModel_ == false {
		t.Errorf( "Created object is not equal to the last entry in Model_[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Model_
	// --------------------------------------------------------------	
	deleteModel_RequestResult := dao.DeleteModel_(uint64(createModel_Obj.ID))

	if deleteModel_RequestResult.Success == false {
			t.Errorf(deleteModel_RequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Model_ success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getModel_RequestResult = dao.GetModel_( uint64(createModel_Obj.ID) )
	
	if getModel_RequestResult.Success == true {
		t.Errorf(getModel_RequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestModelVersionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for ModelVersion
	//----------------------------------------------------------------------------
	ModelVersionObj := model.ModelVersion                                                            {Version:"test value for Version",Lifecycle:0,TrainingStatus:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createModelVersionRequestResult := dao.CreateModelVersion( ModelVersionObj )
	
	if createModelVersionRequestResult.Success == false {
		t.Errorf(createModelVersionRequestResult.Msg)
	} else {
		fmt.Println("Check Create ModelVersion success...")
	}
	
	createModelVersionObj,_ := createModelVersionRequestResult.Data. (model.ModelVersion)

	// --------------------------------------------------------------
	// Check ModelVersion Obj ID
	// --------------------------------------------------------------	
	if createModelVersionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for ModelVersion" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getModelVersionRequestResult := dao.GetModelVersion( uint64(createModelVersionObj.ID) )
	
	if getModelVersionRequestResult.Success == false {
		t.Errorf(getModelVersionRequestResult.Msg)
	} else {
		fmt.Println("Check Get ModelVersion success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getModelVersionObj,_ := getModelVersionRequestResult.Data. (model.ModelVersion)
	compareModelVersion := cmp.Equal(createModelVersionObj.ID, getModelVersionObj.ID)
	
	if  compareModelVersion == false	{
		t.Errorf( "Created ModelVersion object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllModelVersionRequestResult := dao.GetAllModelVersion()

	if getAllModelVersionRequestResult.Success == false {
			t.Errorf(getAllModelVersionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll ModelVersion success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllModelVersionObj []model.ModelVersion = getAllModelVersionRequestResult.Data. ([]model.ModelVersion)
		
	equalModelVersion := cmp.Equal(createModelVersionObj.ID, getAllModelVersionObj[len(getAllModelVersionObj)-1].ID)
		
	if equalModelVersion == false {
		t.Errorf( "Created object is not equal to the last entry in ModelVersion[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for ModelVersion
	// --------------------------------------------------------------	
	deleteModelVersionRequestResult := dao.DeleteModelVersion(uint64(createModelVersionObj.ID))

	if deleteModelVersionRequestResult.Success == false {
			t.Errorf(deleteModelVersionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion ModelVersion success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getModelVersionRequestResult = dao.GetModelVersion( uint64(createModelVersionObj.ID) )
	
	if getModelVersionRequestResult.Success == true {
		t.Errorf(getModelVersionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestEvaluationMetricCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for EvaluationMetric
	//----------------------------------------------------------------------------
	EvaluationMetricObj := model.EvaluationMetric                                                                                    {Name:"test value for Name",Value:"test value"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createEvaluationMetricRequestResult := dao.CreateEvaluationMetric( EvaluationMetricObj )
	
	if createEvaluationMetricRequestResult.Success == false {
		t.Errorf(createEvaluationMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Create EvaluationMetric success...")
	}
	
	createEvaluationMetricObj,_ := createEvaluationMetricRequestResult.Data. (model.EvaluationMetric)

	// --------------------------------------------------------------
	// Check EvaluationMetric Obj ID
	// --------------------------------------------------------------	
	if createEvaluationMetricObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for EvaluationMetric" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getEvaluationMetricRequestResult := dao.GetEvaluationMetric( uint64(createEvaluationMetricObj.ID) )
	
	if getEvaluationMetricRequestResult.Success == false {
		t.Errorf(getEvaluationMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Get EvaluationMetric success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getEvaluationMetricObj,_ := getEvaluationMetricRequestResult.Data. (model.EvaluationMetric)
	compareEvaluationMetric := cmp.Equal(createEvaluationMetricObj.ID, getEvaluationMetricObj.ID)
	
	if  compareEvaluationMetric == false	{
		t.Errorf( "Created EvaluationMetric object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllEvaluationMetricRequestResult := dao.GetAllEvaluationMetric()

	if getAllEvaluationMetricRequestResult.Success == false {
			t.Errorf(getAllEvaluationMetricRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll EvaluationMetric success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllEvaluationMetricObj []model.EvaluationMetric = getAllEvaluationMetricRequestResult.Data. ([]model.EvaluationMetric)
		
	equalEvaluationMetric := cmp.Equal(createEvaluationMetricObj.ID, getAllEvaluationMetricObj[len(getAllEvaluationMetricObj)-1].ID)
		
	if equalEvaluationMetric == false {
		t.Errorf( "Created object is not equal to the last entry in EvaluationMetric[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for EvaluationMetric
	// --------------------------------------------------------------	
	deleteEvaluationMetricRequestResult := dao.DeleteEvaluationMetric(uint64(createEvaluationMetricObj.ID))

	if deleteEvaluationMetricRequestResult.Success == false {
			t.Errorf(deleteEvaluationMetricRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion EvaluationMetric success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getEvaluationMetricRequestResult = dao.GetEvaluationMetric( uint64(createEvaluationMetricObj.ID) )
	
	if getEvaluationMetricRequestResult.Success == true {
		t.Errorf(getEvaluationMetricRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFeatureSetCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FeatureSet
	//----------------------------------------------------------------------------
	FeatureSetObj := model.FeatureSet                                                            {Name:"test value for Name",RefreshSchedule:new CronSchedule(),StoreType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFeatureSetRequestResult := dao.CreateFeatureSet( FeatureSetObj )
	
	if createFeatureSetRequestResult.Success == false {
		t.Errorf(createFeatureSetRequestResult.Msg)
	} else {
		fmt.Println("Check Create FeatureSet success...")
	}
	
	createFeatureSetObj,_ := createFeatureSetRequestResult.Data. (model.FeatureSet)

	// --------------------------------------------------------------
	// Check FeatureSet Obj ID
	// --------------------------------------------------------------	
	if createFeatureSetObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for FeatureSet" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFeatureSetRequestResult := dao.GetFeatureSet( uint64(createFeatureSetObj.ID) )
	
	if getFeatureSetRequestResult.Success == false {
		t.Errorf(getFeatureSetRequestResult.Msg)
	} else {
		fmt.Println("Check Get FeatureSet success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFeatureSetObj,_ := getFeatureSetRequestResult.Data. (model.FeatureSet)
	compareFeatureSet := cmp.Equal(createFeatureSetObj.ID, getFeatureSetObj.ID)
	
	if  compareFeatureSet == false	{
		t.Errorf( "Created FeatureSet object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFeatureSetRequestResult := dao.GetAllFeatureSet()

	if getAllFeatureSetRequestResult.Success == false {
			t.Errorf(getAllFeatureSetRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FeatureSet success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFeatureSetObj []model.FeatureSet = getAllFeatureSetRequestResult.Data. ([]model.FeatureSet)
		
	equalFeatureSet := cmp.Equal(createFeatureSetObj.ID, getAllFeatureSetObj[len(getAllFeatureSetObj)-1].ID)
		
	if equalFeatureSet == false {
		t.Errorf( "Created object is not equal to the last entry in FeatureSet[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FeatureSet
	// --------------------------------------------------------------	
	deleteFeatureSetRequestResult := dao.DeleteFeatureSet(uint64(createFeatureSetObj.ID))

	if deleteFeatureSetRequestResult.Success == false {
			t.Errorf(deleteFeatureSetRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FeatureSet success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFeatureSetRequestResult = dao.GetFeatureSet( uint64(createFeatureSetObj.ID) )
	
	if getFeatureSetRequestResult.Success == true {
		t.Errorf(getFeatureSetRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFeatureCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Feature
	//----------------------------------------------------------------------------
	FeatureObj := model.Feature                                                                            {Name:"test value for Name",Description:"test value for Description",DataType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFeatureRequestResult := dao.CreateFeature( FeatureObj )
	
	if createFeatureRequestResult.Success == false {
		t.Errorf(createFeatureRequestResult.Msg)
	} else {
		fmt.Println("Check Create Feature success...")
	}
	
	createFeatureObj,_ := createFeatureRequestResult.Data. (model.Feature)

	// --------------------------------------------------------------
	// Check Feature Obj ID
	// --------------------------------------------------------------	
	if createFeatureObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Feature" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFeatureRequestResult := dao.GetFeature( uint64(createFeatureObj.ID) )
	
	if getFeatureRequestResult.Success == false {
		t.Errorf(getFeatureRequestResult.Msg)
	} else {
		fmt.Println("Check Get Feature success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFeatureObj,_ := getFeatureRequestResult.Data. (model.Feature)
	compareFeature := cmp.Equal(createFeatureObj.ID, getFeatureObj.ID)
	
	if  compareFeature == false	{
		t.Errorf( "Created Feature object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFeatureRequestResult := dao.GetAllFeature()

	if getAllFeatureRequestResult.Success == false {
			t.Errorf(getAllFeatureRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Feature success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFeatureObj []model.Feature = getAllFeatureRequestResult.Data. ([]model.Feature)
		
	equalFeature := cmp.Equal(createFeatureObj.ID, getAllFeatureObj[len(getAllFeatureObj)-1].ID)
		
	if equalFeature == false {
		t.Errorf( "Created object is not equal to the last entry in Feature[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Feature
	// --------------------------------------------------------------	
	deleteFeatureRequestResult := dao.DeleteFeature(uint64(createFeatureObj.ID))

	if deleteFeatureRequestResult.Success == false {
			t.Errorf(deleteFeatureRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Feature success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFeatureRequestResult = dao.GetFeature( uint64(createFeatureObj.ID) )
	
	if getFeatureRequestResult.Success == true {
		t.Errorf(getFeatureRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestInferenceEndpointCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for InferenceEndpoint
	//----------------------------------------------------------------------------
	InferenceEndpointObj := model.InferenceEndpoint                                                                                            {Name:"test value for Name",EndpointUrl:"test value for EndpointUrl",TrafficShare:new Percentage(),Mode:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createInferenceEndpointRequestResult := dao.CreateInferenceEndpoint( InferenceEndpointObj )
	
	if createInferenceEndpointRequestResult.Success == false {
		t.Errorf(createInferenceEndpointRequestResult.Msg)
	} else {
		fmt.Println("Check Create InferenceEndpoint success...")
	}
	
	createInferenceEndpointObj,_ := createInferenceEndpointRequestResult.Data. (model.InferenceEndpoint)

	// --------------------------------------------------------------
	// Check InferenceEndpoint Obj ID
	// --------------------------------------------------------------	
	if createInferenceEndpointObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for InferenceEndpoint" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getInferenceEndpointRequestResult := dao.GetInferenceEndpoint( uint64(createInferenceEndpointObj.ID) )
	
	if getInferenceEndpointRequestResult.Success == false {
		t.Errorf(getInferenceEndpointRequestResult.Msg)
	} else {
		fmt.Println("Check Get InferenceEndpoint success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getInferenceEndpointObj,_ := getInferenceEndpointRequestResult.Data. (model.InferenceEndpoint)
	compareInferenceEndpoint := cmp.Equal(createInferenceEndpointObj.ID, getInferenceEndpointObj.ID)
	
	if  compareInferenceEndpoint == false	{
		t.Errorf( "Created InferenceEndpoint object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllInferenceEndpointRequestResult := dao.GetAllInferenceEndpoint()

	if getAllInferenceEndpointRequestResult.Success == false {
			t.Errorf(getAllInferenceEndpointRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll InferenceEndpoint success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllInferenceEndpointObj []model.InferenceEndpoint = getAllInferenceEndpointRequestResult.Data. ([]model.InferenceEndpoint)
		
	equalInferenceEndpoint := cmp.Equal(createInferenceEndpointObj.ID, getAllInferenceEndpointObj[len(getAllInferenceEndpointObj)-1].ID)
		
	if equalInferenceEndpoint == false {
		t.Errorf( "Created object is not equal to the last entry in InferenceEndpoint[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for InferenceEndpoint
	// --------------------------------------------------------------	
	deleteInferenceEndpointRequestResult := dao.DeleteInferenceEndpoint(uint64(createInferenceEndpointObj.ID))

	if deleteInferenceEndpointRequestResult.Success == false {
			t.Errorf(deleteInferenceEndpointRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion InferenceEndpoint success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getInferenceEndpointRequestResult = dao.GetInferenceEndpoint( uint64(createInferenceEndpointObj.ID) )
	
	if getInferenceEndpointRequestResult.Success == true {
		t.Errorf(getInferenceEndpointRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestPredictionCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Prediction
	//----------------------------------------------------------------------------
	PredictionObj := model.Prediction                                                                                                                                            {ReferenceKey:"test value for ReferenceKey",PredictedAt:time.Now(),Score:"test value"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createPredictionRequestResult := dao.CreatePrediction( PredictionObj )
	
	if createPredictionRequestResult.Success == false {
		t.Errorf(createPredictionRequestResult.Msg)
	} else {
		fmt.Println("Check Create Prediction success...")
	}
	
	createPredictionObj,_ := createPredictionRequestResult.Data. (model.Prediction)

	// --------------------------------------------------------------
	// Check Prediction Obj ID
	// --------------------------------------------------------------	
	if createPredictionObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Prediction" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getPredictionRequestResult := dao.GetPrediction( uint64(createPredictionObj.ID) )
	
	if getPredictionRequestResult.Success == false {
		t.Errorf(getPredictionRequestResult.Msg)
	} else {
		fmt.Println("Check Get Prediction success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getPredictionObj,_ := getPredictionRequestResult.Data. (model.Prediction)
	comparePrediction := cmp.Equal(createPredictionObj.ID, getPredictionObj.ID)
	
	if  comparePrediction == false	{
		t.Errorf( "Created Prediction object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllPredictionRequestResult := dao.GetAllPrediction()

	if getAllPredictionRequestResult.Success == false {
			t.Errorf(getAllPredictionRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Prediction success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllPredictionObj []model.Prediction = getAllPredictionRequestResult.Data. ([]model.Prediction)
		
	equalPrediction := cmp.Equal(createPredictionObj.ID, getAllPredictionObj[len(getAllPredictionObj)-1].ID)
		
	if equalPrediction == false {
		t.Errorf( "Created object is not equal to the last entry in Prediction[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Prediction
	// --------------------------------------------------------------	
	deletePredictionRequestResult := dao.DeletePrediction(uint64(createPredictionObj.ID))

	if deletePredictionRequestResult.Success == false {
			t.Errorf(deletePredictionRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Prediction success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getPredictionRequestResult = dao.GetPrediction( uint64(createPredictionObj.ID) )
	
	if getPredictionRequestResult.Success == true {
		t.Errorf(getPredictionRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestForecastCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Forecast
	//----------------------------------------------------------------------------
	ForecastObj := model.Forecast                                                                            {Name:"test value for Name",Horizon:100,Granularity:0}

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


func TestTimeSeriesCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for TimeSeries
	//----------------------------------------------------------------------------
	TimeSeriesObj := model.TimeSeries                                                                            {Name:"test value for Name",Timezone:"test value for Timezone",Granularity:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTimeSeriesRequestResult := dao.CreateTimeSeries( TimeSeriesObj )
	
	if createTimeSeriesRequestResult.Success == false {
		t.Errorf(createTimeSeriesRequestResult.Msg)
	} else {
		fmt.Println("Check Create TimeSeries success...")
	}
	
	createTimeSeriesObj,_ := createTimeSeriesRequestResult.Data. (model.TimeSeries)

	// --------------------------------------------------------------
	// Check TimeSeries Obj ID
	// --------------------------------------------------------------	
	if createTimeSeriesObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for TimeSeries" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTimeSeriesRequestResult := dao.GetTimeSeries( uint64(createTimeSeriesObj.ID) )
	
	if getTimeSeriesRequestResult.Success == false {
		t.Errorf(getTimeSeriesRequestResult.Msg)
	} else {
		fmt.Println("Check Get TimeSeries success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTimeSeriesObj,_ := getTimeSeriesRequestResult.Data. (model.TimeSeries)
	compareTimeSeries := cmp.Equal(createTimeSeriesObj.ID, getTimeSeriesObj.ID)
	
	if  compareTimeSeries == false	{
		t.Errorf( "Created TimeSeries object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTimeSeriesRequestResult := dao.GetAllTimeSeries()

	if getAllTimeSeriesRequestResult.Success == false {
			t.Errorf(getAllTimeSeriesRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll TimeSeries success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTimeSeriesObj []model.TimeSeries = getAllTimeSeriesRequestResult.Data. ([]model.TimeSeries)
		
	equalTimeSeries := cmp.Equal(createTimeSeriesObj.ID, getAllTimeSeriesObj[len(getAllTimeSeriesObj)-1].ID)
		
	if equalTimeSeries == false {
		t.Errorf( "Created object is not equal to the last entry in TimeSeries[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for TimeSeries
	// --------------------------------------------------------------	
	deleteTimeSeriesRequestResult := dao.DeleteTimeSeries(uint64(createTimeSeriesObj.ID))

	if deleteTimeSeriesRequestResult.Success == false {
			t.Errorf(deleteTimeSeriesRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion TimeSeries success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTimeSeriesRequestResult = dao.GetTimeSeries( uint64(createTimeSeriesObj.ID) )
	
	if getTimeSeriesRequestResult.Success == true {
		t.Errorf(getTimeSeriesRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAnomalyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Anomaly
	//----------------------------------------------------------------------------
	AnomalyObj := model.Anomaly                                                                                                                    {OccurredAt:time.Now(),Details:"test value for Details",AnomalyType:0,Severity:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAnomalyRequestResult := dao.CreateAnomaly( AnomalyObj )
	
	if createAnomalyRequestResult.Success == false {
		t.Errorf(createAnomalyRequestResult.Msg)
	} else {
		fmt.Println("Check Create Anomaly success...")
	}
	
	createAnomalyObj,_ := createAnomalyRequestResult.Data. (model.Anomaly)

	// --------------------------------------------------------------
	// Check Anomaly Obj ID
	// --------------------------------------------------------------	
	if createAnomalyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Anomaly" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAnomalyRequestResult := dao.GetAnomaly( uint64(createAnomalyObj.ID) )
	
	if getAnomalyRequestResult.Success == false {
		t.Errorf(getAnomalyRequestResult.Msg)
	} else {
		fmt.Println("Check Get Anomaly success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAnomalyObj,_ := getAnomalyRequestResult.Data. (model.Anomaly)
	compareAnomaly := cmp.Equal(createAnomalyObj.ID, getAnomalyObj.ID)
	
	if  compareAnomaly == false	{
		t.Errorf( "Created Anomaly object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAnomalyRequestResult := dao.GetAllAnomaly()

	if getAllAnomalyRequestResult.Success == false {
			t.Errorf(getAllAnomalyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Anomaly success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAnomalyObj []model.Anomaly = getAllAnomalyRequestResult.Data. ([]model.Anomaly)
		
	equalAnomaly := cmp.Equal(createAnomalyObj.ID, getAllAnomalyObj[len(getAllAnomalyObj)-1].ID)
		
	if equalAnomaly == false {
		t.Errorf( "Created object is not equal to the last entry in Anomaly[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Anomaly
	// --------------------------------------------------------------	
	deleteAnomalyRequestResult := dao.DeleteAnomaly(uint64(createAnomalyObj.ID))

	if deleteAnomalyRequestResult.Success == false {
			t.Errorf(deleteAnomalyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Anomaly success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAnomalyRequestResult = dao.GetAnomaly( uint64(createAnomalyObj.ID) )
	
	if getAnomalyRequestResult.Success == true {
		t.Errorf(getAnomalyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestQualityRuleCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for QualityRule
	//----------------------------------------------------------------------------
	QualityRuleObj := model.QualityRule                                                                                                            {Name:"test value for Name",Threshold:new Threshold(),TargetField:"test value for TargetField",Dimension:0,Operator:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createQualityRuleRequestResult := dao.CreateQualityRule( QualityRuleObj )
	
	if createQualityRuleRequestResult.Success == false {
		t.Errorf(createQualityRuleRequestResult.Msg)
	} else {
		fmt.Println("Check Create QualityRule success...")
	}
	
	createQualityRuleObj,_ := createQualityRuleRequestResult.Data. (model.QualityRule)

	// --------------------------------------------------------------
	// Check QualityRule Obj ID
	// --------------------------------------------------------------	
	if createQualityRuleObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for QualityRule" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getQualityRuleRequestResult := dao.GetQualityRule( uint64(createQualityRuleObj.ID) )
	
	if getQualityRuleRequestResult.Success == false {
		t.Errorf(getQualityRuleRequestResult.Msg)
	} else {
		fmt.Println("Check Get QualityRule success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getQualityRuleObj,_ := getQualityRuleRequestResult.Data. (model.QualityRule)
	compareQualityRule := cmp.Equal(createQualityRuleObj.ID, getQualityRuleObj.ID)
	
	if  compareQualityRule == false	{
		t.Errorf( "Created QualityRule object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllQualityRuleRequestResult := dao.GetAllQualityRule()

	if getAllQualityRuleRequestResult.Success == false {
			t.Errorf(getAllQualityRuleRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll QualityRule success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllQualityRuleObj []model.QualityRule = getAllQualityRuleRequestResult.Data. ([]model.QualityRule)
		
	equalQualityRule := cmp.Equal(createQualityRuleObj.ID, getAllQualityRuleObj[len(getAllQualityRuleObj)-1].ID)
		
	if equalQualityRule == false {
		t.Errorf( "Created object is not equal to the last entry in QualityRule[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for QualityRule
	// --------------------------------------------------------------	
	deleteQualityRuleRequestResult := dao.DeleteQualityRule(uint64(createQualityRuleObj.ID))

	if deleteQualityRuleRequestResult.Success == false {
			t.Errorf(deleteQualityRuleRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion QualityRule success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getQualityRuleRequestResult = dao.GetQualityRule( uint64(createQualityRuleObj.ID) )
	
	if getQualityRuleRequestResult.Success == true {
		t.Errorf(getQualityRuleRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestQualityCheckCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for QualityCheck
	//----------------------------------------------------------------------------
	QualityCheckObj := model.QualityCheck                                                                                                                                                            {CheckedAt:time.Now(),ObservedValue:"test value",SampleSize:100,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createQualityCheckRequestResult := dao.CreateQualityCheck( QualityCheckObj )
	
	if createQualityCheckRequestResult.Success == false {
		t.Errorf(createQualityCheckRequestResult.Msg)
	} else {
		fmt.Println("Check Create QualityCheck success...")
	}
	
	createQualityCheckObj,_ := createQualityCheckRequestResult.Data. (model.QualityCheck)

	// --------------------------------------------------------------
	// Check QualityCheck Obj ID
	// --------------------------------------------------------------	
	if createQualityCheckObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for QualityCheck" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getQualityCheckRequestResult := dao.GetQualityCheck( uint64(createQualityCheckObj.ID) )
	
	if getQualityCheckRequestResult.Success == false {
		t.Errorf(getQualityCheckRequestResult.Msg)
	} else {
		fmt.Println("Check Get QualityCheck success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getQualityCheckObj,_ := getQualityCheckRequestResult.Data. (model.QualityCheck)
	compareQualityCheck := cmp.Equal(createQualityCheckObj.ID, getQualityCheckObj.ID)
	
	if  compareQualityCheck == false	{
		t.Errorf( "Created QualityCheck object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllQualityCheckRequestResult := dao.GetAllQualityCheck()

	if getAllQualityCheckRequestResult.Success == false {
			t.Errorf(getAllQualityCheckRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll QualityCheck success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllQualityCheckObj []model.QualityCheck = getAllQualityCheckRequestResult.Data. ([]model.QualityCheck)
		
	equalQualityCheck := cmp.Equal(createQualityCheckObj.ID, getAllQualityCheckObj[len(getAllQualityCheckObj)-1].ID)
		
	if equalQualityCheck == false {
		t.Errorf( "Created object is not equal to the last entry in QualityCheck[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for QualityCheck
	// --------------------------------------------------------------	
	deleteQualityCheckRequestResult := dao.DeleteQualityCheck(uint64(createQualityCheckObj.ID))

	if deleteQualityCheckRequestResult.Success == false {
			t.Errorf(deleteQualityCheckRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion QualityCheck success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getQualityCheckRequestResult = dao.GetQualityCheck( uint64(createQualityCheckObj.ID) )
	
	if getQualityCheckRequestResult.Success == true {
		t.Errorf(getQualityCheckRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestLineageNodeCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for LineageNode
	//----------------------------------------------------------------------------
	LineageNodeObj := model.LineageNode                                                                            {Name:"test value for Name",QualifiedName:"test value for QualifiedName",NodeType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createLineageNodeRequestResult := dao.CreateLineageNode( LineageNodeObj )
	
	if createLineageNodeRequestResult.Success == false {
		t.Errorf(createLineageNodeRequestResult.Msg)
	} else {
		fmt.Println("Check Create LineageNode success...")
	}
	
	createLineageNodeObj,_ := createLineageNodeRequestResult.Data. (model.LineageNode)

	// --------------------------------------------------------------
	// Check LineageNode Obj ID
	// --------------------------------------------------------------	
	if createLineageNodeObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for LineageNode" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getLineageNodeRequestResult := dao.GetLineageNode( uint64(createLineageNodeObj.ID) )
	
	if getLineageNodeRequestResult.Success == false {
		t.Errorf(getLineageNodeRequestResult.Msg)
	} else {
		fmt.Println("Check Get LineageNode success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getLineageNodeObj,_ := getLineageNodeRequestResult.Data. (model.LineageNode)
	compareLineageNode := cmp.Equal(createLineageNodeObj.ID, getLineageNodeObj.ID)
	
	if  compareLineageNode == false	{
		t.Errorf( "Created LineageNode object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllLineageNodeRequestResult := dao.GetAllLineageNode()

	if getAllLineageNodeRequestResult.Success == false {
			t.Errorf(getAllLineageNodeRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll LineageNode success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllLineageNodeObj []model.LineageNode = getAllLineageNodeRequestResult.Data. ([]model.LineageNode)
		
	equalLineageNode := cmp.Equal(createLineageNodeObj.ID, getAllLineageNodeObj[len(getAllLineageNodeObj)-1].ID)
		
	if equalLineageNode == false {
		t.Errorf( "Created object is not equal to the last entry in LineageNode[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for LineageNode
	// --------------------------------------------------------------	
	deleteLineageNodeRequestResult := dao.DeleteLineageNode(uint64(createLineageNodeObj.ID))

	if deleteLineageNodeRequestResult.Success == false {
			t.Errorf(deleteLineageNodeRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion LineageNode success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getLineageNodeRequestResult = dao.GetLineageNode( uint64(createLineageNodeObj.ID) )
	
	if getLineageNodeRequestResult.Success == true {
		t.Errorf(getLineageNodeRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestTagCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Tag
	//----------------------------------------------------------------------------
	TagObj := model.Tag                                            {Name:"test value for Name",Category:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createTagRequestResult := dao.CreateTag( TagObj )
	
	if createTagRequestResult.Success == false {
		t.Errorf(createTagRequestResult.Msg)
	} else {
		fmt.Println("Check Create Tag success...")
	}
	
	createTagObj,_ := createTagRequestResult.Data. (model.Tag)

	// --------------------------------------------------------------
	// Check Tag Obj ID
	// --------------------------------------------------------------	
	if createTagObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Tag" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getTagRequestResult := dao.GetTag( uint64(createTagObj.ID) )
	
	if getTagRequestResult.Success == false {
		t.Errorf(getTagRequestResult.Msg)
	} else {
		fmt.Println("Check Get Tag success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getTagObj,_ := getTagRequestResult.Data. (model.Tag)
	compareTag := cmp.Equal(createTagObj.ID, getTagObj.ID)
	
	if  compareTag == false	{
		t.Errorf( "Created Tag object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllTagRequestResult := dao.GetAllTag()

	if getAllTagRequestResult.Success == false {
			t.Errorf(getAllTagRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Tag success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllTagObj []model.Tag = getAllTagRequestResult.Data. ([]model.Tag)
		
	equalTag := cmp.Equal(createTagObj.ID, getAllTagObj[len(getAllTagObj)-1].ID)
		
	if equalTag == false {
		t.Errorf( "Created object is not equal to the last entry in Tag[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Tag
	// --------------------------------------------------------------	
	deleteTagRequestResult := dao.DeleteTag(uint64(createTagObj.ID))

	if deleteTagRequestResult.Success == false {
			t.Errorf(deleteTagRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Tag success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getTagRequestResult = dao.GetTag( uint64(createTagObj.ID) )
	
	if getTagRequestResult.Success == true {
		t.Errorf(getTagRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAccessPolicyCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for AccessPolicy
	//----------------------------------------------------------------------------
	AccessPolicyObj := model.AccessPolicy                                                                                            {Name:"test value for Name",SubjectName:"test value for SubjectName",AccessLevel:0,SubjectType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAccessPolicyRequestResult := dao.CreateAccessPolicy( AccessPolicyObj )
	
	if createAccessPolicyRequestResult.Success == false {
		t.Errorf(createAccessPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Create AccessPolicy success...")
	}
	
	createAccessPolicyObj,_ := createAccessPolicyRequestResult.Data. (model.AccessPolicy)

	// --------------------------------------------------------------
	// Check AccessPolicy Obj ID
	// --------------------------------------------------------------	
	if createAccessPolicyObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for AccessPolicy" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAccessPolicyRequestResult := dao.GetAccessPolicy( uint64(createAccessPolicyObj.ID) )
	
	if getAccessPolicyRequestResult.Success == false {
		t.Errorf(getAccessPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Get AccessPolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAccessPolicyObj,_ := getAccessPolicyRequestResult.Data. (model.AccessPolicy)
	compareAccessPolicy := cmp.Equal(createAccessPolicyObj.ID, getAccessPolicyObj.ID)
	
	if  compareAccessPolicy == false	{
		t.Errorf( "Created AccessPolicy object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAccessPolicyRequestResult := dao.GetAllAccessPolicy()

	if getAllAccessPolicyRequestResult.Success == false {
			t.Errorf(getAllAccessPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll AccessPolicy success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAccessPolicyObj []model.AccessPolicy = getAllAccessPolicyRequestResult.Data. ([]model.AccessPolicy)
		
	equalAccessPolicy := cmp.Equal(createAccessPolicyObj.ID, getAllAccessPolicyObj[len(getAllAccessPolicyObj)-1].ID)
		
	if equalAccessPolicy == false {
		t.Errorf( "Created object is not equal to the last entry in AccessPolicy[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for AccessPolicy
	// --------------------------------------------------------------	
	deleteAccessPolicyRequestResult := dao.DeleteAccessPolicy(uint64(createAccessPolicyObj.ID))

	if deleteAccessPolicyRequestResult.Success == false {
			t.Errorf(deleteAccessPolicyRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion AccessPolicy success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAccessPolicyRequestResult = dao.GetAccessPolicy( uint64(createAccessPolicyObj.ID) )
	
	if getAccessPolicyRequestResult.Success == true {
		t.Errorf(getAccessPolicyRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestAlertCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Alert
	//----------------------------------------------------------------------------
	AlertObj := model.Alert                                                                                                                    {Title:"test value for Title",CreatedAt:time.Now(),Severity:0,Status:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createAlertRequestResult := dao.CreateAlert( AlertObj )
	
	if createAlertRequestResult.Success == false {
		t.Errorf(createAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Create Alert success...")
	}
	
	createAlertObj,_ := createAlertRequestResult.Data. (model.Alert)

	// --------------------------------------------------------------
	// Check Alert Obj ID
	// --------------------------------------------------------------	
	if createAlertObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Alert" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getAlertRequestResult := dao.GetAlert( uint64(createAlertObj.ID) )
	
	if getAlertRequestResult.Success == false {
		t.Errorf(getAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Get Alert success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getAlertObj,_ := getAlertRequestResult.Data. (model.Alert)
	compareAlert := cmp.Equal(createAlertObj.ID, getAlertObj.ID)
	
	if  compareAlert == false	{
		t.Errorf( "Created Alert object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllAlertRequestResult := dao.GetAllAlert()

	if getAllAlertRequestResult.Success == false {
			t.Errorf(getAllAlertRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Alert success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllAlertObj []model.Alert = getAllAlertRequestResult.Data. ([]model.Alert)
		
	equalAlert := cmp.Equal(createAlertObj.ID, getAllAlertObj[len(getAllAlertObj)-1].ID)
		
	if equalAlert == false {
		t.Errorf( "Created object is not equal to the last entry in Alert[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Alert
	// --------------------------------------------------------------	
	deleteAlertRequestResult := dao.DeleteAlert(uint64(createAlertObj.ID))

	if deleteAlertRequestResult.Success == false {
			t.Errorf(deleteAlertRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Alert success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getAlertRequestResult = dao.GetAlert( uint64(createAlertObj.ID) )
	
	if getAlertRequestResult.Success == true {
		t.Errorf(getAlertRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestSubscriberCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for Subscriber
	//----------------------------------------------------------------------------
	SubscriberObj := model.Subscriber                                                                            {Name:"test value for Name",Address:"test value for Address",Channel:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createSubscriberRequestResult := dao.CreateSubscriber( SubscriberObj )
	
	if createSubscriberRequestResult.Success == false {
		t.Errorf(createSubscriberRequestResult.Msg)
	} else {
		fmt.Println("Check Create Subscriber success...")
	}
	
	createSubscriberObj,_ := createSubscriberRequestResult.Data. (model.Subscriber)

	// --------------------------------------------------------------
	// Check Subscriber Obj ID
	// --------------------------------------------------------------	
	if createSubscriberObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for Subscriber" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getSubscriberRequestResult := dao.GetSubscriber( uint64(createSubscriberObj.ID) )
	
	if getSubscriberRequestResult.Success == false {
		t.Errorf(getSubscriberRequestResult.Msg)
	} else {
		fmt.Println("Check Get Subscriber success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getSubscriberObj,_ := getSubscriberRequestResult.Data. (model.Subscriber)
	compareSubscriber := cmp.Equal(createSubscriberObj.ID, getSubscriberObj.ID)
	
	if  compareSubscriber == false	{
		t.Errorf( "Created Subscriber object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllSubscriberRequestResult := dao.GetAllSubscriber()

	if getAllSubscriberRequestResult.Success == false {
			t.Errorf(getAllSubscriberRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll Subscriber success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllSubscriberObj []model.Subscriber = getAllSubscriberRequestResult.Data. ([]model.Subscriber)
		
	equalSubscriber := cmp.Equal(createSubscriberObj.ID, getAllSubscriberObj[len(getAllSubscriberObj)-1].ID)
		
	if equalSubscriber == false {
		t.Errorf( "Created object is not equal to the last entry in Subscriber[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for Subscriber
	// --------------------------------------------------------------	
	deleteSubscriberRequestResult := dao.DeleteSubscriber(uint64(createSubscriberObj.ID))

	if deleteSubscriberRequestResult.Success == false {
			t.Errorf(deleteSubscriberRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion Subscriber success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getSubscriberRequestResult = dao.GetSubscriber( uint64(createSubscriberObj.ID) )
	
	if getSubscriberRequestResult.Success == true {
		t.Errorf(getSubscriberRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestBusinessGlossaryTermCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for BusinessGlossaryTerm
	//----------------------------------------------------------------------------
	BusinessGlossaryTermObj := model.BusinessGlossaryTerm                                                                                            {Term:"test value for Term",Definition:"test value for Definition",Steward:"test value for Steward"}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createBusinessGlossaryTermRequestResult := dao.CreateBusinessGlossaryTerm( BusinessGlossaryTermObj )
	
	if createBusinessGlossaryTermRequestResult.Success == false {
		t.Errorf(createBusinessGlossaryTermRequestResult.Msg)
	} else {
		fmt.Println("Check Create BusinessGlossaryTerm success...")
	}
	
	createBusinessGlossaryTermObj,_ := createBusinessGlossaryTermRequestResult.Data. (model.BusinessGlossaryTerm)

	// --------------------------------------------------------------
	// Check BusinessGlossaryTerm Obj ID
	// --------------------------------------------------------------	
	if createBusinessGlossaryTermObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for BusinessGlossaryTerm" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getBusinessGlossaryTermRequestResult := dao.GetBusinessGlossaryTerm( uint64(createBusinessGlossaryTermObj.ID) )
	
	if getBusinessGlossaryTermRequestResult.Success == false {
		t.Errorf(getBusinessGlossaryTermRequestResult.Msg)
	} else {
		fmt.Println("Check Get BusinessGlossaryTerm success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getBusinessGlossaryTermObj,_ := getBusinessGlossaryTermRequestResult.Data. (model.BusinessGlossaryTerm)
	compareBusinessGlossaryTerm := cmp.Equal(createBusinessGlossaryTermObj.ID, getBusinessGlossaryTermObj.ID)
	
	if  compareBusinessGlossaryTerm == false	{
		t.Errorf( "Created BusinessGlossaryTerm object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllBusinessGlossaryTermRequestResult := dao.GetAllBusinessGlossaryTerm()

	if getAllBusinessGlossaryTermRequestResult.Success == false {
			t.Errorf(getAllBusinessGlossaryTermRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll BusinessGlossaryTerm success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllBusinessGlossaryTermObj []model.BusinessGlossaryTerm = getAllBusinessGlossaryTermRequestResult.Data. ([]model.BusinessGlossaryTerm)
		
	equalBusinessGlossaryTerm := cmp.Equal(createBusinessGlossaryTermObj.ID, getAllBusinessGlossaryTermObj[len(getAllBusinessGlossaryTermObj)-1].ID)
		
	if equalBusinessGlossaryTerm == false {
		t.Errorf( "Created object is not equal to the last entry in BusinessGlossaryTerm[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for BusinessGlossaryTerm
	// --------------------------------------------------------------	
	deleteBusinessGlossaryTermRequestResult := dao.DeleteBusinessGlossaryTerm(uint64(createBusinessGlossaryTermObj.ID))

	if deleteBusinessGlossaryTermRequestResult.Success == false {
			t.Errorf(deleteBusinessGlossaryTermRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion BusinessGlossaryTerm success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getBusinessGlossaryTermRequestResult = dao.GetBusinessGlossaryTerm( uint64(createBusinessGlossaryTermObj.ID) )
	
	if getBusinessGlossaryTermRequestResult.Success == true {
		t.Errorf(getBusinessGlossaryTermRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestRecommendationScenarioCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for RecommendationScenario
	//----------------------------------------------------------------------------
	RecommendationScenarioObj := model.RecommendationScenario                                                                            {Name:"test value for Name",Objective:"test value for Objective",RecommendationType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createRecommendationScenarioRequestResult := dao.CreateRecommendationScenario( RecommendationScenarioObj )
	
	if createRecommendationScenarioRequestResult.Success == false {
		t.Errorf(createRecommendationScenarioRequestResult.Msg)
	} else {
		fmt.Println("Check Create RecommendationScenario success...")
	}
	
	createRecommendationScenarioObj,_ := createRecommendationScenarioRequestResult.Data. (model.RecommendationScenario)

	// --------------------------------------------------------------
	// Check RecommendationScenario Obj ID
	// --------------------------------------------------------------	
	if createRecommendationScenarioObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for RecommendationScenario" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getRecommendationScenarioRequestResult := dao.GetRecommendationScenario( uint64(createRecommendationScenarioObj.ID) )
	
	if getRecommendationScenarioRequestResult.Success == false {
		t.Errorf(getRecommendationScenarioRequestResult.Msg)
	} else {
		fmt.Println("Check Get RecommendationScenario success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getRecommendationScenarioObj,_ := getRecommendationScenarioRequestResult.Data. (model.RecommendationScenario)
	compareRecommendationScenario := cmp.Equal(createRecommendationScenarioObj.ID, getRecommendationScenarioObj.ID)
	
	if  compareRecommendationScenario == false	{
		t.Errorf( "Created RecommendationScenario object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllRecommendationScenarioRequestResult := dao.GetAllRecommendationScenario()

	if getAllRecommendationScenarioRequestResult.Success == false {
			t.Errorf(getAllRecommendationScenarioRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll RecommendationScenario success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllRecommendationScenarioObj []model.RecommendationScenario = getAllRecommendationScenarioRequestResult.Data. ([]model.RecommendationScenario)
		
	equalRecommendationScenario := cmp.Equal(createRecommendationScenarioObj.ID, getAllRecommendationScenarioObj[len(getAllRecommendationScenarioObj)-1].ID)
		
	if equalRecommendationScenario == false {
		t.Errorf( "Created object is not equal to the last entry in RecommendationScenario[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for RecommendationScenario
	// --------------------------------------------------------------	
	deleteRecommendationScenarioRequestResult := dao.DeleteRecommendationScenario(uint64(createRecommendationScenarioObj.ID))

	if deleteRecommendationScenarioRequestResult.Success == false {
			t.Errorf(deleteRecommendationScenarioRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion RecommendationScenario success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getRecommendationScenarioRequestResult = dao.GetRecommendationScenario( uint64(createRecommendationScenarioObj.ID) )
	
	if getRecommendationScenarioRequestResult.Success == true {
		t.Errorf(getRecommendationScenarioRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFraudScenarioCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FraudScenario
	//----------------------------------------------------------------------------
	FraudScenarioObj := model.FraudScenario                                                                            {Name:"test value for Name",RiskAppetite:"test value for RiskAppetite",DetectionType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFraudScenarioRequestResult := dao.CreateFraudScenario( FraudScenarioObj )
	
	if createFraudScenarioRequestResult.Success == false {
		t.Errorf(createFraudScenarioRequestResult.Msg)
	} else {
		fmt.Println("Check Create FraudScenario success...")
	}
	
	createFraudScenarioObj,_ := createFraudScenarioRequestResult.Data. (model.FraudScenario)

	// --------------------------------------------------------------
	// Check FraudScenario Obj ID
	// --------------------------------------------------------------	
	if createFraudScenarioObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for FraudScenario" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFraudScenarioRequestResult := dao.GetFraudScenario( uint64(createFraudScenarioObj.ID) )
	
	if getFraudScenarioRequestResult.Success == false {
		t.Errorf(getFraudScenarioRequestResult.Msg)
	} else {
		fmt.Println("Check Get FraudScenario success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFraudScenarioObj,_ := getFraudScenarioRequestResult.Data. (model.FraudScenario)
	compareFraudScenario := cmp.Equal(createFraudScenarioObj.ID, getFraudScenarioObj.ID)
	
	if  compareFraudScenario == false	{
		t.Errorf( "Created FraudScenario object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFraudScenarioRequestResult := dao.GetAllFraudScenario()

	if getAllFraudScenarioRequestResult.Success == false {
			t.Errorf(getAllFraudScenarioRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FraudScenario success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFraudScenarioObj []model.FraudScenario = getAllFraudScenarioRequestResult.Data. ([]model.FraudScenario)
		
	equalFraudScenario := cmp.Equal(createFraudScenarioObj.ID, getAllFraudScenarioObj[len(getAllFraudScenarioObj)-1].ID)
		
	if equalFraudScenario == false {
		t.Errorf( "Created object is not equal to the last entry in FraudScenario[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FraudScenario
	// --------------------------------------------------------------	
	deleteFraudScenarioRequestResult := dao.DeleteFraudScenario(uint64(createFraudScenarioObj.ID))

	if deleteFraudScenarioRequestResult.Success == false {
			t.Errorf(deleteFraudScenarioRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FraudScenario success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFraudScenarioRequestResult = dao.GetFraudScenario( uint64(createFraudScenarioObj.ID) )
	
	if getFraudScenarioRequestResult.Success == true {
		t.Errorf(getFraudScenarioRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}


func TestFraudSignalCRUD(t *testing.T) {

	//----------------------------------------------------------------------------
	// Test CRUD for FraudSignal
	//----------------------------------------------------------------------------
	FraudSignalObj := model.FraudSignal                                                                            {Name:"test value for Name",RuleLogic:"test value for RuleLogic",SignalType:0}

	// --------------------------------------------------------------
	// Check Create
	// --------------------------------------------------------------
	createFraudSignalRequestResult := dao.CreateFraudSignal( FraudSignalObj )
	
	if createFraudSignalRequestResult.Success == false {
		t.Errorf(createFraudSignalRequestResult.Msg)
	} else {
		fmt.Println("Check Create FraudSignal success...")
	}
	
	createFraudSignalObj,_ := createFraudSignalRequestResult.Data. (model.FraudSignal)

	// --------------------------------------------------------------
	// Check FraudSignal Obj ID
	// --------------------------------------------------------------	
	if createFraudSignalObj.ID == 0 {
	    t.Errorf( "The ORM failed to assign and ID for FraudSignal" )
	}	

	// --------------------------------------------------------------
	// Check Get
	// --------------------------------------------------------------	
	getFraudSignalRequestResult := dao.GetFraudSignal( uint64(createFraudSignalObj.ID) )
	
	if getFraudSignalRequestResult.Success == false {
		t.Errorf(getFraudSignalRequestResult.Msg)
	} else {
		fmt.Println("Check Get FraudSignal success...")
	}
	
	// --------------------------------------------------------------
	// Check returned struct from Get equals original created obj
	// --------------------------------------------------------------	
	getFraudSignalObj,_ := getFraudSignalRequestResult.Data. (model.FraudSignal)
	compareFraudSignal := cmp.Equal(createFraudSignalObj.ID, getFraudSignalObj.ID)
	
	if  compareFraudSignal == false	{
		t.Errorf( "Created FraudSignal object is not equal to read object." )
	}
	
	// --------------------------------------------------------------
	// Check GetAll
	// --------------------------------------------------------------	
	getAllFraudSignalRequestResult := dao.GetAllFraudSignal()

	if getAllFraudSignalRequestResult.Success == false {
			t.Errorf(getAllFraudSignalRequestResult.Msg)
	} else {
		fmt.Println("Check GetAll FraudSignal success...")
	}
	
	// --------------------------------------------------------------
	// Check GetAll returns an array with zero index equal 
	// to initially created object
	// --------------------------------------------------------------		
	var getAllFraudSignalObj []model.FraudSignal = getAllFraudSignalRequestResult.Data. ([]model.FraudSignal)
		
	equalFraudSignal := cmp.Equal(createFraudSignalObj.ID, getAllFraudSignalObj[len(getAllFraudSignalObj)-1].ID)
		
	if equalFraudSignal == false {
		t.Errorf( "Created object is not equal to the last entry in FraudSignal[] returned by GetAll" )
    }
    
	// --------------------------------------------------------------
	// Check deletion for FraudSignal
	// --------------------------------------------------------------	
	deleteFraudSignalRequestResult := dao.DeleteFraudSignal(uint64(createFraudSignalObj.ID))

	if deleteFraudSignalRequestResult.Success == false {
			t.Errorf(deleteFraudSignalRequestResult.Msg)
	} else {
		fmt.Println("Check Deletion FraudSignal success...")
	}


	// --------------------------------------------------------------
	// Check deletion causes Get to fail
	// --------------------------------------------------------------		
	getFraudSignalRequestResult = dao.GetFraudSignal( uint64(createFraudSignalObj.ID) )
	
	if getFraudSignalRequestResult.Success == true {
		t.Errorf(getFraudSignalRequestResult.Msg)
	} else {
		fmt.Println("Validate deletion success...")
	}	
	
}

