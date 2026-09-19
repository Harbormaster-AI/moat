require "test_helper"

class DistributorControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @distributor = distributors(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create distributor" do
    assert_difference("Distributor.count") do
      post distributors_url, params: { distributor: { name:"test string for name", licenseNumber:"test string for licenseNumber", region:"test string for region", DistributorType:Distributor.DistributorTypes[0] } }
    end

    assert_redirected_to distributors_url
  end

 
  
  test "should destroy distributor" do
    assert_difference("Distributor.count", -1) do
      delete distributor_url(@distributor)
    end

    assert_redirected_to distributors_url
  end
  
end


