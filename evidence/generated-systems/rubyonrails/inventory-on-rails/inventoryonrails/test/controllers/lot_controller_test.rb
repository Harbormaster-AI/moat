require "test_helper"

class LotControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @lot = lots(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create lot" do
    assert_difference("Lot.count") do
      post lots_url, params: { lot: { batchNumber:"test value", manufactureDate:1.week.ago, expirationDate:1.week.ago, LotStatus:Lot.LotStatuss[0] } }
    end

    assert_redirected_to lots_url
  end

 
  
  test "should destroy lot" do
    assert_difference("Lot.count", -1) do
      delete lot_url(@lot)
    end

    assert_redirected_to lots_url
  end
  
end


