require "test_helper"

class ComplianceProgramControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @complianceProgram = compliancePrograms(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create complianceProgram" do
    assert_difference("ComplianceProgram.count") do
      post compliancePrograms_url, params: { complianceProgram: { name:"test string for name", framework:"test string for framework", Status:ComplianceProgram.Statuss[0] } }
    end

    assert_redirected_to compliancePrograms_url
  end

 
  
  test "should destroy complianceProgram" do
    assert_difference("ComplianceProgram.count", -1) do
      delete complianceProgram_url(@complianceProgram)
    end

    assert_redirected_to compliancePrograms_url
  end
  
end


