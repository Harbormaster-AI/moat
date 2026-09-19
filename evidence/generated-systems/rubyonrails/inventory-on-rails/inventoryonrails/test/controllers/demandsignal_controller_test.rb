require "test_helper"

class DemandSignalControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @demandSignal = demandSignals(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create demandSignal" do
    assert_difference("DemandSignal.count") do
      post demandSignals_url, params: { demandSignal: { externalReference:"test string for externalReference", requestedDate:1.week.ago, quantity:"test value", DemandType:DemandSignal.DemandTypes[0] } }
    end

    assert_redirected_to demandSignals_url
  end

 
  
  test "should destroy demandSignal" do
    assert_difference("DemandSignal.count", -1) do
      delete demandSignal_url(@demandSignal)
    end

    assert_redirected_to demandSignals_url
  end
  
end


