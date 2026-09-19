require "test_helper"

class SerialNumberControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @serialNumber = serialNumbers(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create serialNumber" do
    assert_difference("SerialNumber.count") do
      post serialNumbers_url, params: { serialNumber: { serial:"test value", activationDate:1.week.ago, Status:SerialNumber.Statuss[0] } }
    end

    assert_redirected_to serialNumbers_url
  end

 
  
  test "should destroy serialNumber" do
    assert_difference("SerialNumber.count", -1) do
      delete serialNumber_url(@serialNumber)
    end

    assert_redirected_to serialNumbers_url
  end
  
end


