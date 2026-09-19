require "test_helper"

class EmailMessageControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @emailMessage = emailMessages(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create emailMessage" do
    assert_difference("EmailMessage.count") do
      post emailMessages_url, params: { emailMessage: { subject:"test string for subject", body:"test string for body", sentAt:1.week.ago, messageId:"test string for messageId", Direction:EmailMessage.Directions[0], Status:EmailMessage.Statuss[0] } }
    end

    assert_redirected_to emailMessages_url
  end

 
  
  test "should destroy emailMessage" do
    assert_difference("EmailMessage.count", -1) do
      delete emailMessage_url(@emailMessage)
    end

    assert_redirected_to emailMessages_url
  end
  
end


