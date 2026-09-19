require "test_helper"

class BonusPlanControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @bonusPlan = bonusPlans(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create bonusPlan" do
    assert_difference("BonusPlan.count") do
      post bonusPlans_url, params: { bonusPlan: { name:"test string for name", targetPercentage:"test value" } }
    end

    assert_redirected_to bonusPlans_url
  end

 
  
  test "should destroy bonusPlan" do
    assert_difference("BonusPlan.count", -1) do
      delete bonusPlan_url(@bonusPlan)
    end

    assert_redirected_to bonusPlans_url
  end
  
end


