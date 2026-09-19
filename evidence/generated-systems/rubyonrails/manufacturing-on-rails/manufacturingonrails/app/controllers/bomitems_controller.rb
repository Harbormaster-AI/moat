class BOMItemsController < ApplicationController
  def index
    @bOMItems = BOMItem.all
  end
 
  def show
    @bOMItem = BOMItem.find(params[:id])
  end
 
  def new
    @bOMItem = BOMItem.new
  end
 
  def edit
    @bOMItem = BOMItem.find(params[:id])
  end
 
  def create
    @bOMItem = BOMItem.new(bOMItem_params)
 
    if @bOMItem.save
      redirect_to bOMItems_path
    else
      render 'new'
    end
  end
 
  def update
    @bOMItem = BOMItem.find(params[:id])
 
    if @bOMItem.update(bOMItem_params)
      redirect_to bOMItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @bOMItem = BOMItem.find(params[:id])
    @bOMItem.destroy
    redirect_to bOMItems_path
  end

 
  private
    def bOMItem_params
      params.require(:bOMItem).permit(:lineNumber, :quantity, :scrapPercent)
    end
end