class WorkSchedulesController < ApplicationController
  def index
    @workSchedules = WorkSchedule.all
  end
 
  def show
    @workSchedule = WorkSchedule.find(params[:id])
  end
 
  def new
    @workSchedule = WorkSchedule.new
  end
 
  def edit
    @workSchedule = WorkSchedule.find(params[:id])
  end
 
  def create
    @workSchedule = WorkSchedule.new(workSchedule_params)
 
    if @workSchedule.save
      redirect_to workSchedules_path
    else
      render 'new'
    end
  end
 
  def update
    @workSchedule = WorkSchedule.find(params[:id])
 
    if @workSchedule.update(workSchedule_params)
      redirect_to workSchedules_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @workSchedule = WorkSchedule.find(params[:id])
    @workSchedule.destroy
    redirect_to workSchedules_path
  end

 
  private
    def workSchedule_params
      params.require(:workSchedule).permit(:name, :standardHoursPerWeek, :ScheduleType)
    end
end