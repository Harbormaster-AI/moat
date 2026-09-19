class OpportunityLineItemsController < ApplicationController
  def index
    @opportunityLineItems = OpportunityLineItem.all
  end
 
  def show
    @opportunityLineItem = OpportunityLineItem.find(params[:id])
  end
 
  def new
    @opportunityLineItem = OpportunityLineItem.new
  end
 
  def edit
    @opportunityLineItem = OpportunityLineItem.find(params[:id])
  end
 
  def create
    @opportunityLineItem = OpportunityLineItem.new(opportunityLineItem_params)
 
    if @opportunityLineItem.save
      redirect_to opportunityLineItems_path
    else
      render 'new'
    end
  end
 
  def update
    @opportunityLineItem = OpportunityLineItem.find(params[:id])
 
    if @opportunityLineItem.update(opportunityLineItem_params)
      redirect_to opportunityLineItems_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @opportunityLineItem = OpportunityLineItem.find(params[:id])
    @opportunityLineItem.destroy
    redirect_to opportunityLineItems_path
  end

 
  private
    def opportunityLineItem_params
      params.require(:opportunityLineItem).permit(:quantity, :unitPrice, :discountPercent, :totalPrice)
    end
end