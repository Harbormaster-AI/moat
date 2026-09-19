class ProductionLinesController < ApplicationController
  def index
    @productionLines = ProductionLine.all
  end
 
  def show
    @productionLine = ProductionLine.find(params[:id])
  end
 
  def new
    @productionLine = ProductionLine.new
  end
 
  def edit
    @productionLine = ProductionLine.find(params[:id])
  end
 
  def create
    @productionLine = ProductionLine.new(productionLine_params)
 
    if @productionLine.save
      redirect_to productionLines_path
    else
      render 'new'
    end
  end
 
  def update
    @productionLine = ProductionLine.find(params[:id])
 
    if @productionLine.update(productionLine_params)
      redirect_to productionLines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @productionLine = ProductionLine.find(params[:id])
    @productionLine.destroy
    redirect_to productionLines_path
  end

 
  private
    def productionLine_params
      params.require(:productionLine).permit(:name, :LineType)
    end
end