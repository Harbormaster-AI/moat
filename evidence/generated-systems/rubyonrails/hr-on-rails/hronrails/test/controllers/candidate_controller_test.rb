require "test_helper"

class CandidateControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @candidate = candidates(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create candidate" do
    assert_difference("Candidate.count") do
      post candidates_url, params: { candidate: { name:"test value", email:"test value", phone:"test value", Source:Candidate.Sources[0] } }
    end

    assert_redirected_to candidates_url
  end

 
  
  test "should destroy candidate" do
    assert_difference("Candidate.count", -1) do
      delete candidate_url(@candidate)
    end

    assert_redirected_to candidates_url
  end
  
end


