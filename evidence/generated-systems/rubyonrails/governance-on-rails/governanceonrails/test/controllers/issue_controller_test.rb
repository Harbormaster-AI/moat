require "test_helper"

class IssueControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @issue = issues(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create issue" do
    assert_difference("Issue.count") do
      post issues_url, params: { issue: { title:"test string for title", openedDate:1.week.ago, closedDate:1.week.ago, IssueType:Issue.IssueTypes[0], Priority:Issue.Prioritys[0], Status:Issue.Statuss[0] } }
    end

    assert_redirected_to issues_url
  end

 
  
  test "should destroy issue" do
    assert_difference("Issue.count", -1) do
      delete issue_url(@issue)
    end

    assert_redirected_to issues_url
  end
  
end


