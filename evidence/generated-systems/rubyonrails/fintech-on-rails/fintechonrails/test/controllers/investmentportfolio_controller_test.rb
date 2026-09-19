require "test_helper"

class InvestmentPortfolioControllerTest < ActionDispatch::IntegrationTest
  # called before every single test
  setup do
    @investmentPortfolio = investmentPortfolios(:one)
  end

  # called after every single test
  teardown do
    # when controller is using cache it may be a good idea to reset it afterwards
    Rails.cache.clear
  end

  test "should create investmentPortfolio" do
    assert_difference("InvestmentPortfolio.count") do
      post investmentPortfolios_url, params: { investmentPortfolio: { portfolioCode:"test string for portfolioCode", baseCurrency:"test string for baseCurrency", createdAt:1.week.ago, Status:InvestmentPortfolio.Statuss[0] } }
    end

    assert_redirected_to investmentPortfolios_url
  end

 
  
  test "should destroy investmentPortfolio" do
    assert_difference("InvestmentPortfolio.count", -1) do
      delete investmentPortfolio_url(@investmentPortfolio)
    end

    assert_redirected_to investmentPortfolios_url
  end
  
end


