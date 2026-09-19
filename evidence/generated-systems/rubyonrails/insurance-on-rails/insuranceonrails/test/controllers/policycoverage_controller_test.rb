require "test_helper"

class PolicyCoverageControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @policyCoverage = policyCoverages(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create policyCoverage" do
    assert_difference("PolicyCoverage.count") do
      post policyCoverages_url, params: { policyCoverage: { limit:"test value", deductible:"test value", premium:"test value", CoverageType:PolicyCoverage.CoverageTypes[0] } }
    end

    assert_redirected_to policyCoverages_url
  end

 
  
  test "should destroy policyCoverage" do
    assert_difference("PolicyCoverage.count", -1) do
      delete policyCoverage_url(@policyCoverage)
    end

    assert_redirected_to policyCoverages_url
  end
  
end


