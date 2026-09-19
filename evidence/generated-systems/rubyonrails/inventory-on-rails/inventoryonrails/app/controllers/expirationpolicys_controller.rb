class ExpirationPolicysController < ApplicationController
  def index
    @expirationPolicys = ExpirationPolicy.all
  end
 
  def show
    @expirationPolicy = ExpirationPolicy.find(params[:id])
  end
 
  def new
    @expirationPolicy = ExpirationPolicy.new
  end
 
  def edit
    @expirationPolicy = ExpirationPolicy.find(params[:id])
  end
 
  def create
    @expirationPolicy = ExpirationPolicy.new(expirationPolicy_params)
 
    if @expirationPolicy.save
      redirect_to expirationPolicys_path
    else
      render 'new'
    end
  end
 
  def update
    @expirationPolicy = ExpirationPolicy.find(params[:id])
 
    if @expirationPolicy.update(expirationPolicy_params)
      redirect_to expirationPolicys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @expirationPolicy = ExpirationPolicy.find(params[:id])
    @expirationPolicy.destroy
    redirect_to expirationPolicys_path
  end

 
  private
    def expirationPolicy_params
      params.require(:expirationPolicy).permit(:rejectIfDaysToExpireLessThan, :autoQuarantineDaysToExpire, :RotationMethod)
    end
end