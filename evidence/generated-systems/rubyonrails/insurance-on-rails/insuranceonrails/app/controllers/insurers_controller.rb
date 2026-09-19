class InsurersController < ApplicationController
  def index
    @insurers = Insurer.all
  end
 
  def show
    @insurer = Insurer.find(params[:id])
  end
 
  def new
    @insurer = Insurer.new
  end
 
  def edit
    @insurer = Insurer.find(params[:id])
  end
 
  def create
    @insurer = Insurer.new(insurer_params)
 
    if @insurer.save
      redirect_to insurers_path
    else
      render 'new'
    end
  end
 
  def update
    @insurer = Insurer.find(params[:id])
 
    if @insurer.update(insurer_params)
      redirect_to insurers_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @insurer = Insurer.find(params[:id])
    @insurer.destroy
    redirect_to insurers_path
  end

 
  private
    def insurer_params
      params.require(:insurer).permit(:name, :legalName, :domicileCountry, :naicNumber, :website)
    end
end