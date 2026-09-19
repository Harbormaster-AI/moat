require "test_helper"

class CompensationPackageControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @compensationPackage = compensationPackages(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create compensationPackage" do
    assert_difference("CompensationPackage.count") do
      post compensationPackages_url, params: { compensationPackage: { effectiveFrom:1.week.ago, effectiveTo:1.week.ago, currency:"test string for currency" } }
    end

    assert_redirected_to compensationPackages_url
  end

 
  
  test "should destroy compensationPackage" do
    assert_difference("CompensationPackage.count", -1) do
      delete compensationPackage_url(@compensationPackage)
    end

    assert_redirected_to compensationPackages_url
  end
  
end


