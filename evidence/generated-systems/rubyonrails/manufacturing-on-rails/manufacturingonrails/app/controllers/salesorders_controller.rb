class SalesOrdersController < ApplicationController
  def index
    @salesOrders = SalesOrder.all
  end
 
  def show
    @salesOrder = SalesOrder.find(params[:id])
  end
 
  def new
    @salesOrder = SalesOrder.new
  end
 
  def edit
    @salesOrder = SalesOrder.find(params[:id])
  end
 
  def create
    @salesOrder = SalesOrder.new(salesOrder_params)
 
    if @salesOrder.save
      redirect_to salesOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @salesOrder = SalesOrder.find(params[:id])
 
    if @salesOrder.update(salesOrder_params)
      redirect_to salesOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @salesOrder = SalesOrder.find(params[:id])
    @salesOrder.destroy
    redirect_to salesOrders_path
  end

 
  private
    def salesOrder_params
      params.require(:salesOrder).permit(:orderNumber, :orderDate, :totalAmount, :Status)
    end
end