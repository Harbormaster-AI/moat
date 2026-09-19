require "test_helper"

class ChannelControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @channel = channels(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create channel" do
    assert_difference("Channel.count") do
      post channels_url, params: { channel: { name:"test string for name", channelCode:"test string for channelCode", locale:"test string for locale", domain:"test string for domain", asActive:true, defaultCurrency:"test string for defaultCurrency", ChannelType:Channel.ChannelTypes[0] } }
    end

    assert_redirected_to channels_url
  end

 
  
  test "should destroy channel" do
    assert_difference("Channel.count", -1) do
      delete channel_url(@channel)
    end

    assert_redirected_to channels_url
  end
  
end


