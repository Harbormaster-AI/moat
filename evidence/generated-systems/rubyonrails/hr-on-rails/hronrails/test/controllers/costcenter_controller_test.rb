require "test_helper"

class CostCenterControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @costCenter = costCenters(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create costCenter" do
    assert_difference("CostCenter.count") do
      post costCenters_url, params: { costCenter: { code:"test string for code", name:"test string for name" } }
    end

    assert_redirected_to costCenters_url
  end

 
  
  test "should destroy costCenter" do
    assert_difference("CostCenter.count", -1) do
      delete costCenter_url(@costCenter)
    end

    assert_redirected_to costCenters_url
  end
  
end


