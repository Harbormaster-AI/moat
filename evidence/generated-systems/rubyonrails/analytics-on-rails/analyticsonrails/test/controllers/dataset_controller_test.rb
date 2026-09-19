require "test_helper"

class DataSetControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dataSet = dataSets(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dataSet" do
    assert_difference("DataSet.count") do
      post dataSets_url, params: { dataSet: { name:"test string for name", schemaVersion:"test string for schemaVersion", refreshSchedule:"test value", sensitive:true, DataFormat:DataSet.DataFormats[0] } }
    end

    assert_redirected_to dataSets_url
  end

 
  
  test "should destroy dataSet" do
    assert_difference("DataSet.count", -1) do
      delete dataSet_url(@dataSet)
    end

    assert_redirected_to dataSets_url
  end
  
end


