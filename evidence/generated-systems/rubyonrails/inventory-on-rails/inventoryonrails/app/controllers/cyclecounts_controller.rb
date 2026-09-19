class CycleCountsController < ApplicationController
  def index
    @cycleCounts = CycleCount.all
  end
 
  def show
    @cycleCount = CycleCount.find(params[:id])
  end
 
  def new
    @cycleCount = CycleCount.new
  end
 
  def edit
    @cycleCount = CycleCount.find(params[:id])
  end
 
  def create
    @cycleCount = CycleCount.new(cycleCount_params)
 
    if @cycleCount.save
      redirect_to cycleCounts_path
    else
      render 'new'
    end
  end
 
  def update
    @cycleCount = CycleCount.find(params[:id])
 
    if @cycleCount.update(cycleCount_params)
      redirect_to cycleCounts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @cycleCount = CycleCount.find(params[:id])
    @cycleCount.destroy
    redirect_to cycleCounts_path
  end

 
  private
    def cycleCount_params
      params.require(:cycleCount).permit(:countNumber, :scheduledDate, :performedDate, :approvedBy, :Status)
    end
end