require "test_helper"

class TerritoryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @territory = territorys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create territory" do
    assert_difference("Territory.count") do
      post territorys_url, params: { territory: { name:"test string for name", region:"test string for region", TerritoryType:Territory.TerritoryTypes[0] } }
    end

    assert_redirected_to territorys_url
  end

 
  
  test "should destroy territory" do
    assert_difference("Territory.count", -1) do
      delete territory_url(@territory)
    end

    assert_redirected_to territorys_url
  end
  
end


