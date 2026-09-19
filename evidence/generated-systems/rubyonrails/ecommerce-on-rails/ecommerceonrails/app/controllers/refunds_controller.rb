class RefundsController < ApplicationController
  def index
    @refunds = Refund.all
  end
 
  def show
    @refund = Refund.find(params[:id])
  end
 
  def new
    @refund = Refund.new
  end
 
  def edit
    @refund = Refund.find(params[:id])
  end
 
  def create
    @refund = Refund.new(refund_params)
 
    if @refund.save
      redirect_to refunds_path
    else
      render 'new'
    end
  end
 
  def update
    @refund = Refund.find(params[:id])
 
    if @refund.update(refund_params)
      redirect_to refunds_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @refund = Refund.find(params[:id])
    @refund.destroy
    redirect_to refunds_path
  end

 
  private
    def refund_params
      params.require(:refund).permit(:refundNumber, :amount, :reason, :createdAt, :Status)
    end
end