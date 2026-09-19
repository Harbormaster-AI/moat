class SubrogationRecoverysController < ApplicationController
  def index
    @subrogationRecoverys = SubrogationRecovery.all
  end
 
  def show
    @subrogationRecovery = SubrogationRecovery.find(params[:id])
  end
 
  def new
    @subrogationRecovery = SubrogationRecovery.new
  end
 
  def edit
    @subrogationRecovery = SubrogationRecovery.find(params[:id])
  end
 
  def create
    @subrogationRecovery = SubrogationRecovery.new(subrogationRecovery_params)
 
    if @subrogationRecovery.save
      redirect_to subrogationRecoverys_path
    else
      render 'new'
    end
  end
 
  def update
    @subrogationRecovery = SubrogationRecovery.find(params[:id])
 
    if @subrogationRecovery.update(subrogationRecovery_params)
      redirect_to subrogationRecoverys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @subrogationRecovery = SubrogationRecovery.find(params[:id])
    @subrogationRecovery.destroy
    redirect_to subrogationRecoverys_path
  end

 
  private
    def subrogationRecovery_params
      params.require(:subrogationRecovery).permit(:recoveryReference, :amount, :recoveryDate, :Status)
    end
end