class PayoutsController < ApplicationController
  def index
    @payouts = Payout.all
  end
 
  def show
    @payout = Payout.find(params[:id])
  end
 
  def new
    @payout = Payout.new
  end
 
  def edit
    @payout = Payout.find(params[:id])
  end
 
  def create
    @payout = Payout.new(payout_params)
 
    if @payout.save
      redirect_to payouts_path
    else
      render 'new'
    end
  end
 
  def update
    @payout = Payout.find(params[:id])
 
    if @payout.update(payout_params)
      redirect_to payouts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @payout = Payout.find(params[:id])
    @payout.destroy
    redirect_to payouts_path
  end

 
  private
    def payout_params
      params.require(:payout).permit(:payoutNumber, :amount, :scheduledDate, :paidDate, :Status)
    end
end