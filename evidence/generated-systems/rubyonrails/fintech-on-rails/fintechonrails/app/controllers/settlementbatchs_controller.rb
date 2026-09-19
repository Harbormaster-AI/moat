class SettlementBatchsController < ApplicationController
  def index
    @settlementBatchs = SettlementBatch.all
  end
 
  def show
    @settlementBatch = SettlementBatch.find(params[:id])
  end
 
  def new
    @settlementBatch = SettlementBatch.new
  end
 
  def edit
    @settlementBatch = SettlementBatch.find(params[:id])
  end
 
  def create
    @settlementBatch = SettlementBatch.new(settlementBatch_params)
 
    if @settlementBatch.save
      redirect_to settlementBatchs_path
    else
      render 'new'
    end
  end
 
  def update
    @settlementBatch = SettlementBatch.find(params[:id])
 
    if @settlementBatch.update(settlementBatch_params)
      redirect_to settlementBatchs_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @settlementBatch = SettlementBatch.find(params[:id])
    @settlementBatch.destroy
    redirect_to settlementBatchs_path
  end

 
  private
    def settlementBatch_params
      params.require(:settlementBatch).permit(:batchId, :periodStart, :periodEnd, :totalVolume, :totalCount, :Status)
    end
end