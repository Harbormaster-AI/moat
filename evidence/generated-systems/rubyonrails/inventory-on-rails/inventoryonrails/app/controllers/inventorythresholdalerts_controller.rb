class InventoryThresholdAlertsController < ApplicationController
  def index
    @inventoryThresholdAlerts = InventoryThresholdAlert.all
  end
 
  def show
    @inventoryThresholdAlert = InventoryThresholdAlert.find(params[:id])
  end
 
  def new
    @inventoryThresholdAlert = InventoryThresholdAlert.new
  end
 
  def edit
    @inventoryThresholdAlert = InventoryThresholdAlert.find(params[:id])
  end
 
  def create
    @inventoryThresholdAlert = InventoryThresholdAlert.new(inventoryThresholdAlert_params)
 
    if @inventoryThresholdAlert.save
      redirect_to inventoryThresholdAlerts_path
    else
      render 'new'
    end
  end
 
  def update
    @inventoryThresholdAlert = InventoryThresholdAlert.find(params[:id])
 
    if @inventoryThresholdAlert.update(inventoryThresholdAlert_params)
      redirect_to inventoryThresholdAlerts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inventoryThresholdAlert = InventoryThresholdAlert.find(params[:id])
    @inventoryThresholdAlert.destroy
    redirect_to inventoryThresholdAlerts_path
  end

 
  private
    def inventoryThresholdAlert_params
      params.require(:inventoryThresholdAlert).permit(:alertNumber, :detectedAt, :message, :AlertType, :Status)
    end
end