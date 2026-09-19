class AllergysController < ApplicationController
  def index
    @allergys = Allergy.all
  end
 
  def show
    @allergy = Allergy.find(params[:id])
  end
 
  def new
    @allergy = Allergy.new
  end
 
  def edit
    @allergy = Allergy.find(params[:id])
  end
 
  def create
    @allergy = Allergy.new(allergy_params)
 
    if @allergy.save
      redirect_to allergys_path
    else
      render 'new'
    end
  end
 
  def update
    @allergy = Allergy.find(params[:id])
 
    if @allergy.update(allergy_params)
      redirect_to allergys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @allergy = Allergy.find(params[:id])
    @allergy.destroy
    redirect_to allergys_path
  end

 
  private
    def allergy_params
      params.require(:allergy).permit(:substance, :reaction, :Severity, :Status)
    end
end