class StockAdjustmentsController < ApplicationController
  def index
    @stockAdjustments = StockAdjustment.all
  end
 
  def show
    @stockAdjustment = StockAdjustment.find(params[:id])
  end
 
  def new
    @stockAdjustment = StockAdjustment.new
  end
 
  def edit
    @stockAdjustment = StockAdjustment.find(params[:id])
  end
 
  def create
    @stockAdjustment = StockAdjustment.new(stockAdjustment_params)
 
    if @stockAdjustment.save
      redirect_to stockAdjustments_path
    else
      render 'new'
    end
  end
 
  def update
    @stockAdjustment = StockAdjustment.find(params[:id])
 
    if @stockAdjustment.update(stockAdjustment_params)
      redirect_to stockAdjustments_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @stockAdjustment = StockAdjustment.find(params[:id])
    @stockAdjustment.destroy
    redirect_to stockAdjustments_path
  end

 
  private
    def stockAdjustment_params
      params.require(:stockAdjustment).permit(:adjustmentNumber, :reason, :adjustmentDate, :AdjustmentType, :Status)
    end
end