class ContractsController < ApplicationController
  def index
    @contracts = Contract.all
  end
 
  def show
    @contract = Contract.find(params[:id])
  end
 
  def new
    @contract = Contract.new
  end
 
  def edit
    @contract = Contract.find(params[:id])
  end
 
  def create
    @contract = Contract.new(contract_params)
 
    if @contract.save
      redirect_to contracts_path
    else
      render 'new'
    end
  end
 
  def update
    @contract = Contract.find(params[:id])
 
    if @contract.update(contract_params)
      redirect_to contracts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @contract = Contract.find(params[:id])
    @contract.destroy
    redirect_to contracts_path
  end

 
  private
    def contract_params
      params.require(:contract).permit(:contractNumber, :startDate, :endDate, :renewalTermMonths, :autoRenew, :Status)
    end
end