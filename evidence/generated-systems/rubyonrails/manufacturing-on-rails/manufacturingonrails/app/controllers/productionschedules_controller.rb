class ProductionSchedulesController < ApplicationController
  def index
    @productionSchedules = ProductionSchedule.all
  end
 
  def show
    @productionSchedule = ProductionSchedule.find(params[:id])
  end
 
  def new
    @productionSchedule = ProductionSchedule.new
  end
 
  def edit
    @productionSchedule = ProductionSchedule.find(params[:id])
  end
 
  def create
    @productionSchedule = ProductionSchedule.new(productionSchedule_params)
 
    if @productionSchedule.save
      redirect_to productionSchedules_path
    else
      render 'new'
    end
  end
 
  def update
    @productionSchedule = ProductionSchedule.find(params[:id])
 
    if @productionSchedule.update(productionSchedule_params)
      redirect_to productionSchedules_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @productionSchedule = ProductionSchedule.find(params[:id])
    @productionSchedule.destroy
    redirect_to productionSchedules_path
  end

 
  private
    def productionSchedule_params
      params.require(:productionSchedule).permit(:scheduleNumber, :horizonStart, :horizonEnd, :Status)
    end
end