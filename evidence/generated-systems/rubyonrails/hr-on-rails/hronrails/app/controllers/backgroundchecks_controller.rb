class BackgroundChecksController < ApplicationController
  def index
    @backgroundChecks = BackgroundCheck.all
  end
 
  def show
    @backgroundCheck = BackgroundCheck.find(params[:id])
  end
 
  def new
    @backgroundCheck = BackgroundCheck.new
  end
 
  def edit
    @backgroundCheck = BackgroundCheck.find(params[:id])
  end
 
  def create
    @backgroundCheck = BackgroundCheck.new(backgroundCheck_params)
 
    if @backgroundCheck.save
      redirect_to backgroundChecks_path
    else
      render 'new'
    end
  end
 
  def update
    @backgroundCheck = BackgroundCheck.find(params[:id])
 
    if @backgroundCheck.update(backgroundCheck_params)
      redirect_to backgroundChecks_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @backgroundCheck = BackgroundCheck.find(params[:id])
    @backgroundCheck.destroy
    redirect_to backgroundChecks_path
  end

 
  private
    def backgroundCheck_params
      params.require(:backgroundCheck).permit(:checkNumber, :provider, :completedDate, :Status)
    end
end