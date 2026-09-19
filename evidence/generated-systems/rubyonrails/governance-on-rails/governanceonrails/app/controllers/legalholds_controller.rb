class LegalHoldsController < ApplicationController
  def index
    @legalHolds = LegalHold.all
  end
 
  def show
    @legalHold = LegalHold.find(params[:id])
  end
 
  def new
    @legalHold = LegalHold.new
  end
 
  def edit
    @legalHold = LegalHold.find(params[:id])
  end
 
  def create
    @legalHold = LegalHold.new(legalHold_params)
 
    if @legalHold.save
      redirect_to legalHolds_path
    else
      render 'new'
    end
  end
 
  def update
    @legalHold = LegalHold.find(params[:id])
 
    if @legalHold.update(legalHold_params)
      redirect_to legalHolds_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @legalHold = LegalHold.find(params[:id])
    @legalHold.destroy
    redirect_to legalHolds_path
  end

 
  private
    def legalHold_params
      params.require(:legalHold).permit(:name, :reason, :issuedDate, :releaseDate, :HoldStatus)
    end
end