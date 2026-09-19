class InvestmentAccountsController < ApplicationController
  def index
    @investmentAccounts = InvestmentAccount.all
  end
 
  def show
    @investmentAccount = InvestmentAccount.find(params[:id])
  end
 
  def new
    @investmentAccount = InvestmentAccount.new
  end
 
  def edit
    @investmentAccount = InvestmentAccount.find(params[:id])
  end
 
  def create
    @investmentAccount = InvestmentAccount.new(investmentAccount_params)
 
    if @investmentAccount.save
      redirect_to investmentAccounts_path
    else
      render 'new'
    end
  end
 
  def update
    @investmentAccount = InvestmentAccount.find(params[:id])
 
    if @investmentAccount.update(investmentAccount_params)
      redirect_to investmentAccounts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @investmentAccount = InvestmentAccount.find(params[:id])
    @investmentAccount.destroy
    redirect_to investmentAccounts_path
  end

 
  private
    def investmentAccount_params
      params.require(:investmentAccount).permit(:accountNumber, :baseCurrency, :balance, :AccountType, :Status)
    end
end