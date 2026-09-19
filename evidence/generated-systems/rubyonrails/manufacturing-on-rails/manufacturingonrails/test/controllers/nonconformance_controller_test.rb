require "test_helper"

class NonconformanceControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @nonconformance = nonconformances(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create nonconformance" do
    assert_difference("Nonconformance.count") do
      post nonconformances_url, params: { nonconformance: { ncNumber:"test string for ncNumber", description:"test string for description", containmentAction:"test string for containmentAction", NcType:Nonconformance.NcTypes[0], Severity:Nonconformance.Severitys[0], Status:Nonconformance.Statuss[0] } }
    end

    assert_redirected_to nonconformances_url
  end

 
  
  test "should destroy nonconformance" do
    assert_difference("Nonconformance.count", -1) do
      delete nonconformance_url(@nonconformance)
    end

    assert_redirected_to nonconformances_url
  end
  
end


