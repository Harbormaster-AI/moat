class PolicysController < ApplicationController
  def index
    @policys = Policy.all
  end
 
  def show
    @policy = Policy.find(params[:id])
  end
 
  def new
    @policy = Policy.new
  end
 
  def edit
    @policy = Policy.find(params[:id])
  end
 
  def create
    @policy = Policy.new(policy_params)
 
    if @policy.save
      redirect_to policys_path
    else
      render 'new'
    end
  end
 
  def update
    @policy = Policy.find(params[:id])
 
    if @policy.update(policy_params)
      redirect_to policys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @policy = Policy.find(params[:id])
    @policy.destroy
    redirect_to policys_path
  end

 
  private
    def policy_params
      params.require(:policy).permit(:policyNumber, :name, :effectiveDate, :description)
    end
end