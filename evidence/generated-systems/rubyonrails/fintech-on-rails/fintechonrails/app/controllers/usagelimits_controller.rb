class UsageLimitsController < ApplicationController
  def index
    @usageLimits = UsageLimit.all
  end
 
  def show
    @usageLimit = UsageLimit.find(params[:id])
  end
 
  def new
    @usageLimit = UsageLimit.new
  end
 
  def edit
    @usageLimit = UsageLimit.find(params[:id])
  end
 
  def create
    @usageLimit = UsageLimit.new(usageLimit_params)
 
    if @usageLimit.save
      redirect_to usageLimits_path
    else
      render 'new'
    end
  end
 
  def update
    @usageLimit = UsageLimit.find(params[:id])
 
    if @usageLimit.update(usageLimit_params)
      redirect_to usageLimits_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @usageLimit = UsageLimit.find(params[:id])
    @usageLimit.destroy
    redirect_to usageLimits_path
  end

 
  private
    def usageLimit_params
      params.require(:usageLimit).permit(:name, :amount, :count, :Scope, :Period)
    end
end