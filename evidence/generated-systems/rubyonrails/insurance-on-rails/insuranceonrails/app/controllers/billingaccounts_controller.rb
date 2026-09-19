class BillingAccountsController < ApplicationController
  def index
    @billingAccounts = BillingAccount.all
  end
 
  def show
    @billingAccount = BillingAccount.find(params[:id])
  end
 
  def new
    @billingAccount = BillingAccount.new
  end
 
  def edit
    @billingAccount = BillingAccount.find(params[:id])
  end
 
  def create
    @billingAccount = BillingAccount.new(billingAccount_params)
 
    if @billingAccount.save
      redirect_to billingAccounts_path
    else
      render 'new'
    end
  end
 
  def update
    @billingAccount = BillingAccount.find(params[:id])
 
    if @billingAccount.update(billingAccount_params)
      redirect_to billingAccounts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @billingAccount = BillingAccount.find(params[:id])
    @billingAccount.destroy
    redirect_to billingAccounts_path
  end

 
  private
    def billingAccount_params
      params.require(:billingAccount).permit(:accountNumber, :balance, :Status)
    end
end