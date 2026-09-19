class InsurancePayersController < ApplicationController
  def index
    @insurancePayers = InsurancePayer.all
  end
 
  def show
    @insurancePayer = InsurancePayer.find(params[:id])
  end
 
  def new
    @insurancePayer = InsurancePayer.new
  end
 
  def edit
    @insurancePayer = InsurancePayer.find(params[:id])
  end
 
  def create
    @insurancePayer = InsurancePayer.new(insurancePayer_params)
 
    if @insurancePayer.save
      redirect_to insurancePayers_path
    else
      render 'new'
    end
  end
 
  def update
    @insurancePayer = InsurancePayer.find(params[:id])
 
    if @insurancePayer.update(insurancePayer_params)
      redirect_to insurancePayers_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @insurancePayer = InsurancePayer.find(params[:id])
    @insurancePayer.destroy
    redirect_to insurancePayers_path
  end

 
  private
    def insurancePayer_params
      params.require(:insurancePayer).permit(:name, :website, :PayerType)
    end
end