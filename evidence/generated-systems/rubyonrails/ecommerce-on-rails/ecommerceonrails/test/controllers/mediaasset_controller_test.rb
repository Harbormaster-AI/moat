require "test_helper"

class MediaAssetControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @mediaAsset = mediaAssets(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create mediaAsset" do
    assert_difference("MediaAsset.count") do
      post mediaAssets_url, params: { mediaAsset: { url:"test string for url", altText:"test string for altText", position:100, MediaType:MediaAsset.MediaTypes[0] } }
    end

    assert_redirected_to mediaAssets_url
  end

 
  
  test "should destroy mediaAsset" do
    assert_difference("MediaAsset.count", -1) do
      delete mediaAsset_url(@mediaAsset)
    end

    assert_redirected_to mediaAssets_url
  end
  
end


