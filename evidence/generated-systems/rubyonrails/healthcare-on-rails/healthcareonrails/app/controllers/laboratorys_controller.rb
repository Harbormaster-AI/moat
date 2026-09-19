class LaboratorysController < ApplicationController
  def index
    @laboratorys = Laboratory.all
  end
 
  def show
    @laboratory = Laboratory.find(params[:id])
  end
 
  def new
    @laboratory = Laboratory.new
  end
 
  def edit
    @laboratory = Laboratory.find(params[:id])
  end
 
  def create
    @laboratory = Laboratory.new(laboratory_params)
 
    if @laboratory.save
      redirect_to laboratorys_path
    else
      render 'new'
    end
  end
 
  def update
    @laboratory = Laboratory.find(params[:id])
 
    if @laboratory.update(laboratory_params)
      redirect_to laboratorys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @laboratory = Laboratory.find(params[:id])
    @laboratory.destroy
    redirect_to laboratorys_path
  end

 
  private
    def laboratory_params
      params.require(:laboratory).permit(:name, :cliaNumber)
    end
end