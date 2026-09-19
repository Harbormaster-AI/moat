class PriceBookEntrysController < ApplicationController
  def index
    @priceBookEntrys = PriceBookEntry.all
  end
 
  def show
    @priceBookEntry = PriceBookEntry.find(params[:id])
  end
 
  def new
    @priceBookEntry = PriceBookEntry.new
  end
 
  def edit
    @priceBookEntry = PriceBookEntry.find(params[:id])
  end
 
  def create
    @priceBookEntry = PriceBookEntry.new(priceBookEntry_params)
 
    if @priceBookEntry.save
      redirect_to priceBookEntrys_path
    else
      render 'new'
    end
  end
 
  def update
    @priceBookEntry = PriceBookEntry.find(params[:id])
 
    if @priceBookEntry.update(priceBookEntry_params)
      redirect_to priceBookEntrys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @priceBookEntry = PriceBookEntry.find(params[:id])
    @priceBookEntry.destroy
    redirect_to priceBookEntrys_path
  end

 
  private
    def priceBookEntry_params
      params.require(:priceBookEntry).permit(:unitPrice, :effectiveDate, :expirationDate, :asActive)
    end
end