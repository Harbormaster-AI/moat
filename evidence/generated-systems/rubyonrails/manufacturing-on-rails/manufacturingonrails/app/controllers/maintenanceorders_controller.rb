class MaintenanceOrdersController < ApplicationController
  def index
    @maintenanceOrders = MaintenanceOrder.all
  end
 
  def show
    @maintenanceOrder = MaintenanceOrder.find(params[:id])
  end
 
  def new
    @maintenanceOrder = MaintenanceOrder.new
  end
 
  def edit
    @maintenanceOrder = MaintenanceOrder.find(params[:id])
  end
 
  def create
    @maintenanceOrder = MaintenanceOrder.new(maintenanceOrder_params)
 
    if @maintenanceOrder.save
      redirect_to maintenanceOrders_path
    else
      render 'new'
    end
  end
 
  def update
    @maintenanceOrder = MaintenanceOrder.find(params[:id])
 
    if @maintenanceOrder.update(maintenanceOrder_params)
      redirect_to maintenanceOrders_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @maintenanceOrder = MaintenanceOrder.find(params[:id])
    @maintenanceOrder.destroy
    redirect_to maintenanceOrders_path
  end

 
  private
    def maintenanceOrder_params
      params.require(:maintenanceOrder).permit(:orderNumber, :priority, :requestedDate, :completionDate, :Status)
    end
end