require "test_helper"

class Record_ControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @record_ = record_s(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create record_" do
    assert_difference("Record_.count") do
      post record_s_url, params: { record_: { title:"test string for title", creationDate:1.week.ago, RecordType:Record_.RecordTypes[0], Classification:Record_.Classifications[0], Status:Record_.Statuss[0] } }
    end

    assert_redirected_to record_s_url
  end

 
  
  test "should destroy record_" do
    assert_difference("Record_.count", -1) do
      delete record__url(@record_)
    end

    assert_redirected_to record_s_url
  end
  
end


