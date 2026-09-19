class BenefitPlansController < ApplicationController
  def index
    @benefitPlans = BenefitPlan.all
  end
 
  def show
    @benefitPlan = BenefitPlan.find(params[:id])
  end
 
  def new
    @benefitPlan = BenefitPlan.new
  end
 
  def edit
    @benefitPlan = BenefitPlan.find(params[:id])
  end
 
  def create
    @benefitPlan = BenefitPlan.new(benefitPlan_params)
 
    if @benefitPlan.save
      redirect_to benefitPlans_path
    else
      render 'new'
    end
  end
 
  def update
    @benefitPlan = BenefitPlan.find(params[:id])
 
    if @benefitPlan.update(benefitPlan_params)
      redirect_to benefitPlans_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @benefitPlan = BenefitPlan.find(params[:id])
    @benefitPlan.destroy
    redirect_to benefitPlans_path
  end

 
  private
    def benefitPlan_params
      params.require(:benefitPlan).permit(:name, :providerName, :employeeContributionRate, :employerContributionRate, :eligibilityRules, :BenefitType)
    end
end