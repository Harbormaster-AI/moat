class WorkOrdersController < ApplicationController
  def index
    @workOrders = WorkOrder.all
  end
 
  def show
    @workOrder = WorkOrder.find(params[:id])
  end
 
  def new
    @workOrder = WorkOrder.new
  end
 
  def edit
    @workOrder = WorkOrder.find(params[:id])
  end
 
  def create
    @workOrder = WorkOrder.new(workOrder_params)
 
    if @workOrder.save
      redirect_to workOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @workOrder = WorkOrder.find(params[:id])
 
    if @workOrder.update(workOrder_params)
      redirect_to workOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @workOrder = WorkOrder.find(params[:id])
    @workOrder.destroy
    redirect_to workOrders_path
  end

 
  private
    def workOrder_params
      params.require(:workOrder).permit(:workOrderNumber, :plannedStart, :plannedEnd, :quantity, :priority, :Status)
    end
end