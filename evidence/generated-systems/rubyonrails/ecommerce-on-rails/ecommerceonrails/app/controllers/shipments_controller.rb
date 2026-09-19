class ShipmentsController < ApplicationController
  def index
    @shipments = Shipment.all
  end
 
  def show
    @shipment = Shipment.find(params[:id])
  end
 
  def new
    @shipment = Shipment.new
  end
 
  def edit
    @shipment = Shipment.find(params[:id])
  end
 
  def create
    @shipment = Shipment.new(shipment_params)
 
    if @shipment.save
      redirect_to shipments_path
    else
      render 'new'
    end
  end
 
  def update
    @shipment = Shipment.find(params[:id])
 
    if @shipment.update(shipment_params)
      redirect_to shipments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @shipment = Shipment.find(params[:id])
    @shipment.destroy
    redirect_to shipments_path
  end

 
  private
    def shipment_params
      params.require(:shipment).permit(:shipmentNumber, :shippedDate, :deliveredDate, :trackingNumber, :shippingAddress, :Status, :Carrier)
    end
end