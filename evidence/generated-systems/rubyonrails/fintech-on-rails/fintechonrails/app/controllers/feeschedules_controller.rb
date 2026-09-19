class FeeSchedulesController < ApplicationController
  def index
    @feeSchedules = FeeSchedule.all
  end
 
  def show
    @feeSchedule = FeeSchedule.find(params[:id])
  end
 
  def new
    @feeSchedule = FeeSchedule.new
  end
 
  def edit
    @feeSchedule = FeeSchedule.find(params[:id])
  end
 
  def create
    @feeSchedule = FeeSchedule.new(feeSchedule_params)
 
    if @feeSchedule.save
      redirect_to feeSchedules_path
    else
      render 'new'
    end
  end
 
  def update
    @feeSchedule = FeeSchedule.find(params[:id])
 
    if @feeSchedule.update(feeSchedule_params)
      redirect_to feeSchedules_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @feeSchedule = FeeSchedule.find(params[:id])
    @feeSchedule.destroy
    redirect_to feeSchedules_path
  end

 
  private
    def feeSchedule_params
      params.require(:feeSchedule).permit(:name, :amount, :percentage, :minimum, :maximum, :FeeType, :CalculationMethod)
    end
end