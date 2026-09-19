class LoansController < ApplicationController
  def index
    @loans = Loan.all
  end
 
  def show
    @loan = Loan.find(params[:id])
  end
 
  def new
    @loan = Loan.new
  end
 
  def edit
    @loan = Loan.find(params[:id])
  end
 
  def create
    @loan = Loan.new(loan_params)
 
    if @loan.save
      redirect_to loans_path
    else
      render 'new'
    end
  end
 
  def update
    @loan = Loan.find(params[:id])
 
    if @loan.update(loan_params)
      redirect_to loans_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @loan = Loan.find(params[:id])
    @loan.destroy
    redirect_to loans_path
  end

 
  private
    def loan_params
      params.require(:loan).permit(:loanNumber, :principal, :interestRate, :originationDate, :maturityDate, :RateType, :Status)
    end
end