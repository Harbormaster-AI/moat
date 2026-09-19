class InsuredObjectsController < ApplicationController
  def index
    @insuredObjects = InsuredObject.all
  end
 
  def show
    @insuredObject = InsuredObject.find(params[:id])
  end
 
  def new
    @insuredObject = InsuredObject.new
  end
 
  def edit
    @insuredObject = InsuredObject.find(params[:id])
  end
 
  def create
    @insuredObject = InsuredObject.new(insuredObject_params)
 
    if @insuredObject.save
      redirect_to insuredObjects_path
    else
      render 'new'
    end
  end
 
  def update
    @insuredObject = InsuredObject.find(params[:id])
 
    if @insuredObject.update(insuredObject_params)
      redirect_to insuredObjects_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @insuredObject = InsuredObject.find(params[:id])
    @insuredObject.destroy
    redirect_to insuredObjects_path
  end

 
  private
    def insuredObject_params
      params.require(:insuredObject).permit(:description, :serialOrId, :primaryAddress, :ObjectType)
    end
end