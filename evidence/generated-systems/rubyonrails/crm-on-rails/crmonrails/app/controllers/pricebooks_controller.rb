class PriceBooksController < ApplicationController
  def index
    @priceBooks = PriceBook.all
  end
 
  def show
    @priceBook = PriceBook.find(params[:id])
  end
 
  def new
    @priceBook = PriceBook.new
  end
 
  def edit
    @priceBook = PriceBook.find(params[:id])
  end
 
  def create
    @priceBook = PriceBook.new(priceBook_params)
 
    if @priceBook.save
      redirect_to priceBooks_path
    else
      render 'new'
    end
  end
 
  def update
    @priceBook = PriceBook.find(params[:id])
 
    if @priceBook.update(priceBook_params)
      redirect_to priceBooks_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @priceBook = PriceBook.find(params[:id])
    @priceBook.destroy
    redirect_to priceBooks_path
  end

 
  private
    def priceBook_params
      params.require(:priceBook).permit(:name, :asActive, :description)
    end
end