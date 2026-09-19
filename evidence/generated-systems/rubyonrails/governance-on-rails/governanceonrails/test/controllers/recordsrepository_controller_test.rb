require "test_helper"

class RecordsRepositoryControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @recordsRepository = recordsRepositorys(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create recordsRepository" do
    assert_difference("RecordsRepository.count") do
      post recordsRepositorys_url, params: { recordsRepository: { name:"test string for name", location:"test string for location", ownerDepartment:"test string for ownerDepartment", RepositoryType:RecordsRepository.RepositoryTypes[0] } }
    end

    assert_redirected_to recordsRepositorys_url
  end

 
  
  test "should destroy recordsRepository" do
    assert_difference("RecordsRepository.count", -1) do
      delete recordsRepository_url(@recordsRepository)
    end

    assert_redirected_to recordsRepositorys_url
  end
  
end


