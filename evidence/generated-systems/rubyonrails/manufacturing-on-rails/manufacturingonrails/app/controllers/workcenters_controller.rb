class WorkCentersController < ApplicationController
  def index
    @workCenters = WorkCenter.all
  end
 
  def show
    @workCenter = WorkCenter.find(params[:id])
  end
 
  def new
    @workCenter = WorkCenter.new
  end
 
  def edit
    @workCenter = WorkCenter.find(params[:id])
  end
 
  def create
    @workCenter = WorkCenter.new(workCenter_params)
 
    if @workCenter.save
      redirect_to workCenters_path
    else
      render 'new'
    end
  end
 
  def update
    @workCenter = WorkCenter.find(params[:id])
 
    if @workCenter.update(workCenter_params)
      redirect_to workCenters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @workCenter = WorkCenter.find(params[:id])
    @workCenter.destroy
    redirect_to workCenters_path
  end

 
  private
    def workCenter_params
      params.require(:workCenter).permit(:name, :code, :capacityPerHour, :oeeTarget, :WorkCenterType)
    end
end