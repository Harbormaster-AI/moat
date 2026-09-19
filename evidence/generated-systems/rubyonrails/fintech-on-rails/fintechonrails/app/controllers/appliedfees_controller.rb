class AppliedFeesController < ApplicationController
  def index
    @appliedFees = AppliedFee.all
  end
 
  def show
    @appliedFee = AppliedFee.find(params[:id])
  end
 
  def new
    @appliedFee = AppliedFee.new
  end
 
  def edit
    @appliedFee = AppliedFee.find(params[:id])
  end
 
  def create
    @appliedFee = AppliedFee.new(appliedFee_params)
 
    if @appliedFee.save
      redirect_to appliedFees_path
    else
      render 'new'
    end
  end
 
  def update
    @appliedFee = AppliedFee.find(params[:id])
 
    if @appliedFee.update(appliedFee_params)
      redirect_to appliedFees_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @appliedFee = AppliedFee.find(params[:id])
    @appliedFee.destroy
    redirect_to appliedFees_path
  end

 
  private
    def appliedFee_params
      params.require(:appliedFee).permit(:amount, :description, :FeeType)
    end
end