class StockKeepingUnitsController < ApplicationController
  def index
    @stockKeepingUnits = StockKeepingUnit.all
  end
 
  def show
    @stockKeepingUnit = StockKeepingUnit.find(params[:id])
  end
 
  def new
    @stockKeepingUnit = StockKeepingUnit.new
  end
 
  def edit
    @stockKeepingUnit = StockKeepingUnit.find(params[:id])
  end
 
  def create
    @stockKeepingUnit = StockKeepingUnit.new(stockKeepingUnit_params)
 
    if @stockKeepingUnit.save
      redirect_to stockKeepingUnits_path
    else
      render 'new'
    end
  end
 
  def update
    @stockKeepingUnit = StockKeepingUnit.find(params[:id])
 
    if @stockKeepingUnit.update(stockKeepingUnit_params)
      redirect_to stockKeepingUnits_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @stockKeepingUnit = StockKeepingUnit.find(params[:id])
    @stockKeepingUnit.destroy
    redirect_to stockKeepingUnits_path
  end

 
  private
    def stockKeepingUnit_params
      params.require(:stockKeepingUnit).permit(:skuCode, :name, :weight, :weightUnit, :volume, :volumeUnit, :shelfLifeDays, :hazardousMaterial, :ItemType, :UnitOfMeasure)
    end
end