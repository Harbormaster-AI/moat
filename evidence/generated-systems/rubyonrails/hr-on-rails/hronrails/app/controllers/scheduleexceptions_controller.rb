class ScheduleExceptionsController < ApplicationController
  def index
    @scheduleExceptions = ScheduleException.all
  end
 
  def show
    @scheduleException = ScheduleException.find(params[:id])
  end
 
  def new
    @scheduleException = ScheduleException.new
  end
 
  def edit
    @scheduleException = ScheduleException.find(params[:id])
  end
 
  def create
    @scheduleException = ScheduleException.new(scheduleException_params)
 
    if @scheduleException.save
      redirect_to scheduleExceptions_path
    else
      render 'new'
    end
  end
 
  def update
    @scheduleException = ScheduleException.find(params[:id])
 
    if @scheduleException.update(scheduleException_params)
      redirect_to scheduleExceptions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @scheduleException = ScheduleException.find(params[:id])
    @scheduleException.destroy
    redirect_to scheduleExceptions_path
  end

 
  private
    def scheduleException_params
      params.require(:scheduleException).permit(:date, :reason, :hours)
    end
end