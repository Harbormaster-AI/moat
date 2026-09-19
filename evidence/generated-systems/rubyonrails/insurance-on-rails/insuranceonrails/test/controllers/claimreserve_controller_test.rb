require "test_helper"

class ClaimReserveControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @claimReserve = claimReserves(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create claimReserve" do
    assert_difference("ClaimReserve.count") do
      post claimReserves_url, params: { claimReserve: { amount:"test value", setDate:1.week.ago, ReserveType:ClaimReserve.ReserveTypes[0], Status:ClaimReserve.Statuss[0] } }
    end

    assert_redirected_to claimReserves_url
  end

 
  
  test "should destroy claimReserve" do
    assert_difference("ClaimReserve.count", -1) do
      delete claimReserve_url(@claimReserve)
    end

    assert_redirected_to claimReserves_url
  end
  
end


