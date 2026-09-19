class PricingPlansController < ApplicationController
  def index
    @pricingPlans = PricingPlan.all
  end
 
  def show
    @pricingPlan = PricingPlan.find(params[:id])
  end
 
  def new
    @pricingPlan = PricingPlan.new
  end
 
  def edit
    @pricingPlan = PricingPlan.find(params[:id])
  end
 
  def create
    @pricingPlan = PricingPlan.new(pricingPlan_params)
 
    if @pricingPlan.save
      redirect_to pricingPlans_path
    else
      render 'new'
    end
  end
 
  def update
    @pricingPlan = PricingPlan.find(params[:id])
 
    if @pricingPlan.update(pricingPlan_params)
      redirect_to pricingPlans_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @pricingPlan = PricingPlan.find(params[:id])
    @pricingPlan.destroy
    redirect_to pricingPlans_path
  end

 
  private
    def pricingPlan_params
      params.require(:pricingPlan).permit(:name, :planCode, :baseCurrency, :Status)
    end
end