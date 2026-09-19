require "test_helper"

class ModelVersionControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @modelVersion = modelVersions(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create modelVersion" do
    assert_difference("ModelVersion.count") do
      post modelVersions_url, params: { modelVersion: { version:"test string for version", Lifecycle:ModelVersion.Lifecycles[0], TrainingStatus:ModelVersion.TrainingStatuss[0] } }
    end

    assert_redirected_to modelVersions_url
  end

 
  
  test "should destroy modelVersion" do
    assert_difference("ModelVersion.count", -1) do
      delete modelVersion_url(@modelVersion)
    end

    assert_redirected_to modelVersions_url
  end
  
end


