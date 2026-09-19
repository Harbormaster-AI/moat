require "test_helper"

class BusinessGlossaryTermControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @businessGlossaryTerm = businessGlossaryTerms(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create businessGlossaryTerm" do
    assert_difference("BusinessGlossaryTerm.count") do
      post businessGlossaryTerms_url, params: { businessGlossaryTerm: { term:"test string for term", definition:"test string for definition", steward:"test string for steward" } }
    end

    assert_redirected_to businessGlossaryTerms_url
  end

 
  
  test "should destroy businessGlossaryTerm" do
    assert_difference("BusinessGlossaryTerm.count", -1) do
      delete businessGlossaryTerm_url(@businessGlossaryTerm)
    end

    assert_redirected_to businessGlossaryTerms_url
  end
  
end


