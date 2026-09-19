require "test_helper"

class CardTokenizationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @cardTokenization = cardTokenizations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create cardTokenization" do
    assert_difference("CardTokenization.count") do
      post cardTokenizations_url, params: { cardTokenization: { tokenReference:"test string for tokenReference", createdAt:1.week.ago, WalletProvider:CardTokenization.WalletProviders[0], Status:CardTokenization.Statuss[0] } }
    end

    assert_redirected_to cardTokenizations_url
  end

 
  
  test "should destroy cardTokenization" do
    assert_difference("CardTokenization.count", -1) do
      delete cardTokenization_url(@cardTokenization)
    end

    assert_redirected_to cardTokenizations_url
  end
  
end


