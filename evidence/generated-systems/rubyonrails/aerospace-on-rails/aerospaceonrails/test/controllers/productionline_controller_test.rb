require "test_helper"

class ProductionLineControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @productionLine = productionLines(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create productionLine" do
    assert_difference("ProductionLine.count") do
      post productionLines_url, params: { productionLine: { name:"test string for name", LineType:ProductionLine.LineTypes[0] } }
    end

    assert_redirected_to productionLines_url
  end

 
  
  test "should destroy productionLine" do
    assert_difference("ProductionLine.count", -1) do
      delete productionLine_url(@productionLine)
    end

    assert_redirected_to productionLines_url
  end
  
end


