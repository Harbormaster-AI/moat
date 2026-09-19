class PerformanceCyclesController < ApplicationController
  def index
    @performanceCycles = PerformanceCycle.all
  end
 
  def show
    @performanceCycle = PerformanceCycle.find(params[:id])
  end
 
  def new
    @performanceCycle = PerformanceCycle.new
  end
 
  def edit
    @performanceCycle = PerformanceCycle.find(params[:id])
  end
 
  def create
    @performanceCycle = PerformanceCycle.new(performanceCycle_params)
 
    if @performanceCycle.save
      redirect_to performanceCycles_path
    else
      render 'new'
    end
  end
 
  def update
    @performanceCycle = PerformanceCycle.find(params[:id])
 
    if @performanceCycle.update(performanceCycle_params)
      redirect_to performanceCycles_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @performanceCycle = PerformanceCycle.find(params[:id])
    @performanceCycle.destroy
    redirect_to performanceCycles_path
  end

 
  private
    def performanceCycle_params
      params.require(:performanceCycle).permit(:name, :startDate, :endDate, :Status)
    end
end