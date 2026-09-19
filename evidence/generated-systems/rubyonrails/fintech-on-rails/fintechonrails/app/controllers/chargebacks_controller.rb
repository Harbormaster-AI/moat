class ChargebacksController < ApplicationController
  def index
    @chargebacks = Chargeback.all
  end
 
  def show
    @chargeback = Chargeback.find(params[:id])
  end
 
  def new
    @chargeback = Chargeback.new
  end
 
  def edit
    @chargeback = Chargeback.find(params[:id])
  end
 
  def create
    @chargeback = Chargeback.new(chargeback_params)
 
    if @chargeback.save
      redirect_to chargebacks_path
    else
      render 'new'
    end
  end
 
  def update
    @chargeback = Chargeback.find(params[:id])
 
    if @chargeback.update(chargeback_params)
      redirect_to chargebacks_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @chargeback = Chargeback.find(params[:id])
    @chargeback.destroy
    redirect_to chargebacks_path
  end

 
  private
    def chargeback_params
      params.require(:chargeback).permit(:chargebackReference, :amount, :postedAt, :Stage, :Status)
    end
end