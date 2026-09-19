require "test_helper"

class QualitySpecificationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @qualitySpecification = qualitySpecifications(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create qualitySpecification" do
    assert_difference("QualitySpecification.count") do
      post qualitySpecifications_url, params: { qualitySpecification: { specCode:"test string for specCode", name:"test string for name", version:"test string for version" } }
    end

    assert_redirected_to qualitySpecifications_url
  end

 
  
  test "should destroy qualitySpecification" do
    assert_difference("QualitySpecification.count", -1) do
      delete qualitySpecification_url(@qualitySpecification)
    end

    assert_redirected_to qualitySpecifications_url
  end
  
end


