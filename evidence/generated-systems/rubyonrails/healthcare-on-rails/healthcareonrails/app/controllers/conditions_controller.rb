class ConditionsController < ApplicationController
  def index
    @conditions = Condition.all
  end
 
  def show
    @condition = Condition.find(params[:id])
  end
 
  def new
    @condition = Condition.new
  end
 
  def edit
    @condition = Condition.find(params[:id])
  end
 
  def create
    @condition = Condition.new(condition_params)
 
    if @condition.save
      redirect_to conditions_path
    else
      render 'new'
    end
  end
 
  def update
    @condition = Condition.find(params[:id])
 
    if @condition.update(condition_params)
      redirect_to conditions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @condition = Condition.find(params[:id])
    @condition.destroy
    redirect_to conditions_path
  end

 
  private
    def condition_params
      params.require(:condition).permit(:code, :onsetDate, :abatementDate, :ClinicalStatus, :VerificationStatus)
    end
end