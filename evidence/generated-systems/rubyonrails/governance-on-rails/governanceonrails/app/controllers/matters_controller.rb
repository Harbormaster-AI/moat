class MattersController < ApplicationController
  def index
    @matters = Matter.all
  end
 
  def show
    @matter = Matter.find(params[:id])
  end
 
  def new
    @matter = Matter.new
  end
 
  def edit
    @matter = Matter.find(params[:id])
  end
 
  def create
    @matter = Matter.new(matter_params)
 
    if @matter.save
      redirect_to matters_path
    else
      render 'new'
    end
  end
 
  def update
    @matter = Matter.find(params[:id])
 
    if @matter.update(matter_params)
      redirect_to matters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @matter = Matter.find(params[:id])
    @matter.destroy
    redirect_to matters_path
  end

 
  private
    def matter_params
      params.require(:matter).permit(:matterName, :leadCounsel, :MatterType, :Status)
    end
end