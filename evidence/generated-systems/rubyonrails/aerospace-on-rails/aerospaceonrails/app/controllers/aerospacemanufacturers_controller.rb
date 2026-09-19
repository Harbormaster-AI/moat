class AerospaceManufacturersController < ApplicationController
  def index
    @aerospaceManufacturers = AerospaceManufacturer.all
  end
 
  def show
    @aerospaceManufacturer = AerospaceManufacturer.find(params[:id])
  end
 
  def new
    @aerospaceManufacturer = AerospaceManufacturer.new
  end
 
  def edit
    @aerospaceManufacturer = AerospaceManufacturer.find(params[:id])
  end
 
  def create
    @aerospaceManufacturer = AerospaceManufacturer.new(aerospaceManufacturer_params)
 
    if @aerospaceManufacturer.save
      redirect_to aerospaceManufacturers_path
    else
      render 'new'
    end
  end
 
  def update
    @aerospaceManufacturer = AerospaceManufacturer.find(params[:id])
 
    if @aerospaceManufacturer.update(aerospaceManufacturer_params)
      redirect_to aerospaceManufacturers_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aerospaceManufacturer = AerospaceManufacturer.find(params[:id])
    @aerospaceManufacturer.destroy
    redirect_to aerospaceManufacturers_path
  end

 
  private
    def aerospaceManufacturer_params
      params.require(:aerospaceManufacturer).permit(:name, :legalName, :headquartersCountry, :website)
    end
end