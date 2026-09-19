class LandingGearsController < ApplicationController
  def index
    @landingGears = LandingGear.all
  end
 
  def show
    @landingGear = LandingGear.find(params[:id])
  end
 
  def new
    @landingGear = LandingGear.new
  end
 
  def edit
    @landingGear = LandingGear.find(params[:id])
  end
 
  def create
    @landingGear = LandingGear.new(landingGear_params)
 
    if @landingGear.save
      redirect_to landingGears_path
    else
      render 'new'
    end
  end
 
  def update
    @landingGear = LandingGear.find(params[:id])
 
    if @landingGear.update(landingGear_params)
      redirect_to landingGears_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @landingGear = LandingGear.find(params[:id])
    @landingGear.destroy
    redirect_to landingGears_path
  end

 
  private
    def landingGear_params
      params.require(:landingGear).permit(:supplierPartNumber, :GearType)
    end
end