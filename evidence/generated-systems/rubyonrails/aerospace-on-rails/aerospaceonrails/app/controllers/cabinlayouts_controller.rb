class CabinLayoutsController < ApplicationController
  def index
    @cabinLayouts = CabinLayout.all
  end
 
  def show
    @cabinLayout = CabinLayout.find(params[:id])
  end
 
  def new
    @cabinLayout = CabinLayout.new
  end
 
  def edit
    @cabinLayout = CabinLayout.find(params[:id])
  end
 
  def create
    @cabinLayout = CabinLayout.new(cabinLayout_params)
 
    if @cabinLayout.save
      redirect_to cabinLayouts_path
    else
      render 'new'
    end
  end
 
  def update
    @cabinLayout = CabinLayout.find(params[:id])
 
    if @cabinLayout.update(cabinLayout_params)
      redirect_to cabinLayouts_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @cabinLayout = CabinLayout.find(params[:id])
    @cabinLayout.destroy
    redirect_to cabinLayouts_path
  end

 
  private
    def cabinLayout_params
      params.require(:cabinLayout).permit(:layoutCode, :totalSeats, :classLayout)
    end
end