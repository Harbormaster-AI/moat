class RegulationsController < ApplicationController
  def index
    @regulations = Regulation.all
  end
 
  def show
    @regulation = Regulation.find(params[:id])
  end
 
  def new
    @regulation = Regulation.new
  end
 
  def edit
    @regulation = Regulation.find(params[:id])
  end
 
  def create
    @regulation = Regulation.new(regulation_params)
 
    if @regulation.save
      redirect_to regulations_path
    else
      render 'new'
    end
  end
 
  def update
    @regulation = Regulation.find(params[:id])
 
    if @regulation.update(regulation_params)
      redirect_to regulations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @regulation = Regulation.find(params[:id])
    @regulation.destroy
    redirect_to regulations_path
  end

 
  private
    def regulation_params
      params.require(:regulation).permit(:name, :citation, :jurisdiction, :publicationUrl)
    end
end