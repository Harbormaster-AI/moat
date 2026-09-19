class TransferOrderLinesController < ApplicationController
  def index
    @transferOrderLines = TransferOrderLine.all
  end
 
  def show
    @transferOrderLine = TransferOrderLine.find(params[:id])
  end
 
  def new
    @transferOrderLine = TransferOrderLine.new
  end
 
  def edit
    @transferOrderLine = TransferOrderLine.find(params[:id])
  end
 
  def create
    @transferOrderLine = TransferOrderLine.new(transferOrderLine_params)
 
    if @transferOrderLine.save
      redirect_to transferOrderLines_path
    else
      render 'new'
    end
  end
 
  def update
    @transferOrderLine = TransferOrderLine.find(params[:id])
 
    if @transferOrderLine.update(transferOrderLine_params)
      redirect_to transferOrderLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @transferOrderLine = TransferOrderLine.find(params[:id])
    @transferOrderLine.destroy
    redirect_to transferOrderLines_path
  end

 
  private
    def transferOrderLine_params
      params.require(:transferOrderLine).permit(:lineNumber, :quantity, :UnitOfMeasure, :StockStatus)
    end
end