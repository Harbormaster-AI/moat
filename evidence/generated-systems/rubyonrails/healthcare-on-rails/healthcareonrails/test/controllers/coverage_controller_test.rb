require "test_helper"

class CoverageControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @coverage = coverages(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create coverage" do
    assert_difference("Coverage.count") do
      post coverages_url, params: { coverage: { memberId:"test string for memberId", groupNumber:"test string for groupNumber", effectiveDate:1.week.ago, endDate:1.week.ago, CoverageType:Coverage.CoverageTypes[0] } }
    end

    assert_redirected_to coverages_url
  end

 
  
  test "should destroy coverage" do
    assert_difference("Coverage.count", -1) do
      delete coverage_url(@coverage)
    end

    assert_redirected_to coverages_url
  end
  
end


