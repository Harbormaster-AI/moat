class VerifiedAddresssController < ApplicationController
  def index
    @verifiedAddresss = VerifiedAddress.all
  end
 
  def show
    @verifiedAddress = VerifiedAddress.find(params[:id])
  end
 
  def new
    @verifiedAddress = VerifiedAddress.new
  end
 
  def edit
    @verifiedAddress = VerifiedAddress.find(params[:id])
  end
 
  def create
    @verifiedAddress = VerifiedAddress.new(verifiedAddress_params)
 
    if @verifiedAddress.save
      redirect_to verifiedAddresss_path
    else
      render 'new'
    end
  end
 
  def update
    @verifiedAddress = VerifiedAddress.find(params[:id])
 
    if @verifiedAddress.update(verifiedAddress_params)
      redirect_to verifiedAddresss_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @verifiedAddress = VerifiedAddress.find(params[:id])
    @verifiedAddress.destroy
    redirect_to verifiedAddresss_path
  end

 
  private
    def verifiedAddress_params
      params.require(:verifiedAddress).permit(:address, :verifiedAt, :VerificationStatus)
    end
end