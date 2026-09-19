class TerritorysController < ApplicationController
  def index
    @territorys = Territory.all
  end
 
  def show
    @territory = Territory.find(params[:id])
  end
 
  def new
    @territory = Territory.new
  end
 
  def edit
    @territory = Territory.find(params[:id])
  end
 
  def create
    @territory = Territory.new(territory_params)
 
    if @territory.save
      redirect_to territorys_path
    else
      render 'new'
    end
  end
 
  def update
    @territory = Territory.find(params[:id])
 
    if @territory.update(territory_params)
      redirect_to territorys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @territory = Territory.find(params[:id])
    @territory.destroy
    redirect_to territorys_path
  end

 
  private
    def territory_params
      params.require(:territory).permit(:name, :region, :TerritoryType)
    end
end