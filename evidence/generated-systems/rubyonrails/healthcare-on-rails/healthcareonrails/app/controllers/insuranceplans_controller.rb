class InsurancePlansController < ApplicationController
  def index
    @insurancePlans = InsurancePlan.all
  end
 
  def show
    @insurancePlan = InsurancePlan.find(params[:id])
  end
 
  def new
    @insurancePlan = InsurancePlan.new
  end
 
  def edit
    @insurancePlan = InsurancePlan.find(params[:id])
  end
 
  def create
    @insurancePlan = InsurancePlan.new(insurancePlan_params)
 
    if @insurancePlan.save
      redirect_to insurancePlans_path
    else
      render 'new'
    end
  end
 
  def update
    @insurancePlan = InsurancePlan.find(params[:id])
 
    if @insurancePlan.update(insurancePlan_params)
      redirect_to insurancePlans_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @insurancePlan = InsurancePlan.find(params[:id])
    @insurancePlan.destroy
    redirect_to insurancePlans_path
  end

 
  private
    def insurancePlan_params
      params.require(:insurancePlan).permit(:name, :planCode, :PlanType)
    end
end