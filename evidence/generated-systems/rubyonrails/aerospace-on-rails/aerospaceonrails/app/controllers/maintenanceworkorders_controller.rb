class MaintenanceWorkOrdersController < ApplicationController
  def index
    @maintenanceWorkOrders = MaintenanceWorkOrder.all
  end
 
  def show
    @maintenanceWorkOrder = MaintenanceWorkOrder.find(params[:id])
  end
 
  def new
    @maintenanceWorkOrder = MaintenanceWorkOrder.new
  end
 
  def edit
    @maintenanceWorkOrder = MaintenanceWorkOrder.find(params[:id])
  end
 
  def create
    @maintenanceWorkOrder = MaintenanceWorkOrder.new(maintenanceWorkOrder_params)
 
    if @maintenanceWorkOrder.save
      redirect_to maintenanceWorkOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @maintenanceWorkOrder = MaintenanceWorkOrder.find(params[:id])
 
    if @maintenanceWorkOrder.update(maintenanceWorkOrder_params)
      redirect_to maintenanceWorkOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @maintenanceWorkOrder = MaintenanceWorkOrder.find(params[:id])
    @maintenanceWorkOrder.destroy
    redirect_to maintenanceWorkOrders_path
  end

 
  private
    def maintenanceWorkOrder_params
      params.require(:maintenanceWorkOrder).permit(:workOrderNumber, :Status)
    end
end