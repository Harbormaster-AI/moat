require "test_helper"

class DataPipelineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dataPipeline = dataPipelines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dataPipeline" do
    assert_difference("DataPipeline.count") do
      post dataPipelines_url, params: { dataPipeline: { name:"test string for name", schedule:"test value", TriggerType:DataPipeline.TriggerTypes[0], Status:DataPipeline.Statuss[0] } }
    end

    assert_redirected_to dataPipelines_url
  end

 
  
  test "should destroy dataPipeline" do
    assert_difference("DataPipeline.count", -1) do
      delete dataPipeline_url(@dataPipeline)
    end

    assert_redirected_to dataPipelines_url
  end
  
end


