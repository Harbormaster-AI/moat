require "test_helper"

class PharmacyControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @pharmacy = pharmacys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create pharmacy" do
    assert_difference("Pharmacy.count") do
      post pharmacys_url, params: { pharmacy: { name:"test string for name" } }
    end

    assert_redirected_to pharmacys_url
  end

 
  
  test "should destroy pharmacy" do
    assert_difference("Pharmacy.count", -1) do
      delete pharmacy_url(@pharmacy)
    end

    assert_redirected_to pharmacys_url
  end
  
end


