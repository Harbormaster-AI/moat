class ProductionOrdersController < ApplicationController
  def index
    @productionOrders = ProductionOrder.all
  end
 
  def show
    @productionOrder = ProductionOrder.find(params[:id])
  end
 
  def new
    @productionOrder = ProductionOrder.new
  end
 
  def edit
    @productionOrder = ProductionOrder.find(params[:id])
  end
 
  def create
    @productionOrder = ProductionOrder.new(productionOrder_params)
 
    if @productionOrder.save
      redirect_to productionOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @productionOrder = ProductionOrder.find(params[:id])
 
    if @productionOrder.update(productionOrder_params)
      redirect_to productionOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @productionOrder = ProductionOrder.find(params[:id])
    @productionOrder.destroy
    redirect_to productionOrders_path
  end

 
  private
    def productionOrder_params
      params.require(:productionOrder).permit(:orderNumber, :Status)
    end
end