require "test_helper"

class ModelControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @model = models(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create model" do
    assert_difference("Model.count") do
      post models_url, params: { model: { name:"test string for name", taskDescription:"test string for taskDescription", ModelType:Model.ModelTypes[0] } }
    end

    assert_redirected_to models_url
  end

 
  
  test "should destroy model" do
    assert_difference("Model.count", -1) do
      delete model_url(@model)
    end

    assert_redirected_to models_url
  end
  
end


