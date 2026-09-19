class TransferOrdersController < ApplicationController
  def index
    @transferOrders = TransferOrder.all
  end
 
  def show
    @transferOrder = TransferOrder.find(params[:id])
  end
 
  def new
    @transferOrder = TransferOrder.new
  end
 
  def edit
    @transferOrder = TransferOrder.find(params[:id])
  end
 
  def create
    @transferOrder = TransferOrder.new(transferOrder_params)
 
    if @transferOrder.save
      redirect_to transferOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @transferOrder = TransferOrder.find(params[:id])
 
    if @transferOrder.update(transferOrder_params)
      redirect_to transferOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @transferOrder = TransferOrder.find(params[:id])
    @transferOrder.destroy
    redirect_to transferOrders_path
  end

 
  private
    def transferOrder_params
      params.require(:transferOrder).permit(:orderNumber, :requestedShipDate, :requestedReceiveDate, :shippedDate, :receivedDate, :Status)
    end
end