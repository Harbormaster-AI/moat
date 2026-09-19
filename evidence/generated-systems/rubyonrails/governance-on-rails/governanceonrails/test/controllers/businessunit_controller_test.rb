require "test_helper"

class BusinessUnitControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @businessUnit = businessUnits(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create businessUnit" do
    assert_difference("BusinessUnit.count") do
      post businessUnits_url, params: { businessUnit: { name:"test string for name", leader:"test string for leader" } }
    end

    assert_redirected_to businessUnits_url
  end

 
  
  test "should destroy businessUnit" do
    assert_difference("BusinessUnit.count", -1) do
      delete businessUnit_url(@businessUnit)
    end

    assert_redirected_to businessUnits_url
  end
  
end


