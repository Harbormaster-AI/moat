class OpportunitysController < ApplicationController
  def index
    @opportunitys = Opportunity.all
  end
 
  def show
    @opportunity = Opportunity.find(params[:id])
  end
 
  def new
    @opportunity = Opportunity.new
  end
 
  def edit
    @opportunity = Opportunity.find(params[:id])
  end
 
  def create
    @opportunity = Opportunity.new(opportunity_params)
 
    if @opportunity.save
      redirect_to opportunitys_path
    else
      render 'new'
    end
  end
 
  def update
    @opportunity = Opportunity.find(params[:id])
 
    if @opportunity.update(opportunity_params)
      redirect_to opportunitys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @opportunity = Opportunity.find(params[:id])
    @opportunity.destroy
    redirect_to opportunitys_path
  end

 
  private
    def opportunity_params
      params.require(:opportunity).permit(:name, :amount, :closeDate, :probability, :description, :Stage, :Type, :ForecastCategory)
    end
end