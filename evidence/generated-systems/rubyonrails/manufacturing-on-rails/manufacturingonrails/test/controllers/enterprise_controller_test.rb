require "test_helper"

class EnterpriseControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @enterprise = enterprises(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create enterprise" do
    assert_difference("Enterprise.count") do
      post enterprises_url, params: { enterprise: { name:"test string for name", legalName:"test string for legalName", registrationCountry:"test string for registrationCountry", website:"test string for website", taxId:"test string for taxId" } }
    end

    assert_redirected_to enterprises_url
  end

 
  
  test "should destroy enterprise" do
    assert_difference("Enterprise.count", -1) do
      delete enterprise_url(@enterprise)
    end

    assert_redirected_to enterprises_url
  end
  
end


