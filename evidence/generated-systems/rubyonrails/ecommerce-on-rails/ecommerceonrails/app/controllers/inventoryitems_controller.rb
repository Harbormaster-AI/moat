class InventoryItemsController < ApplicationController
  def index
    @inventoryItems = InventoryItem.all
  end
 
  def show
    @inventoryItem = InventoryItem.find(params[:id])
  end
 
  def new
    @inventoryItem = InventoryItem.new
  end
 
  def edit
    @inventoryItem = InventoryItem.find(params[:id])
  end
 
  def create
    @inventoryItem = InventoryItem.new(inventoryItem_params)
 
    if @inventoryItem.save
      redirect_to inventoryItems_path
    else
      render 'new'
    end
  end
 
  def update
    @inventoryItem = InventoryItem.find(params[:id])
 
    if @inventoryItem.update(inventoryItem_params)
      redirect_to inventoryItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inventoryItem = InventoryItem.find(params[:id])
    @inventoryItem.destroy
    redirect_to inventoryItems_path
  end

 
  private
    def inventoryItem_params
      params.require(:inventoryItem).permit(:quantityOnHand, :quantityReserved, :safetyStock, :Status)
    end
end