require "test_helper"

class AssetControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @asset = assets(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create asset" do
    assert_difference("Asset.count") do
      post assets_url, params: { asset: { assetTag:"test string for assetTag", assetName:"test string for assetName", commissioningDate:1.week.ago, AssetStatus:Asset.AssetStatuss[0] } }
    end

    assert_redirected_to assets_url
  end

 
  
  test "should destroy asset" do
    assert_difference("Asset.count", -1) do
      delete asset_url(@asset)
    end

    assert_redirected_to assets_url
  end
  
end


