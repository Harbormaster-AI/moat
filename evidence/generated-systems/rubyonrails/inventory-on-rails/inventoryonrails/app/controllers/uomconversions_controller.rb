class UoMConversionsController < ApplicationController
  def index
    @uoMConversions = UoMConversion.all
  end
 
  def show
    @uoMConversion = UoMConversion.find(params[:id])
  end
 
  def new
    @uoMConversion = UoMConversion.new
  end
 
  def edit
    @uoMConversion = UoMConversion.find(params[:id])
  end
 
  def create
    @uoMConversion = UoMConversion.new(uoMConversion_params)
 
    if @uoMConversion.save
      redirect_to uoMConversions_path
    else
      render 'new'
    end
  end
 
  def update
    @uoMConversion = UoMConversion.find(params[:id])
 
    if @uoMConversion.update(uoMConversion_params)
      redirect_to uoMConversions_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @uoMConversion = UoMConversion.find(params[:id])
    @uoMConversion.destroy
    redirect_to uoMConversions_path
  end

 
  private
    def uoMConversion_params
      params.require(:uoMConversion).permit(:factor, :precision, :FromUnit, :ToUnit)
    end
end