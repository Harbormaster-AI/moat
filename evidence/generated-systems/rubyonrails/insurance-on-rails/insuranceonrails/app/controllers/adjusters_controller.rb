class AdjustersController < ApplicationController
  def index
    @adjusters = Adjuster.all
  end
 
  def show
    @adjuster = Adjuster.find(params[:id])
  end
 
  def new
    @adjuster = Adjuster.new
  end
 
  def edit
    @adjuster = Adjuster.find(params[:id])
  end
 
  def create
    @adjuster = Adjuster.new(adjuster_params)
 
    if @adjuster.save
      redirect_to adjusters_path
    else
      render 'new'
    end
  end
 
  def update
    @adjuster = Adjuster.find(params[:id])
 
    if @adjuster.update(adjuster_params)
      redirect_to adjusters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @adjuster = Adjuster.find(params[:id])
    @adjuster.destroy
    redirect_to adjusters_path
  end

 
  private
    def adjuster_params
      params.require(:adjuster).permit(:firstName, :lastName, :licenseNumber, :AdjusterType)
    end
end