require "test_helper"

class CarrierServiceControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @carrierService = carrierServices(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create carrierService" do
    assert_difference("CarrierService.count") do
      post carrierServices_url, params: { carrierService: { name:"test string for name", code:"test string for code", Carrier:CarrierService.Carriers[0], ServiceLevel:CarrierService.ServiceLevels[0] } }
    end

    assert_redirected_to carrierServices_url
  end

 
  
  test "should destroy carrierService" do
    assert_difference("CarrierService.count", -1) do
      delete carrierService_url(@carrierService)
    end

    assert_redirected_to carrierServices_url
  end
  
end


