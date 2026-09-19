class StockAdjustmentLinesController < ApplicationController
  def index
    @stockAdjustmentLines = StockAdjustmentLine.all
  end
 
  def show
    @stockAdjustmentLine = StockAdjustmentLine.find(params[:id])
  end
 
  def new
    @stockAdjustmentLine = StockAdjustmentLine.new
  end
 
  def edit
    @stockAdjustmentLine = StockAdjustmentLine.find(params[:id])
  end
 
  def create
    @stockAdjustmentLine = StockAdjustmentLine.new(stockAdjustmentLine_params)
 
    if @stockAdjustmentLine.save
      redirect_to stockAdjustmentLines_path
    else
      render 'new'
    end
  end
 
  def update
    @stockAdjustmentLine = StockAdjustmentLine.find(params[:id])
 
    if @stockAdjustmentLine.update(stockAdjustmentLine_params)
      redirect_to stockAdjustmentLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @stockAdjustmentLine = StockAdjustmentLine.find(params[:id])
    @stockAdjustmentLine.destroy
    redirect_to stockAdjustmentLines_path
  end

 
  private
    def stockAdjustmentLine_params
      params.require(:stockAdjustmentLine).permit(:lineNumber, :quantity, :UnitOfMeasure, :StockStatus)
    end
end