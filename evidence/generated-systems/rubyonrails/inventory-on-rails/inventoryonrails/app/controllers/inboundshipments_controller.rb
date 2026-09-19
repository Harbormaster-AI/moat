class InboundShipmentsController < ApplicationController
  def index
    @inboundShipments = InboundShipment.all
  end
 
  def show
    @inboundShipment = InboundShipment.find(params[:id])
  end
 
  def new
    @inboundShipment = InboundShipment.new
  end
 
  def edit
    @inboundShipment = InboundShipment.find(params[:id])
  end
 
  def create
    @inboundShipment = InboundShipment.new(inboundShipment_params)
 
    if @inboundShipment.save
      redirect_to inboundShipments_path
    else
      render 'new'
    end
  end
 
  def update
    @inboundShipment = InboundShipment.find(params[:id])
 
    if @inboundShipment.update(inboundShipment_params)
      redirect_to inboundShipments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inboundShipment = InboundShipment.find(params[:id])
    @inboundShipment.destroy
    redirect_to inboundShipments_path
  end

 
  private
    def inboundShipment_params
      params.require(:inboundShipment).permit(:shipmentNumber, :expectedArrivalDate, :arrivalDate, :carrierName, :Status)
    end
end