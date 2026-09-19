class BuildSchedulesController < ApplicationController
  def index
    @buildSchedules = BuildSchedule.all
  end
 
  def show
    @buildSchedule = BuildSchedule.find(params[:id])
  end
 
  def new
    @buildSchedule = BuildSchedule.new
  end
 
  def edit
    @buildSchedule = BuildSchedule.find(params[:id])
  end
 
  def create
    @buildSchedule = BuildSchedule.new(buildSchedule_params)
 
    if @buildSchedule.save
      redirect_to buildSchedules_path
    else
      render 'new'
    end
  end
 
  def update
    @buildSchedule = BuildSchedule.find(params[:id])
 
    if @buildSchedule.update(buildSchedule_params)
      redirect_to buildSchedules_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @buildSchedule = BuildSchedule.find(params[:id])
    @buildSchedule.destroy
    redirect_to buildSchedules_path
  end

 
  private
    def buildSchedule_params
      params.require(:buildSchedule).permit(:scheduleNumber, :Status)
    end
end