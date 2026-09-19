require "test_helper"

class ThirdPartyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @thirdParty = thirdPartys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create thirdParty" do
    assert_difference("ThirdParty.count") do
      post thirdPartys_url, params: { thirdParty: { name:"test string for name", country:"test string for country", contactEmail:"test value", ThirdPartyType:ThirdParty.ThirdPartyTypes[0], Criticality:ThirdParty.Criticalitys[0] } }
    end

    assert_redirected_to thirdPartys_url
  end

 
  
  test "should destroy thirdParty" do
    assert_difference("ThirdParty.count", -1) do
      delete thirdParty_url(@thirdParty)
    end

    assert_redirected_to thirdPartys_url
  end
  
end


