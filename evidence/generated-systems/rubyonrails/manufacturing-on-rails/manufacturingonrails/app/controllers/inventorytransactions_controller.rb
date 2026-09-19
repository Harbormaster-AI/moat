class InventoryTransactionsController < ApplicationController
  def index
    @inventoryTransactions = InventoryTransaction.all
  end
 
  def show
    @inventoryTransaction = InventoryTransaction.find(params[:id])
  end
 
  def new
    @inventoryTransaction = InventoryTransaction.new
  end
 
  def edit
    @inventoryTransaction = InventoryTransaction.find(params[:id])
  end
 
  def create
    @inventoryTransaction = InventoryTransaction.new(inventoryTransaction_params)
 
    if @inventoryTransaction.save
      redirect_to inventoryTransactions_path
    else
      render 'new'
    end
  end
 
  def update
    @inventoryTransaction = InventoryTransaction.find(params[:id])
 
    if @inventoryTransaction.update(inventoryTransaction_params)
      redirect_to inventoryTransactions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inventoryTransaction = InventoryTransaction.find(params[:id])
    @inventoryTransaction.destroy
    redirect_to inventoryTransactions_path
  end

 
  private
    def inventoryTransaction_params
      params.require(:inventoryTransaction).permit(:transactionNumber, :quantity, :transactionDateTime, :referenceDocument, :TransactionType)
    end
end