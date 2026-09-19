require "test_helper"

class DischargeControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @discharge = discharges(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create discharge" do
    assert_difference("Discharge.count") do
      post discharges_url, params: { discharge: { dischargeDateTime:1.week.ago, Disposition:Discharge.Dispositions[0] } }
    end

    assert_redirected_to discharges_url
  end

 
  
  test "should destroy discharge" do
    assert_difference("Discharge.count", -1) do
      delete discharge_url(@discharge)
    end

    assert_redirected_to discharges_url
  end
  
end


