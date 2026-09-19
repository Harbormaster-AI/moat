class InboundShipmentLinesController < ApplicationController
  def index
    @inboundShipmentLines = InboundShipmentLine.all
  end
 
  def show
    @inboundShipmentLine = InboundShipmentLine.find(params[:id])
  end
 
  def new
    @inboundShipmentLine = InboundShipmentLine.new
  end
 
  def edit
    @inboundShipmentLine = InboundShipmentLine.find(params[:id])
  end
 
  def create
    @inboundShipmentLine = InboundShipmentLine.new(inboundShipmentLine_params)
 
    if @inboundShipmentLine.save
      redirect_to inboundShipmentLines_path
    else
      render 'new'
    end
  end
 
  def update
    @inboundShipmentLine = InboundShipmentLine.find(params[:id])
 
    if @inboundShipmentLine.update(inboundShipmentLine_params)
      redirect_to inboundShipmentLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inboundShipmentLine = InboundShipmentLine.find(params[:id])
    @inboundShipmentLine.destroy
    redirect_to inboundShipmentLines_path
  end

 
  private
    def inboundShipmentLine_params
      params.require(:inboundShipmentLine).permit(:lineNumber, :quantity, :UnitOfMeasure, :StockStatus)
    end
end