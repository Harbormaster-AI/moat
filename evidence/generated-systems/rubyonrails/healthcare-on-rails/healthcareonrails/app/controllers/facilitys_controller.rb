class FacilitysController < ApplicationController
  def index
    @facilitys = Facility.all
  end
 
  def show
    @facility = Facility.find(params[:id])
  end
 
  def new
    @facility = Facility.new
  end
 
  def edit
    @facility = Facility.find(params[:id])
  end
 
  def create
    @facility = Facility.new(facility_params)
 
    if @facility.save
      redirect_to facilitys_path
    else
      render 'new'
    end
  end
 
  def update
    @facility = Facility.find(params[:id])
 
    if @facility.update(facility_params)
      redirect_to facilitys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @facility = Facility.find(params[:id])
    @facility.destroy
    redirect_to facilitys_path
  end

 
  private
    def facility_params
      params.require(:facility).permit(:name, :facilityCode, :address, :FacilityType)
    end
end