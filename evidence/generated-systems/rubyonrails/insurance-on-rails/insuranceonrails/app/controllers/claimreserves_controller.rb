class ClaimReservesController < ApplicationController
  def index
    @claimReserves = ClaimReserve.all
  end
 
  def show
    @claimReserve = ClaimReserve.find(params[:id])
  end
 
  def new
    @claimReserve = ClaimReserve.new
  end
 
  def edit
    @claimReserve = ClaimReserve.find(params[:id])
  end
 
  def create
    @claimReserve = ClaimReserve.new(claimReserve_params)
 
    if @claimReserve.save
      redirect_to claimReserves_path
    else
      render 'new'
    end
  end
 
  def update
    @claimReserve = ClaimReserve.find(params[:id])
 
    if @claimReserve.update(claimReserve_params)
      redirect_to claimReserves_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @claimReserve = ClaimReserve.find(params[:id])
    @claimReserve.destroy
    redirect_to claimReserves_path
  end

 
  private
    def claimReserve_params
      params.require(:claimReserve).permit(:amount, :setDate, :ReserveType, :Status)
    end
end