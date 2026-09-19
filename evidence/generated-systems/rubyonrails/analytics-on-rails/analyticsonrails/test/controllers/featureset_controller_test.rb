require "test_helper"

class FeatureSetControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @featureSet = featureSets(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create featureSet" do
    assert_difference("FeatureSet.count") do
      post featureSets_url, params: { featureSet: { name:"test string for name", refreshSchedule:"test value", StoreType:FeatureSet.StoreTypes[0] } }
    end

    assert_redirected_to featureSets_url
  end

 
  
  test "should destroy featureSet" do
    assert_difference("FeatureSet.count", -1) do
      delete featureSet_url(@featureSet)
    end

    assert_redirected_to featureSets_url
  end
  
end


