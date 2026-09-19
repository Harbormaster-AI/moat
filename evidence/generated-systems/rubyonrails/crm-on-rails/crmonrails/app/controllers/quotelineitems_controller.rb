class QuoteLineItemsController < ApplicationController
  def index
    @quoteLineItems = QuoteLineItem.all
  end
 
  def show
    @quoteLineItem = QuoteLineItem.find(params[:id])
  end
 
  def new
    @quoteLineItem = QuoteLineItem.new
  end
 
  def edit
    @quoteLineItem = QuoteLineItem.find(params[:id])
  end
 
  def create
    @quoteLineItem = QuoteLineItem.new(quoteLineItem_params)
 
    if @quoteLineItem.save
      redirect_to quoteLineItems_path
    else
      render 'new'
    end
  end
 
  def update
    @quoteLineItem = QuoteLineItem.find(params[:id])
 
    if @quoteLineItem.update(quoteLineItem_params)
      redirect_to quoteLineItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @quoteLineItem = QuoteLineItem.find(params[:id])
    @quoteLineItem.destroy
    redirect_to quoteLineItems_path
  end

 
  private
    def quoteLineItem_params
      params.require(:quoteLineItem).permit(:quantity, :unitPrice, :discountAmount, :taxAmount, :totalAmount)
    end
end