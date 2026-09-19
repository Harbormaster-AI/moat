class SalesRegionsController < ApplicationController
  def index
    @salesRegions = SalesRegion.all
  end
 
  def show
    @salesRegion = SalesRegion.find(params[:id])
  end
 
  def new
    @salesRegion = SalesRegion.new
  end
 
  def edit
    @salesRegion = SalesRegion.find(params[:id])
  end
 
  def create
    @salesRegion = SalesRegion.new(salesRegion_params)
 
    if @salesRegion.save
      redirect_to salesRegions_path
    else
      render 'new'
    end
  end
 
  def update
    @salesRegion = SalesRegion.find(params[:id])
 
    if @salesRegion.update(salesRegion_params)
      redirect_to salesRegions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @salesRegion = SalesRegion.find(params[:id])
    @salesRegion.destroy
    redirect_to salesRegions_path
  end

 
  private
    def salesRegion_params
      params.require(:salesRegion).permit(:name, :regionCode)
    end
end