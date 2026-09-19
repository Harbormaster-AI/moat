class PurchaseOrderLinesController < ApplicationController
  def index
    @purchaseOrderLines = PurchaseOrderLine.all
  end
 
  def show
    @purchaseOrderLine = PurchaseOrderLine.find(params[:id])
  end
 
  def new
    @purchaseOrderLine = PurchaseOrderLine.new
  end
 
  def edit
    @purchaseOrderLine = PurchaseOrderLine.find(params[:id])
  end
 
  def create
    @purchaseOrderLine = PurchaseOrderLine.new(purchaseOrderLine_params)
 
    if @purchaseOrderLine.save
      redirect_to purchaseOrderLines_path
    else
      render 'new'
    end
  end
 
  def update
    @purchaseOrderLine = PurchaseOrderLine.find(params[:id])
 
    if @purchaseOrderLine.update(purchaseOrderLine_params)
      redirect_to purchaseOrderLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @purchaseOrderLine = PurchaseOrderLine.find(params[:id])
    @purchaseOrderLine.destroy
    redirect_to purchaseOrderLines_path
  end

 
  private
    def purchaseOrderLine_params
      params.require(:purchaseOrderLine).permit(:lineNumber, :quantity, :unitPrice, :dueDate)
    end
end