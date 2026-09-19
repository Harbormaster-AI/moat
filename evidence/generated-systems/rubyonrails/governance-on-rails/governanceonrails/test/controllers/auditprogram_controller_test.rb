require "test_helper"

class AuditProgramControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @auditProgram = auditPrograms(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create auditProgram" do
    assert_difference("AuditProgram.count") do
      post auditPrograms_url, params: { auditProgram: { name:"test string for name", scope:"test string for scope", Cycle:AuditProgram.Cycles[0], Status:AuditProgram.Statuss[0] } }
    end

    assert_redirected_to auditPrograms_url
  end

 
  
  test "should destroy auditProgram" do
    assert_difference("AuditProgram.count", -1) do
      delete auditProgram_url(@auditProgram)
    end

    assert_redirected_to auditPrograms_url
  end
  
end


