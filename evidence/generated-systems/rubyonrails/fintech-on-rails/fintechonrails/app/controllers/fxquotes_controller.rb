class FXQuotesController < ApplicationController
  def index
    @fXQuotes = FXQuote.all
  end
 
  def show
    @fXQuote = FXQuote.find(params[:id])
  end
 
  def new
    @fXQuote = FXQuote.new
  end
 
  def edit
    @fXQuote = FXQuote.find(params[:id])
  end
 
  def create
    @fXQuote = FXQuote.new(fXQuote_params)
 
    if @fXQuote.save
      redirect_to fXQuotes_path
    else
      render 'new'
    end
  end
 
  def update
    @fXQuote = FXQuote.find(params[:id])
 
    if @fXQuote.update(fXQuote_params)
      redirect_to fXQuotes_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @fXQuote = FXQuote.find(params[:id])
    @fXQuote.destroy
    redirect_to fXQuotes_path
  end

 
  private
    def fXQuote_params
      params.require(:fXQuote).permit(:baseCurrency, :quoteCurrency, :rate, :quotedAt, :expiresAt, :PriceType)
    end
end