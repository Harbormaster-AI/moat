class CycleCountEntrysController < ApplicationController
  def index
    @cycleCountEntrys = CycleCountEntry.all
  end
 
  def show
    @cycleCountEntry = CycleCountEntry.find(params[:id])
  end
 
  def new
    @cycleCountEntry = CycleCountEntry.new
  end
 
  def edit
    @cycleCountEntry = CycleCountEntry.find(params[:id])
  end
 
  def create
    @cycleCountEntry = CycleCountEntry.new(cycleCountEntry_params)
 
    if @cycleCountEntry.save
      redirect_to cycleCountEntrys_path
    else
      render 'new'
    end
  end
 
  def update
    @cycleCountEntry = CycleCountEntry.find(params[:id])
 
    if @cycleCountEntry.update(cycleCountEntry_params)
      redirect_to cycleCountEntrys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @cycleCountEntry = CycleCountEntry.find(params[:id])
    @cycleCountEntry.destroy
    redirect_to cycleCountEntrys_path
  end

 
  private
    def cycleCountEntry_params
      params.require(:cycleCountEntry).permit(:lineNumber, :systemQuantity, :countedQuantity, :varianceQuantity, :recountRequired, :StockStatus)
    end
end