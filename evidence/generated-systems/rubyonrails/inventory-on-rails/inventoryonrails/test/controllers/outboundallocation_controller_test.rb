require "test_helper"

class OutboundAllocationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @outboundAllocation = outboundAllocations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create outboundAllocation" do
    assert_difference("OutboundAllocation.count") do
      post outboundAllocations_url, params: { outboundAllocation: { allocationNumber:"test string for allocationNumber", allocatedQuantity:"test value", allocationDate:1.week.ago, Status:OutboundAllocation.Statuss[0] } }
    end

    assert_redirected_to outboundAllocations_url
  end

 
  
  test "should destroy outboundAllocation" do
    assert_difference("OutboundAllocation.count", -1) do
      delete outboundAllocation_url(@outboundAllocation)
    end

    assert_redirected_to outboundAllocations_url
  end
  
end


