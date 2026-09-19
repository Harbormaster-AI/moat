require "test_helper"

class SalesCampaignControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @salesCampaign = salesCampaigns(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create salesCampaign" do
    assert_difference("SalesCampaign.count") do
      post salesCampaigns_url, params: { salesCampaign: { campaignCode:"test string for campaignCode", Status:SalesCampaign.Statuss[0] } }
    end

    assert_redirected_to salesCampaigns_url
  end

 
  
  test "should destroy salesCampaign" do
    assert_difference("SalesCampaign.count", -1) do
      delete salesCampaign_url(@salesCampaign)
    end

    assert_redirected_to salesCampaigns_url
  end
  
end


