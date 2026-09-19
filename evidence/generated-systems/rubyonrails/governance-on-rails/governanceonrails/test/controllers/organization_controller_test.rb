require "test_helper"

class OrganizationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @organization = organizations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create organization" do
    assert_difference("Organization.count") do
      post organizations_url, params: { organization: { name:"test string for name", legalName:"test string for legalName", jurisdiction:"test string for jurisdiction", industrySector:"test string for industrySector" } }
    end

    assert_redirected_to organizations_url
  end

 
  
  test "should destroy organization" do
    assert_difference("Organization.count", -1) do
      delete organization_url(@organization)
    end

    assert_redirected_to organizations_url
  end
  
end


