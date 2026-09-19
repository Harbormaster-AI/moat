class DistributorsController < ApplicationController
  def index
    @distributors = Distributor.all
  end
 
  def show
    @distributor = Distributor.find(params[:id])
  end
 
  def new
    @distributor = Distributor.new
  end
 
  def edit
    @distributor = Distributor.find(params[:id])
  end
 
  def create
    @distributor = Distributor.new(distributor_params)
 
    if @distributor.save
      redirect_to distributors_path
    else
      render 'new'
    end
  end
 
  def update
    @distributor = Distributor.find(params[:id])
 
    if @distributor.update(distributor_params)
      redirect_to distributors_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @distributor = Distributor.find(params[:id])
    @distributor.destroy
    redirect_to distributors_path
  end

 
  private
    def distributor_params
      params.require(:distributor).permit(:name, :licenseNumber, :region, :DistributorType)
    end
end