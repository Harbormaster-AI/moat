class FraudSignalsController < ApplicationController
  def index
    @fraudSignals = FraudSignal.all
  end
 
  def show
    @fraudSignal = FraudSignal.find(params[:id])
  end
 
  def new
    @fraudSignal = FraudSignal.new
  end
 
  def edit
    @fraudSignal = FraudSignal.find(params[:id])
  end
 
  def create
    @fraudSignal = FraudSignal.new(fraudSignal_params)
 
    if @fraudSignal.save
      redirect_to fraudSignals_path
    else
      render 'new'
    end
  end
 
  def update
    @fraudSignal = FraudSignal.find(params[:id])
 
    if @fraudSignal.update(fraudSignal_params)
      redirect_to fraudSignals_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @fraudSignal = FraudSignal.find(params[:id])
    @fraudSignal.destroy
    redirect_to fraudSignals_path
  end

 
  private
    def fraudSignal_params
      params.require(:fraudSignal).permit(:name, :ruleLogic, :SignalType)
    end
end