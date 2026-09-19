require "test_helper"

class LineageNodeControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @lineageNode = lineageNodes(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create lineageNode" do
    assert_difference("LineageNode.count") do
      post lineageNodes_url, params: { lineageNode: { name:"test string for name", qualifiedName:"test string for qualifiedName", NodeType:LineageNode.NodeTypes[0] } }
    end

    assert_redirected_to lineageNodes_url
  end

 
  
  test "should destroy lineageNode" do
    assert_difference("LineageNode.count", -1) do
      delete lineageNode_url(@lineageNode)
    end

    assert_redirected_to lineageNodes_url
  end
  
end


