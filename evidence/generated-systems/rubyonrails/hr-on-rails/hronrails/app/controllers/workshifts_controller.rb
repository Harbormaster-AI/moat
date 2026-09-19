class WorkShiftsController < ApplicationController
  def index
    @workShifts = WorkShift.all
  end
 
  def show
    @workShift = WorkShift.find(params[:id])
  end
 
  def new
    @workShift = WorkShift.new
  end
 
  def edit
    @workShift = WorkShift.find(params[:id])
  end
 
  def create
    @workShift = WorkShift.new(workShift_params)
 
    if @workShift.save
      redirect_to workShifts_path
    else
      render 'new'
    end
  end
 
  def update
    @workShift = WorkShift.find(params[:id])
 
    if @workShift.update(workShift_params)
      redirect_to workShifts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @workShift = WorkShift.find(params[:id])
    @workShift.destroy
    redirect_to workShifts_path
  end

 
  private
    def workShift_params
      params.require(:workShift).permit(:startTime, :endTime, :breakMinutes, :DayOfWeek)
    end
end