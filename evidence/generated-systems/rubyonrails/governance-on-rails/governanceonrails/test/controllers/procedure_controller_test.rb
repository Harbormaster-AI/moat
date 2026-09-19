require "test_helper"

class ProcedureControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @procedure = procedures(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create procedure" do
    assert_difference("Procedure.count") do
      post procedures_url, params: { procedure: { title:"test string for title", versionLabel:"test string for versionLabel", Status:Procedure.Statuss[0] } }
    end

    assert_redirected_to procedures_url
  end

 
  
  test "should destroy procedure" do
    assert_difference("Procedure.count", -1) do
      delete procedure_url(@procedure)
    end

    assert_redirected_to procedures_url
  end
  
end


