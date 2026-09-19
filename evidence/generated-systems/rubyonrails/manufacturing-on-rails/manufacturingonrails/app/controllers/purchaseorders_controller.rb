class PurchaseOrdersController < ApplicationController
  def index
    @purchaseOrders = PurchaseOrder.all
  end
 
  def show
    @purchaseOrder = PurchaseOrder.find(params[:id])
  end
 
  def new
    @purchaseOrder = PurchaseOrder.new
  end
 
  def edit
    @purchaseOrder = PurchaseOrder.find(params[:id])
  end
 
  def create
    @purchaseOrder = PurchaseOrder.new(purchaseOrder_params)
 
    if @purchaseOrder.save
      redirect_to purchaseOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @purchaseOrder = PurchaseOrder.find(params[:id])
 
    if @purchaseOrder.update(purchaseOrder_params)
      redirect_to purchaseOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @purchaseOrder = PurchaseOrder.find(params[:id])
    @purchaseOrder.destroy
    redirect_to purchaseOrders_path
  end

 
  private
    def purchaseOrder_params
      params.require(:purchaseOrder).permit(:poNumber, :orderDate, :totalAmount, :Status)
    end
end