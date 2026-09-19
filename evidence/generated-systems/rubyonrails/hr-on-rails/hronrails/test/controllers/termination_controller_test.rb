require "test_helper"

class TerminationControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @termination = terminations(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create termination" do
    assert_difference("Termination.count") do
      post terminations_url, params: { termination: { terminationNumber:"test string for terminationNumber", terminationDate:1.week.ago, notes:"test string for notes", eligibleForRehire:true, Reason:Termination.Reasons[0], Type:Termination.Types[0] } }
    end

    assert_redirected_to terminations_url
  end

 
  
  test "should destroy termination" do
    assert_difference("Termination.count", -1) do
      delete termination_url(@termination)
    end

    assert_redirected_to terminations_url
  end
  
end


