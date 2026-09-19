require "test_helper"

class DataSourceControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @dataSource = dataSources(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create dataSource" do
    assert_difference("DataSource.count") do
      post dataSources_url, params: { dataSource: { name:"test string for name", connection:"test value", streaming:true, SourceType:DataSource.SourceTypes[0], Format:DataSource.Formats[0] } }
    end

    assert_redirected_to dataSources_url
  end

 
  
  test "should destroy dataSource" do
    assert_difference("DataSource.count", -1) do
      delete dataSource_url(@dataSource)
    end

    assert_redirected_to dataSources_url
  end
  
end


