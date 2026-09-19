class LoanTransactionsController < ApplicationController
  def index
    @loanTransactions = LoanTransaction.all
  end
 
  def show
    @loanTransaction = LoanTransaction.find(params[:id])
  end
 
  def new
    @loanTransaction = LoanTransaction.new
  end
 
  def edit
    @loanTransaction = LoanTransaction.find(params[:id])
  end
 
  def create
    @loanTransaction = LoanTransaction.new(loanTransaction_params)
 
    if @loanTransaction.save
      redirect_to loanTransactions_path
    else
      render 'new'
    end
  end
 
  def update
    @loanTransaction = LoanTransaction.find(params[:id])
 
    if @loanTransaction.update(loanTransaction_params)
      redirect_to loanTransactions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @loanTransaction = LoanTransaction.find(params[:id])
    @loanTransaction.destroy
    redirect_to loanTransactions_path
  end

 
  private
    def loanTransaction_params
      params.require(:loanTransaction).permit(:transactionId, :amount, :postingDate, :Type, :Status)
    end
end