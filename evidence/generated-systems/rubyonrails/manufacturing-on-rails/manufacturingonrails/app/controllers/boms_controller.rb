class BOMsController < ApplicationController
  def index
    @bOMs = BOM.all
  end
 
  def show
    @bOM = BOM.find(params[:id])
  end
 
  def new
    @bOM = BOM.new
  end
 
  def edit
    @bOM = BOM.find(params[:id])
  end
 
  def create
    @bOM = BOM.new(bOM_params)
 
    if @bOM.save
      redirect_to bOMs_path
    else
      render 'new'
    end
  end
 
  def update
    @bOM = BOM.find(params[:id])
 
    if @bOM.update(bOM_params)
      redirect_to bOMs_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @bOM = BOM.find(params[:id])
    @bOM.destroy
    redirect_to bOMs_path
  end

 
  private
    def bOM_params
      params.require(:bOM).permit(:bomNumber, :revision, :effectivityStart, :effectivityEnd, :Status)
    end
end