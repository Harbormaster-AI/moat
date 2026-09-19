class InvestmentPortfoliosController < ApplicationController
  def index
    @investmentPortfolios = InvestmentPortfolio.all
  end
 
  def show
    @investmentPortfolio = InvestmentPortfolio.find(params[:id])
  end
 
  def new
    @investmentPortfolio = InvestmentPortfolio.new
  end
 
  def edit
    @investmentPortfolio = InvestmentPortfolio.find(params[:id])
  end
 
  def create
    @investmentPortfolio = InvestmentPortfolio.new(investmentPortfolio_params)
 
    if @investmentPortfolio.save
      redirect_to investmentPortfolios_path
    else
      render 'new'
    end
  end
 
  def update
    @investmentPortfolio = InvestmentPortfolio.find(params[:id])
 
    if @investmentPortfolio.update(investmentPortfolio_params)
      redirect_to investmentPortfolios_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @investmentPortfolio = InvestmentPortfolio.find(params[:id])
    @investmentPortfolio.destroy
    redirect_to investmentPortfolios_path
  end

 
  private
    def investmentPortfolio_params
      params.require(:investmentPortfolio).permit(:portfolioCode, :baseCurrency, :createdAt, :Status)
    end
end