class PlannedOrdersController < ApplicationController
  def index
    @plannedOrders = PlannedOrder.all
  end
 
  def show
    @plannedOrder = PlannedOrder.find(params[:id])
  end
 
  def new
    @plannedOrder = PlannedOrder.new
  end
 
  def edit
    @plannedOrder = PlannedOrder.find(params[:id])
  end
 
  def create
    @plannedOrder = PlannedOrder.new(plannedOrder_params)
 
    if @plannedOrder.save
      redirect_to plannedOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @plannedOrder = PlannedOrder.find(params[:id])
 
    if @plannedOrder.update(plannedOrder_params)
      redirect_to plannedOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @plannedOrder = PlannedOrder.find(params[:id])
    @plannedOrder.destroy
    redirect_to plannedOrders_path
  end

 
  private
    def plannedOrder_params
      params.require(:plannedOrder).permit(:plannedOrderNumber, :quantity, :dueDate, :OrderType, :Status)
    end
end