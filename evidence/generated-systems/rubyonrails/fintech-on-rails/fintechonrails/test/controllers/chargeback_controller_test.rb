require "test_helper"

class ChargebackControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @chargeback = chargebacks(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create chargeback" do
    assert_difference("Chargeback.count") do
      post chargebacks_url, params: { chargeback: { chargebackReference:"test string for chargebackReference", amount:"test value", postedAt:1.week.ago, Stage:Chargeback.Stages[0], Status:Chargeback.Statuss[0] } }
    end

    assert_redirected_to chargebacks_url
  end

 
  
  test "should destroy chargeback" do
    assert_difference("Chargeback.count", -1) do
      delete chargeback_url(@chargeback)
    end

    assert_redirected_to chargebacks_url
  end
  
end


