class BankAccountsController < ApplicationController
  def index
    @bankAccounts = BankAccount.all
  end
 
  def show
    @bankAccount = BankAccount.find(params[:id])
  end
 
  def new
    @bankAccount = BankAccount.new
  end
 
  def edit
    @bankAccount = BankAccount.find(params[:id])
  end
 
  def create
    @bankAccount = BankAccount.new(bankAccount_params)
 
    if @bankAccount.save
      redirect_to bankAccounts_path
    else
      render 'new'
    end
  end
 
  def update
    @bankAccount = BankAccount.find(params[:id])
 
    if @bankAccount.update(bankAccount_params)
      redirect_to bankAccounts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @bankAccount = BankAccount.find(params[:id])
    @bankAccount.destroy
    redirect_to bankAccounts_path
  end

 
  private
    def bankAccount_params
      params.require(:bankAccount).permit(:accountHolder, :bankName, :iban, :bic, :accountNumber, :routingNumber)
    end
end