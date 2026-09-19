require "test_helper"

class SemanticModelControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @semanticModel = semanticModels(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create semanticModel" do
    assert_difference("SemanticModel.count") do
      post semanticModels_url, params: { semanticModel: { name:"test string for name", version:"test string for version", grain:"test string for grain" } }
    end

    assert_redirected_to semanticModels_url
  end

 
  
  test "should destroy semanticModel" do
    assert_difference("SemanticModel.count", -1) do
      delete semanticModel_url(@semanticModel)
    end

    assert_redirected_to semanticModels_url
  end
  
end


