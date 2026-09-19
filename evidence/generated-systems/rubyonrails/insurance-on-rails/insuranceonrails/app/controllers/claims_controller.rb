class ClaimsController < ApplicationController
  def index
    @claims = Claim.all
  end
 
  def show
    @claim = Claim.find(params[:id])
  end
 
  def new
    @claim = Claim.new
  end
 
  def edit
    @claim = Claim.find(params[:id])
  end
 
  def create
    @claim = Claim.new(claim_params)
 
    if @claim.save
      redirect_to claims_path
    else
      render 'new'
    end
  end
 
  def update
    @claim = Claim.find(params[:id])
 
    if @claim.update(claim_params)
      redirect_to claims_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @claim = Claim.find(params[:id])
    @claim.destroy
    redirect_to claims_path
  end

 
  private
    def claim_params
      params.require(:claim).permit(:claimNumber, :noticeDate, :lossDate, :reportedBy, :Status, :LossCause)
    end
end