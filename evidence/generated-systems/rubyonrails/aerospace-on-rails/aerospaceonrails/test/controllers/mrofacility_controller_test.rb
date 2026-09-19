require "test_helper"

class MROFacilityControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @mROFacility = mROFacilitys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create mROFacility" do
    assert_difference("MROFacility.count") do
      post mROFacilitys_url, params: { mROFacility: { name:"test string for name", approvalScope:"test string for approvalScope", address:"test value" } }
    end

    assert_redirected_to mROFacilitys_url
  end

 
  
  test "should destroy mROFacility" do
    assert_difference("MROFacility.count", -1) do
      delete mROFacility_url(@mROFacility)
    end

    assert_redirected_to mROFacilitys_url
  end
  
end


