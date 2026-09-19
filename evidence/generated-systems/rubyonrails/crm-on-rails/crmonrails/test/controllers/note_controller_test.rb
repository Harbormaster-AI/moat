require "test_helper"

class NoteControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @note = notes(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create note" do
    assert_difference("Note.count") do
      post notes_url, params: { note: { title:"test string for title", content:"test string for content", createdAt:1.week.ago, updatedAt:1.week.ago } }
    end

    assert_redirected_to notes_url
  end

 
  
  test "should destroy note" do
    assert_difference("Note.count", -1) do
      delete note_url(@note)
    end

    assert_redirected_to notes_url
  end
  
end


