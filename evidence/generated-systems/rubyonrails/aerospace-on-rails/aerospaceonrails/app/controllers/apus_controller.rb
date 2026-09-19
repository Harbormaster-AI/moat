class APUsController < ApplicationController
  def index
    @aPUs = APU.all
  end
 
  def show
    @aPU = APU.find(params[:id])
  end
 
  def new
    @aPU = APU.new
  end
 
  def edit
    @aPU = APU.find(params[:id])
  end
 
  def create
    @aPU = APU.new(aPU_params)
 
    if @aPU.save
      redirect_to aPUs_path
    else
      render 'new'
    end
  end
 
  def update
    @aPU = APU.find(params[:id])
 
    if @aPU.update(aPU_params)
      redirect_to aPUs_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @aPU = APU.find(params[:id])
    @aPU.destroy
    redirect_to aPUs_path
  end

 
  private
    def aPU_params
      params.require(:aPU).permit(:model)
    end
end