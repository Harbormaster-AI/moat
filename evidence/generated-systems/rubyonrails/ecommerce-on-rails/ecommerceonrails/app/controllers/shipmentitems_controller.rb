class ShipmentItemsController < ApplicationController
  def index
    @shipmentItems = ShipmentItem.all
  end
 
  def show
    @shipmentItem = ShipmentItem.find(params[:id])
  end
 
  def new
    @shipmentItem = ShipmentItem.new
  end
 
  def edit
    @shipmentItem = ShipmentItem.find(params[:id])
  end
 
  def create
    @shipmentItem = ShipmentItem.new(shipmentItem_params)
 
    if @shipmentItem.save
      redirect_to shipmentItems_path
    else
      render 'new'
    end
  end
 
  def update
    @shipmentItem = ShipmentItem.find(params[:id])
 
    if @shipmentItem.update(shipmentItem_params)
      redirect_to shipmentItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @shipmentItem = ShipmentItem.find(params[:id])
    @shipmentItem.destroy
    redirect_to shipmentItems_path
  end

 
  private
    def shipmentItem_params
      params.require(:shipmentItem).permit(:quantity)
    end
end