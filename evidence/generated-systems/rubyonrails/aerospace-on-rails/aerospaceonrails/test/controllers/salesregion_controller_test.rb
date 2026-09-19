require "test_helper"

class SalesRegionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @salesRegion = salesRegions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create salesRegion" do
    assert_difference("SalesRegion.count") do
      post salesRegions_url, params: { salesRegion: { name:"test string for name", regionCode:"test string for regionCode" } }
    end

    assert_redirected_to salesRegions_url
  end

 
  
  test "should destroy salesRegion" do
    assert_difference("SalesRegion.count", -1) do
      delete salesRegion_url(@salesRegion)
    end

    assert_redirected_to salesRegions_url
  end
  
end


