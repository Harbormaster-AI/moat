require "test_helper"

class FXDealControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @fXDeal = fXDeals(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create fXDeal" do
    assert_difference("FXDeal.count") do
      post fXDeals_url, params: { fXDeal: { dealReference:"test string for dealReference", baseCurrency:"test string for baseCurrency", quoteCurrency:"test string for quoteCurrency", rate:"test value", amount:"test value", settlementDate:1.week.ago, Status:FXDeal.Statuss[0] } }
    end

    assert_redirected_to fXDeals_url
  end

 
  
  test "should destroy fXDeal" do
    assert_difference("FXDeal.count", -1) do
      delete fXDeal_url(@fXDeal)
    end

    assert_redirected_to fXDeals_url
  end
  
end


