class DemandSignalsController < ApplicationController
  def index
    @demandSignals = DemandSignal.all
  end
 
  def show
    @demandSignal = DemandSignal.find(params[:id])
  end
 
  def new
    @demandSignal = DemandSignal.new
  end
 
  def edit
    @demandSignal = DemandSignal.find(params[:id])
  end
 
  def create
    @demandSignal = DemandSignal.new(demandSignal_params)
 
    if @demandSignal.save
      redirect_to demandSignals_path
    else
      render 'new'
    end
  end
 
  def update
    @demandSignal = DemandSignal.find(params[:id])
 
    if @demandSignal.update(demandSignal_params)
      redirect_to demandSignals_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @demandSignal = DemandSignal.find(params[:id])
    @demandSignal.destroy
    redirect_to demandSignals_path
  end

 
  private
    def demandSignal_params
      params.require(:demandSignal).permit(:externalReference, :requestedDate, :quantity, :DemandType)
    end
end