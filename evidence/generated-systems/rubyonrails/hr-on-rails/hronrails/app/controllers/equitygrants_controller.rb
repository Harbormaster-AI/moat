class EquityGrantsController < ApplicationController
  def index
    @equityGrants = EquityGrant.all
  end
 
  def show
    @equityGrant = EquityGrant.find(params[:id])
  end
 
  def new
    @equityGrant = EquityGrant.new
  end
 
  def edit
    @equityGrant = EquityGrant.find(params[:id])
  end
 
  def create
    @equityGrant = EquityGrant.new(equityGrant_params)
 
    if @equityGrant.save
      redirect_to equityGrants_path
    else
      render 'new'
    end
  end
 
  def update
    @equityGrant = EquityGrant.find(params[:id])
 
    if @equityGrant.update(equityGrant_params)
      redirect_to equityGrants_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @equityGrant = EquityGrant.find(params[:id])
    @equityGrant.destroy
    redirect_to equityGrants_path
  end

 
  private
    def equityGrant_params
      params.require(:equityGrant).permit(:grantId, :grantedUnits, :vestingStart, :GrantType)
    end
end