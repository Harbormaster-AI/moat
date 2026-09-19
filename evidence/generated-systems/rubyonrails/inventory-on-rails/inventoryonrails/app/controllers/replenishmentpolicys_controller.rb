class ReplenishmentPolicysController < ApplicationController
  def index
    @replenishmentPolicys = ReplenishmentPolicy.all
  end
 
  def show
    @replenishmentPolicy = ReplenishmentPolicy.find(params[:id])
  end
 
  def new
    @replenishmentPolicy = ReplenishmentPolicy.new
  end
 
  def edit
    @replenishmentPolicy = ReplenishmentPolicy.find(params[:id])
  end
 
  def create
    @replenishmentPolicy = ReplenishmentPolicy.new(replenishmentPolicy_params)
 
    if @replenishmentPolicy.save
      redirect_to replenishmentPolicys_path
    else
      render 'new'
    end
  end
 
  def update
    @replenishmentPolicy = ReplenishmentPolicy.find(params[:id])
 
    if @replenishmentPolicy.update(replenishmentPolicy_params)
      redirect_to replenishmentPolicys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @replenishmentPolicy = ReplenishmentPolicy.find(params[:id])
    @replenishmentPolicy.destroy
    redirect_to replenishmentPolicys_path
  end

 
  private
    def replenishmentPolicy_params
      params.require(:replenishmentPolicy).permit(:minLevel, :maxLevel, :reorderPoint, :reorderQuantity, :leadTimeDays, :reviewPeriodDays, :PolicyType)
    end
end